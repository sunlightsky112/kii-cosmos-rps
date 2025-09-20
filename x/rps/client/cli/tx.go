package cli

import (
	"strconv"

	"challenge/x/rps/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

func CmdCreateStudent() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-student [name] [age]",
		Short: "Create a new student",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			name := args[0]
			age, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateStudent(clientCtx.GetFromAddress().String(), name, age)
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func CmdDeleteStudent() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-student [id]",
		Short: "Delete a student by ID",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			id := args[0]

			msg := types.NewMsgDeleteStudent(clientCtx.GetFromAddress().String(), id)
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
