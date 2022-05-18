package user

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"rest-app/internal/apperror"
	"rest-app/internal/handlers"
	"rest-app/package/logging"
)

var _ handlers.Handler = &handler{}

const (
	usersUrl = "/users"
	userUrl  = "/user/:id"
)

type handler struct {
	logger *logging.Logger
}

func NewHandler(logger *logging.Logger) handlers.Handler {
	return &handler{
		logger: logger,
	}
}

func (h *handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodGet, usersUrl, apperror.MiddleWare(h.GetList))
	router.HandlerFunc(http.MethodGet, userUrl, apperror.MiddleWare(h.GetUserByID))
	router.HandlerFunc(http.MethodPost, usersUrl, apperror.MiddleWare(h.CreateUser))
	router.HandlerFunc(http.MethodPut, userUrl, apperror.MiddleWare(h.UpdateUserByID))
	router.HandlerFunc(http.MethodPatch, userUrl, apperror.MiddleWare(h.PartiallyUpdateUserByID))
	router.HandlerFunc(http.MethodDelete, userUrl, apperror.MiddleWare(h.DeleteUserByID))

}
func (h *handler) GetList(w http.ResponseWriter, r *http.Request) error {
	return apperror.ErrorNotFound
}
func (h *handler) GetUserByID(w http.ResponseWriter, r *http.Request) error {
	return apperror.NewAppError(nil, "test", "test", "T-123")
}
func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) error {
	return fmt.Errorf("this is API error Created User")
}
func (h *handler) UpdateUserByID(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(204)
	w.Write([]byte("this is fully updated user by ID"))

	return nil
}
func (h *handler) PartiallyUpdateUserByID(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(204)
	w.Write([]byte("this is partially updated user by ID"))

	return nil
}
func (h *handler) DeleteUserByID(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(204)
	w.Write([]byte("this is deleted user by ID"))

	return nil
}
