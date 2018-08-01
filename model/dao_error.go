package model

type DaoErrors struct {
	Errors      []*DaoError `json:"errors"`
}

func (errors *DaoErrors) Status() int {
	return errors.Errors[0].Status
}

type DaoError struct {
	Status      int         `json:"status"`
	Code        string      `json:"code"`
	Title       string      `json:"title"`
	Details     string      `json:"details"`
}

func NewDaoError(status int, code string, title string, details string) *DaoError {
	return &DaoError{
		Status:     status,
		Code:       code,
		Title:      title,
		Details:    details,
	}
}

func NotFoundDaoError(err error) *DaoError {
	return NewDaoError(400, "not_found", "Not Found", err.Error())
}
func NotExistDaoError(err error) *DaoError {
	return NewDaoError(400, "not_exist", "Not Exist", err.Error())
}
func QueryDaoError(err error) *DaoError {
	return NewDaoError(400, "db_query", "DB query error", err.Error())
}
