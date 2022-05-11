// Code generated by protoc-gen-goext. DO NOT EDIT.

package greenplum

import (
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

func (m *Resources) SetResourcePresetId(v string) {
	m.ResourcePresetId = v
}

func (m *Resources) SetDiskSize(v int64) {
	m.DiskSize = v
}

func (m *Resources) SetDiskTypeId(v string) {
	m.DiskTypeId = v
}

func (m *ConnectionPoolerConfig) SetMode(v ConnectionPoolerConfig_PoolMode) {
	m.Mode = v
}

func (m *ConnectionPoolerConfig) SetSize(v *wrapperspb.Int64Value) {
	m.Size = v
}

func (m *ConnectionPoolerConfig) SetClientIdleTimeout(v *wrapperspb.Int64Value) {
	m.ClientIdleTimeout = v
}

func (m *MasterSubclusterConfig) SetResources(v *Resources) {
	m.Resources = v
}

func (m *SegmentSubclusterConfig) SetResources(v *Resources) {
	m.Resources = v
}

func (m *GreenplumConfig6_17) SetMaxConnections(v *wrapperspb.Int64Value) {
	m.MaxConnections = v
}

func (m *GreenplumConfig6_17) SetMaxSlotWalKeepSize(v *wrapperspb.Int64Value) {
	m.MaxSlotWalKeepSize = v
}

func (m *GreenplumConfig6_17) SetGpWorkfileLimitPerSegment(v *wrapperspb.Int64Value) {
	m.GpWorkfileLimitPerSegment = v
}

func (m *GreenplumConfig6_17) SetGpWorkfileLimitPerQuery(v *wrapperspb.Int64Value) {
	m.GpWorkfileLimitPerQuery = v
}

func (m *GreenplumConfig6_17) SetGpWorkfileLimitFilesPerQuery(v *wrapperspb.Int64Value) {
	m.GpWorkfileLimitFilesPerQuery = v
}

func (m *GreenplumConfig6_17) SetMaxPreparedTransactions(v *wrapperspb.Int64Value) {
	m.MaxPreparedTransactions = v
}

func (m *GreenplumConfig6_17) SetGpWorkfileCompression(v *wrapperspb.BoolValue) {
	m.GpWorkfileCompression = v
}

func (m *GreenplumConfig6_19) SetMaxConnections(v *wrapperspb.Int64Value) {
	m.MaxConnections = v
}

func (m *GreenplumConfig6_19) SetMaxSlotWalKeepSize(v *wrapperspb.Int64Value) {
	m.MaxSlotWalKeepSize = v
}

func (m *GreenplumConfig6_19) SetGpWorkfileLimitPerSegment(v *wrapperspb.Int64Value) {
	m.GpWorkfileLimitPerSegment = v
}

func (m *GreenplumConfig6_19) SetGpWorkfileLimitPerQuery(v *wrapperspb.Int64Value) {
	m.GpWorkfileLimitPerQuery = v
}

func (m *GreenplumConfig6_19) SetGpWorkfileLimitFilesPerQuery(v *wrapperspb.Int64Value) {
	m.GpWorkfileLimitFilesPerQuery = v
}

func (m *GreenplumConfig6_19) SetMaxPreparedTransactions(v *wrapperspb.Int64Value) {
	m.MaxPreparedTransactions = v
}

func (m *GreenplumConfig6_19) SetGpWorkfileCompression(v *wrapperspb.BoolValue) {
	m.GpWorkfileCompression = v
}

func (m *GreenplumConfig6_19) SetMaxStatementMem(v *wrapperspb.Int64Value) {
	m.MaxStatementMem = v
}

func (m *GreenplumConfig6_19) SetLogStatement(v LogStatement) {
	m.LogStatement = v
}

func (m *GreenplumConfigSet6_17) SetEffectiveConfig(v *GreenplumConfig6_17) {
	m.EffectiveConfig = v
}

func (m *GreenplumConfigSet6_17) SetUserConfig(v *GreenplumConfig6_17) {
	m.UserConfig = v
}

func (m *GreenplumConfigSet6_17) SetDefaultConfig(v *GreenplumConfig6_17) {
	m.DefaultConfig = v
}

func (m *GreenplumConfigSet6_19) SetEffectiveConfig(v *GreenplumConfig6_19) {
	m.EffectiveConfig = v
}

func (m *GreenplumConfigSet6_19) SetUserConfig(v *GreenplumConfig6_19) {
	m.UserConfig = v
}

func (m *GreenplumConfigSet6_19) SetDefaultConfig(v *GreenplumConfig6_19) {
	m.DefaultConfig = v
}

func (m *ConnectionPoolerConfigSet) SetEffectiveConfig(v *ConnectionPoolerConfig) {
	m.EffectiveConfig = v
}

func (m *ConnectionPoolerConfigSet) SetUserConfig(v *ConnectionPoolerConfig) {
	m.UserConfig = v
}

func (m *ConnectionPoolerConfigSet) SetDefaultConfig(v *ConnectionPoolerConfig) {
	m.DefaultConfig = v
}
