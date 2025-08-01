---
# ----------------------------------------------------------------------------
#
#     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
#
# ----------------------------------------------------------------------------
#
#     This code is generated by Magic Modules using the following:
#
#     Configuration: https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/firebaseapphosting/Backend.yaml
#     Template:      https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.html.markdown.tmpl
#
#     DO NOT EDIT this file directly. Any changes made to this file will be
#     overwritten during the next generation cycle.
#
# ----------------------------------------------------------------------------
subcategory: "Firebase App Hosting"
description: |-
  A Backend is the primary resource of App Hosting.
---

# google_firebase_app_hosting_backend

A Backend is the primary resource of App Hosting.



## Example Usage - Firebase App Hosting Backend Minimal


```hcl
resource "google_firebase_app_hosting_backend" "example" {
  project          = "my-project-name"
  # Choose the region closest to your users

  location         = "us-central1"
  backend_id       = "mini"
  app_id           = "1:0000000000:web:674cde32020e16fbce9dbd"
  serving_locality = "GLOBAL_ACCESS"
  service_account  = google_service_account.service_account.email

  depends_on = [google_project_service.fah]
}

### Include these blocks only once per project if you are starting from scratch ###
resource "google_service_account" "service_account" {
  project = "my-project-name"

  # Must be firebase-app-hosting-compute
  account_id                   = "firebase-app-hosting-compute"
  display_name                 = "Firebase App Hosting compute service account"

  # Do not throw if already exists
  create_ignore_already_exists = true
}

resource "google_project_iam_member" "app_hosting_sa_runner" {
  project = "my-project-name"

  # For App Hosting
  role   = "roles/firebaseapphosting.computeRunner"
  member = google_service_account.service_account.member
}

resource "google_project_service" "fah" {
  project = "my-project-name"
  service = "firebaseapphosting.googleapis.com"

  disable_on_destroy = false
}
###
```
## Example Usage - Firebase App Hosting Backend Full


```hcl
resource "google_firebase_app_hosting_backend" "example" {
  project          = "my-project-name"

  # Choose the region closest to your users
  location         = "us-central1"
  backend_id       = "full"
  app_id           = "1:0000000000:web:674cde32020e16fbce9dbd"
  display_name     = "My Backend"
  serving_locality = "GLOBAL_ACCESS"
  service_account  = google_service_account.service_account.email
  environment      = "prod"

  annotations = {
    "key" = "value"
  }

  labels = {
    "key" = "value"
  }

  depends_on = [google_project_service.fah]
}

### Include these blocks only once per project if you are starting from scratch ###
resource "google_service_account" "service_account" {
  project = "my-project-name"

  # Must be firebase-app-hosting-compute
  account_id                   = "firebase-app-hosting-compute"
  display_name                 = "Firebase App Hosting compute service account"

  # Do not throw if already exists
  create_ignore_already_exists = true
}

resource "google_project_iam_member" "app_hosting_sa_developerconnect" {
  project = "my-project-name"

  # For reading connected Github repos
  role   = "roles/developerconnect.readTokenAccessor"
  member = google_service_account.service_account.member
}

resource "google_project_iam_member" "app_hosting_sa_adminsdk" {
  project = "my-project-name"

  # For Firebase Admin SDK
  role   = "roles/firebase.sdkAdminServiceAgent"
  member = google_service_account.service_account.member
}

resource "google_project_iam_member" "app_hosting_sa_runner" {
  project = "my-project-name"

  # For App Hosting
  role   = "roles/firebaseapphosting.computeRunner"
  member = google_service_account.service_account.member
}

resource "google_project_service" "fah" {
  project = "my-project-name"
  service = "firebaseapphosting.googleapis.com"

  disable_on_destroy = false
}
###
```
## Example Usage - Firebase App Hosting Backend Github


```hcl
resource "google_firebase_app_hosting_backend" "example" {
  project          = "my-project-name"

  # Choose the region closest to your users
  location         = "us-central1"
  backend_id       = "my-backend-gh"
  app_id           = "1:0000000000:web:674cde32020e16fbce9dbd"
  display_name     = "My Backend"
  serving_locality = "GLOBAL_ACCESS"
  service_account  = "firebase-app-hosting-compute@my-project-name.iam.gserviceaccount.com"
  environment      = "prod"

  annotations = {
    "key" = "value"
  }

  labels = {
    "key" = "value"
  }

  codebase {
    repository = google_developer_connect_git_repository_link.my-repository.name
    root_directory = "/"
  }
}

resource "google_developer_connect_git_repository_link" "my-repository" {
  project  = "my-project-name"
  location = "us-central1"

  git_repository_link_id = "my-repo"
  parent_connection = google_developer_connect_connection.my-connection.connection_id
  clone_uri = "https://github.com/myuser/myrepo.git"
}

### Include these blocks only once per project if you are starting from scratch ###
resource "google_project_service_identity" "devconnect-p4sa" {
  provider = google-beta

  project  = "my-project-name"
  service  = "developerconnect.googleapis.com"
}

resource "google_project_iam_member" "devconnect-secret" {
  project  = "my-project-name"
  role     = "roles/secretmanager.admin"
  member   = google_project_service_identity.devconnect-p4sa.member
}
###

### Include these blocks only once per Github account ###
resource "google_developer_connect_connection" "my-connection" {
  project  = "my-project-name"
  location = "us-central1"
  connection_id = "tf-test-connection-new"
  github_config {
    github_app = "FIREBASE"
  }
  depends_on = [google_project_iam_member.devconnect-secret]
}

output "next_steps" {
  description = "Follow the action_uri if present to continue setup"
  value = google_developer_connect_connection.my-connection.installation_state
}
###
```

