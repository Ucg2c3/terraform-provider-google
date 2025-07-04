---
# ----------------------------------------------------------------------------
#
#     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
#
# ----------------------------------------------------------------------------
#
#     This code is generated by Magic Modules using the following:
#
#     Configuration: https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/apigee/Organization.yaml
#     Template:      https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.html.markdown.tmpl
#
#     DO NOT EDIT this file directly. Any changes made to this file will be
#     overwritten during the next generation cycle.
#
# ----------------------------------------------------------------------------
subcategory: "Apigee"
description: |-
  An `Organization` is the top-level container in Apigee.
---

# google_apigee_organization

An `Organization` is the top-level container in Apigee.


To get more information about Organization, see:

* [API documentation](https://cloud.google.com/apigee/docs/reference/apis/apigee/rest/v1/organizations)
* How-to Guides
    * [Creating an API organization](https://cloud.google.com/apigee/docs/api-platform/get-started/create-org)
    * [Setting a custom endpoint (required for data residency)](https://registry.terraform.io/providers/hashicorp/google/latest/docs/guides/provider_reference#advanced-settings-configuration)

## Example Usage - Apigee Organization Cloud Basic


```hcl
data "google_client_config" "current" {}

resource "google_compute_network" "apigee_network" {
  name       = "apigee-network"
}

resource "google_compute_global_address" "apigee_range" {
  name          = "apigee-range"
  purpose       = "VPC_PEERING"
  address_type  = "INTERNAL"
  prefix_length = 16
  network       = google_compute_network.apigee_network.id
}

resource "google_service_networking_connection" "apigee_vpc_connection" {
  network                 = google_compute_network.apigee_network.id
  service                 = "servicenetworking.googleapis.com"
  reserved_peering_ranges = [google_compute_global_address.apigee_range.name]
}

resource "google_apigee_organization" "org" {
  analytics_region   = "us-central1"
  project_id         = data.google_client_config.current.project
  authorized_network = google_compute_network.apigee_network.id
  depends_on         = [google_service_networking_connection.apigee_vpc_connection]
}
```
## Example Usage - Apigee Organization Cloud Basic Disable Vpc Peering


```hcl
data "google_client_config" "current" {}

resource "google_apigee_organization" "org" {
  description         = "Terraform-provisioned basic Apigee Org without VPC Peering."
  analytics_region    = "us-central1"
  project_id          = data.google_client_config.current.project
  disable_vpc_peering = true
}
```
## Example Usage - Apigee Organization Cloud Basic Data Residency


```hcl
provider "google" {
  apigee_custom_endpoint = "https://eu-apigee.googleapis.com/v1/"
}

data "google_client_config" "current" {}

resource "google_apigee_organization" "org" {
  description                = "Terraform-provisioned basic Apigee Org under European Union hosting jurisdiction."
  project_id                 = data.google_client_config.current.project
  api_consumer_data_location = "europe-west1"
  billing_type               = "PAYG"
  disable_vpc_peering        = true
}
```
## Example Usage - Apigee Organization Cloud Full


```hcl
data "google_client_config" "current" {}

resource "google_compute_network" "apigee_network" {
  name       = "apigee-network"
}

resource "google_compute_global_address" "apigee_range" {
  name          = "apigee-range"
  purpose       = "VPC_PEERING"
  address_type  = "INTERNAL"
  prefix_length = 16
  network       = google_compute_network.apigee_network.id
}

resource "google_service_networking_connection" "apigee_vpc_connection" {
  network                 = google_compute_network.apigee_network.id
  service                 = "servicenetworking.googleapis.com"
  reserved_peering_ranges = [google_compute_global_address.apigee_range.name]
}

resource "google_kms_key_ring" "apigee_keyring" {
  name     = "apigee-keyring"
  location = "us-central1"
}

resource "google_kms_crypto_key" "apigee_key" {
  name            = "apigee-key"
  key_ring        = google_kms_key_ring.apigee_keyring.id

  lifecycle {
    prevent_destroy = true
  }
}

resource "google_project_service_identity" "apigee_sa" {
  provider = google-beta
  project  = google_project.project.project_id
  service  = google_project_service.apigee.service
}

resource "google_kms_crypto_key_iam_member" "apigee_sa_keyuser" {
  crypto_key_id = google_kms_crypto_key.apigee_key.id
  role          = "roles/cloudkms.cryptoKeyEncrypterDecrypter"

  member = google_project_service_identity.apigee_sa.member
}

resource "google_apigee_organization" "org" {
  analytics_region                     = "us-central1"
  display_name                         = "apigee-org"
  description                          = "Terraform-provisioned Apigee Org."
  project_id                           = data.google_client_config.current.project
  authorized_network                   = google_compute_network.apigee_network.id
  runtime_database_encryption_key_name = google_kms_crypto_key.apigee_key.id

  depends_on = [
    google_service_networking_connection.apigee_vpc_connection,
    google_kms_crypto_key_iam_member.apigee_sa_keyuser,
  ]
}
```
## Example Usage - Apigee Organization Cloud Full Disable Vpc Peering


```hcl
data "google_client_config" "current" {}

resource "google_kms_key_ring" "apigee_keyring" {
  name     = "apigee-keyring"
  location = "us-central1"
}

resource "google_kms_crypto_key" "apigee_key" {
  name            = "apigee-key"
  key_ring        = google_kms_key_ring.apigee_keyring.id

  lifecycle {
    prevent_destroy = true
  }
}

resource "google_project_service_identity" "apigee_sa" {
  provider = google-beta
  project  = google_project.project.project_id
  service  = google_project_service.apigee.service
}

resource "google_kms_crypto_key_iam_member" "apigee_sa_keyuser" {
  crypto_key_id = google_kms_crypto_key.apigee_key.id
  role          = "roles/cloudkms.cryptoKeyEncrypterDecrypter"

  member = google_project_service_identity.apigee_sa.member
}

resource "google_apigee_organization" "org" {
  analytics_region                     = "us-central1"
  display_name                         = "apigee-org"
  description                          = "Terraform-provisioned Apigee Org without VPC Peering."
  project_id                           = data.google_client_config.current.project
  disable_vpc_peering                  = true
  runtime_database_encryption_key_name = google_kms_crypto_key.apigee_key.id

  depends_on = [
    google_kms_crypto_key_iam_member.apigee_sa_keyuser,
  ]
}
```

## Argument Reference

The following arguments are supported:


* `project_id` -
  (Required)
  The project ID associated with the Apigee organization.


* `display_name` -
  (Optional)
  The display name of the Apigee organization.

* `description` -
  (Optional)
  Description of the Apigee organization.

* `analytics_region` -
  (Optional)
  Primary GCP region for analytics data storage. For valid values, see [Create an Apigee organization](https://cloud.google.com/apigee/docs/api-platform/get-started/create-org).

* `api_consumer_data_location` -
  (Optional)
  This field is needed only for customers using non-default data residency regions.
  Apigee stores some control plane data only in single region.
  This field determines which single region Apigee should use.

* `api_consumer_data_encryption_key_name` -
  (Optional)
  Cloud KMS key name used for encrypting API consumer data.

* `control_plane_encryption_key_name` -
  (Optional)
  Cloud KMS key name used for encrypting control plane data that is stored in a multi region.
  Only used for the data residency region "US" or "EU".

* `authorized_network` -
  (Optional)
  Compute Engine network used for Service Networking to be peered with Apigee runtime instances.
  See [Getting started with the Service Networking API](https://cloud.google.com/service-infrastructure/docs/service-networking/getting-started).
  Valid only when `RuntimeType` is set to CLOUD. The value can be updated only when there are no runtime instances. For example: "default".

* `disable_vpc_peering` -
  (Optional)
  Flag that specifies whether the VPC Peering through Private Google Access should be
  disabled between the consumer network and Apigee. Required if an `authorizedNetwork`
  on the consumer project is not provided, in which case the flag should be set to `true`.
  Valid only when `RuntimeType` is set to CLOUD. The value must be set before the creation
  of any Apigee runtime instance and can be updated only when there are no runtime instances.

* `runtime_type` -
  (Optional)
  Runtime type of the Apigee organization based on the Apigee subscription purchased.
  Default value is `CLOUD`.
  Possible values are: `CLOUD`, `HYBRID`.

* `billing_type` -
  (Optional)
  Billing type of the Apigee organization. See [Apigee pricing](https://cloud.google.com/apigee/pricing).

* `runtime_database_encryption_key_name` -
  (Optional)
  Cloud KMS key name used for encrypting the data that is stored and replicated across runtime instances.
  Update is not allowed after the organization is created.
  If not specified, a Google-Managed encryption key will be used.
  Valid only when `RuntimeType` is CLOUD. For example: `projects/foo/locations/us/keyRings/bar/cryptoKeys/baz`.

* `properties` -
  (Optional)
  Properties defined in the Apigee organization profile.
  Structure is [documented below](#nested_properties).

* `retention` -
  (Optional)
  Optional. This setting is applicable only for organizations that are soft-deleted (i.e., BillingType
  is not EVALUATION). It controls how long Organization data will be retained after the initial delete
  operation completes. During this period, the Organization may be restored to its last known state.
  After this period, the Organization will no longer be able to be restored.
  Default value is `DELETION_RETENTION_UNSPECIFIED`.
  Possible values are: `DELETION_RETENTION_UNSPECIFIED`, `MINIMUM`.



<a name="nested_properties"></a>The `properties` block supports:

* `property` -
  (Optional)
  List of all properties in the object.
  Structure is [documented below](#nested_properties_property).


<a name="nested_properties_property"></a>The `property` block supports:

* `name` -
  (Optional)
  Name of the property.

* `value` -
  (Optional)
  Value of the property.

## Attributes Reference

In addition to the arguments listed above, the following computed attributes are exported:

* `id` - an identifier for the resource with format `organizations/{{name}}`

* `name` -
  Output only. Name of the Apigee organization.

* `subscription_type` -
  Output only. Subscription type of the Apigee organization.
  Valid values include trial (free, limited, and for evaluation purposes only) or paid (full subscription has been purchased).

* `ca_certificate` -
  Output only. Base64-encoded public certificate for the root CA of the Apigee organization.
  Valid only when `RuntimeType` is CLOUD. A base64-encoded string.

* `apigee_project_id` -
  Output only. Project ID of the Apigee Tenant Project.


## Timeouts

This resource provides the following
[Timeouts](https://developer.hashicorp.com/terraform/plugin/sdkv2/resources/retries-and-customizable-timeouts) configuration options:

- `create` - Default is 45 minutes.
- `update` - Default is 45 minutes.
- `delete` - Default is 45 minutes.

## Import


Organization can be imported using any of these accepted formats:

* `organizations/{{name}}`
* `{{name}}`


In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import Organization using one of the formats above. For example:

```tf
import {
  id = "organizations/{{name}}"
  to = google_apigee_organization.default
}
```

When using the [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import), Organization can be imported using one of the formats above. For example:

```
$ terraform import google_apigee_organization.default organizations/{{name}}
$ terraform import google_apigee_organization.default {{name}}
```
