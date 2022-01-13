package service

import (
	"context"
	"image"
	"image/color"
	"image/jpeg"
	"os"
	"testing"

	"github.com/Nikby53/image-converter/internal/repository/mockrepo"
	"github.com/Nikby53/image-converter/internal/storage/mocksstorage"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestService_Conversion(t *testing.T) {
	file, err := os.Create("image.jpg")
	assert.NoError(t, err)

	img := image.NewRGBA(image.Rect(0, 0, 20, 20))

	for x := 0; x < 20; x++ {
		for y := 0; y < 20; y++ {
			img.Set(x, y, color.White)
		}
	}

	picture, err := os.Open("image.jpg")
	assert.NoError(t, err)

	err = jpeg.Encode(file, img, nil)
	assert.NoError(t, err)

	type mockBehavior func(s *mocksstorage.MockStorageInterface, r *mockrepo.MockRepoInterface)

	tests := []struct {
		name          string
		mockBehavior  mockBehavior
		input         ConversionPayLoad
		expectedReqID string
		expectedError error
	}{
		{
			name: "Ok",
			mockBehavior: func(s *mocksstorage.MockStorageInterface, r *mockrepo.MockRepoInterface) {
				r.EXPECT().Transactional(gomock.Any()).Return(nil)
			},
			input: ConversionPayLoad{
				SourceFormat: "jpg",
				TargetFormat: "png",
				Filename:     "image",
				Ratio:        77,
				File:         picture,
				UsersID:      1,
			},
			expectedReqID: "1",
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			repo := mockrepo.NewMockRepoInterface(c)
			st := mocksstorage.NewMockStorageInterface(c)
			tt.mockBehavior(st, repo)
			service := New(repo, st)
			_, err := service.Conversion(context.Background(), tt.input)
			assert.Equal(t, tt.expectedError, err)
			err = file.Close()
			assert.NoError(t, err)
			err = picture.Close()
			assert.NoError(t, err)
			defer func() {
				err = os.Remove("image.jpg")
				if err != nil {
					assert.NoError(t, err)
				}
			}()
		})
	}
}
