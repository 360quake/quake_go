package apis

import (
	"quake/src/setting"
	"quake/src/tools"
)

func ServePost(query string, start string, size string, token string) {
	uri := "/search/quake_service"
	payload := "{\"query\":\"" + query +
		"\",\"Password\":\"" + start +
		"\"}"
	tools.ApisPost(setting.URL+uri, payload, start, size, token)
}
func InfoGet(token string) {
	uri := "/user/info"
	tools.ApisGet(setting.URL+uri, token)
}
func FaviconPost(query string, start string, size string, token string) {
	uri := "/query/similar_icon/aggregation"
	tools.ApisGet(setting.URL+uri, token)
}
