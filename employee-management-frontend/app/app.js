// Declare the main module with the ngRoute dependency
var app = angular.module('employeeApp', ['ngRoute']);

// Configure routing
app.config(function($routeProvider) {
    $routeProvider
        .when('/employees', {
            templateUrl: 'views/employee-list.html',
            controller: 'EmployeeListController'
        })
        .when('/employees/edit/:id', {
            templateUrl: 'views/employee-form.html',
            controller: 'EmployeeFormController'
        })
        .otherwise({
            redirectTo: '/employees'
        });
});

// Service to manage employee data
app.service('EmployeeService', function($http) {
    var apiUrl = '/api/employees'; // Replace with your actual API endpoint

    // Get all employees
    this.getAll = function() {
        return $http.get(apiUrl);
    };

    // Get a single employee by ID
    this.getById = function(id) {
        return $http.get(apiUrl + '/' + id);
    };

    // Create a new employee
    this.create = function(employee) {
        return $http.post(apiUrl, employee);
    };

    // Update an existing employee
    this.update = function(id, employee) {
        return $http.put(apiUrl + '/' + id, employee);
    };

    // Delete an employee
    this.delete = function(id) {
        return $http.delete(apiUrl + '/' + id);
    };
});

// Controller for displaying the list of employees
app.controller('EmployeeListController', function($scope, EmployeeService) {
    // Fetch all employees
    EmployeeService.getAll().then(function(response) {
        $scope.employees = response.data;
    });

    // Search functionality
    $scope.filterEmployees = function(employee) {
        return (!$scope.searchText || employee.firstName.toLowerCase().includes($scope.searchText.toLowerCase()) || 
                employee.lastName.toLowerCase().includes($scope.searchText.toLowerCase()));
    };

    // Delete an employee
    $scope.deleteEmployee = function(id) {
        if (confirm('Are you sure you want to delete this employee?')) {
            EmployeeService.delete(id).then(function() {
                // Refresh employee list after deletion
                $scope.employees = $scope.employees.filter(function(employee) {
                    return employee.id !== id;
                });
            }).catch(function(error) {
                alert('Error deleting employee: ' + error.message);
            });
        }
    };
});

// Controller for creating or editing an employee
app.controller('EmployeeFormController', function($scope, $routeParams, EmployeeService, $location) {
    var employeeId = $routeParams.id;

    if (employeeId) {
        // If editing, fetch the employee data
        EmployeeService.getById(employeeId).then(function(response) {
            $scope.employee = response.data;
        });
    } else {
        // Otherwise, initialize a new employee
        $scope.employee = {};
    }

    // Save or update the employee
    $scope.saveEmployee = function() {
        if (employeeId) {
            // If there's an ID, update the existing employee
            EmployeeService.update(employeeId, $scope.employee).then(function() {
                $location.path('/employees'); // Redirect to employee list after update
            });
        } else {
            // Otherwise, create a new employee
            EmployeeService.create($scope.employee).then(function() {
                $location.path('/employees'); // Redirect to employee list after creation
            });
        }
    };
});

