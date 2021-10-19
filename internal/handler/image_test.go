package handler

import (
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Nikby53/image-converter/internal/storage/mocksstorage"

	"github.com/Nikby53/image-converter/internal/models"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"

	"github.com/Nikby53/image-converter/internal/service/mocks"
)

func TestHandler_requests(t *testing.T) {
	request := []models.Request{{
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
	requestJSON, err := json.MarshalIndent(request, "\t", "")
	if err != nil {
		t.Fatalf("can't marshal: %v", err)
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
				r.EXPECT().GetRequestFromID(0).Return(request, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: string(requestJSON),
		},
		{
			name: "Repo error",
			mockBehavior: func(r *mocks.MockServicesInterface) {
				r.EXPECT().GetRequestFromID(0).Return(nil, fmt.Errorf(""))
			},
			expectedStatusCode:   500,
			expectedResponseBody: "repository error \n",
		},
		{
			name: "Can't get id from token",
			mockBehavior: func(r *mocks.MockServicesInterface) {
				r.EXPECT().GetRequestFromID(0).Return(nil, fmt.Errorf("can't get id from jwt token"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: "repository error can't get id from jwt token\n",
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

func TestHandler_downloadImage(t *testing.T) {
	imageModel := models.Images{
		ID:     "1",
		Format: "png",
		Name:   "image",
	}
	type mockBehavior func(r *mocksstorage.MockStorageInterface, w *mocks.MockServicesInterface, server Server)
	tests := []struct {
		name                 string
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name: "Ok",
			mockBehavior: func(r *mocksstorage.MockStorageInterface, w *mocks.MockServicesInterface, server Server) {
				w.EXPECT().GetImageByID(gomock.Any()).Return(imageModel, nil)
				r.EXPECT().DownloadImageFromID(imageModel.ID).Return("https://images-convert.s3.eu-central-1.amazonaws.com/5?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAVU5FA7EW3LEOGAOA%2F20211019%2Feu-central-1%2Fs3%2Faws4_request&X-Amz-Date=20211019T112019Z&X-Amz-Expires=600&X-Amz-SignedHeaders=host&X-Amz-Signature=599a2445e2e9b37a39c531b125df5216975546aa4c9156f25cc0fb30bda89ea9", nil)
				server.ParseUrl("https://images-convert.s3.eu-central-1.amazonaws.com/5?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAVU5FA7EW3LEOGAOA%2F20211019%2Feu-central-1%2Fs3%2Faws4_request&X-Amz-Date=20211019T112019Z&X-Amz-Expires=600&X-Amz-SignedHeaders=host&X-Amz-Signature=599a2445e2e9b37a39c531b125df5216975546aa4c9156f25cc0fb30bda89ea9")
			},
			expectedStatusCode:   200,
			expectedResponseBody: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			services := mocks.NewMockServicesInterface(c)
			st := mocksstorage.NewMockStorageInterface(c)
			server := Server{}
			tt.mockBehavior(st, services, server)
			broker := Server{messageBroker: nil}
			handler := NewServer(services, st, broker.messageBroker)
			r := mux.NewRouter()
			r.HandleFunc("/image/downloadImage", handler.downloadImage).Methods("GET")
			w := httptest.NewRecorder()
			w.Header().Set("Content-Disposition", "attachment; filename="+imageModel.Name+"."+imageModel.Format)
			req := httptest.NewRequest("GET", "/image/downloadImage", nil)
			r.ServeHTTP(w, req)
			assert.Equal(t, tt.expectedStatusCode, w.Code)
			assert.Equal(t, tt.expectedResponseBody, w.Body.String())
		})
	}
}
