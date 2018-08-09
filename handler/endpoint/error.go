package endpoint

import (
	"github.com/gin-gonic/gin"
	"apibuilder-server/model"
	"log"
)

type ControllerErrors struct {
	Errors      []*ControllerError `json:"errors"`
}

func (errors *ControllerErrors) Status() int {
	return errors.Errors[0].Status
}

type ControllerError struct {
	Status      int         `json:"status"`
	Code        string      `json:"code"`
	Title       string      `json:"title"`
	Details     string      `json:"details"`
}

func NewControllerError(status int, code string, title string, details string) *ControllerError {
	return &ControllerError{
		Status:     status,
		Code:       code,
		Title:      title,
		Details:    details,
	}
}
func NOContentError(err error) *ControllerError {
	return NewControllerError(400, "not_content", "Not Content", err.Error())
}
func NOChangeError(err error) *ControllerError {
	return NewControllerError(400, "not_change", "Not Change", err.Error())
}
func JsonTypeError(err error) *ControllerError {
	return NewControllerError(400, "json_type", "Json type error", err.Error())
}
func ForbidError(err error) *ControllerError {
	return NewControllerError(400, "forbid", "forbid", err.Error())
}

func CatchErrors() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				switch err.(type) {
				case *ControllerError:
					apiError := err.(*ControllerError)
					apiErrors := &ControllerErrors{
						Errors: []*ControllerError{apiError},
					}
					c.JSON(apiError.Status, apiErrors)
				case *ControllerErrors:
					apiErrors := err.(*ControllerErrors)
					c.JSON(apiErrors.Status(), apiErrors)
				case *model.DaoError:
					daoError := err.(*model.DaoError)
					daoErrors := &model.DaoErrors{
						Errors: []*model.DaoError{daoError},
					}
					c.JSON(daoError.Status, daoErrors)
				case *model.DaoErrors:
					daoErrors := err.(*model.DaoErrors)
					c.JSON(daoErrors.Status(), daoErrors)
				default:
					log.Print(err)
					panic(err)
				}
			}
		}()
		c.Next()
	}
}