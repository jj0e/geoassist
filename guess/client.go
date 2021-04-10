package guess

import (
	"fmt"
	"log"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/jj0e/geoguessr-assist/constants"
	"github.com/jj0e/geoguessr-assist/utils"
	"github.com/jasonwinn/geocoder"
	"github.com/pariz/gountries"
)

func New(username string, password string) *GameInstance {
	session := &GameInstance{
		GameClient: resty.New(),
		Username:   username,
		Password:   password,
	}
	session.Login()
	geocoder.SetAPIKey(constants.GeoCoderApiKey)
	return session
}

func (game *GameInstance) Join(uuid string) {
	game.Endpoint = constants.GeoGuessrBattleRoyaleEndpoint + "/" + uuid
}

func (game *GameInstance) GetGameData() *GameResult {
	resp, err := game.GameClient.R().SetResult(&GameResult{}).Get(game.Endpoint)
	if err != nil {
		log.Fatal(err)
	}
	result := resp.Result().(*GameResult)
	return result
}

func (game *GameInstance) Watch() {
	isEnded := false
	checkedRounds := 0
	for !isEnded {
		result := game.GetGameData()
		if checkedRounds < result.RoundNumber {
			currentLat := result.Rounds[result.RoundNumber-1].Latitude
			currentLon := result.Rounds[result.RoundNumber-1].Longitude
			address, err := geocoder.ReverseGeocode(currentLat, currentLon)
			if err != nil {
				fmt.Print(utils.GetTimestamp(), fmt.Sprintf("[Round %d] Unable to locate country\n", result.RoundNumber))
			} else {
				countryCode := address.CountryCode
				query := gountries.New()
				country, _ := query.FindCountryByAlpha(countryCode)
				fmt.Print(utils.GetTimestamp(), fmt.Sprintf("[Round %d] %s\n", result.RoundNumber, country.Name.Common))
			}
			checkedRounds++
		}
		isEnded = result.IsEnded
		time.Sleep(time.Second)
	}
	fmt.Print(utils.GetTimestamp(), "The game ended, hopefully you won!\n")
}
