package generic

// Ordered encompasses all types that can be compared with the < > operators.
type Ordered interface {
	string |
		int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64
}
