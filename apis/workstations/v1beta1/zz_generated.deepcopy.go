//go:build !ignore_autogenerated

// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkstationCluster) DeepCopyInto(out *WorkstationCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkstationCluster.
func (in *WorkstationCluster) DeepCopy() *WorkstationCluster {
	if in == nil {
		return nil
	}
	out := new(WorkstationCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WorkstationCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkstationClusterAnnotation) DeepCopyInto(out *WorkstationClusterAnnotation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkstationClusterAnnotation.
func (in *WorkstationClusterAnnotation) DeepCopy() *WorkstationClusterAnnotation {
	if in == nil {
		return nil
	}
	out := new(WorkstationClusterAnnotation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkstationClusterGCPCondition) DeepCopyInto(out *WorkstationClusterGCPCondition) {
	*out = *in
	if in.Code != nil {
		in, out := &in.Code, &out.Code
		*out = new(int32)
		**out = **in
	}
	if in.Message != nil {
		in, out := &in.Message, &out.Message
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkstationClusterGCPCondition.
func (in *WorkstationClusterGCPCondition) DeepCopy() *WorkstationClusterGCPCondition {
	if in == nil {
		return nil
	}
	out := new(WorkstationClusterGCPCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkstationClusterLabel) DeepCopyInto(out *WorkstationClusterLabel) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkstationClusterLabel.
func (in *WorkstationClusterLabel) DeepCopy() *WorkstationClusterLabel {
	if in == nil {
		return nil
	}
	out := new(WorkstationClusterLabel)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkstationClusterList) DeepCopyInto(out *WorkstationClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]WorkstationCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkstationClusterList.
func (in *WorkstationClusterList) DeepCopy() *WorkstationClusterList {
	if in == nil {
		return nil
	}
	out := new(WorkstationClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WorkstationClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkstationClusterObservedState) DeepCopyInto(out *WorkstationClusterObservedState) {
	*out = *in
	if in.Uid != nil {
		in, out := &in.Uid, &out.Uid
		*out = new(string)
		**out = **in
	}
	if in.Reconciling != nil {
		in, out := &in.Reconciling, &out.Reconciling
		*out = new(bool)
		**out = **in
	}
	if in.CreateTime != nil {
		in, out := &in.CreateTime, &out.CreateTime
		*out = new(string)
		**out = **in
	}
	if in.UpdateTime != nil {
		in, out := &in.UpdateTime, &out.UpdateTime
		*out = new(string)
		**out = **in
	}
	if in.DeleteTime != nil {
		in, out := &in.DeleteTime, &out.DeleteTime
		*out = new(string)
		**out = **in
	}
	if in.Etag != nil {
		in, out := &in.Etag, &out.Etag
		*out = new(string)
		**out = **in
	}
	if in.ControlPlaneIP != nil {
		in, out := &in.ControlPlaneIP, &out.ControlPlaneIP
		*out = new(string)
		**out = **in
	}
	if in.ClusterHostname != nil {
		in, out := &in.ClusterHostname, &out.ClusterHostname
		*out = new(string)
		**out = **in
	}
	if in.ServiceAttachmentURI != nil {
		in, out := &in.ServiceAttachmentURI, &out.ServiceAttachmentURI
		*out = new(string)
		**out = **in
	}
	if in.Degraded != nil {
		in, out := &in.Degraded, &out.Degraded
		*out = new(bool)
		**out = **in
	}
	if in.GCPConditions != nil {
		in, out := &in.GCPConditions, &out.GCPConditions
		*out = make([]WorkstationClusterGCPCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkstationClusterObservedState.
func (in *WorkstationClusterObservedState) DeepCopy() *WorkstationClusterObservedState {
	if in == nil {
		return nil
	}
	out := new(WorkstationClusterObservedState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkstationClusterSpec) DeepCopyInto(out *WorkstationClusterSpec) {
	*out = *in
	out.ProjectRef = in.ProjectRef
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make([]WorkstationClusterAnnotation, len(*in))
		copy(*out, *in)
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make([]WorkstationClusterLabel, len(*in))
		copy(*out, *in)
	}
	out.NetworkRef = in.NetworkRef
	out.SubnetworkRef = in.SubnetworkRef
	if in.PrivateClusterConfig != nil {
		in, out := &in.PrivateClusterConfig, &out.PrivateClusterConfig
		*out = new(WorkstationCluster_PrivateClusterConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkstationClusterSpec.
func (in *WorkstationClusterSpec) DeepCopy() *WorkstationClusterSpec {
	if in == nil {
		return nil
	}
	out := new(WorkstationClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkstationClusterStatus) DeepCopyInto(out *WorkstationClusterStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	if in.ObservedGeneration != nil {
		in, out := &in.ObservedGeneration, &out.ObservedGeneration
		*out = new(int64)
		**out = **in
	}
	if in.ExternalRef != nil {
		in, out := &in.ExternalRef, &out.ExternalRef
		*out = new(string)
		**out = **in
	}
	if in.ObservedState != nil {
		in, out := &in.ObservedState, &out.ObservedState
		*out = new(WorkstationClusterObservedState)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkstationClusterStatus.
func (in *WorkstationClusterStatus) DeepCopy() *WorkstationClusterStatus {
	if in == nil {
		return nil
	}
	out := new(WorkstationClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkstationCluster_PrivateClusterConfig) DeepCopyInto(out *WorkstationCluster_PrivateClusterConfig) {
	*out = *in
	if in.EnablePrivateEndpoint != nil {
		in, out := &in.EnablePrivateEndpoint, &out.EnablePrivateEndpoint
		*out = new(bool)
		**out = **in
	}
	if in.AllowedProjects != nil {
		in, out := &in.AllowedProjects, &out.AllowedProjects
		*out = make([]refsv1beta1.ProjectRef, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkstationCluster_PrivateClusterConfig.
func (in *WorkstationCluster_PrivateClusterConfig) DeepCopy() *WorkstationCluster_PrivateClusterConfig {
	if in == nil {
		return nil
	}
	out := new(WorkstationCluster_PrivateClusterConfig)
	in.DeepCopyInto(out)
	return out
}