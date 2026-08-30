package handlers

import (
	"net/http"

	"github.com/VOVOplay/creatorcoaster.com/src/views"
)

type PageHandler struct {
}

func NewPageHandler() *PageHandler {
	return &PageHandler{}
}

func (h *PageHandler) HandleHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		component := views.NotFound()
		component.Render(r.Context(), w)
		return
	}

	component := views.Home()
	component.Render(r.Context(), w)
}
