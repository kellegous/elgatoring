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

func BoolFrom(b bool) Bool {
	if b {
		return 1
	}
	return 0
}
