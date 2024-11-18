angular.module('employeeApp')
.constant('AppConfig', {
    apiUrl: process.env.API_URL || 'http://localhost:8080/api',
    pageSize: 10,
    dateFormat: 'yyyy-MM-dd',
    toastDuration: 3000
})
.factory('HttpInterceptor', function($q, $injector) {
    return {
        request: function(config) {
            const token = localStorage.getItem('jwt_token');
            if (token) {
                config.headers.Authorization = 'Bearer ' + token;
            }
            return config;
        },
        responseError: function(rejection) {
            const $location = $injector.get('$location');
            
            if (rejection.status === 401) {
                localStorage.removeItem('jwt_token');
                $location.path('/login');
            }
            
            return $q.reject(rejection);
        }
    };
})
.config(function($httpProvider) {
    $httpProvider.interceptors.push('HttpInterceptor');
});
