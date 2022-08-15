// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

package gql

import (
	"context"

	"github.com/Khan/genqlient/graphql"
)

type AddOnType string

const (
	// A Redis database
	AddOnTypeRedis AddOnType = "redis"
)

// CreateAddOnCreateAddOnCreateAddOnPayload includes the requested fields of the GraphQL type CreateAddOnPayload.
// The GraphQL type's documentation follows.
//
// Autogenerated return type of CreateAddOn
type CreateAddOnCreateAddOnCreateAddOnPayload struct {
	AddOn CreateAddOnCreateAddOnCreateAddOnPayloadAddOn `json:"addOn"`
}

// GetAddOn returns CreateAddOnCreateAddOnCreateAddOnPayload.AddOn, and is useful for accessing the field via an interface.
func (v *CreateAddOnCreateAddOnCreateAddOnPayload) GetAddOn() CreateAddOnCreateAddOnCreateAddOnPayloadAddOn {
	return v.AddOn
}

// CreateAddOnCreateAddOnCreateAddOnPayloadAddOn includes the requested fields of the GraphQL type AddOn.
type CreateAddOnCreateAddOnCreateAddOnPayloadAddOn struct {
	Id string `json:"id"`
	// Public URL for this service
	PublicUrl string `json:"publicUrl"`
}

// GetId returns CreateAddOnCreateAddOnCreateAddOnPayloadAddOn.Id, and is useful for accessing the field via an interface.
func (v *CreateAddOnCreateAddOnCreateAddOnPayloadAddOn) GetId() string { return v.Id }

// GetPublicUrl returns CreateAddOnCreateAddOnCreateAddOnPayloadAddOn.PublicUrl, and is useful for accessing the field via an interface.
func (v *CreateAddOnCreateAddOnCreateAddOnPayloadAddOn) GetPublicUrl() string { return v.PublicUrl }

// CreateAddOnResponse is returned by CreateAddOn on success.
type CreateAddOnResponse struct {
	CreateAddOn CreateAddOnCreateAddOnCreateAddOnPayload `json:"createAddOn"`
}

// GetCreateAddOn returns CreateAddOnResponse.CreateAddOn, and is useful for accessing the field via an interface.
func (v *CreateAddOnResponse) GetCreateAddOn() CreateAddOnCreateAddOnCreateAddOnPayload {
	return v.CreateAddOn
}

// DeleteAddOnDeleteAddOnDeleteAddOnPayload includes the requested fields of the GraphQL type DeleteAddOnPayload.
// The GraphQL type's documentation follows.
//
// Autogenerated return type of DeleteAddOn
type DeleteAddOnDeleteAddOnDeleteAddOnPayload struct {
	DeletedAddOnId string `json:"deletedAddOnId"`
}

// GetDeletedAddOnId returns DeleteAddOnDeleteAddOnDeleteAddOnPayload.DeletedAddOnId, and is useful for accessing the field via an interface.
func (v *DeleteAddOnDeleteAddOnDeleteAddOnPayload) GetDeletedAddOnId() string {
	return v.DeletedAddOnId
}

// DeleteAddOnResponse is returned by DeleteAddOn on success.
type DeleteAddOnResponse struct {
	DeleteAddOn DeleteAddOnDeleteAddOnDeleteAddOnPayload `json:"deleteAddOn"`
}

// GetDeleteAddOn returns DeleteAddOnResponse.DeleteAddOn, and is useful for accessing the field via an interface.
func (v *DeleteAddOnResponse) GetDeleteAddOn() DeleteAddOnDeleteAddOnDeleteAddOnPayload {
	return v.DeleteAddOn
}

// GetAddOnAddOn includes the requested fields of the GraphQL type AddOn.
type GetAddOnAddOn struct {
	Id string `json:"id"`
	// The service name according to the provider
	Name string `json:"name"`
	// Public URL for this service
	PublicUrl string `json:"publicUrl"`
	// Region where the primary instance is deployed
	PrimaryRegion string `json:"primaryRegion"`
	// Regions where replica instances are deployed
	ReadRegions []string `json:"readRegions"`
	// The add-on plan
	AddOnPlan GetAddOnAddOnAddOnPlan `json:"addOnPlan"`
}

