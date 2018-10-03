package provider

// Provider defines a video provider
type Provider interface {
	// GetInfoURLs returns URLs for retrieving video infomation
	GetInfoURL(videoID string) string
	// ParseVideoID parses URLs and try to get Video ID
	ParseVideoID(url string) (string, error)
	// DecodeDownloadURL decodes the video download URL from video infomation
	DecodeDownloadURL(url string) string
}

// List of Providers
type List []Provider

// ProviderList is an instance of List
var ProviderList List
