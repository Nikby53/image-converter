package handler

import (
	"bytes"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/gorilla/mux"

	"github.com/stretchr/testify/assert"

	"github.com/Nikby53/image-converter/internal/models"
	"github.com/Nikby53/image-converter/internal/service/mocks"
)

func TestHandler_signUp(t *testing.T) {
	type mockBehavior func(r *mocks.MockServiceInterface, user models.User)

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
			inputBody: `{"email": "email", "password": "qwerty"}`,
			inputUser: models.User{
				Email:    "email",
				Password: "qwerty",
			},
			mockBehavior: func(r *mocks.MockServiceInterface, user models.User) {
				r.EXPECT().CreateUser(user).Return(1, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: "{\"id\":1}\n",
		},
		{
			name:      "Wrong Input",
			inputBody: `{"email": "email", "password": ""}`,
			inputUser: models.User{
				Email:    "email",
				Password: "",
			},
			mockBehavior:         func(r *mocks.MockServiceInterface, user models.User) {},
			expectedStatusCode:   400,
			expectedResponseBody: "password should be not empty\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			services := mocks.NewMockServiceInterface(c)
			tt.mockBehavior(services, tt.inputUser)
			handler := NewServer(services)
			r := mux.NewRouter()
			r.HandleFunc("/sign-up", handler.signUp).Methods("POST")
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/sign-up", bytes.NewBufferString(tt.inputBody))
			r.ServeHTTP(w, req)
			assert.Equal(t, tt.expectedStatusCode, w.Code)
			assert.Equal(t, tt.expectedResponseBody, w.Body.String())
		})
	}
}
