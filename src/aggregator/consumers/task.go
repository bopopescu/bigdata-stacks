package consumers

import (
	"strconv"
	"time"

	"github.com/Shopify/sarama"
	saramaC "github.com/bsm/sarama-cluster"
	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/ovh/metronome/src/metronome/kafka"
	"github.com/ovh/metronome/src/metronome/models"
	"github.com/ovh/metronome/src/metronome/pg"
	"github.com/ovh/metronome/src/metronome/redis"
)

// TaskConsumer consumed tasks messages from a Kafka topic to maintain the tasks database.
type TaskConsumer struct {
	consumer                 *saramaC.Consumer
	doneTasks                int
	lastCommit               time.Time
	taskCounter              *prometheus.CounterVec
	taskUnprocessableCounter *prometheus.CounterVec
	taskProcessedCounter     *prometheus.CounterVec
	taskPublishErrorCounter  *prometheus.CounterVec
}

// NewTaskConsumer returns a new task consumer.
func NewTaskConsumer() (*TaskConsumer, error) {
	brokers := viper.GetStringSlice("kafka.brokers")

	config := saramaC.NewConfig()
	config.Config = *kafka.NewConfig()
	config.ClientID = "metronome-aggregator"
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	consumer, err := saramaC.NewConsumer(brokers, kafka.GroupAggregators(), []string{kafka.TopicTasks()}, config)
	if err != nil {
		return nil, err
	}

	tc := &TaskConsumer{
		consumer:   consumer,
		doneTasks:  0,
		lastCommit: time.Now(),
	}

	// metrics
	tc.taskCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "metronome",
		Subsystem: "aggregator",
		Name:      "tasks",
		Help:      "Number of tasks processed.",
	},
		[]string{"partition"})
	prometheus.MustRegister(tc.taskCounter)
	tc.taskUnprocessableCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "metronome",
		Subsystem: "aggregator",
		Name:      "tasks_unprocessable",
		Help:      "Number of unprocessable tasks.",
	},
		[]string{"partition"})
	prometheus.MustRegister(tc.taskUnprocessableCounter)
	tc.taskProcessedCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "metronome",
		Subsystem: "aggregator",
		Name:      "tasks_processeed",
		Help:      "Number of processeed tasks.",
	},
		[]string{"partition"})
	prometheus.MustRegister(tc.taskProcessedCounter)
	tc.taskPublishErrorCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "metronome",
		Subsystem: "aggregator",
		Name:      "tasks_publish_error",
		Help:      "Number of tasks publish error.",
	},
		[]string{"partition"})
	prometheus.MustRegister(tc.taskPublishErrorCounter)

	// Consume Kafka Tasks
	go func() {
		for {
			select {
			case msg, ok := <-consumer.Messages():
				if !ok { // shuting down
					return
				}
				if err := tc.handleMsg(msg); err != nil {
					log.WithError(err).Warn("Could not handle the task")
					continue
				}
			}
		}
	}()

	return tc, nil
}

// Close the consumer.
func (tc *TaskConsumer) Close() error {
	return tc.consumer.Close()
}

// Handle message from Kafka.
// Apply updates to the database.
func (tc *TaskConsumer) handleMsg(msg *sarama.ConsumerMessage) error {
	tc.taskCounter.WithLabelValues(strconv.Itoa(int(msg.Partition))).Inc()
	var t models.Task
	if err := t.FromKafka(msg); err != nil {
		tc.taskUnprocessableCounter.WithLabelValues(strconv.Itoa(int(msg.Partition))).Inc()
		return err
	}

	db := pg.DB()

	if t.Schedule == "" {
		log.Infof("DELETE task: %s", t.GUID)

		_, err := db.Model(&t).Delete()
		if err != nil {
			return err
		}
	} else {
		log.Infof("UPSERT task: %s", t.GUID)

		_, err := db.Model(&t).OnConflict("(guid) DO UPDATE").
			Set("name = ?name").
			Set("urn = ?urn").
			Set("schedule = ?schedule").
			Set("payload = ?payload").
			Set("id = ?id").
			Insert()
		if err != nil {
			return err
		}
	}
	tc.taskProcessedCounter.WithLabelValues(strconv.Itoa(int(msg.Partition))).Inc()

	body, err := t.ToJSON()
	if err != nil {
		return err
	}

	if err = redis.DB().PublishTopic(t.UserID, "task", string(body)).Err(); err != nil {
		tc.taskPublishErrorCounter.WithLabelValues(strconv.Itoa(int(msg.Partition))).Inc()
		return err
	}

	tc.consumer.MarkOffset(msg, "aggregated")
	tc.doneTasks++
	if tc.doneTasks >= 100 || time.Now().After(tc.lastCommit.Add(time.Duration(time.Second*10))) {
		// If more than 10 seconds since last offset commit ORmore than 100 messages pending
		if err = tc.consumer.CommitOffsets(); err != nil {
			return err
		}

		tc.doneTasks = 0
		tc.lastCommit = time.Now()
	}

	return nil
}
