package product_analytics

type ProductAnalyticsClient interface {
	SetUserID(string)

	SendEvent(string, map[string]any)

	InitUser(string, map[string]any)
	UpdateUser(string, map[string]any)

	Close() error
}