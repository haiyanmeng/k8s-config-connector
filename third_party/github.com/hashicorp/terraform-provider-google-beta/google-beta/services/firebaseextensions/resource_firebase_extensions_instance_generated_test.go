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

package firebaseextensions_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func TestAccFirebaseExtensionsInstance_firebaseExtentionsInstanceResizeImageExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project_id":    envvar.GetTestProjectFromEnv(),
		"location":      "us-central1",
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckFirebaseExtensionsInstanceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccFirebaseExtensionsInstance_firebaseExtentionsInstanceResizeImageExample(context),
			},
			{
				ResourceName:            "google_firebase_extensions_instance.resize_image",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"instance_id"},
			},
		},
	})
}

func testAccFirebaseExtensionsInstance_firebaseExtentionsInstanceResizeImageExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_storage_bucket" "images" {
  provider                    = google-beta
  project                     = "%{project_id}"
  name                        = "tf-test-bucket-id%{random_suffix}"
  location                    = "US"
  uniform_bucket_level_access = true

  # Delete all objects when the bucket is deleted
  force_destroy = true
}

resource "google_firebase_extensions_instance" "resize_image" {
  provider = google-beta
  project  = "%{project_id}"
  instance_id = "tf-test-storage-resize-images%{random_suffix}"
  config {
    extension_ref = "firebase/storage-resize-images"
    extension_version = "0.1.37"

    # The following params apply to the firebase/storage-resize-images extension. 
    # Different extensions may have different params
    params = {
      DELETE_ORIGINAL_FILE = false
      MAKE_PUBLIC          = false
      IMAGE_TYPE           = false
      IS_ANIMATED          = true
      FUNCTION_MEMORY      = 1024
      DO_BACKFILL          = false
      IMG_SIZES            = "200x200"
      IMG_BUCKET           = google_storage_bucket.images.name
      LOCATION             = "%{location}"
    }

    system_params = {
      "firebaseextensions.v1beta.function/maxInstances"               = 3000
      "firebaseextensions.v1beta.function/memory"                     = 256
      "firebaseextensions.v1beta.function/minInstances"               = 0
      "firebaseextensions.v1beta.function/vpcConnectorEgressSettings" = "VPC_CONNECTOR_EGRESS_SETTINGS_UNSPECIFIED"
    }

    allowed_event_types = [
      "firebase.extensions.storage-resize-images.v1.complete"
    ]

    eventarc_channel = "projects/%{project_id}/locations/%{location}/channels/firebase"
  }
}
`, context)
}

func testAccCheckFirebaseExtensionsInstanceDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_firebase_extensions_instance" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{FirebaseExtensionsBasePath}}projects/{{project}}/instances/{{instance_id}}")
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
				return fmt.Errorf("FirebaseExtensionsInstance still exists at %s", url)
			}
		}

		return nil
	}
}