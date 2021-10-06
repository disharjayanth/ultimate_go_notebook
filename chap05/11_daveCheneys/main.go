package main

import (
	"fmt"

	"github.com/pkg/errors"
)

type AppError struct {
	State int
}

func (a *AppError) Error() string {
	return fmt.Sprintf("App Error, State; %d", a.State)
}

func firstCall(i int) error {
	if err := secondCall(i); err != nil {
		return errors.Wrapf(err, "secondCall(%d)", i)
	}
	return nil
}

func secondCall(i int) error {
	return &AppError{99}
}

func main() {
	if err := firstCall(10); err != nil {
		switch v := errors.Cause(err).(type) {
		case *AppError:
			fmt.Println("Custom app error:", v.State)
		default:
			fmt.Println("default error!")
		}

		fmt.Printf("%v\n", err)
	}
}
