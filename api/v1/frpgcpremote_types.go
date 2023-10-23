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

	// Foo is an example field of FrpGCPRemote. Edit frpgcpremote_types.go to remove/update
	Foo string `json:"foo,omitempty"`
}

// FrpGCPRemoteStatus defines the observed state of FrpGCPRemote
type FrpGCPRemoteStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
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

// FrpGCPRemoteList contains a list of FrpGCPRemote
type FrpGCPRemoteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FrpGCPRemote `json:"items"`
}

func init() {
	SchemeBuilder.Register(&FrpGCPRemote{}, &FrpGCPRemoteList{})
}
