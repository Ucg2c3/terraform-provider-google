---
# ----------------------------------------------------------------------------
#
#     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
#
# ----------------------------------------------------------------------------
#
#     This code is generated by Magic Modules using the following:
#
#     Configuration: https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/firebaseappcheck/DeviceCheckConfig.yaml
#     Template:      https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.html.markdown.tmpl
#
#     DO NOT EDIT this file directly. Any changes made to this file will be
#     overwritten during the next generation cycle.
#
# ----------------------------------------------------------------------------
subcategory: "Firebase App Check"
description: |-
  An app's DeviceCheck configuration object.
---

# google_firebase_app_check_device_check_config

An app's DeviceCheck configuration object. Note that the Team ID registered with your
app is used as part of the validation process. Make sure your `google_firebase_apple_app` has a team_id present.


To get more information about DeviceCheckConfig, see:

* [API documentation](https://firebase.google.com/docs/reference/appcheck/rest/v1/projects.apps.deviceCheckConfig)
* How-to Guides
    * [Official Documentation](https://firebase.google.com/docs/app-check)

~> **Warning:** All arguments including the following potentially sensitive
values will be stored in the raw state as plain text: `private_key`.
[Read more about sensitive data in state](https://www.terraform.io/language/state/sensitive-data).

## Example Usage - Firebase App Check Device Check Config Full


```hcl
resource "google_firebase_apple_app" "default" {
  provider = google-beta

  project      = "my-project-name"
  display_name = "Apple app"
  bundle_id    = "bundle.id.devicecheck"
  team_id      = "9987654321"
}

# It takes a while for App Check to recognize the new app
# If your app already exists, you don't have to wait 30 seconds.
resource "time_sleep" "wait_30s" {
  depends_on      = [google_firebase_apple_app.default]
  create_duration = "30s"
}

resource "google_firebase_app_check_device_check_config" "default" {
  provider = google-beta

  project     = "my-project-name"
  app_id      = google_firebase_apple_app.default.app_id
  token_ttl   = "7200s"
  key_id      = "Key ID"
  private_key = file("path/to/private-key.p8")

  depends_on = [time_sleep.wait_30s]

  lifecycle {
    precondition {
      condition     = google_firebase_apple_app.default.team_id != ""
      error_message = "Provide a Team ID on the Apple App to use App Check"
    }
  }
}
```

## Argument Reference

The following arguments are supported:


* `key_id` -
  (Required)
  The key identifier of a private key enabled with DeviceCheck, created in your Apple Developer account.

* `private_key` -
  (Required)
  The contents of the private key (.p8) file associated with the key specified by keyId.
  **Note**: This property is sensitive and will not be displayed in the plan.

* `app_id` -
  (Required)
  The ID of an
  [Apple App](https://firebase.google.com/docs/reference/firebase-management/rest/v1beta1/projects.iosApps#IosApp.FIELDS.app_id).


* `token_ttl` -
  (Optional)
  Specifies the duration for which App Check tokens exchanged from DeviceCheck artifacts will be valid.
  If unset, a default value of 1 hour is assumed. Must be between 30 minutes and 7 days, inclusive.
  A duration in seconds with up to nine fractional digits, ending with 's'. Example: "3.5s".

* `project` - (Optional) The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.



## Attributes Reference

In addition to the arguments listed above, the following computed attributes are exported:

* `id` - an identifier for the resource with format `projects/{{project}}/apps/{{app_id}}/deviceCheckConfig`

* `name` -
  The relative resource name of the DeviceCheck configuration object

* `private_key_set` -
  Whether the privateKey field was previously set. Since App Check will never return the
  privateKey field, this field is the only way to find out whether it was previously set.


## Timeouts

This resource provides the following
[Timeouts](https://developer.hashicorp.com/terraform/plugin/sdkv2/resources/retries-and-customizable-timeouts) configuration options:

- `create` - Default is 20 minutes.
- `update` - Default is 20 minutes.
- `delete` - Default is 20 minutes.

## Import


DeviceCheckConfig can be imported using any of these accepted formats:

* `projects/{{project}}/apps/{{app_id}}/deviceCheckConfig`
* `{{project}}/{{app_id}}`
* `{{app_id}}`


In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import DeviceCheckConfig using one of the formats above. For example:

```tf
import {
  id = "projects/{{project}}/apps/{{app_id}}/deviceCheckConfig"
  to = google_firebase_app_check_device_check_config.default
}
```

When using the [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import), DeviceCheckConfig can be imported using one of the formats above. For example:

```
$ terraform import google_firebase_app_check_device_check_config.default projects/{{project}}/apps/{{app_id}}/deviceCheckConfig
$ terraform import google_firebase_app_check_device_check_config.default {{project}}/{{app_id}}
$ terraform import google_firebase_app_check_device_check_config.default {{app_id}}
```

## User Project Overrides

This resource supports [User Project Overrides](https://registry.terraform.io/providers/hashicorp/google/latest/docs/guides/provider_reference#user_project_override).
