package handler

import (
	"bytes"
	"context"
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

	"github.com/Nikby53/image-converter/internal/models"
	"github.com/Nikby53/image-converter/internal/service/mocks"
	"github.com/Nikby53/image-converter/internal/storage/mocksstorage"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
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
	type mockBehavior func(r *mocks.MockServicesInterface, token string, ctx context.Context)
	tests := []struct {
		name                 string
		mockBehavior         mockBehavior
		headerName           string
		headerValue          string
		token                string
		expectedStatusCode   int
		expectedResponseBody string
		ctx                  context.Context
	}{
		{
			name: "Ok",
			mockBehavior: func(r *mocks.MockServicesInterface, token string, ctx context.Context) {
				r.EXPECT().ParseToken(token).Return(1, nil)
				r.EXPECT().GetRequestFromID(gomock.Any(), 1).Return(request, nil)
			},
			headerName:           "Authorization",
			headerValue:          "Bearer token",
			token:                "token",
			expectedStatusCode:   200,
			expectedResponseBody: string(requestJSON),
		},
		{
			name: "Repo error",
			mockBehavior: func(r *mocks.MockServicesInterface, token string, ctx context.Context) {
				r.EXPECT().ParseToken(token).Return(1, nil)
				r.EXPECT().GetRequestFromID(gomock.Any(), 1).Return(nil, fmt.Errorf("mock error"))
			},
			headerName:           "Authorization",
			headerValue:          "Bearer token",
			token:                "token",
			expectedStatusCode:   500,
			expectedResponseBody: "{\"error\":\"mock error\"}\n",
		},
		{
			name:                 "Invalid token",
			mockBehavior:         func(r *mocks.MockServicesInterface, token string, ctx context.Context) {},
			headerName:           "Authorization",
			headerValue:          "Bearer ",
			token:                "token",
			expectedStatusCode:   401,
			expectedResponseBody: "{\"error\":\"token is empty\"}\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			services := mocks.NewMockServicesInterface(c)
			tt.mockBehavior(services, tt.token, context.Background())
			storage := Server{storage: nil}
			server := NewServer(services, storage.storage)
			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/requests", nil)
			req.Header.Set(tt.headerName, tt.headerValue)
			server.router.ServeHTTP(w, req)
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
	type mockBehavior func(r *mocksstorage.MockStorageInterface, w *mocks.MockServicesInterface, token string)
	tests := []struct {
		name                      string
		mockBehavior              mockBehavior
		checkHeaders              bool
		headerName                string
		headerValue               string
		token                     string
		expectedStatusCode        int
		expectedResponseBody      string
		expectedDispositionHeader string
		expectedConTypeHeader     string
	}{
		{
			name: "Ok",
			mockBehavior: func(r *mocksstorage.MockStorageInterface, w *mocks.MockServicesInterface, token string) {
				w.EXPECT().ParseToken(token).Return(1, nil)
				w.EXPECT().GetImageByID(gomock.Any(), gomock.Any()).Return(imageModel, nil)
				r.EXPECT().DownloadImageFromID(imageModel.ID).Return("https://images-convert.s3.eu-central-1.amazonaws.com/image.png", nil)
			},
			headerName:                "Authorization",
			headerValue:               "Bearer token",
			token:                     "token",
			expectedStatusCode:        200,
			checkHeaders:              true,
			expectedDispositionHeader: "attachment; filename=" + imageModel.Name + "." + imageModel.Format,
			expectedConTypeHeader:     imageModel.Format,
		},
		{
			name: "Repository error",
			mockBehavior: func(r *mocksstorage.MockStorageInterface, w *mocks.MockServicesInterface, token string) {
				w.EXPECT().ParseToken(token).Return(1, nil)
				w.EXPECT().GetImageByID(gomock.Any(), gomock.Any()).Return(imageModel, fmt.Errorf("repo mock error"))
			},
			headerName:           "Authorization",
			headerValue:          "Bearer token",
			token:                "token",
			expectedStatusCode:   500,
			expectedResponseBody: "{\"error\":\"repo mock error\"}\n",
			checkHeaders:         false,
		},
		{
			name: "Can't download image",
			mockBehavior: func(r *mocksstorage.MockStorageInterface, w *mocks.MockServicesInterface, token string) {
				w.EXPECT().ParseToken(token).Return(1, nil)
				w.EXPECT().GetImageByID(gomock.Any(), gomock.Any()).Return(imageModel, nil)
				r.EXPECT().DownloadImageFromID(imageModel.ID).Return("", fmt.Errorf("can't download image from id"))
			},
			headerName:           "Authorization",
			headerValue:          "Bearer token",
			token:                "token",
			expectedStatusCode:   500,
			expectedResponseBody: "{\"error\":\"can't download image from id\"}\n",
			checkHeaders:         false,
		},
		{
			name: "Client get error",
			mockBehavior: func(r *mocksstorage.MockStorageInterface, w *mocks.MockServicesInterface, token string) {
				w.EXPECT().ParseToken(token).Return(1, nil)
				w.EXPECT().GetImageByID(gomock.Any(), gomock.Any()).Return(imageModel, nil)
				r.EXPECT().DownloadImageFromID(imageModel.ID).Return("mock", nil)
			},
			headerName:           "Authorization",
			headerValue:          "Bearer token",
			token:                "token",
			expectedStatusCode:   500,
			expectedResponseBody: "{\"error\":\"Get \\\"mock\\\": unsupported protocol scheme \\\"\\\"\"}\n",
			checkHeaders:         false,
		},
		{
			name:                 "Invalid token",
			mockBehavior:         func(r *mocksstorage.MockStorageInterface, w *mocks.MockServicesInterface, token string) {},
			headerName:           "Authorization",
			headerValue:          "Bearer ",
			token:                "token",
			expectedStatusCode:   401,
			expectedResponseBody: "{\"error\":\"token is empty\"}\n",
			checkHeaders:         false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			services := mocks.NewMockServicesInterface(c)
			st := mocksstorage.NewMockStorageInterface(c)
			tt.mockBehavior(st, services, tt.token)
			server := NewServer(services, st)
			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/image/download/1", nil)
			req.Header.Set(tt.headerName, tt.headerValue)
			server.router.ServeHTTP(w, req)
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

func requestTest(t *testing.T, filename, paramName, url, targetFormat string, params map[string]string) *http.Request {
	newFileName := strings.TrimSuffix(filename, "."+"png")
	file, err := os.Create(newFileName + "." + targetFormat)
	if err != nil {
		assert.NoError(t, err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			assert.NoError(t, err)
		}
	}(file)
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
	assert.NoError(t, err)
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
	type mockBehavior func(w *mocks.MockServicesInterface, token string)
	tests := []struct {
		name                 string
		mockBehavior         mockBehavior
		headerName           string
		headerValue          string
		token                string
		expectedStatusCode   int
		expectedResponseBody string
		params               map[string]string
		targetFormat         string
		formValue            string
	}{
		{
			name: "Ok",
			mockBehavior: func(w *mocks.MockServicesInterface, token string) {
				w.EXPECT().ParseToken(token).Return(1, nil)
				w.EXPECT().Conversion(gomock.Any(), gomock.Any()).
					Return("1", nil)
			},
			headerName:  "Authorization",
			headerValue: "Bearer token",
			token:       "token",
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
			name: "invalid ratio",
			mockBehavior: func(w *mocks.MockServicesInterface, token string) {
				w.EXPECT().ParseToken(token).Return(1, nil).Times(1)
			},
			headerName:  "Authorization",
			headerValue: "Bearer token",
			token:       "token",
			params: map[string]string{
				"sourceFormat": "png",
				"targetFormat": "jpg",
				"ratio":        "qweqwe",
			},
			formValue:            "image",
			expectedStatusCode:   400,
			targetFormat:         "jpg",
			expectedResponseBody: "{\"error\":\"strconv.Atoi: parsing \\\"qweqwe\\\": invalid syntax\"}\n",
		},
		{
			name: "name of the format should not be empty",
			mockBehavior: func(w *mocks.MockServicesInterface, token string) {
				w.EXPECT().ParseToken(token).Return(1, nil)
			},
			headerName:  "Authorization",
			headerValue: "Bearer token",
			token:       "token",
			params: map[string]string{
				"sourceFormat": "png",
				"targetFormat": "",
				"ratio":        "77",
			},
			formValue:            "image",
			expectedStatusCode:   400,
			targetFormat:         "jpg",
			expectedResponseBody: "{\"error\":\"name of the format should not be empty\"}\n",
		},
		{
			name: "invalid header value",
			mockBehavior: func(w *mocks.MockServicesInterface, token string) {
				w.EXPECT().ParseToken(token).Return(1, nil)
			},
			headerName:  "Authorization",
			headerValue: "Bearer token",
			token:       "token",
			params: map[string]string{
				"sourceFormat": "png",
				"targetFormat": "jpg",
				"ratio":        "77",
			},
			formValue:            "rfr",
			expectedStatusCode:   400,
			targetFormat:         "jpg",
			expectedResponseBody: "{\"error\":\"http: no such file\"}\n",
		},
		{
			name: "Source format should be jpg",
			mockBehavior: func(w *mocks.MockServicesInterface, token string) {
				w.EXPECT().ParseToken(token).Return(1, nil)
			},
			headerName:  "Authorization",
			headerValue: "Bearer token",
			token:       "token",
			params: map[string]string{
				"sourceFormat": "erqer",
				"targetFormat": "png",
				"ratio":        "77",
			},
			formValue:            "image",
			expectedStatusCode:   400,
			targetFormat:         "jpg",
			expectedResponseBody: "{\"error\":\"name of source format should be png\"}\n",
		},
		{
			name: "Source format should be png",
			mockBehavior: func(w *mocks.MockServicesInterface, token string) {
				w.EXPECT().ParseToken(token).Return(1, nil)
			},
			headerName:  "Authorization",
			headerValue: "Bearer token",
			token:       "token",
			params: map[string]string{
				"sourceFormat": "qweqeqw",
				"targetFormat": "jpg",
				"ratio":        "77",
			},
			formValue:            "image",
			expectedStatusCode:   400,
			targetFormat:         "jpg",
			expectedResponseBody: "{\"error\":\"name of source format should be png\"}\n",
		},
		{
			name: "can't parse jwt token",
			mockBehavior: func(r *mocks.MockServicesInterface, token string) {
				r.EXPECT().ParseToken(token).Return(1, fmt.Errorf("")).Times(1)
			},
			headerName:  "Authorization",
			headerValue: "Bearer token",
			token:       "token",
			params: map[string]string{
				"sourceFormat": "png",
				"targetFormat": "jpg",
				"ratio":        "77",
			},
			formValue:            "image",
			targetFormat:         "jpg",
			expectedStatusCode:   401,
			expectedResponseBody: "{\"error\":\"can't parse jwt token: \"}\n",
		},
		{
			name: "Can't convert",
			mockBehavior: func(r *mocks.MockServicesInterface, token string) {
				r.EXPECT().ParseToken(token).Return(1, nil)
				r.EXPECT().Conversion(gomock.Any(), gomock.Any()).Return("0", fmt.Errorf("can't convert image"))
			},
			headerName:  "Authorization",
			headerValue: "Bearer token",
			token:       "token",
			params: map[string]string{
				"sourceFormat": "png",
				"targetFormat": "jpg",
				"ratio":        "77",
			},
			formValue:            "image",
			targetFormat:         "jpg",
			expectedStatusCode:   500,
			expectedResponseBody: "{\"error\":\"can't convert image\"}\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			services := mocks.NewMockServicesInterface(c)
			tt.mockBehavior(services, tt.token)
			storage := Server{storage: nil}
			server := NewServer(services, storage.storage)
			w := httptest.NewRecorder()
			req := requestTest(t, "image.png", tt.formValue, "/image/convert", tt.targetFormat, tt.params)
			req.Header.Set(tt.headerName, tt.headerValue)
			defer func() {
				err := os.Remove("image.jpg")
				if err != nil {
					assert.NoError(t, err)
				}
			}()
			server.router.ServeHTTP(w, req)
			assert.Equal(t, tt.expectedStatusCode, w.Code)
			assert.Equal(t, tt.expectedResponseBody, w.Body.String())
		})
	}
}
