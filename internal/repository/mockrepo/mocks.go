package mockrepo

import (
	context "context"
	reflect "reflect"

	models "github.com/Nikby53/image-converter/internal/models"
	repository "github.com/Nikby53/image-converter/internal/repository"
	gomock "github.com/golang/mock/gomock"
)

// MockAuthorizationRepository is a mock of AuthorizationRepository interface.
type MockAuthorizationRepository struct {
	ctrl     *gomock.Controller
	recorder *MockAuthorizationRepositoryMockRecorder
}

// MockAuthorizationRepositoryMockRecorder is the mock recorder for MockAuthorizationRepository.
type MockAuthorizationRepositoryMockRecorder struct {
	mock *MockAuthorizationRepository
}

// NewMockAuthorizationRepository creates a new mock instance.
func NewMockAuthorizationRepository(ctrl *gomock.Controller) *MockAuthorizationRepository {
	mock := &MockAuthorizationRepository{ctrl: ctrl}
	mock.recorder = &MockAuthorizationRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthorizationRepository) EXPECT() *MockAuthorizationRepositoryMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockAuthorizationRepository) CreateUser(ctx context.Context, user models.User) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, user)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockAuthorizationRepositoryMockRecorder) CreateUser(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockAuthorizationRepository)(nil).CreateUser), ctx, user)
}

// GetUser mocks base method.
func (m *MockAuthorizationRepository) GetUser(ctx context.Context, email string) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", ctx, email)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockAuthorizationRepositoryMockRecorder) GetUser(ctx, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockAuthorizationRepository)(nil).GetUser), ctx, email)
}

// MockImagesRepository is a mock of ImagesRepository interface.
type MockImagesRepository struct {
	ctrl     *gomock.Controller
	recorder *MockImagesRepositoryMockRecorder
}

// MockImagesRepositoryMockRecorder is the mock recorder for MockImagesRepository.
type MockImagesRepositoryMockRecorder struct {
	mock *MockImagesRepository
}

// NewMockImagesRepository creates a new mock instance.
func NewMockImagesRepository(ctrl *gomock.Controller) *MockImagesRepository {
	mock := &MockImagesRepository{ctrl: ctrl}
	mock.recorder = &MockImagesRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockImagesRepository) EXPECT() *MockImagesRepositoryMockRecorder {
	return m.recorder
}

// GetImageByID mocks base method.
func (m *MockImagesRepository) GetImageByID(ctx context.Context, id string) (models.Images, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetImageByID", ctx, id)
	ret0, _ := ret[0].(models.Images)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetImageByID indicates an expected call of GetImageByID.
func (mr *MockImagesRepositoryMockRecorder) GetImageByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetImageByID", reflect.TypeOf((*MockImagesRepository)(nil).GetImageByID), ctx, id)
}

// GetRequestFromID mocks base method.
func (m *MockImagesRepository) GetRequestFromID(ctx context.Context, userID int) ([]models.Request, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRequestFromID", ctx, userID)
	ret0, _ := ret[0].([]models.Request)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRequestFromID indicates an expected call of GetRequestFromID.
func (mr *MockImagesRepositoryMockRecorder) GetRequestFromID(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRequestFromID", reflect.TypeOf((*MockImagesRepository)(nil).GetRequestFromID), ctx, userID)
}

// InsertImage mocks base method.
func (m *MockImagesRepository) InsertImage(ctx context.Context, filename, format string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertImage", ctx, filename, format)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertImage indicates an expected call of InsertImage.
func (mr *MockImagesRepositoryMockRecorder) InsertImage(ctx, filename, format interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertImage", reflect.TypeOf((*MockImagesRepository)(nil).InsertImage), ctx, filename, format)
}

// RequestsHistory mocks base method.
func (m *MockImagesRepository) RequestsHistory(ctx context.Context, sourceFormat, targetFormat, imageID, filename string, userID, ratio int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RequestsHistory", ctx, sourceFormat, targetFormat, imageID, filename, userID, ratio)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RequestsHistory indicates an expected call of RequestsHistory.
func (mr *MockImagesRepositoryMockRecorder) RequestsHistory(ctx, sourceFormat, targetFormat, imageID, filename, userID, ratio interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RequestsHistory", reflect.TypeOf((*MockImagesRepository)(nil).RequestsHistory), ctx, sourceFormat, targetFormat, imageID, filename, userID, ratio)
}

// UpdateRequest mocks base method.
func (m *MockImagesRepository) UpdateRequest(ctx context.Context, status, imageID, targetID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRequest", ctx, status, imageID, targetID)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateRequest indicates an expected call of UpdateRequest.
func (mr *MockImagesRepositoryMockRecorder) UpdateRequest(ctx, status, imageID, targetID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRequest", reflect.TypeOf((*MockImagesRepository)(nil).UpdateRequest), ctx, status, imageID, targetID)
}

// MockRepoInterface is a mock of RepoInterface interface.
type MockRepoInterface struct {
	ctrl     *gomock.Controller
	recorder *MockRepoInterfaceMockRecorder
}

// MockRepoInterfaceMockRecorder is the mock recorder for MockRepoInterface.
type MockRepoInterfaceMockRecorder struct {
	mock *MockRepoInterface
}

// NewMockRepoInterface creates a new mock instance.
func NewMockRepoInterface(ctrl *gomock.Controller) *MockRepoInterface {
	mock := &MockRepoInterface{ctrl: ctrl}
	mock.recorder = &MockRepoInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepoInterface) EXPECT() *MockRepoInterfaceMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockRepoInterface) CreateUser(ctx context.Context, user models.User) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, user)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockRepoInterfaceMockRecorder) CreateUser(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockRepoInterface)(nil).CreateUser), ctx, user)
}

