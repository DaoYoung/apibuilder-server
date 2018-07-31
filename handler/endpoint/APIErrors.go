package endpoint

import "github.com/gin-gonic/gin"

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
var (
	ErrDatabase         = NewAPIError(500, "database_error", "Database Error", "An unknown error occurred.", "")
	ErrInvalidSet       = NewAPIError(404, "invalid_set", "Invalid Set", "The set you requested does not exist.", "")
	ErrInvalidGroup     = NewAPIError(404, "invalid_group", "Invalid Group", "The group you requested does not exist.", "")
)

func Recovery() gin.HandlerFunc {
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
				default:
					panic(err)
				}
			}
		}()

		c.Next()
	}
}