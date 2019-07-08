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
		Schema: map[string]*schema.Schema{},
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
