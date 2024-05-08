package bbproto

type Version byte

const VERSION = 1

func (v *Version) populateFromBytes(b []byte) {
	*v = Version(b[0])
}
