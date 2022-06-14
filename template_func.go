package pinv

import "html/template"

func templateFuncMap() template.FuncMap {
	return template.FuncMap{
		"appInfo": appInfo,
	}
}

func appInfo(key string) string {
	v := ""
	switch key {
	case "appname":
		v = Options.AppName()
	case "version":
		v = Options.Version()
	}
	return v
}
