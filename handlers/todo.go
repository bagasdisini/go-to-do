package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go-to-do/dto"
	"go-to-do/models"
	"go-to-do/pkg"
	"go-to-do/util"
	"net/http"
	"strconv"
)

func GetTodos(c *fiber.Ctx) error {
	activityID, _ := strconv.Atoi(c.Query("activity_group_id"))

	db := pkg.DB
	if activityID != 0 {
		db = db.Where("activity_group_id = ?", activityID)
	}
	var result []models.Todo
	db.Find(&result)

	var todos []models.GetTodoResponse
	for _, todo := range result {
		var isActive bool
		if todo.IsActive == "1" {
			isActive = true
		}
		todos = append(todos, models.GetTodoResponse{
			ID:              todo.ID,
			ActivityGroupID: strconv.Itoa(todo.ActivityGroupID),
			Title:           todo.Title,
			IsActive:        isActive,
			Priority:        todo.Priority,
			CreatedAt:       todo.CreatedAt,
			UpdatedAt:       todo.UpdatedAt,
			DeletedAt:       todo.DeletedAt,
		})
	}

	return c.JSON(dto.Response{
		Status:  "Success",
		Message: "Success",
		Data:    todos,
	})
}

func CreateTodo(c *fiber.Ctx) error {
	todo := new(models.Todo)
	if err := c.BodyParser(&todo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.Response{
			Status:  http.StatusText(http.StatusBadRequest),
			Message: err.Error(),
			Data:    map[string]interface{}{},
		})
	}

	errors := util.ValidateStruct(*todo)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.Response{
			Status:  http.StatusText(http.StatusBadRequest),
			Message: fmt.Sprintf("%v cannot be null", errors[0].FailedField),
			Data:    map[string]interface{}{},
		})
	}

	db := pkg.DB
	todo.IsActive = "1"
	todo.Priority = "very-high"

	db.Create(&todo)

	return c.Status(fiber.StatusCreated).JSON(dto.Response{
		Status:  "Success",
		Message: "Success",
		Data: models.CreateTodoResponse{
			ID:              todo.ID,
			ActivityGroupID: todo.ActivityGroupID,
			Title:           todo.Title,
			IsActive:        true,
			Priority:        todo.Priority,
			CreatedAt:       todo.CreatedAt,
			UpdatedAt:       todo.UpdatedAt,
			DeletedAt:       todo.DeletedAt,
		},
	})
}

func GetTodo(c *fiber.Ctx) error {
	db := pkg.DB
	var todo models.Todo
	id, _ := strconv.Atoi(c.Params("id"))
	if err := db.First(&todo, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(dto.Response{
			Status:  http.StatusText(http.StatusNotFound),
			Message: fmt.Sprintf("Todo with ID %v Not Found", id),
			Data:    map[string]interface{}{},
		})
	}

	var isActive bool
	if todo.IsActive == "1" {
		isActive = true
	}

	return c.JSON(dto.Response{
		Status:  "Success",
		Message: "Success",
		Data: models.GetTodoResponse{
			ID:              todo.ID,
			ActivityGroupID: strconv.Itoa(todo.ActivityGroupID),
			Title:           todo.Title,
			IsActive:        isActive,
			Priority:        todo.Priority,
			CreatedAt:       todo.CreatedAt,
			UpdatedAt:       todo.UpdatedAt,
			DeletedAt:       todo.DeletedAt,
		},
	})
}

func UpdateTodo(c *fiber.Ctx) error {
	req := new(models.Todo)
	_ = c.BodyParser(&req)

	db := pkg.DB
	id, _ := strconv.Atoi(c.Params("id"))
	todo := new(models.Todo)
	if err := db.First(&todo, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(dto.Response{
			Status:  http.StatusText(http.StatusNotFound),
			Message: fmt.Sprintf("Todo with ID %v Not Found", id),
			Data:    map[string]interface{}{},
		})
	}

	if req.Title != "" {
		todo.Title = req.Title
	}
	if req.ActivityGroupID != 0 {
		todo.ActivityGroupID = req.ActivityGroupID
	}
	if req.IsActive != "" {
		todo.IsActive = req.IsActive
	}
	if req.Priority != "" {
		todo.Priority = req.Priority
	}

	db.Save(&todo)

	return c.JSON(dto.Response{
		Status:  "Success",
		Message: "Success",
		Data:    todo,
	})
}

func DeleteTodo(c *fiber.Ctx) error {
	db := pkg.DB
	id, _ := strconv.Atoi(c.Params("id"))

	res := db.Delete(&models.Todo{}, id)
	if res.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(dto.Response{
			Status:  http.StatusText(http.StatusNotFound),
			Message: fmt.Sprintf("Todo with ID %v Not Found", id),
			Data:    map[string]interface{}{},
		})
	}

	return c.JSON(dto.Response{
		Status:  "Success",
		Message: "Success",
		Data:    map[string]interface{}{},
	})
}
