// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/discoveryengine/Sitemap.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package discoveryengine

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"regexp"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func ResourceDiscoveryEngineSitemap() *schema.Resource {
	return &schema.Resource{
		Create: resourceDiscoveryEngineSitemapCreate,
		Read:   resourceDiscoveryEngineSitemapRead,
		Delete: resourceDiscoveryEngineSitemapDelete,

		Importer: &schema.ResourceImporter{
			State: resourceDiscoveryEngineSitemapImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
			Delete: schema.DefaultTimeout(60 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			tpgresource.DefaultProviderProject,
		),

		Schema: map[string]*schema.Schema{
			"data_store_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The unique id of the data store.`,
			},
			"location": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `The geographic location where the data store should reside. The value can
only be one of "global", "us" and "eu".`,
			},
			"uri": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `Public URI for the sitemap, e.g. "www.example.com/sitemap.xml".`,
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Timestamp when the sitemap was created.`,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The unique full resource name of the sitemap. Values are of the format
'projects/{project}/locations/{location}/collections/{collection_id}/dataStores/{data_store_id}/siteSearchEngine/sitemaps/{sitemap_id}'.
This field must be a UTF-8 encoded string with a length limit of 1024
characters.`,
			},
			"sitemap_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The unique id of the sitemap.`,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceDiscoveryEngineSitemapCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	uriProp, err := expandDiscoveryEngineSitemapUri(d.Get("uri"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("uri"); !tpgresource.IsEmptyValue(reflect.ValueOf(uriProp)) && (ok || !reflect.DeepEqual(v, uriProp)) {
		obj["uri"] = uriProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{DiscoveryEngineBasePath}}projects/{{project}}/locations/{{location}}/collections/default_collection/dataStores/{{data_store_id}}/siteSearchEngine/sitemaps")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Sitemap: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Sitemap: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "POST",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutCreate),
		Headers:   headers,
	})
	if err != nil {
		return fmt.Errorf("Error creating Sitemap: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Use the resource in the operation response to populate
	// identity fields and d.Id() before read
	var opRes map[string]interface{}
	err = DiscoveryEngineOperationWaitTimeWithResponse(
		config, res, &opRes, project, "Creating Sitemap", userAgent,
		d.Timeout(schema.TimeoutCreate))
	if err != nil {
		// The resource didn't actually create
		d.SetId("")

		return fmt.Errorf("Error waiting to create Sitemap: %s", err)
	}

	if err := d.Set("name", flattenDiscoveryEngineSitemapName(opRes["name"], d, config)); err != nil {
		return err
	}

	// This may have caused the ID to update - update it if so.
	id, err = tpgresource.ReplaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Sitemap %q: %#v", d.Id(), res)

	return resourceDiscoveryEngineSitemapRead(d, meta)
}

func resourceDiscoveryEngineSitemapRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{DiscoveryEngineBasePath}}projects/{{project}}/locations/{{location}}/collections/default_collection/dataStores/{{data_store_id}}/siteSearchEngine/sitemaps:fetch")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Sitemap: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Headers:   headers,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("DiscoveryEngineSitemap %q", d.Id()))
	}
	// Extract the resource name from the ID to match against the sitemap entries
	resourceName := d.Get("name").(string)

	// Find the specific sitemap entry from the response that matches our resource
	var sitemapData map[string]interface{}
	if sitemapsMetadata, ok := res["sitemapsMetadata"].([]interface{}); ok {
		for _, metadata := range sitemapsMetadata {
			metadataMap, ok := metadata.(map[string]interface{})
			if !ok {
				continue
			}

			sitemap, ok := metadataMap["sitemap"].(map[string]interface{})
			if !ok {
				continue
			}

			name, ok := sitemap["name"].(string)
			if !ok {
				continue
			}

			if name == resourceName {
				sitemapData = sitemap
				break
			}
		}
	}

	// If we didn't find a matching sitemap, return a not found error
	if sitemapData == nil {
		return transport_tpg.HandleNotFoundError(fmt.Errorf("Sitemap %q not found", resourceName), d, fmt.Sprintf("DiscoveryEngineSitemap %q", d.Id()))
	}

	res = sitemapData

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Sitemap: %s", err)
	}

	if err := d.Set("name", flattenDiscoveryEngineSitemapName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Sitemap: %s", err)
	}
	if err := d.Set("uri", flattenDiscoveryEngineSitemapUri(res["uri"], d, config)); err != nil {
		return fmt.Errorf("Error reading Sitemap: %s", err)
	}
	if err := d.Set("create_time", flattenDiscoveryEngineSitemapCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading Sitemap: %s", err)
	}

	return nil
}

func resourceDiscoveryEngineSitemapDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Sitemap: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{DiscoveryEngineBasePath}}{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting Sitemap %q", d.Id())
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "DELETE",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutDelete),
		Headers:   headers,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "Sitemap")
	}

	err = DiscoveryEngineOperationWaitTime(
		config, res, project, "Deleting Sitemap", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting Sitemap %q: %#v", d.Id(), res)
	return nil
}

func resourceDiscoveryEngineSitemapImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	// current import_formats can't import fields with forward slashes in their value
	if err := tpgresource.ParseImportId([]string{"(?P<name>.+)"}, d, config); err != nil {
		return nil, err
	}
	re := regexp.MustCompile("^projects/([^/]+)/locations/([^/]+)/collections/default_collection/dataStores/([^/]+)/siteSearchEngine/sitemaps/([^/]+)$")
	match := re.FindStringSubmatch(d.Get("name").(string))
	if len(match) > 0 {
		if err := d.Set("project", match[1]); err != nil {
			return nil, fmt.Errorf("Error setting project: %s", err)
		}
		if err := d.Set("location", match[2]); err != nil {
			return nil, fmt.Errorf("Error setting location: %s", err)
		}
		if err := d.Set("data_store_id", match[3]); err != nil {
			return nil, fmt.Errorf("Error setting data_store_id: %s", err)
		}
	} else {
		return nil, fmt.Errorf("import did not match the regex ^projects/([^/]+)/locations/([^/]+)/collections/default_collection/dataStores/([^/]+)/siteSearchEngine/sitemaps/([^/]+)$")
	}
	// Replace import id for the resource id
	id := d.Get("name").(string)
	d.SetId(id)
	return []*schema.ResourceData{d}, nil
}

func flattenDiscoveryEngineSitemapName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDiscoveryEngineSitemapUri(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDiscoveryEngineSitemapCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandDiscoveryEngineSitemapUri(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
