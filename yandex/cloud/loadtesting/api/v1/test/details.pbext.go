// Code generated by protoc-gen-goext. DO NOT EDIT.

package test

import (
	common "github.com/yandex-cloud/go-genproto/yandex/cloud/loadtesting/api/v1/common"
)

func (m *Details) SetName(v string) {
	m.Name = v
}

func (m *Details) SetDescription(v string) {
	m.Description = v
}

func (m *Details) SetTags(v []*common.Tag) {
	m.Tags = v
}

func (m *Details) SetLoggingLogGroupId(v string) {
	m.LoggingLogGroupId = v
}
