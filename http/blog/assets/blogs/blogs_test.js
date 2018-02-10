'use strict';

describe('myApp.blogs module', function() {

  beforeEach(module('myApp.blogs'));

  describe('blogs controller', function(){

    it('should ....', inject(function($controller) {
      //spec body
      var view2Ctrl = $controller('BlogsCtrl');
      expect(view2Ctrl).toBeDefined();
    }));

  });
});