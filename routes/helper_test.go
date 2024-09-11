package routes

import (
	"os"
	"testing"

	"github.com/likid1412/address/model"
	"github.com/stretchr/testify/assert"
)

func TestIsCSVFile(t *testing.T) {
	tests := []struct {
		filename string
		expected bool
	}{
		{"file.csv", true},
		{"file.txt", false},
		{"file.", false},
		{"file.csv.zip", false},
	}

	for _, tt := range tests {
		t.Run(tt.filename, func(t *testing.T) {
			result := isCSVFile(tt.filename)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestGetFullAddress(t *testing.T) {
	tests := []struct {
		addr     *model.AddressInfoRow
		expected string
	}{
		{
			addr: &model.AddressInfoRow{
				Prefecture:     "神奈川県",
				City:           "横須賀市",
				Town:           "追浜",
				Chome:          "10",
				Banchi:         "11",
				Go:             "15",
				Building:       "東京オペラシティ",
				Price:          "23871590",
				NearestStation: "戸塚駅",
				PropertyType:   "一戸建て",
				LandArea:       "4167",
			},
			expected: "神奈川県横須賀市追浜10-11-15",
		},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			result := getFullAddress(tt.addr)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestGetSavePath(t *testing.T) {
	_savePath = "/my/directory"

	tests := []struct {
		filename string
		expected string
	}{
		{"file.csv", "/my/directory/file.csv"},
	}

	for _, tt := range tests {
		t.Run(tt.filename, func(t *testing.T) {
			result := getSavePath(tt.filename)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestFileExists(t *testing.T) {
	// Create a temporary file for testing
	file, err := os.CreateTemp("", "testfile")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(file.Name()) // Clean up

	tests := []struct {
		path     string
		expected bool
	}{
		{file.Name(), true},
		{"invalid/path", false},
	}

	for _, tt := range tests {
		t.Run(tt.path, func(t *testing.T) {
			result := fileExists(tt.path)
			assert.Equal(t, tt.expected, result)
		})
	}
}
