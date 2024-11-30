package locale

// CLIENT-TODO: Should be updated to use url of the implementing project,
// so should not be left as sundae. (this should be set by auto-check)
const SundaeSourceID = "github.com/snivilised/sundae"

type sundaeTemplData struct{}

func (td sundaeTemplData) SourceID() string {
	return SundaeSourceID
}
