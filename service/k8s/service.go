package k8s

type ServiceSlice []Service

type Service struct {
	Kind   string `json:"kind"`
	Detail struct {
		Name            string `json:"name"`
		Namespace       string `json:"namespace"`
		ResourceVersion string `json:"resourceVersion"`
		Uid             string `json:"uid"`
	} `json:"metadata"`
	Spec struct {
		Ports []struct {
			Port     int    `json:"port"`
			Protocol string `json:"protocol"`
		} `json:"ports"`
	} `json:"spec"`
}
