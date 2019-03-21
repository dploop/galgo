package iterators

type Writeable interface {
	Write(data Data)
}

func Write(w Writeable, data Data) {
	w.Write(data)
}
