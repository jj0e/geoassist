package constants

import "fmt"

const (
	Scheme = "https"
	GeoGuessrHostname = "game-server.geoguessr.com"
)

var (
	GeoGuessrHost = fmt.Sprintf("%s://%s", Scheme, GeoGuessrHostname)
	GeoGuessrBattleRoyaleEndpoint = GeoGuessrHost + "/api/battle-royale"
)
