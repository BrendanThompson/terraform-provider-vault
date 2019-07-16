module github.com/terraform-providers/terraform-provider-vault

go 1.12

replace github.com/hashicorp/vault-plugin-secrets-azure => ../../hashicorp/vault-plugin-secrets-azure

require (
	cloud.google.com/go v0.37.4 // indirect
	github.com/aws/aws-sdk-go v1.19.18
	github.com/gosimple/slug v1.4.1
	github.com/hashicorp/go-cleanhttp v0.5.1
	github.com/hashicorp/terraform v0.12.2
	github.com/hashicorp/vault v1.1.2
	github.com/hashicorp/vault-plugin-secrets-azure v0.5.1
	github.com/mitchellh/go-homedir v1.1.0
	github.com/rainycape/unidecode v0.0.0-20150907023854-cb7f23ec59be // indirect
	github.com/ulikunitz/xz v0.5.6 // indirect
	golang.org/x/oauth2 v0.0.0-20190402181905-9f3314589c9a // indirect
	google.golang.org/api v0.3.2 // indirect
	google.golang.org/grpc v1.20.0 // indirect
)
