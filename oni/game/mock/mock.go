// Automatically generated by MockGen. DO NOT EDIT!
// Source: oniproject/oni/game (interfaces: EffectReceiver,FeatureReceiver,SkillTarget,Walkabler)

package mock

import (
	gomock "code.google.com/p/gomock/gomock"
	geom "github.com/oniproject/geom"
)

// Mock of EffectReceiver interface
type MockEffectReceiver struct {
	ctrl     *gomock.Controller
	recorder *_MockEffectReceiverRecorder
}

// Recorder for MockEffectReceiver (not exported)
type _MockEffectReceiverRecorder struct {
	mock *MockEffectReceiver
}

func NewMockEffectReceiver(ctrl *gomock.Controller) *MockEffectReceiver {
	mock := &MockEffectReceiver{ctrl: ctrl}
	mock.recorder = &_MockEffectReceiverRecorder{mock}
	return mock
}

func (_m *MockEffectReceiver) EXPECT() *_MockEffectReceiverRecorder {
	return _m.recorder
}

func (_m *MockEffectReceiver) AddState(_param0 string) {
	_m.ctrl.Call(_m, "AddState", _param0)
}

func (_mr *_MockEffectReceiverRecorder) AddState(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AddState", arg0)
}

func (_m *MockEffectReceiver) RecoverHP(_param0 int) {
	_m.ctrl.Call(_m, "RecoverHP", _param0)
}

func (_mr *_MockEffectReceiverRecorder) RecoverHP(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RecoverHP", arg0)
}

func (_m *MockEffectReceiver) RecoverMP(_param0 int) {
	_m.ctrl.Call(_m, "RecoverMP", _param0)
}

func (_mr *_MockEffectReceiverRecorder) RecoverMP(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RecoverMP", arg0)
}

func (_m *MockEffectReceiver) RecoverTP(_param0 int) {
	_m.ctrl.Call(_m, "RecoverTP", _param0)
}

func (_mr *_MockEffectReceiverRecorder) RecoverTP(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RecoverTP", arg0)
}

func (_m *MockEffectReceiver) RemoveState(_param0 string) {
	_m.ctrl.Call(_m, "RemoveState", _param0)
}

func (_mr *_MockEffectReceiverRecorder) RemoveState(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RemoveState", arg0)
}

// Mock of FeatureReceiver interface
type MockFeatureReceiver struct {
	ctrl     *gomock.Controller
	recorder *_MockFeatureReceiverRecorder
}

// Recorder for MockFeatureReceiver (not exported)
type _MockFeatureReceiverRecorder struct {
	mock *MockFeatureReceiver
}

func NewMockFeatureReceiver(ctrl *gomock.Controller) *MockFeatureReceiver {
	mock := &MockFeatureReceiver{ctrl: ctrl}
	mock.recorder = &_MockFeatureReceiverRecorder{mock}
	return mock
}

func (_m *MockFeatureReceiver) EXPECT() *_MockFeatureReceiverRecorder {
	return _m.recorder
}

func (_m *MockFeatureReceiver) AddATK(_param0 int) {
	_m.ctrl.Call(_m, "AddATK", _param0)
}

func (_mr *_MockFeatureReceiverRecorder) AddATK(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AddATK", arg0)
}

func (_m *MockFeatureReceiver) AddDEF(_param0 int) {
	_m.ctrl.Call(_m, "AddDEF", _param0)
}

func (_mr *_MockFeatureReceiverRecorder) AddDEF(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AddDEF", arg0)
}

func (_m *MockFeatureReceiver) AddSkill(_param0 string) {
	_m.ctrl.Call(_m, "AddSkill", _param0)
}

func (_mr *_MockFeatureReceiverRecorder) AddSkill(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AddSkill", arg0)
}

func (_m *MockFeatureReceiver) SealSkill(_param0 string) {
	_m.ctrl.Call(_m, "SealSkill", _param0)
}

func (_mr *_MockFeatureReceiverRecorder) SealSkill(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SealSkill", arg0)
}

// Mock of SkillTarget interface
type MockSkillTarget struct {
	ctrl     *gomock.Controller
	recorder *_MockSkillTargetRecorder
}

// Recorder for MockSkillTarget (not exported)
type _MockSkillTargetRecorder struct {
	mock *MockSkillTarget
}

func NewMockSkillTarget(ctrl *gomock.Controller) *MockSkillTarget {
	mock := &MockSkillTarget{ctrl: ctrl}
	mock.recorder = &_MockSkillTargetRecorder{mock}
	return mock
}

func (_m *MockSkillTarget) EXPECT() *_MockSkillTargetRecorder {
	return _m.recorder
}

func (_m *MockSkillTarget) AddState(_param0 string) {
	_m.ctrl.Call(_m, "AddState", _param0)
}

func (_mr *_MockSkillTargetRecorder) AddState(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AddState", arg0)
}

func (_m *MockSkillTarget) Race() int {
	ret := _m.ctrl.Call(_m, "Race")
	ret0, _ := ret[0].(int)
	return ret0
}

func (_mr *_MockSkillTargetRecorder) Race() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Race")
}

func (_m *MockSkillTarget) RecoverHP(_param0 int) {
	_m.ctrl.Call(_m, "RecoverHP", _param0)
}

func (_mr *_MockSkillTargetRecorder) RecoverHP(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RecoverHP", arg0)
}

func (_m *MockSkillTarget) RecoverMP(_param0 int) {
	_m.ctrl.Call(_m, "RecoverMP", _param0)
}

func (_mr *_MockSkillTargetRecorder) RecoverMP(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RecoverMP", arg0)
}

func (_m *MockSkillTarget) RecoverTP(_param0 int) {
	_m.ctrl.Call(_m, "RecoverTP", _param0)
}

func (_mr *_MockSkillTargetRecorder) RecoverTP(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RecoverTP", arg0)
}

func (_m *MockSkillTarget) RemoveState(_param0 string) {
	_m.ctrl.Call(_m, "RemoveState", _param0)
}

func (_mr *_MockSkillTargetRecorder) RemoveState(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RemoveState", arg0)
}

// Mock of Walkabler interface
type MockWalkabler struct {
	ctrl     *gomock.Controller
	recorder *_MockWalkablerRecorder
}

// Recorder for MockWalkabler (not exported)
type _MockWalkablerRecorder struct {
	mock *MockWalkabler
}

func NewMockWalkabler(ctrl *gomock.Controller) *MockWalkabler {
	mock := &MockWalkabler{ctrl: ctrl}
	mock.recorder = &_MockWalkablerRecorder{mock}
	return mock
}

func (_m *MockWalkabler) EXPECT() *_MockWalkablerRecorder {
	return _m.recorder
}

func (_m *MockWalkabler) Walkable(_param0 geom.Coord) bool {
	ret := _m.ctrl.Call(_m, "Walkable", _param0)
	ret0, _ := ret[0].(bool)
	return ret0
}

func (_mr *_MockWalkablerRecorder) Walkable(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Walkable", arg0)
}
