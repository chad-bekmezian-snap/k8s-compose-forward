package app

import "test/fdcas/cmd/k9s-automation/service"

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
