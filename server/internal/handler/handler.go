package handler

import (
	"encoding/json"
	"net/http"

	catalog "landlord-app/server/internal/catalog"

	"github.com/sirupsen/logrus"
)

type (
	// HTTPResponder is a utility for http responses
	HTTPResponder struct {
		Logger *logrus.Logger
	}

	// Responder is the interface to send a response to the user.
	Responder interface {
		// JSON(w http.ResponseWriter, code int, body Response)
		NotFoundHandler(w http.ResponseWriter, resp *http.Request)
		// UserNotFoundHandler(w http.ResponseWriter, resp *http.Request)
		// ChallengeNotFoundHandler(w http.ResponseWriter, resp *http.Request)
		RespondSuccess(w http.ResponseWriter, response Response)
		// AbortInternalServerError(w http.ResponseWriter)
		// RespondNoContent(w http.ResponseWriter)
		// AbortBadRequestWithError(w http.ResponseWriter, err error)
		// AbortRequestEntityTooLarge(w http.ResponseWriter)
		// RespondConflict(w http.ResponseWriter)
	}

	// ErrorResponse is an interface for the body of a json error response
	ErrorResponse struct {
		StatusCode int    `json:"statusCode"`
		ErrorCode  string `json:"errorCode"`
		Message    string `json:"message"`
	}

	// Response is an interface for the body of a json response
	Response interface{}
)

// NewResponder creates a new responder
func NewResponder(l *logrus.Logger) *HTTPResponder {
	return &HTTPResponder{
		Logger: l,
	}
}

// JSON is responsible for returning the context of the http request in a JSON format
func (r *HTTPResponder) JSON(w http.ResponseWriter, code int, body Response) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		r.Logger.WithError(err).Warn(catalog.ErrorMarshalJSONResponse)

		if code >= 200 && code < 300 {
			w.WriteHeader(http.StatusNoContent)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(jsonBody)
}

// NotFoundHandler returns a 404 and a route not found response
func (r *HTTPResponder) NotFoundHandler(w http.ResponseWriter, resp *http.Request) {
	r.JSON(w, http.StatusNotFound,
		&ErrorResponse{
			StatusCode: http.StatusNotFound,
			ErrorCode:  catalog.CodeRouteNotFound,
			Message:    catalog.ErrorRouteNotFound,
		})
}

// RespondSuccess returns a 200 and the response
func (r *HTTPResponder) RespondSuccess(w http.ResponseWriter, response Response) {
	r.JSON(w, http.StatusOK, response)
}

// abortBadRequest returns a 400 and the response
func (r *HTTPResponder) abortBadRequest(w http.ResponseWriter) {
	resp := ErrorResponse{
		StatusCode: http.StatusBadRequest,
		ErrorCode:  catalog.CodeBadRequest,
		Message:    catalog.ErrorBadRequest,
	}

	r.JSON(w, http.StatusBadRequest, &resp)
}
