package handler

import (
	"encoding/json"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Nikby53/image-converter/internal/models"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"

	"github.com/Nikby53/image-converter/internal/service/mocks"
)

func TestHandler_requests(t *testing.T) {
	req := []models.Request{{
		Filename:      "img",
		Status:        "done",
		SourceFormat:  "png",
		TargetFormat:  "jpg",
		Ratio:         99,
		Created:       time.Time{},
		Updated:       time.Time{},
		OriginalImgID: "4",
		TargetImgID:   "5",
	},
		{
			Filename:      "image",
			Status:        "done",
			SourceFormat:  "jpg",
			TargetFormat:  "png",
			Ratio:         56,
			Created:       time.Time{},
			Updated:       time.Time{},
			OriginalImgID: "7",
			TargetImgID:   "8",
		},
	}
	reqJSON, err := json.MarshalIndent(req, "\t", "")
	if err != nil {
		t.Fatalf("can't marshal response body: %v", err)
	}
	type mockBehavior func(r *mocks.MockServicesInterface)
	tests := []struct {
		name                 string
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name: "Ok",
			mockBehavior: func(r *mocks.MockServicesInterface) {
				r.EXPECT().GetRequestFromID(2).Return(req, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: string(reqJSON),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			services := mocks.NewMockServicesInterface(c)
			tt.mockBehavior(services)
			storage := Server{storage: nil}
			broker := Server{messageBroker: nil}
			handler := NewServer(services, storage.storage, broker.messageBroker)
			r := mux.NewRouter()
			r.HandleFunc("/image/requests", handler.requests).Methods("GET")
			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/image/requests", nil)
			r.ServeHTTP(w, req)
			assert.Equal(t, tt.expectedStatusCode, w.Code)
			assert.Equal(t, tt.expectedResponseBody, w.Body.String())
		})
	}
}
