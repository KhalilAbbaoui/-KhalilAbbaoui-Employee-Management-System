// backend/tests/employee_test.go
package tests

import (
    "testing"
    "net/http"
    "net/http/httptest"
    "encoding/json"
    "bytes"
)

func TestEmployeeAPI(t *testing.T) {
    // Configuration du router pour les tests
    router := setupTestRouter()
    
    t.Run("GetEmployees", func(t *testing.T) {
        w := httptest.NewRecorder()
        req, _ := http.NewRequest("GET", "/api/employees", nil)
        router.ServeHTTP(w, req)
        
        if w.Code != http.StatusOK {
            t.Errorf("Expected status %d; got %d", http.StatusOK, w.Code)
        }
    })
    
    t.Run("CreateEmployee", func(t *testing.T) {
        employee := Employee{
            FirstName:  "John",
            LastName:   "Doe",
            Email:      "john@example.com",
            Phone:      "1234567890",
            Position:   "Developer",
            Department: "IT",
            HireDate:   time.Now(),
        }
        
        payload, _ := json.Marshal(employee)
        w := httptest.NewRecorder()
        req, _ := http.NewRequest("POST", "/api/employees", bytes.NewBuffer(payload))
        req.Header.Set("Content-Type", "application/json")
        router.ServeHTTP(w, req)
        
        if w.Code != http.StatusCreated {
            t.Errorf("Expected status %d; got %d", http.StatusCreated, w.Code)
        }
    })
}

// frontend/tests/employee.controller.test.js
describe('EmployeeListController', function() {
    var $controller, $scope, EmployeeService;
    
    beforeEach(module('employeeApp'));
    
    beforeEach(inject(function(_$controller_, _$rootScope_, _EmployeeService_) {
        $controller = _$controller_;
        $scope = _$rootScope_.$new();
        EmployeeService = _EmployeeService_;
        
        spyOn(EmployeeService, 'getAll').and.returnValue(Promise.resolve({
            data: [{
                id: 1,
                firstName: 'John',
                lastName: 'Doe',
                email: 'john@example.com'
            }]
        }));
    }));
    
    it('should load employees on initialization', function(done) {
        $controller('EmployeeListController', { $scope: $scope });
        
        setTimeout(function() {
            expect($scope.employees.length).toBe(1);
            expect($scope.employees[0].firstName).toBe('John');
            done();
        });
    });
});
