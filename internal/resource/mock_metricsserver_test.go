// Code generated by pegomock. DO NOT EDIT.
// Source: github.com/GusTheSadGeek/k9s/internal/resource (interfaces: MetricsServer)

package resource_test

import (
	k8s "github.com/GusTheSadGeek/k9s/internal/k8s"
	pegomock "github.com/petergtz/pegomock"
	v1beta1 "k8s.io/metrics/pkg/apis/metrics/v1beta1"
	"reflect"
	"time"
)

type MockMetricsServer struct {
	fail func(message string, callerSkip ...int)
}

func NewMockMetricsServer(options ...pegomock.Option) *MockMetricsServer {
	mock := &MockMetricsServer{}
	for _, option := range options {
		option.Apply(mock)
	}
	return mock
}

func (mock *MockMetricsServer) SetFailHandler(fh pegomock.FailHandler) { mock.fail = fh }
func (mock *MockMetricsServer) FailHandler() pegomock.FailHandler      { return mock.fail }

func (mock *MockMetricsServer) ClusterLoad(_param0 k8s.Collection, _param1 k8s.Collection, _param2 *k8s.ClusterMetrics) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockMetricsServer().")
	}
	params := []pegomock.Param{_param0, _param1, _param2}
	pegomock.GetGenericMockFrom(mock).Invoke("ClusterLoad", params, []reflect.Type{})
}

func (mock *MockMetricsServer) FetchNodesMetrics() (*v1beta1.NodeMetricsList, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockMetricsServer().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("FetchNodesMetrics", params, []reflect.Type{reflect.TypeOf((**v1beta1.NodeMetricsList)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 *v1beta1.NodeMetricsList
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(*v1beta1.NodeMetricsList)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockMetricsServer) FetchPodsMetrics(_param0 string) (*v1beta1.PodMetricsList, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockMetricsServer().")
	}
	params := []pegomock.Param{_param0}
	result := pegomock.GetGenericMockFrom(mock).Invoke("FetchPodsMetrics", params, []reflect.Type{reflect.TypeOf((**v1beta1.PodMetricsList)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 *v1beta1.PodMetricsList
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(*v1beta1.PodMetricsList)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockMetricsServer) HasMetrics() bool {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockMetricsServer().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("HasMetrics", params, []reflect.Type{reflect.TypeOf((*bool)(nil)).Elem()})
	var ret0 bool
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(bool)
		}
	}
	return ret0
}

func (mock *MockMetricsServer) NodesMetrics(_param0 k8s.Collection, _param1 *v1beta1.NodeMetricsList, _param2 k8s.NodesMetrics) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockMetricsServer().")
	}
	params := []pegomock.Param{_param0, _param1, _param2}
	pegomock.GetGenericMockFrom(mock).Invoke("NodesMetrics", params, []reflect.Type{})
}

func (mock *MockMetricsServer) PodsMetrics(_param0 *v1beta1.PodMetricsList, _param1 k8s.PodsMetrics) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockMetricsServer().")
	}
	params := []pegomock.Param{_param0, _param1}
	pegomock.GetGenericMockFrom(mock).Invoke("PodsMetrics", params, []reflect.Type{})
}

func (mock *MockMetricsServer) VerifyWasCalledOnce() *VerifierMockMetricsServer {
	return &VerifierMockMetricsServer{
		mock:                   mock,
		invocationCountMatcher: pegomock.Times(1),
	}
}

func (mock *MockMetricsServer) VerifyWasCalled(invocationCountMatcher pegomock.Matcher) *VerifierMockMetricsServer {
	return &VerifierMockMetricsServer{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
	}
}

func (mock *MockMetricsServer) VerifyWasCalledInOrder(invocationCountMatcher pegomock.Matcher, inOrderContext *pegomock.InOrderContext) *VerifierMockMetricsServer {
	return &VerifierMockMetricsServer{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		inOrderContext:         inOrderContext,
	}
}

func (mock *MockMetricsServer) VerifyWasCalledEventually(invocationCountMatcher pegomock.Matcher, timeout time.Duration) *VerifierMockMetricsServer {
	return &VerifierMockMetricsServer{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		timeout:                timeout,
	}
}

type VerifierMockMetricsServer struct {
	mock                   *MockMetricsServer
	invocationCountMatcher pegomock.Matcher
	inOrderContext         *pegomock.InOrderContext
	timeout                time.Duration
}

func (verifier *VerifierMockMetricsServer) ClusterLoad(_param0 k8s.Collection, _param1 k8s.Collection, _param2 *k8s.ClusterMetrics) *MockMetricsServer_ClusterLoad_OngoingVerification {
	params := []pegomock.Param{_param0, _param1, _param2}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "ClusterLoad", params, verifier.timeout)
	return &MockMetricsServer_ClusterLoad_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockMetricsServer_ClusterLoad_OngoingVerification struct {
	mock              *MockMetricsServer
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockMetricsServer_ClusterLoad_OngoingVerification) GetCapturedArguments() (k8s.Collection, k8s.Collection, *k8s.ClusterMetrics) {
	_param0, _param1, _param2 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1], _param1[len(_param1)-1], _param2[len(_param2)-1]
}

