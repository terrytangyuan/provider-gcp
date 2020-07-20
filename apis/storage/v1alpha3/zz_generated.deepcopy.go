// +build !ignore_autogenerated

/*
Copyright 2019 The Crossplane Authors.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha3

import (
	"github.com/crossplane/crossplane-runtime/apis/core/v1alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ACLRule) DeepCopyInto(out *ACLRule) {
	*out = *in
	if in.ProjectTeam != nil {
		in, out := &in.ProjectTeam, &out.ProjectTeam
		*out = new(ProjectTeam)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ACLRule.
func (in *ACLRule) DeepCopy() *ACLRule {
	if in == nil {
		return nil
	}
	out := new(ACLRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Bucket) DeepCopyInto(out *Bucket) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Bucket.
func (in *Bucket) DeepCopy() *Bucket {
	if in == nil {
		return nil
	}
	out := new(Bucket)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Bucket) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketClass) DeepCopyInto(out *BucketClass) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.SpecTemplate.DeepCopyInto(&out.SpecTemplate)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketClass.
func (in *BucketClass) DeepCopy() *BucketClass {
	if in == nil {
		return nil
	}
	out := new(BucketClass)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BucketClass) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketClassList) DeepCopyInto(out *BucketClassList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BucketClass, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketClassList.
func (in *BucketClassList) DeepCopy() *BucketClassList {
	if in == nil {
		return nil
	}
	out := new(BucketClassList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BucketClassList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketClassSpecTemplate) DeepCopyInto(out *BucketClassSpecTemplate) {
	*out = *in
	out.ClassSpecTemplate = in.ClassSpecTemplate
	in.BucketParameters.DeepCopyInto(&out.BucketParameters)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketClassSpecTemplate.
func (in *BucketClassSpecTemplate) DeepCopy() *BucketClassSpecTemplate {
	if in == nil {
		return nil
	}
	out := new(BucketClassSpecTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketEncryption) DeepCopyInto(out *BucketEncryption) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketEncryption.
func (in *BucketEncryption) DeepCopy() *BucketEncryption {
	if in == nil {
		return nil
	}
	out := new(BucketEncryption)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketList) DeepCopyInto(out *BucketList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Bucket, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketList.
func (in *BucketList) DeepCopy() *BucketList {
	if in == nil {
		return nil
	}
	out := new(BucketList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BucketList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketLogging) DeepCopyInto(out *BucketLogging) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketLogging.
func (in *BucketLogging) DeepCopy() *BucketLogging {
	if in == nil {
		return nil
	}
	out := new(BucketLogging)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketOutputAttrs) DeepCopyInto(out *BucketOutputAttrs) {
	*out = *in
	if in.BucketPolicyOnly != nil {
		in, out := &in.BucketPolicyOnly, &out.BucketPolicyOnly
		*out = new(BucketPolicyOnly)
		(*in).DeepCopyInto(*out)
	}
	in.Created.DeepCopyInto(&out.Created)
	if in.RetentionPolicy != nil {
		in, out := &in.RetentionPolicy, &out.RetentionPolicy
		*out = new(RetentionPolicyStatus)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketOutputAttrs.
func (in *BucketOutputAttrs) DeepCopy() *BucketOutputAttrs {
	if in == nil {
		return nil
	}
	out := new(BucketOutputAttrs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketParameters) DeepCopyInto(out *BucketParameters) {
	*out = *in
	in.BucketSpecAttrs.DeepCopyInto(&out.BucketSpecAttrs)
	if in.ServiceAccountSecretRef != nil {
		in, out := &in.ServiceAccountSecretRef, &out.ServiceAccountSecretRef
		*out = new(v1alpha1.SecretReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketParameters.
func (in *BucketParameters) DeepCopy() *BucketParameters {
	if in == nil {
		return nil
	}
	out := new(BucketParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketPolicyOnly) DeepCopyInto(out *BucketPolicyOnly) {
	*out = *in
	in.LockedTime.DeepCopyInto(&out.LockedTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketPolicyOnly.
func (in *BucketPolicyOnly) DeepCopy() *BucketPolicyOnly {
	if in == nil {
		return nil
	}
	out := new(BucketPolicyOnly)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketSpec) DeepCopyInto(out *BucketSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.BucketParameters.DeepCopyInto(&out.BucketParameters)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketSpec.
func (in *BucketSpec) DeepCopy() *BucketSpec {
	if in == nil {
		return nil
	}
	out := new(BucketSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketSpecAttrs) DeepCopyInto(out *BucketSpecAttrs) {
	*out = *in
	in.BucketUpdatableAttrs.DeepCopyInto(&out.BucketUpdatableAttrs)
	if in.ACL != nil {
		in, out := &in.ACL, &out.ACL
		*out = make([]ACLRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DefaultObjectACL != nil {
		in, out := &in.DefaultObjectACL, &out.DefaultObjectACL
		*out = make([]ACLRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketSpecAttrs.
func (in *BucketSpecAttrs) DeepCopy() *BucketSpecAttrs {
	if in == nil {
		return nil
	}
	out := new(BucketSpecAttrs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketStatus) DeepCopyInto(out *BucketStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.BucketOutputAttrs.DeepCopyInto(&out.BucketOutputAttrs)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketStatus.
func (in *BucketStatus) DeepCopy() *BucketStatus {
	if in == nil {
		return nil
	}
	out := new(BucketStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketUpdatableAttrs) DeepCopyInto(out *BucketUpdatableAttrs) {
	*out = *in
	if in.BucketPolicyOnly != nil {
		in, out := &in.BucketPolicyOnly, &out.BucketPolicyOnly
		*out = new(BucketPolicyOnly)
		(*in).DeepCopyInto(*out)
	}
	if in.CORS != nil {
		in, out := &in.CORS, &out.CORS
		*out = make([]CORS, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Encryption != nil {
		in, out := &in.Encryption, &out.Encryption
		*out = new(BucketEncryption)
		**out = **in
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.Lifecycle.DeepCopyInto(&out.Lifecycle)
	if in.Logging != nil {
		in, out := &in.Logging, &out.Logging
		*out = new(BucketLogging)
		**out = **in
	}
	if in.RetentionPolicy != nil {
		in, out := &in.RetentionPolicy, &out.RetentionPolicy
		*out = new(RetentionPolicy)
		**out = **in
	}
	if in.Website != nil {
		in, out := &in.Website, &out.Website
		*out = new(BucketWebsite)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketUpdatableAttrs.
func (in *BucketUpdatableAttrs) DeepCopy() *BucketUpdatableAttrs {
	if in == nil {
		return nil
	}
	out := new(BucketUpdatableAttrs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketWebsite) DeepCopyInto(out *BucketWebsite) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketWebsite.
func (in *BucketWebsite) DeepCopy() *BucketWebsite {
	if in == nil {
		return nil
	}
	out := new(BucketWebsite)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CORS) DeepCopyInto(out *CORS) {
	*out = *in
	out.MaxAge = in.MaxAge
	if in.Methods != nil {
		in, out := &in.Methods, &out.Methods
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Origins != nil {
		in, out := &in.Origins, &out.Origins
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ResponseHeaders != nil {
		in, out := &in.ResponseHeaders, &out.ResponseHeaders
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CORS.
func (in *CORS) DeepCopy() *CORS {
	if in == nil {
		return nil
	}
	out := new(CORS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Lifecycle) DeepCopyInto(out *Lifecycle) {
	*out = *in
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]LifecycleRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Lifecycle.
func (in *Lifecycle) DeepCopy() *Lifecycle {
	if in == nil {
		return nil
	}
	out := new(Lifecycle)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LifecycleAction) DeepCopyInto(out *LifecycleAction) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LifecycleAction.
func (in *LifecycleAction) DeepCopy() *LifecycleAction {
	if in == nil {
		return nil
	}
	out := new(LifecycleAction)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LifecycleCondition) DeepCopyInto(out *LifecycleCondition) {
	*out = *in
	in.CreatedBefore.DeepCopyInto(&out.CreatedBefore)
	if in.MatchesStorageClasses != nil {
		in, out := &in.MatchesStorageClasses, &out.MatchesStorageClasses
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LifecycleCondition.
func (in *LifecycleCondition) DeepCopy() *LifecycleCondition {
	if in == nil {
		return nil
	}
	out := new(LifecycleCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LifecycleRule) DeepCopyInto(out *LifecycleRule) {
	*out = *in
	out.Action = in.Action
	in.Condition.DeepCopyInto(&out.Condition)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LifecycleRule.
func (in *LifecycleRule) DeepCopy() *LifecycleRule {
	if in == nil {
		return nil
	}
	out := new(LifecycleRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectTeam) DeepCopyInto(out *ProjectTeam) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectTeam.
func (in *ProjectTeam) DeepCopy() *ProjectTeam {
	if in == nil {
		return nil
	}
	out := new(ProjectTeam)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RetentionPolicy) DeepCopyInto(out *RetentionPolicy) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RetentionPolicy.
func (in *RetentionPolicy) DeepCopy() *RetentionPolicy {
	if in == nil {
		return nil
	}
	out := new(RetentionPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RetentionPolicyStatus) DeepCopyInto(out *RetentionPolicyStatus) {
	*out = *in
	in.EffectiveTime.DeepCopyInto(&out.EffectiveTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RetentionPolicyStatus.
func (in *RetentionPolicyStatus) DeepCopy() *RetentionPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(RetentionPolicyStatus)
	in.DeepCopyInto(out)
	return out
}
