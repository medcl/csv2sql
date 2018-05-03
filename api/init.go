package api

import (
	"github.com/infinitbyte/framework/core/api"
)

// API namespace
type API struct {
	api.Handler
}

// InitAPI init apis
func InitAPI() {

	apis := API{}

	//Index
	api.HandleAPIMethod(api.GET, "/favicon.ico", apis.FaviconAction)

	//Stats APIs
	api.HandleAPIMethod(api.GET, "/_proxy/stats", apis.StatsAction)
	api.HandleAPIMethod(api.POST, "/_proxy/queue/resume", apis.QueueResumeAction)
	api.HandleAPIMethod(api.GET, "/_proxy/queue/stats", apis.QueueStatsAction)

}
