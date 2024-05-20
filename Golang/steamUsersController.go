package main

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
)

type SteamResponse struct {
	Response struct {
		SteamID string `json:"steamid"`
	} `json:"response"`
}

var db, err = connectToDB()

func steamUsers(w http.ResponseWriter, r *http.Request) {
	vanityNames := getVanityNamesArr()

	// log DB connection error
	if err != nil {
		// try to connect to the database again
		db, err = connectToDB()

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var wg sync.WaitGroup
	for _, vanityName := range vanityNames {
		wg.Add(1)
		go func(vanityName string) {
			defer wg.Done()
			response := fetchSteam(vanityName)

			_, err = db.Exec("UPDATE SteamUsers SET SteamID = ? WHERE VanityName LIKE ?", response, "%"+vanityName+"%")
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}(vanityName)
	}

	wg.Wait()

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(vanityNames)
}

func getVanityNamesArr() []string {
	if db == nil {
		var err error
		db, err = connectToDB()
		if err != nil {
			log.Fatal(err)
		}
	}

	rows, err := db.Query("SELECT VanityName FROM SteamUsers")
	if err != nil {
		log.Fatal(err)
	}

	var vanityNames []string
	for rows.Next() {
		var vanityName string
		if err := rows.Scan(&vanityName); err != nil {
			log.Fatal(err)
		}
		vanityNames = append(vanityNames, getUsernameFromSteamUrl(vanityName))
	}

	return vanityNames
}
func getUsernameFromSteamUrl(url string) string {
	parts := strings.Split(url, "/")
	if parts[len(parts)-1] == "" {
		parts = parts[:len(parts)-1]
	}

	return parts[len(parts)-1]
}

func fetchSteam(vanityName string) string {
	url := "http://api.steampowered.com/ISteamUser/ResolveVanityURL/v0001/?key=" + getEnvVar("STEAM_API_KEY") + "&vanityurl=" + vanityName

	// check if the SteamID is already in Redis
	if existsInRedis(url) {
		return getFromRedis(url)
	}

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var sr SteamResponse
	json.Unmarshal(body, &sr)

	// update the APIsCalled table
	updateAPIsCalled(url)

	// set the SteamID in redis
	setToRedis(url, sr.Response.SteamID)

	return sr.Response.SteamID
}

func updateAPIsCalled(url string) {
	// Check if the URL exists in the APIsCalled table
	var count int
	err := db.QueryRow("SELECT count FROM APIsCalled WHERE Api_url = ?", url).Scan(&count)

	switch {
	case err == sql.ErrNoRows:
		// URL does not exist, create a new record
		_, err := db.Exec("INSERT INTO APIsCalled (Api_url, count) VALUES (?, ?)", url, 1)
		if err != nil {
			log.Fatal(err)
		}
	case err != nil:
		log.Fatal(err)
	default:
		// URL exists, increment the count
		_, err := db.Exec("UPDATE APIsCalled SET count = count + 1 WHERE Api_url = ?", url)
		if err != nil {
			log.Fatal(err)
		}
	}
}
