package models

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestUser_Struct(t *testing.T) {
	user := User{
		ID:        1,
		Username:  "testuser",
		Email:     "test@example.com",
		Password:  "hashedpassword",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	assert.Equal(t, 1, user.ID)
	assert.Equal(t, "testuser", user.Username)
	assert.Equal(t, "test@example.com", user.Email)
	assert.Equal(t, "hashedpassword", user.Password)
}

func TestUserRequest_Validation(t *testing.T) {
	req := UserRequest{
		Username: "testuser",
		Email:    "test@example.com",
		Password: "password123",
	}

	assert.Equal(t, "testuser", req.Username)
	assert.Equal(t, "test@example.com", req.Email)
	assert.Equal(t, "password123", req.Password)
}