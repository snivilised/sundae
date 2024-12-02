package locale

import (
	"github.com/snivilised/li18ngo"
)

const (
	SundaeSourceID = "github.com/snivilised/sundae"
)

func Use(options ...li18ngo.UseOptionFn) error {
	return li18ngo.Use(options...)
}

// sundaeTemplData
type sundaeTemplData struct{}

func (td sundaeTemplData) SourceID() string {
	return SundaeSourceID
}