func (c *MockMetricsServer_ClusterLoad_OngoingVerification) GetAllCapturedArguments() (_param0 []k8s.Collection, _param1 []k8s.Collection, _param2 []*k8s.ClusterMetrics) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]k8s.Collection, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(k8s.Collection)
		}
		_param1 = make([]k8s.Collection, len(params[1]))
		for u, param := range params[1] {
			_param1[u] = param.(k8s.Collection)
		}
		_param2 = make([]*k8s.ClusterMetrics, len(params[2]))
		for u, param := range params[2] {
			_param2[u] = param.(*k8s.ClusterMetrics)
		}
	}
	return
}

func (verifier *VerifierMockMetricsServer) FetchNodesMetrics() *MockMetricsServer_FetchNodesMetrics_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "FetchNodesMetrics", params, verifier.timeout)
	return &MockMetricsServer_FetchNodesMetrics_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockMetricsServer_FetchNodesMetrics_OngoingVerification struct {
	mock              *MockMetricsServer
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockMetricsServer_FetchNodesMetrics_OngoingVerification) GetCapturedArguments() {
}

func (c *MockMetricsServer_FetchNodesMetrics_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierMockMetricsServer) FetchPodsMetrics(_param0 string) *MockMetricsServer_FetchPodsMetrics_OngoingVerification {
	params := []pegomock.Param{_param0}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "FetchPodsMetrics", params, verifier.timeout)
	return &MockMetricsServer_FetchPodsMetrics_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockMetricsServer_FetchPodsMetrics_OngoingVerification struct {
	mock              *MockMetricsServer
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockMetricsServer_FetchPodsMetrics_OngoingVerification) GetCapturedArguments() string {
	_param0 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1]
}

func (c *MockMetricsServer_FetchPodsMetrics_OngoingVerification) GetAllCapturedArguments() (_param0 []string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
	}
	return
}

func (verifier *VerifierMockMetricsServer) HasMetrics() *MockMetricsServer_HasMetrics_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "HasMetrics", params, verifier.timeout)
	return &MockMetricsServer_HasMetrics_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockMetricsServer_HasMetrics_OngoingVerification struct {
	mock              *MockMetricsServer
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockMetricsServer_HasMetrics_OngoingVerification) GetCapturedArguments() {
}

func (c *MockMetricsServer_HasMetrics_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierMockMetricsServer) NodesMetrics(_param0 k8s.Collection, _param1 *v1beta1.NodeMetricsList, _param2 k8s.NodesMetrics) *MockMetricsServer_NodesMetrics_OngoingVerification {
	params := []pegomock.Param{_param0, _param1, _param2}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "NodesMetrics", params, verifier.timeout)
	return &MockMetricsServer_NodesMetrics_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockMetricsServer_NodesMetrics_OngoingVerification struct {
	mock              *MockMetricsServer
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockMetricsServer_NodesMetrics_OngoingVerification) GetCapturedArguments() (k8s.Collection, *v1beta1.NodeMetricsList, k8s.NodesMetrics) {
	_param0, _param1, _param2 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1], _param1[len(_param1)-1], _param2[len(_param2)-1]
}

func (c *MockMetricsServer_NodesMetrics_OngoingVerification) GetAllCapturedArguments() (_param0 []k8s.Collection, _param1 []*v1beta1.NodeMetricsList, _param2 []k8s.NodesMetrics) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]k8s.Collection, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(k8s.Collection)
		}
		_param1 = make([]*v1beta1.NodeMetricsList, len(params[1]))
		for u, param := range params[1] {
			_param1[u] = param.(*v1beta1.NodeMetricsList)
		}
		_param2 = make([]k8s.NodesMetrics, len(params[2]))
		for u, param := range params[2] {
			_param2[u] = param.(k8s.NodesMetrics)
		}
	}
	return
}

func (verifier *VerifierMockMetricsServer) PodsMetrics(_param0 *v1beta1.PodMetricsList, _param1 k8s.PodsMetrics) *MockMetricsServer_PodsMetrics_OngoingVerification {
	params := []pegomock.Param{_param0, _param1}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "PodsMetrics", params, verifier.timeout)
	return &MockMetricsServer_PodsMetrics_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockMetricsServer_PodsMetrics_OngoingVerification struct {
	mock              *MockMetricsServer
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockMetricsServer_PodsMetrics_OngoingVerification) GetCapturedArguments() (*v1beta1.PodMetricsList, k8s.PodsMetrics) {
	_param0, _param1 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1], _param1[len(_param1)-1]
}

func (c *MockMetricsServer_PodsMetrics_OngoingVerification) GetAllCapturedArguments() (_param0 []*v1beta1.PodMetricsList, _param1 []k8s.PodsMetrics) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]*v1beta1.PodMetricsList, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(*v1beta1.PodMetricsList)
		}
		_param1 = make([]k8s.PodsMetrics, len(params[1]))
		for u, param := range params[1] {
			_param1[u] = param.(k8s.PodsMetrics)
		}
	}
	return
}
