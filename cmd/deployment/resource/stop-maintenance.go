// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package cmddeploymentresource

import (
	"fmt"

	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
	"github.com/spf13/cobra"

	cmdutil "github.com/elastic/ecctl/cmd/util"
	"github.com/elastic/ecctl/pkg/deployment"
	"github.com/elastic/ecctl/pkg/deployment/depresource"
	"github.com/elastic/ecctl/pkg/ecctl"
)

var stopMaintCmd = &cobra.Command{
	Use:     "stop-maintenance <deployment id> --type <type> [--all|--i <instance-id>,<instance-id>]",
	Short:   "Stops maintenance mode on a deployment resource",
	PreRunE: cmdutil.MinimumNArgsAndUUID(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		resType, _ := cmd.Flags().GetString("type")
		refID, _ := cmd.Flags().GetString("ref-id")
		instanceID, _ := cmd.Flags().GetStringSlice("instance-id")
		ignoreMissing, _ := cmd.Flags().GetBool("ignore-missing")
		all, _ := cmd.Flags().GetBool("all")

		if err := cmdutil.IncompatibleFlags(cmd, "all", "instance-id"); err != nil {
			fmt.Fprintln(cmd.OutOrStderr(), err)
		}

		_, err := depresource.StopMaintenanceModeAllOrSpecified(depresource.StopInstancesParams{
			StopParams: depresource.StopParams{
				ResourceParams: deployment.ResourceParams{
					API:          ecctl.Get().API,
					DeploymentID: args[0],
					Type:         resType,
					RefID:        refID,
				},
				All: all,
			},
			InstanceIDs:   instanceID,
			IgnoreMissing: ec.Bool(ignoreMissing),
		})
		if err != nil {
			return err
		}

		return nil

	},
}

func init() {
	Command.AddCommand(stopMaintCmd)
	stopMaintCmd.Flags().Bool("all", false, "Stops maintenance mode on all instances of a defined resource type")
	stopMaintCmd.Flags().Bool("ignore-missing", false, "If set and the specified instance does not exist, then quietly proceed to the next instance")
	cmdutil.AddTypeFlag(stopMaintCmd, "Required", true)
	stopMaintCmd.MarkFlagRequired("type")
	stopMaintCmd.Flags().String("ref-id", "", "Optional deployment RefId, if not set, the RefId will be auto-discovered")
	stopMaintCmd.Flags().StringSliceP("instance-id", "i", nil, "Deployment instance IDs to use (e.g. instance-0000000001)")
}
