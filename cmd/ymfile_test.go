package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/yzimhao/ymfile"
)

var router *gin.Engine

func init() {
	router = ymfile.SetupRouter()
}

func Test_ping(t *testing.T) {

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	Convey("ping测试", t, func() {
		So(w.Code, ShouldEqual, http.StatusOK)
		So(w.Body.String(), ShouldEqual, "pong")
	})

}
