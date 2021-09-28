package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

func validateConvert(sourceFormat, filename, targetFormat string, ratio int) error {
	if len(filename) == 0 {
		return fmt.Errorf("name of the file should not be empty")
	}
	if len(sourceFormat) == 0 || len(targetFormat) == 0 {
		return fmt.Errorf("name of the format should not be empty")
	}
	if ratio < 1 || ratio > 99 {
		return fmt.Errorf("ratio should be in range of 1 to 99")
	}
	// TODO finish validate func
	//if strings.Contains(filename, "")
	return nil
}

const (
	queued     = "queued"
	processing = "processing"
	done       = "done"
)

type RequestID struct {
	ID string `json:"id"`
}

func (s *Server) convert(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("image")
	if err != nil {
		logrus.Printf("error retrieving the file %v", err)
		return
	}

	defer file.Close()
	sourceFormat := r.FormValue("sourceFormat")
	targetFormat := r.FormValue("targetFormat")
	filename := strings.TrimSuffix(header.Filename, "."+sourceFormat)
	ratio, err := strconv.Atoi(r.FormValue("ratio"))
	if err != nil {
		http.Error(w, "invalid ratio", http.StatusBadRequest)
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
	usersID, err := s.GetIdFromToken(r)
	if err != nil {
		http.Error(w, "can't get id from jwt token", http.StatusInternalServerError)
		return
	}
	err = s.services.UpdateRequest(queued, usersID)
	if err != nil {
		http.Error(w, fmt.Sprintf("can't update request: %v", err), http.StatusInternalServerError)
	}
	sourceFile, err := s.storage.DownloadFile(imageID)
	if err != nil {
		http.Error(w, "can't download image", http.StatusInternalServerError)
		return
	}
	convImageBytes, err := s.services.ConvertImage(sourceFile, targetFormat, ratio)
	if err != nil {
		logrus.Printf("can't convert image: %v", err)
		return
	}
	err = s.services.UpdateRequest(processing, usersID)
	if err != nil {
		http.Error(w, fmt.Sprintf("can't update request: %v", err), http.StatusInternalServerError)
	}
	err = ioutil.WriteFile(filename+"."+targetFormat, convImageBytes, 0644)
	if err != nil {
		return
	}
	fmt.Fprintf(w, "successfully uploaded file\n")
	requestID, err := s.services.RequestsHistory(sourceFormat, targetFormat, imageID, filename, usersID, ratio)
	if err != nil {
		http.Error(w, fmt.Sprintf("repository error: %v", err), http.StatusInternalServerError)
		return
	}
	err = s.services.UpdateRequest(done, usersID)
	if err != nil {
		http.Error(w, fmt.Sprintf("can't update request: %v", err), http.StatusInternalServerError)
	}
	err = json.NewEncoder(w).Encode(RequestID{ID: requestID})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logrus.Printf("signUp: error encoding json: %v", err)
		return
	}
}

func (s *Server) requestHistory(w http.ResponseWriter, r *http.Request) {
	usersID, err := s.GetIdFromToken(r)
	if err != nil {
		http.Error(w, "can't get id from jwt token", http.StatusInternalServerError)
		return
	}
	history, err := s.services.GetRequestFromId(usersID)
	if err != nil {
		http.Error(w, fmt.Sprintf("can't get request history %v", err), http.StatusInternalServerError)
		return
	}
	historyJSON, err := json.MarshalIndent(&history, "\t", "")
	fmt.Fprint(w, string(historyJSON))
}

func (s *Server) downloadImage(w http.ResponseWriter, r *http.Request) {

}
