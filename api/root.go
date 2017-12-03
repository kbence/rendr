package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kbence/rendr/util"
)

var ApiV1ObjectHandlers map[string]JSONObjectHandler = map[string]JSONObjectHandler{
	"/job": JSONObjectHandler{
		ModelFunc: NewUntypedJobModel,
		Get:       GetJob,
		Post:      PostJob,
		Put:       PutJob,
		Delete:    DeleteJob},
}

type Test struct {
}

func (t *Test) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte(fmt.Sprintf("Path = '%s'", req.URL.Path)))
}

func getListenAddress() string {
	return fmt.Sprintf("%s:%s", util.GetEnvWithDefault("HOST", ""), util.GetEnvWithDefault("PORT", "5678"))
}

func NewApiRouter(apiPrefix string, handlers map[string]JSONObjectHandler) *mux.Router {
	router := mux.NewRouter().StrictSlash(false)

	for path, handler := range handlers {
		prefix := fmt.Sprintf("%s%s", apiPrefix, path)
		SetupJSONObjectRouter(router.PathPrefix(prefix).Subrouter(), handler)
	}

	return router
}

func Serve() {
	http.ListenAndServe(getListenAddress(), NewApiRouter("/api/v1", ApiV1ObjectHandlers))
}
