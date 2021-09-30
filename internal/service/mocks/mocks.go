package mocks

import (
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
func (m *MockImages) ConvertImage(imageBytes []byte, targetFormat string, ratio int) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConvertImage", imageBytes, targetFormat, ratio)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConvertImage indicates an expected call of ConvertImage.
func (mr *MockImagesMockRecorder) ConvertImage(imageBytes, targetFormat, ratio interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConvertImage", reflect.TypeOf((*MockImages)(nil).ConvertImage), imageBytes, targetFormat, ratio)
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

// GetRequestFromId mocks base method.
func (m *MockImages) GetRequestFromId(userID int) ([]models.Request, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRequestFromId", userID)
	ret0, _ := ret[0].([]models.Request)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRequestFromId indicates an expected call of GetRequestFromId.
func (mr *MockImagesMockRecorder) GetRequestFromId(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRequestFromId", reflect.TypeOf((*MockImages)(nil).GetRequestFromId), userID)
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
func (m *MockImages) RequestsHistory(sourceFormat, targetFormat, imagesId, filename string, userId, ratio int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RequestsHistory", sourceFormat, targetFormat, imagesId, filename, userId, ratio)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RequestsHistory indicates an expected call of RequestsHistory.
func (mr *MockImagesMockRecorder) RequestsHistory(sourceFormat, targetFormat, imagesId, filename, userId, ratio interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RequestsHistory", reflect.TypeOf((*MockImages)(nil).RequestsHistory), sourceFormat, targetFormat, imagesId, filename, userId, ratio)
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

// MockServiceInterface is a mock of ServiceInterface interface.
type MockServiceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockServiceInterfaceMockRecorder
}

// MockServiceInterfaceMockRecorder is the mock recorder for MockServiceInterface.
type MockServiceInterfaceMockRecorder struct {
	mock *MockServiceInterface
}

// NewMockServiceInterface creates a new mock instance.
func NewMockServiceInterface(ctrl *gomock.Controller) *MockServiceInterface {
	mock := &MockServiceInterface{ctrl: ctrl}
	mock.recorder = &MockServiceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceInterface) EXPECT() *MockServiceInterfaceMockRecorder {
	return m.recorder
}

// ConvertImage mocks base method.
func (m *MockServiceInterface) ConvertImage(imageBytes []byte, targetFormat string, ratio int) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConvertImage", imageBytes, targetFormat, ratio)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConvertImage indicates an expected call of ConvertImage.
func (mr *MockServiceInterfaceMockRecorder) ConvertImage(imageBytes, targetFormat, ratio interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConvertImage", reflect.TypeOf((*MockServiceInterface)(nil).ConvertImage), imageBytes, targetFormat, ratio)
}

// CreateUser mocks base method.
func (m *MockServiceInterface) CreateUser(user models.User) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", user)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockServiceInterfaceMockRecorder) CreateUser(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockServiceInterface)(nil).CreateUser), user)
}

// GenerateToken mocks base method.
func (m *MockServiceInterface) GenerateToken(email, password string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateToken", email, password)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateToken indicates an expected call of GenerateToken.
func (mr *MockServiceInterfaceMockRecorder) GenerateToken(email, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateToken", reflect.TypeOf((*MockServiceInterface)(nil).GenerateToken), email, password)
}

// GetImageID mocks base method.
func (m *MockServiceInterface) GetImageID(id string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetImageID", id)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetImageID indicates an expected call of GetImageID.
func (mr *MockServiceInterfaceMockRecorder) GetImageID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetImageID", reflect.TypeOf((*MockServiceInterface)(nil).GetImageID), id)
}

// GetRequestFromId mocks base method.
func (m *MockServiceInterface) GetRequestFromId(userID int) ([]models.Request, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRequestFromId", userID)
	ret0, _ := ret[0].([]models.Request)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRequestFromId indicates an expected call of GetRequestFromId.
func (mr *MockServiceInterfaceMockRecorder) GetRequestFromId(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRequestFromId", reflect.TypeOf((*MockServiceInterface)(nil).GetRequestFromId), userID)
}

// InsertImage mocks base method.
func (m *MockServiceInterface) InsertImage(filename, format string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertImage", filename, format)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertImage indicates an expected call of InsertImage.
func (mr *MockServiceInterfaceMockRecorder) InsertImage(filename, format interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertImage", reflect.TypeOf((*MockServiceInterface)(nil).InsertImage), filename, format)
}

// ParseToken mocks base method.
func (m *MockServiceInterface) ParseToken(accessToken string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseToken", accessToken)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParseToken indicates an expected call of ParseToken.
func (mr *MockServiceInterfaceMockRecorder) ParseToken(accessToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseToken", reflect.TypeOf((*MockServiceInterface)(nil).ParseToken), accessToken)
}

// RequestsHistory mocks base method.
func (m *MockServiceInterface) RequestsHistory(sourceFormat, targetFormat, imagesId, filename string, userId, ratio int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RequestsHistory", sourceFormat, targetFormat, imagesId, filename, userId, ratio)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RequestsHistory indicates an expected call of RequestsHistory.
func (mr *MockServiceInterfaceMockRecorder) RequestsHistory(sourceFormat, targetFormat, imagesId, filename, userId, ratio interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RequestsHistory", reflect.TypeOf((*MockServiceInterface)(nil).RequestsHistory), sourceFormat, targetFormat, imagesId, filename, userId, ratio)
}

// UpdateRequest mocks base method.
func (m *MockServiceInterface) UpdateRequest(status, imageID, targetID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRequest", status, imageID, targetID)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateRequest indicates an expected call of UpdateRequest.
func (mr *MockServiceInterfaceMockRecorder) UpdateRequest(status, imageID, targetID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRequest", reflect.TypeOf((*MockServiceInterface)(nil).UpdateRequest), status, imageID, targetID)
}
