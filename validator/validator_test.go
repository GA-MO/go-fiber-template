package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type role struct {
	Level int `validate:"required,min=1,max=10"`
}

type user struct {
	Name  string `validate:"required"`
	Email string `validate:"required,email=invalid"`
	Age   int    `validate:"required,min=18"`
	Role  role
}

func TestValidateStruct(t *testing.T) {
	t.Run("test validate struct", func(t *testing.T) {
		user := &user{
			Name:  "test",
			Email: "test@test.com",
			Age:   18,
			Role: role{
				Level: 5,
			},
		}
		err := ValidateStruct(user)
		assert.NoError(t, err)
	})

	t.Run("test validate struct with error", func(t *testing.T) {
		user := &user{
			Name:  "test",
			Email: "test",
			Age:   25,
			Role: role{
				Level: 11,
			},
		}
		err := ValidateStruct(user)
		assert.Error(t, err)
	})
}
