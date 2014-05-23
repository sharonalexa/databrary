module.directive('volumeEditPublicationsForm', [
	'pageService', function (page) {
		var link = function ($scope) {
			var form = $scope.volumeEditPublicationsForm;

			form.data = {};
			form.volume = undefined;
			var backup = {};

			form.saveFn = undefined;
			form.successFn = undefined;
			form.errorFn = undefined;
			form.resetFn = undefined;
			form.cancelFn = undefined;

			//

			form.init = function (data, volume) {
				form.data = data;
				form.volume = form.volume || volume;
				backup = $.extend(true, {}, data);
			};

			//

			form.save = function () {
				if (angular.isFunction(form.saveFn)) {
					form.saveFn(form);
				}

				page.models.Volume.save(form.data,
					function (res) {
						form.messages.add({
							type: 'green',
							countdown: 3000,
							body: page.constants.message('volume.edit.publications.success'),
						});

						//update backup so a future revert goes to current state, not pageload state
						backup = $.extend(true, {}, form.data);

						if (angular.isFunction(form.successFn)) {
							form.successFn(form, res);
						}

						form.$setPristine();
						page.models.Volume.$cache.removeAll();
					}, function (res) {
						form.messages.addError({
							body: page.constants.message('volume.edit.publications.error'),
							report: res
						});

						if (angular.isFunction(form.errorFn)) {
							form.errorFn(form, res);
						}
					});
			};

			form.reset = function() { //reset to last saved state
				if (angular.isFunction(form.resetFn))
					form.resetFn(form);

				form.data = $.extend(true, {}, backup);
				form.$setPristine();

				if(form.repeater)
					form.repeater.repeats = form.data.citation;
			};

			//

			form.autoDOI = function (target) {
				if (!target.url || target.head) {
					return;
				}

				var doi = page.constants.data.regex.doi.exec(target.url);

				if (!doi || !doi[1]) {
					return;
				}

				if (!target.name) {
					page.models.CrossCite
						.json(doi[1])
						.then(function (res) {
							if (!res.title) {
								form.messages.add({
									type: 'red',
									countdown: 3000,
									body: page.constants.message('volume.edit.autodoi.name.error'),
								});
							} else {
								target.head = res.title;
								target.url = doi[1];

								form.messages.add({
									type: 'green',
									countdown: 3000,
									body: page.constants.message('volume.edit.autodoi.name.success'),
								});
							}
						}, function (res) {
							form.messages.add({
								type: 'red',
								countdown: 3000,
								body: page.constants.message('volume.edit.autodoi.name.error'),
							});
						});
				}
			};

			//

			var changeFn = function () {
				form.$setDirty();
			};

			form.retrieveRepeater = function (repeater) {
				form.repeater = repeater;
				form.repeater.autoDOI = form.autoDOI;
				form.repeater.repeats = form.data.citation || [];
				form.repeater.addFn = changeFn;
				form.repeater.removeFn = changeFn;
			};

			//

			page.events.talk('volumeEditPublicationsForm-init', form, $scope);
		};

		//

		return {
			restrict: 'E',
			templateUrl: 'volumeEditPublicationsForm.html',
			scope: false,
			replace: true,
			link: link
		};
	}
]);
