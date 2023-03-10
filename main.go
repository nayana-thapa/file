package file

import (
	"errors"
	"fmt"
	"path/filepath"
	"regexp"
	"strings"

	"golang.org/x/exp/slices"
)

const (
	// Regex to validaete the file name without extension.
	fileExpr = "^[a-zA-Z0-9][a-zA-Z0-9-_.]+$"
)

// Valid checks if file has a valid extension or not.
func Valid(filename string) error {

	fileExts := []string{"txt", "png", "pdf", "jpeg", "jpg", "zip"}

	filename = strings.TrimSpace(filename)

	if len(filename) == 0 {
		return errors.New("filename is empty")
	}

	fileExtension := filepath.Ext(filename)
	filePrefix := strings.TrimSuffix(filename, fileExtension)

	fmt.Printf("This is file name: %v\n", filePrefix)

	match, err := regexp.MatchString(fileExpr, filePrefix)
	if err != nil {
		return err
	}

	if !match {
		return errors.New("regular expression didn't match with filename")
	}

	contains := slices.Contains(fileExts, strings.Trim(fileExtension, "."))
	if !contains {
		return errors.New("invalid extension")
	}

	return nil

}
