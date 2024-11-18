angular.module('employeeApp')
.service('AuthService', function($http, AppConfig) {
    return {
        login: function(credentials) {
            return $http.post(AppConfig.apiUrl + '/auth/login', credentials)
                .then(function(response) {
                    // Save the token in localStorage
                    localStorage.setItem('jwt_token', response.data.token);
                    console.log("Token saved: ", response.data.token);  // Add this for debugging
                    return response.data;
                });
        },        
        
        logout: function() {
            // Remove token from localStorage
            localStorage.removeItem('jwt_token');
        },
        
        isAuthenticated: function() {
            // Check if there's a valid token in localStorage
            return !!localStorage.getItem('jwt_token');
        },
        
        getToken: function() {
            return localStorage.getItem('jwt_token');
        }
    };
});
