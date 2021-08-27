// Code generated by entc, DO NOT EDIT.

package http

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/mailru/easyjson"
	"github.com/masseelch/elk/internal/openapi/ent"
	"github.com/masseelch/elk/internal/openapi/ent/category"
	"github.com/masseelch/elk/internal/openapi/ent/owner"
	"github.com/masseelch/elk/internal/openapi/ent/pet"
	"go.uber.org/zap"
)

// Read fetches the ent.Category identified by a given url-parameter from the
// database and renders it to the client.
func (h *CategoryHandler) Read(w http.ResponseWriter, r *http.Request) {
	l := h.log.With(zap.String("method", "Read"))
	// ID is URL parameter.
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		l.Error("error getting id from url parameter", zap.String("id", chi.URLParam(r, "id")), zap.Error(err))
		BadRequest(w, "id must be an integer greater zero")
		return
	}
	// Create the query to fetch the Category
	q := h.client.Category.Query().Where(category.ID(id))
	e, err := q.Only(r.Context())
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			msg := stripEntError(err)
			l.Info(msg, zap.Error(err), zap.Int("id", id))
			NotFound(w, msg)
		case ent.IsNotSingular(err):
			msg := stripEntError(err)
			l.Error(msg, zap.Error(err), zap.Int("id", id))
			BadRequest(w, msg)
		default:
			l.Error("could not read category", zap.Error(err), zap.Int("id", id))
			InternalServerError(w, nil)
		}
		return
	}
	l.Info("category rendered", zap.Int("id", id))
	easyjson.MarshalToHTTPResponseWriter(NewCategory4094953247View(e), w)
}

// Read fetches the ent.Owner identified by a given url-parameter from the
// database and renders it to the client.
func (h *OwnerHandler) Read(w http.ResponseWriter, r *http.Request) {
	l := h.log.With(zap.String("method", "Read"))
	// ID is URL parameter.
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		l.Error("error getting id from url parameter", zap.String("id", chi.URLParam(r, "id")), zap.Error(err))
		BadRequest(w, "id must be an integer greater zero")
		return
	}
	// Create the query to fetch the Owner
	q := h.client.Owner.Query().Where(owner.ID(id))
	e, err := q.Only(r.Context())
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			msg := stripEntError(err)
			l.Info(msg, zap.Error(err), zap.Int("id", id))
			NotFound(w, msg)
		case ent.IsNotSingular(err):
			msg := stripEntError(err)
			l.Error(msg, zap.Error(err), zap.Int("id", id))
			BadRequest(w, msg)
		default:
			l.Error("could not read owner", zap.Error(err), zap.Int("id", id))
			InternalServerError(w, nil)
		}
		return
	}
	l.Info("owner rendered", zap.Int("id", id))
	easyjson.MarshalToHTTPResponseWriter(NewOwner139708381View(e), w)
}

// Read fetches the ent.Pet identified by a given url-parameter from the
// database and renders it to the client.
func (h *PetHandler) Read(w http.ResponseWriter, r *http.Request) {
	l := h.log.With(zap.String("method", "Read"))
	// ID is URL parameter.
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		l.Error("error getting id from url parameter", zap.String("id", chi.URLParam(r, "id")), zap.Error(err))
		BadRequest(w, "id must be an integer greater zero")
		return
	}
	// Create the query to fetch the Pet
	q := h.client.Pet.Query().Where(pet.ID(id))
	// Eager load edges that are required on read operation.
	q.WithOwner().WithFriends(func(q *ent.PetQuery) {
		q.WithOwner().WithFriends(func(q *ent.PetQuery) {
			q.WithOwner().WithFriends(func(q *ent.PetQuery) {
				q.WithOwner()
			})
		})
	})
	e, err := q.Only(r.Context())
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			msg := stripEntError(err)
			l.Info(msg, zap.Error(err), zap.Int("id", id))
			NotFound(w, msg)
		case ent.IsNotSingular(err):
			msg := stripEntError(err)
			l.Error(msg, zap.Error(err), zap.Int("id", id))
			BadRequest(w, msg)
		default:
			l.Error("could not read pet", zap.Error(err), zap.Int("id", id))
			InternalServerError(w, nil)
		}
		return
	}
	l.Info("pet rendered", zap.Int("id", id))
	easyjson.MarshalToHTTPResponseWriter(NewPet1876743790View(e), w)
}
