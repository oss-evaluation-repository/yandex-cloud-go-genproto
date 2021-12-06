// Code generated by protoc-gen-goext. DO NOT EDIT.

package clickhouse

import (
	config "github.com/yandex-cloud/go-genproto/yandex/cloud/mdb/clickhouse/v1/config"
	operation "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	timeofday "google.golang.org/genproto/googleapis/type/timeofday"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

func (m *GetClusterRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *ListClustersRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *ListClustersRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListClustersRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListClustersRequest) SetFilter(v string) {
	m.Filter = v
}

func (m *ListClustersResponse) SetClusters(v []*Cluster) {
	m.Clusters = v
}

func (m *ListClustersResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *CreateClusterRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *CreateClusterRequest) SetName(v string) {
	m.Name = v
}

func (m *CreateClusterRequest) SetDescription(v string) {
	m.Description = v
}

func (m *CreateClusterRequest) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *CreateClusterRequest) SetEnvironment(v Cluster_Environment) {
	m.Environment = v
}

func (m *CreateClusterRequest) SetConfigSpec(v *ConfigSpec) {
	m.ConfigSpec = v
}

func (m *CreateClusterRequest) SetDatabaseSpecs(v []*DatabaseSpec) {
	m.DatabaseSpecs = v
}

func (m *CreateClusterRequest) SetUserSpecs(v []*UserSpec) {
	m.UserSpecs = v
}

func (m *CreateClusterRequest) SetHostSpecs(v []*HostSpec) {
	m.HostSpecs = v
}

func (m *CreateClusterRequest) SetNetworkId(v string) {
	m.NetworkId = v
}

func (m *CreateClusterRequest) SetShardName(v string) {
	m.ShardName = v
}

func (m *CreateClusterRequest) SetServiceAccountId(v string) {
	m.ServiceAccountId = v
}

func (m *CreateClusterRequest) SetSecurityGroupIds(v []string) {
	m.SecurityGroupIds = v
}

func (m *CreateClusterRequest) SetDeletionProtection(v bool) {
	m.DeletionProtection = v
}

func (m *CreateClusterMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *UpdateClusterRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *UpdateClusterRequest) SetUpdateMask(v *fieldmaskpb.FieldMask) {
	m.UpdateMask = v
}

func (m *UpdateClusterRequest) SetDescription(v string) {
	m.Description = v
}

func (m *UpdateClusterRequest) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *UpdateClusterRequest) SetConfigSpec(v *ConfigSpec) {
	m.ConfigSpec = v
}

func (m *UpdateClusterRequest) SetName(v string) {
	m.Name = v
}

func (m *UpdateClusterRequest) SetServiceAccountId(v string) {
	m.ServiceAccountId = v
}

func (m *UpdateClusterRequest) SetMaintenanceWindow(v *MaintenanceWindow) {
	m.MaintenanceWindow = v
}

func (m *UpdateClusterRequest) SetSecurityGroupIds(v []string) {
	m.SecurityGroupIds = v
}

func (m *UpdateClusterRequest) SetDeletionProtection(v bool) {
	m.DeletionProtection = v
}

func (m *UpdateClusterMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *DeleteClusterRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *DeleteClusterMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *StartClusterRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *StartClusterMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *StopClusterRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *StopClusterMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *MoveClusterRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *MoveClusterRequest) SetDestinationFolderId(v string) {
	m.DestinationFolderId = v
}

func (m *MoveClusterMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *MoveClusterMetadata) SetSourceFolderId(v string) {
	m.SourceFolderId = v
}

func (m *MoveClusterMetadata) SetDestinationFolderId(v string) {
	m.DestinationFolderId = v
}

func (m *AddClusterZookeeperRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *AddClusterZookeeperRequest) SetResources(v *Resources) {
	m.Resources = v
}

func (m *AddClusterZookeeperRequest) SetHostSpecs(v []*HostSpec) {
	m.HostSpecs = v
}

