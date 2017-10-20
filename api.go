package pork

import (
	"github.com/mspaulding06/nap"
	"github.com/spf13/viper"
)

var api *nap.API

func GitHubAPI() *nap.API {
	if api == nil {
		api = nap.NewAPI("https://api.github.com")
		token := viper.GetString("token")
		api.SetAuth(nap.NewAuthToken(token))
		api.AddResource("fork", GetForkResource())
		api.AddResource("search", GetSearchResource())
		api.AddResource("docs", GetReadmeResource())
	}
	return api
}
