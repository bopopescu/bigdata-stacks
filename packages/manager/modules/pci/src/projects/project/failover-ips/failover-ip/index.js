import angular from 'angular';
import '@uirouter/angularjs';

import edit from './edit';

const moduleName = 'ovhManagerPciProjectFailoverIp';

angular.module(moduleName, [edit, 'ui.router']);

export default moduleName;
