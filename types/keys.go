package types

type KeyBind struct {
	Keys    KeyCombo `json:"keys"`
	Command string   `json:"cmd"`
	Desc    string   `json:"desc,omitempty"`
}

type KeyGroup struct {
	Name string    `json:"name"`
	Keys []KeyBind `json:"keys"`
}

type KeyCombo struct {
	Keys   []string `json:"keys"`
	Meta   bool     `json:"meta"`
	Shift  bool     `json:"shift"`
	Ctrl   bool     `json:"ctrl"`
	AltKey bool     `json:"alt"`
}

func (kc KeyCombo) String() string {
	result := ""
	if kc.Meta {
		result += "Meta +"
	}
	if kc.Shift {
		result += "Shift +"
	}
	if kc.Ctrl {
		result += "Ctrl +"
	}
	if kc.AltKey {
		result += "Alt +"
	}

	for i := 0; i < len(kc.Keys); i++ {
		if i != (len(kc.Keys) - 1) {
			result += kc.Keys[i] + "+ "
		} else {
			result += kc.Keys[i]
		}
	}
	return result

}

type KeySet interface {
	toKeyCombo(ident string) KeyCombo
}
