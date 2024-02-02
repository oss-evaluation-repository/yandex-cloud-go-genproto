// Code generated by protoc-gen-goext. DO NOT EDIT.

package endpoint

import (
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func (m *ClickhouseShard) SetName(v string) {
	m.Name = v
}

func (m *ClickhouseShard) SetHosts(v []string) {
	m.Hosts = v
}

func (m *OnPremiseClickhouse) SetShards(v []*ClickhouseShard) {
	m.Shards = v
}

func (m *OnPremiseClickhouse) SetHttpPort(v int64) {
	m.HttpPort = v
}

func (m *OnPremiseClickhouse) SetNativePort(v int64) {
	m.NativePort = v
}

func (m *OnPremiseClickhouse) SetTlsMode(v *TLSMode) {
	m.TlsMode = v
}

type ClickhouseConnectionOptions_Address = isClickhouseConnectionOptions_Address

func (m *ClickhouseConnectionOptions) SetAddress(v ClickhouseConnectionOptions_Address) {
	m.Address = v
}

func (m *ClickhouseConnectionOptions) SetMdbClusterId(v string) {
	m.Address = &ClickhouseConnectionOptions_MdbClusterId{
		MdbClusterId: v,
	}
}

func (m *ClickhouseConnectionOptions) SetOnPremise(v *OnPremiseClickhouse) {
	m.Address = &ClickhouseConnectionOptions_OnPremise{
		OnPremise: v,
	}
}

func (m *ClickhouseConnectionOptions) SetUser(v string) {
	m.User = v
}

func (m *ClickhouseConnectionOptions) SetPassword(v *Secret) {
	m.Password = v
}

func (m *ClickhouseConnectionOptions) SetDatabase(v string) {
	m.Database = v
}

type ClickhouseConnection_Connection = isClickhouseConnection_Connection

func (m *ClickhouseConnection) SetConnection(v ClickhouseConnection_Connection) {
	m.Connection = v
}

func (m *ClickhouseConnection) SetConnectionOptions(v *ClickhouseConnectionOptions) {
	m.Connection = &ClickhouseConnection_ConnectionOptions{
		ConnectionOptions: v,
	}
}

type ClickhouseSharding_Sharding = isClickhouseSharding_Sharding

func (m *ClickhouseSharding) SetSharding(v ClickhouseSharding_Sharding) {
	m.Sharding = v
}

func (m *ClickhouseSharding) SetColumnValueHash(v *ClickhouseSharding_ColumnValueHash) {
	m.Sharding = &ClickhouseSharding_ColumnValueHash_{
		ColumnValueHash: v,
	}
}

func (m *ClickhouseSharding) SetCustomMapping(v *ClickhouseSharding_ColumnValueMapping) {
	m.Sharding = &ClickhouseSharding_CustomMapping{
		CustomMapping: v,
	}
}

func (m *ClickhouseSharding) SetTransferId(v *emptypb.Empty) {
	m.Sharding = &ClickhouseSharding_TransferId{
		TransferId: v,
	}
}

func (m *ClickhouseSharding) SetRoundRobin(v *emptypb.Empty) {
	m.Sharding = &ClickhouseSharding_RoundRobin{
		RoundRobin: v,
	}
}

func (m *ClickhouseSharding_ColumnValueHash) SetColumnName(v string) {
	m.ColumnName = v
}

func (m *ClickhouseSharding_ColumnValueMapping) SetColumnName(v string) {
	m.ColumnName = v
}

func (m *ClickhouseSharding_ColumnValueMapping) SetMapping(v []*ClickhouseSharding_ColumnValueMapping_ValueToShard) {
	m.Mapping = v
}

func (m *ClickhouseSharding_ColumnValueMapping_ValueToShard) SetColumnValue(v *ColumnValue) {
	m.ColumnValue = v
}

func (m *ClickhouseSharding_ColumnValueMapping_ValueToShard) SetShardName(v string) {
	m.ShardName = v
}

func (m *ClickhouseSource) SetConnection(v *ClickhouseConnection) {
	m.Connection = v
}

func (m *ClickhouseSource) SetIncludeTables(v []string) {
	m.IncludeTables = v
}

func (m *ClickhouseSource) SetExcludeTables(v []string) {
	m.ExcludeTables = v
}

func (m *ClickhouseSource) SetSubnetId(v string) {
	m.SubnetId = v
}

func (m *ClickhouseSource) SetSecurityGroups(v []string) {
	m.SecurityGroups = v
}

func (m *ClickhouseTarget) SetConnection(v *ClickhouseConnection) {
	m.Connection = v
}

func (m *ClickhouseTarget) SetSubnetId(v string) {
	m.SubnetId = v
}

func (m *ClickhouseTarget) SetAltNames(v []*AltName) {
	m.AltNames = v
}

func (m *ClickhouseTarget) SetCleanupPolicy(v ClickhouseCleanupPolicy) {
	m.CleanupPolicy = v
}

func (m *ClickhouseTarget) SetSharding(v *ClickhouseSharding) {
	m.Sharding = v
}

func (m *ClickhouseTarget) SetClickhouseClusterName(v string) {
	m.ClickhouseClusterName = v
}

func (m *ClickhouseTarget) SetSecurityGroups(v []string) {
	m.SecurityGroups = v
}