// GetId returns GetAddOnAddOn.Id, and is useful for accessing the field via an interface.
func (v *GetAddOnAddOn) GetId() string { return v.Id }

// GetName returns GetAddOnAddOn.Name, and is useful for accessing the field via an interface.
func (v *GetAddOnAddOn) GetName() string { return v.Name }

// GetPublicUrl returns GetAddOnAddOn.PublicUrl, and is useful for accessing the field via an interface.
func (v *GetAddOnAddOn) GetPublicUrl() string { return v.PublicUrl }

// GetPrimaryRegion returns GetAddOnAddOn.PrimaryRegion, and is useful for accessing the field via an interface.
func (v *GetAddOnAddOn) GetPrimaryRegion() string { return v.PrimaryRegion }

// GetReadRegions returns GetAddOnAddOn.ReadRegions, and is useful for accessing the field via an interface.
func (v *GetAddOnAddOn) GetReadRegions() []string { return v.ReadRegions }

// GetAddOnPlan returns GetAddOnAddOn.AddOnPlan, and is useful for accessing the field via an interface.
func (v *GetAddOnAddOn) GetAddOnPlan() GetAddOnAddOnAddOnPlan { return v.AddOnPlan }

// GetAddOnAddOnAddOnPlan includes the requested fields of the GraphQL type AddOnPlan.
type GetAddOnAddOnAddOnPlan struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
}

// GetId returns GetAddOnAddOnAddOnPlan.Id, and is useful for accessing the field via an interface.
func (v *GetAddOnAddOnAddOnPlan) GetId() string { return v.Id }

// GetName returns GetAddOnAddOnAddOnPlan.Name, and is useful for accessing the field via an interface.
func (v *GetAddOnAddOnAddOnPlan) GetName() string { return v.Name }

// GetDisplayName returns GetAddOnAddOnAddOnPlan.DisplayName, and is useful for accessing the field via an interface.
func (v *GetAddOnAddOnAddOnPlan) GetDisplayName() string { return v.DisplayName }

// GetAddOnResponse is returned by GetAddOn on success.
type GetAddOnResponse struct {
	// Find an add-on by ID
	AddOn GetAddOnAddOn `json:"addOn"`
}

// GetAddOn returns GetAddOnResponse.AddOn, and is useful for accessing the field via an interface.
func (v *GetAddOnResponse) GetAddOn() GetAddOnAddOn { return v.AddOn }

// ListAddOnPlansAddOnPlansAddOnPlanConnection includes the requested fields of the GraphQL type AddOnPlanConnection.
// The GraphQL type's documentation follows.
//
// The connection type for AddOnPlan.
type ListAddOnPlansAddOnPlansAddOnPlanConnection struct {
	// A list of nodes.
	Nodes []ListAddOnPlansAddOnPlansAddOnPlanConnectionNodesAddOnPlan `json:"nodes"`
}

// GetNodes returns ListAddOnPlansAddOnPlansAddOnPlanConnection.Nodes, and is useful for accessing the field via an interface.
func (v *ListAddOnPlansAddOnPlansAddOnPlanConnection) GetNodes() []ListAddOnPlansAddOnPlansAddOnPlanConnectionNodesAddOnPlan {
	return v.Nodes
}

// ListAddOnPlansAddOnPlansAddOnPlanConnectionNodesAddOnPlan includes the requested fields of the GraphQL type AddOnPlan.
type ListAddOnPlansAddOnPlansAddOnPlanConnectionNodesAddOnPlan struct {
	Id            string `json:"id"`
	DisplayName   string `json:"displayName"`
	MaxDataSize   string `json:"maxDataSize"`
	PricePerMonth int    `json:"pricePerMonth"`
}

