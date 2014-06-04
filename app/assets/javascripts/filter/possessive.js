module.filter('possessive', [
	'pageService', function (page) {
		return function (key, party, name) {
			var replace;

			if (angular.isString(party)) {
				replace = party + "'s";
			}
			else if (page.auth.user.id == party.id) {
				replace = 'my';
			}
			else {
				replace = (name ? name : party.name) + "'s";
			}

			return page.constants.message(key, replace);
		};
	}
]);
