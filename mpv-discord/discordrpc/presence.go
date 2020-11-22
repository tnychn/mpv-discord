package discordrpc

import (
	"time"
)

type Activity struct {
	State          string
	Details        string
	LargeImageKey  string
	LargeImageText string
	SmallImageKey  string
	SmallImageText string

	Party      *ActivityParty
	Secrets    *ActivitySecrets
	Timestamps *ActivityTimestamps
}

type ActivityParty struct {
	ID         string
	Players    int
	MaxPlayers int
}

type ActivitySecrets struct {
	Match    string
	Join     string
	Spectate string
}

type ActivityTimestamps struct {
	Start *time.Time
	End   *time.Time
}

type Presence struct{ *Client }

func NewPresence(cid string) *Presence {
	return &Presence{NewClient(cid)}
}

func (p *Presence) Update(activity Activity) error {
	payload := newActivityPayload()
	payload.Args.Activity = mapActivityMainPayload(activity)
	return p.send(1, payload)
}
