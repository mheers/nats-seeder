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

	// UserNameFlag used for the username
	UserNameFlag string
	// AllowPubFlag used to define channels the user will be allowed to publish
	AllowPubFlag []string
	// AllowSubFlag used to define channels the user will be allowed to subscribe
	AllowSubFlag []string

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

	mqCreateUserNkeyCmd = &cobra.Command{
		Use:   "user-nkey",
		Short: "creates a nkey for a user",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			if OperatorSeedFlag == "" {
				return errors.New("no value for parameter --operator-seed found")
			}
			if AccountSeedFlag == "" {
				return errors.New("no value for parameter --account-seed found")
			}
			name := UserNameFlag
			allowPub := AllowPubFlag
			allowSub := AllowSubFlag
			userJWT, uSeed, err := helpers.CreateUser([]byte(OperatorSeedFlag), []byte(AccountSeedFlag), name, allowPub, allowSub)
			if err != nil {
				return err
			}
			fmt.Println("-----BEGIN NATS USER JWT-----")
			fmt.Printf("%s\n", userJWT)
			fmt.Println("-----END NATS USER JWT-----")
			fmt.Println("")
			fmt.Println("-----BEGIN USER NKEY SEED-----")
			fmt.Printf("%s\n", uSeed)
			fmt.Println("-----END USER NKEY SEED-----")
			return nil
		},
	}

	mqCreateUserJWTCmd = &cobra.Command{
		Use:   "user-jwt",
		Short: "creates a jwt for a user",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			if OperatorSeedFlag == "" {
				return errors.New("no value for parameter --operator-seed found")
			}
			if AccountSeedFlag == "" {
				return errors.New("no value for parameter --account-seed found")
			}
			name := UserNameFlag
			allowPub := AllowPubFlag
			allowSub := AllowSubFlag
			userJWT, _, err := helpers.CreateUser([]byte(OperatorSeedFlag), []byte(AccountSeedFlag), name, allowPub, allowSub)
			if err != nil {
				return err
			}
			fmt.Printf("%s", userJWT)
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

	mqUserPublicKeyCmd = &cobra.Command{
		Use:   "user-public-key",
		Short: "calculates the public-key for an user",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			if OperatorSeedFlag == "" {
				return errors.New("no value for parameter --operator-seed found")
			}
			if AccountSeedFlag == "" {
				return errors.New("no value for parameter --account-seed found")
			}
			name := UserNameFlag
			allowPub := AllowPubFlag
			allowSub := AllowSubFlag
			_, uSeed, err := helpers.CreateUser([]byte(OperatorSeedFlag), []byte(AccountSeedFlag), name, allowPub, allowSub)
			if err != nil {
				return err
			}
			fmt.Printf("%s", uSeed)
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
	rootCmd.AddCommand(mqCreateUserJWTCmd)

	// mqCreateUserJWTCmd.PersistentFlags().StringArrayP()
	mqCreateUserJWTCmd.PersistentFlags().StringVarP(&UserNameFlag, "user-name", "u", "", "name for the user")
	mqCreateUserJWTCmd.PersistentFlags().StringArrayVarP(&AllowPubFlag, "allow-pub", "p", []string{}, "channels the user will be allowed to publish")
	mqCreateUserJWTCmd.PersistentFlags().StringArrayVarP(&AllowSubFlag, "allow-sub", "s", []string{}, "channels the user will be allowed to subscribe")
	rootCmd.AddCommand(mqOperatorPublicKeyCmd)
	rootCmd.AddCommand(mqAccountPublicKeyCmd)
	rootCmd.AddCommand(mqCreateUserNkeyCmd)
	rootCmd.AddCommand(mqUserPublicKeyCmd)
}
