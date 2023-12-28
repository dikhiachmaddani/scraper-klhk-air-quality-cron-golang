package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"quality-air-golang/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetQualityAir(db *gorm.DB) error {
	resp, err := http.Get("https://ispu.menlhk.go.id/apimobile/v1/getDetail/stasiun/MALANG")
	if err != nil {
		fmt.Println("No response from request")
	}
	defer resp.Body.Close()
	byteValue, err := ioutil.ReadAll(resp.Body) // response body is []byte
	var rows models.Rows
	json.Unmarshal(byteValue, &rows)

	pm10, _ := strconv.Atoi(rows.Rows[0].Pm10)
	pm25, _ := strconv.Atoi(rows.Rows[0].Pm25)
	so2, _ := strconv.Atoi(rows.Rows[0].So2)
	co, _ := strconv.Atoi(rows.Rows[0].Co)
	o3, _ := strconv.Atoi(rows.Rows[0].O3)
	no2, _ := strconv.Atoi(rows.Rows[0].No2)
	hc, _ := strconv.Atoi(rows.Rows[0].Hc)
	val, _ := strconv.Atoi(rows.Rows[0].Val)

	qualityAir := models.QualityAir{
		Pm10:      pm10,
		Pm25:      pm25,
		So2:       so2,
		Co:        co,
		O3:        o3,
		No2:       no2,
		Hc:        hc,
		Station:   rows.Rows[0].Station,
		Longitude: rows.Rows[0].Longitude,
		Latitude:  rows.Rows[0].Latitude,
		Waktu:     rows.Rows[0].Waktu,
		Ispu:      val,
		Category:  rows.Rows[0].Cat,
	}

	db.Create(&qualityAir) // pass pointer of data to Create
	return nil
}

func FindQualityAir(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var qualityAir []models.QualityAir
	db.Find(&qualityAir)

	c.JSON(http.StatusOK, gin.H{"data": qualityAir})
}
