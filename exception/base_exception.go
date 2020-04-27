package exception

type BaseException interface {
	error
	GetCode() int
}
