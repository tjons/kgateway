// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/client-go/applyconfigurations/meta/v1"
)

// UpstreamStatusApplyConfiguration represents a declarative configuration of the UpstreamStatus type for use
// with apply.
type UpstreamStatusApplyConfiguration struct {
	Conditions []v1.ConditionApplyConfiguration `json:"conditions,omitempty"`
}

// UpstreamStatusApplyConfiguration constructs a declarative configuration of the UpstreamStatus type for use with
// apply.
func UpstreamStatus() *UpstreamStatusApplyConfiguration {
	return &UpstreamStatusApplyConfiguration{}
}

// WithConditions adds the given value to the Conditions field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Conditions field.
func (b *UpstreamStatusApplyConfiguration) WithConditions(values ...*v1.ConditionApplyConfiguration) *UpstreamStatusApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithConditions")
		}
		b.Conditions = append(b.Conditions, *values[i])
	}
	return b
}
