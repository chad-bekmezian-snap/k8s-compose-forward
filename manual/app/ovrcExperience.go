package app

import "github.com/chad-bekmezian-snap/k8s-port-forwarding/manual/service"

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
