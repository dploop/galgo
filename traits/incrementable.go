package traits

type Incrementable interface {
	Next() Incrementable
}
