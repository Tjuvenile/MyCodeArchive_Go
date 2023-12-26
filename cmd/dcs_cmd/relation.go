package dcs_cmd

import (
	"MyCodeArchive_Go/cmd"
	"fmt"
	"github.com/spf13/cobra"
)

var haGroupCmd = &cobra.Command{
	Use:   "ha-group",
	Short: "use --help for more info.",
	Long:  `use --help for more info.`,
	Args:  cobra.OnlyValidArgs, //verify positonal parameters

	ValidArgs: []string{"create", "delete", "list", "query", "add-gw", "del-gw",
		"add-vip", "del-vip", "rename"},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			return
		}
	},
}

var createHaGroupCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a HA group",
	Long:  `Create a HA group`,
	Args: func(cmd *cobra.Command, args []string) error {
		fmt.Println(args)
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		// encode arguments to a JSON string
		//data, err := json.Marshal(&haGroupParams)
		//if err != nil {
		//	fmt.Printf("Encode arguments to json failed")
		//	return
		//}
		//op_entry.LeaderOpeartionEntry3("COMP_CSC", "create_ha_group",
		//	string(data))
		fmt.Println(args)
	},
}

func init() {
	cmd.RootCmd.AddCommand(haGroupCmd)

	//create
	haGroupCmd.AddCommand(createHaGroupCmd)
}
