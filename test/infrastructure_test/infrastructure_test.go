package infrastructuretest

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zaahidali/task_manager_api/Infrastructure"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestGenerateToken(t *testing.T) {
	infra := &Infrastructure.Infrastructure{}

	// Test data
	userName := "testuser"
	userId := primitive.NewObjectID()
	role := "user"

	token, err := infra.GenerateToken(userName, userId, role)

	// Ensure no error and token is not empty
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	// Validate token (token should be valid)
	valid, err := infra.ValidateToken(token)
	assert.NoError(t, err)
	assert.True(t, valid)
}

func TestValidateToken(t *testing.T) {
	infra := &Infrastructure.Infrastructure{}

	// Generate a valid token
	userName := "testuser"
	userId := primitive.NewObjectID()
	role := "user"
	token, err := infra.GenerateToken(userName, userId, role)
	assert.NoError(t, err)

	// Validate the token (should be valid)
	valid, err := infra.ValidateToken(token)
	assert.NoError(t, err)
	assert.True(t, valid)

	// Test with an invalid token
	invalidToken := token + "invalid"
	valid, err = infra.ValidateToken(invalidToken)
	assert.Error(t, err)
	assert.False(t, valid)
}

func TestHashPassword(t *testing.T) {
	infra := &Infrastructure.Infrastructure{}

	password := "testpassword"
	hashedPassword, err := infra.HashPassword(password)

	// Ensure no error and hashed password is not empty
	assert.NoError(t, err)
	assert.NotEmpty(t, hashedPassword)

	// Compare hashed password with the original password
	err = infra.ComparePasswords(hashedPassword, password)
	assert.NoError(t, err)
}

func TestComparePasswords(t *testing.T) {
	infra := &Infrastructure.Infrastructure{}

	password := "testpassword"
	hashedPassword, err := infra.HashPassword(password)
	assert.NoError(t, err)

	// Test with the correct password
	err = infra.ComparePasswords(hashedPassword, password)
	assert.NoError(t, err)

	// Test with an incorrect password
	err = infra.ComparePasswords(hashedPassword, "wrongpassword")
	assert.Error(t, err)
}
