package manifests

type bootkubeTemplateData struct {
	AggregatorCaCert                string
	AggregatorCaKey                 string
	ApiserverCert                   string
	ApiserverKey                    string
	ApiserverProxyCert              string
	ApiserverProxyKey               string
	Base64encodeCloudProviderConfig string
	ClusterapiCaCert                string
	ClusterapiCaKey                 string
	EtcdCaCert                      string
	EtcdClientCert                  string
	EtcdClientKey                   string
	KubeCaCert                      string
	KubeCaKey                       string
	McsTLSCert                      string
	McsTLSKey                       string
	OidcCaCert                      string
	OpenshiftApiserverCert          string
	OpenshiftApiserverKey           string
	OpenshiftLoopbackKubeconfig     string
	PullSecret                      string
	RootCaCert                      string
	ServiceaccountKey               string
	ServiceaccountPub               string
	ServiceServingCaCert            string
	ServiceServingCaKey             string
	TectonicNetworkOperatorImage    string
	WorkerIgnConfig                 string
	CVOClusterID                    string
	EtcdEndpointHostnames           []string
	EtcdEndpointDNSSuffix           string
}

type tectonicTemplateData struct {
	KubeAddonOperatorImage string
	PullSecret             string
}
