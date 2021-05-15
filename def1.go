package header

type Headers struct {
	h map[string][]string
}

func NewHeaders() *Headers {
	return &Headers{
		h: make(map[string][]string),
	}
}

func (h *Headers) Build() map[string][]string {
	return h.h
}

func (h *Headers) SetAccept(str ...string) *Headers {
	h.h[Accept] = str
	return h
}
