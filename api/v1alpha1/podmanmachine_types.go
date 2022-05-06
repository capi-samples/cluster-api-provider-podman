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

const (
	// MachineFinalizer allows ReconcilePodmanMachine to clean up resources associated with
	// PodmanMachine before removing it from the apiserver.
	MachineFinalizer = "podmanmachine.infrastructure.cluster.x-k8s.io"
)

// PodmanMachineSpec defines the desired state of PodmanMachine
type PodmanMachineSpec struct {
	// ProviderID will be the container name in ProviderID format (podman:////<containername>)
	// +optional
	ProviderID *string `json:"providerID,omitempty"`

	// ExtraMounts describes additional mount points for the node container
	// These may be used to bind a hostPath
	// +optional
	ExtraMounts []Mount `json:"extraMounts,omitempty"`
}

// Mount specifies a host volume to mount into a container.
// This is a simplified version of kind v1alpha4.Mount types.
type Mount struct {
	// Path of the mount within the container.
	ContainerPath string `json:"containerPath,omitempty"`

	// Path of the mount on the host. If the hostPath doesn't exist, then runtimes
	// should report error. If the hostpath is a symbolic link, runtimes should
	// follow the symlink and mount the real destination to container.
	HostPath string `json:"hostPath,omitempty"`

	// If set, the mount is read-only.
	// +optional
	Readonly bool `json:"readOnly,omitempty"`
}

// PodmanMachineStatus defines the observed state of PodmanMachine
type PodmanMachineStatus struct {
	// Ready denotes that the machine (podman container) is ready
	// +optional
	Ready bool `json:"ready"`

	// Addresses contains the associated addresses for the docker machine.
	// +optional
	Addresses []clusterv1.MachineAddress `json:"addresses,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// PodmanMachine is the Schema for the podmanmachines API
type PodmanMachine struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PodmanMachineSpec   `json:"spec,omitempty"`
	Status PodmanMachineStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// PodmanMachineList contains a list of PodmanMachine
type PodmanMachineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PodmanMachine `json:"items"`
}

func init() {
	SchemeBuilder.Register(&PodmanMachine{}, &PodmanMachineList{})
}
