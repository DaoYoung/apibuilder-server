package model

type ApiCommit struct {
	BaseFields
	TaskId        int
	ApiId         int
	AuthorId      int
	CommitMessage string
	Changes       JSON
}

func (model *ApiCommit) UpdateStruct() interface{} {
	return nil
}

func (model *ApiCommit) InitDao() *Dao {
	dao := &Dao{}
	dao.MainResource = model
	dao.SliceResource = &[]ApiCommit{}
	return dao
}

func CreateCommit(chs []byte, msg string, taskId int , apiId int, authorId int) interface{} {
	commitInfo := new(ApiCommit)
	commitInfo.Changes = chs
	commitInfo.CommitMessage = msg
	commitInfo.ApiId = apiId
	commitInfo.TaskId = taskId
	commitInfo.AuthorId = authorId
	return (&(ApiCommit{})).InitDao().Create(commitInfo)
}

type CommitChange struct {
	Before interface{} `json:"before"`
	After interface{} `json:"after"`
}
type CommitChangeJson struct {
	ChangeJson map[string]CommitChange `json:"change_json"`
}