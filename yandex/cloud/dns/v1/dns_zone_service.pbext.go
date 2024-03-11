// Code generated by protoc-gen-goext. DO NOT EDIT.

package dns

import (
	operation "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
)

func (m *UpdateDnsZonePrivateNetworksRequest) SetDnsZoneId(v string) {
	m.DnsZoneId = v
}

func (m *UpdateDnsZonePrivateNetworksRequest) SetPrivateNetworkIdAdditions(v []string) {
	m.PrivateNetworkIdAdditions = v
}

func (m *UpdateDnsZonePrivateNetworksRequest) SetPrivateNetworkIdDeletions(v []string) {
	m.PrivateNetworkIdDeletions = v
}

func (m *UpdateDnsZonePrivateNetworksMetadata) SetDnsZoneId(v string) {
	m.DnsZoneId = v
}

func (m *GetDnsZoneRequest) SetDnsZoneId(v string) {
	m.DnsZoneId = v
}

func (m *ListDnsZonesRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *ListDnsZonesRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListDnsZonesRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListDnsZonesRequest) SetFilter(v string) {
	m.Filter = v
}

func (m *ListDnsZonesResponse) SetDnsZones(v []*DnsZone) {
	m.DnsZones = v
}

func (m *ListDnsZonesResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *CreateDnsZoneRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *CreateDnsZoneRequest) SetName(v string) {
	m.Name = v
}

func (m *CreateDnsZoneRequest) SetDescription(v string) {
	m.Description = v
}

func (m *CreateDnsZoneRequest) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *CreateDnsZoneRequest) SetZone(v string) {
	m.Zone = v
}

func (m *CreateDnsZoneRequest) SetPrivateVisibility(v *PrivateVisibility) {
	m.PrivateVisibility = v
}

func (m *CreateDnsZoneRequest) SetPublicVisibility(v *PublicVisibility) {
	m.PublicVisibility = v
}

func (m *CreateDnsZoneRequest) SetDeletionProtection(v bool) {
	m.DeletionProtection = v
}

func (m *CreateDnsZoneMetadata) SetDnsZoneId(v string) {
	m.DnsZoneId = v
}

func (m *UpdateDnsZoneRequest) SetDnsZoneId(v string) {
	m.DnsZoneId = v
}

func (m *UpdateDnsZoneRequest) SetUpdateMask(v *fieldmaskpb.FieldMask) {
	m.UpdateMask = v
}

func (m *UpdateDnsZoneRequest) SetName(v string) {
	m.Name = v
}

func (m *UpdateDnsZoneRequest) SetDescription(v string) {
	m.Description = v
}

func (m *UpdateDnsZoneRequest) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *UpdateDnsZoneRequest) SetPrivateVisibility(v *PrivateVisibility) {
	m.PrivateVisibility = v
}

func (m *UpdateDnsZoneRequest) SetPublicVisibility(v *PublicVisibility) {
	m.PublicVisibility = v
}

func (m *UpdateDnsZoneRequest) SetDeletionProtection(v bool) {
	m.DeletionProtection = v
}

func (m *UpdateDnsZoneMetadata) SetDnsZoneId(v string) {
	m.DnsZoneId = v
}

func (m *DeleteDnsZoneRequest) SetDnsZoneId(v string) {
	m.DnsZoneId = v
}

func (m *DeleteDnsZoneMetadata) SetDnsZoneId(v string) {
	m.DnsZoneId = v
}

func (m *GetDnsZoneRecordSetRequest) SetDnsZoneId(v string) {
	m.DnsZoneId = v
}

func (m *GetDnsZoneRecordSetRequest) SetName(v string) {
	m.Name = v
}

func (m *GetDnsZoneRecordSetRequest) SetType(v string) {
	m.Type = v
}

func (m *ListDnsZoneRecordSetsRequest) SetDnsZoneId(v string) {
	m.DnsZoneId = v
}

func (m *ListDnsZoneRecordSetsRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListDnsZoneRecordSetsRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListDnsZoneRecordSetsRequest) SetFilter(v string) {
	m.Filter = v
}

func (m *ListDnsZoneRecordSetsResponse) SetRecordSets(v []*RecordSet) {
	m.RecordSets = v
}

func (m *ListDnsZoneRecordSetsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *UpdateRecordSetsRequest) SetDnsZoneId(v string) {
	m.DnsZoneId = v
}

func (m *UpdateRecordSetsRequest) SetDeletions(v []*RecordSet) {
	m.Deletions = v
}

func (m *UpdateRecordSetsRequest) SetAdditions(v []*RecordSet) {
	m.Additions = v
}

func (m *UpsertRecordSetsRequest) SetDnsZoneId(v string) {
	m.DnsZoneId = v
}

func (m *UpsertRecordSetsRequest) SetDeletions(v []*RecordSet) {
	m.Deletions = v
}

func (m *UpsertRecordSetsRequest) SetReplacements(v []*RecordSet) {
	m.Replacements = v
}

func (m *UpsertRecordSetsRequest) SetMerges(v []*RecordSet) {
	m.Merges = v
}

func (m *RecordSetDiff) SetAdditions(v []*RecordSet) {
	m.Additions = v
}

func (m *RecordSetDiff) SetDeletions(v []*RecordSet) {
	m.Deletions = v
}

func (m *ListDnsZoneOperationsRequest) SetDnsZoneId(v string) {
	m.DnsZoneId = v
}

func (m *ListDnsZoneOperationsRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListDnsZoneOperationsRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListDnsZoneOperationsRequest) SetFilter(v string) {
	m.Filter = v
}

func (m *ListDnsZoneOperationsResponse) SetOperations(v []*operation.Operation) {
	m.Operations = v
}

func (m *ListDnsZoneOperationsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}
