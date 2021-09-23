package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"
)

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
	convImageBytes, err := s.services.Convert(fileBytes, targetFormat)
	if err != nil {
		logrus.Printf("can't convert image: %v", err)
		return
	}
	ioutil.WriteFile(filename+"."+targetFormat, convImageBytes, 0644)
	fmt.Fprintf(w, "successfully uploaded file\n")
}

func (s *Server) requestHistory(w http.ResponseWriter, r *http.Request) {

}
