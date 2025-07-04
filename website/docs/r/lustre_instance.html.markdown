---
# ----------------------------------------------------------------------------
#
#     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
#
# ----------------------------------------------------------------------------
#
#     This code is generated by Magic Modules using the following:
#
#     Configuration: https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/lustre/Instance.yaml
#     Template:      https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.html.markdown.tmpl
#
#     DO NOT EDIT this file directly. Any changes made to this file will be
#     overwritten during the next generation cycle.
#
# ----------------------------------------------------------------------------
subcategory: "Google Cloud Managed Lustre"
description: |-
  A Managed Lustre instance
---

# google_lustre_instance

A Managed Lustre instance


To get more information about Instance, see:

* [API documentation](https://cloud.google.com/managed-lustre/docs/reference/rest/v1/projects.locations.instances)
* How-to Guides
    * [Official Documentation](https://cloud.google.com/managed-lustre/docs/create-instance)

<div class = "oics-button" style="float: right; margin: 0 0 -15px">
  <a href="https://console.cloud.google.com/cloudshell/open?cloudshell_git_repo=https%3A%2F%2Fgithub.com%2Fterraform-google-modules%2Fdocs-examples.git&cloudshell_image=gcr.io%2Fcloudshell-images%2Fcloudshell%3Alatest&cloudshell_print=.%2Fmotd&cloudshell_tutorial=.%2Ftutorial.md&cloudshell_working_dir=lustre_instance_basic&open_in_editor=main.tf" target="_blank">
    <img alt="Open in Cloud Shell" src="//gstatic.com/cloudssh/images/open-btn.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
  </a>
</div>
## Example Usage - Lustre Instance Basic


```hcl
resource "google_lustre_instance" "instance" {
  instance_id                 = "my-instance"
  location                    = "us-central1-a"
  description                 = "test lustre instance"
  filesystem                  = "testfs"
  capacity_gib                = 18000
  network                     = data.google_compute_network.lustre-network.id
  per_unit_storage_throughput = 1000
  labels                      = {
    test = "value"
  }
  timeouts {
		create = "120m"
	}
}

// This example assumes this network already exists.
// The API creates a tenant network per network authorized for a
// Lustre instance and that network is not deleted when the user-created
// network (authorized_network) is deleted, so this prevents issues
// with tenant network quota.
// If this network hasn't been created and you are using this example in your
// config, add an additional network resource or change
// this from "data"to "resource"
data "google_compute_network" "lustre-network" {
  name = "my-network"
}
```

## Argument Reference

The following arguments are supported:


* `capacity_gib` -
  (Required)
  The storage capacity of the instance in gibibytes (GiB). Allowed values
  are from `18000` to `954000`, in increments of 9000.

* `filesystem` -
  (Required)
  The filesystem name for this instance. This name is used by client-side
  tools, including when mounting the instance. Must be eight characters or
  less and can only contain letters and numbers.

* `network` -
  (Required)
  The full name of the VPC network to which the instance is connected.
  Must be in the format
  `projects/{project_id}/global/networks/{network_name}`.

* `per_unit_storage_throughput` -
  (Required)
  The throughput of the instance in MB/s/TiB.
  Valid values are 125, 250, 500, 1000.

* `location` -
  (Required)
  Resource ID segment making up resource `name`. It identifies the resource within its parent collection as described in https://google.aip.dev/122.

* `instance_id` -
  (Required)
  The name of the Managed Lustre instance.
  * Must contain only lowercase letters, numbers, and hyphens.
  * Must start with a letter.
  * Must be between 1-63 characters.
  * Must end with a number or a letter.


* `gke_support_enabled` -
  (Optional)
  Indicates whether you want to enable support for GKE clients. By default,
  GKE clients are not supported.

* `description` -
  (Optional)
  A user-readable description of the instance.

* `labels` -
  (Optional)
  Labels as key value pairs.
  **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
  Please refer to the field `effective_labels` for all of the labels present on the resource.

* `project` - (Optional) The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.



## Attributes Reference

In addition to the arguments listed above, the following computed attributes are exported:

* `id` - an identifier for the resource with format `projects/{{project}}/locations/{{location}}/instances/{{instance_id}}`

* `update_time` -
  Timestamp when the instance was last updated.

* `state` -
  The state of the instance.
  Possible values:
  STATE_UNSPECIFIED
  ACTIVE
  CREATING
  DELETING
  UPGRADING
  REPAIRING
  STOPPED

* `mount_point` -
  Mount point of the instance in the format `IP_ADDRESS@tcp:/FILESYSTEM`.

* `create_time` -
  Timestamp when the instance was created.

* `name` -
  Identifier. The name of the instance.

* `terraform_labels` -
  The combination of labels configured directly on the resource
   and default labels configured on the provider.

* `effective_labels` -
  All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other clients and services.


## Timeouts

This resource provides the following
[Timeouts](https://developer.hashicorp.com/terraform/plugin/sdkv2/resources/retries-and-customizable-timeouts) configuration options:

- `create` - Default is 20 minutes.
- `update` - Default is 20 minutes.
- `delete` - Default is 20 minutes.

## Import


Instance can be imported using any of these accepted formats:

* `projects/{{project}}/locations/{{location}}/instances/{{instance_id}}`
* `{{project}}/{{location}}/{{instance_id}}`
* `{{location}}/{{instance_id}}`


In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import Instance using one of the formats above. For example:

```tf
import {
  id = "projects/{{project}}/locations/{{location}}/instances/{{instance_id}}"
  to = google_lustre_instance.default
}
```

When using the [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import), Instance can be imported using one of the formats above. For example:

```
$ terraform import google_lustre_instance.default projects/{{project}}/locations/{{location}}/instances/{{instance_id}}
$ terraform import google_lustre_instance.default {{project}}/{{location}}/{{instance_id}}
$ terraform import google_lustre_instance.default {{location}}/{{instance_id}}
```

## User Project Overrides

This resource supports [User Project Overrides](https://registry.terraform.io/providers/hashicorp/google/latest/docs/guides/provider_reference#user_project_override).
