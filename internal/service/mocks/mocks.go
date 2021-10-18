package mocks

import (
	io "io"
	reflect "reflect"

	models "github.com/Nikby53/image-converter/internal/models"
	gomock "github.com/golang/mock/gomock"
)

// MockAuthorization is a mock of Authorization interface.
type MockAuthorization struct {
	ctrl     *gomock.Controller
	recorder *MockAuthorizationMockRecorder
}

// MockAuthorizationMockRecorder is the mock recorder for MockAuthorization.
type MockAuthorizationMockRecorder struct {
	mock *MockAuthorization
}

// NewMockAuthorization creates a new mock instance.
func NewMockAuthorization(ctrl *gomock.Controller) *MockAuthorization {
	mock := &MockAuthorization{ctrl: ctrl}
	mock.recorder = &MockAuthorizationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthorization) EXPECT() *MockAuthorizationMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockAuthorization) CreateUser(user models.User) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", user)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockAuthorizationMockRecorder) CreateUser(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockAuthorization)(nil).CreateUser), user)
}

// GenerateToken mocks base method.
func (m *MockAuthorization) GenerateToken(email, password string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateToken", email, password)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateToken indicates an expected call of GenerateToken.
func (mr *MockAuthorizationMockRecorder) GenerateToken(email, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateToken", reflect.TypeOf((*MockAuthorization)(nil).GenerateToken), email, password)
}

// ParseToken mocks base method.
func (m *MockAuthorization) ParseToken(accessToken string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseToken", accessToken)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParseToken indicates an expected call of ParseToken.
func (mr *MockAuthorizationMockRecorder) ParseToken(accessToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseToken", reflect.TypeOf((*MockAuthorization)(nil).ParseToken), accessToken)
}

// MockImages is a mock of Images interface.
type MockImages struct {
	ctrl     *gomock.Controller
	recorder *MockImagesMockRecorder
}

// MockImagesMockRecorder is the mock recorder for MockImages.
type MockImagesMockRecorder struct {
	mock *MockImages
}

// NewMockImages creates a new mock instance.
func NewMockImages(ctrl *gomock.Controller) *MockImages {
	mock := &MockImages{ctrl: ctrl}
	mock.recorder = &MockImagesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockImages) EXPECT() *MockImagesMockRecorder {
	return m.recorder
}

// ConvertImage mocks base method.
func (m *MockImages) ConvertImage(sourceImage io.ReadSeeker, targetFormat string, ratio int) (io.ReadSeeker, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConvertImage", sourceImage, targetFormat, ratio)
	ret0, _ := ret[0].(io.ReadSeeker)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConvertImage indicates an expected call of ConvertImage.
func (mr *MockImagesMockRecorder) ConvertImage(sourceImage, targetFormat, ratio interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConvertImage", reflect.TypeOf((*MockImages)(nil).ConvertImage), sourceImage, targetFormat, ratio)
}

// GetImage mocks base method.
func (m *MockImages) GetImage(id string) (string, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetImage", id)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetImage indicates an expected call of GetImage.
func (mr *MockImagesMockRecorder) GetImage(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetImage", reflect.TypeOf((*MockImages)(nil).GetImage), id)
}

// GetImageID mocks base method.
func (m *MockImages) GetImageID(id string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetImageID", id)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetImageID indicates an expected call of GetImageID.
func (mr *MockImagesMockRecorder) GetImageID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetImageID", reflect.TypeOf((*MockImages)(nil).GetImageID), id)
}

// GetRequestFromID mocks base method.
func (m *MockImages) GetRequestFromID(userID int) ([]models.Request, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRequestFromID", userID)
	ret0, _ := ret[0].([]models.Request)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRequestFromID indicates an expected call of GetRequestFromID.
func (mr *MockImagesMockRecorder) GetRequestFromID(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRequestFromID", reflect.TypeOf((*MockImages)(nil).GetRequestFromID), userID)
}

// InsertImage mocks base method.
func (m *MockImages) InsertImage(filename, format string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertImage", filename, format)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertImage indicates an expected call of InsertImage.
func (mr *MockImagesMockRecorder) InsertImage(filename, format interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertImage", reflect.TypeOf((*MockImages)(nil).InsertImage), filename, format)
}

// RequestsHistory mocks base method.
func (m *MockImages) RequestsHistory(sourceFormat, targetFormat, imageID, filename string, userID, ratio int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RequestsHistory", sourceFormat, targetFormat, imageID, filename, userID, ratio)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RequestsHistory indicates an expected call of RequestsHistory.
func (mr *MockImagesMockRecorder) RequestsHistory(sourceFormat, targetFormat, imageID, filename, userID, ratio interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RequestsHistory", reflect.TypeOf((*MockImages)(nil).RequestsHistory), sourceFormat, targetFormat, imageID, filename, userID, ratio)
}

// UpdateRequest mocks base method.
func (m *MockImages) UpdateRequest(status, imageID, targetID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRequest", status, imageID, targetID)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateRequest indicates an expected call of UpdateRequest.
func (mr *MockImagesMockRecorder) UpdateRequest(status, imageID, targetID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRequest", reflect.TypeOf((*MockImages)(nil).UpdateRequest), status, imageID, targetID)
}

// MockServicesInterface is a mock of ServicesInterface interface.
type MockServicesInterface struct {
	ctrl     *gomock.Controller
	recorder *MockServicesInterfaceMockRecorder
}

// MockServicesInterfaceMockRecorder is the mock recorder for MockServicesInterface.
type MockServicesInterfaceMockRecorder struct {
	mock *MockServicesInterface
}

// NewMockServicesInterface creates a new mock instance.
func NewMockServicesInterface(ctrl *gomock.Controller) *MockServicesInterface {
	mock := &MockServicesInterface{ctrl: ctrl}
	mock.recorder = &MockServicesInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServicesInterface) EXPECT() *MockServicesInterfaceMockRecorder {
	return m.recorder
}

// ConvertImage mocks base method.
func (m *MockServicesInterface) ConvertImage(sourceImage io.ReadSeeker, targetFormat string, ratio int) (io.ReadSeeker, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConvertImage", sourceImage, targetFormat, ratio)
	ret0, _ := ret[0].(io.ReadSeeker)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConvertImage indicates an expected call of ConvertImage.
func (mr *MockServicesInterfaceMockRecorder) ConvertImage(sourceImage, targetFormat, ratio interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConvertImage", reflect.TypeOf((*MockServicesInterface)(nil).ConvertImage), sourceImage, targetFormat, ratio)
}

// CreateUser mocks base method.
func (m *MockServicesInterface) CreateUser(user models.User) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", user)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockServicesInterfaceMockRecorder) CreateUser(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockServicesInterface)(nil).CreateUser), user)
}

// GenerateToken mocks base method.
func (m *MockServicesInterface) GenerateToken(email, password string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateToken", email, password)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateToken indicates an expected call of GenerateToken.
func (mr *MockServicesInterfaceMockRecorder) GenerateToken(email, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateToken", reflect.TypeOf((*MockServicesInterface)(nil).GenerateToken), email, password)
}

// GetImage mocks base method.
func (m *MockServicesInterface) GetImage(id string) (string, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetImage", id)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetImage indicates an expected call of GetImage.
func (mr *MockServicesInterfaceMockRecorder) GetImage(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetImage", reflect.TypeOf((*MockServicesInterface)(nil).GetImage), id)
}

// GetImageID mocks base method.
func (m *MockServicesInterface) GetImageID(id string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetImageID", id)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetImageID indicates an expected call of GetImageID.
func (mr *MockServicesInterfaceMockRecorder) GetImageID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetImageID", reflect.TypeOf((*MockServicesInterface)(nil).GetImageID), id)
}

// GetRequestFromID mocks base method.
func (m *MockServicesInterface) GetRequestFromID(userID int) ([]models.Request, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRequestFromID", userID)
	ret0, _ := ret[0].([]models.Request)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRequestFromID indicates an expected call of GetRequestFromID.
func (mr *MockServicesInterfaceMockRecorder) GetRequestFromID(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRequestFromID", reflect.TypeOf((*MockServicesInterface)(nil).GetRequestFromID), userID)
}

// InsertImage mocks base method.
func (m *MockServicesInterface) InsertImage(filename, format string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertImage", filename, format)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertImage indicates an expected call of InsertImage.
func (mr *MockServicesInterfaceMockRecorder) InsertImage(filename, format interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertImage", reflect.TypeOf((*MockServicesInterface)(nil).InsertImage), filename, format)
}

// ParseToken mocks base method.
func (m *MockServicesInterface) ParseToken(accessToken string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseToken", accessToken)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParseToken indicates an expected call of ParseToken.
func (mr *MockServicesInterfaceMockRecorder) ParseToken(accessToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseToken", reflect.TypeOf((*MockServicesInterface)(nil).ParseToken), accessToken)
}

// RequestsHistory mocks base method.
func (m *MockServicesInterface) RequestsHistory(sourceFormat, targetFormat, imageID, filename string, userID, ratio int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RequestsHistory", sourceFormat, targetFormat, imageID, filename, userID, ratio)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RequestsHistory indicates an expected call of RequestsHistory.
func (mr *MockServicesInterfaceMockRecorder) RequestsHistory(sourceFormat, targetFormat, imageID, filename, userID, ratio interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RequestsHistory", reflect.TypeOf((*MockServicesInterface)(nil).RequestsHistory), sourceFormat, targetFormat, imageID, filename, userID, ratio)
}

// UpdateRequest mocks base method.
func (m *MockServicesInterface) UpdateRequest(status, imageID, targetID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRequest", status, imageID, targetID)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateRequest indicates an expected call of UpdateRequest.
func (mr *MockServicesInterfaceMockRecorder) UpdateRequest(status, imageID, targetID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRequest", reflect.TypeOf((*MockServicesInterface)(nil).UpdateRequest), status, imageID, targetID)
}
