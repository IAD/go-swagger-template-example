package service

import (
	"fmt"
	"strings"
	"sync"

	"github.com/IAD/go-swagger-template-example/internal/server/models"
	"github.com/IAD/go-swagger-template-example/internal/server/restapi/operations/todos"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
)

type Item struct {
	Completed   bool
	Description string
}

type Handlers struct {
	items map[int64]Item
	lock  sync.RWMutex
}

func NewHandlers() *Handlers {
	return &Handlers{
		items: make(map[int64]Item),
	}
}

// Handler for POST /
func (h *Handlers) AddOneHandler(
	params *todos.AddOneParams,
	addOneCreated todos.NewAddOneCreatedFunc,
	addOneInternalServerError todos.NewAddOneInternalServerErrorFunc,
) middleware.Responder {
	h.lock.Lock()
	defer h.lock.Unlock()

	item := Item{
		Completed:   params.Body.Completed,
		Description: swag.StringValue(params.Body.Description),
	}

	id := params.Body.ID

	h.items[id] = item

	return addOneCreated().WithPayload(&models.Item{
		Completed:   item.Completed,
		Description: swag.String(item.Description),
		ID:          id,
	})
}

// Handler for DELETE /{id}
func (h *Handlers) DestroyOneHandler(
	params *todos.DestroyOneParams,
	destroyOneNoContent todos.NewDestroyOneNoContentFunc,
	destroyOneNotFound todos.NewDestroyOneNotFoundFunc,
	destroyOneInternalServerError todos.NewDestroyOneInternalServerErrorFunc,
) middleware.Responder {
	h.lock.Lock()
	defer h.lock.Unlock()

	id := params.ID

	if _, found := h.items[id]; !found {
		return destroyOneNotFound().WithErr(fmt.Errorf("not found"))
	}

	delete(h.items, id)

	return destroyOneNoContent()
}

// Handler for GET /
func (h *Handlers) FindHandler(
	params *todos.FindParams,
	findOK todos.NewFindOKFunc,
	findNotFound todos.NewFindNotFoundFunc,
	findInternalServerError todos.NewFindInternalServerErrorFunc,
) middleware.Responder {
	h.lock.RLock()
	defer h.lock.RUnlock()

	searchCriteria := params.Words
	if len(searchCriteria) == 0 {
		return findInternalServerError().WithErr(fmt.Errorf("unable to search for an empty list"))
	}

	limit := params.Limit
	if limit <= 0 {
		return findInternalServerError().WithErr(fmt.Errorf("unable to search for not positive limit"))
	}

	result := make([]*models.Item, 0, limit)

	foundItems := int64(0)

SEARCH:
	for id, item := range h.items {
		for _, word := range searchCriteria {
			if strings.Contains(item.Description, word) {
				result = append(result, &models.Item{
					Completed:   item.Completed,
					Description: swag.String(item.Description),
					ID:          id,
				})
				foundItems++
				if foundItems >= limit {
					break SEARCH
				}
			}
		}
	}

	if foundItems == 0 {
		return findNotFound().WithErr(fmt.Errorf("not found"))
	}

	return findOK().WithPayload(result)
}

// Handler for PUT /{id}
func (h *Handlers) UpdateOneHandler(
	params *todos.UpdateOneParams,
	updateOneOK todos.NewUpdateOneOKFunc,
	updateOneNotFound todos.NewUpdateOneNotFoundFunc,
	updateOneInternalServerError todos.NewUpdateOneInternalServerErrorFunc,
) middleware.Responder {
	h.lock.Lock()
	defer h.lock.Unlock()

	id := params.ID

	item, found := h.items[id]
	if !found {
		return updateOneNotFound().WithErr(fmt.Errorf("not found"))
	}

	item.Description = swag.StringValue(params.Body.Description)
	item.Completed = params.Body.Completed

	h.items[id] = item

	return updateOneOK().WithPayload(&models.Item{
		Completed:   item.Completed,
		Description: swag.String(item.Description),
		ID:          id,
	})
}
