package handlers

import (
	"net/http"

	"github.com/dsniels/noobe/components"
)

func HomepageHandler(w http.ResponseWriter, r *http.Request) {
	components.Homepage().Render(r.Context(), w)
}
