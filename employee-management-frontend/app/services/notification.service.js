angular.module('employeeApp')
.service('NotificationService', function($timeout, AppConfig) {
    var notifications = [];
    
    return {
        show: function(message, type) {
            var notification = {
                message: message,
                type: type || 'info',
                id: Date.now()
            };
            
            notifications.push(notification);
            
            $timeout(function() {
                var index = notifications.indexOf(notification);
                if (index > -1) {
                    notifications.splice(index, 1);
                }
            }, AppConfig.toastDuration);
        },
        
        success: function(message) {
            this.show(message, 'success');
        },
        
        error: function(message) {
            this.show(message, 'error');
        },
        
        getNotifications: function() {
            return notifications;
        }
    };
});
