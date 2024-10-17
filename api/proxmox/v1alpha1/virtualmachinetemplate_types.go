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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// VirtualMachineTemplateSpec defines the desired state of VirtualMachineTemplate
type VirtualMachineTemplateSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Name is the name of the VM
	Name string `json:"name"`
	// NodeName is the node name
	NodeName string `json:"node"`
	// +kubebuilder:default:=false
	DeletionProtection bool `json:"deletionProtection,omitempty"`

	// VirtualMachineConfig is the configuration of the VM
	VirtualMachineConfig VirtualMachineConfig `json:"virtualMachineConfig,omitempty"`

	// Image config
	ImageConfig StorageDownloadURLSpec `json:"imageConfig"`

	// Cloud Init Config
	CloudInitConfig CloudInitConfig `json:"cloudInitConfig,omitempty"`
}

type VirtualMachineConfig struct {
	// Sockets
	// +kubebuilder:default:=1
	Sockets int `json:"sockets,omitempty"`
	// Cores
	// +kubebuilder:default:=2
	Cores int `json:"cores,omitempty"`
	// Memory as MB
	// +kubebuilder:default:=2048
	Memory  int               `json:"memory,omitempty"`
	Network VMTemplateNetwork `json:"network,omitempty"`
}

type VMTemplateNetwork struct {
	// +kubebuilder:default:="virtio"
	Model string `json:"model,omitempty"`
	// +kubebuilder:default:="vmbr0"
	Bridge string `json:"bridge,omitempty"`
}

type CloudInitConfig struct {

	// User is the user name for the template
	User string `json:"user,omitempty"`
	// Password is the password for the template
	Password string `json:"password,omitempty"`
	// DNS Domain
	DNSDomain string `json:"dnsDomain,omitempty"`
	// DNS Servers
	DNSServers []string `json:"dnsServers,omitempty"`
	// SSH Keys -- suppose to be on openSSH format
	SSHKeys []string `json:"sshKeys,omitempty"`
	// Upgrade Packages
	// +kubebuilder:default:=true
	UpgradePackages bool `json:"upgradePackages,omitempty"`
	// IPConfig is the IP configuration for the VM
	IPConfig IPConfig `json:"ipConfig,omitempty"`
	// TODO: Implement the following
	// cicustom
	// ipconfig[n]
}

type IPConfig struct {
	// Gateway
	Gateway string `json:"gateway,omitempty"`
	// GatewayIPv6
	GatewayIPv6 string `json:"gatewayIPv6,omitempty"`
	// IP Address
	IP string `json:"ip,omitempty"`
	// IPv6 Address
	IPv6 string `json:"ipv6,omitempty"`
	// Subnet Mask
	CIDR string `json:"cidr,omitempty"`
}

// VirtualMachineTemplateStatus defines the observed state of VirtualMachineTemplate
type VirtualMachineTemplateStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Conditions []metav1.Condition `json:"condition,omitempty"`
	Status     string             `json:"status,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:resource:shortName="vmt"
//+kubebuilder:printcolumn:name="Node",type="string",JSONPath=".spec.node",description="The node name"
//+kubebuilder:printcolumn:name="Cores",type="string",JSONPath=".spec.virtualMachineConfig.cores",description="The number of cores"
//+kubebuilder:printcolumn:name="Memory",type="string",JSONPath=".spec.virtualMachineConfig.memory",description="The amount of memory"
//+kubebuilder:printcolumn:name="Image",type="string",JSONPath=".spec.imageConfig.filename",description="The name of the image"
//+kubebuilder:printcolumn:name="Username",type="string",JSONPath=".spec.cloudInitConfig.user",description="The username"
//+kubebuilder:printcolumn:name="Password",type="string",JSONPath=".spec.cloudInitConfig.password",description="The password"
//+kubebuilder:printcolumn:name="Status",type="string",JSONPath=".status.condition[0].type",description="The status of the VM"
//+kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"

// VirtualMachineTemplate is the Schema for the virtualmachinetemplates API
type VirtualMachineTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VirtualMachineTemplateSpec   `json:"spec,omitempty"`
	Status VirtualMachineTemplateStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// VirtualMachineTemplateList contains a list of VirtualMachineTemplate
type VirtualMachineTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualMachineTemplate `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VirtualMachineTemplate{}, &VirtualMachineTemplateList{})
}