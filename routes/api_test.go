package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPingRoute(t *testing.T) {
	gin.SetMode(gin.TestMode)

	Convey("Given a HTTP request for /ping", t, func() {
		req, _ := http.NewRequest("GET", "/ping", nil)
		w := httptest.NewRecorder()

		Convey("When the request is handled by the router", func() {
			r := gin.Default()
			API(r)
			r.ServeHTTP(w, req)

			Convey("Then response should be ok and include 'pong'", func() {
				So(w.Code, ShouldEqual, 200)
				So(w.Body.String(), ShouldContainSubstring, "pong")
			})
		})
	})
}
