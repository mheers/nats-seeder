package cmd

import (
	"errors"

	"github.com/mheers/nats-seeder/helpers"
	"github.com/spf13/cobra"
)

var (
	// OperatorSeedFlag used for setting the operator seed
	OperatorSeedFlag string
	// AccountSeedFlag used for setting the account seed
	AccountSeedFlag string

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

	mqCreateOperatorJWTCmd = &cobra.Command{
		Use:   "operator-jwt",
		Short: "creates a jwt for an account",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			if OperatorSeedFlag == "" {
				return errors.New("no value for parameter --operator-seed found")
			}
			_, operatorJWT, err := helpers.CreateOperator([]byte(OperatorSeedFlag))
			if err != nil {
				return err
			}
			cmd.Printf("%s", operatorJWT)
			return nil
		},
	}

	mqCreateAccountJWTCmd = &cobra.Command{
		Use:   "account-jwt",
		Short: "creates a jwt for an account",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			if OperatorSeedFlag == "" {
				return errors.New("no value for parameter --operator-seed found")
			}
			operatorKP, _, err := helpers.CreateOperator([]byte(OperatorSeedFlag))
			if err != nil {
				return err
			}
			if AccountSeedFlag == "" {
				return errors.New("no value for parameter --account-seed found")
			}
			_, accountJWT, err := helpers.CreateAccount([]byte(AccountSeedFlag), operatorKP)
			if err != nil {
				return err
			}
			cmd.Printf("%s", accountJWT)
			return nil
		},
	}
)

func init() {
	mqCmd.PersistentFlags().StringVarP(&OperatorSeedFlag, "operator-seed", "o", "", "seed for the operator")
	mqCmd.PersistentFlags().StringVarP(&AccountSeedFlag, "account-seed", "a", "", "seed for the account")
	mqCmd.AddCommand(mqCreateSeedsCmd)
	mqCmd.AddCommand(mqCreateOperatorJWTCmd)
	mqCmd.AddCommand(mqCreateAccountJWTCmd)
}
