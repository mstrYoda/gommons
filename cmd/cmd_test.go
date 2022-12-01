package cmd

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestExec(t *testing.T) {
	out := Exec("echo", "test")

	assert.Equal(t, []byte("test\n"), out)
}

func TestExecPipe(t *testing.T) {
	strReader := strings.NewReader("hello world")

	outWriter := bytes.NewBuffer(nil)
	errWriter := bytes.NewBuffer(nil)

	ExecPipe(strReader, outWriter, errWriter, "echo", "test")
	outputStr := outWriter.String()

	assert.Equal(t, "test\n", outputStr)
}
