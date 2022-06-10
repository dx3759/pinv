package ymfile

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type response struct {
	Ok     bool        `json:"ok"`
	Reason string      `json:"reason"`
	Data   interface{} `json:"data"`
}

type FileInfo struct {
	Name         string `json:"name"`
	IsDir        bool   `json:"is_dir"`
	Size         string `json:"size"`
	LastModified int64  `json:"last_modified"`
	ContentType  string `json:"content_type"`
}

func Run() {
	// http.Handle("/", http.FileServer(http.Dir(opt.Dir)))
	// http.ListenAndServe(":8080", nil)
	startGin()
}

func startGin() {
	r := gin.Default()
	gin.SetMode(gin.DebugMode)

	r.LoadHTMLGlob("./*.html")

	r.GET("/", index)
	r.GET("/filelist", fileList)
	r.POST("/upload", upload)
	r.Run(":8080")
}

func index(c *gin.Context) {
	if gin.Mode() == gin.DebugMode {
		c.HTML(http.StatusOK, "index.html", nil)
	} else {
		c.Data(http.StatusOK, "text/html", []byte(indexHtmlString))
	}
}

func upload(c *gin.Context) {
	file, _ := c.FormFile("file")

	c.SaveUploadedFile(file, "./upload/"+file.Filename)

	c.JSON(http.StatusOK, &response{Ok: true, Reason: "", Data: nil})
}

func fileList(c *gin.Context) {

	root := GloOptions.Dir
	path := c.Query("path")

	if path == "" || path == "/" {
		path = root
	}
	//todo 安全验证

	c.JSON(http.StatusOK, &response{Ok: true, Reason: "", Data: gin.H{"path": path, "files": getFile(path)}})
}

func getFile(pathName string) []FileInfo {
	files := make([]FileInfo, 0)

	rd, err := ioutil.ReadDir(pathName)
	if err != nil {
		return files
	}

	for _, info := range rd {
		files = append(files, FileInfo{
			Name:         info.Name(),
			IsDir:        info.IsDir(),
			Size:         fileSizeHuman(info.Size()),
			LastModified: info.ModTime().Unix(),
			ContentType:  getContentType(pathName + "/" + info.Name()),
		})
	}
	return files
}

func getContentType(fileName string) string {
	f, err := os.Open(fileName)
	if err != nil {
		return "unknown"
	}
	defer f.Close()

	buffer := make([]byte, 512)
	_, err = f.Read(buffer)
	if err != nil {
		return "unknown"
	}
	return http.DetectContentType(buffer)
}

func fileSizeHuman(fbyte int64) string {
	str := ""
	if fbyte < 1048576 {
		str = fmt.Sprintf("%.2fKB", float64(fbyte/1024))
	} else if fbyte == 1048576 {
		str = "1MB"
	} else if fbyte > 1048576 && fbyte < 1073741824 {
		str = fmt.Sprintf("%.2fMB", float64(fbyte/(1024*1024)))
	} else if fbyte > 1048576 && fbyte == 1073741824 {
		str = "1GB"
	} else if fbyte > 1073741824 && fbyte < 1099511627776 {
		str = fmt.Sprintf("%.2fGB", float64(fbyte/(1024*1024*1024)))
	} else {
		str = ">1TB"
	}
	return str
}
