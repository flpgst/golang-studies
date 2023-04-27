package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("John Doe", "jd@test.com", "123456")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "John Doe", user.Name)
	assert.Equal(t, "jd@test.com", user.Email)
}

func TestUser_ValidadePassword(t *testing.T) {
	user, err := NewUser("John Doe", "jd@test.com", "123456")
	assert.Nil(t, err)
	assert.True(t, user.ValidadePassword("123456"))
	assert.False(t, user.ValidadePassword("InvalidPassword"))
	assert.NotEqual(t, "123456", user.Password)
}
