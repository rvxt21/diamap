package services

import "diamap/api/pkg/enteties"

type storage interface {
	GetAllPodcasts() ([]enteties.PodcastShortInfo, error)
}

type Service struct {
	Storage storage
}

func (service Service) GetAllPodcasts() ([]enteties.PodcastShortInfo, error) {
	podcasts, err := service.Storage.GetAllPodcasts()
	return podcasts, err
}
