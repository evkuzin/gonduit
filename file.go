package gonduit

import (
	"github.com/evkuzin/gonduit/requests"
	"github.com/evkuzin/gonduit/responses"
)

// FileDownload performs a call to file.download.
func (c *Conn) FileDownload(
	req requests.FileDownloadRequest,
) (*responses.FileDownloadResponse, error) {
	var res responses.FileDownloadResponse

	if err := c.Call("file.download", &req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
