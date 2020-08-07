package handler

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type (
	// User handles requests to the user endpoint
	User struct {
		context   context.Context
		responder *HTTPResponder
	}

	// UserResponse is the struct that contains the data needed for the response
	UserResponse struct {
		FirstName string `json:"firstname"`
		LastName  string `json:"lastname"`
		Nickname  string `json:"nickname"`
		Birthday  string `json:"birthday"`
		Email     string `json:"email"`
		Phone     string `json:"phone"`
	}

	// UserTenantsResponse is the struct that contains the data needed for the response
	UserTenantsResponse struct {
		Tenants []int `json:"tenants"`
	}

	// UserInvoicesResponse is the struct that contains the data needed for the response
	UserInvoicesResponse struct {
		Invoices []int `json:"invoices"`
	}
)

// NewUser creates a new NewUserHandler
func NewUser(ctx context.Context, r *HTTPResponder) *User {
	return &User{
		context:   ctx,
		responder: r,
	}
}

// Handle receives a context and handles response for the user endpoint
func (h *User) Handle(w http.ResponseWriter, r *http.Request) {

	// from a route param /user/{userID}
	userIDSlug := chi.URLParam(r, "userID")

	userID, err := strconv.Atoi(userIDSlug)
	if err != nil {
		h.responder.abortBadRequest(w)
		return
	}

	fmt.Println(userID)

	// challenges, err := h.finder.FindUserChallenges(h.context, userID)
	// if err != nil {
	// 	h.responder.Logger.Error(err)
	// 	h.responder.AbortInternalServerError(w)
	// 	return
	// }

	// if len(challenges) == 0 {
	// 	h.responder.UserNotFoundHandler(w, r)
	// 	return
	// }

	h.responder.RespondSuccess(w, &UserResponse{})

	return
}

// HandleTenants receives a context and handles response for the user tenant endpoint
func (h *User) HandleTenants(w http.ResponseWriter, r *http.Request) {

	// from a route param /user/{userID}/tenants
	userIDSlug := chi.URLParam(r, "userID")

	userID, err := strconv.Atoi(userIDSlug)
	if err != nil {
		h.responder.abortBadRequest(w)
		return
	}

	fmt.Println(userID)

	// challenges, err := h.finder.FindUserChallenges(h.context, userID)
	// if err != nil {
	// 	h.responder.Logger.Error(err)
	// 	h.responder.AbortInternalServerError(w)
	// 	return
	// }

	// if len(challenges) == 0 {
	// 	h.responder.UserNotFoundHandler(w, r)
	// 	return
	// }

	h.responder.RespondSuccess(w, &UserTenantsResponse{})

	return
}

// HandleInvoices receives a context and handles response for the user invoices endpoint
func (h *User) HandleInvoices(w http.ResponseWriter, r *http.Request) {

	// from a route param /user/{userID}/invoices
	userIDSlug := chi.URLParam(r, "userID")

	userID, err := strconv.Atoi(userIDSlug)
	if err != nil {
		h.responder.abortBadRequest(w)
		return
	}

	fmt.Println(userID)

	// challenges, err := h.finder.FindUserChallenges(h.context, userID)
	// if err != nil {
	// 	h.responder.Logger.Error(err)
	// 	h.responder.AbortInternalServerError(w)
	// 	return
	// }

	// if len(challenges) == 0 {
	// 	h.responder.UserNotFoundHandler(w, r)
	// 	return
	// }

	h.responder.RespondSuccess(w, &UserInvoicesResponse{})

	return
}
