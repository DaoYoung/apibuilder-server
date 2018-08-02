package model

type ApiNote struct {
	BaseFields
	ApiId         int    `json:"api_id"`
	AuthorId      int    `json:"author_id"`
	Status        int    `json:"status"`
	NoteUpdate
}

type NoteUpdate struct {
	Fkey          string `json:"fkey"`
	FkeyParent    string `json:"fkey_parent"`
	FkeyToken     string `json:"fkey_token"`
	ModelId       int    `json:"model_id"`
	ParentModelId int    `json:"parent_model_id"`
} 

func (model *ApiNote) UpdateStruct() interface{} {
	return NoteUpdate{}
}


func CreateApiNote(chs []byte, msg string, taskId int, apiId int, authorId int) interface{} {
	commitInfo := new(ApiCommit)
	commitInfo.Changes = chs
	commitInfo.CommitMessage = msg
	commitInfo.ApiId = apiId
	commitInfo.TaskId = taskId
	commitInfo.AuthorId = authorId
	return Create(commitInfo)
}
 