// Code generated by MockGen. DO NOT EDIT.
// Source: metric.go

// Package metrics is a generated GoMock package.
package metrics

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	time "time"
)

// MockMetric is a mock of Metric interface
type MockMetric struct {
	ctrl     *gomock.Controller
	recorder *MockMetricMockRecorder
}

// MockMetricMockRecorder is the mock recorder for MockMetric
type MockMetricMockRecorder struct {
	mock *MockMetric
}

// NewMockMetric creates a new mock instance
func NewMockMetric(ctrl *gomock.Controller) *MockMetric {
	mock := &MockMetric{ctrl: ctrl}
	mock.recorder = &MockMetricMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMetric) EXPECT() *MockMetricMockRecorder {
	return m.recorder
}

// Name mocks base method
func (m *MockMetric) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name
func (mr *MockMetricMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockMetric)(nil).Name))
}

// Tags mocks base method
func (m *MockMetric) Tags() map[string]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Tags")
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

// Tags indicates an expected call of Tags
func (mr *MockMetricMockRecorder) Tags() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tags", reflect.TypeOf((*MockMetric)(nil).Tags))
}

// Type mocks base method
func (m *MockMetric) Type() MetricType {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Type")
	ret0, _ := ret[0].(MetricType)
	return ret0
}

// Type indicates an expected call of Type
func (mr *MockMetricMockRecorder) Type() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Type", reflect.TypeOf((*MockMetric)(nil).Type))
}

// MockGaugeMetric is a mock of GaugeMetric interface
type MockGaugeMetric struct {
	ctrl     *gomock.Controller
	recorder *MockGaugeMetricMockRecorder
}

// MockGaugeMetricMockRecorder is the mock recorder for MockGaugeMetric
type MockGaugeMetricMockRecorder struct {
	mock *MockGaugeMetric
}

// NewMockGaugeMetric creates a new mock instance
func NewMockGaugeMetric(ctrl *gomock.Controller) *MockGaugeMetric {
	mock := &MockGaugeMetric{ctrl: ctrl}
	mock.recorder = &MockGaugeMetricMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGaugeMetric) EXPECT() *MockGaugeMetricMockRecorder {
	return m.recorder
}

// Name mocks base method
func (m *MockGaugeMetric) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name
func (mr *MockGaugeMetricMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockGaugeMetric)(nil).Name))
}

// Tags mocks base method
func (m *MockGaugeMetric) Tags() map[string]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Tags")
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

// Tags indicates an expected call of Tags
func (mr *MockGaugeMetricMockRecorder) Tags() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tags", reflect.TypeOf((*MockGaugeMetric)(nil).Tags))
}

// Type mocks base method
func (m *MockGaugeMetric) Type() MetricType {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Type")
	ret0, _ := ret[0].(MetricType)
	return ret0
}

// Type indicates an expected call of Type
func (mr *MockGaugeMetricMockRecorder) Type() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Type", reflect.TypeOf((*MockGaugeMetric)(nil).Type))
}

// Value mocks base method
func (m *MockGaugeMetric) Value() float64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Value")
	ret0, _ := ret[0].(float64)
	return ret0
}

// Value indicates an expected call of Value
func (mr *MockGaugeMetricMockRecorder) Value() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Value", reflect.TypeOf((*MockGaugeMetric)(nil).Value))
}

// MockHistogramMetric is a mock of HistogramMetric interface
type MockHistogramMetric struct {
	ctrl     *gomock.Controller
	recorder *MockHistogramMetricMockRecorder
}

// MockHistogramMetricMockRecorder is the mock recorder for MockHistogramMetric
type MockHistogramMetricMockRecorder struct {
	mock *MockHistogramMetric
}

// NewMockHistogramMetric creates a new mock instance
func NewMockHistogramMetric(ctrl *gomock.Controller) *MockHistogramMetric {
	mock := &MockHistogramMetric{ctrl: ctrl}
	mock.recorder = &MockHistogramMetricMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHistogramMetric) EXPECT() *MockHistogramMetricMockRecorder {
	return m.recorder
}

// Name mocks base method
func (m *MockHistogramMetric) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name
func (mr *MockHistogramMetricMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockHistogramMetric)(nil).Name))
}

// Tags mocks base method
func (m *MockHistogramMetric) Tags() map[string]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Tags")
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

// Tags indicates an expected call of Tags
func (mr *MockHistogramMetricMockRecorder) Tags() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tags", reflect.TypeOf((*MockHistogramMetric)(nil).Tags))
}

// Type mocks base method
func (m *MockHistogramMetric) Type() MetricType {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Type")
	ret0, _ := ret[0].(MetricType)
	return ret0
}

