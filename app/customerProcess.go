package app

import "test/fdcas/cmd/k9s-automation/service"

var CustomerProcess = App{
	{
		Service: service.Account,
	},
	{
		Service: service.Consumer,
	},
	{
		Service: service.Permission,
	},
	{
		Service: service.Controller,
	},
	{
		Service: service.Token,
	},
	{
		Service: service.Event,
	},
	{
		Service:     service.DealerProcess,
		DefaultPort: 8059,
	},
	{
		Service: service.Certificate,
	},
	{
		Service: service.Dealer,
	},
	{
		Service: service.Device,
	},
	{
		Service: service.Authentication,
	},
	{
		Service: service.User,
	},
}
