// Code generated by protoc-gen-goext. DO NOT EDIT.

package dataproc_manager

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
)

type Job_JobSpec = isJob_JobSpec

func (m *Job) SetJobSpec(v Job_JobSpec) {
	m.JobSpec = v
}

func (m *Job) SetId(v string) {
	m.Id = v
}

func (m *Job) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *Job) SetCreatedAt(v *timestamp.Timestamp) {
	m.CreatedAt = v
}

func (m *Job) SetStartedAt(v *timestamp.Timestamp) {
	m.StartedAt = v
}

func (m *Job) SetFinishedAt(v *timestamp.Timestamp) {
	m.FinishedAt = v
}

func (m *Job) SetName(v string) {
	m.Name = v
}

func (m *Job) SetCreatedBy(v string) {
	m.CreatedBy = v
}

func (m *Job) SetStatus(v Job_Status) {
	m.Status = v
}

func (m *Job) SetMapreduceJob(v *MapreduceJob) {
	m.JobSpec = &Job_MapreduceJob{
		MapreduceJob: v,
	}
}

func (m *Job) SetSparkJob(v *SparkJob) {
	m.JobSpec = &Job_SparkJob{
		SparkJob: v,
	}
}

func (m *Job) SetPysparkJob(v *PysparkJob) {
	m.JobSpec = &Job_PysparkJob{
		PysparkJob: v,
	}
}

func (m *Job) SetHiveJob(v *HiveJob) {
	m.JobSpec = &Job_HiveJob{
		HiveJob: v,
	}
}

func (m *Job) SetApplicationInfo(v *ApplicationInfo) {
	m.ApplicationInfo = v
}

func (m *ApplicationAttempt) SetId(v string) {
	m.Id = v
}

func (m *ApplicationAttempt) SetAmContainerId(v string) {
	m.AmContainerId = v
}

func (m *ApplicationInfo) SetId(v string) {
	m.Id = v
}

func (m *ApplicationInfo) SetApplicationAttempts(v []*ApplicationAttempt) {
	m.ApplicationAttempts = v
}

type MapreduceJob_Driver = isMapreduceJob_Driver

func (m *MapreduceJob) SetDriver(v MapreduceJob_Driver) {
	m.Driver = v
}

func (m *MapreduceJob) SetArgs(v []string) {
	m.Args = v
}

func (m *MapreduceJob) SetJarFileUris(v []string) {
	m.JarFileUris = v
}

func (m *MapreduceJob) SetFileUris(v []string) {
	m.FileUris = v
}

func (m *MapreduceJob) SetArchiveUris(v []string) {
	m.ArchiveUris = v
}

func (m *MapreduceJob) SetProperties(v map[string]string) {
	m.Properties = v
}

func (m *MapreduceJob) SetMainJarFileUri(v string) {
	m.Driver = &MapreduceJob_MainJarFileUri{
		MainJarFileUri: v,
	}
}

func (m *MapreduceJob) SetMainClass(v string) {
	m.Driver = &MapreduceJob_MainClass{
		MainClass: v,
	}
}

func (m *SparkJob) SetArgs(v []string) {
	m.Args = v
}

func (m *SparkJob) SetJarFileUris(v []string) {
	m.JarFileUris = v
}

func (m *SparkJob) SetFileUris(v []string) {
	m.FileUris = v
}

func (m *SparkJob) SetArchiveUris(v []string) {
	m.ArchiveUris = v
}

func (m *SparkJob) SetProperties(v map[string]string) {
	m.Properties = v
}

func (m *SparkJob) SetMainJarFileUri(v string) {
	m.MainJarFileUri = v
}

func (m *SparkJob) SetMainClass(v string) {
	m.MainClass = v
}

func (m *PysparkJob) SetArgs(v []string) {
	m.Args = v
}

func (m *PysparkJob) SetJarFileUris(v []string) {
	m.JarFileUris = v
}

func (m *PysparkJob) SetFileUris(v []string) {
	m.FileUris = v
}

func (m *PysparkJob) SetArchiveUris(v []string) {
	m.ArchiveUris = v
}

func (m *PysparkJob) SetProperties(v map[string]string) {
	m.Properties = v
}

func (m *PysparkJob) SetMainPythonFileUri(v string) {
	m.MainPythonFileUri = v
}

func (m *PysparkJob) SetPythonFileUris(v []string) {
	m.PythonFileUris = v
}

func (m *QueryList) SetQueries(v []string) {
	m.Queries = v
}

type HiveJob_QueryType = isHiveJob_QueryType

func (m *HiveJob) SetQueryType(v HiveJob_QueryType) {
	m.QueryType = v
}

func (m *HiveJob) SetProperties(v map[string]string) {
	m.Properties = v
}

func (m *HiveJob) SetContinueOnFailure(v bool) {
	m.ContinueOnFailure = v
}

func (m *HiveJob) SetScriptVariables(v map[string]string) {
	m.ScriptVariables = v
}

func (m *HiveJob) SetJarFileUris(v []string) {
	m.JarFileUris = v
}

func (m *HiveJob) SetQueryFileUri(v string) {
	m.QueryType = &HiveJob_QueryFileUri{
		QueryFileUri: v,
	}
}

func (m *HiveJob) SetQueryList(v *QueryList) {
	m.QueryType = &HiveJob_QueryList{
		QueryList: v,
	}
}
