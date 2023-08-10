// Code generated by protoc-gen-goext. DO NOT EDIT.

package ydb

import (
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

type Database_DatabaseType = isDatabase_DatabaseType

func (m *Database) SetDatabaseType(v Database_DatabaseType) {
	m.DatabaseType = v
}

func (m *Database) SetId(v string) {
	m.Id = v
}

func (m *Database) SetFolderId(v string) {
	m.FolderId = v
}

func (m *Database) SetCreatedAt(v *timestamppb.Timestamp) {
	m.CreatedAt = v
}

func (m *Database) SetName(v string) {
	m.Name = v
}

func (m *Database) SetDescription(v string) {
	m.Description = v
}

func (m *Database) SetStatus(v Database_Status) {
	m.Status = v
}

func (m *Database) SetEndpoint(v string) {
	m.Endpoint = v
}

func (m *Database) SetResourcePresetId(v string) {
	m.ResourcePresetId = v
}

func (m *Database) SetStorageConfig(v *StorageConfig) {
	m.StorageConfig = v
}

func (m *Database) SetScalePolicy(v *ScalePolicy) {
	m.ScalePolicy = v
}

func (m *Database) SetNetworkId(v string) {
	m.NetworkId = v
}

func (m *Database) SetSubnetIds(v []string) {
	m.SubnetIds = v
}

func (m *Database) SetZonalDatabase(v *ZonalDatabase) {
	m.DatabaseType = &Database_ZonalDatabase{
		ZonalDatabase: v,
	}
}

func (m *Database) SetRegionalDatabase(v *RegionalDatabase) {
	m.DatabaseType = &Database_RegionalDatabase{
		RegionalDatabase: v,
	}
}

func (m *Database) SetDedicatedDatabase(v *DedicatedDatabase) {
	m.DatabaseType = &Database_DedicatedDatabase{
		DedicatedDatabase: v,
	}
}

func (m *Database) SetServerlessDatabase(v *ServerlessDatabase) {
	m.DatabaseType = &Database_ServerlessDatabase{
		ServerlessDatabase: v,
	}
}

func (m *Database) SetAssignPublicIps(v bool) {
	m.AssignPublicIps = v
}

func (m *Database) SetLocationId(v string) {
	m.LocationId = v
}

func (m *Database) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *Database) SetBackupConfig(v *BackupConfig) {
	m.BackupConfig = v
}

func (m *Database) SetDocumentApiEndpoint(v string) {
	m.DocumentApiEndpoint = v
}

func (m *Database) SetKinesisApiEndpoint(v string) {
	m.KinesisApiEndpoint = v
}

func (m *Database) SetKafkaApiEndpoint(v string) {
	m.KafkaApiEndpoint = v
}

func (m *Database) SetMonitoringConfig(v *MonitoringConfig) {
	m.MonitoringConfig = v
}

func (m *Database) SetDeletionProtection(v bool) {
	m.DeletionProtection = v
}

type AlertParameter_Parameter = isAlertParameter_Parameter

func (m *AlertParameter) SetParameter(v AlertParameter_Parameter) {
	m.Parameter = v
}

func (m *AlertParameter) SetDoubleParameterValue(v *AlertParameter_DoubleParameterValue) {
	m.Parameter = &AlertParameter_DoubleParameterValue_{
		DoubleParameterValue: v,
	}
}

func (m *AlertParameter) SetIntegerParameterValue(v *AlertParameter_IntegerParameterValue) {
	m.Parameter = &AlertParameter_IntegerParameterValue_{
		IntegerParameterValue: v,
	}
}

func (m *AlertParameter) SetTextParameterValue(v *AlertParameter_TextParameterValue) {
	m.Parameter = &AlertParameter_TextParameterValue_{
		TextParameterValue: v,
	}
}

func (m *AlertParameter) SetTextListParameterValue(v *AlertParameter_TextListParameterValue) {
	m.Parameter = &AlertParameter_TextListParameterValue_{
		TextListParameterValue: v,
	}
}

func (m *AlertParameter) SetLabelListParameterValue(v *AlertParameter_LabelListParameterValue) {
	m.Parameter = &AlertParameter_LabelListParameterValue_{
		LabelListParameterValue: v,
	}
}

func (m *AlertParameter_DoubleParameterValue) SetName(v string) {
	m.Name = v
}

func (m *AlertParameter_DoubleParameterValue) SetValue(v float64) {
	m.Value = v
}

