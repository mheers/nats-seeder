package cmd

import (
	"errors"
	"fmt"

	"github.com/mheers/nats-seeder/helpers"
	"github.com/spf13/cobra"
)

var (
	// OperatorSeedFlag used for setting the operator seed
	OperatorSeedFlag string
	// AccountSeedFlag used for setting the account seed
	AccountSeedFlag string
	// SeedFlag used for setting a seed
	SeedFlag string

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
			fmt.Printf("OPERATOR_SEED=\"" + string(operatorSeed) + "\"\n")
			fmt.Printf("ACCOUNT_SEED=\"" + string(accountSeed) + "\"\n")
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
			fmt.Printf("%s", operatorJWT)
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
			fmt.Printf("%s", accountJWT)
			return nil
		},
	}

	mqOperatorPublicKeyCmd = &cobra.Command{
		Use:   "operator-public-key",
		Short: "calculates the public-key for an operator",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			if OperatorSeedFlag == "" {
				return errors.New("no value for parameter --operator-seed found")
			}
			operatorKP, _, err := helpers.CreateOperator([]byte(OperatorSeedFlag))
			if err != nil {
				return err
			}
			publicKey, err := operatorKP.PublicKey()
			if err != nil {
				return err
			}
			fmt.Printf("%s", publicKey)
			return nil
		},
	}

	mqAccountPublicKeyCmd = &cobra.Command{
		Use:   "account-public-key",
		Short: "calculates the public-key for an account",
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
			accountKP, _, err := helpers.CreateAccount([]byte(AccountSeedFlag), operatorKP)
			if err != nil {
				return err
			}
			publicKey, err := accountKP.PublicKey()
			if err != nil {
				return err
			}
			fmt.Printf("%s", publicKey)
			return nil
		},
	}
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&OperatorSeedFlag, "operator-seed", "o", "", "seed for the operator")
	rootCmd.PersistentFlags().StringVarP(&AccountSeedFlag, "account-seed", "a", "", "seed for the account")
	rootCmd.AddCommand(mqCreateSeedsCmd)
	rootCmd.AddCommand(mqCreateOperatorJWTCmd)
	rootCmd.AddCommand(mqCreateAccountJWTCmd)
	rootCmd.AddCommand(mqOperatorPublicKeyCmd)
	rootCmd.AddCommand(mqAccountPublicKeyCmd)
}
