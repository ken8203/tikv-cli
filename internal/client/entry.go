package client

type Entry struct {
	K, V []byte
}

func (e Entry) Key() string {
	return string(e.K)
}

func (e Entry) Value() string {
	return string(e.V)
}
