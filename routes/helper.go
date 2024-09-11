package routes

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/likid1412/address/model"
)

var (
	_savePath = "uploaded_files"
)

// Error variables to represent file existence states
var (
	ErrExist    = errors.New("file already exists")
	ErrNotExist = errors.New("file does not exist")
)

// isCSVFile checks if the file is a CSV.
//
// Returns true if the file is a CSV, otherwise false.
func isCSVFile(filename string) bool {
	return filepath.Ext(filename) == ".csv"
}

// getFullAddress constructs the full address as a single string
//
// Returns the full address as a string.
func getFullAddress(addr *model.AddressInfoRow) string {
	builder := strings.Builder{}
	builder.WriteString(addr.Prefecture)
	builder.WriteString(addr.City)
	builder.WriteString(addr.Town)
	builder.WriteString(addr.Chome)
	builder.WriteString("-")
	builder.WriteString(addr.Banchi)
	builder.WriteString("-")
	builder.WriteString(addr.Go)
	return builder.String()
}

// getSavePath builds the full save path for a file
//
// Returns the full path as a string.
func getSavePath(filename string) string {
	// Constant representing the directory where files are saved
	return fmt.Sprintf("%s/%s", _savePath, filename)
}

// fileExists checks if a file exists
//
// Returns true if the file exists, otherwise false.
func fileExists(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return err == nil
}
