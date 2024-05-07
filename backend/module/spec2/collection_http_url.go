package spec2

const NODE_HTTP_URL = "apicat-http-url"

type CollectionHttpUrl struct {
	Type  string       `json:"type" yaml:"type"`
	Attrs HttpUrlAttrs `json:"attr" yaml:"attrs"`
}

type HttpUrlAttrs struct {
	Path   string `json:"path" yaml:"path"`
	Method string `json:"method" yaml:"method"`
}

func NewCollectionHttpUrl(p, m string) *CollectionHttpUrl {
	return &CollectionHttpUrl{
		Type: NODE_HTTP_URL,
		Attrs: HttpUrlAttrs{
			Path:   p,
			Method: m,
		},
	}
}

func (u *CollectionHttpUrl) NodeType() string {
	return u.Type
}
