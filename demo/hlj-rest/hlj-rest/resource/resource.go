package resource

import (
	"net/http"
	"hlj-rest/rest"
)

var json = rest.GetJson()

type Resource interface {
	GetList(offset, limit int, sort string, count bool, conditions ...interface{}) (list []Resource, total int)
	GetOne(id int) Resource
	Insert() error
	Update() error
	Delete() error
}

type Authorizator func(resource Resource) bool

type Handler struct {
	resource       Resource
	responseWriter rest.ResponseWriter
	request        *rest.Request
}

func NewHandler(resource Resource, w rest.ResponseWriter, r *rest.Request) *Handler {
	return &Handler{
		resource:       resource,
		responseWriter: w,
		request:        r,
	}
}

func (h *Handler) List(conditions ...interface{}) {
	sort := h.request.QueryParam("sort").String()
	if sort == "" {
		sort = "-id"
	}
	page := h.request.QueryParam("page").Default("1").Int()
	if page <=0 {
		page = 1
	}
	pageSize := h.request.QueryParam("per_page").Default("10").Int()
	offset := (page - 1) * pageSize

	if conditions == nil {
		conditions = []interface{}{}
	}
	entities, total := h.resource.GetList(offset, pageSize, sort, true, conditions...)
	relationNodes := ParseResult(h.request.Request)
	list, _ := LoadRelation(entities, relationNodes)
	meta := make(map[string]interface{})
	meta["Pagination-Total"] = total
	meta["Pagination-PageSize"] = pageSize

	rest.Ok(h.responseWriter, list, meta)
}

func (h *Handler) One(id int, authorizator Authorizator) {
	entity := h.resource.GetOne(id)
	if entity == nil {
		rest.NotFound(h.responseWriter)
		return
	}

	relationNodes := ParseResult(h.request.Request)
	one, _ := LoadRelation(entity, relationNodes)

	rest.Ok(h.responseWriter, one, nil)
}

func (h *Handler) Add(authorizator Authorizator) {
	entity := h.resource
	if err := h.request.DecodeJsonPayload(entity); err != nil {
		rest.Error(h.responseWriter, http.StatusUnprocessableEntity, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	if err := entity.Insert(); err != nil {
		rest.ServerUnavailabe(h.responseWriter)
		return
	}

	rest.Created(h.responseWriter, entity, nil)
}

func (h *Handler) Update(id int, authorizator Authorizator) {
	entity := h.resource.GetOne(id)
	if entity == nil {
		rest.NotFound(h.responseWriter)
		return
	}

	if authorizator != nil && !authorizator(entity) {
		rest.Unauthorized(h.responseWriter)
		return
	}

	if err := h.request.DecodeJsonPayload(entity); err != nil {
		rest.Error(h.responseWriter, http.StatusUnprocessableEntity, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	if err := entity.Update(); err != nil {
		rest.ServerUnavailabe(h.responseWriter)
		return
	}

	rest.Ok(h.responseWriter, entity, nil)
}

func (h *Handler) Delete(id int, authorizator Authorizator) {
	entity := h.resource.GetOne(id)
	if entity == nil {
		rest.NotFound(h.responseWriter)
		return
	}

	if err := entity.Delete(); err != nil {
		rest.ServerUnavailabe(h.responseWriter)
		return
	}

	rest.NoContent(h.responseWriter)
}