// GetId returns ListAddOnPlansAddOnPlansAddOnPlanConnectionNodesAddOnPlan.Id, and is useful for accessing the field via an interface.
func (v *ListAddOnPlansAddOnPlansAddOnPlanConnectionNodesAddOnPlan) GetId() string { return v.Id }

// GetDisplayName returns ListAddOnPlansAddOnPlansAddOnPlanConnectionNodesAddOnPlan.DisplayName, and is useful for accessing the field via an interface.
func (v *ListAddOnPlansAddOnPlansAddOnPlanConnectionNodesAddOnPlan) GetDisplayName() string {
	return v.DisplayName
}

// GetMaxDataSize returns ListAddOnPlansAddOnPlansAddOnPlanConnectionNodesAddOnPlan.MaxDataSize, and is useful for accessing the field via an interface.
func (v *ListAddOnPlansAddOnPlansAddOnPlanConnectionNodesAddOnPlan) GetMaxDataSize() string {
	return v.MaxDataSize
}

// GetPricePerMonth returns ListAddOnPlansAddOnPlansAddOnPlanConnectionNodesAddOnPlan.PricePerMonth, and is useful for accessing the field via an interface.
func (v *ListAddOnPlansAddOnPlansAddOnPlanConnectionNodesAddOnPlan) GetPricePerMonth() int {
	return v.PricePerMonth
}

// ListAddOnPlansResponse is returned by ListAddOnPlans on success.
type ListAddOnPlansResponse struct {
	// List add-on service plans
	AddOnPlans ListAddOnPlansAddOnPlansAddOnPlanConnection `json:"addOnPlans"`
}

// GetAddOnPlans returns ListAddOnPlansResponse.AddOnPlans, and is useful for accessing the field via an interface.
func (v *ListAddOnPlansResponse) GetAddOnPlans() ListAddOnPlansAddOnPlansAddOnPlanConnection {
	return v.AddOnPlans
}

// ListAddOnsAddOnsAddOnConnection includes the requested fields of the GraphQL type AddOnConnection.
// The GraphQL type's documentation follows.
//
// The connection type for AddOn.
type ListAddOnsAddOnsAddOnConnection struct {
	// A list of nodes.
	Nodes []ListAddOnsAddOnsAddOnConnectionNodesAddOn `json:"nodes"`
}

// GetNodes returns ListAddOnsAddOnsAddOnConnection.Nodes, and is useful for accessing the field via an interface.
func (v *ListAddOnsAddOnsAddOnConnection) GetNodes() []ListAddOnsAddOnsAddOnConnectionNodesAddOn {
	return v.Nodes
}

// ListAddOnsAddOnsAddOnConnectionNodesAddOn includes the requested fields of the GraphQL type AddOn.
type ListAddOnsAddOnsAddOnConnectionNodesAddOn struct {
	Id string `json:"id"`
	// The service name according to the provider
	Name string `json:"name"`
	// The add-on plan
	AddOnPlan ListAddOnsAddOnsAddOnConnectionNodesAddOnAddOnPlan `json:"addOnPlan"`
	// Region where the primary instance is deployed
	PrimaryRegion string `json:"primaryRegion"`
	// Regions where replica instances are deployed
	ReadRegions []string `json:"readRegions"`
	// Organization that owns this service
	Organization ListAddOnsAddOnsAddOnConnectionNodesAddOnOrganization `json:"organization"`
}

// GetId returns ListAddOnsAddOnsAddOnConnectionNodesAddOn.Id, and is useful for accessing the field via an interface.
func (v *ListAddOnsAddOnsAddOnConnectionNodesAddOn) GetId() string { return v.Id }

// GetName returns ListAddOnsAddOnsAddOnConnectionNodesAddOn.Name, and is useful for accessing the field via an interface.
func (v *ListAddOnsAddOnsAddOnConnectionNodesAddOn) GetName() string { return v.Name }

// GetAddOnPlan returns ListAddOnsAddOnsAddOnConnectionNodesAddOn.AddOnPlan, and is useful for accessing the field via an interface.
func (v *ListAddOnsAddOnsAddOnConnectionNodesAddOn) GetAddOnPlan() ListAddOnsAddOnsAddOnConnectionNodesAddOnAddOnPlan {
	return v.AddOnPlan
}

