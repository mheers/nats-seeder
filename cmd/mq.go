package cmd

import (
	"github.com/mheers/nats-seeder/helpers"
	"github.com/spf13/cobra"
)

var (
	mqCmd = &cobra.Command{
		Use:   "mq",
		Short: "manages the mq",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	mqCreateSeedsCmd = &cobra.Command{
		Use:   "seeds",
		Short: "creates seeds for the mq that can be directly added to the .env file",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			operator, account, err := helpers.Create()
			if err != nil {
				return err
			}
			operatorSeed, err := operator.Seed()
			if err != nil {
				return err
			}
			accountSeed, err := account.Seed()
			if err != nil {
				return err
			}
			cmd.Printf("OPERATOR_SEED=\"" + string(operatorSeed) + "\"\n")
			cmd.Printf("ACCOUNT_SEED=\"" + string(accountSeed) + "\"\n")
			return nil
		},
	}
)

func init() {
	// mqCmd.AddCommand(mqMigrateCmd)
	mqCmd.AddCommand(mqCreateSeedsCmd)
}
