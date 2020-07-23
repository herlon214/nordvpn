package nordvpn

type Logger interface {
	Printf(format string, args ...interface{})
}
