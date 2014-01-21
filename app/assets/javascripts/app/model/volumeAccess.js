define(['app/config/module'], function (module) {
	'use strict';

	module.factory('VolumeAccess', ['$rootScope', '$resource', '$route', function ($rootScope, $resource, $route) {
		return $resource('/api/volume/:id/access/:partyId', {
			id: function () {
				return $route.current.params.id || undefined;
			},
			partyId: function () {
				return $route.current.params.partyId || undefined;
			}
		}, {
			'query': {method: 'GET', isArray: false}
		});
	}]);
});
