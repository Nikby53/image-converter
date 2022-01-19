package main

import (
	"os"

	"github.com/Nikby53/image-converter/internal/handler"
	"github.com/Nikby53/image-converter/internal/repository"
	"github.com/Nikby53/image-converter/internal/service"
	"github.com/Nikby53/image-converter/internal/storage"
	"github.com/krzysztofreczek/go-structurizr/pkg/scraper"
	"github.com/krzysztofreczek/go-structurizr/pkg/view"
)

func main() {
	s := buildScraper()

	collector := handler.NewServer(service.New(&repository.Repository{}, &storage.Storage{}))
	svcStructure := s.Scrape(collector)

	f, err := os.Create("tools/c4/output.plantuml")
	if err != nil {
		panic(err)
	}

	defer func() {
		_ = f.Close()
	}()

	v := buildView()

	err = v.RenderStructureTo(svcStructure, f)
	if err != nil {
		panic(err)
	}
}

func buildScraper() scraper.Scraper {
	s, err := scraper.NewScraperFromConfigFile("tools/c4/scraper.yml")
	if err != nil {
		panic(err)
	}

	return s
}

func buildView() view.View {
	v, err := view.NewViewFromConfigFile("tools/c4/view.yml")
	if err != nil {
		panic(err)
	}

	return v
}
