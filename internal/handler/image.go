package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
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
	processing   = "processing"
	done         = "done"
	defaultRatio = 100
)

// RequestID is for id output in JSON.
type RequestID struct {
	ID string `json:"id"`
}

func (s *Server) convert(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("image")
	if err != nil {
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
	imageID, err := s.services.InsertImage(filename, sourceFormat)
	if err != nil {
		http.Error(w, fmt.Sprintf("repository error: %v", err), http.StatusInternalServerError)
		return
	}
	err = s.storage.UploadFile(file, imageID)
	if err != nil {
		logrus.Printf("can't upload image: %v", err)
		return
	}
	sourceFile, err := s.storage.DownloadFile(imageID)
	if err != nil {
		http.Error(w, "can't download image", http.StatusInternalServerError)
		return
	}
	_, err = s.services.ConvertImage(sourceFile, targetFormat, ratio)
	if err != nil {
		logrus.Printf("can't convert image: %v", err)
		return
	}
	usersID, err := s.GetIDFromToken(r)
	if err != nil {
		http.Error(w, "can't get id from jwt token", http.StatusInternalServerError)
		return
	}
	s.logger.Infof("user with id %d converted image", usersID)
	requestID, err := s.services.RequestsHistory(sourceFormat, targetFormat, imageID, filename, usersID, ratio)
	if err != nil {
		http.Error(w, fmt.Sprintf("repository error: %v", err), http.StatusInternalServerError)
		return
	}
	targetImageID, err := s.services.InsertImage(filename, targetFormat)
	if err != nil {
		http.Error(w, fmt.Sprintf("repository error: %v", err), http.StatusInternalServerError)
		return
	}
	err = s.services.UpdateRequest(processing, imageID, targetImageID)
	if err != nil {
		http.Error(w, fmt.Sprintf("can't update request: %v", err), http.StatusInternalServerError)
	}
	err = s.storage.UploadTargetFile(filename+"."+targetFormat, targetImageID)
	if err != nil {
		s.logger.Printf("can't upload image: %v", err)
		return
	}
	err = s.services.UpdateRequest(done, imageID, targetImageID)
	if err != nil {
		http.Error(w, fmt.Sprintf("can't update request: %v", err), http.StatusInternalServerError)

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
		http.Error(w, fmt.Sprintf("can't get request history %v", err), http.StatusInternalServerError)
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
	imageID, err := s.services.GetImageID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	url, err := s.storage.DownloadImageFromID(imageID)
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
	name, format := s.services.GetImage(imageID)
	w.Header().Set("Content-Disposition", "attachment; filename="+name+"."+format)
	w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
	w.Header().Set("Content-Length", r.Header.Get("Content-Length"))
	io.Copy(w, resp.Body)
}
