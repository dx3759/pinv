package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/gookit/goutil/fsutil"
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

func Test_base(t *testing.T) {
	w := httptest.NewRecorder()

	testMap := map[string][]string{
		"根目录下创建文件夹":  []string{"/", "test1", "/test1/"},
		"子目录下创建文件夹":  []string{"/test1", "test2", "/test1/test2/"},
		"目录名称包含.":    []string{"/", "./test3", "/test3"},
		"目录名称包含..":   []string{"/", "../abc", "/abc"},
		"目录名称包含多个..": []string{"/", "../../../ab2", "/ab2"},
	}

	for tn, item := range testMap {
		body := bytes.NewBufferString("current_path=" + item[0] + "&dirname=" + item[1])
		req, _ := http.NewRequest("POST", "/api/v1/createdir", body)
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
		router.ServeHTTP(w, req)

		Convey(tn, t, func() {
			newDir := ymfile.GloOptions.RootDir + item[2]
			So(fsutil.DirExist(newDir), ShouldEqual, true)
			fsutil.DeleteIfExist(newDir)
		})
	}

	Convey("删除文件夹", t, func() {
		body := bytes.NewBufferString("current_path=/&dirname=test1")
		req, _ := http.NewRequest("POST", "/api/v1/delete", body)
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
		router.ServeHTTP(w, req)

		So(fsutil.DeleteIfExist(ymfile.GloOptions.RootDir+"/test1"), ShouldBeNil)
		So(fsutil.DirExist(ymfile.GloOptions.RootDir+"/test1"), ShouldBeFalse)
	})
}
