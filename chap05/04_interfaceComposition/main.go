package main

import (
	"errors"
	"fmt"
	"io"
	"math/rand"
	"time"
)

type Data struct {
	Line string
}

type Xenia struct {
	Host    string
	Timeout time.Duration
}

func (x *Xenia) Pull(d *Data) error {
	switch rand.Intn(10) {
	case 9:
		return io.EOF
	case 5:
		return errors.New("error reading data")
	default:
		d.Line = "data"
		fmt.Println("In:", d.Line)
		return nil
	}
}

type Pillar struct {
	Host    string
	Timeout time.Duration
}

func (p *Pillar) Store(d *Data) error {
	fmt.Println("Out:", d.Line)
	return nil
}

type Puller interface {
	Pull(d *Data) error
}

type Storer interface {
	Store(d *Data) error
}

type System struct {
	Puller
	Storer
}

func Pull(sys System, data []Data) (int, error) {
	for i := range data {
		if err := sys.Pull(&data[i]); err != nil {
			return i, err
		}
	}
	return len(data), nil
}

func Store(sys System, data []Data) (int, error) {
	for i := range data {
		if err := sys.Store(&data[i]); err != nil {
			return i, err
		}
	}
	return len(data), nil
}

func Copy(sys System, batch int) error {
	data := make([]Data, batch)

	for {
		i, err := Pull(sys, data)
		if err != nil {
			return err
		}
		if i > 0 {
			_, err = Store(sys, data[:i])
			if err != nil {
				return err
			}
		}
	}
}

func main() {
	sys := System{
		Puller: &Xenia{
			Host:    "localhost:8000",
			Timeout: time.Second,
		},
		Storer: &Pillar{
			Host:    "localhost:9000",
			Timeout: time.Second,
		},
	}

	if err := Copy(sys, 3); err != io.EOF {
		fmt.Println("Error:", err)
	}
}
