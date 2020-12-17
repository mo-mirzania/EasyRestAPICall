package nap

type RestResource struct {
	Endpoint string
	Method   string
	Router   *CBRouter
}
