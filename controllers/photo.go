package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"create-api/database"
	"create-api/helpers"
	"create-api/models"

	"github.com/dgrijalva/jwt-go"

	"github.com/gin-gonic/gin"
)

func ListPhoto(c *gin.Context) {
	var (
		photos []models.Photo
	)

	db := database.StartDB()
	err := db.Preload("User").Find(&photos).Error

	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": &photos,
	})
}

func CreatePhoto(c *gin.Context) {
	db := database.StartDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Photo := models.Photo{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserID = userID

	err := db.Debug().Create(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":         Photo.ID,
		"title":      Photo.Title,
		"caption":    Photo.Caption,
		"photo_url":  Photo.PhotoUrl,
		"user_id":    Photo.UserID,
		"created_at": Photo.CreatedAt,
	})
}

func UpdatePhoto(c *gin.Context) {
	db := database.StartDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType

	Photo := models.Photo{}
	NewPhoto := models.Photo{}

	id := c.Param("userId")

	err := db.First(&Photo, id).Error
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	err = json.NewDecoder(c.Request.Body).Decode(&NewPhoto)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	err = db.Model(&Photo).Updates(NewPhoto).Error
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":         Photo.ID,
		"title":      Photo.Title,
		"caption":    Photo.Caption,
		"photo_url":  Photo.PhotoUrl,
		"updated_at": Photo.UpdatedAt,
	})
}

func DeletePhoto(c *gin.Context) {
	db := database.StartDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	Photo := models.Photo{}

	id := c.Param("photoId")

	err := db.First(&Photo, id).Error
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	err = db.Delete(&Photo).Error
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your photo has been successfully deleted",
	})
}
