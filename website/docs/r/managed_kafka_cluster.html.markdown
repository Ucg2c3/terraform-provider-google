---
# ----------------------------------------------------------------------------
#
#     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
#
# ----------------------------------------------------------------------------
#
#     This code is generated by Magic Modules using the following:
#
#     Configuration: https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/managedkafka/Cluster.yaml
#     Template:      https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.html.markdown.tmpl
#
#     DO NOT EDIT this file directly. Any changes made to this file will be
#     overwritten during the next generation cycle.
#
# ----------------------------------------------------------------------------
subcategory: "Managed Kafka"
description: |-
  A Managed Service for Apache Kafka cluster.
---

# google_managed_kafka_cluster

A Managed Service for Apache Kafka cluster. Apache Kafka is a trademark owned by the Apache Software Foundation.



<div class = "oics-button" style="float: right; margin: 0 0 -15px">
  <a href="https://console.cloud.google.com/cloudshell/open?cloudshell_git_repo=https%3A%2F%2Fgithub.com%2Fterraform-google-modules%2Fdocs-examples.git&cloudshell_image=gcr.io%2Fcloudshell-images%2Fcloudshell%3Alatest&cloudshell_print=.%2Fmotd&cloudshell_tutorial=.%2Ftutorial.md&cloudshell_working_dir=managedkafka_cluster_basic&open_in_editor=main.tf" target="_blank">
    <img alt="Open in Cloud Shell" src="//gstatic.com/cloudssh/images/open-btn.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
  </a>
</div>
## Example Usage - Managedkafka Cluster Basic


```hcl
resource "google_managed_kafka_cluster" "example" {
  cluster_id = "my-cluster"
  location = "us-central1"
  capacity_config {
    vcpu_count = 3
    memory_bytes = 3221225472
  }
  gcp_config {
    access_config {
      network_configs {
        subnet = "projects/${data.google_project.project.number}/regions/us-central1/subnetworks/default"
      }
    }
  }
  rebalance_config {
    mode = "AUTO_REBALANCE_ON_SCALE_UP"
  }
  labels = {
    key = "value"
  }
}

data "google_project" "project" {
}
```
<div class = "oics-button" style="float: right; margin: 0 0 -15px">
  <a href="https://console.cloud.google.com/cloudshell/open?cloudshell_git_repo=https%3A%2F%2Fgithub.com%2Fterraform-google-modules%2Fdocs-examples.git&cloudshell_image=gcr.io%2Fcloudshell-images%2Fcloudshell%3Alatest&cloudshell_print=.%2Fmotd&cloudshell_tutorial=.%2Ftutorial.md&cloudshell_working_dir=managedkafka_cluster_mtls&open_in_editor=main.tf" target="_blank">
    <img alt="Open in Cloud Shell" src="//gstatic.com/cloudssh/images/open-btn.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
  </a>
</div>
## Example Usage - Managedkafka Cluster Mtls


```hcl
resource "google_managed_kafka_cluster" "example" {
  cluster_id = "my-cluster"
  location = "us-central1"
  capacity_config {
    vcpu_count = 3
    memory_bytes = 3221225472
  }
  gcp_config {
    access_config {
      network_configs {
        subnet = "projects/${data.google_project.project.number}/regions/us-central1/subnetworks/default"
      }
    }
  }
  tls_config {
    trust_config {
      cas_configs {
        ca_pool = google_privateca_ca_pool.ca_pool.id 
      }
    }
    ssl_principal_mapping_rules = "RULE:pattern/replacement/L,DEFAULT"
  }
}

resource "google_privateca_ca_pool" "ca_pool" {
  name = "my-ca-pool"
  location = "us-central1"
  tier = "ENTERPRISE"
  publishing_options {
    publish_ca_cert = true
    publish_crl = true
  }
}

data "google_project" "project" {
}
```
<div class = "oics-button" style="float: right; margin: 0 0 -15px">
  <a href="https://console.cloud.google.com/cloudshell/open?cloudshell_git_repo=https%3A%2F%2Fgithub.com%2Fterraform-google-modules%2Fdocs-examples.git&cloudshell_image=gcr.io%2Fcloudshell-images%2Fcloudshell%3Alatest&cloudshell_print=.%2Fmotd&cloudshell_tutorial=.%2Ftutorial.md&cloudshell_working_dir=managedkafka_cluster_cmek&open_in_editor=main.tf" target="_blank">
    <img alt="Open in Cloud Shell" src="//gstatic.com/cloudssh/images/open-btn.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
  </a>
