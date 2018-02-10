'use strict';

angular.module('myApp.createblog', ['ngRoute'])

.config(['$routeProvider', function($routeProvider) {
	$routeProvider.when('/createblog', {
		templateUrl: 'createblog/createblog.html',
		controller: 'CreateBlogCtrl'
	});
}])

.controller('CreateBlogCtrl', ['$scope', '$mdDialog', '$http', function($scope,
	$mdDialog, $http) {

	$scope.blog = {};
	$scope.showHints = {};

	function isValuedObject(obj) {

		if (angular.isObject(obj) && Object.keys(obj).length > 0) {
			return true;
		}

		return false;
	}

	function doSave(ev) {

		$http.post("/blog/new", $scope.blog).
		then(function success(response) {

			$mdDialog.show(
				$mdDialog.alert()
				.parent(angular.element(document.querySelector('#popupContainer')))
				.clickOutsideToClose(true)
				.title('提示')
				.textContent('创建日志成功.')
				.ariaLabel('创建日志成功')
				.targetEvent(ev)
				.ok('关闭')
			);

		}, function failure(response) {

			$mdDialog.show(
				$mdDialog.alert()
				.parent(angular.element(document.querySelector('#popupContainer')))
				.clickOutsideToClose(true)
				.title('提示')
				.textContent('创建日志失败：' + response.error)
				.ariaLabel('创建日志失败')
				.targetEvent(ev)
				.ok('关闭')
				.targetEvent(ev)
			);
		})
	}

	$scope.save = function(ev) {

		$scope.showHints = {
			subject: isValuedObject($scope.blogForm.subject.$error),
			author: isValuedObject($scope.blogForm.author.$error),
			blog: isValuedObject($scope.blogForm.blog.$error)
		};

		if (Object.keys($scope.blogForm.$error).length == 0) {
			doSave(ev)
		};

	}

}]);
