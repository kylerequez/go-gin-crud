package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kylerequez/go-gin-api/initializers"
	"github.com/kylerequez/go-gin-api/models"
	"github.com/kylerequez/go-gin-api/services"
)

func RegistrationHandler(c *gin.Context) {
	fmt.Println(":::-:::\tREGISTRATION HANDLER\t:::-:::")
	var body struct {
		Name  string
		Email string
		Age   uint8
	}

	c.Bind(&body)

	name := body.Name
	email := body.Email
	var age uint8 = body.Age

	if name == "" || email == "" || age <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error",
			"error":   "The form is incomplete. Please fill out all the necessary information",
		})
		return
	}

	user := models.User{Name: name, Email: &email, Age: &age}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error",
			"error":   result.Error,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"user":    user,
	})
}

func LoginHandler(c *gin.Context) {
	fmt.Println(":::-:::\tLOGIN HANDLER\t:::-:::")
	var body struct {
		Email string
	}

	c.Bind(&body)

	if body.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error",
			"error":   "You have entered an incorrect email",
		})
		return
	}

	var user models.User
	userFound := initializers.DB.First(&user, "email = ?", body.Email)
	if userFound.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error",
			"error":   userFound.Error,
		})
		return
	}

	if userFound.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Error",
			"error":   userFound.Error,
		})
		return
	}

	tokenString, err := services.GenerateJWT(*user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error",
			"error":   "There was an error in generating your authentication token",
		})
		return
	}

	c.SetCookie("go-gin-crud-token", tokenString, 3600, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"token":   tokenString,
	})
}

func GetUsers(c *gin.Context) {
	fmt.Println(":::-:::\tGET ALL USERS\t:::-:::")
	var users []models.User
	result := initializers.DB.Find(&users)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error",
			"error":   result.Error,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"users":   users,
	})
}

func GetUserById(c *gin.Context) {
	fmt.Println(":::-:::\tGET USER BY ID\t:::-:::")
	var uri struct {
		ID string `uri:"id" binding:"required"`
	}

	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error",
			"error":   err,
		})
		return
	}

	var user models.User
	initializers.DB.Model(&models.User{}).First(&user, "id = ?", uri.ID)

	if (user == models.User{}) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Success",
			"user":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"user":    user,
	})
}

func CreateUser(c *gin.Context) {
	fmt.Println(":::-:::\tPOST USER CREATE\t:::-:::")
	var body struct {
		Name  string
		Email string
		Age   uint8
	}

	c.Bind(&body)

	name := body.Name
	email := body.Email
	var age uint8 = body.Age

	user := models.User{Name: name, Email: &email, Age: &age}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error",
			"error":   result.Error,
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Success",
		"user":    user,
	})
}

func PutUpdateUser(c *gin.Context) {
	fmt.Println(":::-:::\tPUT USER UPDATE\t:::-:::")
	var uri struct {
		ID string `uri:"id" binding:"required"`
	}

	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error",
			"error":   err,
		})
		return
	}

	var body struct {
		Name  string
		Email string
		Age   uint8
	}

	c.Bind(&body)

	user := models.User{ID: uri.ID}
	userFound := initializers.DB.First(&user)

	if userFound.Error != nil && userFound.RowsAffected != 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error",
			"error":   userFound.Error,
		})
		return
	}

	if userFound.RowsAffected == 0 {
		user = models.User{
			Name:  body.Name,
			Email: &body.Email,
			Age:   &body.Age,
		}
	} else {
		user.Name = body.Name
		user.Email = &body.Email
		user.Age = &body.Age
	}

	isSaved := initializers.DB.Save(&user)
	if isSaved.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error",
			"error":   isSaved.Error,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Success",
		"user":    user,
	})
}

func PatchUpdateUser(c *gin.Context) {
	fmt.Println(":::-:::\tPATCH USER UPDATE\t:::-:::")
	var uri struct {
		ID string `uri:"id" binding:"required"`
	}

	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error",
			"error":   err,
		})
		return
	}

	var body struct {
		Name  string
		Email string
		Age   uint8
	}

	c.Bind(&body)

	user := models.User{ID: uri.ID}
	userFound := initializers.DB.First(&user)

	if userFound.Error != nil && userFound.RowsAffected != 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error",
			"error":   userFound.Error,
		})
		return
	}

	if userFound.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "User not found",
		})
		return
	}

	user.Name = body.Name
	user.Email = &body.Email
	user.Age = &body.Age

	isSaved := initializers.DB.Save(&user)
	if isSaved.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error",
			"error":   isSaved.Error,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Success",
		"user":    user,
	})
}

func DeleteUserById(c *gin.Context) {
	var uri struct {
		ID string `uri:"id" binding:"required"`
	}

	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error",
			"error":   err,
		})
		return
	}

	var user models.User = models.User{ID: uri.ID}
	result := initializers.DB.Delete(&user)

	if result.Error != nil {
		c.JSON(http.StatusConflict, gin.H{
			"message": "Error",
			"error":   result.Error,
		})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusConflict, gin.H{
			"message": "Error",
			"error":   result.Error,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
	})
}