// GetImageByID mocks base method.
func (m *MockRepoInterface) GetImageByID(ctx context.Context, id string) (models.Images, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetImageByID", ctx, id)
	ret0, _ := ret[0].(models.Images)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetImageByID indicates an expected call of GetImageByID.
func (mr *MockRepoInterfaceMockRecorder) GetImageByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetImageByID", reflect.TypeOf((*MockRepoInterface)(nil).GetImageByID), ctx, id)
}

// GetRequestFromID mocks base method.
func (m *MockRepoInterface) GetRequestFromID(ctx context.Context, userID int) ([]models.Request, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRequestFromID", ctx, userID)
	ret0, _ := ret[0].([]models.Request)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRequestFromID indicates an expected call of GetRequestFromID.
func (mr *MockRepoInterfaceMockRecorder) GetRequestFromID(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRequestFromID", reflect.TypeOf((*MockRepoInterface)(nil).GetRequestFromID), ctx, userID)
}

// GetUser mocks base method.
func (m *MockRepoInterface) GetUser(ctx context.Context, email string) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", ctx, email)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockRepoInterfaceMockRecorder) GetUser(ctx, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockRepoInterface)(nil).GetUser), ctx, email)
}

// InsertImage mocks base method.
func (m *MockRepoInterface) InsertImage(ctx context.Context, filename, format string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertImage", ctx, filename, format)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertImage indicates an expected call of InsertImage.
func (mr *MockRepoInterfaceMockRecorder) InsertImage(ctx, filename, format interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertImage", reflect.TypeOf((*MockRepoInterface)(nil).InsertImage), ctx, filename, format)
}

// RequestsHistory mocks base method.
func (m *MockRepoInterface) RequestsHistory(ctx context.Context, sourceFormat, targetFormat, imageID, filename string, userID, ratio int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RequestsHistory", ctx, sourceFormat, targetFormat, imageID, filename, userID, ratio)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RequestsHistory indicates an expected call of RequestsHistory.
func (mr *MockRepoInterfaceMockRecorder) RequestsHistory(ctx, sourceFormat, targetFormat, imageID, filename, userID, ratio interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RequestsHistory", reflect.TypeOf((*MockRepoInterface)(nil).RequestsHistory), ctx, sourceFormat, targetFormat, imageID, filename, userID, ratio)
}

// Transactional mocks base method.
func (m *MockRepoInterface) Transactional(f func(repository.RepoInterface) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Transactional", f)
	ret0, _ := ret[0].(error)
	return ret0
}

// Transactional indicates an expected call of Transactional.
func (mr *MockRepoInterfaceMockRecorder) Transactional(f interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Transactional", reflect.TypeOf((*MockRepoInterface)(nil).Transactional), f)
}

// UpdateRequest mocks base method.
func (m *MockRepoInterface) UpdateRequest(ctx context.Context, status, imageID, targetID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRequest", ctx, status, imageID, targetID)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateRequest indicates an expected call of UpdateRequest.
func (mr *MockRepoInterfaceMockRecorder) UpdateRequest(ctx, status, imageID, targetID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRequest", reflect.TypeOf((*MockRepoInterface)(nil).UpdateRequest), ctx, status, imageID, targetID)
}
