package webresources

import (
	"net/http"
)

// HomeRoutes to home enpoints
var HomeRoutes = Routes{
	Route{
		Name:        "Home",
		Method:      "GET",
		Path:        "/",
		HandlerFunc: index,
	},
}

func index(w http.ResponseWriter, r *http.Request) {
	header := w.Header()

	header.Add("Content-type", "text/plain")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Welcome People!"))
}
