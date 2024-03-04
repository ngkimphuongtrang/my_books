package httpkit

type Route struct {
	Name         string
	Method       string
	Path         string
	AuthRequired bool
}
