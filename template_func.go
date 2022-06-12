package ymfile

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
		v = GloOptions.AppName()
	case "version":
		v = GloOptions.Version()
	}
	return v
}
