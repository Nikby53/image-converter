package storage

import (
	"os"
	"testing"

	"github.com/Nikby53/image-converter/internal/storage/mocksstorage"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestStorage_UploadFile(t *testing.T) {
	file, _ := os.Create("image")
	type mockBehavior func(s *mocksstorage.MockStorageInterface)
	tests := []struct {
		name         string
		mockBehavior mockBehavior
	}{
		{
			name: "Ok",
			mockBehavior: func(s *mocksstorage.MockStorageInterface) {
				s.EXPECT().UploadFile(gomock.Any(), gomock.Any()).Return(nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			storage := mocksstorage.NewMockStorageInterface(c)
			tt.mockBehavior(storage)
			err := storage.UploadFile(file, "1")
			defer os.Remove("image")
			assert.NoError(t, err)
		})
	}
}
