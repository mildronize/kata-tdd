package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculate(t *testing.T) {
	// Arrange
	assert := assert.New(t)
	// Act
	var actual = SayHelloWorld()
	// Assert
	assert.Equal(actual, "Hello Wolrd!")
}