## Argument Reference

The following arguments are supported:


* `serving_locality` -
  (Required)
  Immutable. Specifies how App Hosting will serve the content for this backend. It will
  either be contained to a single region (REGIONAL_STRICT) or allowed to use
  App Hosting's global-replicated serving infrastructure (GLOBAL_ACCESS).
  Possible values are: `REGIONAL_STRICT`, `GLOBAL_ACCESS`.

* `app_id` -
  (Required)
  The [ID of a Web
  App](https://firebase.google.com/docs/reference/firebase-management/rest/v1beta1/projects.webApps#WebApp.FIELDS.app_id)
  associated with the backend.

* `service_account` -
  (Required)
  The name of the service account used for Cloud Build and Cloud Run.
  Should have the role roles/firebaseapphosting.computeRunner
  or equivalent permissions.

* `location` -
  (Required)
  The canonical IDs of a Google Cloud location such as "us-east1".

* `backend_id` -
  (Required)
  Id of the backend. Also used as the service ID for Cloud Run, and as part
  of the default domain name.


* `annotations` -
  (Optional)
  Unstructured key value map that may be set by external tools to
  store and arbitrary metadata. They are not queryable and should be
  preserved when modifying objects.
  **Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.
  Please refer to the field `effective_annotations` for all of the annotations present on the resource.

* `display_name` -
  (Optional)
  Human-readable name. 63 character limit.

* `environment` -
  (Optional)
  The environment name of the backend, used to load environment variables
  from environment specific configuration.

* `labels` -
  (Optional)
  Unstructured key value map that can be used to organize and categorize
  objects.
  **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
  Please refer to the field `effective_labels` for all of the labels present on the resource.

* `codebase` -
  (Optional)
  The connection to an external source repository to watch for event-driven
  updates to the backend.
  Structure is [documented below](#nested_codebase).

* `project` - (Optional) The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.



<a name="nested_codebase"></a>The `codebase` block supports:

* `repository` -
  (Required)
  The resource name for the Developer Connect
  [`gitRepositoryLink`](https://cloud.google.com/developer-connect/docs/api/reference/rest/v1/projects.locations.connections.gitRepositoryLinks)
  connected to this backend, in the format:
  projects/{project}/locations/{location}/connections/{connection}/gitRepositoryLinks/{repositoryLink}

* `root_directory` -
  (Optional)
  If `repository` is provided, the directory relative to the root of the
  repository to use as the root for the deployed web app.

## Attributes Reference

In addition to the arguments listed above, the following computed attributes are exported:

* `id` - an identifier for the resource with format `projects/{{project}}/locations/{{location}}/backends/{{backend_id}}`

* `etag` -
  Server-computed checksum based on other values; may be sent
  on update or delete to ensure operation is done on expected resource.

* `name` -
  Identifier. The resource name of the backend.
  Format:
  `projects/{project}/locations/{locationId}/backends/{backendId}`.

* `create_time` -
  Time at which the backend was created.

* `delete_time` -
  Time at which the backend was deleted.

* `update_time` -
  Time at which the backend was last updated.

* `uid` -
  System-assigned, unique identifier.

* `uri` -
  The primary URI to communicate with the backend.

* `managed_resources` -
  A list of the resources managed by this backend.
  Structure is [documented below](#nested_managed_resources).

* `effective_annotations` -
  All of annotations (key/value pairs) present on the resource in GCP, including the annotations configured through Terraform, other clients and services.

* `terraform_labels` -
  The combination of labels configured directly on the resource
   and default labels configured on the provider.

* `effective_labels` -
  All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other clients and services.


<a name="nested_managed_resources"></a>The `managed_resources` block contains:

* `run_service` -
  (Output)
  A managed Cloud Run
  [`service`](https://cloud.google.com/run/docs/reference/rest/v2/projects.locations.services#resource:-service).
  Structure is [documented below](#nested_managed_resources_managed_resources_run_service).


<a name="nested_managed_resources_managed_resources_run_service"></a>The `run_service` block contains:

* `service` -
  (Output)
  The name of the Cloud Run
  [`service`](https://cloud.google.com/run/docs/reference/rest/v2/projects.locations.services#resource:-service),
  in the format:
  projects/{project}/locations/{location}/services/{serviceId}

## Timeouts

This resource provides the following
[Timeouts](https://developer.hashicorp.com/terraform/plugin/sdkv2/resources/retries-and-customizable-timeouts) configuration options:

- `create` - Default is 20 minutes.
- `update` - Default is 20 minutes.
- `delete` - Default is 20 minutes.

## Import


Backend can be imported using any of these accepted formats:

* `projects/{{project}}/locations/{{location}}/backends/{{backend_id}}`
* `{{project}}/{{location}}/{{backend_id}}`
* `{{location}}/{{backend_id}}`


In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import Backend using one of the formats above. For example:

```tf
import {
  id = "projects/{{project}}/locations/{{location}}/backends/{{backend_id}}"
  to = google_firebase_app_hosting_backend.default
}
```

When using the [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import), Backend can be imported using one of the formats above. For example:

```
$ terraform import google_firebase_app_hosting_backend.default projects/{{project}}/locations/{{location}}/backends/{{backend_id}}
$ terraform import google_firebase_app_hosting_backend.default {{project}}/{{location}}/{{backend_id}}
$ terraform import google_firebase_app_hosting_backend.default {{location}}/{{backend_id}}
```

## User Project Overrides

This resource supports [User Project Overrides](https://registry.terraform.io/providers/hashicorp/google/latest/docs/guides/provider_reference#user_project_override).
