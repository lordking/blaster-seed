'use strict';

angular.module('myApp.blogs', ['ngRoute'])

.config(['$routeProvider', function($routeProvider) {
  $routeProvider.when('/blogs', {
    templateUrl: 'blogs/blogs.html',
    controller: 'BlogsCtrl'
  });
}])

.controller('BlogsCtrl', ['$scope', '$mdDialog', '$http', '$log', '$auth', function($scope,
  $mdDialog, $http, $log, $auth) {

  if (!$auth.isAuthorized()) {
    $scope.totalDesc = "认证失败";
    $auth.logout();
    return
  }

  var start = 0;
  $scope.showButtons = {};
  if (!$scope.blogs) {
    loadBlogs(start);
  };

  function loadBlogs(start) {

    $http.get("/blog/" + start + "/10?random=" + Math.random()). //加入随机数，保持最新
    then(function success(response) {

      if (!response.data.result) {
        $scope.totalDesc = "当前无记录";
        $scope.blogs = [];

      } else {
        var result = response.data.result;

        if (result.length > 0) {
          $scope.totalDesc = '当前' + result.length + '记录';
        } else {
          $scope.totalDesc = "当前无记录";
        }

        $scope.blogs = result;
      }

      if (start == 0) {
        $scope.showButtons.prePage = false;
      } else {
        $scope.showButtons.prePage = true;
      }

      if ($scope.blogs < 10) {
        $scope.showButtons.nextPage = false;
      } else {
        $scope.showButtons.nextPage = true;
      }

    }, function failure(response) {

      $scope.totalDesc = "导入日志失败";

    });
  }

  function doDelete(item) {

    $http.delete("/blog/delete/" + item.id, {
      cache: false
    }).
    then(function success(response) {

      var result = response.data;
      $scope.status = '你已经成功删除这条记录';
      loadBlogs(start);

    }, function failure(response) {

      $scope.status = '删除失败：' + response.data.error;
      loadBlogs(start);
    });
  }

  function doUpdate(item) {

    $http.put("/blog/update/" + item.id, item).
    then(function success(response) {

      var result = response.data;
      $scope.status = '你已经成功更新这条记录';
      loadBlogs(start);

    }, function failure(response) {

      $scope.status = '更新失败：' + response.data.error;
      loadBlogs(start);
    });
  }

  $scope.prePage = function() {
    start -= 10;
    loadBlogs(start);
  }

  $scope.nextPage = function() {
    start += 10;
    loadBlogs(start);
  }

  $scope.delete = function(item, ev) {

    var confirm = $mdDialog.confirm()
      .title('警告')
      .textContent('是否删除这条记录')
      .ariaLabel('Delete item')
      .targetEvent(ev)
      .ok('删除')
      .cancel('取消');

    $mdDialog.show(confirm).then(function() {

      doDelete(item);

    }, function() {
      $scope.status = '你已经取消删除这条记录';
    });
  }

  $scope.edit = function(item, ev) {

    $mdDialog.show({
        controller: UpdateCtrl,
        templateUrl: 'blogs/update.tmpl.html?random=' + Math.random(), //加入随机数，保持最新
        parent: angular.element(document.body),
        targetEvent: ev,
        clickOutsideToClose: true,
        fullscreen: true,
        locals: {
          item: angular.copy(item)
        }
      })
      .then(function(answer) {

        doUpdate(answer)

      }, function() {
        $scope.status = '你已经放弃修改这条记录';
      });
  }

  function UpdateCtrl($scope, $mdDialog, item) {
    $scope.blog = item;

    $scope.cancel = function() {
      $mdDialog.cancel();
    }

    $scope.answer = function() {
      $mdDialog.hide($scope.blog);
    }
  }

}]);
