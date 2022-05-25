package main

import (
	"encoding/json"
	"log"

	"github.com/jasonmccallister/envoy-xds-types/cluster"
)

func main() {
	// create the cluster resources
	resources := []cluster.Resource{
		{
			ResourceType:    "type.googleapis.com/envoy.config.cluster.v3.Cluster",
			Type:            "LOGICAL_DNS",
			Name:            "cluster_name",
			DNSLookupFamily: "V4_ONLY",
			LoadAssignment: cluster.LoadAssignment{
				ClusterName: "cluster_name",
				Endpoints: []cluster.Endpoints{
					{
						LbEndpoints: cluster.LBEndpoints{
							Endpoint: cluster.LBEndpoint{
								HealthCheckConfig: &cluster.HealthCheckConfig{
									Hostname:  "example.com",
									PortValue: 443,
								},
								Address: cluster.Address{
									SocketAddress: cluster.SocketAddress{
										Address:   "example.com",
										PortValue: 443,
									},
								},
							},
						},
					},
				},
			},
			TransportSocket: &cluster.TransportSocket{
				Name: "envoy.transport_sockets.tls",
				TypedConfig: cluster.TransportSocketTypedConfig{
					Type: "type.googleapis.com/envoy.extensions.transport_sockets.tls.v3.UpstreamTlsContext",
					SNI:  "example.com",
				},
			},
		},
	}

	// generate the cluster response with those resources
	response := cluster.NewResponse("V3", cluster.ControlPlane{Identifier: "control-plane-id"}, resources)

	// marshal the response to JSON
	data, err := json.Marshal(response)
	if err != nil {
		log.Fatal("Error marshalling response: ", err)
	}

	log.Println(string(data))
}
