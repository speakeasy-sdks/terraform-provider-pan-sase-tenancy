<div align="center">
    <img src="https://github.com/speakeasy-sdks/pan-sase-tenancy-ts/assets/6267663/e0205c2a-fa61-4b1f-aaa0-599896e022da" width="300">
    <h1>SASE Tenancy Terraform Provider</h1>
   <p>Containers used to build your tenant hierachy.</p>
   <a href="https://pan.dev/category/sase/api/tenancy/tenant-service-group/"><img src="https://img.shields.io/static/v1?label=Docs&message=API Ref&color=000&style=for-the-badge" /></a>
  <a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/License-MIT-blue.svg?style=for-the-badge" /></a>
</div>

<!-- Start SDK Installation -->
## Notice of Alpha Status

This is a Alpha release of the pan-sase-tenancy Terraform Provider.

It is under active development and you may experience breaking changes. Please pin to a version if trialing in production.

## SDK Installation

To install this provider, copy and paste this code into your Terraform configuration. Then, run `terraform init`.

```hcl
terraform {
  required_providers {
    pan-sase-tenancy = {
      source  = "speakeasy/pan-sase-tenancy"
      version = "0.0.1"
    }
  }
}

provider "pan-sase-tenancy" {
  # Configuration options
}
```
<!-- End SDK Installation -->

<!-- Start SDK Example Usage -->
## Testing the provider locally

Should you want to validate a change locally, the `--debug` flag allows you to execute the provider against a terraform instance locally.

This also allows for debuggers (e.g. delve) to be attached to the provider.

```sh
go run main.go --debug
# Copy the TF_REATTACH_PROVIDERS env var
# In a new terminal
cd examples/your-example
TF_REATTACH_PROVIDERS=... terraform init
TF_REATTACH_PROVIDERS=... terraform apply
```
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->

<!-- End SDK Available Operations -->

Terraform allows you to use local provider builds by setting a `dev_overrides` block in a configuration file called `.terraformrc`. This block overrides all other configured installation methods.

Terraform searches for the `.terraformrc` file in your home directory and applies any configuration settings you set.

```
provider_installation {

  dev_overrides {
      "registry.terraform.io/terraform/scaffolding" = "<PATH>"
  }

  # For all other providers, install them directly from their origin provider
  # registries as normal. If you omit this, Terraform will _only_ use
  # the dev_overrides block, and so no other providers will be available.
  direct {}
}
```

Your `<PATH>` may vary depending on how your Go environment variables are configured. Execute `go env GOBIN` to set it, then set the `<PATH>` to the value returned. If nothing is returned, set it to the default location, `$HOME/go/bin`.

### Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release !

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
