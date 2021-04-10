package constants

import "fmt"

const (
	Scheme = "https"
	GeoGuessrHostname = "www.geoguessr.com"
	GeoGuessrGameHostname = "game-server.geoguessr.com"
	GeoCoderApiKey = "USLblbtiXCvYxK15p2Y1ckAoopm4rkG3"
)

var (
	GeoGuessrHost = fmt.Sprintf("%s://%s", Scheme, GeoGuessrHostname)
	GeoGuessrGameHost = fmt.Sprintf("%s://%s", Scheme, GeoGuessrGameHostname)
	GeoGuessrLoginEndpoint = GeoGuessrHost + "/api/v3/accounts/signin"
	GeoGuessrBattleRoyaleEndpoint = GeoGuessrGameHost + "/api/battle-royale"
)
