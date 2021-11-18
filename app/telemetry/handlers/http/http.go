package handler

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/flussrd/fluss-back/app/accounts/shared/httputils"
	"github.com/flussrd/fluss-back/app/telemetry/service"
	"github.com/gorilla/mux"
	"golang.org/x/net/context"
)

type HTTPHandler struct {
	service service.Service
	router  *mux.Router
}

func NewHTTPHandler(service service.Service, router *mux.Router) HTTPHandler {
	return HTTPHandler{
		service: service,
		router:  router,
	}
}

func (handler HTTPHandler) Init(ctx context.Context) {
	handler.router.Handle("/messages/{source}", handler.handleReceiveMessage(ctx)).Methods(http.MethodPost)
}

func (handler HTTPHandler) handleReceiveMessage(ctx context.Context) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		source := mux.Vars(r)["source"]
		if source == "" {
			log.Println("missing_source_in_path")
			return
		}

		// TODO: validate the host/headers etc. Where should it be validated?
		requestBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("reading_response_body_failed: " + err.Error())
			httputils.RespondInternalServerError(rw)
			return
		}

		handler.service.HandleHTTPMessage(ctx, source, string(requestBody))

		// TODO: manage the responses according the source. Twillio required the response to be text/plan ( or other but not JSON)
		httputils.RespondText(rw, http.StatusOK, "")
	}
}
