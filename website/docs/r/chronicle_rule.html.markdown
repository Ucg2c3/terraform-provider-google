---
# ----------------------------------------------------------------------------
#
#     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
#
# ----------------------------------------------------------------------------
#
#     This code is generated by Magic Modules using the following:
#
#     Configuration: https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/chronicle/Rule.yaml
#     Template:      https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.html.markdown.tmpl
#
#     DO NOT EDIT this file directly. Any changes made to this file will be
#     overwritten during the next generation cycle.
#
# ----------------------------------------------------------------------------
subcategory: "Chronicle"
description: |-
  The Rule resource represents a user-created rule.
---

# google_chronicle_rule

The Rule resource represents a user-created rule.


To get more information about Rule, see:

* [API documentation](https://cloud.google.com/chronicle/docs/reference/rest/v1/projects.locations.instances.rules)
* How-to Guides
    * [Google SecOps Guides](https://cloud.google.com/chronicle/docs/secops/secops-overview)

## Example Usage - Chronicle Rule Basic


```hcl
resource "google_chronicle_rule" "example" {
 location = "us"
 instance = "00000000-0000-0000-0000-000000000000"
 deletion_policy = "DEFAULT"
 text = <<-EOT
             rule test_rule { meta: events:  $userid = $e.principal.user.userid  match: $userid over 10m condition: $e }
         EOT
}
```
## Example Usage - Chronicle Rule With Force Deletion


```hcl
resource "google_chronicle_rule" "example" {
 location = "us"
 instance = "00000000-0000-0000-0000-000000000000"
 deletion_policy = "FORCE"
 text = <<-EOT
             rule test_rule { meta: events:  $userid = $e.principal.user.userid  match: $userid over 10m condition: $e }
         EOT
}
```
## Example Usage - Chronicle Rule With Data Access Scope


```hcl
resource "google_chronicle_data_access_scope" "data_access_scope_test" {
 location = "us"
 instance = "00000000-0000-0000-0000-000000000000"
 data_access_scope_id = "scope-name"
 description = "scope-description"
 allowed_data_access_labels {
   log_type = "GCP_CLOUDAUDIT"
 }
}

resource "google_chronicle_rule" "example" {
 location = "us"
 instance = "00000000-0000-0000-0000-000000000000"
 scope = resource.google_chronicle_data_access_scope.data_access_scope_test.name
 text = <<-EOT
             rule test_rule { meta: events:  $userid = $e.principal.user.userid  match: $userid over 10m condition: $e }
         EOT
}
```

## Argument Reference

The following arguments are supported:


* `location` -
  (Required)
  The location of the resource. This is the geographical region where the Chronicle instance resides, such as "us" or "europe-west2".

* `instance` -
  (Required)
  The unique identifier for the Chronicle instance, which is the same as the customer ID.


* `text` -
  (Optional)
  The YARA-L content of the rule.
  Populated in FULL view.

* `scope` -
  (Optional)
  Resource name of the DataAccessScope bound to this rule.
  Populated in BASIC view and FULL view.
  If reference lists are used in the rule, validations will be performed
  against this scope to ensure that the reference lists are compatible with
  both the user's and the rule's scopes.
  The scope should be in the format:
  "projects/{project}/locations/{location}/instances/{instance}/dataAccessScopes/{scope}".

* `etag` -
  (Optional)
  The etag for this rule.
  If this is provided on update, the request will succeed if and only if it
  matches the server-computed value, and will fail with an ABORTED error
  otherwise.
  Populated in BASIC view and FULL view.

* `project` - (Optional) The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.

* `deletion_policy` - (Optional) Policy to determine if the rule should be deleted forcefully.
If deletion_policy = "FORCE", any retrohunts and any detections associated with the rule
will also be deleted. If deletion_policy = "DEFAULT", the call will only succeed if the
rule has no associated retrohunts, including completed retrohunts, and no
associated detections. Regardless of this field's value, the rule
deployment associated with this rule will also be deleted.
Possible values: DEFAULT, FORCE



## Attributes Reference

In addition to the arguments listed above, the following computed attributes are exported:

* `id` - an identifier for the resource with format `projects/{{project}}/locations/{{location}}/instances/{{instance}}/rules/{{rule_id}}`

* `name` -
  Full resource name for the rule. This unique identifier is generated using values provided for the URL parameters.
  Format:
  projects/{project}/locations/{location}/instances/{instance}/rules/{rule}

* `rule_id` -
  Rule Id is the ID of the Rule.

* `metadata` -
  Output only. Additional metadata specified in the meta section of text.
  Populated in FULL view.

* `near_real_time_live_rule_eligible` -
  Output only. Indicate the rule can run in near real time live rule.
  If this is true, the rule uses the near real time live rule when the run
  frequency is set to LIVE.

* `revision_id` -
  Output only. The revision ID of the rule.
  A new revision is created whenever the rule text is changed in any way.
  Format: v_{10 digits}_{9 digits}
  Populated in REVISION_METADATA_ONLY view and FULL view.

* `severity` -
  Severity represents the severity level of the rule.
  Structure is [documented below](#nested_severity).

* `revision_create_time` -
  Output only. The timestamp of when the rule revision was created.
  Populated in FULL, REVISION_METADATA_ONLY views.

* `compilation_state` -
  Output only. The current compilation state of the rule.
  Populated in FULL view.
  Possible values:
  COMPILATION_STATE_UNSPECIFIED
  SUCCEEDED
  FAILED

* `type` -
  Possible values:
  RULE_TYPE_UNSPECIFIED
  SINGLE_EVENT
  MULTI_EVENT

* `reference_lists` -
  Output only. Resource names of the reference lists used in this rule.
  Populated in FULL view.

* `display_name` -
  Output only. Display name of the rule.
  Populated in BASIC view and FULL view.

* `create_time` -
  Output only. The timestamp of when the rule was created.
  Populated in FULL view.

* `author` -
  Output only. The author of the rule. Extracted from the meta section of text.
  Populated in BASIC view and FULL view.

* `allowed_run_frequencies` -
  Output only. The run frequencies that are allowed for the rule.
  Populated in BASIC view and FULL view.

* `compilation_diagnostics` -
  Output only. A list of a rule's corresponding compilation diagnostic messages
  such as compilation errors and compilation warnings.
  Populated in FULL view.
  Structure is [documented below](#nested_compilation_diagnostics).

* `data_tables` -
  Output only. Resource names of the data tables used in this rule.


<a name="nested_severity"></a>The `severity` block contains:

* `display_name` -
  (Optional)
  The display name of the severity level. Extracted from the meta section of
  the rule text.

<a name="nested_compilation_diagnostics"></a>The `compilation_diagnostics` block contains:

* `message` -
  (Output)
  Output only. The diagnostic message.

* `position` -
  (Optional)
  CompilationPosition represents the location of a compilation diagnostic in
  rule text.
  Structure is [documented below](#nested_compilation_diagnostics_compilation_diagnostics_position).

* `severity` -
  (Output)
  Output only. The severity of a rule's compilation diagnostic.
  Possible values:
  SEVERITY_UNSPECIFIED
  WARNING
  ERROR

* `uri` -
  (Output)
  Output only. Link to documentation that describes a diagnostic in more detail.


<a name="nested_compilation_diagnostics_compilation_diagnostics_position"></a>The `position` block supports:

* `start_line` -
  (Output)
  Output only. Start line number, beginning at 1.

* `start_column` -
  (Output)
  Output only. Start column number, beginning at 1.

* `end_line` -
  (Output)
  Output only. End line number, beginning at 1.

* `end_column` -
  (Output)
  Output only. End column number, beginning at 1.

## Timeouts

This resource provides the following
[Timeouts](https://developer.hashicorp.com/terraform/plugin/sdkv2/resources/retries-and-customizable-timeouts) configuration options:

- `create` - Default is 20 minutes.
- `update` - Default is 20 minutes.
- `delete` - Default is 20 minutes.

## Import


Rule can be imported using any of these accepted formats:

* `projects/{{project}}/locations/{{location}}/instances/{{instance}}/rules/{{rule_id}}`
* `{{project}}/{{location}}/{{instance}}/{{rule_id}}`
* `{{location}}/{{instance}}/{{rule_id}}`


In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import Rule using one of the formats above. For example:

```tf
import {
  id = "projects/{{project}}/locations/{{location}}/instances/{{instance}}/rules/{{rule_id}}"
  to = google_chronicle_rule.default
}
```

When using the [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import), Rule can be imported using one of the formats above. For example:

```
$ terraform import google_chronicle_rule.default projects/{{project}}/locations/{{location}}/instances/{{instance}}/rules/{{rule_id}}
$ terraform import google_chronicle_rule.default {{project}}/{{location}}/{{instance}}/{{rule_id}}
$ terraform import google_chronicle_rule.default {{location}}/{{instance}}/{{rule_id}}
```

## User Project Overrides

This resource supports [User Project Overrides](https://registry.terraform.io/providers/hashicorp/google/latest/docs/guides/provider_reference#user_project_override).
