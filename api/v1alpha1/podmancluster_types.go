/*
Copyright 2022.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clusterv1 "sigs.k8s.io/cluster-api/api/v1beta1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// PodmanClusterSpec defines the desired state of PodmanCluster
type PodmanClusterSpec struct {
	// ControlPlaneEndpoint represents the endpoint used to communicate with the control plane.
	//
	// See https://cluster-api.sigs.k8s.io/developer/architecture/controllers/cluster.html
	// for more details.
	//
	// +optional
	ControlPlaneEndpoint clusterv1.APIEndpoint `json:"controlPlaneEndpoint"`
}

// PodmanClusterStatus defines the observed state of PodmanCluster
type PodmanClusterStatus struct {
	// Ready indicates that the cluster is ready.
	// +optional
	// +kubebuilder:default=false
	Ready bool `json:"ready"`

	// FailureDomains is a list of the failure domains that CAPI should spread the machines across. For
	// the CAPPOD provider this doesn't mean anything.
	FailureDomains clusterv1.FailureDomains `json:"failureDomains,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// PodmanCluster is the Schema for the podmanclusters API
type PodmanCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PodmanClusterSpec   `json:"spec,omitempty"`
	Status PodmanClusterStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// PodmanClusterList contains a list of PodmanCluster
type PodmanClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PodmanCluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&PodmanCluster{}, &PodmanClusterList{})
}
