// Code generated by protoc-gen-goext. DO NOT EDIT.

package datasphere

import (
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

func (m *CreateProjectJobRequest) SetProjectId(v string) {
	m.ProjectId = v
}

func (m *CreateProjectJobRequest) SetJobParameters(v *JobParameters) {
	m.JobParameters = v
}

func (m *CreateProjectJobRequest) SetConfig(v string) {
	m.Config = v
}

func (m *CreateProjectJobRequest) SetName(v string) {
	m.Name = v
}

func (m *CreateProjectJobRequest) SetDesc(v string) {
	m.Desc = v
}

func (m *CreateProjectJobMetadata) SetProjectId(v string) {
	m.ProjectId = v
}

func (m *CreateProjectJobResponse) SetJobId(v string) {
	m.JobId = v
}

func (m *CreateProjectJobResponse) SetUploadFiles(v []*StorageFile) {
	m.UploadFiles = v
}

func (m *ExecuteProjectJobRequest) SetJobId(v string) {
	m.JobId = v
}

func (m *ExecuteProjectJobResponse) SetOutputFiles(v []*StorageFile) {
	m.OutputFiles = v
}

func (m *ExecuteProjectJobResponse) SetResult(v *JobResult) {
	m.Result = v
}

func (m *ExecuteProjectJobMetadata) SetJob(v *Job) {
	m.Job = v
}

func (m *CancelProjectJobRequest) SetJobId(v string) {
	m.JobId = v
}

func (m *CancelProjectJobRequest) SetReason(v string) {
	m.Reason = v
}

func (m *FinalizeProjectJobRequest) SetJobId(v string) {
	m.JobId = v
}

func (m *FinalizeProjectJobResponse) SetJobId(v string) {
	m.JobId = v
}

func (m *ReadProjectJobStdLogsRequest) SetJobId(v string) {
	m.JobId = v
}

func (m *ReadProjectJobStdLogsRequest) SetOffset(v int64) {
	m.Offset = v
}

func (m *ReadProjectJobStdLogsResponse) SetLogs(v []*StdLog) {
	m.Logs = v
}

func (m *ReadProjectJobStdLogsResponse) SetOffset(v int64) {
	m.Offset = v
}

func (m *ReadProjectJobLogsRequest) SetJobId(v string) {
	m.JobId = v
}

func (m *ReadProjectJobLogsRequest) SetOffset(v int64) {
	m.Offset = v
}

func (m *ReadProjectJobLogsResponse) SetLogs(v []*LogMessage) {
	m.Logs = v
}

func (m *ReadProjectJobLogsResponse) SetOffset(v int64) {
	m.Offset = v
}

func (m *ListProjectJobRequest) SetProjectId(v string) {
	m.ProjectId = v
}

func (m *ListProjectJobRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListProjectJobRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListProjectJobResponse) SetJobs(v []*Job) {
	m.Jobs = v
}

func (m *ListProjectJobResponse) SetPageToken(v string) {
	m.PageToken = v
}

func (m *GetProjectJobRequest) SetJobId(v string) {
	m.JobId = v
}

func (m *DeleteProjectJobRequest) SetJobId(v string) {
	m.JobId = v
}

func (m *DeleteProjectJobMetadata) SetJobId(v string) {
	m.JobId = v
}

func (m *StdLog) SetContent(v []byte) {
	m.Content = v
}

func (m *StdLog) SetType(v StdLog_Type) {
	m.Type = v
}

type LogMessage_Source = isLogMessage_Source

func (m *LogMessage) SetSource(v LogMessage_Source) {
	m.Source = v
}

func (m *LogMessage) SetContent(v []byte) {
	m.Content = v
}

func (m *LogMessage) SetCreatedAt(v *timestamppb.Timestamp) {
	m.CreatedAt = v
}

func (m *LogMessage) SetStandardStream(v StandardStream) {
	m.Source = &LogMessage_StandardStream{
		StandardStream: v,
	}
}

func (m *LogMessage) SetFilePath(v string) {
	m.Source = &LogMessage_FilePath{
		FilePath: v,
	}
}
