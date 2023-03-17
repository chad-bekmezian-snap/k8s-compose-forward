package service

const (
	namespaceExperienceServices = "experience-services"
	namespaceBootServices       = "boot-services"
)

var (
	OvrCExperience = Service{
		K8sName:      "service/cese-ovrc-experience-k8s",
		K8sNamespace: namespaceExperienceServices,
		DefaultPort:  5013,
	}

	CustomerProcess = Service{
		K8sName:      "service/cs-customer-process-boot",
		K8sNamespace: namespaceBootServices,
		DefaultPort:  8076,
	}

	DealerProcess = Service{
		K8sName:      "service/cs-dealer-process-boot",
		K8sNamespace: namespaceBootServices,
		DefaultPort:  8079,
	}

	LicenseProcess = Service{
		K8sName:      "service/cs-license-process-boot",
		K8sNamespace: namespaceBootServices,
		DefaultPort:  8061,
	}

	Account = Service{
		K8sName:      "service/cs-account-boot",
		K8sNamespace: namespaceBootServices,
		DefaultPort:  8031,
	}

	Consumer = Service{
		K8sName:      "service/cs-consumer-boot",
		K8sNamespace: namespaceBootServices,
		DefaultPort:  8041,
	}

	Permission = Service{
		K8sName:      "service/cs-permission-boot",
		K8sNamespace: namespaceBootServices,
		DefaultPort:  8056,
	}

	Controller = Service{
		K8sName:      "service/cs-controller-boot",
		K8sNamespace: namespaceBootServices,
		DefaultPort:  8047,
	}

	Token = Service{
		K8sName:      "service/cs-token-boot",
		K8sNamespace: namespaceBootServices,
		DefaultPort:  8073,
	}

	Event = Service{
		K8sName:      "service/cs-event-boot",
		K8sNamespace: namespaceBootServices,
		DefaultPort:  8053,
	}

	Certificate = Service{
		K8sName:      "service/cs-certificate-boot",
		K8sNamespace: namespaceBootServices,
		DefaultPort:  7788,
	}

	Dealer = Service{
		K8sName:      "service/cs-dealer-boot",
		K8sNamespace: namespaceBootServices,
		DefaultPort:  8028,
	}

	Authentication = Service{
		K8sName:      "service/cs-authentication-boot",
		K8sNamespace: namespaceBootServices,
		DefaultPort:  9080,
	}

	User = Service{
		K8sName:      "service/cs-user-boot",
		K8sNamespace: namespaceBootServices,
		DefaultPort:  8038,
	}

	Device = Service{
		K8sName:      "service/cs-device-boot",
		K8sNamespace: namespaceBootServices,
		DefaultPort:  8555,
	}

	Application = Service{
		K8sName:      "service/cs-application-boot",
		K8sNamespace: namespaceBootServices,
		DefaultPort:  8765,
	}

	Payment = Service{
		K8sName:      "service/cs-payment-service",
		K8sNamespace: namespaceBootServices,
		DefaultPort:  8064,
	}

	GeolocationProcess = Service{
		K8sName:      "service/cs-geolocation-process-boot",
		K8sNamespace: namespaceBootServices,
		DefaultPort:  8069,
	}

	License = Service{
		K8sName:      "service/cs-license-boot",
		K8sNamespace: namespaceBootServices,
		DefaultPort:  8050,
	}
)
