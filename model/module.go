package model

type Module struct {
	BaseFields
	ModuleUpdate
}

type ModuleUpdate struct {
	Pid      int    `json:"pid"`
	Spid     string `json:"spid"`
	AuthorId int    `json:"author_id"`
	Title    string `json:"title"`
}

func (model Module) UpdateStruct() interface{} {
	return ModuleUpdate{}
}
