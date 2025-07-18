---
# ----------------------------------------------------------------------------
#
#     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
#
# ----------------------------------------------------------------------------
#
#     This code is generated by Magic Modules using the following:
#
#     Configuration: https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/apigee/EnvgroupAttachment.yaml
#     Template:      https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.html.markdown.tmpl
#
#     DO NOT EDIT this file directly. Any changes made to this file will be
#     overwritten during the next generation cycle.
#
# ----------------------------------------------------------------------------
subcategory: "Apigee"
description: |-
  An `Environment Group attachment` in Apigee.
---

# google_apigee_envgroup_attachment

An `Environment Group attachment` in Apigee.


To get more information about EnvgroupAttachment, see:

* [API documentation](https://cloud.google.com/apigee/docs/reference/apis/apigee/rest/v1/organizations.envgroups.attachments/create)
* How-to Guides
    * [Creating an environment](https://cloud.google.com/apigee/docs/api-platform/get-started/create-environment)

## Example Usage - Apigee Environment Group Attachment Basic


```hcl
resource "google_project" "project" {
  project_id      = "my-project"
  name            = "my-project"
  org_id          = "123456789"
  billing_account = "000000-0000000-0000000-000000"
  deletion_policy = "DELETE"
}

resource "google_project_service" "apigee" {
  project = google_project.project.project_id
  service = "apigee.googleapis.com"
}

resource "google_project_service" "compute" {
  project = google_project.project.project_id
  service = "compute.googleapis.com"
}

resource "google_project_service" "servicenetworking" {
  project = google_project.project.project_id
  service = "servicenetworking.googleapis.com"
}

resource "google_compute_network" "apigee_network" {
  name       = "apigee-network"
  project    = google_project.project.project_id
  depends_on = [google_project_service.compute]
}

resource "google_compute_global_address" "apigee_range" {
  name          = "apigee-range"
  purpose       = "VPC_PEERING"
  address_type  = "INTERNAL"
  prefix_length = 16
  network       = google_compute_network.apigee_network.id
  project       = google_project.project.project_id
}

resource "google_service_networking_connection" "apigee_vpc_connection" {
  network                 = google_compute_network.apigee_network.id
  service                 = "servicenetworking.googleapis.com"
  reserved_peering_ranges = [google_compute_global_address.apigee_range.name]
  depends_on              = [google_project_service.servicenetworking]
}

resource "google_apigee_organization" "apigee_org" {
  analytics_region   = "us-central1"
  project_id         = google_project.project.project_id
  authorized_network = google_compute_network.apigee_network.id
  depends_on         = [
    google_service_networking_connection.apigee_vpc_connection,
    google_project_service.apigee,
  ]
}

resource "google_apigee_envgroup" "apigee_envgroup" {
  org_id    = google_apigee_organization.apigee_org.id
  name      = "my-envgroup"
  hostnames = ["abc.foo.com"]
}

resource "google_apigee_environment" "apigee_env" {
  org_id       = google_apigee_organization.apigee_org.id
  name         = "my-environment"
  description  = "Apigee Environment"
  display_name = "my-environment"
}

resource "google_apigee_envgroup_attachment" "" {
  envgroup_id  = google_apigee_envgroup.apigee_envgroup.id
  environment  = google_apigee_environment.apigee_env.name
}
```

## Argument Reference

The following arguments are supported:


* `environment` -
  (Required)
  The resource ID of the environment.

* `envgroup_id` -
  (Required)
  The Apigee environment group associated with the Apigee environment,
  in the format `organizations/{{org_name}}/envgroups/{{envgroup_name}}`.




## Attributes Reference

In addition to the arguments listed above, the following computed attributes are exported:

* `id` - an identifier for the resource with format `{{envgroup_id}}/attachments/{{name}}`

* `name` -
  The name of the newly created  attachment (output parameter).


## Timeouts

This resource provides the following
[Timeouts](https://developer.hashicorp.com/terraform/plugin/sdkv2/resources/retries-and-customizable-timeouts) configuration options:

- `create` - Default is 30 minutes.
- `delete` - Default is 30 minutes.

## Import


EnvgroupAttachment can be imported using any of these accepted formats:

* `{{envgroup_id}}/attachments/{{name}}`
* `{{envgroup_id}}/{{name}}`


In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import EnvgroupAttachment using one of the formats above. For example:

```tf
import {
  id = "{{envgroup_id}}/attachments/{{name}}"
  to = google_apigee_envgroup_attachment.default
}
```

When using the [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import), EnvgroupAttachment can be imported using one of the formats above. For example:

```
$ terraform import google_apigee_envgroup_attachment.default {{envgroup_id}}/attachments/{{name}}
$ terraform import google_apigee_envgroup_attachment.default {{envgroup_id}}/{{name}}
```
