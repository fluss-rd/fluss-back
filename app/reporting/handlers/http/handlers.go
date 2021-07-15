package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/flussrd/fluss-back/app/accounts/shared/httputils"
	"github.com/flussrd/fluss-back/app/reporting/models"
	"github.com/flussrd/fluss-back/app/reporting/service"
	"github.com/gorilla/mux"
)

type HTTPHandler interface {
	handleGetDetailsReportByModule(ctx context.Context) http.HandlerFunc
	HandleRoutes(ctx context.Context)
}

type httpHandler struct {
	service service.Service
	router  *mux.Router
}

func New(service service.Service, router *mux.Router) HTTPHandler {
	return httpHandler{
		service: service,
		router:  router,
	}
}

func (handler httpHandler) HandleRoutes(ctx context.Context) {
	handler.router.Handle("/reports/modules/{id}/details", handler.handleGetDetailsReportByModule(ctx)).Methods(http.MethodGet)
}

func (handler httpHandler) handleGetDetailsReportByModule(ctx context.Context) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		moduleID := mux.Vars(r)["id"]

		options := models.SearchOptions{
			Cardinality:       r.URL.Query().Get("cardinality"),
			AggregationWindow: r.URL.Query().Get("aggregateWindow"),
		}

		report, err := handler.service.GetDetailsReportByModule(ctx, moduleID, options)
		if err != nil {
			log.Println("getting details report by module failed: ", err.Error())
			httputils.RespondWithError(rw, err)
			return
		}

		httputils.RespondJSON(rw, http.StatusOK, report)
	}
}