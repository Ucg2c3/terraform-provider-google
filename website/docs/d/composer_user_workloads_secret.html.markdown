---
# ----------------------------------------------------------------------------
#
#     ***     AUTO GENERATED CODE    ***    Type: Handwritten     ***
#
# ----------------------------------------------------------------------------
#
#     This code is generated by Magic Modules using the following:
#
#     Source file: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/third_party/terraform/website/docs/d/composer_user_workloads_secret.html.markdown
#
#     DO NOT EDIT this file directly. Any changes made to this file will be
#     overwritten during the next generation cycle.
#
# ----------------------------------------------------------------------------
subcategory: "Cloud Composer"
description: |-
  User workloads Secret used by Airflow tasks that run with Kubernetes Executor or KubernetesPodOperator.
---

# google_composer_user_workloads_secret

Provides access to Kubernetes Secret configuration for a given project, region and Composer Environment.

To get more information about Composer User Workloads Secrets, see:

* [API documentation](https://cloud.google.com/composer/docs/reference/rest/v1/projects.locations.environments.userWorkloadsSecrets)
* How-to Guides
    * [Official Documentation](https://cloud.google.com/artifact-registry/docs/overview)
    
## Example Usage

```hcl
resource "google_composer_environment" "example" {
    name = "example-environment"
    config{
        software_config {
            image_version = "composer-3-airflow-2"
        }
    }
}

resource "google_composer_user_workloads_secret" "example" {
    environment = google_composer_environment.example.name
    name = "example-secret"
    data = {
        username: base64encode("username"),
        password: base64encode("password"),
    }
}

data "google_composer_user_workloads_secret" "example" {
    environment = google_composer_environment.example.name
    name = resource.google_composer_user_workloads_secret.example.name
}

output "debug" {
    value = data.google_composer_user_workloads_secret.example
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of the Secret.

* `environment` - (Required) Environment where the Secret is stored.

* `project` - (Optional) The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.

* `region` - (Optional) The location or Compute Engine region of the environment.

## Attributes Reference

See [google_composer_user_workloads_secret](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/composer_user_workloads_secret) resource for details of the available attributes.
