package vault

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/vault-plugin-secrets-azure"
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
			"name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: "Unique name for the role."
			},
			"backend": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: "The path of the Azure Secret Backend the role belongs to."
			}
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
	client := meta.(*api.Client)

	backend := d.Get("backend").(string)
	name := d.Get("name").(string)
	applicationObjectID := d.Get("application_object_id").(string)
	ttl := d.Get("ttl").(string)
	maxTtl := d.Get("max_ttl").(string)

	azureRoles, err := expandAzureRoles(d.Get("azure_roles").([]interface{}))
	if err != nil {
		return fmt.Errorf("Failed to exapand azure roles: %s", err)
	}

	if applicationObjectID != "" {
		data["application_object_id"] = applicationObjectID
	}

	if ttl != "" {
		data["ttl"] = ttl
	}

	if maxTtl != "" {
		data["max_ttl"] = maxTtl
	}

	data := path_roles.Role{
		// Statically assigning as thats how it is done in the api client
		CredentialType: 0,
		AzureRoles: *azureRoles,
		// Commenting out this as not required at this stage
		//ApplicationID: ,
		ApplicationObjectID: applicationObjectID,
		TTL: ttl,
		MaxTTL: maxTtl,
	}

	_, err := client.Logical().Write(backend+"/roles/"+name, data)
	if err != nil {
		return fmt.Errorf("error creating role %q for backend %q: %s", name, backend, err)
	}

	d.SetId(backend + "/roles/" + name)
	return azureSecretBackendRoleRead(d, meta)
}

func azureSecretBackendRoleRead(d *schema.ResourceData, meta interface{}) error {

}

func azureSecretBackendRoleUpdate(d *schema.ResourceData, meta interface{}) error {

}

func azureSecretBackendRoleDelete(d *schema.ResourceData, meta interface{}) error {

}

func azureSecretBackendRoleExists(d *schema.ResourceData, meta interface{}) (bool, error) {

}

func flattenAzureRoles(in []path_roles.azureRole, d *schema.ResourceData) ([]interface{}, error) {
	att := make(map[string]interface{})

	if in.RoleName != nil {
		att["role_name"] = *in.RoleName
	}

	if in.RoleID != nil {
		att["role_id"] = *in.RoleID
	}

	att["scope"] = in.Scope

	if err != nil {
		return nil, err
	}

	return []interface{}{att}, nil
}

func expandAzureRoles(azureRole []interface{}) (*path_roles.azureRole, error) {
	obj := &path_roles.azureRole{}

	if len(azureRole) == 0 || azureRole[0] == nil {
		return obj, nil
	}

	in := azureRole[0].(map[string]interface{})

	obj.RoleName = in["role_name"]
	obj.RoleID = in["role_id"]
	obj.Scope = in["scope"]

	return obj, nil
}