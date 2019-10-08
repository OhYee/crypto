package blockhole

// BlackHole writer, but do nothing
type BlackHole struct{}

func (bh BlackHole) Write(b []byte) (n int, err error) {
	return len(b), nil
}
