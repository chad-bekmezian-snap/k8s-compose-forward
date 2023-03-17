package app

import "test/fdcas/cmd/k9s-automation/service"

var OvrCExperience = App{
	{
		Service: service.CustomerProcess,
	},
	{
		Service: service.DealerProcess,
	},
	{
		Service: service.LicenseProcess,
	},
	{
		Service: service.Payment,
	},
}
