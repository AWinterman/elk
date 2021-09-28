// Code generated by entc, DO NOT EDIT.

package http

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/mailru/easyjson"
	"github.com/masseelch/elk/internal/pets/ent"
	badge "github.com/masseelch/elk/internal/pets/ent/badge"
	pet "github.com/masseelch/elk/internal/pets/ent/pet"
	playgroup "github.com/masseelch/elk/internal/pets/ent/playgroup"
	toy "github.com/masseelch/elk/internal/pets/ent/toy"
	"go.uber.org/zap"
)

// Update updates a given ent.Badge and saves the changes to the database.
func (h BadgeHandler) Update(w http.ResponseWriter, r *http.Request) {
	l := h.log.With(zap.String("method", "Update"))
	// ID is URL parameter.
	id64, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 0)
	if err != nil {
		l.Error("error getting id from url parameter", zap.String("id", chi.URLParam(r, "id")), zap.Error(err))
		BadRequest(w, "id must be an integer greater zero")
		return
	}
	id := uint32(id64)
	// Get the post data.
	var d BadgeUpdateRequest
	if err := easyjson.UnmarshalFromReader(r.Body, &d); err != nil {
		l.Error("error decoding json", zap.Error(err))
		BadRequest(w, "invalid json string")
		return
	}
	// Save the data.
	b := h.client.Badge.UpdateOneID(id)
	if d.Color != nil {
		b.SetColor(*d.Color)
	}
	if d.Material != nil {
		b.SetMaterial(*d.Material)
	}
	if d.Wearer != nil {
		b.SetWearerID(*d.Wearer)

	}
	// Store in database.
	e, err := b.Save(r.Context())
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			msg := stripEntError(err)
			l.Info(msg, zap.Error(err), zap.Uint32("id", id))
			NotFound(w, msg)
		case ent.IsNotSingular(err):
			msg := stripEntError(err)
			l.Error(msg, zap.Error(err), zap.Uint32("id", id))
			BadRequest(w, msg)
		default:
			l.Error("could-not-update-badge", zap.Error(err), zap.Uint32("id", id))
			InternalServerError(w, nil)
		}
		return
	}
	// Reload entry.
	q := h.client.Badge.Query().Where(badge.ID(e.ID))
	e, err = q.Only(r.Context())
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			msg := stripEntError(err)
			l.Info(msg, zap.Error(err), zap.Uint32("id", id))
			NotFound(w, msg)
		case ent.IsNotSingular(err):
			msg := stripEntError(err)
			l.Error(msg, zap.Error(err), zap.Uint32("id", id))
			BadRequest(w, msg)
		default:
			l.Error("could-not-read-badge", zap.Error(err), zap.Uint32("id", id))
			InternalServerError(w, nil)
		}
		return
	}
	l.Info("badge rendered", zap.Uint32("id", id))
	easyjson.MarshalToHTTPResponseWriter(NewBadge2492344257View(e), w)
}

// Update updates a given ent.Pet and saves the changes to the database.
func (h PetHandler) Update(w http.ResponseWriter, r *http.Request) {
	l := h.log.With(zap.String("method", "Update"))
	// ID is URL parameter.
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		l.Error("error getting id from url parameter", zap.String("id", chi.URLParam(r, "id")), zap.Error(err))
		BadRequest(w, "id must be an integer")
		return
	}
	// Get the post data.
	var d PetUpdateRequest
	if err := easyjson.UnmarshalFromReader(r.Body, &d); err != nil {
		l.Error("error decoding json", zap.Error(err))
		BadRequest(w, "invalid json string")
		return
	}
	// Validate the data.
	errs := make(map[string]string)
	if d.Height == nil {
		errs["height"] = `missing required field: "height"`
	} else if err := pet.HeightValidator(*d.Height); err != nil {
		errs["height"] = strings.TrimPrefix(err.Error(), "pet: ")
	}
	if d.Weight != nil {
		if err := pet.WeightValidator(*d.Weight); err != nil {
			errs["weight"] = strings.TrimPrefix(err.Error(), "pet: ")
		}
	}
	if d.Castrated == nil {
		errs["castrated"] = `missing required field: "castrated"`
	}
	if d.Name == nil {
		errs["name"] = `missing required field: "name"`
	} else if err := pet.NameValidator(*d.Name); err != nil {
		errs["name"] = strings.TrimPrefix(err.Error(), "pet: ")
	}
	if d.Badge == nil {
		errs["badge"] = `missing required edge: "badge"`
	}
	if len(errs) > 0 {
		l.Info("validation failed", zapFields(errs)...)
		BadRequest(w, errs)
		return
	}
	// Save the data.
	b := h.client.Pet.UpdateOneID(id)
	if d.Height != nil {
		b.SetHeight(*d.Height)
	}
	if d.Weight != nil {
		b.SetWeight(*d.Weight)
	}
	if d.Castrated != nil {
		b.SetCastrated(*d.Castrated)
	}
	if d.Name != nil {
		b.SetName(*d.Name)
	}
	if d.Birthday != nil {
		b.SetBirthday(*d.Birthday)
	}
	if d.Nicknames != nil {
		b.SetNicknames(*d.Nicknames)
	}
	if d.Chip != nil {
		b.SetChip(*d.Chip)
	}
	if d.Badge != nil {
		b.SetBadgeID(*d.Badge)

	}
	if d.Protege != nil {
		b.SetProtegeID(*d.Protege)

	}
	if d.Mentor != nil {
		b.SetMentorID(*d.Mentor)

	}
	if d.Spouse != nil {
		b.SetSpouseID(*d.Spouse)

	}
	if d.Toys != nil {
		b.ClearToys().AddToyIDs(d.Toys...)
	}
	if d.Parent != nil {
		b.SetParentID(*d.Parent)

	}
	if d.Children != nil {
		b.ClearChildren().AddChildIDs(d.Children...)
	}
	if d.PlayGroups != nil {
		b.ClearPlayGroups().AddPlayGroupIDs(d.PlayGroups...)
	}
	if d.Friends != nil {
		b.ClearFriends().AddFriendIDs(d.Friends...)
	}
	// Store in database.
	e, err := b.Save(r.Context())
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
			l.Error("could-not-update-pet", zap.Error(err), zap.Int("id", id))
			InternalServerError(w, nil)
		}
		return
	}
	// Reload entry.
	q := h.client.Pet.Query().Where(pet.ID(e.ID))
	e, err = q.Only(r.Context())
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
	l.Info("pet rendered", zap.Int("id", id))
	easyjson.MarshalToHTTPResponseWriter(NewPet340207500View(e), w)
}

