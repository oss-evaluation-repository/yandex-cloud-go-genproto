// Code generated by protoc-gen-goext. DO NOT EDIT.

package endpoint

import (
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func (m *AltName) SetFromName(v string) {
	m.FromName = v
}

func (m *AltName) SetToName(v string) {
	m.ToName = v
}

type Secret_Value = isSecret_Value

func (m *Secret) SetValue(v Secret_Value) {
	m.Value = v
}

func (m *Secret) SetRaw(v string) {
	m.Value = &Secret_Raw{
		Raw: v,
	}
}

func (m *ColSchema) SetName(v string) {
	m.Name = v
}

func (m *ColSchema) SetType(v ColumnType) {
	m.Type = v
}

func (m *ColSchema) SetKey(v bool) {
	m.Key = v
}

func (m *ColSchema) SetRequired(v bool) {
	m.Required = v
}

func (m *ColSchema) SetPath(v string) {
	m.Path = v
}

type TLSMode_TlsMode = isTLSMode_TlsMode

func (m *TLSMode) SetTlsMode(v TLSMode_TlsMode) {
	m.TlsMode = v
}

func (m *TLSMode) SetDisabled(v *emptypb.Empty) {
	m.TlsMode = &TLSMode_Disabled{
		Disabled: v,
	}
}

func (m *TLSMode) SetEnabled(v *TLSConfig) {
	m.TlsMode = &TLSMode_Enabled{
		Enabled: v,
	}
}

func (m *TLSConfig) SetCaCertificate(v string) {
	m.CaCertificate = v
}

type ColumnValue_Value = isColumnValue_Value

func (m *ColumnValue) SetValue(v ColumnValue_Value) {
	m.Value = v
}

func (m *ColumnValue) SetStringValue(v string) {
	m.Value = &ColumnValue_StringValue{
		StringValue: v,
	}
}

func (m *DataTransformationOptions) SetCloudFunction(v string) {
	m.CloudFunction = v
}

func (m *DataTransformationOptions) SetNumberOfRetries(v int64) {
	m.NumberOfRetries = v
}

func (m *DataTransformationOptions) SetBufferSize(v string) {
	m.BufferSize = v
}

func (m *DataTransformationOptions) SetBufferFlushInterval(v string) {
	m.BufferFlushInterval = v
}

func (m *DataTransformationOptions) SetInvocationTimeout(v string) {
	m.InvocationTimeout = v
}

func (m *DataTransformationOptions) SetServiceAccountId(v string) {
	m.ServiceAccountId = v
}

func (m *FieldList) SetFields(v []*ColSchema) {
	m.Fields = v
}

type DataSchema_Schema = isDataSchema_Schema

func (m *DataSchema) SetSchema(v DataSchema_Schema) {
	m.Schema = v
}

func (m *DataSchema) SetFields(v *FieldList) {
	m.Schema = &DataSchema_Fields{
		Fields: v,
	}
}

func (m *DataSchema) SetJsonFields(v string) {
	m.Schema = &DataSchema_JsonFields{
		JsonFields: v,
	}
}
