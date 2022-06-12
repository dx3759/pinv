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
		RootDir: "./www/",
	}
}

func (o Options) AppName() string {
	return "YMFile"
}

func (o Options) Version() string {
	return "v0.0.1"
}
