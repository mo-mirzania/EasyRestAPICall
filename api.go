package nap

type API struct {
	BaseURL       string
	Resources     map[string]RestResource
	DefaultRouter *CBRouter
	Client        *Client
}
