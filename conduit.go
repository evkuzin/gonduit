package gonduit

import (
	"github.com/evkuzin/gonduit/requests"
	"github.com/evkuzin/gonduit/responses"
)

// ConduitQuery performs a call to conduit.query.
func (c *Conn) ConduitQuery() (*responses.ConduitQueryResponse, error) {
	var res responses.ConduitQueryResponse

	if err := c.Call("conduit.query", &requests.Request{}, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