</div>
## Example Usage - Managedkafka Cluster Cmek


```hcl
resource "google_managed_kafka_cluster" "example" {
  cluster_id = "my-cluster"
  location = "us-central1"
  capacity_config {
    vcpu_count = 3
    memory_bytes = 3221225472
  }
  gcp_config {
    access_config {
      network_configs {
        subnet = "projects/${data.google_project.project.number}/regions/us-central1/subnetworks/default"
      }
    }
    kms_key = google_kms_crypto_key.key.id
  }

  provider = google-beta
}

resource "google_project_service_identity" "kafka_service_identity" {
  project  = data.google_project.project.project_id
  service  = "managedkafka.googleapis.com"

  provider = google-beta
}

resource "google_kms_crypto_key" "key" {
  name     = "example-key"
  key_ring = google_kms_key_ring.key_ring.id

  provider = google-beta
}

resource "google_kms_key_ring" "key_ring" {
  name     = "example-key-ring"
  location = "us-central1"

  provider = google-beta
}

resource "google_kms_crypto_key_iam_binding" "crypto_key_binding" {
  crypto_key_id = google_kms_crypto_key.key.id
  role          = "roles/cloudkms.cryptoKeyEncrypterDecrypter"

  members = [
    "serviceAccount:service-${data.google_project.project.number}@gcp-sa-managedkafka.iam.gserviceaccount.com",
  ]

  provider = google-beta
}

data "google_project" "project" {
  provider = google-beta
}
```

## Argument Reference

The following arguments are supported:


