//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AttributeFeatureSet) DeepCopyInto(out *AttributeFeatureSet) {
	*out = *in
	if in.Elements != nil {
		in, out := &in.Elements, &out.Elements
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AttributeFeatureSet.
func (in *AttributeFeatureSet) DeepCopy() *AttributeFeatureSet {
	if in == nil {
		return nil
	}
	out := new(AttributeFeatureSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in FeatureMatcher) DeepCopyInto(out *FeatureMatcher) {
	{
		in := &in
		*out = make(FeatureMatcher, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeatureMatcher.
func (in FeatureMatcher) DeepCopy() FeatureMatcher {
	if in == nil {
		return nil
	}
	out := new(FeatureMatcher)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FeatureMatcherTerm) DeepCopyInto(out *FeatureMatcherTerm) {
	*out = *in
	if in.MatchExpressions != nil {
		in, out := &in.MatchExpressions, &out.MatchExpressions
		*out = new(map[string]*MatchExpression)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[string]*MatchExpression, len(*in))
			for key, val := range *in {
				var outVal *MatchExpression
				if val == nil {
					(*out)[key] = nil
				} else {
					in, out := &val, &outVal
					*out = new(MatchExpression)
					(*in).DeepCopyInto(*out)
				}
				(*out)[key] = outVal
			}
		}
	}
	if in.MatchName != nil {
		in, out := &in.MatchName, &out.MatchName
		*out = new(MatchExpression)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeatureMatcherTerm.
func (in *FeatureMatcherTerm) DeepCopy() *FeatureMatcherTerm {
	if in == nil {
		return nil
	}
	out := new(FeatureMatcherTerm)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Features) DeepCopyInto(out *Features) {
	*out = *in
	if in.Flags != nil {
		in, out := &in.Flags, &out.Flags
		*out = make(map[string]FlagFeatureSet, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.Attributes != nil {
		in, out := &in.Attributes, &out.Attributes
		*out = make(map[string]AttributeFeatureSet, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.Instances != nil {
		in, out := &in.Instances, &out.Instances
		*out = make(map[string]InstanceFeatureSet, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Features.
func (in *Features) DeepCopy() *Features {
	if in == nil {
		return nil
	}
	out := new(Features)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlagFeatureSet) DeepCopyInto(out *FlagFeatureSet) {
	*out = *in
	if in.Elements != nil {
		in, out := &in.Elements, &out.Elements
		*out = make(map[string]Nil, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlagFeatureSet.
func (in *FlagFeatureSet) DeepCopy() *FlagFeatureSet {
	if in == nil {
		return nil
	}
	out := new(FlagFeatureSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceFeature) DeepCopyInto(out *InstanceFeature) {
	*out = *in
	if in.Attributes != nil {
		in, out := &in.Attributes, &out.Attributes
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceFeature.
func (in *InstanceFeature) DeepCopy() *InstanceFeature {
	if in == nil {
		return nil
	}
	out := new(InstanceFeature)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceFeatureSet) DeepCopyInto(out *InstanceFeatureSet) {
	*out = *in
	if in.Elements != nil {
		in, out := &in.Elements, &out.Elements
		*out = make([]InstanceFeature, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceFeatureSet.
func (in *InstanceFeatureSet) DeepCopy() *InstanceFeatureSet {
	if in == nil {
		return nil
	}
	out := new(InstanceFeatureSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MatchAnyElem) DeepCopyInto(out *MatchAnyElem) {
	*out = *in
	if in.MatchFeatures != nil {
		in, out := &in.MatchFeatures, &out.MatchFeatures
		*out = make(FeatureMatcher, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MatchAnyElem.
func (in *MatchAnyElem) DeepCopy() *MatchAnyElem {
	if in == nil {
		return nil
	}
	out := new(MatchAnyElem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MatchExpression) DeepCopyInto(out *MatchExpression) {
	*out = *in
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = make(MatchValue, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MatchExpression.
func (in *MatchExpression) DeepCopy() *MatchExpression {
	if in == nil {
		return nil
	}
	out := new(MatchExpression)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in MatchExpressionSet) DeepCopyInto(out *MatchExpressionSet) {
	{
		in := &in
		*out = make(MatchExpressionSet, len(*in))
		for key, val := range *in {
			var outVal *MatchExpression
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(MatchExpression)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MatchExpressionSet.
func (in MatchExpressionSet) DeepCopy() MatchExpressionSet {
	if in == nil {
		return nil
	}
	out := new(MatchExpressionSet)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in MatchValue) DeepCopyInto(out *MatchValue) {
	{
		in := &in
		*out = make(MatchValue, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MatchValue.
func (in MatchValue) DeepCopy() MatchValue {
	if in == nil {
		return nil
	}
	out := new(MatchValue)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Nil) DeepCopyInto(out *Nil) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Nil.
func (in *Nil) DeepCopy() *Nil {
	if in == nil {
		return nil
	}
	out := new(Nil)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeFeature) DeepCopyInto(out *NodeFeature) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeFeature.
func (in *NodeFeature) DeepCopy() *NodeFeature {
	if in == nil {
		return nil
	}
	out := new(NodeFeature)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodeFeature) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeFeatureList) DeepCopyInto(out *NodeFeatureList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NodeFeature, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeFeatureList.
func (in *NodeFeatureList) DeepCopy() *NodeFeatureList {
	if in == nil {
		return nil
	}
	out := new(NodeFeatureList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodeFeatureList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeFeatureRule) DeepCopyInto(out *NodeFeatureRule) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeFeatureRule.
func (in *NodeFeatureRule) DeepCopy() *NodeFeatureRule {
	if in == nil {
		return nil
	}
	out := new(NodeFeatureRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodeFeatureRule) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeFeatureRuleList) DeepCopyInto(out *NodeFeatureRuleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NodeFeatureRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeFeatureRuleList.
func (in *NodeFeatureRuleList) DeepCopy() *NodeFeatureRuleList {
	if in == nil {
		return nil
	}
	out := new(NodeFeatureRuleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodeFeatureRuleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeFeatureRuleSpec) DeepCopyInto(out *NodeFeatureRuleSpec) {
	*out = *in
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]Rule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeFeatureRuleSpec.
func (in *NodeFeatureRuleSpec) DeepCopy() *NodeFeatureRuleSpec {
	if in == nil {
		return nil
	}
	out := new(NodeFeatureRuleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeFeatureSpec) DeepCopyInto(out *NodeFeatureSpec) {
	*out = *in
	in.Features.DeepCopyInto(&out.Features)
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeFeatureSpec.
func (in *NodeFeatureSpec) DeepCopy() *NodeFeatureSpec {
	if in == nil {
		return nil
	}
	out := new(NodeFeatureSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Rule) DeepCopyInto(out *Rule) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Vars != nil {
		in, out := &in.Vars, &out.Vars
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Taints != nil {
		in, out := &in.Taints, &out.Taints
		*out = make([]v1.Taint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ExtendedResources != nil {
		in, out := &in.ExtendedResources, &out.ExtendedResources
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.MatchFeatures != nil {
		in, out := &in.MatchFeatures, &out.MatchFeatures
		*out = make(FeatureMatcher, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MatchAny != nil {
		in, out := &in.MatchAny, &out.MatchAny
		*out = make([]MatchAnyElem, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.labelsTemplate != nil {
		in, out := &in.labelsTemplate, &out.labelsTemplate
		*out = (*in).DeepCopy()
	}
	if in.varsTemplate != nil {
		in, out := &in.varsTemplate, &out.varsTemplate
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Rule.
func (in *Rule) DeepCopy() *Rule {
	if in == nil {
		return nil
	}
	out := new(Rule)
	in.DeepCopyInto(out)
	return out
}
