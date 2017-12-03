package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type JSONModelFunc func() interface{}
type JSONGetFunc func(key string) (interface{}, error)
type JSONPostFunc func(key string, content interface{}) error
type JSONPutFunc func(key string, content interface{}) error
type JSONDeleteFunc func(key string) error

type JSONObjectHandler struct {
	ModelFunc JSONModelFunc
	Get       JSONGetFunc
	Post      JSONPostFunc
	Put       JSONPutFunc
	Delete    JSONDeleteFunc
}

type JSONResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func SetupJSONObjectRouter(r *mux.Router, objectHandler JSONObjectHandler) {
	r.PathPrefix("/{key}").HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		object, err := objectHandler.Get(mux.Vars(req)["key"])
		response := JSONResponse{}
		response.Data = object

		if err != nil {
			w.WriteHeader(500)
			response.Message = fmt.Sprintf("%s", err)
		} else if object == nil {
			w.WriteHeader(404)
			response.Message = "Not found"
		} else {
			response.Success = true
		}

		json.NewEncoder(w).Encode(response)
	}).Methods("GET")

	r.PathPrefix("").HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		object := objectHandler.ModelFunc()
		response := JSONResponse{}

		err := json.NewDecoder(req.Body).Decode(object)

		if err == nil {
			err := objectHandler.Post(mux.Vars(req)["key"], object)

			if err == nil {
				w.WriteHeader(201)
				response.Success = true
			} else {
				w.WriteHeader(500)
				response.Message = fmt.Sprintf("%s", err)
			}
		} else {
			w.WriteHeader(500)
			response.Message = fmt.Sprintf("Cannot unmarshal JSON body: %s!", err)
		}

		json.NewEncoder(w).Encode(response)
	}).Methods("POST")
}
