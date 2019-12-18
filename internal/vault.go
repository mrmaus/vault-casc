package internal

import (
	"fmt"
	"github.com/hashicorp/vault/api"
	"net/http"
	"os"
	"time"
)

//todo:
type VaultClient struct {
	client *api.Client
}

var httpClient = &http.Client{
	Timeout: 10 * time.Second,
}

func NewVaultClient(flags *Flags) *VaultClient {
	vaultAddress := vaultAddress(flags)

	fmt.Println("Connecting to vault at " + vaultAddress)

	token := flags.VaultToken

	client, err := api.NewClient(&api.Config{Address: vaultAddress, HttpClient: httpClient})
	if err != nil {
		panic(err)
	}

	client.SetToken(token)

	return &VaultClient{client}
}

func vaultAddress(flags *Flags) string {
	addr := os.Getenv("VAULT_ADDR")
	if addr != "" {
		return addr
	}

	addr = flags.VaultAddress
	if addr != "" {
		return addr
	}

	return "http://localhost:8200" //default value
}

//todo:
func (c *VaultClient) Sys() *api.Sys {
	return c.client.Sys()
}

//todo:
func (c *VaultClient) Policies(callback func(name string, content string)) {
	sys := c.client.Sys()
	policies, err := sys.ListPolicies()
	if err != nil {
		panic(err)
	}

	for _, policy := range policies {
		if policy == "root" { //deliberately skipping root policy
			continue
		}
		content, err := sys.GetPolicy(policy)
		if err != nil {
			panic(err) //todo:
		}
		callback(policy, content)
	}
}
