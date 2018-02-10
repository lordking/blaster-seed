'use strict';

// Declare app level module which depends on views, and components
angular.module('myApp', [
	'ngRoute',
	'ngMaterial',
	'myApp.version',
	'myApp.auth'
])

.config(function($mdThemingProvider, $mdIconProvider) {

	$mdThemingProvider.theme('default')
		.primaryPalette('teal')
		.accentPalette('red');
})

.controller('loginCtrl', ['$scope', '$http', '$auth', function($scope, $http, $auth) {

	$scope.user = {};
	$scope.showHints = {};
	$scope.loginForm = {};

	function isValuedObject(obj) {

		if (angular.isObject(obj) && Object.keys(obj).length > 0) {
			return true;
		}

		return false;
	}

	function doLogin() {

		var username = $scope.user.username;
		var password = $scope.user.password;
		var headers = {
			authorization: "Basic " + btoa(username + ":" + password)
		};

		$http.post("/user/login?random=" + Math.random(), $scope.user, {
				headers: headers
			}). //加入随机数，保持最新
		then(function success(response) {

			var token = response.data.result.token;
			var expireTime = response.data.result.expireTime;

			if (!token || token == "") {
				$scope.loginResult = "认证失败。失败原因: 服务器内部错误";
				return
			}

			$auth.afterSuccess(username, password, token, expireTime);

		}, function failure(response) {

			$scope.loginResult = "认证失败。失败原因:" + response.data.error

		});
	}

	$scope.login = function(ev) {

		$scope.showHints = {
			username: isValuedObject($scope.loginForm.username.$error),
			password: isValuedObject($scope.loginForm.password.$error),
		};

		if (Object.keys($scope.loginForm.$error).length == 0) {
			doLogin()
		};
	}

}]);
