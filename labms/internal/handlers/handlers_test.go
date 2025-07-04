package handlers

import (
	"net/http/httptest"
	"strings"

	"github.com/gin-gonic/gin"
)

type TestHandler struct {
	Router *gin.Engine
}

func GetTestHandlerRouter() *TestHandler {
	router := gin.Default()
	return &TestHandler{Router: router}

}
func (th *TestHandler) PerformRouteTest(method, API, body string) *httptest.ResponseRecorder {
	rq := httptest.NewRequest(method, API, strings.NewReader(body))
	rq.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	th.Router.ServeHTTP(w, rq)
	return w
}
