package endpoint

import (
	"github.com/gin-gonic/gin"
	"apibuilder-server/model"
)

type APIErrors struct {
	Errors      []*APIError `json:"errors"`
}

func (errors *APIErrors) Status() int {
	return errors.Errors[0].Status
}

type APIError struct {
	Status      int         `json:"status"`
	Code        string      `json:"code"`
	Title       string      `json:"title"`
	Details     string      `json:"details"`
	Href        string      `json:"href"`
}

func NewAPIError(status int, code string, title string, details string, href string) *APIError {
	return &APIError{
		Status:     status,
		Code:       code,
		Title:      title,
		Details:    details,
		Href:       href,
	}
}
func NOContentError(err error) *APIError {
	return NewAPIError(400, "not_content", "Not Content", err.Error(), "")
}
func NOChangeError(err error) *APIError {
	return NewAPIError(400, "not_change", "Not Change", err.Error(), "")
}
func JsonTypeError(err error) *APIError {
	return NewAPIError(400, "json_type", "Json type error", err.Error(), "")
}
func ForbidError(err error) *APIError {
	return NewAPIError(400, "forbid", "forbid", err.Error(), "")
}

func CatchErrors() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				switch err.(type) {
				case *APIError:
					apiError := err.(*APIError)
					apiErrors := &APIErrors{
						Errors: []*APIError{apiError},
					}
					c.JSON(apiError.Status, apiErrors)
				case *APIErrors:
					apiErrors := err.(*APIErrors)
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
					panic(err)
				}
			}
		}()
		c.Next()
	}
}