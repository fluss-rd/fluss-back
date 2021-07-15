package handlers

import (
	"net/http"
	"time"

	"github.com/flussrd/fluss-back/app/accounts/shared/httputils"
	"github.com/flussrd/fluss-back/app/reporting/models"
)

var (
	ErrInvalidDate        = httputils.NewBadRequestError("invalid date")
	ErrInvalidCardinality = httputils.NewBadRequestError("invalid cardinality")
)

// TODO: make cleaner
func getSearchOptions(r *http.Request) (models.SearchOptions, error) {
	start := time.Time{}
	end := start

	var err error

	startStr := r.URL.Query().Get("start")
	if startStr != "" {
		start, err = time.Parse(time.RFC3339, startStr)
		if err != nil {
			return models.SearchOptions{}, ErrInvalidDate
		}
	}

	endStr := r.URL.Query().Get("end")
	if endStr != "" {
		end, err = time.Parse(time.RFC3339, endStr)
		if err != nil {
			return models.SearchOptions{}, ErrInvalidDate
		}
	}

	cardinality := r.URL.Query().Get("cardinality")
	if cardinality != "" && !models.IsValidAggregationType(cardinality) {
		return models.SearchOptions{}, ErrInvalidCardinality
	}

	aggregationWindow := r.URL.Query().Get("aggregationWindow")
	if aggregationWindow == "" {
		aggregationWindow = "mean"
	}

	return models.SearchOptions{
		Cardinality:       cardinality,
		AggregationWindow: aggregationWindow,
		Start:             start,
		End:               end,
	}, nil
}
