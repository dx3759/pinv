package pinv

type options struct {
	Host    string
	Port    int
	RootDir string

	AllowDelete bool
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
	return "Pinv"
}

func (o options) Version() string {
	return VERSION
}
