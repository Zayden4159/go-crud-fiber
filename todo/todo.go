package todo

import (
	"strings"

	"github.com/Zayden4159/crud-fiber/response"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

var Users = []User{}

func GetUsers(c *fiber.Ctx) error {
	return response.Success(c, fiber.StatusOK, response.GetUsersSuccess, Users)
}

func GetUser(c *fiber.Ctx) error {
	userId := c.Params("id")

	for _, user := range Users {
		if user.ID == userId {
			return response.Success(c, fiber.StatusOK, response.GetUserSuccess, user)
		}
	}

	return response.Error(c, fiber.StatusNotFound, response.UserNotFound, nil)
}

func CreateUser(c *fiber.Ctx) error {
	user := new(User)

	if err := c.BodyParser(user); err != nil {
		return response.Error(c, fiber.StatusBadRequest, response.ValidationError, nil)
	}

	if user.Name == "" || user.Email == "" {
		return response.Error(c, fiber.StatusBadRequest, response.UserNotFound, nil)
	}
	id := strings.ReplaceAll(uuid.NewString(), "-", "")
	user.ID = id
	Users = append(Users, *user)

	return response.Success(c, fiber.StatusCreated, response.CreateUserSuccess, user)
}

func UpdateUser(c *fiber.Ctx) error {
	userId := c.Params("id")

	userUpdate := new(User)
	if err := c.BodyParser(userUpdate); err != nil {
		return response.Error(c, fiber.StatusBadRequest, response.ValidationError, nil)
	}

	for i, user := range Users {
		if user.ID == userId {
			Users[i].Name = userUpdate.Name
			Users[i].Email = userUpdate.Email
			Users[i].Phone = userUpdate.Phone
			return response.Success(c, fiber.StatusOK, response.UpdateUserSuccess, Users[i])
		}
	}

	return response.Error(c, fiber.StatusNotFound, response.UserNotFound, nil)
}

func DeleteUser(c *fiber.Ctx) error {
	userId := c.Params("id")

	for i, user := range Users {
		if user.ID == userId {
			Users = append(Users[:i], Users[i+1:]...)
			return response.Success(c, fiber.StatusOK, response.DeleteUserSuccess, user)
		}
	}

	return response.Error(c, fiber.StatusNotFound, response.UserNotFound, nil)
}

func UploadFile(c *fiber.Ctx) error {
	file, err := c.FormFile("image")

	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, response.ValidationError, nil)
	}

	err = c.SaveFile(file, "./upload/"+file.Filename)

	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, response.InternalServerError, nil)
	}

	return c.SendString("File upload complete!")
}
