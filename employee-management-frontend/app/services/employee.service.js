angular.module('employeeApp')
.service('EmployeeService', function($http) {
    const apiUrl = 'your-api-url-here'; // Update with actual API URL

    this.getAll = function() {
        return $http.get(apiUrl + '/employees');
    };

    this.get = function(id) {
        return $http.get(apiUrl + '/employees/' + id);
    };

    this.create = function(employee) {
        return $http.post(apiUrl + '/employees', employee);
    };

    this.update = function(id, employee) {
        return $http.put(apiUrl + '/employees/' + id, employee);
    };

    this.delete = function(id) {
        return $http.delete(apiUrl + '/employees/' + id);
    };
});
