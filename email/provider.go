package email

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"email_email": resourceEmail(),
		},
		DataSourcesMap: map[string]*schema.Resource{},
	}
}
