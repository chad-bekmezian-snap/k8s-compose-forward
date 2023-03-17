package app

import "github.com/chad-bekmezian-snap/k8s-port-forwarding/manual/service"

var LicenseProcess = App{
	{
		Service: service.Account,
	},
	{
		Service: service.Payment,
	},
	{
		Service: service.User,
	},
	{
		Service: service.Consumer,
	},
	{
		Service: service.Application,
	},
	{
		Service: service.License,
	},
}
