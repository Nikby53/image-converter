package handler

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Nikby53/image-converter/internal/service/mocks"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func testHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("1"))
}

func TestHandler_userIdentity(t *testing.T) {
	type mockBehavior func(r *mocks.MockServicesInterface, token string)

	testTable := []struct {
		name                 string
		headerName           string
		headerValue          string
		token                string
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:        "Ok",
			headerName:  "Authorization",
			headerValue: "Bearer token",
			token:       "token",
			mockBehavior: func(r *mocks.MockServicesInterface, token string) {
				r.EXPECT().ParseToken(token).Return(1, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: "1",
		},
		{
			name:                 "Invalid Header Name",
			headerName:           "",
			headerValue:          "Bearer token",
			token:                "token",
			mockBehavior:         func(r *mocks.MockServicesInterface, token string) {},
			expectedStatusCode:   401,
			expectedResponseBody: "empty authorization handler\n",
		},
		{
			name:                 "Invalid Header Value",
			headerName:           "Authorization",
			headerValue:          "Bearr token",
			token:                "token",
			mockBehavior:         func(r *mocks.MockServicesInterface, token string) {},
			expectedStatusCode:   401,
			expectedResponseBody: "invalid auth header\n",
		},
		{
			name:                 "Empty Token",
			headerName:           "Authorization",
			headerValue:          "Bearer ",
			token:                "token",
			mockBehavior:         func(r *mocks.MockServicesInterface, token string) {},
			expectedStatusCode:   401,
			expectedResponseBody: "token is empty\n",
		},
		{
			name:        "Parse Error",
			headerName:  "Authorization",
			headerValue: "Bearer token",
			token:       "token",
			mockBehavior: func(r *mocks.MockServicesInterface, token string) {
				r.EXPECT().ParseToken(token).Return(0, errors.New("can't parse jwt token\n"))
			},
			expectedStatusCode:   401,
			expectedResponseBody: "can't parse jwt token\n",
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			services := mocks.NewMockServicesInterface(c)
			test.mockBehavior(services, test.token)
			s := &Server{services: services}
			r := mux.NewRouter()
			r.Use(s.userIdentity)
			r.HandleFunc("/identity", testHandler).Methods("GET")

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/identity", nil)
			req.Header.Set(test.headerName, test.headerValue)

			r.ServeHTTP(w, req)

			assert.Equal(t, test.expectedStatusCode, w.Code)
			assert.Equal(t, test.expectedResponseBody, w.Body.String())
		})
	}
}
