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
//	Source file: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/third_party/terraform/services/resourcemanager/data_source_google_iam_role.go
//
//	DO NOT EDIT this file directly. Any changes made to this file will be
//	overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------
package resourcemanager

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func DataSourceGoogleIamRole() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGoogleIamRoleRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"title": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"included_permissions": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"stage": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceGoogleIamRoleRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	roleName := d.Get("name").(string)
	role, err := config.NewIamClient(userAgent).Roles.Get(roleName).Do()
	if err != nil {
		return transport_tpg.HandleDataSourceNotFoundError(err, d, fmt.Sprintf("Error reading IAM Role %s: %s", roleName, err), roleName)
	}

	d.SetId(role.Name)
	if err := d.Set("title", role.Title); err != nil {
		return fmt.Errorf("Error setting title: %s", err)
	}
	if err := d.Set("stage", role.Stage); err != nil {
		return fmt.Errorf("Error setting stage: %s", err)
	}
	if err := d.Set("included_permissions", role.IncludedPermissions); err != nil {
		return fmt.Errorf("Error setting included_permissions: %s", err)
	}

	return nil
}
