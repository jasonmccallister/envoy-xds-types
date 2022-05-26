package route

type Response struct {
	VersionInfo  string       `json:"version_info"`
	ControlPlane ControlPlane `json:"control_plane"`
	Resources    []Resource   `json:"resources"`
}

type ControlPlane struct {
	Identifier string `json:"identifier"`
}

type Routes struct {
	Name     string    `json:"name"`
	Match    Match     `json:"match"`
	Route    *Route    `json:"route,omitempty"`
	Redirect *Redirect `json:"redirect,omitempty"`
}

type Match struct {
	Prefix string `json:"prefix"`
}

type Route struct {
	HostRewriteLiteral string `json:"host_rewrite_literal,omitempty"`
	Cluster            string `json:"cluster,omitempty"`
}

type Redirect struct {
	HostRedirect  string `json:"host_redirect,omitempty"`
	PrefixRewrite string `json:"prefix_rewrite,omitempty"`
}

type VirtualHost struct {
	Name    string   `json:"name"`
	Domains []string `json:"domains"`
	Routes  []Routes `json:"routes"`
}

type Resource struct {
	Type         string        `json:"@type"`
	Name         string        `json:"name"`
	VirtualHosts []VirtualHost `json:"virtual_hosts"`
}

// NewResponse creates a new route response with the given version info and control plane identifier.
func NewResponse(versionInfo string, controlPlane ControlPlane, resources []Resource) Response {
	return Response{
		VersionInfo:  versionInfo,
		ControlPlane: controlPlane,
		Resources:    resources,
	}
}
