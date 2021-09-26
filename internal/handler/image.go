package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

func validateConvert(sourceFormat, filename, targetFormat string) error {
	if len(filename) == 0 {
		return fmt.Errorf("name of the file should not be empty")
	}
	if len(sourceFormat) == 0 || len(targetFormat) == 0 {
		return fmt.Errorf("name of the format should not be empty")
	}
	//if strings.Contains(filename, "")
	return nil
}

func (s *Server) convert(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("image")
	if err != nil {
		logrus.Printf("error retrieving the File %v", err)
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
	err = validateConvert(sourceFormat, filename, targetFormat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	sourceFileID, err := s.services.InsertImage(filename, sourceFormat)
	if err != nil {
		http.Error(w, fmt.Sprintf("repository error: %v", err), http.StatusInternalServerError)
		return
	}
	err = s.storage.UploadFile(file, sourceFileID)
	if err != nil {
		logrus.Printf("can't upload image: %v", err)
		return
	}

	sourceFile, err := s.storage.DownloadFile(sourceFileID)
	if err != nil {
		return
	}
	convImageBytes, err := s.services.Convert(sourceFile, targetFormat, ratio)
	if err != nil {
		logrus.Printf("can't convert image: %v", err)
		return
	}
	err = ioutil.WriteFile(filename+"."+targetFormat, convImageBytes, 0644)
	if err != nil {
		return
	}
	fmt.Fprintf(w, "successfully uploaded file\n")
}

func (s *Server) requestHistory(w http.ResponseWriter, r *http.Request) {

}
