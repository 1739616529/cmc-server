package resp

type Error struct {
	Code int
	Msg  string
	error
}
