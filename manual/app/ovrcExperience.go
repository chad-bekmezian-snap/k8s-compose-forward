package app

import "github.com/chad-bekmezian-snap/cs-port-forwarding/service"

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
