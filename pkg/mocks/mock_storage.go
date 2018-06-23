// Code generated by MockGen. DO NOT EDIT.
// Source: storage.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	uploader "github.com/schepelin/imgupload/pkg/uploader"
	reflect "reflect"
)

// MockImagesStorage is a mock of ImagesStorage interface
type MockImagesStorage struct {
	ctrl     *gomock.Controller
	recorder *MockImagesStorageMockRecorder
}

// MockImagesStorageMockRecorder is the mock recorder for MockImagesStorage
type MockImagesStorageMockRecorder struct {
	mock *MockImagesStorage
}

// NewMockImagesStorage creates a new mock instance
func NewMockImagesStorage(ctrl *gomock.Controller) *MockImagesStorage {
	mock := &MockImagesStorage{ctrl: ctrl}
	mock.recorder = &MockImagesStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockImagesStorage) EXPECT() *MockImagesStorageMockRecorder {
	return m.recorder
}

// SaveImage mocks base method
func (m *MockImagesStorage) SaveImage(ctx context.Context, img uploader.Image) error {
	ret := m.ctrl.Call(m, "SaveImage", ctx, img)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveImage indicates an expected call of SaveImage
func (mr *MockImagesStorageMockRecorder) SaveImage(ctx, img interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveImage", reflect.TypeOf((*MockImagesStorage)(nil).SaveImage), ctx, img)
}

// GetImage mocks base method
func (m *MockImagesStorage) GetImage(ctx context.Context, imgID string) (*uploader.Image, error) {
	ret := m.ctrl.Call(m, "GetImage", ctx, imgID)
	ret0, _ := ret[0].(*uploader.Image)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetImage indicates an expected call of GetImage
func (mr *MockImagesStorageMockRecorder) GetImage(ctx, imgID interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetImage", reflect.TypeOf((*MockImagesStorage)(nil).GetImage), ctx, imgID)
}