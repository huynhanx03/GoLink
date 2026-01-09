package widecolumn

import "go-link/common/pkg/settings"

// New creates a new Wide Column DB connection
func New(config *settings.WideColumn) (*Client, error) {
	client := &Client{
		config: config,
	}

	if err := client.connect(); err != nil {
		return nil, err
	}

	return client, nil
}