func (m *AddClusterZookeeperMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *BackupClusterRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *BackupClusterMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *RestoreClusterRequest) SetBackupId(v string) {
	m.BackupId = v
}

func (m *RestoreClusterRequest) SetAdditionalBackupIds(v []string) {
	m.AdditionalBackupIds = v
}

func (m *RestoreClusterRequest) SetName(v string) {
	m.Name = v
}

func (m *RestoreClusterRequest) SetDescription(v string) {
	m.Description = v
}

func (m *RestoreClusterRequest) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *RestoreClusterRequest) SetEnvironment(v Cluster_Environment) {
	m.Environment = v
}

func (m *RestoreClusterRequest) SetConfigSpec(v *ConfigSpec) {
	m.ConfigSpec = v
}

func (m *RestoreClusterRequest) SetHostSpecs(v []*HostSpec) {
	m.HostSpecs = v
}

func (m *RestoreClusterRequest) SetNetworkId(v string) {
	m.NetworkId = v
}

func (m *RestoreClusterRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *RestoreClusterRequest) SetServiceAccountId(v string) {
	m.ServiceAccountId = v
}

func (m *RestoreClusterRequest) SetSecurityGroupIds(v []string) {
	m.SecurityGroupIds = v
}

func (m *RestoreClusterMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *RestoreClusterMetadata) SetBackupId(v string) {
	m.BackupId = v
}

func (m *RescheduleMaintenanceRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *RescheduleMaintenanceRequest) SetRescheduleType(v RescheduleMaintenanceRequest_RescheduleType) {
	m.RescheduleType = v
}

func (m *RescheduleMaintenanceRequest) SetDelayedUntil(v *timestamppb.Timestamp) {
	m.DelayedUntil = v
}

func (m *RescheduleMaintenanceMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *RescheduleMaintenanceMetadata) SetDelayedUntil(v *timestamppb.Timestamp) {
	m.DelayedUntil = v
}

func (m *LogRecord) SetTimestamp(v *timestamppb.Timestamp) {
	m.Timestamp = v
}

func (m *LogRecord) SetMessage(v map[string]string) {
	m.Message = v
}

func (m *ListClusterLogsRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *ListClusterLogsRequest) SetColumnFilter(v []string) {
	m.ColumnFilter = v
}

func (m *ListClusterLogsRequest) SetServiceType(v ListClusterLogsRequest_ServiceType) {
	m.ServiceType = v
}

func (m *ListClusterLogsRequest) SetFromTime(v *timestamppb.Timestamp) {
	m.FromTime = v
}

func (m *ListClusterLogsRequest) SetToTime(v *timestamppb.Timestamp) {
	m.ToTime = v
}

func (m *ListClusterLogsRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListClusterLogsRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListClusterLogsResponse) SetLogs(v []*LogRecord) {
	m.Logs = v
}

func (m *ListClusterLogsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *StreamLogRecord) SetRecord(v *LogRecord) {
	m.Record = v
}

func (m *StreamLogRecord) SetNextRecordToken(v string) {
	m.NextRecordToken = v
}

func (m *StreamClusterLogsRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *StreamClusterLogsRequest) SetColumnFilter(v []string) {
	m.ColumnFilter = v
}

func (m *StreamClusterLogsRequest) SetServiceType(v StreamClusterLogsRequest_ServiceType) {
	m.ServiceType = v
}

func (m *StreamClusterLogsRequest) SetFromTime(v *timestamppb.Timestamp) {
	m.FromTime = v
}

func (m *StreamClusterLogsRequest) SetToTime(v *timestamppb.Timestamp) {
	m.ToTime = v
}

func (m *StreamClusterLogsRequest) SetRecordToken(v string) {
	m.RecordToken = v
}

func (m *StreamClusterLogsRequest) SetFilter(v string) {
	m.Filter = v
}

func (m *ListClusterOperationsRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *ListClusterOperationsRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListClusterOperationsRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListClusterOperationsResponse) SetOperations(v []*operation.Operation) {
	m.Operations = v
}

