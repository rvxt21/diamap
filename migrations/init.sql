\c podcastsDiaMap;

CREATE TABLE IF NOT EXISTS continents(
    continentId SERIAL PRIMARY KEY,
    continentName VARCHAR(20)
);

CREATE TABLE IF NOT EXISTS countries(
    countryId SERIAL PRIMARY KEY,
    countryName VARCHAR(100) NOT NULL,
    continentId INT,
    flag_emoji TEXT,
    FOREIGN KEY (continentId) REFERENCES countries(countryId) ON DELETE SET NULL
);

CREATE TABLE IF NOT EXISTS podcasts (
    podcastId SERIAL PRIMARY KEY,
    podcastName VARCHAR(100) NOT NULL,
    duration VARCHAR(20) NOT NULL,
    podcastDescription TEXT NOT NULL,
    audioName VARCHAR(100) NOT NULL,
    photoRoute VARCHAR,
    audioRoute VARCHAR,
    countryId INT NOT NULL,
    FOREIGN KEY (countryId) REFERENCES countries(countryId) ON DELETE SET NULL
);

