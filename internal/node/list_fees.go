package node

import (
	"fmt"
	"os"

	"github.com/lightningnetwork/lnd/lnrpc"
	"github.com/montanaflynn/stats"
	"github.com/olekukonko/tablewriter"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func ListFees(client lnrpc.LightningClient) {

	channelsResponse, err := FetchChannelData(client)
	if err != nil {
		fmt.Print("failed to fetch channels: ", err)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Pubkey", "Channel ID", "Fee PPM", "Capacity"})
	table.SetCenterSeparator("┼")
	table.SetRowLine(true)
	table.SetRowSeparator("─")
	table.SetColumnSeparator("│")
	table.SetAlignment(tablewriter.ALIGN_RIGHT)

	p := message.NewPrinter(language.English)

	ONE_MILLION := 1000000.0

	channelBalances := []float64{}
	for _, c := range channelsResponse.ChannelMap {
		tableRow := []string{c.RemotePubkey,
			fmt.Sprintf("%d", c.ChanId),
			p.Sprintf("%d", int64(channelsResponse.FeeMap[c.ChannelPoint].FeeRate*ONE_MILLION)),
			p.Sprintf("%d", c.LocalBalance+c.RemoteBalance),
		}
		table.Append(tableRow)
		channelBalances = append(channelBalances, float64(c.LocalBalance+c.RemoteBalance))
	}

	mc, _ := stats.Median(channelBalances)

	table.SetFooter([]string{" ", " ", " ", p.Sprintf("Median: %d", int64(mc))})

	table.Render()

}
