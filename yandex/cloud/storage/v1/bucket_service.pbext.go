// Code generated by protoc-gen-goext. DO NOT EDIT.

package storage

import (
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	structpb "google.golang.org/protobuf/types/known/structpb"
)

func (m *GetBucketRequest) SetName(v string) {
	m.Name = v
}

func (m *GetBucketRequest) SetView(v GetBucketRequest_View) {
	m.View = v
}

func (m *ListBucketsRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *ListBucketsResponse) SetBuckets(v []*Bucket) {
	m.Buckets = v
}

func (m *CreateBucketRequest) SetName(v string) {
	m.Name = v
}

func (m *CreateBucketRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *CreateBucketRequest) SetDefaultStorageClass(v string) {
	m.DefaultStorageClass = v
}

func (m *CreateBucketRequest) SetMaxSize(v int64) {
	m.MaxSize = v
}

func (m *CreateBucketRequest) SetAnonymousAccessFlags(v *AnonymousAccessFlags) {
	m.AnonymousAccessFlags = v
}

func (m *CreateBucketRequest) SetAcl(v *ACL) {
	m.Acl = v
}

func (m *CreateBucketRequest) SetTags(v []*Tag) {
	m.Tags = v
}

func (m *CreateBucketMetadata) SetName(v string) {
	m.Name = v
}

func (m *UpdateBucketRequest) SetName(v string) {
	m.Name = v
}

func (m *UpdateBucketRequest) SetFieldMask(v *fieldmaskpb.FieldMask) {
	m.FieldMask = v
}

func (m *UpdateBucketRequest) SetAnonymousAccessFlags(v *AnonymousAccessFlags) {
	m.AnonymousAccessFlags = v
}

func (m *UpdateBucketRequest) SetDefaultStorageClass(v string) {
	m.DefaultStorageClass = v
}

func (m *UpdateBucketRequest) SetMaxSize(v int64) {
	m.MaxSize = v
}

func (m *UpdateBucketRequest) SetCors(v []*CorsRule) {
	m.Cors = v
}

func (m *UpdateBucketRequest) SetWebsiteSettings(v *WebsiteSettings) {
	m.WebsiteSettings = v
}

func (m *UpdateBucketRequest) SetVersioning(v Versioning) {
	m.Versioning = v
}

func (m *UpdateBucketRequest) SetLifecycleRules(v []*LifecycleRule) {
	m.LifecycleRules = v
}

func (m *UpdateBucketRequest) SetPolicy(v *structpb.Struct) {
	m.Policy = v
}

func (m *UpdateBucketRequest) SetAcl(v *ACL) {
	m.Acl = v
}

func (m *UpdateBucketRequest) SetTags(v []*Tag) {
	m.Tags = v
}

func (m *UpdateBucketMetadata) SetName(v string) {
	m.Name = v
}

func (m *DeleteBucketRequest) SetName(v string) {
	m.Name = v
}

func (m *DeleteBucketMetadata) SetName(v string) {
	m.Name = v
}

func (m *GetBucketStatsRequest) SetName(v string) {
	m.Name = v
}

func (m *GetBucketHTTPSConfigRequest) SetName(v string) {
	m.Name = v
}

func (m *SelfManagedHTTPSConfigParams) SetCertificatePem(v string) {
	m.CertificatePem = v
}

func (m *SelfManagedHTTPSConfigParams) SetPrivateKeyPem(v string) {
	m.PrivateKeyPem = v
}

func (m *CertificateManagerHTTPSConfigParams) SetCertificateId(v string) {
	m.CertificateId = v
}

type SetBucketHTTPSConfigRequest_Params = isSetBucketHTTPSConfigRequest_Params

func (m *SetBucketHTTPSConfigRequest) SetParams(v SetBucketHTTPSConfigRequest_Params) {
	m.Params = v
}

func (m *SetBucketHTTPSConfigRequest) SetName(v string) {
	m.Name = v
}

func (m *SetBucketHTTPSConfigRequest) SetSelfManaged(v *SelfManagedHTTPSConfigParams) {
	m.Params = &SetBucketHTTPSConfigRequest_SelfManaged{
		SelfManaged: v,
	}
}

func (m *SetBucketHTTPSConfigRequest) SetCertificateManager(v *CertificateManagerHTTPSConfigParams) {
	m.Params = &SetBucketHTTPSConfigRequest_CertificateManager{
		CertificateManager: v,
	}
}

func (m *SetBucketHTTPSConfigMetadata) SetName(v string) {
	m.Name = v
}

func (m *DeleteBucketHTTPSConfigRequest) SetName(v string) {
	m.Name = v
}

func (m *DeleteBucketHTTPSConfigMetadata) SetName(v string) {
	m.Name = v
}
