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

package cmddeployment

import (
	"fmt"

	"github.com/spf13/cobra"

	cmdutil "github.com/elastic/ecctl/cmd/util"
	"github.com/elastic/ecctl/pkg/deployment"
	"github.com/elastic/ecctl/pkg/ecctl"
)

var resyncDeploymentCmd = &cobra.Command{
	Use:     "resync {<deployment id> | --all}",
	Short:   "Resynchronizes the search index and cache for the selected deployment or all",
	PreRunE: cmdutil.CheckInputHas1ArgsOr0ArgAndAll,
	RunE: func(cmd *cobra.Command, args []string) error {
		all, _ := cmd.Flags().GetBool("all")

		if all {
			fmt.Println("Resynchronizing all deployments")
			res, err := deployment.ResyncAll(deployment.ResyncAllParams{
				API: ecctl.Get().API,
			})
			if err != nil {
				return err
			}

			return ecctl.Get().Formatter.Format("", res)
		}

		fmt.Printf("Resynchronizing deployment: %s\n", args[0])
		return deployment.Resync(deployment.ResyncParams{
			API: ecctl.Get().API,
			ID:  args[0],
		})
	},
}

func init() {
	Command.AddCommand(resyncDeploymentCmd)
	resyncDeploymentCmd.Flags().Bool("all", false, "Resynchronizes the search index for all deployments")
}
