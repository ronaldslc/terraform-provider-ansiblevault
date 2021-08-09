module github.com/ronaldslc/terraform-provider-ansiblevault/v2

go 1.12

require (
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.7.0
	github.com/sosedoff/ansible-vault-go v0.0.0-20201201002713-782dc5c40224
	golang.org/x/crypto v0.0.0-20210711020723-a769d52b0f97 // indirect
	gopkg.in/yaml.v2 v2.4.0
)

// replace git.apache.org/thrift.git => github.com/apache/thrift v0.0.0-20180902110319-2566ecd5d999
