/*
Copyright 2021 The Kubernetes Authors.

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

// Features is the collection of all discovered features.
type Features struct {
	// Flags contains all the flag-type features of the node.
	// +optional
	Flags map[string]FlagFeatureSet `json:"flags" protobuf:"bytes,1,rep,name=flags"`
	// Attributes contains all the attribute-type features of the node.
	// +optional
	Attributes map[string]AttributeFeatureSet `json:"attributes" protobuf:"bytes,2,rep,name=vattributes"`
	// Instances contains all the instance-type features of the node.
	// +optional
	Instances map[string]InstanceFeatureSet `json:"instances" protobuf:"bytes,3,rep,name=instances"`
}

// FlagFeatureSet is a set of simple features only containing names without values.
type FlagFeatureSet struct {
	Elements map[string]Nil `json:"elements" protobuf:"bytes,1,rep,name=elements"`
}

// AttributeFeatureSet is a set of features having string value.
type AttributeFeatureSet struct {
	Elements map[string]string `json:"elements" protobuf:"bytes,1,rep,name=elements"`
}

// InstanceFeatureSet is a set of features each of which is an instance having multiple attributes.
type InstanceFeatureSet struct {
	Elements []InstanceFeature `json:"elements" protobuf:"bytes,1,rep,name=elements"`
}

// InstanceFeature represents one instance of a complex features, e.g. a device.
type InstanceFeature struct {
	Attributes map[string]string `json:"attributes" protobuf:"bytes,1,rep,name=attributes"`
}

// Nil is a dummy empty struct for protobuf compatibility
type Nil struct{}
