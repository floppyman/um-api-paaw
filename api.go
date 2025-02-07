package paaw

import (
	"github.com/spf13/viper"

	"github.com/umbrella-sh/um-api-paaw/base"
	"github.com/umbrella-sh/um-api-paaw/endpoints/attendances"
	"github.com/umbrella-sh/um-api-paaw/endpoints/roots"
)

//goland:noinspection GoUnusedExportedFunction
func NewClient(url string, clientId string, clientSecret string) PaawApiClient {
	return newClient(base.PaawOptions{
		ApiUrl:          url,
		ApiClientId:     clientId,
		ApiClientSecret: clientSecret,
	})
}

func NewClientFromOptions(conf *viper.Viper) PaawApiClient {
	return newClient(base.PaawOptions{
		ApiUrl:          conf.GetString("api_url"),
		ApiClientId:     conf.GetString("api_client_id"),
		ApiClientSecret: conf.GetString("api_client_secret"),
	})
}

func newClient(options base.PaawOptions) PaawApiClient {
	base.Init(options)
	return PaawApiClient{
		Attendance: attendances.AttendanceEndPoint{},
		Root:       roots.RootEndPoint{},
	}
}

//goland:noinspection GoNameStartsWithPackageName
type PaawApiClient struct {
	Attendance attendances.AttendanceEndPoint
	Root       roots.RootEndPoint
}
