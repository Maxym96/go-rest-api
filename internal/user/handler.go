package user

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
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
	router.GET(usersUrl, h.GetList)
	router.GET(userUrl, h.GetUserByID)
	router.POST(usersUrl, h.CreateUser)
	router.PUT(userUrl, h.UpdateUserByID)
	router.PATCH(userUrl, h.PartiallyUpdateUserByID)
	router.DELETE(userUrl, h.DeleteUserByID)

}
func (h *handler) GetList(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(200)
	h.logger.Info("this is list of Users")
	w.Write([]byte("this is list of Users"))
}
func (h *handler) GetUserByID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(200)
	w.Write([]byte("this is one of the user by ID"))
}
func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(201)
	w.Write([]byte("this is created user by id"))
}
func (h *handler) UpdateUserByID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(204)
	w.Write([]byte("this is fully updated user by ID"))
}
func (h *handler) PartiallyUpdateUserByID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(204)
	w.Write([]byte("this is partially updated user by ID"))
}
func (h *handler) DeleteUserByID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(204)
	w.Write([]byte("this is deleted user by ID"))
}