func (m *AlertParameter_IntegerParameterValue) SetName(v string) {
	m.Name = v
}

func (m *AlertParameter_IntegerParameterValue) SetValue(v int64) {
	m.Value = v
}

func (m *AlertParameter_TextParameterValue) SetName(v string) {
	m.Name = v
}

func (m *AlertParameter_TextParameterValue) SetValue(v string) {
	m.Value = v
}

func (m *AlertParameter_TextListParameterValue) SetName(v string) {
	m.Name = v
}

func (m *AlertParameter_TextListParameterValue) SetValues(v []string) {
	m.Values = v
}

func (m *AlertParameter_LabelListParameterValue) SetName(v string) {
	m.Name = v
}

func (m *AlertParameter_LabelListParameterValue) SetValues(v []string) {
	m.Values = v
}

func (m *NotificationChannel) SetNotificationChannelId(v string) {
	m.NotificationChannelId = v
}

func (m *NotificationChannel) SetNotifyAboutStatuses(v []AlertEvaluationStatus) {
	m.NotifyAboutStatuses = v
}

func (m *NotificationChannel) SetRepeateNotifyDelayMs(v int64) {
	m.RepeateNotifyDelayMs = v
}

func (m *Alert) SetAlertId(v string) {
	m.AlertId = v
}

func (m *Alert) SetAlertTemplateId(v string) {
	m.AlertTemplateId = v
}

func (m *Alert) SetName(v string) {
	m.Name = v
}

func (m *Alert) SetDescription(v string) {
	m.Description = v
}

func (m *Alert) SetNotificationChannels(v []*NotificationChannel) {
	m.NotificationChannels = v
}

func (m *Alert) SetAlertParameters(v []*AlertParameter) {
	m.AlertParameters = v
}

func (m *Alert) SetAlertThresholds(v []*AlertParameter) {
	m.AlertThresholds = v
}

func (m *MonitoringConfig) SetAlerts(v []*Alert) {
	m.Alerts = v
}

func (m *DedicatedDatabase) SetResourcePresetId(v string) {
	m.ResourcePresetId = v
}

func (m *DedicatedDatabase) SetStorageConfig(v *StorageConfig) {
	m.StorageConfig = v
}

func (m *DedicatedDatabase) SetScalePolicy(v *ScalePolicy) {
	m.ScalePolicy = v
}

func (m *DedicatedDatabase) SetNetworkId(v string) {
	m.NetworkId = v
}

func (m *DedicatedDatabase) SetSubnetIds(v []string) {
	m.SubnetIds = v
}

func (m *DedicatedDatabase) SetAssignPublicIps(v bool) {
	m.AssignPublicIps = v
}

func (m *ServerlessDatabase) SetThrottlingRcuLimit(v int64) {
	m.ThrottlingRcuLimit = v
}

func (m *ServerlessDatabase) SetStorageSizeLimit(v int64) {
	m.StorageSizeLimit = v
}

func (m *ServerlessDatabase) SetEnableThrottlingRcuLimit(v bool) {
	m.EnableThrottlingRcuLimit = v
}

func (m *ServerlessDatabase) SetProvisionedRcuLimit(v int64) {
	m.ProvisionedRcuLimit = v
}

func (m *ServerlessDatabase) SetTopicWriteQuota(v int64) {
	m.TopicWriteQuota = v
}

func (m *ZonalDatabase) SetZoneId(v string) {
	m.ZoneId = v
}

func (m *RegionalDatabase) SetRegionId(v string) {
	m.RegionId = v
}

type ScalePolicy_ScaleType = isScalePolicy_ScaleType

func (m *ScalePolicy) SetScaleType(v ScalePolicy_ScaleType) {
	m.ScaleType = v
}

func (m *ScalePolicy) SetFixedScale(v *ScalePolicy_FixedScale) {
	m.ScaleType = &ScalePolicy_FixedScale_{
		FixedScale: v,
	}
}

func (m *ScalePolicy_FixedScale) SetSize(v int64) {
	m.Size = v
}

func (m *StorageConfig) SetStorageOptions(v []*StorageOption) {
	m.StorageOptions = v
}

func (m *StorageConfig) SetStorageSizeLimit(v int64) {
	m.StorageSizeLimit = v
}

func (m *StorageOption) SetStorageTypeId(v string) {
	m.StorageTypeId = v
}

func (m *StorageOption) SetGroupCount(v int64) {
	m.GroupCount = v
}
