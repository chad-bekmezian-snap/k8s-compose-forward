package app

import "github.com/chad-bekmezian-snap/cs-port-forwarding/service"

var MobileExperience = App{
	{
		Service: service.LicenseProcess,
	},
	{
		Service: service.CustomerProcess,
	},
	{
		Service: service.Authentication,
	},
	{
		Service: service.Payment,
	},
	{
		Service: service.GeolocationProcess,
	},
	{
		Service: service.DealerProcess,
	},
	{
		Service: service.Application,
	},
}