// Update updates a given ent.PlayGroup and saves the changes to the database.
func (h PlayGroupHandler) Update(w http.ResponseWriter, r *http.Request) {
	l := h.log.With(zap.String("method", "Update"))
	// ID is URL parameter.
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		l.Error("error getting id from url parameter", zap.String("id", chi.URLParam(r, "id")), zap.Error(err))
		BadRequest(w, "id must be an integer")
		return
	}
	// Get the post data.
	var d PlayGroupUpdateRequest
	if err := easyjson.UnmarshalFromReader(r.Body, &d); err != nil {
		l.Error("error decoding json", zap.Error(err))
		BadRequest(w, "invalid json string")
		return
	}
	// Save the data.
	b := h.client.PlayGroup.UpdateOneID(id)
	if d.Title != nil {
		b.SetTitle(*d.Title)
	}
	if d.Description != nil {
		b.SetDescription(*d.Description)
	}
	if d.Weekday != nil {
		b.SetWeekday(*d.Weekday)
	}
	if d.Participants != nil {
		b.ClearParticipants().AddParticipantIDs(d.Participants...)
	}
	// Store in database.
	e, err := b.Save(r.Context())
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
			l.Error("could-not-update-play-group", zap.Error(err), zap.Int("id", id))
			InternalServerError(w, nil)
		}
		return
	}
	// Reload entry.
	q := h.client.PlayGroup.Query().Where(playgroup.ID(e.ID))
	e, err = q.Only(r.Context())
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
			l.Error("could-not-read-play-group", zap.Error(err), zap.Int("id", id))
			InternalServerError(w, nil)
		}
		return
	}
	l.Info("play-group rendered", zap.Int("id", id))
	easyjson.MarshalToHTTPResponseWriter(NewPlayGroup3432834655View(e), w)
}

// Update updates a given ent.Toy and saves the changes to the database.
func (h ToyHandler) Update(w http.ResponseWriter, r *http.Request) {
	l := h.log.With(zap.String("method", "Update"))
	// ID is URL parameter.
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		l.Error("error getting id from url parameter", zap.String("id", chi.URLParam(r, "id")), zap.Error(err))
		BadRequest(w, "id must be a valid UUID")
		return
	}
	// Get the post data.
	var d ToyUpdateRequest
	if err := easyjson.UnmarshalFromReader(r.Body, &d); err != nil {
		l.Error("error decoding json", zap.Error(err))
		BadRequest(w, "invalid json string")
		return
	}
	// Save the data.
	b := h.client.Toy.UpdateOneID(uuid.UUID(id))
	if d.Color != nil {
		b.SetColor(*d.Color)
	}
	if d.Material != nil {
		b.SetMaterial(*d.Material)
	}
	if d.Title != nil {
		b.SetTitle(*d.Title)
	}
	if d.Owner != nil {
		b.SetOwnerID(*d.Owner)

	}
	// Store in database.
	e, err := b.Save(r.Context())
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			msg := stripEntError(err)
			l.Info(msg, zap.Error(err), zap.String("id", id.String()))
			NotFound(w, msg)
		case ent.IsNotSingular(err):
			msg := stripEntError(err)
			l.Error(msg, zap.Error(err), zap.String("id", id.String()))
			BadRequest(w, msg)
		default:
			l.Error("could-not-update-toy", zap.Error(err), zap.String("id", id.String()))
			InternalServerError(w, nil)
		}
		return
	}
	// Reload entry.
	q := h.client.Toy.Query().Where(toy.ID(e.ID))
	e, err = q.Only(r.Context())
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			msg := stripEntError(err)
			l.Info(msg, zap.Error(err), zap.String("id", id.String()))
			NotFound(w, msg)
		case ent.IsNotSingular(err):
			msg := stripEntError(err)
			l.Error(msg, zap.Error(err), zap.String("id", id.String()))
			BadRequest(w, msg)
		default:
			l.Error("could-not-read-toy", zap.Error(err), zap.String("id", id.String()))
			InternalServerError(w, nil)
		}
		return
	}
	l.Info("toy rendered", zap.String("id", id.String()))
	easyjson.MarshalToHTTPResponseWriter(NewToy36157710View(e), w)
}
