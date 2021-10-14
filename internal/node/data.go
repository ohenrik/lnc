package node

import (
	"context"
	"fmt"

	"github.com/lightningnetwork/lnd/lnrpc"
)

type ChannelMap map[string]*lnrpc.Channel
type FeeMap map[string]*lnrpc.ChannelFeeReport

type ChannelFeeMap struct {
	ChannelMap
	FeeMap
}

func createChannelMap(c []*lnrpc.Channel) ChannelMap {
	r := ChannelMap{}
	for _, v := range c {
		r[v.ChannelPoint] = v
	}
	return r
}

func createFeeMap(f []*lnrpc.ChannelFeeReport) FeeMap {
	r := FeeMap{}
	for _, v := range f {
		r[v.ChannelPoint] = v
	}
	return r
}

func FetchChannelData(client lnrpc.LightningClient) (ChannelFeeMap, error) {

	ctx := context.Background()

	channelsResponse, err := client.ListChannels(ctx, &lnrpc.ListChannelsRequest{})
	if err != nil {
		return ChannelFeeMap{}, fmt.Errorf("cannot list channels report from node: %v", err)
	}

	feeReportResponse, err := client.FeeReport(ctx, &lnrpc.FeeReportRequest{})
	if err != nil {
		return ChannelFeeMap{}, fmt.Errorf("cannot fee report from node: %v", err)
	}

	cm := createChannelMap(channelsResponse.Channels)
	fm := createFeeMap(feeReportResponse.ChannelFees)

	return ChannelFeeMap{cm, fm}, nil
}
