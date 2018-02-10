'use strict';

// Declare app level module which depends on views, and components
angular.module('myApp', [
  'ngRoute',
  'ngMaterial',
  'myApp.createblog',
  'myApp.blogs',
  'myApp.version',
  'myApp.auth'
])

.config(['$routeProvider', function($routeProvider) {
  $routeProvider.otherwise({
    redirectTo: '/blogs'
  });
}])

.config(['$mdThemingProvider', '$mdIconProvider', function($mdThemingProvider,
  $mdIconProvider) {
  $mdThemingProvider.theme('default')
    .primaryPalette('teal')
    .accentPalette('red');
}])

.controller('MainCtrl', ['$scope', '$auth', function($scope, $auth) {

  $scope.logout = function() {
    $auth.logout();
  }

}])

.run(['$auth', function($auth) {

  $auth.authorize();

}]);
