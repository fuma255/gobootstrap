package config

import (
	"log"
	"github.com/alehano/gobootstrap/sys/cmd"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/yaml.v2"
)

func init() {

	cmd.RootCmd.AddCommand(&cobra.Command{
		Use:   "version",
		Short: "Show version",
		Run: func(cmd *cobra.Command, args []string) {
			log.Printf("Version: %s\n", Version)
		},
	})

	cmd.RootCmd.AddCommand(&cobra.Command{
		Use:   "dumpconfig",
		Short: "Print config",
		Run: func(cmd *cobra.Command, args []string) {
			out, err := yaml.Marshal(Get())
			if err != nil {
				log.Println(err)
			} else {
				log.Println(string(out))
			}
		},
	})

	cmd.RootCmd.AddCommand(&cobra.Command{
		Use:   "admin_pwd [password]",
		Short: "Generate Admin password hash",
		Long: "Generate Admin password hash (BCrypt) for using in config file.",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 1 {
				hash, err := bcrypt.GenerateFromPassword([]byte(args[0]), bcrypt.DefaultCost)
				if err != nil {
					return err
				} else {
					log.Printf("Password Hash: %s\n", hash)
				}
			} else {
				log.Println("Error: password string wasn't provided")
			}
			return nil
		},
	})

}
