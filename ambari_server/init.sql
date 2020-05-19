set password for 'root'@'localhost' = password('hjc123456.com');
grant all privileges on *.* to 'root'@'%' identified by 'hjc123456.com';
CREATE DATABASE ambari;  
use ambari;  
CREATE USER 'ambari'@'%' IDENTIFIED BY 'ambarizk123';  
GRANT ALL PRIVILEGES ON *.* TO 'ambari'@'%';  
CREATE USER 'ambari'@'localhost' IDENTIFIED BY 'ambarizk123';  
GRANT ALL PRIVILEGES ON *.* TO 'ambari'@'localhost';  
CREATE USER 'ambari'@'master' IDENTIFIED BY 'ambarizk123';  
GRANT ALL PRIVILEGES ON *.* TO 'ambari'@'master';  
FLUSH PRIVILEGES;  
source /var/lib/ambari-server/resources/Ambari-DDL-MySQL-CREATE.sql  
show tables;  
use mysql;  
select Host User Password from user where user='ambari';  
CREATE DATABASE hive;  
use hive;  
CREATE USER 'hive'@'%' IDENTIFIED BY 'hive';  
GRANT ALL PRIVILEGES ON *.* TO 'hive'@'%';  
CREATE USER 'hive'@'localhost' IDENTIFIED BY 'hive';  
GRANT ALL PRIVILEGES ON *.* TO 'hive'@'localhost';  
CREATE USER 'hive'@'master' IDENTIFIED BY 'hive';  
GRANT ALL PRIVILEGES ON *.* TO 'hive'@'master';  
FLUSH PRIVILEGES;  
CREATE DATABASE oozie;  
use oozie;  
CREATE USER 'oozie'@'%' IDENTIFIED BY 'oozie';  
GRANT ALL PRIVILEGES ON *.* TO 'oozie'@'%';  
CREATE USER 'oozie'@'localhost' IDENTIFIED BY 'oozie';  
GRANT ALL PRIVILEGES ON *.* TO 'oozie'@'localhost';  
CREATE USER 'oozie'@'master' IDENTIFIED BY 'oozie';  
GRANT ALL PRIVILEGES ON *.* TO 'oozie'@'master';  
FLUSH PRIVILEGES;  
