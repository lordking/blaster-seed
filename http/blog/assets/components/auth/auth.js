'use strict';

angular.module('myApp.auth', ['ngCookies', 'base64'])

.provider('$auth', AuthProvider);

function AuthProvider() {

  this.$get = ['$cookies', '$base64', '$http',  function($cookies, $base64, $http) {
    return new Auth($cookies, $base64, $http)
  }]

}

function Auth($cookies, $base64, $http) {

  /**
   * 登录成功后的处理
   * @param  {[type]} username [description]
   * @param  {[type]} password [description]
   * @param  {[type]} token    [description]
   * @return {[type]}          [description]
   */
  this.afterSuccess = function(username, password, token, expireTime) {

    $cookies.putObject("userinfo", {
      "username": username,
      "password": password,
      "token": token,
      "expireTime": expireTime
    });

    window.location.href = '/blog.html';
  }

  /**
   * 运行过程中，认证失败后的处理
   * @return {[type]} [description]
   */
  this.logout = function() {

    $cookies.remove("userinfo");
    window.location.href = '/login.html';
  };

  /**
   * 获取用户信息
   * @return {[type]} [description]
   */
  this.userinfo = function() {
    return $cookies.getObject("userinfo");
  }

  /**
   * 使用登录成功后保存在cookie的内容，自动进行登录授权
   * @return {[type]} [description]
   */
  this.authorize = function() {

    var userinfo = $cookies.getObject("userinfo");

    if (userinfo) {
      var tokens = $base64.encode(userinfo.token + ":123");
      $http.defaults.headers.common.Authorization = 'Basic ' + tokens;

    } else {
      window.location.href = '/login.html'
    }

  }

  this.isAuthorized = function() {

    var userinfo = $cookies.getObject("userinfo");

    if (!userinfo.token || userinfo.token != "") {
      return true;
    }

    return false;
  }


}
