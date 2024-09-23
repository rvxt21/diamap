package enteties

type Podcast struct {
	id          uint64
	podcastName string
	countryId   string
	duration    string
	audioName   string
	audioRoute  string
	photoRoute  string
}

type Country struct {
	countryId   uint64
	countryName string
	continentId uint64
	emoji       string
}

type Continent struct {
	continentId   uint64
	continentName string
}

type PodcastShortInfo struct {
	PodcastId    uint64
	PodcastName  string
	Country      string
	CountryEmoji string
	Duration     string
	Photoroute   string
}
