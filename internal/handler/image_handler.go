package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"

	"github.com/Nikby53/image-converter/internal/service"
	"github.com/gorilla/mux"
)

var (
	errShouldBeJpg  = errors.New("name of source format should be jpg")
	errShouldBeJpeg = errors.New("name of source format should be jpeg")
	errShouldBePng  = errors.New("name of source format should be png")
	errEmptyFormat  = errors.New("name of the format should not be empty")
	errInvalidRatio = errors.New("ratio should be in range of 1 to 100")
)

func validateConvert(sourceFormat, filename, targetFormat string, ratio int) error {
	if strings.Contains(filename, ".jpg") {
		if sourceFormat != "jpg" {
			return errShouldBeJpg
		}
	}

	if strings.Contains(filename, ".jpeg") {
		if sourceFormat != "jpeg" {
			return errShouldBeJpeg
		}
	}

	if strings.Contains(filename, ".png") {
		if sourceFormat != "png" {
			return errShouldBePng
		}
	}

	if sourceFormat == "" || targetFormat == "" {
		return errEmptyFormat
	}

	if ratio < 1 || ratio > 100 {
		return errInvalidRatio
	}

	return nil
}

const (
	defaultRatio = 100
)

// RequestID is for id output in JSON.
type RequestID struct {
	ID string `json:"id"`
}

func (s *Server) convert(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("image")
	if err != nil {
		s.errorJSON(w, http.StatusBadRequest, err)
		s.logger.Errorf("error retrieving the file %v", err)

		return
	}

	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			s.logger.Errorf("failed to close file %v", err)
		}
	}(file)

	sourceFormat := r.FormValue("sourceFormat")
	targetFormat := r.FormValue("targetFormat")
	filename := strings.TrimSuffix(header.Filename, "."+sourceFormat)
	ratio := defaultRatio

	if r.FormValue("ratio") != "" {
		ratio, err = strconv.Atoi(r.FormValue("ratio"))
		if err != nil {
			s.errorJSON(w, http.StatusBadRequest, err)
			s.logger.Errorf("invalid ratio %v", err)

			return
		}
	}

	err = validateConvert(sourceFormat, header.Filename, targetFormat, ratio)
	if err != nil {
		s.errorJSON(w, http.StatusBadRequest, err)
		s.logger.Errorf("error in validation %v", err.Error())

		return
	}

	userID, err := s.GetIDFromContext(r.Context())
	if err != nil {
		s.errorJSON(w, http.StatusInternalServerError, err)
		s.logger.Errorf("can't get id from jwt token %v", err)

		return
	}

	payload := service.ConversionPayLoad{
		SourceFormat: sourceFormat,
		TargetFormat: targetFormat,
		Filename:     filename,
		Ratio:        ratio,
		File:         file,
		UsersID:      userID,
	}

	requestID, err := s.service.Conversion(r.Context(), payload)
	if err != nil {
		s.errorJSON(w, http.StatusInternalServerError, err)
		s.logger.Errorf("conversion error %v", err.Error())

		return
	}

	s.logger.Infof("user successfully convert image with request id %v", requestID)

	err = json.NewEncoder(w).Encode(RequestID{ID: requestID})
	if err != nil {
		s.errorJSON(w, http.StatusInternalServerError, err)
		s.logger.Printf("error encoding json: %v", err)

		return
	}
}

func (s *Server) requests(w http.ResponseWriter, r *http.Request) {
	usersID, err := s.GetIDFromContext(r.Context())
	if err != nil {
		s.errorJSON(w, http.StatusInternalServerError, err)
		s.logger.Errorf("can't get id from jwt token %v", err)

		return
	}

	history, err := s.service.GetRequestFromID(r.Context(), usersID)
	if err != nil {
		s.errorJSON(w, http.StatusInternalServerError, err)
		s.logger.Errorf("request repository error %v", err.Error())

		return
	}

	historyJSON, err := json.MarshalIndent(&history, "\t", "")
	if err != nil {
		s.errorJSON(w, http.StatusInternalServerError, err)
		s.logger.Errorf("error in marshal %v", err)

		return
	}

	fmt.Fprint(w, string(historyJSON))
}

func (s *Server) downloadImage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	image, err := s.service.GetImageByID(r.Context(), id)
	if err != nil {
		s.errorJSON(w, http.StatusInternalServerError, err)
		s.logger.Errorf("download repository error %v", err.Error())

		return
	}

	url, err := s.service.DownloadImageFromID(image.ID)
	if err != nil {
		s.errorJSON(w, http.StatusInternalServerError, err)
		s.logger.Errorf("download image error %v", err)

		return
	}

	client := &http.Client{}

	resp, err := client.Get(url)
	if err != nil {
		s.errorJSON(w, http.StatusInternalServerError, err)
		s.logger.Errorf("client error %v", err)

		return
	}

	defer func() {
		err := resp.Body.Close()
		if err != nil {
			s.logger.Errorf("can't close body %v", err)
		}
	}()

	w.Header().Set("Content-Disposition", "attachment; filename="+image.Name+"."+image.Format)
	w.Header().Set("Content-Type", image.Format)
	w.Header().Set("Content-Length", r.Header.Get("Content-Length"))

	_, err = io.Copy(w, resp.Body)
	if err != nil {
		s.errorJSON(w, http.StatusInternalServerError, err)
		s.logger.Errorf("download copy error %v", err)

		return
	}

	s.logger.Infof("user successfully download image with id %v", image.ID)
}
