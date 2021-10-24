package handler

import (
	"bytes"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/Nikby53/image-converter/internal/models"
	"github.com/Nikby53/image-converter/internal/service/mocks"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestHandler_signUp(t *testing.T) {
	type mockBehavior func(r *mocks.MockServicesInterface, user models.User)

	tests := []struct {
		name                 string
		inputBody            string
		inputUser            models.User
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "Ok",
			inputBody: `{"email": "email@mail.ru", "password": "qwertyuiop"}`,
			inputUser: models.User{
				Email:    "email@mail.ru",
				Password: "qwertyuiop",
			},
			mockBehavior: func(r *mocks.MockServicesInterface, user models.User) {
				r.EXPECT().CreateUser(user).Return(1, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: "{\"id\":1}\n",
		},
		{
			name:      "Empty password",
			inputBody: `{"email": "email@email.com", "password": ""}`,
			inputUser: models.User{
				Email:    "email@email.com",
				Password: "",
			},
			mockBehavior:         func(r *mocks.MockServicesInterface, user models.User) {},
			expectedStatusCode:   400,
			expectedResponseBody: "password should be not empty\n",
		},
		{
			name:      "Empty email",
			inputBody: `{"email": "", "password": "12312312441"}`,
			inputUser: models.User{
				Email:    "",
				Password: "12312312441",
			},
			mockBehavior:         func(r *mocks.MockServicesInterface, user models.User) {},
			expectedStatusCode:   400,
			expectedResponseBody: "email should be not empty\n",
		},
		{
			name:      "Invalid email",
			inputBody: `{"email": "retwerwe", "password": "12312312441"}`,
			inputUser: models.User{
				Email:    "retwerwe",
				Password: "12312312441",
			},
			mockBehavior:         func(r *mocks.MockServicesInterface, user models.User) {},
			expectedStatusCode:   400,
			expectedResponseBody: "invalid email\n",
		},
		{
			name:      "Invalid email",
			inputBody: `{"email": "retwerwe", "password": "12312312441"}`,
			inputUser: models.User{
				Email:    "retwe.rwe@gmail.com",
				Password: "12312312441",
			},
			mockBehavior:         func(r *mocks.MockServicesInterface, user models.User) {},
			expectedStatusCode:   400,
			expectedResponseBody: "invalid email\n",
		},
		{
			name:      "Similar user",
			inputBody: `{"email": "email@mail.ru", "password": "qwertyuiop"}`,
			inputUser: models.User{
				Email:    "email@mail.ru",
				Password: "qwertyuiop",
			},
			mockBehavior: func(r *mocks.MockServicesInterface, user models.User) {
				r.EXPECT().CreateUser(user).Return(0, fmt.Errorf("A similar user is already registered in the system"))
			},
			expectedStatusCode:   409,
			expectedResponseBody: "A similar user is already registered in the system\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			services := mocks.NewMockServicesInterface(c)
			tt.mockBehavior(services, tt.inputUser)
			storage := Server{storage: nil}
			broker := Server{messageBroker: nil}
			handler := NewServer(services, storage.storage, broker.messageBroker)
			r := mux.NewRouter()
			r.HandleFunc("/user/signup", handler.signUp).Methods("POST")
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/user/signup", bytes.NewBufferString(tt.inputBody))
			r.ServeHTTP(w, req)
			assert.Equal(t, tt.expectedStatusCode, w.Code)
			assert.Equal(t, tt.expectedResponseBody, w.Body.String())
		})
	}
}

func TestHandler_login(t *testing.T) {
	type mockBehavior func(r *mocks.MockServicesInterface, user models.User)

	tests := []struct {
		name                 string
		inputBody            string
		inputUser            models.User
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "Ok",
			inputBody: `{"email": "12sdrfsdf3121", "password": "12332ferfwf1"}`,
			inputUser: models.User{
				Email:    "12sdrfsdf3121",
				Password: "12332ferfwf1",
			},
			mockBehavior: func(r *mocks.MockServicesInterface, user models.User) {
				r.EXPECT().GenerateToken(user.Email, user.Password)
			},
			expectedStatusCode:   200,
			expectedResponseBody: "{\"token\":\"\"}\n",
		},
		{
			name:      "Empty password",
			inputBody: `{"email": "12sdrfsdf3121@gmail.com", "password": ""}`,
			inputUser: models.User{
				Email:    "12sdrfsdf3121@gmail.com",
				Password: "",
			},
			mockBehavior:         func(r *mocks.MockServicesInterface, user models.User) {},
			expectedStatusCode:   400,
			expectedResponseBody: "password should be not empty\n",
		},
		{
			name:      "Empty email",
			inputBody: `{"email": "", "password": "1233123123"}`,
			inputUser: models.User{
				Email:    "",
				Password: "1233123123",
			},
			mockBehavior:         func(r *mocks.MockServicesInterface, user models.User) {},
			expectedStatusCode:   400,
			expectedResponseBody: "email should be not empty\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			services := mocks.NewMockServicesInterface(c)
			tt.mockBehavior(services, tt.inputUser)
			storage := Server{storage: nil}
			broker := Server{messageBroker: nil}
			handler := NewServer(services, storage.storage, broker.messageBroker)
			r := mux.NewRouter()
			r.HandleFunc("/user/login", handler.login).Methods("POST")
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/user/login", bytes.NewBufferString(tt.inputBody))
			r.ServeHTTP(w, req)
			assert.Equal(t, tt.expectedStatusCode, w.Code)
			assert.Equal(t, tt.expectedResponseBody, w.Body.String())
		})
	}
}