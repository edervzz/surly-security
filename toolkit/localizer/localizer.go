package localizer

const EN string = "en"
const ES string = "es"

type ILocalizer interface {
	Localize(id string) string
	SetLanguage(langu string)
}

type localizer struct {
	messages map[string]map[string]string
	langu    string
}

func NewLocalizer(langu string, messageMap map[string]map[string]string) *localizer {
	return &localizer{
		messages: messageMap,
		langu:    langu,
	}
}

func (l *localizer) SetLanguage(langu string) {
	l.langu = langu
}

func (l localizer) Localize(id string) string {
	mess := l.messages[id][l.langu]
	if mess == "" {
		mess = l.messages[id][EN]
	}
	if mess == "" {
		mess = id
	}

	return mess
}
