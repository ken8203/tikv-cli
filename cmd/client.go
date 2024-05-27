package cmd

import "github.com/ken8203/tikv-cli/internal/client"

var c client.Client

// newClient creates a tikv client.
func newClient() (client.Client, error) {
	var v client.APIVersion
	switch APIVersion {
	case "v1":
		v = client.APIVersion1
	case "v1ttl":
		v = client.APIVersion1TTL
	case "v2":
		v = client.APIVersion2
	default:
		v = client.APIVersion2
	}

	c, err := client.New([]string{addr(Host, Port)}, client.Mode(Mode), v, KeySpace)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func addr(host, port string) string {
	return host + ":" + port
}
