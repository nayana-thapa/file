package file

import (
	"errors"
	"path/filepath"
	"regexp"
	"strings"

	"golang.org/x/exp/slices"
)

const (
	fileExpr = "^[a-zA-Z0-9][a-zA-Z0-9-_.]+$"
)

func Valid(filename string) error {

	fileExts := []string{"xls", "xlsx", "sds", "eds", "txt", "zip"}

	filename = strings.TrimSpace(filename)

	if len(filename) == 0 {
		return errors.New("filename is empty")
	}

	fileExtension := filepath.Ext(filename)
	filePrefix := strings.TrimSuffix(filename, fileExtension)

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
