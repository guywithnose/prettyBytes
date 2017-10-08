package prettyBytes

import "fmt"

type byteSize float64

const (
	_           = iota // ignore first value by assigning to blank identifier
	kb byteSize = 1 << (10 * iota)
	mb
	gb
	tb
	pb
	yb
)

// Format formats bytes as a human readable string
func Format(bytes float64) string {
	b := byteSize(bytes)
	if b > (10 * yb) {
		return fmt.Sprintf("%.2fYB", b/yb)
	} else if b > pb/10 {
		return fmt.Sprintf("%.2fPB", b/pb)
	} else if b > tb/10 {
		return fmt.Sprintf("%.2fTB", b/tb)
	} else if b > gb/10 {
		return fmt.Sprintf("%.2fGB", b/gb)
	} else if b > mb/10 {
		return fmt.Sprintf("%.2fMB", b/mb)
	} else if b > kb/10 {
		return fmt.Sprintf("%.2fKB", b/kb)
	}

	return fmt.Sprintf("%dB", int(bytes))
}