func (m *ListClusterOperationsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *ListClusterBackupsRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *ListClusterBackupsRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListClusterBackupsRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListClusterBackupsResponse) SetBackups(v []*Backup) {
	m.Backups = v
}

func (m *ListClusterBackupsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *ListClusterHostsRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *ListClusterHostsRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListClusterHostsRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListClusterHostsResponse) SetHosts(v []*Host) {
	m.Hosts = v
}

func (m *ListClusterHostsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *AddClusterHostsRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *AddClusterHostsRequest) SetHostSpecs(v []*HostSpec) {
	m.HostSpecs = v
}

func (m *AddClusterHostsRequest) SetCopySchema(v *wrapperspb.BoolValue) {
	m.CopySchema = v
}

func (m *AddClusterHostsMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *AddClusterHostsMetadata) SetHostNames(v []string) {
	m.HostNames = v
}

func (m *DeleteClusterHostsRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *DeleteClusterHostsRequest) SetHostNames(v []string) {
	m.HostNames = v
}

func (m *DeleteClusterHostsMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *DeleteClusterHostsMetadata) SetHostNames(v []string) {
	m.HostNames = v
}

func (m *GetClusterShardRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *GetClusterShardRequest) SetShardName(v string) {
	m.ShardName = v
}

func (m *ListClusterShardsRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *ListClusterShardsRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListClusterShardsRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListClusterShardsResponse) SetShards(v []*Shard) {
	m.Shards = v
}

func (m *ListClusterShardsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *AddClusterShardRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *AddClusterShardRequest) SetShardName(v string) {
	m.ShardName = v
}

func (m *AddClusterShardRequest) SetConfigSpec(v *ShardConfigSpec) {
	m.ConfigSpec = v
}

func (m *AddClusterShardRequest) SetHostSpecs(v []*HostSpec) {
	m.HostSpecs = v
}

func (m *AddClusterShardRequest) SetCopySchema(v *wrapperspb.BoolValue) {
	m.CopySchema = v
}

func (m *AddClusterShardMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *AddClusterShardMetadata) SetShardName(v string) {
	m.ShardName = v
}

func (m *UpdateClusterShardRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *UpdateClusterShardRequest) SetShardName(v string) {
	m.ShardName = v
}

func (m *UpdateClusterShardRequest) SetUpdateMask(v *fieldmaskpb.FieldMask) {
	m.UpdateMask = v
}

func (m *UpdateClusterShardRequest) SetConfigSpec(v *ShardConfigSpec) {
	m.ConfigSpec = v
}

func (m *UpdateClusterShardMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *UpdateClusterShardMetadata) SetShardName(v string) {
	m.ShardName = v
}

func (m *DeleteClusterShardRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *DeleteClusterShardRequest) SetShardName(v string) {
	m.ShardName = v
}

func (m *DeleteClusterShardMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *DeleteClusterShardMetadata) SetShardName(v string) {
	m.ShardName = v
}

func (m *GetClusterShardGroupRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *GetClusterShardGroupRequest) SetShardGroupName(v string) {
	m.ShardGroupName = v
}

func (m *ListClusterShardGroupsRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *ListClusterShardGroupsRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListClusterShardGroupsRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListClusterShardGroupsResponse) SetShardGroups(v []*ShardGroup) {
	m.ShardGroups = v
}

func (m *ListClusterShardGroupsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *CreateClusterShardGroupRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *CreateClusterShardGroupRequest) SetShardGroupName(v string) {
	m.ShardGroupName = v
}

func (m *CreateClusterShardGroupRequest) SetDescription(v string) {
	m.Description = v
}

func (m *CreateClusterShardGroupRequest) SetShardNames(v []string) {
	m.ShardNames = v
}

func (m *CreateClusterShardGroupMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *CreateClusterShardGroupMetadata) SetShardGroupName(v string) {
	m.ShardGroupName = v
}

func (m *UpdateClusterShardGroupRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *UpdateClusterShardGroupRequest) SetShardGroupName(v string) {
	m.ShardGroupName = v
}

