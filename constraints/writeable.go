package constraints

import (
	"github.com/dploop/gostl/types"
)

type Writeable interface {
	Write(data types.Data)
}