// GetPrimaryRegion returns ListAddOnsAddOnsAddOnConnectionNodesAddOn.PrimaryRegion, and is useful for accessing the field via an interface.
func (v *ListAddOnsAddOnsAddOnConnectionNodesAddOn) GetPrimaryRegion() string { return v.PrimaryRegion }

// GetReadRegions returns ListAddOnsAddOnsAddOnConnectionNodesAddOn.ReadRegions, and is useful for accessing the field via an interface.
func (v *ListAddOnsAddOnsAddOnConnectionNodesAddOn) GetReadRegions() []string { return v.ReadRegions }

// GetOrganization returns ListAddOnsAddOnsAddOnConnectionNodesAddOn.Organization, and is useful for accessing the field via an interface.
func (v *ListAddOnsAddOnsAddOnConnectionNodesAddOn) GetOrganization() ListAddOnsAddOnsAddOnConnectionNodesAddOnOrganization {
	return v.Organization
}

// ListAddOnsAddOnsAddOnConnectionNodesAddOnAddOnPlan includes the requested fields of the GraphQL type AddOnPlan.
type ListAddOnsAddOnsAddOnConnectionNodesAddOnAddOnPlan struct {
	DisplayName string `json:"displayName"`
}

// GetDisplayName returns ListAddOnsAddOnsAddOnConnectionNodesAddOnAddOnPlan.DisplayName, and is useful for accessing the field via an interface.
func (v *ListAddOnsAddOnsAddOnConnectionNodesAddOnAddOnPlan) GetDisplayName() string {
	return v.DisplayName
}

// ListAddOnsAddOnsAddOnConnectionNodesAddOnOrganization includes the requested fields of the GraphQL type Organization.
type ListAddOnsAddOnsAddOnConnectionNodesAddOnOrganization struct {
	Id string `json:"id"`
	// Unique organization slug
	Slug string `json:"slug"`
}

// GetId returns ListAddOnsAddOnsAddOnConnectionNodesAddOnOrganization.Id, and is useful for accessing the field via an interface.
func (v *ListAddOnsAddOnsAddOnConnectionNodesAddOnOrganization) GetId() string { return v.Id }

// GetSlug returns ListAddOnsAddOnsAddOnConnectionNodesAddOnOrganization.Slug, and is useful for accessing the field via an interface.
func (v *ListAddOnsAddOnsAddOnConnectionNodesAddOnOrganization) GetSlug() string { return v.Slug }

// ListAddOnsResponse is returned by ListAddOns on success.
type ListAddOnsResponse struct {
	// List add-ons associated with an organization
	AddOns ListAddOnsAddOnsAddOnConnection `json:"addOns"`
}

// GetAddOns returns ListAddOnsResponse.AddOns, and is useful for accessing the field via an interface.
func (v *ListAddOnsResponse) GetAddOns() ListAddOnsAddOnsAddOnConnection { return v.AddOns }

// UpdateAddOnResponse is returned by UpdateAddOn on success.
type UpdateAddOnResponse struct {
	UpdateAddOn UpdateAddOnUpdateAddOnUpdateAddOnPayload `json:"updateAddOn"`
}

// GetUpdateAddOn returns UpdateAddOnResponse.UpdateAddOn, and is useful for accessing the field via an interface.
func (v *UpdateAddOnResponse) GetUpdateAddOn() UpdateAddOnUpdateAddOnUpdateAddOnPayload {
	return v.UpdateAddOn
}

// UpdateAddOnUpdateAddOnUpdateAddOnPayload includes the requested fields of the GraphQL type UpdateAddOnPayload.
// The GraphQL type's documentation follows.
//
// Autogenerated return type of UpdateAddOn
type UpdateAddOnUpdateAddOnUpdateAddOnPayload struct {
	AddOn UpdateAddOnUpdateAddOnUpdateAddOnPayloadAddOn `json:"addOn"`
}

