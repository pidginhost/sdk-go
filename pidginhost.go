// Package pidginhostsdk provides a simple wrapper around the generated PidginHost API client.
//
// Usage:
//
//	client := pidginhostsdk.New("your-api-token")
//	servers, _, err := client.Cloud.CloudServersList(ctx).Execute()
package pidginhostsdk

// New creates a PidginHost API client authenticated with the given token.
// Optionally pass a custom host URL (defaults to https://www.pidginhost.com).
func New(token string, host ...string) *APIClient {
	cfg := NewConfiguration()
	if len(host) > 0 && host[0] != "" {
		cfg.Servers = ServerConfigurations{
			{URL: host[0], Description: "Custom"},
		}
	}
	cfg.AddDefaultHeader("Authorization", "Token "+token)
	return NewAPIClient(cfg)
}
