package pinv

import (
	"fmt"
	"os"
)

var (
	VERSION     = "v0.0.0"
	BUILD       = "2016-11-18 16:40:00"
	COMMIT_SHA1 = ""
	GO_VERSION  = ""
)

func ShowVersion() {
	fmt.Printf("version: %s\nbuild: %s\ncommit: %s\n%s\n", VERSION, BUILD, COMMIT_SHA1, GO_VERSION)
	os.Exit(0)
}
