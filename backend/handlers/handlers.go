package handlers

import (
	"net/http"

	"github.com/cyarie/cyarie.com/backend/settings"
	"fmt"
)

type AppHandler struct {
	*settings.AppContext
	AuthRoute			bool
	HandlerName			string
	H					func(*settings.AppContext, http.ResponseWriter, *http.Request) (int, error)
}

func (fn *AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	status, err := fn.H(fn.AppContext, w, r)

	if err != nil {
		fmt.Fprintln(w, err)
	}

	fmt.Println(status)
}

func IndexHandler(a *settings.AppContext, w http.ResponseWriter, r *http.Request) (int, error) {
	var err error

	fmt.Fprint(w, "WELCOME TO CYARIE DOT COM")

	if err != nil {
		return 500, err
	}

	return 200, nil
}
