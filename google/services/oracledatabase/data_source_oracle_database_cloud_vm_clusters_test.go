// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
// ----------------------------------------------------------------------------
//
//	***     AUTO GENERATED CODE    ***    Type: Handwritten     ***
//
// ----------------------------------------------------------------------------
//
//	This code is generated by Magic Modules using the following:
//
//	Source file: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/third_party/terraform/services/oracledatabase/data_source_oracle_database_cloud_vm_clusters_test.go
//
//	DO NOT EDIT this file directly. Any changes made to this file will be
//	overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------
package oracledatabase_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-google/google/acctest"
)

func TestAccOracleDatabaseCloudVmClusters_basic(t *testing.T) {
	t.Parallel()
	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccOracleDatabaseCloudVmClusters_basic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.google_oracle_database_cloud_vm_clusters.my_vmclusters", "cloud_vm_clusters.#"),
					resource.TestCheckResourceAttrSet("data.google_oracle_database_cloud_vm_clusters.my_vmclusters", "cloud_vm_clusters.0.display_name"),
					resource.TestCheckResourceAttrSet("data.google_oracle_database_cloud_vm_clusters.my_vmclusters", "cloud_vm_clusters.0.properties.#"),
				),
			},
		},
	})
}

func testAccOracleDatabaseCloudVmClusters_basic() string {
	return fmt.Sprintf(`
data "google_oracle_database_cloud_vm_clusters" "my_vmclusters"{
  location = "us-east4"
  project = "oci-terraform-testing-prod"
}
`)
}
