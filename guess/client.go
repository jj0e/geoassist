package guess

import (
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/jj0e/geoguessr-assist/constants"
	"github.com/jj0e/geoguessr-assist/utils"
	"github.com/kelvins/geocoder"
)

func New(username string, password string) *GameInstance {
	session := &GameInstance{
		GameClient: resty.New(),
		Username:   username,
		Password:   password,
	}
	session.Login()
	return session
}

func (game *GameInstance) Join(uuid string) {
	game.Endpoint = constants.GeoGuessrBattleRoyaleEndpoint + "/" + uuid
}

func (game *GameInstance) GetGameData() *GameResult {
	resp, _ := game.GameClient.R().SetResult(&GameResult{}).Get(game.Endpoint)
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
			location := geocoder.Location{
				Latitude:  currentLat,
				Longitude: currentLon,
			}
			addresses, err := geocoder.GeocodingReverse(location)
			if err != nil {
				fmt.Print(utils.GetTimestamp(), fmt.Sprintf("[Round %d] Unable to locate country\n", result.RoundNumber))
			} else {
				address := addresses[0]
				fmt.Print(utils.GetTimestamp(), fmt.Sprintf("[Round %d] %s\n", result.RoundNumber, address.Country))
			}
			checkedRounds++
		}
		isEnded = result.IsEnded
		time.Sleep(time.Second)
	}
	fmt.Print(utils.GetTimestamp(), "The game ended, hopefully you won!\n")
}
