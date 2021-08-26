// Code generated by entc, DO NOT EDIT.

package http

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/mailru/easyjson"
	"github.com/masseelch/elk/internal/unit/pets/ent"
	"github.com/masseelch/elk/internal/unit/pets/ent/category"
	"github.com/masseelch/elk/internal/unit/pets/ent/owner"
	"github.com/masseelch/elk/internal/unit/pets/ent/pet"
	"go.uber.org/zap"
)

// Pets fetches the ent.pets attached to the ent.Category
// identified by a given url-parameter from the database and renders it to the client.
func (h CategoryHandler) Pets(w http.ResponseWriter, r *http.Request) {
	l := h.log.With(zap.String("method", "Pets"))
	// ID is URL parameter.
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		l.Error("error getting id from url parameter", zap.String("id", chi.URLParam(r, "id")), zap.Error(err))
		BadRequest(w, "id must be an integer greater zero")
		return
	}
	// Create the query to fetch the pets attached to this category
	q := h.client.Category.Query().Where(category.ID(id)).QueryPets()
	page := 1
	if d := r.URL.Query().Get("page"); d != "" {
		page, err = strconv.Atoi(d)
		if err != nil {
			l.Info("error parsing query parameter 'page'", zap.String("page", d), zap.Error(err))
			BadRequest(w, "page must be an integer greater zero")
			return
		}
	}
	itemsPerPage := 30
	if d := r.URL.Query().Get("itemsPerPage"); d != "" {
		itemsPerPage, err = strconv.Atoi(d)
		if err != nil {
			l.Info("error parsing query parameter 'itemsPerPage'", zap.String("itemsPerPage", d), zap.Error(err))
			BadRequest(w, "itemsPerPage must be an integer greater zero")
			return
		}
	}
	es, err := q.Limit(itemsPerPage).Offset((page - 1) * itemsPerPage).All(r.Context())
	if err != nil {
		l.Error("error fetching pets from db", zap.Error(err))
		InternalServerError(w, nil)
		return
	}
	l.Info("pets rendered", zap.Int("amount", len(es)))
	easyjson.MarshalToHTTPResponseWriter(NewPet359800019Views(es), w)
}

// Pets fetches the ent.pets attached to the ent.Owner
// identified by a given url-parameter from the database and renders it to the client.
func (h OwnerHandler) Pets(w http.ResponseWriter, r *http.Request) {
	l := h.log.With(zap.String("method", "Pets"))
	// ID is URL parameter.
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		l.Error("error getting id from url parameter", zap.String("id", chi.URLParam(r, "id")), zap.Error(err))
		BadRequest(w, "id must be an integer greater zero")
		return
	}
	// Create the query to fetch the pets attached to this owner
	q := h.client.Owner.Query().Where(owner.ID(id)).QueryPets()
	page := 1
	if d := r.URL.Query().Get("page"); d != "" {
		page, err = strconv.Atoi(d)
		if err != nil {
			l.Info("error parsing query parameter 'page'", zap.String("page", d), zap.Error(err))
			BadRequest(w, "page must be an integer greater zero")
			return
		}
	}
	itemsPerPage := 30
	if d := r.URL.Query().Get("itemsPerPage"); d != "" {
		itemsPerPage, err = strconv.Atoi(d)
		if err != nil {
			l.Info("error parsing query parameter 'itemsPerPage'", zap.String("itemsPerPage", d), zap.Error(err))
			BadRequest(w, "itemsPerPage must be an integer greater zero")
			return
		}
	}
	es, err := q.Limit(itemsPerPage).Offset((page - 1) * itemsPerPage).All(r.Context())
	if err != nil {
		l.Error("error fetching pets from db", zap.Error(err))
		InternalServerError(w, nil)
		return
	}
	l.Info("pets rendered", zap.Int("amount", len(es)))
	easyjson.MarshalToHTTPResponseWriter(NewPet359800019Views(es), w)
}

