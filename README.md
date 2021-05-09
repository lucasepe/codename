# Codename

> an RFC1178 implementation to generate pronounceable, sometimes even memorable, _"superheroe like"_ codenames, consisting of a random combination of adjective and a noun.


## Usage

Codename it's a [package](https://golang.org/doc/code#ImportingRemote), so all you need is all you need to do is import it into your code

```sh
$ git clone https://github.com/lucasepe/terraform-provider-codename
```

Enter the provider directory:

```sh
$ cd terraform-provider-codename
```

Build it:

```sh
$ make build
```

To install the plugin locally, type:

```sh
$ make install
```

## Usage

Define a provider requirement:

```hcl
terraform {
  required_providers {
    codename = {
      version = ">= 0.1.0"
      source  = "github.com/lucasepe/codename"
    }
  }
}
```

Declare the provider:

```hcl
provider "codename" {}
```

Use the `codename` resource defined in this provider:

```hcl
resource "codename" "example1" {
  snakefy      = true
  token_length = 4
}

resource "codename" "example2" {
  prefix = "it->"
}

output "codename1" {
  value = codename.example1.id
}

output "codename2" {
  value = codename.example2.id
}
```

Initialize a working directory:

```sh
$ terraform init
Initializing the backend...

Initializing provider plugins...
- Finding github.com/lucasepe/codename versions matching ">= 0.1.0"...
- Installing github.com/lucasepe/codename v0.1.0...
- Installed github.com/lucasepe/codename v0.1.0 (unauthenticated)

...
```

Execute the terraform script:

```sh
$ terraform apply -auto-approve
...
Apply complete! Resources: 2 added, 0 changed, 0 destroyed.

Outputs:

codename1 = "adequate_retro_girl_8157"
codename2 = "it->still-karate"
```