package routes

import (
	"errors"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	. "github.com/bytedance/mockey"
)

func TestUploadSuccess(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockFileHeader := &multipart.FileHeader{
		Filename: "test.csv",
	}

	req := httptest.NewRequest(http.MethodPost, "/", nil)
	rec := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rec)
	c.Request = req

	PatchConvey("TestUploadSuccess", t, func() {
		Mock(checkFile).Return(mockFileHeader.Filename, mockFileHeader, false).Build()
		Mock(saveFile).Return(false).Build()

		Upload(c)

		assert.Equal(t, http.StatusOK, rec.Code)
	})
}

func TestUploadBadRequest(t *testing.T) {
	gin.SetMode(gin.TestMode)

	req := httptest.NewRequest(http.MethodPost, "/", nil)
	rec := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rec)
	c.Request = req

	PatchConvey("TestUploadBadRequest", t, func() {
		Upload(c)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.JSONEq(t, `{"error": "Unable to retrieve file from request: request Content-Type isn't multipart/form-data"}`,
			rec.Body.String())
	})

}

func TestUploadUnsupportedFormat(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockFileHeader := &multipart.FileHeader{
		Filename: "test.txt",
	}

	req := httptest.NewRequest(http.MethodPost, "/", nil)
	rec := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rec)
	c.Request = req

	PatchConvey("TestUploadUnsupportedFormat", t, func() {
		Mock((*gin.Context).FormFile).Return(mockFileHeader, nil).Build()

		Upload(c)

		assert.Equal(t, http.StatusUnsupportedMediaType, rec.Code)
		assert.JSONEq(t, `{"error": "Unsupported file format, only CSV files are allowed"}`,
			rec.Body.String())
	})
}

func TestUploadFileExists(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockFileHeader := &multipart.FileHeader{
		Filename: "test.txt",
	}

	req := httptest.NewRequest(http.MethodPost, "/", nil)
	rec := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rec)
	c.Request = req

	PatchConvey("TestUploadFileExists", t, func() {
		Mock(checkFile).Return(mockFileHeader.Filename, mockFileHeader, false).Build()
		Mock(fileExists).Return(true).Build()

		Upload(c)

		assert.Equal(t, http.StatusConflict, rec.Code)
		assert.JSONEq(t, `{"error": "file already exists"}`,
			rec.Body.String())
	})
}

func TestUploadSaveFailed(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockFileHeader := &multipart.FileHeader{
		Filename: "test.txt",
	}

	req := httptest.NewRequest(http.MethodPost, "/", nil)
	rec := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rec)
	c.Request = req

	PatchConvey("TestUploadSaveFailed", t, func() {
		Mock(checkFile).Return(mockFileHeader.Filename, mockFileHeader, false).Build()
		Mock(fileExists).Return(false).Build()
		Mock((*gin.Context).SaveUploadedFile).Return(errors.New("mock failed")).Build()

		Upload(c)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.JSONEq(t, `{"error": "Save file failed: mock failed"}`,
			rec.Body.String())
	})
}
