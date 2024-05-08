package bbproto

type StatusCode byte

const (
	Status_Code_UNSPECIFIED  StatusCode = 0
	Status_Code_OK           StatusCode = 0
	Status_Code_SERVER_ERROR StatusCode = 1
	// TODO: more as they come
)

func (statusCode *StatusCode) populateFromBytes(b []byte) {
	*statusCode = StatusCode(b[3])
}
