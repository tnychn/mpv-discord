package payloads

type Payload interface {
	payload()
}

type Handshake struct {
	Payload `json:"-"`

	V        string `json:"v"`
	ClientID string `json:"client_id"`
}

type Activity struct {
	Payload `json:"-"`

	Cmd   string `json:"cmd"`
	Nonce string `json:"nonce"`
	Args  struct {
		Pid      int           `json:"pid"`
		Activity *ActivityMain `json:"activity,omitempty"`
	} `json:"args"`
}

type ActivityMain struct {
	State   string `json:"state,omitempty"`
	Details string `json:"details,omitempty"`
	Assets  struct {
		LargeImage string `json:"large_image,omitempty"`
		LargeText  string `json:"large_text,omitempty"`
		SmallImage string `json:"small_image,omitempty"`
		SmallText  string `json:"small_text,omitempty"`
	} `json:"assets,omitempty"`
	Party *struct {
		ID   string `json:"id,omitempty"`
		Size [2]int `json:"size,omitempty"`
	} `json:"party,omitempty"`
	Secrets *struct {
		Match    string `json:"match,omitempty"`
		Join     string `json:"join,omitempty"`
		Spectate string `json:"spectate,omitempty"`
	} `json:"secrets,omitempty"`
	Timestamps *struct {
		Start int64 `json:"start,omitempty"`
		End   int64 `json:"end,omitempty"`
	} `json:"timestamps,omitempty"`
}
