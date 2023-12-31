# CLI
## Commands
#### init
```
s init <prefix> <region> --profile <profile_name>
```
- prefix: The prefix for the project your creating
- region: The AWS region to create the project in
- profile: The AWS profile to use, uses default when not provided

Initialises s in your repo.

##### add
```
s add <name> <value>
```
- name: The name of the secret
- value: The value of the secret

Adds a secret.

#### update
```
s update <name> <value>
```
- name: The name of the secret
- value: The value of the secret

Updates a secret.

#### list
```
s list
```
Returns a list of all secrets.

#### get
```
s get <name> <value>
```
- name: The name of the secret
- value: The value of the secret

Gets a secret.

#### delete
```
s delete <name>
```
- name: The name of the secret

Deletes a secret.

#### export
```
eval $(s export)
```
Evalutes the output of s for environment variables.

#### help
```
s -h | --help
```
Shows the helper for s.

## How to use
For now I don't use have any official way of distributing s. The latest version
can be downloaded from the releases page or you can clone the repo and build from
source. The command for building can be seen below:
```
go build -o build/s
```
