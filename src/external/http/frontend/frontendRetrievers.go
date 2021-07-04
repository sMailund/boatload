package frontend

type IFrontendRetriever interface {
	getFrontend() ([]byte, error)
}

type TextRetriever struct {
	Message string
}

func (t TextRetriever) getFrontend() ([]byte, error) {
	if t.Message == "" {
		t.Message = "<p>hello, world!</p>"
	}
	return []byte(t.Message), nil
}

