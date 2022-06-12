package ymfile

type options struct {
	Host    string
	Port    int
	RootDir string
}

var GloOptions options

func init() {
	GloOptions = options{
		Host:    "0.0.0.0",
		Port:    8080,
		RootDir: "./www/",
	}
}

func (o options) AppName() string {
	return "YMFile"
}

func (o options) Version() string {
	return "v0.0.1"
}
