package controllers

import (
	"context"
	"employee-management-system/config"
	"employee-management-system/models"
	"employee-management-system/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"go.mongodb.org/mongo-driver/mongo"
)

// Login handles the user login and generates a JWT
func Login(c *gin.Context) {
    var creds models.Credentials
    if err := c.ShouldBindJSON(&creds); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }

    // Mock user for testing purposes (or create dynamically)
    mockUser := models.Employee{
        Email:    creds.Email,
        Password: creds.Password, // In real-world, you'd hash the password
        ID:       primitive.NewObjectID(),
    }

    // Simulate checking the password (without actual DB lookup)
    if creds.Password != "Test1234!" {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    // Generate token using the string representation of ObjectID
    token, err := utils.GenerateToken(mockUser.ID.Hex())
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": token})
}

// GetEmployees retrieves all employees
func GetEmployees(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := config.GetEmployeeCollection().Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching employees"})
		return
	}
	defer cursor.Close(ctx)

	var employees []models.Employee
	if err := cursor.All(ctx, &employees); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error decoding employees"})
		return
	}

	c.JSON(http.StatusOK, employees)
}

// GetEmployee retrieves a specific employee by ID
func GetEmployee(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var employee models.Employee
	err = config.GetEmployeeCollection().FindOne(ctx, bson.M{"_id": objID}).Decode(&employee)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		return
	}

	c.JSON(http.StatusOK, employee)
}

// ValidateEmail checks if the email is valid
func ValidateEmail(email string) bool {
	if email == "" {
		return false
	}
	re := `^[a-zA-Z0-9._%+-]+@[a-zAZ0-9.-]+\.[a-zA-Z]{2,}$`
	match, _ := regexp.MatchString(re, email)
	return match
}

// CreateEmployee creates a new employee
func CreateEmployee(c *gin.Context) {
	var employee models.Employee
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate email
	if !ValidateEmail(employee.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email address"})
		return
	}

	// Generate a new ObjectID for the employee
	employee.ID = primitive.NewObjectID()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Check if employee with the same email already exists
	var existingEmployee models.Employee
	err := config.GetEmployeeCollection().FindOne(ctx, bson.M{"email": employee.Email}).Decode(&existingEmployee)
	if err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Email already exists"})
		return
	}
	if err != mongo.ErrNoDocuments {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error checking if employee exists"})
		return
	}

	// Insert the employee without password handling
	_, err = config.GetEmployeeCollection().InsertOne(ctx, employee)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating employee"})
		return
	}

	c.JSON(http.StatusCreated, employee)
}


// UpdateEmployee updates an existing employee's information
// UpdateEmployee updates an existing employee's information without modifying password
func UpdateEmployee(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var employee models.Employee
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the provided email is valid
	if !ValidateEmail(employee.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email address"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Check if employee exists before updating
	var existingEmployee models.Employee
	err = config.GetEmployeeCollection().FindOne(ctx, bson.M{"_id": objID}).Decode(&existingEmployee)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		return
	}

	// Create an update document with fields that have been provided (excluding password)
	updateData := bson.M{
		"$set": bson.M{
			"firstName":  employee.FirstName,
			"lastName":   employee.LastName,
			"email":      employee.Email,
			"phone":      employee.Phone,
			"position":   employee.Position,
			"department": employee.Department,
			"hireDate":   employee.HireDate,
		},
	}

	// Perform the update operation
	_, err = config.GetEmployeeCollection().UpdateOne(
		ctx,
		bson.M{"_id": objID},
		updateData,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating employee"})
		return
	}

	// Respond with the updated employee details
	c.JSON(http.StatusOK, gin.H{"message": "Employee updated successfully", "employee": employee})
}

// DeleteEmployee deletes an employee
func DeleteEmployee(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := config.GetEmployeeCollection().DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting employee"})
		return
	}

	if result.DeletedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Employee deleted successfully"})
}
