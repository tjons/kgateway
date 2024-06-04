// Code generated by solo-kit. DO NOT EDIT.

package v1

import (
	"encoding/json"

	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"github.com/solo-io/solo-kit/pkg/utils/protoutils"

	api "github.com/solo-io/gloo/projects/gateway/pkg/api/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type metaOnly struct {
	v1.TypeMeta   `json:",inline"`
	v1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +resourceName=gateways
// +genclient
type Gateway struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Spec defines the implementation of this definition.
	// +optional
	Spec   api.Gateway             `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status core.NamespacedStatuses `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

func (o *Gateway) MarshalJSON() ([]byte, error) {
	spec, err := protoutils.MarshalMap(&o.Spec)
	if err != nil {
		return nil, err
	}
	delete(spec, "metadata")
	delete(spec, "namespacedStatuses")
	asMap := map[string]interface{}{
		"metadata":   o.ObjectMeta,
		"apiVersion": o.TypeMeta.APIVersion,
		"kind":       o.TypeMeta.Kind,
		"status":     o.Status,
		"spec":       spec,
	}
	return json.Marshal(asMap)
}

func (o *Gateway) UnmarshalJSON(data []byte) error {
	var metaOnly metaOnly
	if err := json.Unmarshal(data, &metaOnly); err != nil {
		return err
	}
	var spec api.Gateway
	if err := protoutils.UnmarshalResource(data, &spec); err != nil {
		return err
	}
	spec.Metadata = nil
	*o = Gateway{
		ObjectMeta: metaOnly.ObjectMeta,
		TypeMeta:   metaOnly.TypeMeta,
		Spec:       spec,
	}
	if spec.GetNamespacedStatuses() != nil {
		o.Status = *spec.NamespacedStatuses
		o.Spec.NamespacedStatuses = nil
	}

	return nil
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// GatewayList is a collection of Gateways.
type GatewayList struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items       []Gateway `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +resourceName=httplisteneroptions
// +genclient
// +genclient:noStatus
type HttpListenerOption struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Spec defines the implementation of this definition.
	// +optional
	Spec api.HttpListenerOption `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
}

func (o *HttpListenerOption) MarshalJSON() ([]byte, error) {
	spec, err := protoutils.MarshalMap(&o.Spec)
	if err != nil {
		return nil, err
	}
	delete(spec, "metadata")
	asMap := map[string]interface{}{
		"metadata":   o.ObjectMeta,
		"apiVersion": o.TypeMeta.APIVersion,
		"kind":       o.TypeMeta.Kind,
		"spec":       spec,
	}
	return json.Marshal(asMap)
}

func (o *HttpListenerOption) UnmarshalJSON(data []byte) error {
	var metaOnly metaOnly
	if err := json.Unmarshal(data, &metaOnly); err != nil {
		return err
	}
	var spec api.HttpListenerOption
	if err := protoutils.UnmarshalResource(data, &spec); err != nil {
		return err
	}
	spec.Metadata = nil
	*o = HttpListenerOption{
		ObjectMeta: metaOnly.ObjectMeta,
		TypeMeta:   metaOnly.TypeMeta,
		Spec:       spec,
	}

	return nil
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// HttpListenerOptionList is a collection of HttpListenerOptions.
type HttpListenerOptionList struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items       []HttpListenerOption `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +resourceName=listeneroptions
// +genclient
// +genclient:noStatus
type ListenerOption struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Spec defines the implementation of this definition.
	// +optional
	Spec api.ListenerOption `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
}

func (o *ListenerOption) MarshalJSON() ([]byte, error) {
	spec, err := protoutils.MarshalMap(&o.Spec)
	if err != nil {
		return nil, err
	}
	delete(spec, "metadata")
	asMap := map[string]interface{}{
		"metadata":   o.ObjectMeta,
		"apiVersion": o.TypeMeta.APIVersion,
		"kind":       o.TypeMeta.Kind,
		"spec":       spec,
	}
	return json.Marshal(asMap)
}

func (o *ListenerOption) UnmarshalJSON(data []byte) error {
	var metaOnly metaOnly
	if err := json.Unmarshal(data, &metaOnly); err != nil {
		return err
	}
	var spec api.ListenerOption
	if err := protoutils.UnmarshalResource(data, &spec); err != nil {
		return err
	}
	spec.Metadata = nil
	*o = ListenerOption{
		ObjectMeta: metaOnly.ObjectMeta,
		TypeMeta:   metaOnly.TypeMeta,
		Spec:       spec,
	}

	return nil
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ListenerOptionList is a collection of ListenerOptions.
type ListenerOptionList struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items       []ListenerOption `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +resourceName=httpgateways
// +genclient
type MatchableHttpGateway struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Spec defines the implementation of this definition.
	// +optional
	Spec   api.MatchableHttpGateway `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status core.NamespacedStatuses  `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

func (o *MatchableHttpGateway) MarshalJSON() ([]byte, error) {
	spec, err := protoutils.MarshalMap(&o.Spec)
	if err != nil {
		return nil, err
	}
	delete(spec, "metadata")
	delete(spec, "namespacedStatuses")
	asMap := map[string]interface{}{
		"metadata":   o.ObjectMeta,
		"apiVersion": o.TypeMeta.APIVersion,
		"kind":       o.TypeMeta.Kind,
		"status":     o.Status,
		"spec":       spec,
	}
	return json.Marshal(asMap)
}

func (o *MatchableHttpGateway) UnmarshalJSON(data []byte) error {
	var metaOnly metaOnly
	if err := json.Unmarshal(data, &metaOnly); err != nil {
		return err
	}
	var spec api.MatchableHttpGateway
	if err := protoutils.UnmarshalResource(data, &spec); err != nil {
		return err
	}
	spec.Metadata = nil
	*o = MatchableHttpGateway{
		ObjectMeta: metaOnly.ObjectMeta,
		TypeMeta:   metaOnly.TypeMeta,
		Spec:       spec,
	}
	if spec.GetNamespacedStatuses() != nil {
		o.Status = *spec.NamespacedStatuses
		o.Spec.NamespacedStatuses = nil
	}

	return nil
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// MatchableHttpGatewayList is a collection of MatchableHttpGateways.
type MatchableHttpGatewayList struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items       []MatchableHttpGateway `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +resourceName=tcpgateways
// +genclient
type MatchableTcpGateway struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Spec defines the implementation of this definition.
	// +optional
	Spec   api.MatchableTcpGateway `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status core.NamespacedStatuses `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

func (o *MatchableTcpGateway) MarshalJSON() ([]byte, error) {
	spec, err := protoutils.MarshalMap(&o.Spec)
	if err != nil {
		return nil, err
	}
	delete(spec, "metadata")
	delete(spec, "namespacedStatuses")
	asMap := map[string]interface{}{
		"metadata":   o.ObjectMeta,
		"apiVersion": o.TypeMeta.APIVersion,
		"kind":       o.TypeMeta.Kind,
		"status":     o.Status,
		"spec":       spec,
	}
	return json.Marshal(asMap)
}

func (o *MatchableTcpGateway) UnmarshalJSON(data []byte) error {
	var metaOnly metaOnly
	if err := json.Unmarshal(data, &metaOnly); err != nil {
		return err
	}
	var spec api.MatchableTcpGateway
	if err := protoutils.UnmarshalResource(data, &spec); err != nil {
		return err
	}
	spec.Metadata = nil
	*o = MatchableTcpGateway{
		ObjectMeta: metaOnly.ObjectMeta,
		TypeMeta:   metaOnly.TypeMeta,
		Spec:       spec,
	}
	if spec.GetNamespacedStatuses() != nil {
		o.Status = *spec.NamespacedStatuses
		o.Spec.NamespacedStatuses = nil
	}

	return nil
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// MatchableTcpGatewayList is a collection of MatchableTcpGateways.
type MatchableTcpGatewayList struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items       []MatchableTcpGateway `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +resourceName=routeoptions
// +genclient
type RouteOption struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Spec defines the implementation of this definition.
	// +optional
	Spec   api.RouteOption         `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status core.NamespacedStatuses `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

func (o *RouteOption) MarshalJSON() ([]byte, error) {
	spec, err := protoutils.MarshalMap(&o.Spec)
	if err != nil {
		return nil, err
	}
	delete(spec, "metadata")
	delete(spec, "namespacedStatuses")
	asMap := map[string]interface{}{
		"metadata":   o.ObjectMeta,
		"apiVersion": o.TypeMeta.APIVersion,
		"kind":       o.TypeMeta.Kind,
		"status":     o.Status,
		"spec":       spec,
	}
	return json.Marshal(asMap)
}

func (o *RouteOption) UnmarshalJSON(data []byte) error {
	var metaOnly metaOnly
	if err := json.Unmarshal(data, &metaOnly); err != nil {
		return err
	}
	var spec api.RouteOption
	if err := protoutils.UnmarshalResource(data, &spec); err != nil {
		return err
	}
	spec.Metadata = nil
	*o = RouteOption{
		ObjectMeta: metaOnly.ObjectMeta,
		TypeMeta:   metaOnly.TypeMeta,
		Spec:       spec,
	}
	if spec.GetNamespacedStatuses() != nil {
		o.Status = *spec.NamespacedStatuses
		o.Spec.NamespacedStatuses = nil
	}

	return nil
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// RouteOptionList is a collection of RouteOptions.
type RouteOptionList struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items       []RouteOption `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +resourceName=routetables
// +genclient
type RouteTable struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Spec defines the implementation of this definition.
	// +optional
	Spec   api.RouteTable          `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status core.NamespacedStatuses `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

func (o *RouteTable) MarshalJSON() ([]byte, error) {
	spec, err := protoutils.MarshalMap(&o.Spec)
	if err != nil {
		return nil, err
	}
	delete(spec, "metadata")
	delete(spec, "namespacedStatuses")
	asMap := map[string]interface{}{
		"metadata":   o.ObjectMeta,
		"apiVersion": o.TypeMeta.APIVersion,
		"kind":       o.TypeMeta.Kind,
		"status":     o.Status,
		"spec":       spec,
	}
	return json.Marshal(asMap)
}

func (o *RouteTable) UnmarshalJSON(data []byte) error {
	var metaOnly metaOnly
	if err := json.Unmarshal(data, &metaOnly); err != nil {
		return err
	}
	var spec api.RouteTable
	if err := protoutils.UnmarshalResource(data, &spec); err != nil {
		return err
	}
	spec.Metadata = nil
	*o = RouteTable{
		ObjectMeta: metaOnly.ObjectMeta,
		TypeMeta:   metaOnly.TypeMeta,
		Spec:       spec,
	}
	if spec.GetNamespacedStatuses() != nil {
		o.Status = *spec.NamespacedStatuses
		o.Spec.NamespacedStatuses = nil
	}

	return nil
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// RouteTableList is a collection of RouteTables.
type RouteTableList struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items       []RouteTable `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +resourceName=virtualhostoptions
// +genclient
type VirtualHostOption struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Spec defines the implementation of this definition.
	// +optional
	Spec   api.VirtualHostOption   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status core.NamespacedStatuses `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

func (o *VirtualHostOption) MarshalJSON() ([]byte, error) {
	spec, err := protoutils.MarshalMap(&o.Spec)
	if err != nil {
		return nil, err
	}
	delete(spec, "metadata")
	delete(spec, "namespacedStatuses")
	asMap := map[string]interface{}{
		"metadata":   o.ObjectMeta,
		"apiVersion": o.TypeMeta.APIVersion,
		"kind":       o.TypeMeta.Kind,
		"status":     o.Status,
		"spec":       spec,
	}
	return json.Marshal(asMap)
}

func (o *VirtualHostOption) UnmarshalJSON(data []byte) error {
	var metaOnly metaOnly
	if err := json.Unmarshal(data, &metaOnly); err != nil {
		return err
	}
	var spec api.VirtualHostOption
	if err := protoutils.UnmarshalResource(data, &spec); err != nil {
		return err
	}
	spec.Metadata = nil
	*o = VirtualHostOption{
		ObjectMeta: metaOnly.ObjectMeta,
		TypeMeta:   metaOnly.TypeMeta,
		Spec:       spec,
	}
	if spec.GetNamespacedStatuses() != nil {
		o.Status = *spec.NamespacedStatuses
		o.Spec.NamespacedStatuses = nil
	}

	return nil
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// VirtualHostOptionList is a collection of VirtualHostOptions.
type VirtualHostOptionList struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items       []VirtualHostOption `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +resourceName=virtualservices
// +genclient
type VirtualService struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Spec defines the implementation of this definition.
	// +optional
	Spec   api.VirtualService      `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status core.NamespacedStatuses `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

func (o *VirtualService) MarshalJSON() ([]byte, error) {
	spec, err := protoutils.MarshalMap(&o.Spec)
	if err != nil {
		return nil, err
	}
	delete(spec, "metadata")
	delete(spec, "namespacedStatuses")
	asMap := map[string]interface{}{
		"metadata":   o.ObjectMeta,
		"apiVersion": o.TypeMeta.APIVersion,
		"kind":       o.TypeMeta.Kind,
		"status":     o.Status,
		"spec":       spec,
	}
	return json.Marshal(asMap)
}

func (o *VirtualService) UnmarshalJSON(data []byte) error {
	var metaOnly metaOnly
	if err := json.Unmarshal(data, &metaOnly); err != nil {
		return err
	}
	var spec api.VirtualService
	if err := protoutils.UnmarshalResource(data, &spec); err != nil {
		return err
	}
	spec.Metadata = nil
	*o = VirtualService{
		ObjectMeta: metaOnly.ObjectMeta,
		TypeMeta:   metaOnly.TypeMeta,
		Spec:       spec,
	}
	if spec.GetNamespacedStatuses() != nil {
		o.Status = *spec.NamespacedStatuses
		o.Spec.NamespacedStatuses = nil
	}

	return nil
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// VirtualServiceList is a collection of VirtualServices.
type VirtualServiceList struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items       []VirtualService `json:"items" protobuf:"bytes,2,rep,name=items"`
}
