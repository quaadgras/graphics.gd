package gdreference

type Type int

const (
	TypeUnsafe Type = iota // Raw pointer.
	TypePooled             // main thread owns the pointer.
	TypePinned             // main thread owns the pointer, Kept alive.
	TypeThread             // off-thread owns the pointer.
	TypeBorrow             // Engine owns the pointer.
	TypeStatic             // Allocated at init.
)
