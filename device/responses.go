package device

type Hello struct {
	UID   string         `json:"uid"`
	Names []string       `json:"names"`
	Info  map[string]any `json:"info"`
}

type Error struct {
	Err string `json:"error"`
}

type GetVolume struct {
	Value int `json:"value"`
}

type GetMute struct {
	Value bool `json:"value"`
}
