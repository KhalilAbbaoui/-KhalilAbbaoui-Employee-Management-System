package models

import (
    "go.mongodb.org/mongo-driver/bson/primitive"
    "time"
)

// Employee represents the structure of an employee's data
type Employee struct {
    ID         primitive.ObjectID `json:"id" bson:"_id,omitempty"` //MongoDB's ObjectID
    FirstName  string             `json:"firstName" binding:"required"`
    LastName   string             `json:"lastName" binding:"required"`
    Email      string             `json:"email" binding:"required,email" bson:"email"`
    Phone      string             `json:"phone" binding:"required"`
    Position   string             `json:"position" binding:"required"`
    Department string             `json:"department" binding:"required"`
    HireDate   time.Time          `json:"hireDate" binding:"required"`
   // Password   string             `json:"password"`
}
