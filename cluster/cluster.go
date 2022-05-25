package xds

type Response struct {
	VersionInfo  string       `json:"version_info"`
	ControlPlane ControlPlane `json:"control_plane"`
	Resources    []Resource   `json:"resources"`
}

type ControlPlane struct {
	Identifier string `json:"identifier"`
}

type Resource struct {
	ResourceType    string           `json:"@type"`
	Name            string           `json:"name"`
	DNSLookupFamily string           `json:"dns_lookup_family"`
	Type            string           `json:"type"`
	LoadAssignment  LoadAssignment   `json:"load_assignment"`
	TransportSocket *transportSocket `json:"transport_socket,omitempty"`
}

type transportSocket struct {
	Name        string                     `json:"name"`
	TypedConfig TransportSocketTypedConfig `json:"typed_config"`
}

type TransportSocketTypedConfig struct {
	Type string `json:"@type"`
	SNI  string `json:"sni"`
}

type LoadAssignment struct {
	ClusterName string      `json:"cluster_name"`
	Endpoints   []Endpoints `json:"endpoints"`
}

type Endpoints struct {
	LbEndpoints LBEndpoints `json:"lb_endpoints"`
}

type LBEndpoints struct {
	Endpoint LBEndpoint `json:"endpoint"`
}

type LBEndpoint struct {
	HealthCheckConfig *HealthCheckConfig `json:"health_check_config,omitempty"`
	Address           Address            `json:"address"`
}

type Address struct {
	SocketAddress SocketAddress `json:"socket_address"`
}

type SocketAddress struct {
	Address   string `json:"address"`
	PortValue int    `json:"port_value"`
}

type HealthCheckConfig struct {
	Hostname  string `json:"hostname"`
	PortValue int    `json:"port_value"`
}

// NewResponse creates a new xDS response for a cluster resource (e.g. https://www.envoyproxy.io/docs/envoy/latest/configuration/overview/xds_api#post--v3-discovery-clusters).
func NewResponse(versionInfo string, controlPlane ControlPlane, resources []Resource) Response {
	return Response{
		VersionInfo:  versionInfo,
		ControlPlane: controlPlane,
		Resources:    resources,
	}
}
