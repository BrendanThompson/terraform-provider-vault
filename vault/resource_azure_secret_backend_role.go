package vault

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func azureSecretBackendRoleResource() *schema.Resource {
	return &schema.Resource{
		Create: azureSecretBackendRoleCreate,
		Read:   azureSecretBackendRoleRead,
		Update: azureSecretBackendRoleUpdate,
		Delete: azureSecretBackendRoleDelete,
		Exists: azureSecretBackendRoleExists,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"role_name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: "The name of the Vault role."
			},
			"azure_roles": {
				Type: schema.TypeList,
				Required: true,
				Elem: &schema.Resource {
					Schema: map[string]*schema.Schema {
						"role_name": {
							Type: schema.TypeString,
							Optional: true,
							Description: "The case-insensitive name of the Azure role.",
						},
						"role_id": {
							Type: schema.TypeString,
							Optional: true,
							Description: "The explicit Azure role id",
						},
						"scope": {
							Type: schema.TypeString,
							Required: true,
							Description: "The scope of the Vault role",
						},
					},
				},

			},
			"application_object_id": {
				Type: schema.TypeString,
				Optional: true,
				Description: "The Application Object ID of the SP being used to provision the Vault role.",
			},
			"ttl": {
				Type: schema.TypeString,
				Optional: true,
			},
			"max_ttl": {
				Type: schema.TypeString,
				Optional: true,
			},
		},
	}
}

func azureSecretBackendRoleCreate(d *schema.ResourceData, meta interface{}) error {

}

func azureSecretBackendRoleRead(d *schema.ResourceData, meta interface{}) error {

}

func azureSecretBackendRoleUpdate(d *schema.ResourceData, meta interface{}) error {

}

func azureSecretBackendRoleDelete(d *schema.ResourceData, meta interface{}) error {

}

func azureSecretBackendRoleExists(d *schema.ResourceData, meta interface{}) (bool, error) {

}
