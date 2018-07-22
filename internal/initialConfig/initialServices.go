package initialConfig

import "strings"

func GetApplicationSite() string {
	var appSite string
	appSite = strings.Join(configFromJsonFile.Site,"")
	return appSite
}
