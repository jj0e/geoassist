package guess

import (
	"github.com/go-resty/resty/v2"
)

type GameInstance struct {
	Endpoint   string
	GameClient *resty.Client
}

type GameResult struct {
	GameId      string `json:"gameId"`
	RoundNumber int    `json:"currentRoundNumber"`
	Players     []struct {
		ID      string `json:"playerId"`
		Guesses []struct {
			ID          string `json:"id"`
			RoundNumber int    `json:"roundNumber"`
			CountryCode string `json:"countryCode"`
			Correct     bool   `json:"wasCorrect"`
		} `json:"guesses"`
	} `json:"players"`
}
