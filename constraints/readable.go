package constraints

import (
	"github.com/dploop/gostl/types"
)

type Readable interface {
	Read() types.Data
}
