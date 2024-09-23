package storage

import (
	"database/sql"
	"diamap/api/pkg/enteties"
	"fmt"
	"sync"

	"github.com/rs/zerolog/log"
)

type Database struct {
	DB *sql.DB
	m  sync.Mutex
}

func NewDBStorage(db *sql.DB) *Database {
	return &Database{DB: db}
}

func New(connStr string) (*Database, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("openning database: %w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("pinging database: %w", err)
	}
	return &Database{DB: db}, nil
}

func (storage *Database) GetAllPodcasts() ([]enteties.PodcastShortInfo, error) {
	const op = "storage.GetAllpodcasts"
	storage.m.Lock()
	defer storage.m.Unlock()

	query := `SELECT 
						p.podcastid AS podcastId,
					 	p.podcastname AS podcastName,
					 	p.duration,
					 	p.photoroute AS photoRoute,
					 	c.countryname AS country,
				     	c.flag_emoji AS emoji
			  FROM podcasts p
			  JOIN countries c ON p.countryid = c.countryid;
	`
	rows, err := storage.DB.Query(query)
	if err != nil {
		log.Error().Err(err).Msgf("%s: unable to get all podcasts", op)
		return nil, err
	}
	defer rows.Close()

	var podcasts []enteties.PodcastShortInfo
	for rows.Next() {
		var p enteties.PodcastShortInfo
		if err := rows.Scan(&p.PodcastId, &p.PodcastName, &p.Duration, &p.Photoroute, &p.Country, &p.CountryEmoji); err != nil {
			log.Error().Err(err).Msgf("%s: unable to scan podcast", op)
			return nil, err
		}
		podcasts = append(podcasts, p)
	}

	return podcasts, nil
}
