package info

import (
	"github.com/Benbentwo/ubuntu-server-util/pkg/cmd/common"
	"github.com/Benbentwo/utils/log"
	"github.com/digitalocean/go-smbios/smbios"
	"github.com/spf13/cobra"
)

// options for the command
type InfoOptions struct {
	*common.CommonOptions
}

func NewCmdInfo(commonOpts *common.CommonOptions) *cobra.Command {
	options := &InfoOptions{
		CommonOptions: commonOpts,
	}

	cmd := &cobra.Command{
		Use: "info",
		Run: func(cmd *cobra.Command, args []string) {
			options.Cmd = cmd
			options.Args = args
			err := options.Run()
			common.CheckErr(err)
		},
	}
	options.AddInfoFlags(cmd)
	// Section to add commands to:

	return cmd
}

// Run implements this command
func (o *InfoOptions) Run() error {
	// Find SMBIOS data in operating system-specific location.
	rc, ep, err := smbios.Stream()
	if err != nil {
		log.Logger().Fatalf("failed to open stream: %v", err)
	}
	// Be sure to close the stream!
	defer rc.Close()

	// Decode SMBIOS structures from the stream.
	d := smbios.NewDecoder(rc)
	ss, err := d.Decode()
	if err != nil {
		log.Logger().Fatalf("failed to decode structures: %v", err)
	}

	// Determine SMBIOS version and table location from entry point.
	major, minor, rev := ep.Version()
	addr, size := ep.Table()

	log.Logger().Infof("SMBIOS %d.%d.%d - table: address: %#x, size: %d\n",
		major, minor, rev, addr, size)

	for _, s := range ss {
		log.Logger().Println(s)
	}

	return nil
}

func (o *InfoOptions) AddInfoFlags(cmd *cobra.Command) {
	o.Cmd = cmd
}
