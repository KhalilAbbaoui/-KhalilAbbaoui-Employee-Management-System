angular.module('employeeApp')
.controller('EmployeeFormController', function($scope, $routeParams, $location, EmployeeService) {
    $scope.employee = {};
    $scope.isEditing = !!$routeParams.id;

    if ($scope.isEditing) {
        EmployeeService.getById($routeParams.id)
            .then(function(response) {
                $scope.employee = response.data;
            })
            .catch(function(error) {
                alert('Error loading employee: ' + error.message);
            });
    }

    $scope.saveEmployee = function() {
        var promise = $scope.isEditing ?
            EmployeeService.update($routeParams.id, $scope.employee) :
            EmployeeService.create($scope.employee);

        promise
            .then(function() {
                $location.path('/');
            })
            .catch(function(error) {
                alert('Error saving employee: ' + error.message);
            });
    };
});