package cmd

import (
	"os"
	"strconv"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/spf13/cobra"

	"hcc/clarinet/action/graphql/mutationParser"
	"hcc/clarinet/action/graphql/queryParser"
	"hcc/clarinet/lib/config"
	"hcc/clarinet/lib/logger"
	"hcc/clarinet/model"
)

var startIP, endIP, publicIP, privateIP, netmask, extIfaceAddr string
var protocol, description string
var externalPort, internalPort int

// aipCmd represents the aip command
var aipCmd = &cobra.Command{
	Use:     "aip",
	Short:   "Show current Adaptive IP setting",
	Long:    ``,
	Args:    cobra.MinimumNArgs(0),
	PreRunE: checkToken,
	Run: func(cmd *cobra.Command, args []string) {
		query := make(map[string]string)
		query["token"] = config.User.Token

		data, err := queryParser.AdaptiveIP(query)
		if err != nil {
			err.Println()
			return
		}

		aipData := data.(model.AdaptiveIP)
		if len(aipData.Errors) > 0 {
			for _, hrr := range aipData.Errors {
				hrr.Println()
			}
			return
		}

		t := table.NewWriter()

		ts := table.Style{
			Box: table.StyleBoxLight,
			Format: table.FormatOptions{
				Header: text.FormatUpper,
			},
			Options: table.Options{
				DrawBorder:      true,
				SeparateColumns: true,
				SeparateFooter:  true,
				SeparateHeader:  true,
				SeparateRows:    false,
			},
		}

		t.SetStyle(ts)
		t.SetOutputMirror(os.Stdout)
		t.SetTitle("Adaptive IP Settings")

		t.AppendRow(table.Row{"External iface IP", aipData.ExtIfaceAddress})
		t.AppendRow(table.Row{"Netmask", aipData.Netmask})
		t.AppendRow(table.Row{"Gateway", aipData.Gateway})
		t.AppendRow(table.Row{"Start IP Address", aipData.StartIPAddress})
		t.AppendRow(table.Row{"End IP Address", aipData.EndIPAddress})

		t.Render()

	},
}

var aipCreate = &cobra.Command{
	Use:     "create",
	Short:   "Create Adaptive IP",
	Long:    ``,
	Args:    cobra.MinimumNArgs(0),
	PreRunE: checkToken,
	Run: func(cmd *cobra.Command, args []string) {
		queryArgs := make(map[string]string)
		queryArgs["server_uuid"] = serverUUID
		queryArgs["public_ip"] = publicIP
		queryArgs["token"] = config.User.Token
		data, err := mutationParser.CreateAdaptiveIPServer(queryArgs)
		if err != nil {
			err.Println()
			return
		}

		aipServerData := data.(model.AdaptiveIPServer)
		if len(aipServerData.Errors) > 0 {
			for _, hrr := range aipServerData.Errors {
				hrr.Println()
			}
			return
		}
	},
}

var aipDelete = &cobra.Command{
	Use:     "delete",
	Short:   "Delete Adaptive IP",
	Long:    ``,
	Args:    cobra.MinimumNArgs(0),
	PreRunE: checkToken,
	Run: func(cmd *cobra.Command, args []string) {
		queryArgs := make(map[string]string)
		queryArgs["server_uuid"] = serverUUID
		queryArgs["token"] = config.User.Token
		data, err := mutationParser.DeleteAdaptiveIPServer(queryArgs)
		if err != nil {
			err.Println()
			return
		}

		aipServerData := data.(model.AdaptiveIPServer)
		if len(aipServerData.Errors) > 0 {
			for _, hrr := range aipServerData.Errors {
				hrr.Println()
			}
			return
		}

		logger.Logger.Printf("Successfully Deleted AdaptiveIP - %s\n", serverUUID)
	},
}

var aipList = &cobra.Command{
	Use:     "list",
	Short:   "Show Adaptive IP List",
	Long:    ``,
	Args:    cobra.MinimumNArgs(0),
	PreRunE: checkToken,
	Run: func(cmd *cobra.Command, args []string) {
		queryArgs := make(map[string]string)
		queryArgs["server_uuid"] = serverUUID
		queryArgs["public_ip"] = publicIP
		queryArgs["private_ip"] = privateIP
		queryArgs["private_gateway"] = gateway
		queryArgs["token"] = config.User.Token

		data, err := queryParser.ListAdaptiveIPServer(queryArgs)
		if err != nil {
			err.Println()
			return
		}

		aipServerList := data.(model.AdaptiveIPServers)
		if len(aipServerList.Errors) > 0 {
			for _, hrr := range aipServerList.Errors {
				hrr.Println()
			}
			return
		}

		t := table.NewWriter()
		t.SetStyle(table.Style{
			Name: "clarinetTableStyle",
			Box:  table.StyleBoxLight,
			Format: table.FormatOptions{
				Header: text.FormatUpper,
			},
			Options: table.Options{
				DrawBorder:      true,
				SeparateColumns: true,
				SeparateFooter:  true,
				SeparateHeader:  true,
				SeparateRows:    false,
			},
		})
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"No", "Server UUID", "Public IP", "Private IP", "Private Gateway", "Created At"})

		for index, aipServer := range aipServerList.AdaptiveIPServers {
			t.AppendRow([]interface{}{index + 1, aipServer.ServerUUID,
				aipServer.PublicIP, aipServer.PrivateIP, aipServer.PrivateGateway, aipServer.CreatedAt})
		}

		t.AppendFooter(table.Row{"Total", len(aipServerList.AdaptiveIPServers)})
		t.Render()

	},
}

