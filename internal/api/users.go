package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/Gabriel-Valin/products-api/internal/middlewares"
	"github.com/Gabriel-Valin/products-api/internal/users"
)

type UsersHandler struct {
	service users.Service
}

func NewUsersHandler(service users.Service) *UsersHandler {
	return &UsersHandler{
		service: service,
	}
}

func (h *UsersHandler) Users(
	w http.ResponseWriter,
	r *http.Request,
) {
	switch r.Method {
	case http.MethodGet:
		h.list(w, r)

	case http.MethodPost:
		h.create(w, r)

	default:
		_ = writeError(
			w,
			http.StatusMethodNotAllowed,
			http.StatusText(http.StatusMethodNotAllowed),
		)
	}
}

func (h *UsersHandler) list(
	w http.ResponseWriter,
	r *http.Request,
) {
	usersList, err := h.service.List(r.Context())
	if err != nil {
		middlewares.GetLogger(r.Context()).Error(
			"failed to list users",
			"error", err,
		)

		_ = writeError(
			w,
			statusFromError(err),
			messageFromError(err),
		)
		return
	}

	if err := writeJSON(
		w,
		http.StatusOK,
		usersList,
	); err != nil {
		middlewares.GetLogger(r.Context()).Error(
			"failed to write response",
			"error", err,
		)
	}
}

func (h *UsersHandler) create(
	w http.ResponseWriter,
	r *http.Request,
) {
	var req users.CreateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		_ = writeError(
			w,
			http.StatusBadRequest,
			"invalid request body",
		)
		return
	}

	user, err := h.service.Create(
		r.Context(),
		req,
	)
	if err != nil {
		middlewares.GetLogger(r.Context()).Error(
			"failed to create user",
			"error", err,
		)

		_ = writeError(
			w,
			statusFromError(err),
			messageFromError(err),
		)
		return
	}

	if err := writeJSON(
		w,
		http.StatusCreated,
		user,
	); err != nil {
		middlewares.GetLogger(r.Context()).Error(
			"failed to write response",
			"error", err,
		)
	}
}

func (h *UsersHandler) GetByID(
	w http.ResponseWriter,
	r *http.Request,
) {
	if r.Method != http.MethodGet {
		_ = writeError(
			w,
			http.StatusMethodNotAllowed,
			http.StatusText(http.StatusMethodNotAllowed),
		)
		return
	}

	idText := strings.TrimPrefix(r.URL.Path, "/users/")
	if idText == "" {
		_ = writeError(
			w,
			http.StatusBadRequest,
			"user ID is required",
		)
		return
	}

	id, err := strconv.Atoi(idText)
	if err != nil {
		_ = writeError(
			w,
			http.StatusBadRequest,
			"invalid user ID",
		)
		return
	}

	user, err := h.service.GetByID(r.Context(), id)
	if err != nil {
		middlewares.GetLogger(r.Context()).Error(
			"failed to get user",
			"error", err,
			"user_id", id,
		)

		_ = writeError(
			w,
			statusFromError(err),
			messageFromError(err),
		)
		return
	}

	if err := writeJSON(
		w,
		http.StatusOK,
		user,
	); err != nil {
		middlewares.GetLogger(r.Context()).Error(
			"failed to write response",
			"error", err,
		)
	}
}