// GetAddOn returns UpdateAddOnUpdateAddOnUpdateAddOnPayload.AddOn, and is useful for accessing the field via an interface.
func (v *UpdateAddOnUpdateAddOnUpdateAddOnPayload) GetAddOn() UpdateAddOnUpdateAddOnUpdateAddOnPayloadAddOn {
	return v.AddOn
}

// UpdateAddOnUpdateAddOnUpdateAddOnPayloadAddOn includes the requested fields of the GraphQL type AddOn.
type UpdateAddOnUpdateAddOnUpdateAddOnPayloadAddOn struct {
	Id string `json:"id"`
}

// GetId returns UpdateAddOnUpdateAddOnUpdateAddOnPayloadAddOn.Id, and is useful for accessing the field via an interface.
func (v *UpdateAddOnUpdateAddOnUpdateAddOnPayloadAddOn) GetId() string { return v.Id }

// __CreateAddOnInput is used internally by genqlient
type __CreateAddOnInput struct {
	OrganizationId string      `json:"organizationId"`
	PrimaryRegion  string      `json:"primaryRegion"`
	PlanId         string      `json:"planId"`
	ReadRegions    []string    `json:"readRegions"`
	Options        interface{} `json:"options"`
}

// GetOrganizationId returns __CreateAddOnInput.OrganizationId, and is useful for accessing the field via an interface.
func (v *__CreateAddOnInput) GetOrganizationId() string { return v.OrganizationId }

// GetPrimaryRegion returns __CreateAddOnInput.PrimaryRegion, and is useful for accessing the field via an interface.
func (v *__CreateAddOnInput) GetPrimaryRegion() string { return v.PrimaryRegion }

// GetPlanId returns __CreateAddOnInput.PlanId, and is useful for accessing the field via an interface.
func (v *__CreateAddOnInput) GetPlanId() string { return v.PlanId }

// GetReadRegions returns __CreateAddOnInput.ReadRegions, and is useful for accessing the field via an interface.
func (v *__CreateAddOnInput) GetReadRegions() []string { return v.ReadRegions }

// GetOptions returns __CreateAddOnInput.Options, and is useful for accessing the field via an interface.
func (v *__CreateAddOnInput) GetOptions() interface{} { return v.Options }

// __DeleteAddOnInput is used internally by genqlient
type __DeleteAddOnInput struct {
	AddOnId string `json:"addOnId"`
}

// GetAddOnId returns __DeleteAddOnInput.AddOnId, and is useful for accessing the field via an interface.
func (v *__DeleteAddOnInput) GetAddOnId() string { return v.AddOnId }

// __GetAddOnInput is used internally by genqlient
type __GetAddOnInput struct {
	Id string `json:"id"`
}

// GetId returns __GetAddOnInput.Id, and is useful for accessing the field via an interface.
func (v *__GetAddOnInput) GetId() string { return v.Id }

// __ListAddOnsInput is used internally by genqlient
type __ListAddOnsInput struct {
	AddOnType AddOnType `json:"addOnType"`
}

// GetAddOnType returns __ListAddOnsInput.AddOnType, and is useful for accessing the field via an interface.
func (v *__ListAddOnsInput) GetAddOnType() AddOnType { return v.AddOnType }

// __UpdateAddOnInput is used internally by genqlient
type __UpdateAddOnInput struct {
	AddOnId     string   `json:"addOnId"`
	PlanId      string   `json:"planId"`
	ReadRegions []string `json:"readRegions"`
}

// GetAddOnId returns __UpdateAddOnInput.AddOnId, and is useful for accessing the field via an interface.
func (v *__UpdateAddOnInput) GetAddOnId() string { return v.AddOnId }

// GetPlanId returns __UpdateAddOnInput.PlanId, and is useful for accessing the field via an interface.
func (v *__UpdateAddOnInput) GetPlanId() string { return v.PlanId }

// GetReadRegions returns __UpdateAddOnInput.ReadRegions, and is useful for accessing the field via an interface.
func (v *__UpdateAddOnInput) GetReadRegions() []string { return v.ReadRegions }

