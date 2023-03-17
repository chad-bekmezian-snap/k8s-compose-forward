package service

const (
	namespaceExperienceServices = "experience-services"
	namespaceBootServices       = "boot-services"
)

var (
	OvrCExperience = Service{
		Name:        "service/cese-ovrc-experience-k8s",
		Namespace:   namespaceExperienceServices,
		DefaultPort: 5013,
	}

	CustomerProcess = Service{
		Name:        "service/cs-customer-process-boot",
		Namespace:   namespaceBootServices,
		DefaultPort: 8076,
	}

	DealerProcess = Service{
		Name:        "service/cs-dealer-process-boot",
		Namespace:   namespaceBootServices,
		DefaultPort: 8079,
	}

	LicenseProcess = Service{
		Name:        "service/cs-license-process-boot",
		Namespace:   namespaceBootServices,
		DefaultPort: 8061,
	}

	Account = Service{
		Name:        "service/cs-account-boot",
		Namespace:   namespaceBootServices,
		DefaultPort: 8031,
	}

	Consumer = Service{
		Name:        "service/cs-consumer-boot",
		Namespace:   namespaceBootServices,
		DefaultPort: 8041,
	}

	Permission = Service{
		Name:        "service/cs-permission-boot",
		Namespace:   namespaceBootServices,
		DefaultPort: 8056,
	}

	Controller = Service{
		Name:        "service/cs-controller-boot",
		Namespace:   namespaceBootServices,
		DefaultPort: 8047,
	}

	Token = Service{
		Name:        "service/cs-token-boot",
		Namespace:   namespaceBootServices,
		DefaultPort: 8073,
	}

	Event = Service{
		Name:        "service/cs-event-boot",
		Namespace:   namespaceBootServices,
		DefaultPort: 8053,
	}

	Certificate = Service{
		Name:        "service/cs-certificate-boot",
		Namespace:   namespaceBootServices,
		DefaultPort: 7788,
	}

	Dealer = Service{
		Name:        "service/cs-dealer-boot",
		Namespace:   namespaceBootServices,
		DefaultPort: 8028,
	}

	Authentication = Service{
		Name:        "service/cs-authentication-boot",
		Namespace:   namespaceBootServices,
		DefaultPort: 9080,
	}

	User = Service{
		Name:        "service/cs-user-boot",
		Namespace:   namespaceBootServices,
		DefaultPort: 8038,
	}

	Device = Service{
		Name:        "service/cs-device-boot",
		Namespace:   namespaceBootServices,
		DefaultPort: 8555,
	}

	Application = Service{
		Name:        "service/cs-application-boot",
		Namespace:   namespaceBootServices,
		DefaultPort: 8765,
	}

	Payment = Service{
		Name:        "service/cs-payment-service",
		Namespace:   namespaceBootServices,
		DefaultPort: 8064,
	}

	GeolocationProcess = Service{
		Name:        "service/cs-geolocation-process-boot",
		Namespace:   namespaceBootServices,
		DefaultPort: 8069,
	}

	License = Service{
		Name:        "service/cs-license-boot",
		Namespace:   namespaceBootServices,
		DefaultPort: 8050,
	}
)
