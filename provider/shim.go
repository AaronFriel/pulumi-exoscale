package provider

import (
	"context"

	pftfbridge "github.com/pulumi/pulumi-terraform-bridge/pf/tfbridge"
	shim "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
	upstream "github.com/exoscale/terraform-provider-exoscale/pulumi-shim"
)

func ShimmedProvider() shim.Provider {
	return pftfbridge.MuxShimWithPF(
		context.Background(),
		shimv2.NewProvider(upstream.SDKProvider()),
		upstream.PFProvider(),
	)
}

func TfbridgeMain(pulumiSchema []byte, bridgeMetadata []byte) {
	meta := pftfbridge.ProviderMetadata{
		PackageSchema:  pulumiSchema,
		BridgeMetadata: bridgeMetadata,
	}
	pftfbridge.Main(context.Background(), "exoscale", Provider(), meta)
}
