package bbproto

type Byteable interface {
	Bytes() ([]byte, error)
	PopulateFromBytes(b []byte) error
}
