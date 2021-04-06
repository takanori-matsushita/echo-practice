package controllers

import (
	"app/database"
	Models "app/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUsers(c echo.Context) error {
	db := database.Connect()
	// defer db.Close()
	var users []Models.User
	db.Find(&users)
	return c.JSON(http.StatusOK, users)
}

func GetUser(c echo.Context) error {
	db := database.Connect()
	// defer db.Close()
	id := c.Param("id")
	var user Models.User
	db.First(&user, id)
	// if err := db.First(&user, id).Error; err != nil {
	// 	// ここでエラーハンドリング
	// 	return c.JSON(http.StatusNotFound, map[string]interface{}{"message": "user not found"})
	// }
	return c.JSON(http.StatusOK, &user)
}

func CreateUser(c echo.Context) error {
	db := database.Connect()
	// defer db.Close()
	name, email, password := c.FormValue("name"), c.FormValue("email"), c.FormValue("password")
	user := Models.User{Name: name, Email: email, Password: password}
	db.Create(&user)
	if err := db.Create(&user).Error; err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}
	return c.JSON(http.StatusOK, &user)
}

func UpdateUser(c echo.Context) error {
	db := database.Connect()
	// defer db.Close()
	id := c.Param("id")
	var user Models.User
	if err := db.First(&user, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{"message": "user not found"})
	}
	user.Name, user.Email, user.Password = c.FormValue("name"), c.FormValue("email"), c.FormValue("password")
	if err := db.Save(&user).Error; err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	return c.JSON(http.StatusOK, &user)
}

func DeleteUser(c echo.Context) error {
	db := database.Connect()
	// defer db.Close()
	id := c.Param("id")
	var user Models.User
	if err := db.Delete(&user, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{"message": "user not found"})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"message": "user deleted"})
}
