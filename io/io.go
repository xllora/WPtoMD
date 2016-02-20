// Package io contains basic IO helper routines.
package io

import (
	"io/ioutil"
	"os"
)

// ReadFile returns the context of the file.
func ReadFile(filePath string) ([]byte, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return ioutil.ReadAll(f)
}
