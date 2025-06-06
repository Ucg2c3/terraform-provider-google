---
# ----------------------------------------------------------------------------
#
#     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
#
# ----------------------------------------------------------------------------
#
#     This code is generated by Magic Modules using the following:
#
#     Configuration: https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/notebooks/Instance.yaml
#     Template:      https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/datasource_iam.html.markdown.tmpl
#
#     DO NOT EDIT this file directly. Any changes made to this file will be
#     overwritten during the next generation cycle.
#
# ----------------------------------------------------------------------------
subcategory: "Cloud AI Notebooks"
description: |-
  A datasource to retrieve the IAM policy state for Cloud AI Notebooks Instance
---


# google_notebooks_instance_iam_policy

Retrieves the current IAM policy data for instance


## Example Usage


```hcl
data "google_notebooks_instance_iam_policy" "policy" {
  project = google_notebooks_instance.instance.project
  location = google_notebooks_instance.instance.location
  instance_name = google_notebooks_instance.instance.name
}
```

## Argument Reference

The following arguments are supported:

* `location` - (Optional) A reference to the zone where the machine resides. Used to find the parent resource to bind the IAM policy to. If not specified,
  the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
  location is specified, it is taken from the provider configuration.
* `instance_name` - (Required) Used to find the parent resource to bind the IAM policy to

* `project` - (Optional) The ID of the project in which the resource belongs.
    If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.

## Attributes Reference

The attributes are exported:

* `etag` - (Computed) The etag of the IAM policy.

* `policy_data` - (Required only by `google_notebooks_instance_iam_policy`) The policy data generated by
  a `google_iam_policy` data source.
