package nap

type CBRouter struct {
	Routers       map[int]RouterFunc
	DefaultRouter RouterFunc
}
