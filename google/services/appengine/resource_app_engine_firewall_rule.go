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
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/appengine/FirewallRule.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package appengine

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
	"github.com/hashicorp/terraform-provider-google/google/verify"
)

func ResourceAppEngineFirewallRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceAppEngineFirewallRuleCreate,
		Read:   resourceAppEngineFirewallRuleRead,
		Update: resourceAppEngineFirewallRuleUpdate,
		Delete: resourceAppEngineFirewallRuleDelete,

		Importer: &schema.ResourceImporter{
			State: resourceAppEngineFirewallRuleImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			tpgresource.DefaultProviderProject,
		),

		Schema: map[string]*schema.Schema{
			"action": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: verify.ValidateEnum([]string{"UNSPECIFIED_ACTION", "ALLOW", "DENY"}),
				Description:  `The action to take if this rule matches. Possible values: ["UNSPECIFIED_ACTION", "ALLOW", "DENY"]`,
			},
			"source_range": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `IP address or range, defined using CIDR notation, of requests that this rule applies to.`,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `An optional string description of this rule.`,
			},
			"priority": {
				Type:     schema.TypeInt,
				Optional: true,
				Description: `A positive integer that defines the order of rule evaluation.
Rules with the lowest priority are evaluated first.

A default rule at priority Int32.MaxValue matches all IPv4 and
IPv6 traffic when no previous rule matches. Only the action of
this rule can be modified by the user.`,
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

func resourceAppEngineFirewallRuleCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	descriptionProp, err := expandAppEngineFirewallRuleDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	sourceRangeProp, err := expandAppEngineFirewallRuleSourceRange(d.Get("source_range"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("source_range"); !tpgresource.IsEmptyValue(reflect.ValueOf(sourceRangeProp)) && (ok || !reflect.DeepEqual(v, sourceRangeProp)) {
		obj["sourceRange"] = sourceRangeProp
	}
	actionProp, err := expandAppEngineFirewallRuleAction(d.Get("action"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("action"); !tpgresource.IsEmptyValue(reflect.ValueOf(actionProp)) && (ok || !reflect.DeepEqual(v, actionProp)) {
		obj["action"] = actionProp
	}
	priorityProp, err := expandAppEngineFirewallRulePriority(d.Get("priority"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("priority"); !tpgresource.IsEmptyValue(reflect.ValueOf(priorityProp)) && (ok || !reflect.DeepEqual(v, priorityProp)) {
		obj["priority"] = priorityProp
	}

	lockName, err := tpgresource.ReplaceVars(d, config, "apps/{{project}}")
	if err != nil {
		return err
	}
	transport_tpg.MutexStore.Lock(lockName)
	defer transport_tpg.MutexStore.Unlock(lockName)

	url, err := tpgresource.ReplaceVars(d, config, "{{AppEngineBasePath}}apps/{{project}}/firewall/ingressRules")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new FirewallRule: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for FirewallRule: %s", err)
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
		return fmt.Errorf("Error creating FirewallRule: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "apps/{{project}}/firewall/ingressRules/{{priority}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = transport_tpg.PollingWaitTime(resourceAppEngineFirewallRulePollRead(d, meta), transport_tpg.PollCheckForExistence, "Creating FirewallRule", d.Timeout(schema.TimeoutCreate), 1)
	if err != nil {
		return fmt.Errorf("Error waiting to create FirewallRule: %s", err)
	}

	log.Printf("[DEBUG] Finished creating FirewallRule %q: %#v", d.Id(), res)

	return resourceAppEngineFirewallRuleRead(d, meta)
}

func resourceAppEngineFirewallRulePollRead(d *schema.ResourceData, meta interface{}) transport_tpg.PollReadFunc {
	return func() (map[string]interface{}, error) {
		config := meta.(*transport_tpg.Config)

		url, err := tpgresource.ReplaceVars(d, config, "{{AppEngineBasePath}}apps/{{project}}/firewall/ingressRules/{{priority}}")

		if err != nil {
			return nil, err
		}

		billingProject := ""

		project, err := tpgresource.GetProject(d, config)
		if err != nil {
			return nil, fmt.Errorf("Error fetching project for FirewallRule: %s", err)
		}
		billingProject = project

		// err == nil indicates that the billing_project value was found
		if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
			billingProject = bp
		}

		userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
		if err != nil {
			return nil, err
		}

		res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
			Config:    config,
			Method:    "GET",
			Project:   billingProject,
			RawURL:    url,
			UserAgent: userAgent,
		})
		if err != nil {
			return res, err
		}
		return res, nil
	}
}

func resourceAppEngineFirewallRuleRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{AppEngineBasePath}}apps/{{project}}/firewall/ingressRules/{{priority}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for FirewallRule: %s", err)
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("AppEngineFirewallRule %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading FirewallRule: %s", err)
	}

	if err := d.Set("description", flattenAppEngineFirewallRuleDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading FirewallRule: %s", err)
	}
	if err := d.Set("source_range", flattenAppEngineFirewallRuleSourceRange(res["sourceRange"], d, config)); err != nil {
		return fmt.Errorf("Error reading FirewallRule: %s", err)
	}
	if err := d.Set("action", flattenAppEngineFirewallRuleAction(res["action"], d, config)); err != nil {
		return fmt.Errorf("Error reading FirewallRule: %s", err)
	}
	if err := d.Set("priority", flattenAppEngineFirewallRulePriority(res["priority"], d, config)); err != nil {
		return fmt.Errorf("Error reading FirewallRule: %s", err)
	}

	return nil
}

func resourceAppEngineFirewallRuleUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for FirewallRule: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	descriptionProp, err := expandAppEngineFirewallRuleDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	sourceRangeProp, err := expandAppEngineFirewallRuleSourceRange(d.Get("source_range"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("source_range"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, sourceRangeProp)) {
		obj["sourceRange"] = sourceRangeProp
	}
	actionProp, err := expandAppEngineFirewallRuleAction(d.Get("action"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("action"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, actionProp)) {
		obj["action"] = actionProp
	}
	priorityProp, err := expandAppEngineFirewallRulePriority(d.Get("priority"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("priority"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, priorityProp)) {
		obj["priority"] = priorityProp
	}

	lockName, err := tpgresource.ReplaceVars(d, config, "apps/{{project}}")
	if err != nil {
		return err
	}
	transport_tpg.MutexStore.Lock(lockName)
	defer transport_tpg.MutexStore.Unlock(lockName)

	url, err := tpgresource.ReplaceVars(d, config, "{{AppEngineBasePath}}apps/{{project}}/firewall/ingressRules/{{priority}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating FirewallRule %q: %#v", d.Id(), obj)
	headers := make(http.Header)
	updateMask := []string{}

	if d.HasChange("description") {
		updateMask = append(updateMask, "description")
	}

	if d.HasChange("source_range") {
		updateMask = append(updateMask, "sourceRange")
	}

	if d.HasChange("action") {
		updateMask = append(updateMask, "action")
	}

	if d.HasChange("priority") {
		updateMask = append(updateMask, "priority")
	}
	// updateMask is a URL parameter but not present in the schema, so ReplaceVars
	// won't set it
	url, err = transport_tpg.AddQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	// if updateMask is empty we are not updating anything so skip the post
	if len(updateMask) > 0 {
		res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
			Config:    config,
			Method:    "PATCH",
			Project:   billingProject,
			RawURL:    url,
			UserAgent: userAgent,
			Body:      obj,
			Timeout:   d.Timeout(schema.TimeoutUpdate),
			Headers:   headers,
		})

		if err != nil {
			return fmt.Errorf("Error updating FirewallRule %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating FirewallRule %q: %#v", d.Id(), res)
		}

	}

	return resourceAppEngineFirewallRuleRead(d, meta)
}

func resourceAppEngineFirewallRuleDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for FirewallRule: %s", err)
	}
	billingProject = project

	lockName, err := tpgresource.ReplaceVars(d, config, "apps/{{project}}")
	if err != nil {
		return err
	}
	transport_tpg.MutexStore.Lock(lockName)
	defer transport_tpg.MutexStore.Unlock(lockName)

	url, err := tpgresource.ReplaceVars(d, config, "{{AppEngineBasePath}}apps/{{project}}/firewall/ingressRules/{{priority}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting FirewallRule %q", d.Id())
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
		return transport_tpg.HandleNotFoundError(err, d, "FirewallRule")
	}

	log.Printf("[DEBUG] Finished deleting FirewallRule %q: %#v", d.Id(), res)
	return nil
}

func resourceAppEngineFirewallRuleImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^apps/(?P<project>[^/]+)/firewall/ingressRules/(?P<priority>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<priority>[^/]+)$",
		"^(?P<priority>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "apps/{{project}}/firewall/ingressRules/{{priority}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenAppEngineFirewallRuleDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAppEngineFirewallRuleSourceRange(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAppEngineFirewallRuleAction(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAppEngineFirewallRulePriority(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := tpgresource.StringToFixed64(strVal); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func expandAppEngineFirewallRuleDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineFirewallRuleSourceRange(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineFirewallRuleAction(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineFirewallRulePriority(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
