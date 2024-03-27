package product_analytics

import (
    "github.com/posthog/posthog-go"
)

type client_wrap struct {
	posthogClient posthog.Client
	userID string
}

func NewProductAnalyticsClient(api_key string, host string) (ProductAnalyticsClient, error) {
	client, err := posthog.NewWithConfig(api_key, posthog.Config{
		Endpoint: host,
	})

	if err != nil {
		return nil, err
	}

	return &client_wrap{
		posthogClient: client,
	}, nil
}

func (c *client_wrap) SetUserID(id string) {
	c.userID = id
}

func (c *client_wrap) SendEvent(name string, props map[string]any) {
	c.posthogClient.Enqueue(posthog.Capture{
		DistinctId: c.userID,
		Event:      name,
		Properties: props,
	})
}

func (c *client_wrap) InitUser(name string, props map[string]any) {
	c.posthogClient.Enqueue(posthog.Capture{
		DistinctId: c.userID,
		Event:      name,
		Properties: map[string]interface{}{
			"$set_once": props,
		},
	})
}

func (c *client_wrap) UpdateUser(name string, props map[string]any) {
	c.posthogClient.Enqueue(posthog.Capture{
		DistinctId: c.userID,
		Event:      name,
		Properties: map[string]interface{}{
			"$set": props,
		},
	})
}

func (c *client_wrap) Close() error {
	return c.posthogClient.Close()
}