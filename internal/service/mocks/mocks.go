package mocks

import (
	context "context"
	reflect "reflect"

	models "github.com/Nikby53/image-converter/internal/models"
	service "github.com/Nikby53/image-converter/internal/service"
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
func (m *MockAuthorization) CreateUser(ctx context.Context, user models.User) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, user)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockAuthorizationMockRecorder) CreateUser(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockAuthorization)(nil).CreateUser), ctx, user)
}

// GenerateToken mocks base method.
func (m *MockAuthorization) GenerateToken(ctx context.Context, email, password string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateToken", ctx, email, password)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateToken indicates an expected call of GenerateToken.
func (mr *MockAuthorizationMockRecorder) GenerateToken(ctx, email, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateToken", reflect.TypeOf((*MockAuthorization)(nil).GenerateToken), ctx, email, password)
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

// Conversion mocks base method.
func (m *MockImages) Conversion(ctx context.Context, payload service.ConversionPayLoad) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Conversion", ctx, payload)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Conversion indicates an expected call of Conversion.
func (mr *MockImagesMockRecorder) Conversion(ctx, payload interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Conversion", reflect.TypeOf((*MockImages)(nil).Conversion), ctx, payload)
}

// DownloadImageFromID mocks base method.
func (m *MockImages) DownloadImageFromID(fileID string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DownloadImageFromID", fileID)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DownloadImageFromID indicates an expected call of DownloadImageFromID.
func (mr *MockImagesMockRecorder) DownloadImageFromID(fileID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DownloadImageFromID", reflect.TypeOf((*MockImages)(nil).DownloadImageFromID), fileID)
}

// GetImageByID mocks base method.
func (m *MockImages) GetImageByID(ctx context.Context, id string) (models.Images, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetImageByID", ctx, id)
	ret0, _ := ret[0].(models.Images)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetImageByID indicates an expected call of GetImageByID.
func (mr *MockImagesMockRecorder) GetImageByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetImageByID", reflect.TypeOf((*MockImages)(nil).GetImageByID), ctx, id)
}

// GetRequestFromID mocks base method.
func (m *MockImages) GetRequestFromID(ctx context.Context, userID int) ([]models.Request, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRequestFromID", ctx, userID)
	ret0, _ := ret[0].([]models.Request)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRequestFromID indicates an expected call of GetRequestFromID.
func (mr *MockImagesMockRecorder) GetRequestFromID(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRequestFromID", reflect.TypeOf((*MockImages)(nil).GetRequestFromID), ctx, userID)
}

// InsertImage mocks base method.
func (m *MockImages) InsertImage(ctx context.Context, filename, format string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertImage", ctx, filename, format)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertImage indicates an expected call of InsertImage.
func (mr *MockImagesMockRecorder) InsertImage(ctx, filename, format interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertImage", reflect.TypeOf((*MockImages)(nil).InsertImage), ctx, filename, format)
}

// RequestsHistory mocks base method.
func (m *MockImages) RequestsHistory(ctx context.Context, sourceFormat, targetFormat, imageID, filename string, userID, ratio int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RequestsHistory", ctx, sourceFormat, targetFormat, imageID, filename, userID, ratio)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RequestsHistory indicates an expected call of RequestsHistory.
func (mr *MockImagesMockRecorder) RequestsHistory(ctx, sourceFormat, targetFormat, imageID, filename, userID, ratio interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RequestsHistory", reflect.TypeOf((*MockImages)(nil).RequestsHistory), ctx, sourceFormat, targetFormat, imageID, filename, userID, ratio)
}

// UpdateRequest mocks base method.
func (m *MockImages) UpdateRequest(ctx context.Context, status, imageID, targetID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRequest", ctx, status, imageID, targetID)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateRequest indicates an expected call of UpdateRequest.
func (mr *MockImagesMockRecorder) UpdateRequest(ctx, status, imageID, targetID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRequest", reflect.TypeOf((*MockImages)(nil).UpdateRequest), ctx, status, imageID, targetID)
}

// MockInterface is a mock of Interface interface.
type MockInterface struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceMockRecorder
}

// MockInterfaceMockRecorder is the mock recorder for MockInterface.
type MockInterfaceMockRecorder struct {
	mock *MockInterface
}

// NewMockInterface creates a new mock instance.
func NewMockInterface(ctrl *gomock.Controller) *MockInterface {
	mock := &MockInterface{ctrl: ctrl}
	mock.recorder = &MockInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInterface) EXPECT() *MockInterfaceMockRecorder {
	return m.recorder
}

// Conversion mocks base method.
func (m *MockInterface) Conversion(ctx context.Context, payload service.ConversionPayLoad) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Conversion", ctx, payload)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Conversion indicates an expected call of Conversion.
func (mr *MockInterfaceMockRecorder) Conversion(ctx, payload interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Conversion", reflect.TypeOf((*MockInterface)(nil).Conversion), ctx, payload)
}

// CreateUser mocks base method.
func (m *MockInterface) CreateUser(ctx context.Context, user models.User) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, user)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockInterfaceMockRecorder) CreateUser(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockInterface)(nil).CreateUser), ctx, user)
}

// DownloadImageFromID mocks base method.
func (m *MockInterface) DownloadImageFromID(fileID string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DownloadImageFromID", fileID)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DownloadImageFromID indicates an expected call of DownloadImageFromID.
func (mr *MockInterfaceMockRecorder) DownloadImageFromID(fileID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DownloadImageFromID", reflect.TypeOf((*MockInterface)(nil).DownloadImageFromID), fileID)
}

// GenerateToken mocks base method.
func (m *MockInterface) GenerateToken(ctx context.Context, email, password string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateToken", ctx, email, password)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateToken indicates an expected call of GenerateToken.
func (mr *MockInterfaceMockRecorder) GenerateToken(ctx, email, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateToken", reflect.TypeOf((*MockInterface)(nil).GenerateToken), ctx, email, password)
}

// GetImageByID mocks base method.
func (m *MockInterface) GetImageByID(ctx context.Context, id string) (models.Images, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetImageByID", ctx, id)
	ret0, _ := ret[0].(models.Images)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetImageByID indicates an expected call of GetImageByID.
func (mr *MockInterfaceMockRecorder) GetImageByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetImageByID", reflect.TypeOf((*MockInterface)(nil).GetImageByID), ctx, id)
}

// GetRequestFromID mocks base method.
func (m *MockInterface) GetRequestFromID(ctx context.Context, userID int) ([]models.Request, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRequestFromID", ctx, userID)
	ret0, _ := ret[0].([]models.Request)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRequestFromID indicates an expected call of GetRequestFromID.
func (mr *MockInterfaceMockRecorder) GetRequestFromID(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRequestFromID", reflect.TypeOf((*MockInterface)(nil).GetRequestFromID), ctx, userID)
}

// InsertImage mocks base method.
func (m *MockInterface) InsertImage(ctx context.Context, filename, format string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertImage", ctx, filename, format)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertImage indicates an expected call of InsertImage.
func (mr *MockInterfaceMockRecorder) InsertImage(ctx, filename, format interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertImage", reflect.TypeOf((*MockInterface)(nil).InsertImage), ctx, filename, format)
}

// ParseToken mocks base method.
func (m *MockInterface) ParseToken(accessToken string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseToken", accessToken)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParseToken indicates an expected call of ParseToken.
func (mr *MockInterfaceMockRecorder) ParseToken(accessToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseToken", reflect.TypeOf((*MockInterface)(nil).ParseToken), accessToken)
}

// RequestsHistory mocks base method.
func (m *MockInterface) RequestsHistory(ctx context.Context, sourceFormat, targetFormat, imageID, filename string, userID, ratio int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RequestsHistory", ctx, sourceFormat, targetFormat, imageID, filename, userID, ratio)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RequestsHistory indicates an expected call of RequestsHistory.
func (mr *MockInterfaceMockRecorder) RequestsHistory(ctx, sourceFormat, targetFormat, imageID, filename, userID, ratio interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RequestsHistory", reflect.TypeOf((*MockInterface)(nil).RequestsHistory), ctx, sourceFormat, targetFormat, imageID, filename, userID, ratio)
}

// UpdateRequest mocks base method.
func (m *MockInterface) UpdateRequest(ctx context.Context, status, imageID, targetID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRequest", ctx, status, imageID, targetID)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateRequest indicates an expected call of UpdateRequest.
func (mr *MockInterfaceMockRecorder) UpdateRequest(ctx, status, imageID, targetID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRequest", reflect.TypeOf((*MockInterface)(nil).UpdateRequest), ctx, status, imageID, targetID)
}
