//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1beta1

import (
	v1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataCatalogPolicyTag) DeepCopyInto(out *DataCatalogPolicyTag) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataCatalogPolicyTag.
func (in *DataCatalogPolicyTag) DeepCopy() *DataCatalogPolicyTag {
	if in == nil {
		return nil
	}
	out := new(DataCatalogPolicyTag)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DataCatalogPolicyTag) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataCatalogPolicyTagList) DeepCopyInto(out *DataCatalogPolicyTagList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DataCatalogPolicyTag, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataCatalogPolicyTagList.
func (in *DataCatalogPolicyTagList) DeepCopy() *DataCatalogPolicyTagList {
	if in == nil {
		return nil
	}
	out := new(DataCatalogPolicyTagList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DataCatalogPolicyTagList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataCatalogPolicyTagSpec) DeepCopyInto(out *DataCatalogPolicyTagSpec) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.ParentPolicyTagRef != nil {
		in, out := &in.ParentPolicyTagRef, &out.ParentPolicyTagRef
		*out = new(v1alpha1.ResourceRef)
		**out = **in
	}
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	out.TaxonomyRef = in.TaxonomyRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataCatalogPolicyTagSpec.
func (in *DataCatalogPolicyTagSpec) DeepCopy() *DataCatalogPolicyTagSpec {
	if in == nil {
		return nil
	}
	out := new(DataCatalogPolicyTagSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataCatalogPolicyTagStatus) DeepCopyInto(out *DataCatalogPolicyTagStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	if in.ChildPolicyTags != nil {
		in, out := &in.ChildPolicyTags, &out.ChildPolicyTags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.ObservedGeneration != nil {
		in, out := &in.ObservedGeneration, &out.ObservedGeneration
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataCatalogPolicyTagStatus.
func (in *DataCatalogPolicyTagStatus) DeepCopy() *DataCatalogPolicyTagStatus {
	if in == nil {
		return nil
	}
	out := new(DataCatalogPolicyTagStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataCatalogTaxonomy) DeepCopyInto(out *DataCatalogTaxonomy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataCatalogTaxonomy.
func (in *DataCatalogTaxonomy) DeepCopy() *DataCatalogTaxonomy {
	if in == nil {
		return nil
	}
	out := new(DataCatalogTaxonomy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DataCatalogTaxonomy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataCatalogTaxonomyList) DeepCopyInto(out *DataCatalogTaxonomyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DataCatalogTaxonomy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataCatalogTaxonomyList.
func (in *DataCatalogTaxonomyList) DeepCopy() *DataCatalogTaxonomyList {
	if in == nil {
		return nil
	}
	out := new(DataCatalogTaxonomyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DataCatalogTaxonomyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataCatalogTaxonomySpec) DeepCopyInto(out *DataCatalogTaxonomySpec) {
	*out = *in
	if in.ActivatedPolicyTypes != nil {
		in, out := &in.ActivatedPolicyTypes, &out.ActivatedPolicyTypes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	out.ProjectRef = in.ProjectRef
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataCatalogTaxonomySpec.
func (in *DataCatalogTaxonomySpec) DeepCopy() *DataCatalogTaxonomySpec {
	if in == nil {
		return nil
	}
	out := new(DataCatalogTaxonomySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataCatalogTaxonomyStatus) DeepCopyInto(out *DataCatalogTaxonomyStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.ObservedGeneration != nil {
		in, out := &in.ObservedGeneration, &out.ObservedGeneration
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataCatalogTaxonomyStatus.
func (in *DataCatalogTaxonomyStatus) DeepCopy() *DataCatalogTaxonomyStatus {
	if in == nil {
		return nil
	}
	out := new(DataCatalogTaxonomyStatus)
	in.DeepCopyInto(out)
	return out
}
