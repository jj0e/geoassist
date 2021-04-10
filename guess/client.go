package guess

import (
	"github.com/go-resty/resty/v2"
	"github.com/jj0e/geoguessr-assist/constants"
)

func New(uuid string) *GameInstance {
	return &GameInstance{
		Endpoint: constants.GeoGuessrBattleRoyaleEndpoint + "/" + uuid,
		GameClient: resty.New(),
	}
}

func (game GameInstance) QueryCorrectGuesses() *GameResult {
	resp, _ := game.GameClient.R().SetResult(&GameResult{}).Get(game.Endpoint)
	result := resp.Result().(*GameResult)
	return result
}