var aipListAvailable = &cobra.Command{
	Use:     "available",
	Short:   "Show available Adaptive IP List",
	Long:    ``,
	Args:    cobra.MinimumNArgs(0),
	PreRunE: checkToken,
	Run: func(cmd *cobra.Command, args []string) {
		queryArgs := make(map[string]string)
		queryArgs["token"] = config.User.Token
		data, err := queryParser.ListAdaptiveIP(queryArgs)
		if err != nil {
			err.Println()
			return
		}

		availableIPList := data.(model.AvailableIPList)
		if len(availableIPList.Errors) > 0 {
			for _, hrr := range availableIPList.Errors {
				hrr.Println()
			}
			return
		}

		t := table.NewWriter()
		t.SetStyle(table.Style{
			Name: "clarinetTableStyle",
			Box:  table.StyleBoxLight,
			Format: table.FormatOptions{
				Header: text.FormatUpper,
			},
			Options: table.Options{
				DrawBorder:      true,
				SeparateColumns: false,
				SeparateFooter:  true,
				SeparateHeader:  true,
				SeparateRows:    false,
			},
		})
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"Available IP"})

		for _, aip := range availableIPList.AvailableIPs {
			t.AppendRow([]interface{}{aip})
		}

		t.AppendFooter(table.Row{"Total\t" + strconv.Itoa(len(availableIPList.AvailableIPs))})
		t.Render()
	},
}

var aipSetting = &cobra.Command{
	Use:     "setting",
	Short:   "Change Adaptive IP setting",
	Long:    ``,
	Args:    cobra.MinimumNArgs(0),
	PreRunE: checkToken,
	Run: func(cmd *cobra.Command, args []string) {
		queryArgs := make(map[string]string)
		queryArgs["ext_ifaceip_address"] = extIfaceAddr
		queryArgs["netmask"] = netmask
		queryArgs["gateway_address"] = gateway
		queryArgs["start_ip_address"] = startIP
		queryArgs["end_ip_address"] = endIP
		queryArgs["token"] = config.User.Token
		data, err := mutationParser.CreateAdaptiveIPSetting(queryArgs)
		if err != nil {
			err.Println()
			return
		}

		aipData := data.(model.AdaptiveIP)
		if len(aipData.Errors) > 0 {
			for _, hrr := range aipData.Errors {
				hrr.Println()
			}
			return
		}
	},
}

var portForwarding = &cobra.Command{
	Use:   "port",
	Short: "Create or Delete Port Forwarding",
	Long:  ``,
	Args:  cobra.MinimumNArgs(1),
}

var portForwardingCreate = &cobra.Command{
	Use:     "create",
	Short:   "Create Port Forwarding",
	Long:    ``,
	Args:    cobra.MinimumNArgs(0),
	PreRunE: checkToken,
	Run: func(cmd *cobra.Command, args []string) {
		queryArgs := make(map[string]string)
		queryArgs["server_uuid"] = serverUUID
		queryArgs["protocol"] = protocol
		queryArgs["external_port"] = strconv.Itoa(externalPort)
		queryArgs["internal_port"] = strconv.Itoa(internalPort)
		queryArgs["description"] = description
		queryArgs["token"] = config.User.Token
		data, err := mutationParser.CreatePortForwarding(queryArgs)
		if err != nil {
			err.Println()
			return
		}

		portForwardingData := data.(model.PortForwarding)
		if len(portForwardingData.Errors) > 0 {
			for _, hrr := range portForwardingData.Errors {
				hrr.Println()
			}
			return
		}
	},
}

var portForwardingDelete = &cobra.Command{
	Use:     "delete",
	Short:   "Delete Port Forwarding",
	Long:    ``,
	Args:    cobra.MinimumNArgs(0),
	PreRunE: checkToken,
	Run: func(cmd *cobra.Command, args []string) {
		queryArgs := make(map[string]string)
		queryArgs["server_uuid"] = serverUUID
		queryArgs["external_port"] = strconv.Itoa(externalPort)
		queryArgs["token"] = config.User.Token
		data, err := mutationParser.DeletePortForwarding(queryArgs)
		if err != nil {
			err.Println()
			return
		}

		portForwardingData := data.(model.PortForwarding)
		if len(portForwardingData.Errors) > 0 {
			for _, hrr := range portForwardingData.Errors {
				hrr.Println()
			}
			return
		}

		logger.Logger.Printf("Successfully Deleted port forwarding - %s\n", serverUUID)
	},
}

