package ymfile

type Options struct {
	Host    string
	Port    int
	RootDir string
}

var GloOptions Options

func init() {
	GloOptions = Options{
		Host:    "localhost",
		Port:    8080,
		RootDir: "./test",
	}
}
