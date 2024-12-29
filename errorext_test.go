package errorext

import "testing"

func TestCurrentFrame(t *testing.T) {
	frame := CurrentFrame()

	t.Logf("frame info: %v, %s, %d %n", frame, frame, frame, frame)
}
