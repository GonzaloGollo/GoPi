package user

import (
	"context"
	"encoding/json" //transformamos a json
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
			GetAllUser(w)
			//status = 200          // indicamos el valor que queremos que reproduzca por status
			//w.WriteHeader(status) // aqui escribimos en el header el valor del status
			// fmt.Fprintf(w, `{"status": %d, "messege": "%s"}`, status, "success in get")
		case http.MethodPost:
			decode := json.NewDecoder(r.Body)
			var u User
			if err := decode.Decode(&u); err != nil {
				MsgResponse(w, http.StatusBadRequest, err.Error())
				return
			}
			PostUser(w, u)
			// status = 200
			// w.WriteHeader(status)
			// fmt.Fprintf(w, `{"status": %d, "messege": "%s"}`, status, "success in post")
		default:
			InvalidMethod(w)
			// status = 404
			// w.WriteHeader(status)
			// fmt.Fprintf(w, `{"status": %d, "messege": "%s"}`, status, "not found")
		}
	}
}

func UserServer(w http.ResponseWriter, r *http.Request) {
	// var status int
	switch r.Method {
	case http.MethodGet:
		GetAllUser(w)
		//status = 200          // indicamos el valor que queremos que reproduzca por status
		//w.WriteHeader(status) // aqui escribimos en el header el valor del status
		// fmt.Fprintf(w, `{"status": %d, "messege": "%s"}`, status, "success in get")
	case http.MethodPost:
		decode := json.NewDecoder(r.Body)
		var u User
		if err := decode.Decode(&u); err != nil {
			MsgResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		PostUser(w, u)
		// status = 200
		// w.WriteHeader(status)
		// fmt.Fprintf(w, `{"status": %d, "messege": "%s"}`, status, "success in post")
	default:
		InvalidMethod(w)
		// status = 404
		// w.WriteHeader(status)
		// fmt.Fprintf(w, `{"status": %d, "messege": "%s"}`, status, "not found")
	}
}

func GetAllUser(w http.ResponseWriter) {
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
	fmt.Fprintf(w, `{"status": %d, "data":%s}`, status, value) //Fprintf para sobreescribir sobre el response
}

func PostUser(w http.ResponseWriter, data interface{}) {
	user := data.(User) //  .() es para castear data a user

	if user.FirstName == "" {
		MsgResponse(w, http.StatusBadRequest, "First name is required")
		return

	}
	if user.LastName == "" {
		MsgResponse(w, http.StatusBadRequest, "Last name is required")
		return

	}
	if user.Email == "" {
		MsgResponse(w, http.StatusBadRequest, "Email is required")
		return

	}

	maxId++
	user.ID = maxId
	users = append(users, user)
	DataResponse(w, http.StatusCreated, user)
}
