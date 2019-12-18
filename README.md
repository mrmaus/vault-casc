# Vault Configuration-as-Code

Simple CLI utility that enables managing Hashicorp Vault configuration as files outside of Vault.

## Global Configuration Flags

Name | Default Value | Description 
--- | --- | ---
vault-address | http://localhost:8200 | The address of Vault instance
vault-token | | Vault access token
dir | ./ | The local configuration folder


## Creating Local Config Repository
First step in the process would be creating local configuration folder. The folder will contain
all Vault configuration files structures according to Vault internal layout. 
You'd need a running Vault instance for this step since by default it already has some default
configs that will be copied to local folder on the initialization.

Example command:
```
vault-casc init --dir ./casc --vault-address http://localhost:8200 --vault-token=TOKEN
```
Which will create new local configuration repository within ./casc folder.
Valid vault address and token must be provided

## Applying Local Config to Vault
Run the 'apply' command to copy local configuration into specified vault instance, for example:
```
vault-casc apply --dir ./casc --vault-address http://localhost:8200 --vault-token=TOKEN
```
Valid vault address and token must be provided
The command will traverse through all configuration files and apply them to the corresponding Vault endpoints.
Please note that any existing changes in Vault will be overwritten without any warning

## Supported Vault Configuration
The following Vault configuration items are currently supported by the tool:

Name | Location | Details
--- | --- | ---
Policy | /sys/policy | One 'hcl' file created per policy


## TODOs:
- Support for LDAP configuration
- Support for LDAP login
- Removal of missing policies from Vault