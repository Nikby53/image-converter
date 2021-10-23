package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/Nikby53/image-converter/internal/service"

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
				r.EXPECT().GetRequestFromID(gomock.Any()).Return(nil, fmt.Errorf("can't get id from jwt token"))
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
	type mockBehavior func(r *mocksstorage.MockStorageInterface, w *mocks.MockServicesInterface)
	tests := []struct {
		name                      string
		mockBehavior              mockBehavior
		checkHeaders              bool
		expectedStatusCode        int
		expectedResponseBody      string
		expectedDispositionHeader string
		expectedConTypeHeader     string
	}{
		{
			name: "Ok",
			mockBehavior: func(r *mocksstorage.MockStorageInterface, w *mocks.MockServicesInterface) {
				w.EXPECT().GetImageByID(gomock.Any()).Return(imageModel, nil)
				r.EXPECT().DownloadImageFromID(imageModel.ID).Return("https://images-convert.s3.eu-central-1.amazonaws.com/image.png", nil)
			},
			expectedStatusCode:        200,
			checkHeaders:              true,
			expectedDispositionHeader: "attachment; filename=" + imageModel.Name + "." + imageModel.Format,
			expectedConTypeHeader:     imageModel.Format,
		},
		{
			name: "Repository error",
			mockBehavior: func(r *mocksstorage.MockStorageInterface, w *mocks.MockServicesInterface) {
				w.EXPECT().GetImageByID(gomock.Any()).Return(imageModel, fmt.Errorf(""))
			},
			expectedStatusCode:   500,
			expectedResponseBody: "repository error, \n",
			checkHeaders:         false,
		},
		{
			name: "Can't download image",
			mockBehavior: func(r *mocksstorage.MockStorageInterface, w *mocks.MockServicesInterface) {
				w.EXPECT().GetImageByID(gomock.Any()).Return(imageModel, nil)
				r.EXPECT().DownloadImageFromID(imageModel.ID).Return("", fmt.Errorf("can't download image from id"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: "can't download image from id\n",
			checkHeaders:         false,
		},
		{
			name: "Client get error",
			mockBehavior: func(r *mocksstorage.MockStorageInterface, w *mocks.MockServicesInterface) {
				w.EXPECT().GetImageByID(gomock.Any()).Return(imageModel, nil)
				r.EXPECT().DownloadImageFromID(imageModel.ID).Return("mock", nil)
			},
			expectedStatusCode:   500,
			expectedResponseBody: "Get \"mock\": unsupported protocol scheme \"\"\n",
			checkHeaders:         false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			services := mocks.NewMockServicesInterface(c)
			st := mocksstorage.NewMockStorageInterface(c)
			tt.mockBehavior(st, services)
			broker := Server{messageBroker: nil}
			handler := NewServer(services, st, broker.messageBroker)
			r := mux.NewRouter()
			r.HandleFunc("/image/downloadImage", handler.downloadImage).Methods("GET")
			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/image/downloadImage", nil)
			r.ServeHTTP(w, req)
			actualDisposition := w.Result().Header.Get("Content-Disposition")
			actualContentType := w.Result().Header.Get("Content-Type")
			assert.Equal(t, tt.expectedStatusCode, w.Code)
			if tt.checkHeaders {
				assert.Equal(t, tt.expectedDispositionHeader, actualDisposition)
				assert.Equal(t, tt.expectedConTypeHeader, actualContentType)
			} else {
				resultBody, _ := ioutil.ReadAll(w.Result().Body)
				assert.Equal(t, tt.expectedResponseBody, string(resultBody))
			}
		})
	}
}

func requestTest(t *testing.T, filename, paramName, url string, params map[string]string, targetFormat string) *http.Request {
	newfilename := strings.TrimSuffix(filename, "."+"png")
	file, err := os.Create(newfilename + "." + targetFormat)
	if err != nil {
		assert.NoError(t, err)
	}
	defer file.Close()
	img := image.NewRGBA(image.Rect(0, 0, 20, 20))
	for x := 0; x < 20; x++ {
		for y := 0; y < 20; y++ {
			img.Set(x, y, color.White)
		}
	}
	err = jpeg.Encode(file, img, nil)
	assert.NoError(t, err)
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(paramName, filepath.Base(filename))
	if err != nil {
		assert.NoError(t, err)
	}
	_, err = io.Copy(part, file)

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	err = writer.Close()
	if err != nil {
		assert.NoError(t, err)
	}

	req, err := http.NewRequest("POST", url, body)
	assert.NoError(t, err)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	return req
}

func TestHandler_convert(t *testing.T) {
	type payload service.ConvertPayLoad
	type mockBehavior func(w *mocks.MockServicesInterface)
	tests := []struct {
		name                 string
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
		params               map[string]string
		targetFormat         string
		payload              payload
		formValue            string
	}{
		{
			name: "Ok",
			mockBehavior: func(w *mocks.MockServicesInterface) {
				w.EXPECT().Convert(gomock.Any()).
					Return("1", nil)
			},
			params: map[string]string{
				"sourceFormat": "png",
				"targetFormat": "jpg",
				"ratio":        "99",
			},
			formValue:            "image",
			expectedStatusCode:   200,
			expectedResponseBody: "{\"id\":\"1\"}\n",
			targetFormat:         "jpg",
		},
		{
			name:         "invalid ratio",
			mockBehavior: func(w *mocks.MockServicesInterface) {},
			params: map[string]string{
				"sourceFormat": "png",
				"targetFormat": "jpg",
				"ratio":        "qweqwe",
			},
			formValue:            "image",
			expectedStatusCode:   400,
			targetFormat:         "jpg",
			expectedResponseBody: "invalid ratio\n",
		},
		{
			name:         "name of the format should not be empty",
			mockBehavior: func(w *mocks.MockServicesInterface) {},
			params: map[string]string{
				"sourceFormat": "png",
				"targetFormat": "",
				"ratio":        "77",
			},
			formValue:            "image",
			expectedStatusCode:   400,
			targetFormat:         "jpg",
			expectedResponseBody: "name of the format should not be empty\n",
		},
		{
			name:         "name of the format should not be empty",
			mockBehavior: func(w *mocks.MockServicesInterface) {},

			params: map[string]string{
				"sourceFormat": "png",
				"targetFormat": "jpg",
				"ratio":        "77",
			},
			formValue:            "rfr",
			expectedStatusCode:   400,
			targetFormat:         "jpg",
			expectedResponseBody: "name of the format should not be empty\n",
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
			r.HandleFunc("/image/convert", handler.convert).Methods("POST")
			w := httptest.NewRecorder()
			req := requestTest(t, "image.png", tt.formValue, "/image/convert", tt.params, tt.targetFormat)
			defer os.Remove("image.jpg")
			r.ServeHTTP(w, req)
			assert.Equal(t, tt.expectedStatusCode, w.Code)
			assert.Equal(t, tt.expectedResponseBody, w.Body.String())
		})
	}
}
