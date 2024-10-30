package user

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type (
	Controller func(w http.ResponseWriter, r *http.Request) //se define un Type como Controler que sera una func con response Wrtiter y una request

	Endpoints struct {
		Create Controller
		GetAll Controller
	}
	CreateReq struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
	}
)

func MakeEndpoints(ctx context.Context, s Service) Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			GetAllUser(ctx, s, w)

		case http.MethodPost:
			decode := json.NewDecoder(r.Body)
			var structUser CreateReq // la estructura que tenemos que revisar es CreateReq
			if err := decode.Decode(&structUser); err != nil {
				MsgResponse(w, http.StatusBadRequest, err.Error())
				return
			}
			PostUser(ctx, s, w, structUser)

		default:
			InvalidMethod(w)

		}
	}
}

func GetAllUser(ctx context.Context, s Service, w http.ResponseWriter) {
	users, err := s.GetAll(ctx)
	if err != nil {
		MsgResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	DataResponse(w, http.StatusOK, users)
}

func MsgResponse(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	fmt.Fprintf(w, `{"status": %d, "data":"%s"}`, status, message)
}

func InvalidMethod(w http.ResponseWriter) {
	status := http.StatusNotFound
	w.WriteHeader(status)
	fmt.Fprintf(w, `{"status": %d, "message":"method doesn't exist"}`, status)
}

func DataResponse(w http.ResponseWriter, status int, users interface{}) {
	value, err := json.Marshal(users)
	if err != nil {
		MsgResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	w.WriteHeader(status)
	fmt.Fprintf(w, `{"status": %d, "data":%s}`, status, value)
}

func PostUser(ctx context.Context, s Service, w http.ResponseWriter, data interface{}) {
	req := data.(CreateReq) // aca se castea Data a la forma de CreateReq para poder comunicarse con los mismo parametros
	if req.FirstName == "" {
		MsgResponse(w, http.StatusBadRequest, "First name is required")
		return

	}
	if req.LastName == "" {
		MsgResponse(w, http.StatusBadRequest, "Last name is required")
		return

	}
	if req.Email == "" {
		MsgResponse(w, http.StatusBadRequest, "Email is required")
		return

	}

	user, err := s.Create(ctx, req.FirstName, req.LastName, req.Email)
	if err != nil {
		MsgResponse(w, http.StatusInternalServerError, err.Error())
	}

	// maxId++
	// req.ID = maxId
	// users = append(users, req)
	DataResponse(w, http.StatusCreated, user)
}
