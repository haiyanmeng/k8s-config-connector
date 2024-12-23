// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package updatetypes

import (
	"context"
	"fmt"
	"os"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/options"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/typeupdater"

	"github.com/spf13/cobra"
)

type UpdateTypeOptions struct {
	*options.GenerateOptions

	parentNessage string // The fully qualified name of the parent prroto message of the field to be inserted
	fieldToInsert string
	ignoredFields string // TODO: could be part of GenerateOptions
	apiDirectory  string
	goPackagePath string
}

func (o *UpdateTypeOptions) InitDefaults() error {
	root, err := options.RepoRoot()
	if err != nil {
		return err
	}
	o.apiDirectory = root + "/apis/"
	o.goPackagePath = "github.com/GoogleCloudPlatform/k8s-config-connector/apis/"
	return nil
}

func (o *UpdateTypeOptions) BindFlags(cmd *cobra.Command) {
	cmd.Flags().StringVar(&o.parentNessage, "parent-message", o.parentNessage, "Fully qualified name of the proto message holding the new field. e.g. `google.cloud.bigquery.datatransfer.v1.TransferConfig`")
	cmd.Flags().StringVar(&o.fieldToInsert, "field-to-insert", o.fieldToInsert, "Name of the new field to be inserted, e.g. `schedule_options_v2`")
	// TODO: Update this flag to accept a file path pointing to the ignored fields YAML file.
	cmd.Flags().StringVar(&o.ignoredFields, "ignored-fields", o.ignoredFields, "Comma-separated list of fields to ignore")
	cmd.Flags().StringVar(&o.apiDirectory, "api-dir", o.apiDirectory, "Base directory for APIs")
	cmd.Flags().StringVar(&o.goPackagePath, "api-go-package-path", o.goPackagePath, "Package path")
}

func BuildCommand(baseOptions *options.GenerateOptions) *cobra.Command {
	opt := &UpdateTypeOptions{
		GenerateOptions: baseOptions,
	}

	if err := opt.InitDefaults(); err != nil {
		fmt.Fprintf(os.Stderr, "Error initializing defaults: %v\n", err)
		os.Exit(1)
	}

	cmd := &cobra.Command{
		Use:   "update-types",
		Short: "update KRM types for a proto service",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			if err := runTypeUpdater(ctx, opt); err != nil {
				return err
			}
			return nil
		},
	}

	opt.BindFlags(cmd)

	return cmd
}

func runTypeUpdater(ctx context.Context, opt *UpdateTypeOptions) error {
	if opt.apiDirectory == "" {
		return fmt.Errorf("--api-dir is required")
	}

	typeUpdaterOpts := &typeupdater.UpdaterOptions{
		ProtoSourcePath:       opt.GenerateOptions.ProtoSourcePath,
		ParentMessageFullName: opt.parentNessage,
		FieldToInsert:         opt.fieldToInsert,
		IgnoredFields:         opt.ignoredFields,
		APIDirectory:          opt.apiDirectory,
		GoPackagePath:         opt.goPackagePath,
	}
	updater := typeupdater.NewTypeUpdater(typeUpdaterOpts)
	if err := updater.Run(); err != nil {
		return err
	}
	return nil
}
