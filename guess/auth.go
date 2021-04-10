package guess

import (
	"github.com/jj0e/geoguessr-assist/constants"
)

func (game *GameInstance) Login() {
	resp, _ := game.GameClient.R().
	SetFormData(map[string]string{
		"email": game.Username,
		"password": game.Password,
	}).
	Post(constants.GeoGuessrLoginEndpoint)
	for _, cookie := range resp.Cookies() {
		if cookie.Name == "_ncfa" {
			game.AuthCookie = cookie.Value
		}
	}
}
