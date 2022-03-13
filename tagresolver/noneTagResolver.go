package tagresolver

import "github.com/adimax2953/structproto/internal"

var _ internal.TagResolver = NoneTagResolver

func NoneTagResolver(fieldname, token string) (*internal.Tag, error) {
	var tag = &internal.Tag{
		Name: fieldname,
	}
	return tag, nil
}
