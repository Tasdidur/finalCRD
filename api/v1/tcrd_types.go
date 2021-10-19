/*
Copyright 2021.

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

// TCrdSpec defines the desired state of TCrd
type TCrdSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	//+kubebuilder:validation:MinLength=2
	Name string `json:"name"`

	//+kubebuilder:validation:MinLength=1
	Finder string `json:"finder"`

	//+kubebuilder:validatin:Pattern=[a-z]+[.](com|org|net)
	Domain string `json:"domain"`

	//+kubebuilder:validatin:MinLenght=5
	Image string `json:"image"`

	//+kubebuilder:validation:Minimum=3000
	Port int `json:"port"`

	//+kubebuilder:validation:Minimum=3000
	TargetPort int `json:"target-port"`

	//+kubebuilder:validation:MinItems=1
	Paths []string `json:"paths"`
}

// TCrdStatus defines the observed state of TCrd
type TCrdStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// TCrd is the Schema for the tcrds API
type TCrd struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TCrdSpec   `json:"spec,omitempty"`
	Status TCrdStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// TCrdList contains a list of TCrd
type TCrdList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TCrd `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TCrd{}, &TCrdList{})
}