func (m *UpdateClusterShardGroupRequest) SetUpdateMask(v *fieldmaskpb.FieldMask) {
	m.UpdateMask = v
}

func (m *UpdateClusterShardGroupRequest) SetDescription(v string) {
	m.Description = v
}

func (m *UpdateClusterShardGroupRequest) SetShardNames(v []string) {
	m.ShardNames = v
}

func (m *UpdateClusterShardGroupMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *UpdateClusterShardGroupMetadata) SetShardGroupName(v string) {
	m.ShardGroupName = v
}

func (m *DeleteClusterShardGroupRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *DeleteClusterShardGroupRequest) SetShardGroupName(v string) {
	m.ShardGroupName = v
}

func (m *DeleteClusterShardGroupMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *DeleteClusterShardGroupMetadata) SetShardGroupName(v string) {
	m.ShardGroupName = v
}

func (m *CreateClusterExternalDictionaryRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *CreateClusterExternalDictionaryRequest) SetExternalDictionary(v *config.ClickhouseConfig_ExternalDictionary) {
	m.ExternalDictionary = v
}

func (m *CreateClusterExternalDictionaryMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *DeleteClusterExternalDictionaryRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *DeleteClusterExternalDictionaryRequest) SetExternalDictionaryName(v string) {
	m.ExternalDictionaryName = v
}

func (m *DeleteClusterExternalDictionaryMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *HostSpec) SetZoneId(v string) {
	m.ZoneId = v
}

func (m *HostSpec) SetType(v Host_Type) {
	m.Type = v
}

func (m *HostSpec) SetSubnetId(v string) {
	m.SubnetId = v
}

func (m *HostSpec) SetAssignPublicIp(v bool) {
	m.AssignPublicIp = v
}

func (m *HostSpec) SetShardName(v string) {
	m.ShardName = v
}

func (m *ConfigSpec) SetVersion(v string) {
	m.Version = v
}

func (m *ConfigSpec) SetClickhouse(v *ConfigSpec_Clickhouse) {
	m.Clickhouse = v
}

func (m *ConfigSpec) SetZookeeper(v *ConfigSpec_Zookeeper) {
	m.Zookeeper = v
}

func (m *ConfigSpec) SetBackupWindowStart(v *timeofday.TimeOfDay) {
	m.BackupWindowStart = v
}

func (m *ConfigSpec) SetAccess(v *Access) {
	m.Access = v
}

func (m *ConfigSpec) SetCloudStorage(v *CloudStorage) {
	m.CloudStorage = v
}

func (m *ConfigSpec) SetSqlDatabaseManagement(v *wrapperspb.BoolValue) {
	m.SqlDatabaseManagement = v
}

func (m *ConfigSpec) SetSqlUserManagement(v *wrapperspb.BoolValue) {
	m.SqlUserManagement = v
}

func (m *ConfigSpec) SetAdminPassword(v string) {
	m.AdminPassword = v
}

func (m *ConfigSpec) SetEmbeddedKeeper(v *wrapperspb.BoolValue) {
	m.EmbeddedKeeper = v
}

func (m *ConfigSpec_Clickhouse) SetConfig(v *config.ClickhouseConfig) {
	m.Config = v
}

func (m *ConfigSpec_Clickhouse) SetResources(v *Resources) {
	m.Resources = v
}

func (m *ConfigSpec_Zookeeper) SetResources(v *Resources) {
	m.Resources = v
}

func (m *ShardConfigSpec) SetClickhouse(v *ShardConfigSpec_Clickhouse) {
	m.Clickhouse = v
}

func (m *ShardConfigSpec_Clickhouse) SetConfig(v *config.ClickhouseConfig) {
	m.Config = v
}

func (m *ShardConfigSpec_Clickhouse) SetResources(v *Resources) {
	m.Resources = v
}

func (m *ShardConfigSpec_Clickhouse) SetWeight(v *wrapperspb.Int64Value) {
	m.Weight = v
}
