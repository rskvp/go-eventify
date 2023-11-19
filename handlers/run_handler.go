package handlers

import (
	"net/http"
	"strings"

	"assalielmehdi/eventify/runners"
)

type RunHandler struct {
	Runner *runners.HttpRunner
}

func NewRunHandler(r *runners.HttpRunner) *RunHandler {
	return &RunHandler{r}
}

func (h *RunHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost && strings.HasPrefix(r.URL.Path, "/api/flow/run/") {
		id, _ := strings.CutPrefix(r.URL.Path, "/api/flow/run/")

		h.Runner.Run(id)
	}
}