var portForwardingList = &cobra.Command{
	Use:     "list",
	Short:   "Show Port Forwarding List",
	Long:    ``,
	Args:    cobra.MinimumNArgs(0),
	PreRunE: checkToken,
	Run: func(cmd *cobra.Command, args []string) {
		queryArgs := make(map[string]string)
		queryArgs["server_uuid"] = serverUUID
		queryArgs["protocol"] = protocol
		queryArgs["external_port"] = strconv.Itoa(externalPort)
		queryArgs["internal_port"] = strconv.Itoa(internalPort)
		queryArgs["description"] = description
		queryArgs["token"] = config.User.Token

		data, err := queryParser.ListPortForwarding(queryArgs)
		if err != nil {
			err.Println()
			return
		}

		portForwardingList := data.(model.PortForwardingList)
		if len(portForwardingList.Errors) > 0 {
			for _, hrr := range portForwardingList.Errors {
				hrr.Println()
			}
			return
		}

		t := table.NewWriter()
		t.SetStyle(table.Style{
			Name: "clarinetTableStyle",
			Box:  table.StyleBoxLight,
			Format: table.FormatOptions{
				Header: text.FormatUpper,
			},
			Options: table.Options{
				DrawBorder:      true,
				SeparateColumns: true,
				SeparateFooter:  true,
				SeparateHeader:  true,
				SeparateRows:    false,
			},
		})
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"No", "Server UUID", "Protocol", "External Port", "Internal Port", "Description"})

		for index, portForwarding := range portForwardingList.PortForwardings {
			t.AppendRow([]interface{}{index + 1, portForwarding.ServerUUID,
				portForwarding.Protocol, portForwarding.ExternalPort, portForwarding.InternalPort, portForwarding.Description})
		}

		t.AppendFooter(table.Row{"Total", len(portForwardingList.PortForwardings)})
		t.Render()

	},
}

func ReadyAIPCmd() {
	aipCreate.Flags().StringVar(&serverUUID, "server_uuid", "", "UUID of Server")
	aipCreate.Flags().StringVar(&publicIP, "public_ip", "", "Public IP")
	aipCreate.MarkFlagRequired("server_uuid")
	aipCreate.MarkFlagRequired("public_ip")

	aipDelete.Flags().StringVar(&serverUUID, "server_uuid", "", "UUID of Server")
	aipDelete.MarkFlagRequired("server_uuid")

	aipList.Flags().StringVar(&serverUUID, "server_uuid", "", "UUID of Server")
	aipList.Flags().StringVar(&publicIP, "public_ip", "", "Public IP of Server")
	aipList.Flags().StringVar(&privateIP, "private_ip", "", "Private IP of Server")
	aipList.Flags().StringVar(&gateway, "private_gateway", "", "Private Gateway IP")

	aipSetting.Flags().StringVar(&extIfaceAddr, "ext_iface_addr", "", "")
	aipSetting.Flags().StringVar(&netmask, "netmask", "", "")
	aipSetting.Flags().StringVar(&gateway, "gateway", "", "")
	aipSetting.Flags().StringVar(&startIP, "start", "", "")
	aipSetting.Flags().StringVar(&endIP, "end", "", "")
	aipSetting.MarkFlagRequired("ext_iface_addr")
	aipSetting.MarkFlagRequired("netmask")
	aipSetting.MarkFlagRequired("gateway")
	aipSetting.MarkFlagRequired("start")
	aipSetting.MarkFlagRequired("end")

	portForwardingCreate.Flags().StringVar(&serverUUID, "server_uuid", "", "")
	portForwardingCreate.Flags().StringVar(&protocol, "protocol", "", "")
	portForwardingCreate.Flags().IntVar(&externalPort, "external_port", 0, "")
	portForwardingCreate.Flags().IntVar(&internalPort, "internal_port", 0, "")
	portForwardingCreate.Flags().StringVar(&description, "description", "", "")
	portForwardingCreate.MarkFlagRequired("server_uuid")
	portForwardingCreate.MarkFlagRequired("protocol")
	portForwardingCreate.MarkFlagRequired("external_port")
	portForwardingCreate.MarkFlagRequired("internal_port")
	portForwardingCreate.MarkFlagRequired("description")

	portForwardingDelete.Flags().StringVar(&serverUUID, "server_uuid", "", "")
	portForwardingDelete.Flags().IntVar(&externalPort, "external_port", 0, "")
	portForwardingDelete.MarkFlagRequired("server_uuid")
	portForwardingDelete.MarkFlagRequired("external_port")

	portForwardingList.Flags().StringVar(&serverUUID, "server_uuid", "", "")
	portForwardingList.Flags().StringVar(&protocol, "protocol", "", "")
	portForwardingList.Flags().IntVar(&externalPort, "external_port", 0, "")
	portForwardingList.Flags().IntVar(&internalPort, "internal_port", 0, "")
	portForwardingList.Flags().StringVar(&description, "description", "", "")
	portForwardingList.MarkFlagRequired("server_uuid")

	portForwarding.AddCommand(portForwardingCreate, portForwardingDelete, portForwardingList)

	aipCmd.AddCommand(aipCreate, aipDelete, aipList, aipListAvailable, aipSetting, portForwarding)
}
