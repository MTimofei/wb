package cache

type Cache interface {
	Set(Payload) error
	SetList([]Payload) error
	Single(kay string) (Payload, error)
}

type Payload struct {
	Key   string
	Value []byte
}

func NewPayload(key string, value []byte) Payload {
	return Payload{
		Key:   key,
		Value: value,
	}
}
