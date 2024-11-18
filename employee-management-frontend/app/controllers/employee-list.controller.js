angular.module('employeeApp')
.controller('EmployeeListController', function($scope, EmployeeService) {
    $scope.employees = [];
    $scope.searchText = '';

    function loadEmployees() {
        EmployeeService.getAll()
            .then(function(response) {
                $scope.employees = response.data;
            })
            .catch(function(error) {
                alert('Error loading employees: ' + error.message);
            });
    }

    $scope.deleteEmployee = function(id) {
        if (confirm('Are you sure you want to delete this employee?')) {
            EmployeeService.delete(id)
                .then(function() {
                    loadEmployees();
                })
                .catch(function(error) {
                    alert('Error deleting employee: ' + error.message);
                });
        }
    };

    $scope.filterEmployees = function(employee) {
        var searchText = $scope.searchText.toLowerCase();
        return !searchText || 
               employee.firstName.toLowerCase().includes(searchText) ||
               employee.lastName.toLowerCase().includes(searchText) ||
               employee.department.toLowerCase().includes(searchText);
    };

    loadEmployees();
});