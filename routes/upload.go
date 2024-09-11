package routes

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/likid1412/address/model"
	"github.com/rs/zerolog/log"
)

// Upload godoc
// @Summary 	CSV Upload
// @Description Accept a CSV file containing Japanese addresses and property information
// @Tags 		Address
// @Accept 		mpfd
// @Produce		plain
// @Param		file formData file true "CSV file to upload"
// @Success		200 {string} string ""
// @Failure		400 {object} model.ErrorResponse "Unable to retrieve file from request"
// @Failure		415 {object} model.ErrorResponse "Unsupported file format, only CSV files are allowed"
// @Failure		409 {object} model.ErrorResponse "File already exists"
// @Failure		500 {object} model.ErrorResponse "Save file failed"
// @Router		/address/upload [post]
func Upload(c *gin.Context) {
	filename, file, hasErr := checkFile(c)
	if hasErr {
		return
	}

	hasErr = saveFile(c, filename, file)
	if hasErr {
		return
	}

	c.String(http.StatusOK, "")
}

// checkFile checks if the uploaded file is a valid CSV file.
func checkFile(c *gin.Context) (
	filename string, file *multipart.FileHeader, hasErr bool) {

	file, err := c.FormFile("file")
	if err != nil {
		msg := fmt.Sprintf("Unable to retrieve file from request: %v", err)
		log.Error().Msg(msg)
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Error: msg,
		})
		hasErr = true
		return
	}

	filename = filepath.Base(file.Filename)
	log.Info().Msgf("filename: %s", filename)

	if !isCSVFile(filename) {
		msg := "Unsupported file format, only CSV files are allowed"
		log.Error().Msgf("%s, filename: %s", msg, filename)
		c.JSON(http.StatusUnsupportedMediaType, model.ErrorResponse{
			Error: msg,
		})
		hasErr = true
		return
	}

	return
}

// saveFile saves the uploaded file to the server, ensuring the file doesn't already exist.
func saveFile(c *gin.Context, filename string, file *multipart.FileHeader) (
	hasErr bool) {
	// Here are three solutions to handle same filenames:
	//
	// 1. indicate that the file is already existed, and return 409 Conflict
	// 2. allow users to overwrite existing files, and return 200 OK
	// 3. using unique name, such as UUIDs
	//
	// use the first solution here, return 409 Conflict
	filePath := getSavePath(filename)
	if fileExists(filePath) {
		msg := ErrExist.Error()
		log.Error().Msgf("%s, filePath: %s", msg, filePath)
		c.JSON(http.StatusConflict, model.ErrorResponse{
			Error: msg,
		})
		hasErr = true
		return
	}

	if err := c.SaveUploadedFile(file, filePath); err != nil {
		msg := fmt.Sprintf("Save file failed: %v", err)
		log.Error().Msgf("%s, filePath: %s", msg, filePath)
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Error: msg,
		})
		hasErr = true
		return
	}

	return
}
