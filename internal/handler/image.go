package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"
)

func ValidateConvert(sourceFormat, filename, targetFormat string) error {
	if len(filename) == 0 {
		return fmt.Errorf("name of the file should not be empty")
	}
	if len(sourceFormat) == 0 {
		return fmt.Errorf("name of the format should not be empty")
	}
	return nil
}

func (s *Server) convert(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("file")
	if err != nil {
		logrus.Printf("error retrieving the File %v", err)
		return
	}

	defer file.Close()
	sourceFormat := r.FormValue("sourceFormat")
	targetFormat := r.FormValue("targetFormat")
	filename := strings.TrimSuffix(header.Filename, "."+sourceFormat)
	fileBytes, _ := ioutil.ReadAll(file)
	err = ValidateConvert(sourceFormat, filename, targetFormat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	convImageBytes, err := s.services.Convert(fileBytes, targetFormat)
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
