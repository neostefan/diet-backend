package db

type Error struct {
	errMsg string
}

func (e Error) Error() string {
	return e.errMsg
}
