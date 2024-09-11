package routes

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/gocarina/gocsv"
	"github.com/likid1412/address/model"
	"github.com/stretchr/testify/assert"

	. "github.com/bytedance/mockey"
)

func testSuccessRows() []*model.AddressInfoRow {
	return []*model.AddressInfoRow{
		{
			ID:             "1",
			FullAddress:    "神奈川県横須賀市追浜10-11-15",
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
	}
}
func testSuccessRowsString() string {
	return `[{"id":"1","full_address":"神奈川県横須賀市追浜10-11-15","prefecture":"神奈川県","city":"横須賀市","town":"追浜","chome":"10","banchi":"11","go":"15","building":"東京オペラシティ","price":"23871590","nearest_station":"戸塚駅","property_type":"一戸建て","land_area":"4167"}]`
}

func TestRetrievalAllInfoSuccess(t *testing.T) {
	// Initialize Gin router for testing
	gin.SetMode(gin.TestMode)

	// Create a mock context
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Params = []gin.Param{
		{
			Key:   "filename",
			Value: "test.csv",
		},
	}

	PatchConvey("TestRetrievalAllInfoSuccess", t, func() {
		expectedRows := testSuccessRows()

		Mock(getAddressInfosFromFile).Return(expectedRows, nil).Build()

		RetrievalAllInfo(c)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.JSONEq(t, testSuccessRowsString(), w.Body.String())
	})
}

func TestRetrievalAllInfoFileNotFound(t *testing.T) {
	// Initialize Gin router for testing
	gin.SetMode(gin.TestMode)

	// Create a mock context
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Params = []gin.Param{
		{
			Key:   "filename",
			Value: "nonexistent.csv",
		},
	}

	PatchConvey("TestRetrievalAllInfoFileNotFound", t, func() {
		Mock(getAddressInfosFromFile).Return(nil, ErrNotExist).Build()

		// Call the handler function
		RetrievalAllInfo(c)

		// Check the status code and response
		assert.Equal(t, http.StatusNotFound, w.Code)
		assert.JSONEq(t, `{"error": "file does not exist"}`, w.Body.String())
	})
}

func TestRetrievalAllInfoInternalError(t *testing.T) {
	// Initialize Gin router for testing
	gin.SetMode(gin.TestMode)

	// Create a mock context
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Params = []gin.Param{
		{
			Key:   "filename",
			Value: "error.csv",
		},
	}

	PatchConvey("TestRetrievalAllInfoInternalError", t, func() {
		Mock(getAddressInfosFromFile).Return(nil, errors.New("internal error")).Build()

		RetrievalAllInfo(c)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Contains(t, w.Body.String(), "Get addresses infos failed")
	})
}

func TestGetAddressInfosFromFileSuccess(t *testing.T) {
	// Setup a temporary CSV file with test data
	testData := testSuccessRows()

	file, err := os.CreateTemp("", "test_*.csv")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(file.Name()) // Clean up

	if err := gocsv.MarshalFile(&testData, file); err != nil {
		t.Fatal(err)
	}
	file.Close()

	_savePath = filepath.Dir(file.Name())
	filename := filepath.Base(file.Name())
	rows, err := getAddressInfosFromFile(filename)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	assert.ObjectsAreEqual(testData, rows)
}

func TestGetAddressInfosFromFileUnmarshalFailed(t *testing.T) {
	// Setup a temporary empty file with test data
	file, err := os.CreateTemp("", "test_*.csv")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(file.Name()) // Clean up

	_savePath = filepath.Dir(file.Name())
	filename := filepath.Base(file.Name())
	_, err = getAddressInfosFromFile(filename)
	assert.NotNil(t, err)
}
