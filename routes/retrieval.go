package routes

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gocarina/gocsv"
	"github.com/likid1412/address/model"
	"github.com/rs/zerolog/log"
)

// RetrievalAllInfo godoc
// @Summary 	Property Information Retrieval
// @Description Retrieve all address information in the specified format from a CSV file
// and 			returns it as a JSON response.
// @Tags		Address
// @Produce		json
// @Param		filename path string true "Name of the CSV file"
// @Success		200 {array} model.AddressInfoRow "Array of address information"
// @Failure		404 {object} model.ErrorResponse "File does not exist"
// @Failure		500 {object} model.ErrorResponse "Get addresses infos failed"
// @Router		/address/retrieval/{filename} [get]
func RetrievalAllInfo(c *gin.Context) {
	filename := c.Param("filename")
	log.Info().Msgf("filename: %s", filename)

	rows, err := getAddressInfosFromFile(filename)
	if err != nil {
		if errors.Is(err, ErrNotExist) {
			c.JSON(http.StatusNotFound, model.ErrorResponse{
				Error: err.Error(),
			})
			return
		}

		msg := fmt.Sprintf("Get addresses infos failed, err: %v", err)
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Error: msg,
		})
		return
	}

	c.JSON(http.StatusOK, rows)
}

// getAddressInfosFromFile reads address information from a CSV file
// and parses it
func getAddressInfosFromFile(filename string) ([]*model.AddressInfoRow, error) {
	filePath := getSavePath(filename)
	if !fileExists(filePath) {
		err := ErrNotExist
		log.Error().Msgf("filePath: %s, %v", filePath, err)
		return nil, err
	}

	file, err := os.Open(filePath)
	if err != nil {
		err = fmt.Errorf("open file failed: %w", err)
		log.Error().Msgf("filePath: %s, %v", filePath, err)
		return nil, err
	}
	defer file.Close()

	// parse csv
	rows := []*model.AddressInfoRow{}
	if err := gocsv.UnmarshalFile(file, &rows); err != nil {
		err = fmt.Errorf("unmarshal csv file failed: %w", err)
		log.Error().Msgf("filePath: %s, %v", filePath, err)
		return nil, err
	}

	// Fill IDs and generate full addresses for each row
	for index, item := range rows {
		// Assign ID starting from 1
		item.ID = strconv.Itoa(index + 1)
		item.FullAddress = getFullAddress(item)
	}

	return rows, nil
}
