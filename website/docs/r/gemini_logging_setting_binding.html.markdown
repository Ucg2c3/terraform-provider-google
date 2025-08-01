---
# ----------------------------------------------------------------------------
#
#     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
#
# ----------------------------------------------------------------------------
#
#     This code is generated by Magic Modules using the following:
#
#     Configuration: https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/gemini/LoggingSettingBinding.yaml
#     Template:      https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.html.markdown.tmpl
#
#     DO NOT EDIT this file directly. Any changes made to this file will be
#     overwritten during the next generation cycle.
#
# ----------------------------------------------------------------------------
subcategory: "Gemini for Google Cloud"
description: |-
  The resource for managing Logging setting bindings for Admin Control.
---

# google_gemini_logging_setting_binding

The resource for managing Logging setting bindings for Admin Control.


To get more information about LoggingSettingBinding, see:
* How-to Guides
    * [Gemini Cloud Assist overview](https://cloud.google.com/gemini/docs/cloud-assist/overview)

## Example Usage - Gemini Logging Setting Binding Basic


```hcl
resource "google_gemini_logging_setting" "basic" {
    logging_setting_id = "ls-tf1"
    location = "global"
    labels = {"my_key": "my_value"}
    log_prompts_and_responses = true
}

resource "google_gemini_logging_setting_binding" "example" {
    logging_setting_id = google_gemini_logging_setting.basic.logging_setting_id
    setting_binding_id = "ls-tf1b1"
    location = "global"
    target = "projects/980109375338"
}
```

## Argument Reference

The following arguments are supported:


* `target` -
  (Required)
  Target of the binding.

* `logging_setting_id` -
  (Required)
  Resource ID segment making up resource `name`. It identifies the resource within its parent collection as described in https://google.aip.dev/122.

* `setting_binding_id` -
  (Required)
  Id of the setting binding.


* `labels` -
  (Optional)
  Labels as key value pairs.
  **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
  Please refer to the field `effective_labels` for all of the labels present on the resource.

* `product` -
  (Optional)
  Product type of the setting binding.
  Possible values are: `GEMINI_CODE_ASSIST`.

* `location` -
  (Optional)
  Resource ID segment making up resource `name`. It identifies the resource within its parent collection as described in https://google.aip.dev/122.

* `project` - (Optional) The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.



## Attributes Reference

In addition to the arguments listed above, the following computed attributes are exported:

* `id` - an identifier for the resource with format `projects/{{project}}/locations/{{location}}/loggingSettings/{{logging_setting_id}}/settingBindings/{{setting_binding_id}}`

* `name` -
  Identifier. Name of the resource.
  Format:projects/{project}/locations/{location}/loggingSettings/{setting}/settingBindings/{setting_binding}

* `create_time` -
  Create time stamp.

* `update_time` -
  Update time stamp.

* `terraform_labels` -
  The combination of labels configured directly on the resource
   and default labels configured on the provider.

* `effective_labels` -
  All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other clients and services.


## Timeouts

This resource provides the following
[Timeouts](https://developer.hashicorp.com/terraform/plugin/sdkv2/resources/retries-and-customizable-timeouts) configuration options:

- `create` - Default is 90 minutes.
- `update` - Default is 90 minutes.
- `delete` - Default is 90 minutes.

## Import


LoggingSettingBinding can be imported using any of these accepted formats:

* `projects/{{project}}/locations/{{location}}/loggingSettings/{{logging_setting_id}}/settingBindings/{{setting_binding_id}}`
* `{{project}}/{{location}}/{{logging_setting_id}}/{{setting_binding_id}}`
* `{{location}}/{{logging_setting_id}}/{{setting_binding_id}}`


In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import LoggingSettingBinding using one of the formats above. For example:

```tf
import {
  id = "projects/{{project}}/locations/{{location}}/loggingSettings/{{logging_setting_id}}/settingBindings/{{setting_binding_id}}"
  to = google_gemini_logging_setting_binding.default
}
```

When using the [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import), LoggingSettingBinding can be imported using one of the formats above. For example:

```
$ terraform import google_gemini_logging_setting_binding.default projects/{{project}}/locations/{{location}}/loggingSettings/{{logging_setting_id}}/settingBindings/{{setting_binding_id}}
$ terraform import google_gemini_logging_setting_binding.default {{project}}/{{location}}/{{logging_setting_id}}/{{setting_binding_id}}
$ terraform import google_gemini_logging_setting_binding.default {{location}}/{{logging_setting_id}}/{{setting_binding_id}}
```

## User Project Overrides

This resource supports [User Project Overrides](https://registry.terraform.io/providers/hashicorp/google/latest/docs/guides/provider_reference#user_project_override).
