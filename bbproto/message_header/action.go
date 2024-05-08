package bbproto

type Action byte

const (
	Action_UNSPECIFIED Action = 0
	Action_BCST        Action = 1 // Broadcast
	Action_REQ         Action = 2 // Request
	Action_RESP        Action = 8 // Response
	Action_AWK         Action = 9 // Acknowledgement
)

func (action *Action) populateFromBytes(b []byte) {
	*action = Action(b[1])
}
