// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package contactcenterinsights_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/hashicorp/terraform-provider-google/google/acctest"
	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func TestAccContactCenterInsightsView_contactCenterInsightsViewBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckContactCenterInsightsViewDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccContactCenterInsightsView_contactCenterInsightsViewBasicExample(context),
			},
			{
				ResourceName:            "google_contact_center_insights_view.basic_view",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location"},
			},
		},
	})
}

func testAccContactCenterInsightsView_contactCenterInsightsViewBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_contact_center_insights_view" "basic_view" {
  location = "us-central1"
  display_name = "view-display-name"
  value    = "medium=\"CHAT\""
}
`, context)
}

func TestAccContactCenterInsightsView_contactCenterInsightsViewFullExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckContactCenterInsightsViewDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccContactCenterInsightsView_contactCenterInsightsViewFullExample(context),
			},
			{
				ResourceName:            "google_contact_center_insights_view.full_view",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location"},
			},
		},
	})
}

func testAccContactCenterInsightsView_contactCenterInsightsViewFullExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_contact_center_insights_view" "full_view" {
  location = "us-central1"
  display_name = "view-display-name"
  value    = "medium=\"PHONE_CALL\""
}
`, context)
}

func testAccCheckContactCenterInsightsViewDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_contact_center_insights_view" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{ContactCenterInsightsBasePath}}projects/{{project}}/locations/{{location}}/views/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
				Config:    config,
				Method:    "GET",
				Project:   billingProject,
				RawURL:    url,
				UserAgent: config.UserAgent,
			})
			if err == nil {
				return fmt.Errorf("ContactCenterInsightsView still exists at %s", url)
			}
		}

		return nil
	}
}
