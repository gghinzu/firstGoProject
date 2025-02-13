package enum

type UserStatus int

const (
	Active UserStatus = iota
	Passive
	Deleted
	NotInitialized
)
