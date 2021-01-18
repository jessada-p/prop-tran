package todos

import (
	"net/http"

	"github.com/pallat/todos/auth"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/pallat/todos/logger"
)

// NewNewTaskHandler create
func NewNewTaskHandler(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {

		logger := logger.Extract(c)
		logger.Info("new task todo........")

		if ok, err := auth.BearerAuthKey(c); !ok {
			return err
		}

		db.AutoMigrate(Task{})

		var todo struct {
			Task string `json:"task"`
		}

		if err := c.Bind(&todo); err != nil {
			log.Error(err.Error())
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": errors.Wrap(err, "new task").Error(),
			})
		}

		if err := db.Create(&Task{
			Task: todo.Task,
		}).Error; err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"err": errors.Wrap(err, "create task").Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]string{})
	}
}

// NewGetTaskHandler get
func NewGetTaskHandler(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {

		logger := logger.Extract(c)
		logger.Info("get task todo........")

		if ok, err := auth.BearerAuthKey(c); !ok {
			return err
		}

		id := c.Param("id")

		if id == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"err": "id cannot be empty",
			})
		}

		todo := Task{}

		if err := db.Model(&Task{}).Where("id = ?", id).First(&todo).Error; err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"err": errors.Wrap(err, "get task").Error(),
			})
		}

		return c.JSON(http.StatusOK, todo)
	}
}

// NewGetAllTaskHandler get all
func NewGetAllTaskHandler(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {

		logger := logger.Extract(c)
		logger.Info("get all task todo........")

		if ok, err := auth.BearerAuthKey(c); !ok {
			return err
		}

		todo := []Task{}

		if err := db.Model(&Task{}).Find(&todo).Error; err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"err": errors.Wrap(err, "get all task").Error(),
			})
		}

		return c.JSON(http.StatusOK, todo)
	}
}

// NewPutTaskHandler update
func NewPutTaskHandler(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {

		logger := logger.Extract(c)
		logger.Info("put task todo........")

		if ok, err := auth.BearerAuthKey(c); !ok {
			return err
		}

		id := c.Param("id")

		if id == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"err": "id cannot be empty",
			})
		}

		if err := db.Model(&Task{}).Where("id = ?", id).Update("Processed", true).Error; err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"err": errors.Wrap(err, "update task").Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]string{
			"msg": "success",
		})
	}
}

// NewDeleteTaskHandler delete
func NewDeleteTaskHandler(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {

		logger := logger.Extract(c)
		logger.Info("delete task todo........")

		if ok, err := auth.BearerAuthKey(c); !ok {
			return err
		}

		id := c.Param("id")

		if id == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"err": "id cannot be empty",
			})
		}

		if err := db.Model(&Task{}).Where("id = ?", id).Delete(id).Error; err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"err": errors.Wrap(err, "delete task").Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]string{
			"msg": "success",
		})
	}
}

// Task struct
type Task struct {
	gorm.Model
	Task      string
	Processed bool
}

// TableName todos
func (Task) TableName() string {
	return "todos"
}
