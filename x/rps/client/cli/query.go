package cli

import (
	"context"

	"challenge/x/rps/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"
)

func CmdGetStudent() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-student [id]",
		Short: "Get a student by ID",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			queryClient := types.NewQueryClient(clientCtx)
			id := args[0]

			res, err := queryClient.GetStudent(context.Background(), &types.QueryGetStudentRequest{Id: id})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}
	return cmd
}

func CmdListStudents() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-students",
		Short: "List all students",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.GetStudents(context.Background(), &types.QueryGetStudentsRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}
	return cmd
}
