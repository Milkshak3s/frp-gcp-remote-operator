/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// FrpGCPRemoteSpec defines the desired state of FrpGCPRemote
type FrpGCPRemoteSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// GCP Project ID that this will be created in (this will be moved to initial controller config in the future)
	GCPProjectID string `json:"gcp_project_id"`

	// A record name for this service, ex. "my-service" if the full DNS name would be "my-service.my-zone.example.com"
	DNSAName string `json:"dns_a_name"`

	// DNS Zone configured in GCP, must match subdomain A name, ex. "my-zone" if the full subdomain is "my-zone.example.com"
	DNSZone string `json:"dns_zone"`

	// Base domain that the DNS zone sits on, ex. "example.com" if the full subdomain is "my-zone.example.com"
	DNSBaseDomain string `json:"dns_base_domain"`

	// Port that remote FRP instance will listen for frp client connections on
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=65535
	FrpServerPort int `json:"frp_server_port"`

	// Addr of local service that FRP client will proxy connections to
	FrpLocalServiceAddr string `json:"frp_local_service_addr"`

	// Port of local service that FRP client will proxy connections to
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=65535
	FrpLocalServicePort int `json:"frp_local_service_port"`

	// Port exposed by remote FRP instance
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=65535
	FrpRemotePort int `json:"frp_remote_port"`
}

// FrpGCPRemoteStatus defines the observed state of FrpGCPRemote
type FrpGCPRemoteStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Shows values like "active", "provisioning", "error"
	// +optional
	Active string `json:"active"`

	// Shows "healthy" or a related error as the result of health checks (soon tm)
	// +optional
	Health string `json:"healthy"`

	// Shows provisioning step or error, up to "complete"
	// +optional
	ProvisionStatus string `json:"provision_status"`

	// IP Address remote proxy is listening on
	// +optional
	RemoteAddress string `json:"address"`

	// Full DNS name for the remote proxy address
	// +optional
	RemoteDNSName string `json:"dns_name"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// FrpGCPRemote is the Schema for the frpgcpremotes API
type FrpGCPRemote struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FrpGCPRemoteSpec   `json:"spec,omitempty"`
	Status FrpGCPRemoteStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// FrpGCPRemoteList contains a list of FrpGCPRemote
type FrpGCPRemoteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FrpGCPRemote `json:"items"`
}

func init() {
	SchemeBuilder.Register(&FrpGCPRemote{}, &FrpGCPRemoteList{})
}
