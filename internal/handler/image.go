package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/Nikby53/image-converter/internal/service"

	"github.com/gorilla/mux"
)

func validateConvert(sourceFormat, filename, targetFormat string, ratio int) error {
	if filename == "" {
		return fmt.Errorf("name of the file should not be empty")
	}
	if sourceFormat == "" || targetFormat == "" {
		return fmt.Errorf("name of the format should not be empty")
	}
	if ratio < 1 || ratio > 100 {
		return fmt.Errorf("ratio should be in range of 1 to 100")
	}

	// TODO finish validate func
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
		http.Error(w, fmt.Sprintf("invalid header value ( %v", err), http.StatusBadRequest)
		s.logger.Printf("error retrieving the file %v", err)
		return
	}
	defer file.Close()
	sourceFormat := r.FormValue("sourceFormat")
	targetFormat := r.FormValue("targetFormat")
	filename := strings.TrimSuffix(header.Filename, "."+sourceFormat)
	ratio := defaultRatio
	if r.FormValue("ratio") != "" {
		ratio, err = strconv.Atoi(r.FormValue("ratio"))
		if err != nil {
			http.Error(w, "invalid ratio", http.StatusBadRequest)
			return
		}
	}
	err = validateConvert(sourceFormat, filename, targetFormat, ratio)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	usersID, err := s.GetIDFromToken(r)
	if err != nil {
		http.Error(w, "can't get id from jwt token", http.StatusInternalServerError)
		return
	}
	payload := service.ConvertPayLoad{
		SourceFormat: sourceFormat,
		TargetFormat: targetFormat,
		Filename:     filename,
		Ratio:        ratio,
		File:         file,
		UsersID:      usersID,
	}
	requestID, err := s.services.Convert(payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(RequestID{ID: requestID})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		s.logger.Printf("error encoding json: %v", err)
		return
	}
}

func (s *Server) requests(w http.ResponseWriter, r *http.Request) {
	usersID, err := s.GetIDFromToken(r)
	if err != nil {
		http.Error(w, "can't get id from jwt token", http.StatusInternalServerError)
		return
	}
	history, err := s.services.GetRequestFromID(usersID)
	if err != nil {
		http.Error(w, fmt.Sprintf("repository error %v", err), http.StatusInternalServerError)
		return
	}
	historyJSON, err := json.MarshalIndent(&history, "\t", "")
	if err != nil {
		http.Error(w, fmt.Sprintf("error in output in JSON %v", err), http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, string(historyJSON))
}

func (s *Server) downloadImage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	image, err := s.services.GetImageByID(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("repository error, %v", err), http.StatusInternalServerError)
		return
	}
	url, err := s.storage.DownloadImageFromID(image.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	client := &http.Client{}
	resp, err := client.Get(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	w.Header().Set("Content-Disposition", "attachment; filename="+image.Name+"."+image.Format)
	w.Header().Set("Content-Type", image.Format)
	w.Header().Set("Content-Length", r.Header.Get("Content-Length"))
	io.Copy(w, resp.Body)
	s.logger.Infof("user successfully download image with id %v", image.ID)
}
