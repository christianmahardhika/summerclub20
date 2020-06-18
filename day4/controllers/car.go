package controllers

import (
	"day4/models"
	"fmt"

	"github.com/gin-gonic/gin"
)

func CreateCar(context *gin.Context) {
	type carModel struct {
		Name  string `json:"name"`
		Year  int    `json:"year"`
		Color string `json:"color"`
	}

	// Validate input
	var input carModel
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(500, gin.H{"success": false, "result": nil, "errorMessage": err.Error()})
		return
	}

	// Create todo
	car := models.Car{Name: input.Name, Year: input.Year, Color: input.Color}
	models.DB.Create(&car)

	context.JSON(201, gin.H{"success": true, "result": car, "errorMessage": nil})
}

func UpdateCar(context *gin.Context) {
	type carUpdateModel struct {
		Name  string `json:"name"`
		Year  int    `json:"year"`
		Color string `json:"color"`
	}

	// Get model if exist
	var car models.Car
	if err := models.DB.Where("id = ?", context.Query("id")).First(&car).Error; err != nil {
		context.JSON(404, gin.H{"success": false, "result": nil, "errorMessage": "Record not found!"})
		return
	}

	// Validate input
	var input carUpdateModel
	fmt.Println(input)
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(500, gin.H{"success": false, "result": nil, "errorMessage": err.Error()})
		return
	}

	models.DB.Model(&car).Updates(input)

	context.JSON(201, gin.H{"success": true, "result": car, "errorMessage": nil})
}

func DeleteCar(context *gin.Context) {
	// Get model if exist
	var car models.Car
	if err := models.DB.Where("id = ?", context.Query("id")).First(&car).Error; err != nil {
		context.JSON(404, gin.H{"success": false, "result": nil, "errorMessage": "Record not found!"})
		return
	}

	models.DB.Delete(&car)

	context.JSON(201, gin.H{"success": true, "result": context.Query("id"), "errorMessage": nil})
}

func FindCars(context *gin.Context) {
	var cars []models.Car
	models.DB.Find(&cars)

	context.JSON(200, gin.H{"success": true, "result": cars, "errorMessage": nil})
}
