// Code generated by protoc-gen-goext. DO NOT EDIT.

package containerregistry

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
)

func (m *GetLifecyclePolicyRequest) SetLifecyclePolicyId(v string) {
	m.LifecyclePolicyId = v
}

type ListLifecyclePoliciesRequest_Id = isListLifecyclePoliciesRequest_Id

func (m *ListLifecyclePoliciesRequest) SetId(v ListLifecyclePoliciesRequest_Id) {
	m.Id = v
}

func (m *ListLifecyclePoliciesRequest) SetRegistryId(v string) {
	m.Id = &ListLifecyclePoliciesRequest_RegistryId{
		RegistryId: v,
	}
}

func (m *ListLifecyclePoliciesRequest) SetRepositoryId(v string) {
	m.Id = &ListLifecyclePoliciesRequest_RepositoryId{
		RepositoryId: v,
	}
}

func (m *ListLifecyclePoliciesRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListLifecyclePoliciesRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListLifecyclePoliciesRequest) SetFilter(v string) {
	m.Filter = v
}

func (m *ListLifecyclePoliciesRequest) SetOrderBy(v string) {
	m.OrderBy = v
}

func (m *ListLifecyclePoliciesResponse) SetLifecyclePolicies(v []*LifecyclePolicy) {
	m.LifecyclePolicies = v
}

func (m *ListLifecyclePoliciesResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *CreateLifecyclePolicyRequest) SetRepositoryId(v string) {
	m.RepositoryId = v
}

func (m *CreateLifecyclePolicyRequest) SetName(v string) {
	m.Name = v
}

func (m *CreateLifecyclePolicyRequest) SetDescription(v string) {
	m.Description = v
}

func (m *CreateLifecyclePolicyRequest) SetStatus(v LifecyclePolicy_Status) {
	m.Status = v
}

func (m *CreateLifecyclePolicyRequest) SetRules(v []*LifecycleRule) {
	m.Rules = v
}

func (m *UpdateLifecyclePolicyRequest) SetLifecyclePolicyId(v string) {
	m.LifecyclePolicyId = v
}

func (m *UpdateLifecyclePolicyRequest) SetUpdateMask(v *field_mask.FieldMask) {
	m.UpdateMask = v
}

func (m *UpdateLifecyclePolicyRequest) SetName(v string) {
	m.Name = v
}

func (m *UpdateLifecyclePolicyRequest) SetDescription(v string) {
	m.Description = v
}

func (m *UpdateLifecyclePolicyRequest) SetStatus(v LifecyclePolicy_Status) {
	m.Status = v
}

func (m *UpdateLifecyclePolicyRequest) SetRules(v []*LifecycleRule) {
	m.Rules = v
}

func (m *DeleteLifecyclePolicyRequest) SetLifecyclePolicyId(v string) {
	m.LifecyclePolicyId = v
}

func (m *CreateLifecyclePolicyMetadata) SetLifecyclePolicyId(v string) {
	m.LifecyclePolicyId = v
}

func (m *UpdateLifecyclePolicyMetadata) SetLifecyclePolicyId(v string) {
	m.LifecyclePolicyId = v
}

func (m *DeleteLifecyclePolicyMetadata) SetLifecyclePolicyId(v string) {
	m.LifecyclePolicyId = v
}

func (m *DryRunLifecyclePolicyRequest) SetLifecyclePolicyId(v string) {
	m.LifecyclePolicyId = v
}

func (m *DryRunLifecyclePolicyMetadata) SetDryRunLifecyclePolicyResultId(v string) {
	m.DryRunLifecyclePolicyResultId = v
}

func (m *DryRunLifecyclePolicyMetadata) SetLifecyclePolicyId(v string) {
	m.LifecyclePolicyId = v
}

func (m *DryRunLifecyclePolicyResult) SetDryRunLifecyclePolicyResultId(v string) {
	m.DryRunLifecyclePolicyResultId = v
}

func (m *DryRunLifecyclePolicyResult) SetLifecyclePolicyId(v string) {
	m.LifecyclePolicyId = v
}

func (m *DryRunLifecyclePolicyResult) SetRunAt(v *timestamp.Timestamp) {
	m.RunAt = v
}

func (m *DryRunLifecyclePolicyResult) SetAffectedImagesCount(v int64) {
	m.AffectedImagesCount = v
}

func (m *ListDryRunLifecyclePolicyResultsRequest) SetLifecyclePolicyId(v string) {
	m.LifecyclePolicyId = v
}

func (m *ListDryRunLifecyclePolicyResultsRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListDryRunLifecyclePolicyResultsRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListDryRunLifecyclePolicyResultsRequest) SetFilter(v string) {
	m.Filter = v
}

func (m *ListDryRunLifecyclePolicyResultsRequest) SetOrderBy(v string) {
	m.OrderBy = v
}

func (m *ListRunLifecyclePolicyResultsResponse) SetDryRunLifecyclePolicyResults(v []*DryRunLifecyclePolicyResult) {
	m.DryRunLifecyclePolicyResults = v
}

func (m *ListRunLifecyclePolicyResultsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *ListDryRunLifecyclePolicyResultAffectedImagesRequest) SetDryRunLifecyclePolicyResultId(v string) {
	m.DryRunLifecyclePolicyResultId = v
}

func (m *ListDryRunLifecyclePolicyResultAffectedImagesRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListDryRunLifecyclePolicyResultAffectedImagesRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListDryRunLifecyclePolicyResultAffectedImagesRequest) SetFilter(v string) {
	m.Filter = v
}

func (m *ListDryRunLifecyclePolicyResultAffectedImagesRequest) SetOrderBy(v string) {
	m.OrderBy = v
}

func (m *ListDryRunLifecyclePolicyResultAffectedImagesResponse) SetAffectedImages(v []*Image) {
	m.AffectedImages = v
}

func (m *ListDryRunLifecyclePolicyResultAffectedImagesResponse) SetDryRunLifecyclePolicyResultId(v string) {
	m.DryRunLifecyclePolicyResultId = v
}

func (m *ListDryRunLifecyclePolicyResultAffectedImagesResponse) SetLifecyclePolicyId(v string) {
	m.LifecyclePolicyId = v
}

func (m *ListDryRunLifecyclePolicyResultAffectedImagesResponse) SetRunAt(v *timestamp.Timestamp) {
	m.RunAt = v
}

func (m *ListDryRunLifecyclePolicyResultAffectedImagesResponse) SetAffectedImagesCount(v int64) {
	m.AffectedImagesCount = v
}

func (m *ListDryRunLifecyclePolicyResultAffectedImagesResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}
