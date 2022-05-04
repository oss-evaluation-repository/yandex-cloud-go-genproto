// Code generated by protoc-gen-goext. DO NOT EDIT.

package mysql

import (
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

func (m *MysqlConfig8_0) SetInnodbBufferPoolSize(v *wrapperspb.Int64Value) {
	m.InnodbBufferPoolSize = v
}

func (m *MysqlConfig8_0) SetMaxConnections(v *wrapperspb.Int64Value) {
	m.MaxConnections = v
}

func (m *MysqlConfig8_0) SetLongQueryTime(v *wrapperspb.DoubleValue) {
	m.LongQueryTime = v
}

func (m *MysqlConfig8_0) SetGeneralLog(v *wrapperspb.BoolValue) {
	m.GeneralLog = v
}

func (m *MysqlConfig8_0) SetAuditLog(v *wrapperspb.BoolValue) {
	m.AuditLog = v
}

func (m *MysqlConfig8_0) SetSqlMode(v []MysqlConfig8_0_SQLMode) {
	m.SqlMode = v
}

func (m *MysqlConfig8_0) SetMaxAllowedPacket(v *wrapperspb.Int64Value) {
	m.MaxAllowedPacket = v
}

func (m *MysqlConfig8_0) SetDefaultAuthenticationPlugin(v MysqlConfig8_0_AuthPlugin) {
	m.DefaultAuthenticationPlugin = v
}

func (m *MysqlConfig8_0) SetInnodbFlushLogAtTrxCommit(v *wrapperspb.Int64Value) {
	m.InnodbFlushLogAtTrxCommit = v
}

func (m *MysqlConfig8_0) SetInnodbLockWaitTimeout(v *wrapperspb.Int64Value) {
	m.InnodbLockWaitTimeout = v
}

func (m *MysqlConfig8_0) SetTransactionIsolation(v MysqlConfig8_0_TransactionIsolation) {
	m.TransactionIsolation = v
}

func (m *MysqlConfig8_0) SetInnodbPrintAllDeadlocks(v *wrapperspb.BoolValue) {
	m.InnodbPrintAllDeadlocks = v
}

func (m *MysqlConfig8_0) SetNetReadTimeout(v *wrapperspb.Int64Value) {
	m.NetReadTimeout = v
}

func (m *MysqlConfig8_0) SetNetWriteTimeout(v *wrapperspb.Int64Value) {
	m.NetWriteTimeout = v
}

func (m *MysqlConfig8_0) SetGroupConcatMaxLen(v *wrapperspb.Int64Value) {
	m.GroupConcatMaxLen = v
}

func (m *MysqlConfig8_0) SetTmpTableSize(v *wrapperspb.Int64Value) {
	m.TmpTableSize = v
}

func (m *MysqlConfig8_0) SetMaxHeapTableSize(v *wrapperspb.Int64Value) {
	m.MaxHeapTableSize = v
}

func (m *MysqlConfig8_0) SetDefaultTimeZone(v string) {
	m.DefaultTimeZone = v
}

func (m *MysqlConfig8_0) SetCharacterSetServer(v string) {
	m.CharacterSetServer = v
}

func (m *MysqlConfig8_0) SetCollationServer(v string) {
	m.CollationServer = v
}

func (m *MysqlConfig8_0) SetInnodbAdaptiveHashIndex(v *wrapperspb.BoolValue) {
	m.InnodbAdaptiveHashIndex = v
}

func (m *MysqlConfig8_0) SetInnodbNumaInterleave(v *wrapperspb.BoolValue) {
	m.InnodbNumaInterleave = v
}

func (m *MysqlConfig8_0) SetInnodbLogBufferSize(v *wrapperspb.Int64Value) {
	m.InnodbLogBufferSize = v
}

func (m *MysqlConfig8_0) SetInnodbLogFileSize(v *wrapperspb.Int64Value) {
	m.InnodbLogFileSize = v
}

func (m *MysqlConfig8_0) SetInnodbIoCapacity(v *wrapperspb.Int64Value) {
	m.InnodbIoCapacity = v
}

func (m *MysqlConfig8_0) SetInnodbIoCapacityMax(v *wrapperspb.Int64Value) {
	m.InnodbIoCapacityMax = v
}

func (m *MysqlConfig8_0) SetInnodbReadIoThreads(v *wrapperspb.Int64Value) {
	m.InnodbReadIoThreads = v
}

func (m *MysqlConfig8_0) SetInnodbWriteIoThreads(v *wrapperspb.Int64Value) {
	m.InnodbWriteIoThreads = v
}

func (m *MysqlConfig8_0) SetInnodbPurgeThreads(v *wrapperspb.Int64Value) {
	m.InnodbPurgeThreads = v
}

func (m *MysqlConfig8_0) SetInnodbThreadConcurrency(v *wrapperspb.Int64Value) {
	m.InnodbThreadConcurrency = v
}

func (m *MysqlConfig8_0) SetInnodbTempDataFileMaxSize(v *wrapperspb.Int64Value) {
	m.InnodbTempDataFileMaxSize = v
}

func (m *MysqlConfig8_0) SetThreadCacheSize(v *wrapperspb.Int64Value) {
	m.ThreadCacheSize = v
}

func (m *MysqlConfig8_0) SetThreadStack(v *wrapperspb.Int64Value) {
	m.ThreadStack = v
}

func (m *MysqlConfig8_0) SetJoinBufferSize(v *wrapperspb.Int64Value) {
	m.JoinBufferSize = v
}

func (m *MysqlConfig8_0) SetSortBufferSize(v *wrapperspb.Int64Value) {
	m.SortBufferSize = v
}

func (m *MysqlConfig8_0) SetTableDefinitionCache(v *wrapperspb.Int64Value) {
	m.TableDefinitionCache = v
}

func (m *MysqlConfig8_0) SetTableOpenCache(v *wrapperspb.Int64Value) {
	m.TableOpenCache = v
}

func (m *MysqlConfig8_0) SetTableOpenCacheInstances(v *wrapperspb.Int64Value) {
	m.TableOpenCacheInstances = v
}

func (m *MysqlConfig8_0) SetExplicitDefaultsForTimestamp(v *wrapperspb.BoolValue) {
	m.ExplicitDefaultsForTimestamp = v
}

func (m *MysqlConfig8_0) SetAutoIncrementIncrement(v *wrapperspb.Int64Value) {
	m.AutoIncrementIncrement = v
}

func (m *MysqlConfig8_0) SetAutoIncrementOffset(v *wrapperspb.Int64Value) {
	m.AutoIncrementOffset = v
}

func (m *MysqlConfig8_0) SetSyncBinlog(v *wrapperspb.Int64Value) {
	m.SyncBinlog = v
}

func (m *MysqlConfig8_0) SetBinlogCacheSize(v *wrapperspb.Int64Value) {
	m.BinlogCacheSize = v
}

func (m *MysqlConfig8_0) SetBinlogGroupCommitSyncDelay(v *wrapperspb.Int64Value) {
	m.BinlogGroupCommitSyncDelay = v
}

func (m *MysqlConfig8_0) SetBinlogRowImage(v MysqlConfig8_0_BinlogRowImage) {
	m.BinlogRowImage = v
}

func (m *MysqlConfig8_0) SetBinlogRowsQueryLogEvents(v *wrapperspb.BoolValue) {
	m.BinlogRowsQueryLogEvents = v
}

func (m *MysqlConfig8_0) SetRplSemiSyncMasterWaitForSlaveCount(v *wrapperspb.Int64Value) {
	m.RplSemiSyncMasterWaitForSlaveCount = v
}

func (m *MysqlConfig8_0) SetSlaveParallelType(v MysqlConfig8_0_SlaveParallelType) {
	m.SlaveParallelType = v
}

func (m *MysqlConfig8_0) SetSlaveParallelWorkers(v *wrapperspb.Int64Value) {
	m.SlaveParallelWorkers = v
}

func (m *MysqlConfig8_0) SetRegexpTimeLimit(v *wrapperspb.Int64Value) {
	m.RegexpTimeLimit = v
}

func (m *MysqlConfig8_0) SetMdbPreserveBinlogBytes(v *wrapperspb.Int64Value) {
	m.MdbPreserveBinlogBytes = v
}

func (m *MysqlConfig8_0) SetInteractiveTimeout(v *wrapperspb.Int64Value) {
	m.InteractiveTimeout = v
}

func (m *MysqlConfig8_0) SetWaitTimeout(v *wrapperspb.Int64Value) {
	m.WaitTimeout = v
}

func (m *MysqlConfig8_0) SetMdbOfflineModeEnableLag(v *wrapperspb.Int64Value) {
	m.MdbOfflineModeEnableLag = v
}

func (m *MysqlConfig8_0) SetMdbOfflineModeDisableLag(v *wrapperspb.Int64Value) {
	m.MdbOfflineModeDisableLag = v
}

func (m *MysqlConfig8_0) SetRangeOptimizerMaxMemSize(v *wrapperspb.Int64Value) {
	m.RangeOptimizerMaxMemSize = v
}

func (m *MysqlConfig8_0) SetSlowQueryLog(v *wrapperspb.BoolValue) {
	m.SlowQueryLog = v
}

func (m *MysqlConfig8_0) SetSlowQueryLogAlwaysWriteTime(v *wrapperspb.DoubleValue) {
	m.SlowQueryLogAlwaysWriteTime = v
}

func (m *MysqlConfig8_0) SetLogSlowRateType(v MysqlConfig8_0_LogSlowRateType) {
	m.LogSlowRateType = v
}

func (m *MysqlConfig8_0) SetLogSlowRateLimit(v *wrapperspb.Int64Value) {
	m.LogSlowRateLimit = v
}

func (m *MysqlConfig8_0) SetLogSlowSpStatements(v *wrapperspb.BoolValue) {
	m.LogSlowSpStatements = v
}

func (m *MysqlConfig8_0) SetLogSlowFilter(v []MysqlConfig8_0_LogSlowFilterType) {
	m.LogSlowFilter = v
}

func (m *MysqlConfig8_0) SetMdbPriorityChoiceMaxLag(v *wrapperspb.Int64Value) {
	m.MdbPriorityChoiceMaxLag = v
}

func (m *MysqlConfig8_0) SetInnodbPageSize(v *wrapperspb.Int64Value) {
	m.InnodbPageSize = v
}

func (m *MysqlConfig8_0) SetInnodbOnlineAlterLogMaxSize(v *wrapperspb.Int64Value) {
	m.InnodbOnlineAlterLogMaxSize = v
}

func (m *MysqlConfig8_0) SetInnodbFtMinTokenSize(v *wrapperspb.Int64Value) {
	m.InnodbFtMinTokenSize = v
}

func (m *MysqlConfig8_0) SetInnodbFtMaxTokenSize(v *wrapperspb.Int64Value) {
	m.InnodbFtMaxTokenSize = v
}

func (m *MysqlConfig8_0) SetLowerCaseTableNames(v *wrapperspb.Int64Value) {
	m.LowerCaseTableNames = v
}

func (m *MysqlConfig8_0) SetMaxSpRecursionDepth(v *wrapperspb.Int64Value) {
	m.MaxSpRecursionDepth = v
}

func (m *MysqlConfig8_0) SetInnodbCompressionLevel(v *wrapperspb.Int64Value) {
	m.InnodbCompressionLevel = v
}

func (m *MysqlConfigSet8_0) SetEffectiveConfig(v *MysqlConfig8_0) {
	m.EffectiveConfig = v
}

func (m *MysqlConfigSet8_0) SetUserConfig(v *MysqlConfig8_0) {
	m.UserConfig = v
}

func (m *MysqlConfigSet8_0) SetDefaultConfig(v *MysqlConfig8_0) {
	m.DefaultConfig = v
}
