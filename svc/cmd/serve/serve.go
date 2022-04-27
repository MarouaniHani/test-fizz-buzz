package serve

import (
	"fmt"
	"log"
	"os"
	"test-fizz-buzz/svc/configs"
	"test-fizz-buzz/svc/server"

	"github.com/spf13/cobra"
)

var (
	Cmd *cobra.Command

	argAddress string
)

func init() {
	Cmd = &cobra.Command{
		Use:   "serve",
		Short: "Connect and begin serving requests.",
		Long:  ``,
		Run: func(Cmd *cobra.Command, args []string) {
			if err := serve(Cmd, args); err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(2)
			}
		},
	}

	Cmd.Flags().StringVarP(&argAddress, "address", "a", ":8080", "address to listen on")

}

func serve(cmd *cobra.Command, args []string) error {
	svr, err := server.NewServer(&configs.Config{
		HostPort: argAddress,
	})
	if err != nil {
		return err
	}

	log.Fatalln(svr.Run())

	return nil
}
