package handlers

import (
	"context"
	"errors"
	"net/http"

	"github.com/cauabernardino/bucket-list/db"
	"github.com/cauabernardino/bucket-list/models"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

var itemIDKey = "itemID"

// items handles the creation of routes for CRUD items
func items(router chi.Router) {
	router.Get("/", listItems)
	router.Post("/", newItem)
	router.Route("/{itemID}", func(router chi.Router) {
		router.Use(ItemContext)
		router.Get("/", getItem)
		router.Put("/", updateItem)
		router.Delete("/", deleteItem)
	})
}

// ItemContext extracts the item ID from request
func ItemContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		itemID := chi.URLParam(r, "itemID")
		if itemID == "" {
			render.Render(w, r, ErrorRenderer(errors.New("item ID is required")))
			return
		}

		ctx := context.WithValue(r.Context(), itemIDKey, itemID)
		next.ServeHTTP(w, r.WithContext(ctx))

	})
}

// newItem handles the creation route for a new item
func newItem(w http.ResponseWriter, r *http.Request) {
	item := &models.Item{}

	if err := render.Bind(r, item); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}

	if err := dbInstance.NewItem(item); err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}

	if err := render.Render(w, r, item); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}

// listItems handles the listing route for a all items
func listItems(w http.ResponseWriter, r *http.Request) {
	items, err := dbInstance.GetItems()
	if err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}

	if err := render.Render(w, r, items); err != nil {
		render.Render(w, r, ErrorRenderer(err))
	}
}

// getItem handles the search for a single item by ID
func getItem(w http.ResponseWriter, r *http.Request) {
	itemID := r.Context().Value(itemIDKey).(string)

	item, err := dbInstance.SearchItemByID(itemID)
	if err != nil {
		if err == db.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ErrorRenderer(err))
		}
		return
	}

	if err := render.Render(w, r, &item); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}

// deleteItem handles the deletion of a single item by ID
func deleteItem(w http.ResponseWriter, r *http.Request) {
	itemID := r.Context().Value(itemIDKey).(string)

	err := dbInstance.DeleteItem(itemID)
	if err != nil {
		if err == db.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ServerErrorRenderer(err))
		}
		return
	}
}

// deleteItem handles the update of a single item by ID
func updateItem(w http.ResponseWriter, r *http.Request) {
	itemID := r.Context().Value(itemIDKey).(string)
	itemData := models.Item{}

	if err := render.Bind(r, &itemData); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}

	item, err := dbInstance.UpdateItem(itemID, itemData)
	if err != nil {
		if err == db.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ServerErrorRenderer(err))
		}
		return
	}

	if err := render.Render(w, r, &item); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}
