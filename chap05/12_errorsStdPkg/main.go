package main

import (
	"errors"
	"fmt"
)

type AppErr struct {
	State int
}

func (a *AppErr) Error() string {
	return fmt.Sprintf("App Error, state: %d", a.State)
}

func Cause(err error) error {
	root := err
	for {
		if err = errors.Unwrap(root); err == nil {
			return root
		}
		root = err
	}
}

func firstCall(i int) error {
	if err := secondCall(i); err != nil {
		return fmt.Errorf("secondCall (%d): %w", i, err)
	}
	return nil
}

func secondCall(i int) error {
	return &AppErr{99}
}

func main() {
	if err := firstCall(10); err != nil {
		var ap *AppErr
		if errors.As(err, &ap) {
			fmt.Println("As it says it is an AppErr")
		}

		switch v := Cause(err).(type) {
		case *AppErr:
			fmt.Println("Custom error type:", v.State)
		default:
			fmt.Println("Default error")
		}

		fmt.Println(err)
	}
}
