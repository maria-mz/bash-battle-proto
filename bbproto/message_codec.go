package bbproto

func EncodeMessage(header Byteable, payload Byteable) ([]byte, error) {
	headerBytes, err := header.Bytes()

	if err != nil {
		return nil, err
	}

	payloadBytes, err := payload.Bytes()

	if err != nil {
		return nil, err
	}

	message := append(headerBytes, payloadBytes...)

	return message, nil
}

func DecodeMessage(b []byte, headerLen int, header Byteable, payload Byteable) error {
	if err := header.PopulateFromBytes(b[:headerLen]); err != nil {
		return err
	}

	if err := payload.PopulateFromBytes(b[headerLen:]); err != nil {
		return err
	}

	return nil
}
