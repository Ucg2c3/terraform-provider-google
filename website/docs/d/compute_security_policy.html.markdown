---
# ----------------------------------------------------------------------------
#
#     ***     AUTO GENERATED CODE    ***    Type: Handwritten     ***
#
# ----------------------------------------------------------------------------
#
#     This code is generated by Magic Modules using the following:
#
#     Source file: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/third_party/terraform/website/docs/d/compute_security_policy.html.markdown
#
#     DO NOT EDIT this file directly. Any changes made to this file will be
#     overwritten during the next generation cycle.
#
# ----------------------------------------------------------------------------
subcategory: "Compute Engine"
description: |-
  Get information about a Google Compute Security Policy.
---

# google_compute_security_policy

To get more information about Google Compute Security Policy, see:

* [API documentation](https://cloud.google.com/compute/docs/reference/rest/beta/securityPolicies)
* How-to Guides
    * [Official Documentation](https://cloud.google.com/armor/docs/configure-security-policies)

## Example Usage

```hcl
data "google_compute_security_policy" "sp1" {
  name = "my-policy"
  project = "my-project"
}

data "google_compute_security_policy" "sp2" {
  self_link = "https://www.googleapis.com/compute/v1/projects/my-project/global/securityPolicies/my-policy"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) The name of the security policy. Provide either this or a `self_link`.

* `project` - (Optional) The project in which the resource belongs. If it is not provided, the provider project is used.

* `self_link` - (Optional) The self_link of the security policy. Provide either this or a `name`

## Attributes Reference

See [google_compute_security_policy](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/compute_security_policy) resource for details of the available attributes.
