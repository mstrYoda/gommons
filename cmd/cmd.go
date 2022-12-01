package cmd

import (
	"fmt"
	"io"
	"os/exec"
)

type Output struct {
	result        string
	executionTime int
}

func Exec(command string, args ...string) []byte {
	ps := exec.Command(command, args...)

	if ps.Err != nil {
		panic(ps.Err)
	}

	out, err := ps.Output()
	if err != nil {
		panic(err)
	}

	return out
}

func ExecPipe(stdIn io.Reader, stdOut, stdErr io.Writer, command string, args ...string) error {
	ps := exec.Command(command, args...)

	if ps.Err != nil {
		return fmt.Errorf("could not create command, err: %w", ps.Err)
	}

	ps.Stdin = stdIn
	ps.Stdout = stdOut
	ps.Stderr = stdErr

	if err := ps.Start(); err != nil {
		return fmt.Errorf("could not start command, err: %w", err)
	}

	if err := ps.Wait(); err != nil {
		return fmt.Errorf("could not complete command, err: %w", err)
	}

	return nil
}