func CreateAddOn(
	ctx context.Context,
	client graphql.Client,
	organizationId string,
	primaryRegion string,
	planId string,
	readRegions []string,
	options interface{},
) (*CreateAddOnResponse, error) {
	req := &graphql.Request{
		OpName: "CreateAddOn",
		Query: `
mutation CreateAddOn ($organizationId: ID!, $primaryRegion: String!, $planId: ID!, $readRegions: [String!], $options: JSON!) {
	createAddOn(input: {organizationId:$organizationId,type:redis,planId:$planId,primaryRegion:$primaryRegion,readRegions:$readRegions,options:$options}) {
		addOn {
			id
			publicUrl
		}
	}
}
`,
		Variables: &__CreateAddOnInput{
			OrganizationId: organizationId,
			PrimaryRegion:  primaryRegion,
			PlanId:         planId,
			ReadRegions:    readRegions,
			Options:        options,
		},
	}
	var err error

	var data CreateAddOnResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}

func DeleteAddOn(
	ctx context.Context,
	client graphql.Client,
	addOnId string,
) (*DeleteAddOnResponse, error) {
	req := &graphql.Request{
		OpName: "DeleteAddOn",
		Query: `
mutation DeleteAddOn ($addOnId: ID!) {
	deleteAddOn(input: {addOnId:$addOnId}) {
		deletedAddOnId
	}
}
`,
		Variables: &__DeleteAddOnInput{
			AddOnId: addOnId,
		},
	}
	var err error

	var data DeleteAddOnResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}

func GetAddOn(
	ctx context.Context,
	client graphql.Client,
	id string,
) (*GetAddOnResponse, error) {
	req := &graphql.Request{
		OpName: "GetAddOn",
		Query: `
query GetAddOn ($id: ID!) {
	addOn(id: $id) {
		id
		name
		publicUrl
		primaryRegion
		readRegions
		addOnPlan {
			id
			name
			displayName
		}
	}
}
`,
		Variables: &__GetAddOnInput{
			Id: id,
		},
	}
	var err error

	var data GetAddOnResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}

func ListAddOnPlans(
	ctx context.Context,
	client graphql.Client,
) (*ListAddOnPlansResponse, error) {
	req := &graphql.Request{
		OpName: "ListAddOnPlans",
		Query: `
query ListAddOnPlans {
	addOnPlans {
		nodes {
			id
			displayName
			maxDataSize
			pricePerMonth
		}
	}
}
`,
	}
	var err error

	var data ListAddOnPlansResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}

func ListAddOns(
	ctx context.Context,
	client graphql.Client,
	addOnType AddOnType,
) (*ListAddOnsResponse, error) {
	req := &graphql.Request{
		OpName: "ListAddOns",
		Query: `
query ListAddOns ($addOnType: AddOnType) {
	addOns(type: $addOnType) {
		nodes {
			id
			name
			addOnPlan {
				displayName
			}
			primaryRegion
			readRegions
			organization {
				id
				slug
			}
		}
	}
}
`,
		Variables: &__ListAddOnsInput{
			AddOnType: addOnType,
		},
	}
	var err error

	var data ListAddOnsResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}

func UpdateAddOn(
	ctx context.Context,
	client graphql.Client,
	addOnId string,
	planId string,
	readRegions []string,
) (*UpdateAddOnResponse, error) {
	req := &graphql.Request{
		OpName: "UpdateAddOn",
		Query: `
mutation UpdateAddOn ($addOnId: ID!, $planId: ID!, $readRegions: [String!]) {
	updateAddOn(input: {addOnId:$addOnId,planId:$planId,readRegions:$readRegions}) {
		addOn {
			id
		}
	}
}
`,
		Variables: &__UpdateAddOnInput{
			AddOnId:     addOnId,
			PlanId:      planId,
			ReadRegions: readRegions,
		},
	}
	var err error

	var data UpdateAddOnResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}