* `gcp_config` -
  (Required)
  Configuration properties for a Kafka cluster deployed to Google Cloud Platform.
  Structure is [documented below](#nested_gcp_config).

* `capacity_config` -
  (Required)
  A capacity configuration of a Kafka cluster.
  Structure is [documented below](#nested_capacity_config).

* `location` -
  (Required)
  ID of the location of the Kafka resource. See https://cloud.google.com/managed-kafka/docs/locations for a list of supported locations.

* `cluster_id` -
  (Required)
  The ID to use for the cluster, which will become the final component of the cluster's name. The ID must be 1-63 characters long, and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` to comply with RFC 1035. This value is structured like: `my-cluster-id`.


* `labels` -
  (Optional)
  List of label KEY=VALUE pairs to add. Keys must start with a lowercase character and contain only hyphens (-), underscores ( ), lowercase characters, and numbers. Values must contain only hyphens (-), underscores ( ), lowercase characters, and numbers.
  **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
  Please refer to the field `effective_labels` for all of the labels present on the resource.

* `rebalance_config` -
  (Optional)
  Defines rebalancing behavior of a Kafka cluster.
  Structure is [documented below](#nested_rebalance_config).

* `tls_config` -
  (Optional)
  TLS configuration for the Kafka cluster. This is used to configure mTLS authentication. To clear our a TLS configuration that has been previously set, please explicitly add an empty `tls_config` block.
  Structure is [documented below](#nested_tls_config).

* `project` - (Optional) The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.



<a name="nested_gcp_config"></a>The `gcp_config` block supports:

* `access_config` -
  (Required)
  The configuration of access to the Kafka cluster.
  Structure is [documented below](#nested_gcp_config_access_config).

* `kms_key` -
  (Optional)
  The Cloud KMS Key name to use for encryption. The key must be located in the same region as the cluster and cannot be changed. Must be in the format `projects/PROJECT_ID/locations/LOCATION/keyRings/KEY_RING/cryptoKeys/KEY`.


<a name="nested_gcp_config_access_config"></a>The `access_config` block supports:

* `network_configs` -
  (Required)
  Virtual Private Cloud (VPC) subnets where IP addresses for the Kafka cluster are allocated. To make the cluster available in a VPC, you must specify at least one `network_configs` block. Max of 10 subnets per cluster. Additional subnets may be specified with additional `network_configs` blocks.
  Structure is [documented below](#nested_gcp_config_access_config_network_configs).


<a name="nested_gcp_config_access_config_network_configs"></a>The `network_configs` block supports:

* `subnet` -
  (Required)
  Name of the VPC subnet from which the cluster is accessible. Both broker and bootstrap server IP addresses and DNS entries are automatically created in the subnet. There can only be one subnet per network, and the subnet must be located in the same region as the cluster. The project may differ. The name of the subnet must be in the format `projects/PROJECT_ID/regions/REGION/subnetworks/SUBNET`.

<a name="nested_capacity_config"></a>The `capacity_config` block supports:

* `vcpu_count` -
  (Required)
  The number of vCPUs to provision for the cluster. The minimum is 3.

* `memory_bytes` -
  (Required)
  The memory to provision for the cluster in bytes. The value must be between 1 GiB and 8 GiB per vCPU. Ex. 1024Mi, 4Gi.

<a name="nested_rebalance_config"></a>The `rebalance_config` block supports:

* `mode` -
  (Optional)
  The rebalance behavior for the cluster. When not specified, defaults to `NO_REBALANCE`. Possible values: `MODE_UNSPECIFIED`, `NO_REBALANCE`, `AUTO_REBALANCE_ON_SCALE_UP`.

<a name="nested_tls_config"></a>The `tls_config` block supports:

* `trust_config` -
  (Optional)
  The configuration of the broker truststore. If specified, clients can use mTLS for authentication.
  Structure is [documented below](#nested_tls_config_trust_config).

* `ssl_principal_mapping_rules` -
  (Optional)
  The rules for mapping mTLS certificate Distinguished Names (DNs) to shortened principal names for Kafka ACLs. This field corresponds exactly to the ssl.principal.mapping.rules broker config and matches the format and syntax defined in the Apache Kafka documentation. Setting or modifying this field will trigger a rolling restart of the Kafka brokers to apply the change. An empty string means that the default Kafka behavior is used. Example: `RULE:^CN=(.?),OU=ServiceUsers.$/$1@example.com/,DEFAULT`


<a name="nested_tls_config_trust_config"></a>The `trust_config` block supports:

* `cas_configs` -
  (Optional)
  Configuration for the Google Certificate Authority Service. To support mTLS, you must specify at least one `cas_configs` block. A maximum of 10 CA pools can be specified. Additional CA pools may be specified with additional `cas_configs` blocks.
  Structure is [documented below](#nested_tls_config_trust_config_cas_configs).


<a name="nested_tls_config_trust_config_cas_configs"></a>The `cas_configs` block supports:

* `ca_pool` -
  (Required)
  The name of the CA pool to pull CA certificates from. The CA pool does not need to be in the same project or location as the Kafka cluster. Must be in the format `projects/PROJECT_ID/locations/LOCATION/caPools/CA_POOL_ID.

## Attributes Reference

In addition to the arguments listed above, the following computed attributes are exported:

* `id` - an identifier for the resource with format `projects/{{project}}/locations/{{location}}/clusters/{{cluster_id}}`

* `name` -
  The name of the cluster. Structured like: `projects/PROJECT_ID/locations/LOCATION/clusters/CLUSTER_ID`.

* `create_time` -
  The time when the cluster was created.

* `update_time` -
  The time when the cluster was last updated.

* `state` -
  The current state of the cluster. Possible values: `STATE_UNSPECIFIED`, `CREATING`, `ACTIVE`, `DELETING`.

* `terraform_labels` -
  The combination of labels configured directly on the resource
   and default labels configured on the provider.

* `effective_labels` -
  All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other clients and services.


## Timeouts

This resource provides the following
[Timeouts](https://developer.hashicorp.com/terraform/plugin/sdkv2/resources/retries-and-customizable-timeouts) configuration options:

- `create` - Default is 60 minutes.
- `update` - Default is 30 minutes.
- `delete` - Default is 30 minutes.

## Import


Cluster can be imported using any of these accepted formats:

* `projects/{{project}}/locations/{{location}}/clusters/{{cluster_id}}`
* `{{project}}/{{location}}/{{cluster_id}}`
* `{{location}}/{{cluster_id}}`


In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import Cluster using one of the formats above. For example:

```tf
import {
  id = "projects/{{project}}/locations/{{location}}/clusters/{{cluster_id}}"
  to = google_managed_kafka_cluster.default
}
```

When using the [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import), Cluster can be imported using one of the formats above. For example:

```
$ terraform import google_managed_kafka_cluster.default projects/{{project}}/locations/{{location}}/clusters/{{cluster_id}}
$ terraform import google_managed_kafka_cluster.default {{project}}/{{location}}/{{cluster_id}}
$ terraform import google_managed_kafka_cluster.default {{location}}/{{cluster_id}}
```

## User Project Overrides

This resource supports [User Project Overrides](https://registry.terraform.io/providers/hashicorp/google/latest/docs/guides/provider_reference#user_project_override).
