package discordrpc

import (
	"math/rand"
	"os"
	"time"

	"github.com/tnychn/mpv-discord/discordrpc/payloads"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letters = []rune("0123456789abcdefghijklmnopqrstuvwxyz")

func nonce(length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func newActivityPayload() payloads.Activity {
	return payloads.Activity{
		Cmd:   "SET_ACTIVITY",
		Nonce: nonce(20),
		Args: struct {
			Pid      int                    `json:"pid"`
			Activity *payloads.ActivityMain `json:"activity,omitempty"`
		}{os.Getpid(), nil},
	}
}

func mapActivityMainPayload(activity Activity) *payloads.ActivityMain {
	payload := new(payloads.ActivityMain)
	payload.State = activity.State
	payload.Details = activity.Details
	payload.Assets = struct {
		LargeImage string `json:"large_image,omitempty"`
		LargeText  string `json:"large_text,omitempty"`
		SmallImage string `json:"small_image,omitempty"`
		SmallText  string `json:"small_text,omitempty"`
	}{activity.LargeImageKey, activity.LargeImageText, activity.SmallImageKey, activity.SmallImageText}
	if activity.Party != nil {
		payload.Party = &struct {
			ID   string `json:"id,omitempty"`
			Size [2]int `json:"size,omitempty"`
		}{ID: activity.Party.ID, Size: [2]int{activity.Party.Players, activity.Party.MaxPlayers}}
	}
	if activity.Secrets != nil {
		payload.Secrets = &struct {
			Match    string `json:"match,omitempty"`
			Join     string `json:"join,omitempty"`
			Spectate string `json:"spectate,omitempty"`
		}{Match: activity.Secrets.Match, Join: activity.Secrets.Join, Spectate: activity.Secrets.Spectate}
	}
	if activity.Timestamps != nil {
		var start, end int64
		if activity.Timestamps.Start != nil {
			start = activity.Timestamps.Start.Unix()
		}
		if activity.Timestamps.End != nil {
			end = activity.Timestamps.End.Unix()
		}
		payload.Timestamps = &struct {
			Start int64 `json:"start,omitempty"`
			End   int64 `json:"end,omitempty"`
		}{start, end}
	}
	return payload
}
