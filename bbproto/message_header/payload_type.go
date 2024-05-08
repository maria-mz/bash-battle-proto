package bbproto

type PayloadType byte

const (
	PayloadType_UNSPECIFIED      PayloadType = 0
	PayloadType_CREATE_GAME_REQ  PayloadType = 1
	PayloadType_CREATE_GAME_RESP PayloadType = 1
	// TODO: more as they come
)

func (payloadType *PayloadType) populateFromBytes(b []byte) {
	*payloadType = PayloadType(b[2])
}
