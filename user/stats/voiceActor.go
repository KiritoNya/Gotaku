package stats

import "github.com/KiritoNya/gotaku"

//VoiceActor - Struct that describes the stats by voice actors
type VoiceActor struct {
	stat
	VoiceActor   *gotaku.Staff
	CharacterIds []int
}