// Type indicates an expected call of Type
func (mr *MockHistogramMetricMockRecorder) Type() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Type", reflect.TypeOf((*MockHistogramMetric)(nil).Type))
}

// Value mocks base method
func (m *MockHistogramMetric) Value() float64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Value")
	ret0, _ := ret[0].(float64)
	return ret0
}

// Value indicates an expected call of Value
func (mr *MockHistogramMetricMockRecorder) Value() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Value", reflect.TypeOf((*MockHistogramMetric)(nil).Value))
}

// MockCountMetric is a mock of CountMetric interface
type MockCountMetric struct {
	ctrl     *gomock.Controller
	recorder *MockCountMetricMockRecorder
}

// MockCountMetricMockRecorder is the mock recorder for MockCountMetric
type MockCountMetricMockRecorder struct {
	mock *MockCountMetric
}

// NewMockCountMetric creates a new mock instance
func NewMockCountMetric(ctrl *gomock.Controller) *MockCountMetric {
	mock := &MockCountMetric{ctrl: ctrl}
	mock.recorder = &MockCountMetricMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCountMetric) EXPECT() *MockCountMetricMockRecorder {
	return m.recorder
}

// Name mocks base method
func (m *MockCountMetric) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name
func (mr *MockCountMetricMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockCountMetric)(nil).Name))
}

// Tags mocks base method
func (m *MockCountMetric) Tags() map[string]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Tags")
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

// Tags indicates an expected call of Tags
func (mr *MockCountMetricMockRecorder) Tags() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tags", reflect.TypeOf((*MockCountMetric)(nil).Tags))
}

// Type mocks base method
func (m *MockCountMetric) Type() MetricType {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Type")
	ret0, _ := ret[0].(MetricType)
	return ret0
}

// Type indicates an expected call of Type
func (mr *MockCountMetricMockRecorder) Type() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Type", reflect.TypeOf((*MockCountMetric)(nil).Type))
}

// Value mocks base method
func (m *MockCountMetric) Value() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Value")
	ret0, _ := ret[0].(int64)
	return ret0
}

// Value indicates an expected call of Value
func (mr *MockCountMetricMockRecorder) Value() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Value", reflect.TypeOf((*MockCountMetric)(nil).Value))
}

// MockTimingMetric is a mock of TimingMetric interface
type MockTimingMetric struct {
	ctrl     *gomock.Controller
	recorder *MockTimingMetricMockRecorder
}

// MockTimingMetricMockRecorder is the mock recorder for MockTimingMetric
type MockTimingMetricMockRecorder struct {
	mock *MockTimingMetric
}

// NewMockTimingMetric creates a new mock instance
func NewMockTimingMetric(ctrl *gomock.Controller) *MockTimingMetric {
	mock := &MockTimingMetric{ctrl: ctrl}
	mock.recorder = &MockTimingMetricMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTimingMetric) EXPECT() *MockTimingMetricMockRecorder {
	return m.recorder
}

// Name mocks base method
func (m *MockTimingMetric) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name
func (mr *MockTimingMetricMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockTimingMetric)(nil).Name))
}

// Tags mocks base method
func (m *MockTimingMetric) Tags() map[string]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Tags")
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

// Tags indicates an expected call of Tags
func (mr *MockTimingMetricMockRecorder) Tags() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tags", reflect.TypeOf((*MockTimingMetric)(nil).Tags))
}

// Type mocks base method
func (m *MockTimingMetric) Type() MetricType {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Type")
	ret0, _ := ret[0].(MetricType)
	return ret0
}

// Type indicates an expected call of Type
func (mr *MockTimingMetricMockRecorder) Type() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Type", reflect.TypeOf((*MockTimingMetric)(nil).Type))
}

// Value mocks base method
func (m *MockTimingMetric) Value() time.Duration {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Value")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// Value indicates an expected call of Value
func (mr *MockTimingMetricMockRecorder) Value() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Value", reflect.TypeOf((*MockTimingMetric)(nil).Value))
}

// MockWriter is a mock of Writer interface
type MockWriter struct {
	ctrl     *gomock.Controller
	recorder *MockWriterMockRecorder
}

// MockWriterMockRecorder is the mock recorder for MockWriter
type MockWriterMockRecorder struct {
	mock *MockWriter
}

// NewMockWriter creates a new mock instance
func NewMockWriter(ctrl *gomock.Controller) *MockWriter {
	mock := &MockWriter{ctrl: ctrl}
	mock.recorder = &MockWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockWriter) EXPECT() *MockWriterMockRecorder {
	return m.recorder
}

// Send mocks base method
func (m *MockWriter) Send(arg0 Metric) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send
func (mr *MockWriterMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockWriter)(nil).Send), arg0)
}