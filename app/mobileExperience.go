package app

import "test/fdcas/cmd/k9s-automation/service"

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
