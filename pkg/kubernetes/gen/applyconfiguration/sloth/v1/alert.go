// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// AlertApplyConfiguration represents a declarative configuration of the Alert type for use
// with apply.
type AlertApplyConfiguration struct {
	Disable     *bool             `json:"disable,omitempty"`
	Labels      map[string]string `json:"labels,omitempty"`
	Annotations map[string]string `json:"annotations,omitempty"`
}

// AlertApplyConfiguration constructs a declarative configuration of the Alert type for use with
// apply.
func Alert() *AlertApplyConfiguration {
	return &AlertApplyConfiguration{}
}

// WithDisable sets the Disable field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Disable field is set to the value of the last call.
func (b *AlertApplyConfiguration) WithDisable(value bool) *AlertApplyConfiguration {
	b.Disable = &value
	return b
}

// WithLabels puts the entries into the Labels field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Labels field,
// overwriting an existing map entries in Labels field with the same key.
func (b *AlertApplyConfiguration) WithLabels(entries map[string]string) *AlertApplyConfiguration {
	if b.Labels == nil && len(entries) > 0 {
		b.Labels = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.Labels[k] = v
	}
	return b
}

// WithAnnotations puts the entries into the Annotations field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Annotations field,
// overwriting an existing map entries in Annotations field with the same key.
func (b *AlertApplyConfiguration) WithAnnotations(entries map[string]string) *AlertApplyConfiguration {
	if b.Annotations == nil && len(entries) > 0 {
		b.Annotations = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.Annotations[k] = v
	}
	return b
}