// Category fetches the ent.category attached to the ent.Pet
// identified by a given url-parameter from the database and renders it to the client.
func (h PetHandler) Category(w http.ResponseWriter, r *http.Request) {
	l := h.log.With(zap.String("method", "Category"))
	// ID is URL parameter.
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		l.Error("error getting id from url parameter", zap.String("id", chi.URLParam(r, "id")), zap.Error(err))
		BadRequest(w, "id must be an integer greater zero")
		return
	}
	// Create the query to fetch the category attached to this pet
	q := h.client.Pet.Query().Where(pet.ID(id)).QueryCategory()
	page := 1
	if d := r.URL.Query().Get("page"); d != "" {
		page, err = strconv.Atoi(d)
		if err != nil {
			l.Info("error parsing query parameter 'page'", zap.String("page", d), zap.Error(err))
			BadRequest(w, "page must be an integer greater zero")
			return
		}
	}
	itemsPerPage := 30
	if d := r.URL.Query().Get("itemsPerPage"); d != "" {
		itemsPerPage, err = strconv.Atoi(d)
		if err != nil {
			l.Info("error parsing query parameter 'itemsPerPage'", zap.String("itemsPerPage", d), zap.Error(err))
			BadRequest(w, "itemsPerPage must be an integer greater zero")
			return
		}
	}
	es, err := q.Limit(itemsPerPage).Offset((page - 1) * itemsPerPage).All(r.Context())
	if err != nil {
		l.Error("error fetching categories from db", zap.Error(err))
		InternalServerError(w, nil)
		return
	}
	l.Info("categories rendered", zap.Int("amount", len(es)))
	easyjson.MarshalToHTTPResponseWriter(NewCategory4094953247Views(es), w)
}

// Owner fetches the ent.owner attached to the ent.Pet
// identified by a given url-parameter from the database and renders it to the client.
func (h PetHandler) Owner(w http.ResponseWriter, r *http.Request) {
	l := h.log.With(zap.String("method", "Owner"))
	// ID is URL parameter.
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		l.Error("error getting id from url parameter", zap.String("id", chi.URLParam(r, "id")), zap.Error(err))
		BadRequest(w, "id must be an integer greater zero")
		return
	}
	// Create the query to fetch the owner attached to this pet
	q := h.client.Pet.Query().Where(pet.ID(id)).QueryOwner()
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
			l.Error("could-not-read-pet", zap.Error(err), zap.Int("id", id))
			InternalServerError(w, nil)
		}
		return
	}
	l.Info("owner rendered", zap.Int("id", e.ID))
	easyjson.MarshalToHTTPResponseWriter(NewOwner139708381View(e), w)
}

// Friends fetches the ent.friends attached to the ent.Pet
// identified by a given url-parameter from the database and renders it to the client.
func (h PetHandler) Friends(w http.ResponseWriter, r *http.Request) {
	l := h.log.With(zap.String("method", "Friends"))
	// ID is URL parameter.
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		l.Error("error getting id from url parameter", zap.String("id", chi.URLParam(r, "id")), zap.Error(err))
		BadRequest(w, "id must be an integer greater zero")
		return
	}
	// Create the query to fetch the friends attached to this pet
	q := h.client.Pet.Query().Where(pet.ID(id)).QueryFriends()
	page := 1
	if d := r.URL.Query().Get("page"); d != "" {
		page, err = strconv.Atoi(d)
		if err != nil {
			l.Info("error parsing query parameter 'page'", zap.String("page", d), zap.Error(err))
			BadRequest(w, "page must be an integer greater zero")
			return
		}
	}
	itemsPerPage := 30
	if d := r.URL.Query().Get("itemsPerPage"); d != "" {
		itemsPerPage, err = strconv.Atoi(d)
		if err != nil {
			l.Info("error parsing query parameter 'itemsPerPage'", zap.String("itemsPerPage", d), zap.Error(err))
			BadRequest(w, "itemsPerPage must be an integer greater zero")
			return
		}
	}
	es, err := q.Limit(itemsPerPage).Offset((page - 1) * itemsPerPage).All(r.Context())
	if err != nil {
		l.Error("error fetching pets from db", zap.Error(err))
		InternalServerError(w, nil)
		return
	}
	l.Info("pets rendered", zap.Int("amount", len(es)))
	easyjson.MarshalToHTTPResponseWriter(NewPet359800019Views(es), w)
}