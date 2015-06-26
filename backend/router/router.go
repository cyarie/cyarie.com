package router

import (
	"github.com/cyarie/cyarie.com/backend/handlers"
	"github.com/cyarie/cyarie.com/backend/settings"
	"github.com/gorilla/mux"
)

func ApiRouter(ac *settings.AppContext) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	IndexFunc := handlers.AppHandler{
		ac,
		false,
		"IndexHandler",
		handlers.IndexHandler,
	}
	router.Methods("GET").Path("/").Name(IndexFunc.HandlerName).Handler(&IndexFunc)

	return router
}
