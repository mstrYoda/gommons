package unsafe

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnsafeString(t *testing.T) {
	str := String([]byte("test"))

	assert.Equal(t, "test", str)
}

func TestUnsafeStringToByte(t *testing.T) {
	byteArr := Byte("test")

	assert.Equal(t, []byte("test"), byteArr)
}
