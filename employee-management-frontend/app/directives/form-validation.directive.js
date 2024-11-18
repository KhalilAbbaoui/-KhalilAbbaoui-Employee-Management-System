angular.module('employeeApp')
.directive('formValidation', function() {
    return {
        restrict: 'A',
        require: 'form',
        link: function(scope, element, attrs, formCtrl) {
            element.on('submit', function() {
                if (formCtrl.$invalid) {
                    Object.keys(formCtrl).forEach(function(key) {
                        if (key[0] !== '$') {
                            var control = formCtrl[key];
                            control.$setTouched();
                        }
                    });
                    return false;
                }
            });
        }
    };
});
