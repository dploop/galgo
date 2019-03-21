package iterators

type Readable interface {
	Read() Data
}

func Read(r Readable) Data {
	return r.Read()
}
