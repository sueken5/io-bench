package io_bench

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type test struct {
	Message string `json:"message"`
}

//not use io interface.
func Ex1() error {
	f, err := os.Open("test.json")
	if err != nil {
		return fmt.Errorf("ex1 err: %w", err)
	}

	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		return fmt.Errorf("ex1 err: %w", err)
	}

	t := &test{}
	if err := json.Unmarshal(data, t); err != nil {
		return fmt.Errorf("ex1 err: %w", err)
	}

	return nil
}


//use io interface.
func Ex2() error {
	f, err := os.Open("test.json")
	if err != nil {
		return fmt.Errorf("ex2 err: %w", err)
	}

	defer f.Close()

	t := &test{}
	if err := json.NewDecoder(f).Decode(t); err != nil {
		return fmt.Errorf("ex1 err: %w", err)
	}

	return nil
}
