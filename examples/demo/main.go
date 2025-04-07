package main

import (
	"fmt"
	"os"
	"time"
)

// User represents a user in the system
type User struct {
	ID        int
	Name      string
	Email     string
	CreatedAt time.Time
}

// NewUser creates a new user
func NewUser(name, email string) *User {
	return &User{
		ID:        generateID(),
		Name:      name,
		Email:     email,
		CreatedAt: time.Now(),
	}
}

// generateID creates a simple random ID
func generateID() int {
	return int(time.Now().UnixNano() % 10000)
}

// SaveUser saves a user to a database (simulated)
func SaveUser(user *User) error {
	// Simulate database operation
	fmt.Printf("Saving user: %s\n", user.Name)
	return nil
}

// GetUserByID retrieves a user by ID
func GetUserByID(id int) (*User, error) {
	// Simulate database lookup
	if id < 0 {
		return nil, fmt.Errorf("invalid user ID: %d", id)
	}
	
	// Mock data
	return &User{
		ID:        id,
		Name:      "Example User",
		Email:     "user@example.com",
		CreatedAt: time.Now(),
	}, nil
}

func main() {
	// Create a new user
	user := NewUser("John Doe", "john@example.com")
	
	// Save the user
	if err := SaveUser(user); err != nil {
		fmt.Printf("Error saving user: %v\n", err)
		os.Exit(1)
	}
	
	// Retrieve the user
	retrievedUser, err := GetUserByID(user.ID)
	if err != nil {
		fmt.Printf("Error retrieving user: %v\n", err)
		os.Exit(1)
	}
	
	// Display user information
	fmt.Println("User Information:")
	fmt.Printf("ID: %d\n", retrievedUser.ID)
	fmt.Printf("Name: %s\n", retrievedUser.Name)
	fmt.Printf("Email: %s\n", retrievedUser.Email)
	fmt.Printf("Created: %s\n", retrievedUser.CreatedAt.Format(time.RFC3339))
} 