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
	case 1, 9:
		return io.EOF
	case 5:
		return errors.New("error reading data")
	default:
		d.Line = "Data"
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

func Pull(ps PullerStorer, data []Data) (int, error) {
	for i := range data {
		if err := ps.Pull(&data[i]); err != nil {
			return i, err
		}
	}
	return len(data), nil
}

func Store(ps PullerStorer, data []Data) (int, error) {
	for i := range data {
		if err := ps.Store(&data[i]); err != nil {
			return i, err
		}
	}
	return len(data), nil
}

type Puller interface {
	Pull(d *Data) error
}

type Storer interface {
	Store(d *Data) error
}

type PullerStorer interface {
	Puller
	Storer
}

type System struct {
	Xenia
	Pillar
}

func Copy(ps PullerStorer, batch int) error {
	data := make([]Data, batch)

	for {
		i, err := Pull(ps, data)
		if err != nil {
			return err
		}
		if i > 0 {
			if _, err = Store(ps, data[:i]); err != nil {
				return err
			}
		}
	}
}

func main() {
	sys := System{
		Xenia: Xenia{
			Host:    "localhost:8000",
			Timeout: time.Second,
		},
		Pillar: Pillar{
			Host:    "localhost:8000",
			Timeout: time.Second,
		},
	}

	if err := Copy(&sys, 10); err != io.EOF {
		fmt.Println(err)
	}
}
