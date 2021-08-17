package soumuradio

// TargetInfo indicates what information is searched
// set ST
type TargetInfo int

const (
	// License means search for License information
	License TargetInfo = 1
	// Registration means search for registration information
	Registration TargetInfo = 2
)
