// Code generated by entc, DO NOT EDIT.

package http

import (
	"net/http"
	"strconv"

	"github.com/liip/sheriff"
	"github.com/masseelch/render"
	"go.uber.org/zap"
)

// Read fetches the Category model identified by a given url-parameter from the
// database and returns it to the client.
func (h *CategoryHandler) List(w http.ResponseWriter, r *http.Request) {
	l := h.log.With(zap.String("method", "List"))
	q := h.client.Category.Query()
	var err error
	page := 1
	if d := r.URL.Query().Get("page"); d != "" {
		page, err = strconv.Atoi(d)
		if err != nil {
			l.Info("error parsing query parameter 'page'", zap.String("page", d), zap.Error(err))
			render.BadRequest(w, r, "page must be an integer greater zero")
			return
		}
	}
	itemsPerPage := 30
	if d := r.URL.Query().Get("itemsPerPage"); d != "" {
		itemsPerPage, err = strconv.Atoi(d)
		if err != nil {
			l.Info("error parsing query parameter 'itemsPerPage'", zap.String("itemsPerPage", d), zap.Error(err))
			render.BadRequest(w, r, "itemsPerPage must be an integer greater zero")
			return
		}
	}
	es, err := q.Limit(itemsPerPage).Offset((page - 1) * itemsPerPage).All(r.Context())
	if err != nil {
		l.Error("error fetching categories from db", zap.Error(err))
		render.InternalServerError(w, r, nil)
		return
	}
	d, err := sheriff.Marshal(&sheriff.Options{
		IncludeEmptyTag: true,
		Groups:          []string{"category"},
	}, es)
	if err != nil {
		l.Error("serialization error", zap.Error(err))
		render.InternalServerError(w, r, nil)
		return
	}
	l.Info("categories rendered", zap.Int("amount", len(es)))
	render.OK(w, r, d)
}

// Read fetches the Owner model identified by a given url-parameter from the
// database and returns it to the client.
func (h *OwnerHandler) List(w http.ResponseWriter, r *http.Request) {
	l := h.log.With(zap.String("method", "List"))
	q := h.client.Owner.Query()
	var err error
	page := 1
	if d := r.URL.Query().Get("page"); d != "" {
		page, err = strconv.Atoi(d)
		if err != nil {
			l.Info("error parsing query parameter 'page'", zap.String("page", d), zap.Error(err))
			render.BadRequest(w, r, "page must be an integer greater zero")
			return
		}
	}
	itemsPerPage := 30
	if d := r.URL.Query().Get("itemsPerPage"); d != "" {
		itemsPerPage, err = strconv.Atoi(d)
		if err != nil {
			l.Info("error parsing query parameter 'itemsPerPage'", zap.String("itemsPerPage", d), zap.Error(err))
			render.BadRequest(w, r, "itemsPerPage must be an integer greater zero")
			return
		}
	}
	es, err := q.Limit(itemsPerPage).Offset((page - 1) * itemsPerPage).All(r.Context())
	if err != nil {
		l.Error("error fetching owners from db", zap.Error(err))
		render.InternalServerError(w, r, nil)
		return
	}
	d, err := sheriff.Marshal(&sheriff.Options{
		IncludeEmptyTag: true,
		Groups:          []string{"owner"},
	}, es)
	if err != nil {
		l.Error("serialization error", zap.Error(err))
		render.InternalServerError(w, r, nil)
		return
	}
	l.Info("owners rendered", zap.Int("amount", len(es)))
	render.OK(w, r, d)
}

// Read fetches the Pet model identified by a given url-parameter from the
// database and returns it to the client.
func (h *PetHandler) List(w http.ResponseWriter, r *http.Request) {
	l := h.log.With(zap.String("method", "List"))
	q := h.client.Pet.Query()
	var err error
	page := 1
	if d := r.URL.Query().Get("page"); d != "" {
		page, err = strconv.Atoi(d)
		if err != nil {
			l.Info("error parsing query parameter 'page'", zap.String("page", d), zap.Error(err))
			render.BadRequest(w, r, "page must be an integer greater zero")
			return
		}
	}
	itemsPerPage := 30
	if d := r.URL.Query().Get("itemsPerPage"); d != "" {
		itemsPerPage, err = strconv.Atoi(d)
		if err != nil {
			l.Info("error parsing query parameter 'itemsPerPage'", zap.String("itemsPerPage", d), zap.Error(err))
			render.BadRequest(w, r, "itemsPerPage must be an integer greater zero")
			return
		}
	}
	es, err := q.Limit(itemsPerPage).Offset((page - 1) * itemsPerPage).All(r.Context())
	if err != nil {
		l.Error("error fetching pets from db", zap.Error(err))
		render.InternalServerError(w, r, nil)
		return
	}
	d, err := sheriff.Marshal(&sheriff.Options{
		IncludeEmptyTag: true,
		Groups:          []string{"pet"},
	}, es)
	if err != nil {
		l.Error("serialization error", zap.Error(err))
		render.InternalServerError(w, r, nil)
		return
	}
	l.Info("pets rendered", zap.Int("amount", len(es)))
	render.OK(w, r, d)
}
