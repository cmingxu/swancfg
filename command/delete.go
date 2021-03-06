package command

import (
	"fmt"
	"github.com/urfave/cli"
	"net/http"
	"os"
)

// NewDeleteCommand returns the CLI command for "delete"
func NewDeleteCommand() cli.Command {
	return cli.Command{
		Name:      "delete",
		Usage:     "delete application",
		ArgsUsage: "[name]",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "user",
				Usage: "Delete apps belong to user [USER]",
			},
			cli.StringFlag{
				Name:  "cluster",
				Usage: "Delete apps belong to cluster [CLUSTER]",
			},
		},
		Action: func(c *cli.Context) error {
			if err := deleteApp(c); err != nil {
				fmt.Fprintln(os.Stderr, "Error:", err)
			}
			return nil
		},
	}
}

// deleteApplication executes the "delete" command.
func deleteApp(c *cli.Context) error {
	if len(c.Args()) == 0 {
		return fmt.Errorf("name required")
	}

	if c.String("cluster") == "" {
		return fmt.Errorf("cluster required")
	}

	cluster, err := getCluster(c.String("cluster"))
	if err != nil {
		return err
	}

	if cluster == "" {
		return fmt.Errorf("cluster not found")
	}

	client := &http.Client{}
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/apps/%s", cluster, c.Args()[0]), nil)
	if err != nil {
		return fmt.Errorf("Make new request failed: %s", err.Error())
	}

	_, err = client.Do(req)

	return err
}
