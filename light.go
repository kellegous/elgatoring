package elgatoring

type Light struct {
	On          Bool        `json:"on"`
	Brightness  int         `json:"brightness"`
	Temperature Temperature `json:"temperature"`
}

type Bool int

func (b Bool) Value() bool {
	return b != 0
}
