// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package zerotier

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Data source for ZeroTier networks, allowing you to find a network by ID
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-zerotier/sdk/go/zerotier"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := zerotier.GerNetwork(ctx, &GerNetworkArgs{
// 			Id: pulumi.StringRef(zerotier_network.Bobs_garage.Id),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GerNetwork(ctx *pulumi.Context, args *GerNetworkArgs, opts ...pulumi.InvokeOption) (*GerNetworkResult, error) {
	var rv GerNetworkResult
	err := ctx.Invoke("zerotier:index/gerNetwork:gerNetwork", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking gerNetwork.
type GerNetworkArgs struct {
	// IPv4 Assignment RuleSets
	AssignIpv4s []GerNetworkAssignIpv4 `pulumi:"assignIpv4s"`
	// IPv6 Assignment RuleSets
	AssignIpv6s []GerNetworkAssignIpv6 `pulumi:"assignIpv6s"`
	// Rules regarding IPv4 and IPv6 assignments
	AssignmentPools []GerNetworkAssignmentPool `pulumi:"assignmentPools"`
	// The description of the network
	Description *string `pulumi:"description"`
	// Enable broadcast packets on the network
	EnableBroadcast *bool `pulumi:"enableBroadcast"`
	// The layer 2 flow rules to apply to packets traveling across this network. Please see https://www.zerotier.com/manual/#3*4*1 for more information.
	FlowRules *string `pulumi:"flowRules"`
	// ZeroTier's internal network identifier, aka NetworkID
	Id *string `pulumi:"id"`
	// Maximum number of recipients per multicast or broadcast. Warning - Setting this to 0 will disable IPv4 communication on your network!
	MulticastLimit *int `pulumi:"multicastLimit"`
	// The name of the network
	Name *string `pulumi:"name"`
	// Whether or not the network is private.  If false, members will *NOT* need to be authorized to join.
	Private *bool `pulumi:"private"`
	// A ipv4 or ipv6 network route
	Routes []GerNetworkRoute `pulumi:"routes"`
}

// A collection of values returned by gerNetwork.
type GerNetworkResult struct {
	// IPv4 Assignment RuleSets
	AssignIpv4s []GerNetworkAssignIpv4 `pulumi:"assignIpv4s"`
	// IPv6 Assignment RuleSets
	AssignIpv6s []GerNetworkAssignIpv6 `pulumi:"assignIpv6s"`
	// Rules regarding IPv4 and IPv6 assignments
	AssignmentPools []GerNetworkAssignmentPool `pulumi:"assignmentPools"`
	// The time at which this network was created, in epoch seconds
	CreationTime int `pulumi:"creationTime"`
	// The description of the network
	Description *string `pulumi:"description"`
	// Enable broadcast packets on the network
	EnableBroadcast *bool `pulumi:"enableBroadcast"`
	// The layer 2 flow rules to apply to packets traveling across this network. Please see https://www.zerotier.com/manual/#3*4*1 for more information.
	FlowRules *string `pulumi:"flowRules"`
	// ZeroTier's internal network identifier, aka NetworkID
	Id string `pulumi:"id"`
	// Maximum number of recipients per multicast or broadcast. Warning - Setting this to 0 will disable IPv4 communication on your network!
	MulticastLimit *int `pulumi:"multicastLimit"`
	// The name of the network
	Name string `pulumi:"name"`
	// Whether or not the network is private.  If false, members will *NOT* need to be authorized to join.
	Private *bool `pulumi:"private"`
	// A ipv4 or ipv6 network route
	Routes []GerNetworkRoute `pulumi:"routes"`
}

func GerNetworkOutput(ctx *pulumi.Context, args GerNetworkOutputArgs, opts ...pulumi.InvokeOption) GerNetworkResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GerNetworkResult, error) {
			args := v.(GerNetworkArgs)
			r, err := GerNetwork(ctx, &args, opts...)
			return *r, err
		}).(GerNetworkResultOutput)
}

// A collection of arguments for invoking gerNetwork.
type GerNetworkOutputArgs struct {
	// IPv4 Assignment RuleSets
	AssignIpv4s GerNetworkAssignIpv4ArrayInput `pulumi:"assignIpv4s"`
	// IPv6 Assignment RuleSets
	AssignIpv6s GerNetworkAssignIpv6ArrayInput `pulumi:"assignIpv6s"`
	// Rules regarding IPv4 and IPv6 assignments
	AssignmentPools GerNetworkAssignmentPoolArrayInput `pulumi:"assignmentPools"`
	// The description of the network
	Description pulumi.StringPtrInput `pulumi:"description"`
	// Enable broadcast packets on the network
	EnableBroadcast pulumi.BoolPtrInput `pulumi:"enableBroadcast"`
	// The layer 2 flow rules to apply to packets traveling across this network. Please see https://www.zerotier.com/manual/#3*4*1 for more information.
	FlowRules pulumi.StringPtrInput `pulumi:"flowRules"`
	// ZeroTier's internal network identifier, aka NetworkID
	Id pulumi.StringPtrInput `pulumi:"id"`
	// Maximum number of recipients per multicast or broadcast. Warning - Setting this to 0 will disable IPv4 communication on your network!
	MulticastLimit pulumi.IntPtrInput `pulumi:"multicastLimit"`
	// The name of the network
	Name pulumi.StringPtrInput `pulumi:"name"`
	// Whether or not the network is private.  If false, members will *NOT* need to be authorized to join.
	Private pulumi.BoolPtrInput `pulumi:"private"`
	// A ipv4 or ipv6 network route
	Routes GerNetworkRouteArrayInput `pulumi:"routes"`
}

func (GerNetworkOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GerNetworkArgs)(nil)).Elem()
}

// A collection of values returned by gerNetwork.
type GerNetworkResultOutput struct{ *pulumi.OutputState }

func (GerNetworkResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GerNetworkResult)(nil)).Elem()
}

func (o GerNetworkResultOutput) ToGerNetworkResultOutput() GerNetworkResultOutput {
	return o
}

func (o GerNetworkResultOutput) ToGerNetworkResultOutputWithContext(ctx context.Context) GerNetworkResultOutput {
	return o
}

// IPv4 Assignment RuleSets
func (o GerNetworkResultOutput) AssignIpv4s() GerNetworkAssignIpv4ArrayOutput {
	return o.ApplyT(func(v GerNetworkResult) []GerNetworkAssignIpv4 { return v.AssignIpv4s }).(GerNetworkAssignIpv4ArrayOutput)
}

// IPv6 Assignment RuleSets
func (o GerNetworkResultOutput) AssignIpv6s() GerNetworkAssignIpv6ArrayOutput {
	return o.ApplyT(func(v GerNetworkResult) []GerNetworkAssignIpv6 { return v.AssignIpv6s }).(GerNetworkAssignIpv6ArrayOutput)
}

// Rules regarding IPv4 and IPv6 assignments
func (o GerNetworkResultOutput) AssignmentPools() GerNetworkAssignmentPoolArrayOutput {
	return o.ApplyT(func(v GerNetworkResult) []GerNetworkAssignmentPool { return v.AssignmentPools }).(GerNetworkAssignmentPoolArrayOutput)
}

// The time at which this network was created, in epoch seconds
func (o GerNetworkResultOutput) CreationTime() pulumi.IntOutput {
	return o.ApplyT(func(v GerNetworkResult) int { return v.CreationTime }).(pulumi.IntOutput)
}

// The description of the network
func (o GerNetworkResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GerNetworkResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Enable broadcast packets on the network
func (o GerNetworkResultOutput) EnableBroadcast() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GerNetworkResult) *bool { return v.EnableBroadcast }).(pulumi.BoolPtrOutput)
}

// The layer 2 flow rules to apply to packets traveling across this network. Please see https://www.zerotier.com/manual/#3*4*1 for more information.
func (o GerNetworkResultOutput) FlowRules() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GerNetworkResult) *string { return v.FlowRules }).(pulumi.StringPtrOutput)
}

// ZeroTier's internal network identifier, aka NetworkID
func (o GerNetworkResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GerNetworkResult) string { return v.Id }).(pulumi.StringOutput)
}

// Maximum number of recipients per multicast or broadcast. Warning - Setting this to 0 will disable IPv4 communication on your network!
func (o GerNetworkResultOutput) MulticastLimit() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GerNetworkResult) *int { return v.MulticastLimit }).(pulumi.IntPtrOutput)
}

// The name of the network
func (o GerNetworkResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GerNetworkResult) string { return v.Name }).(pulumi.StringOutput)
}

// Whether or not the network is private.  If false, members will *NOT* need to be authorized to join.
func (o GerNetworkResultOutput) Private() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GerNetworkResult) *bool { return v.Private }).(pulumi.BoolPtrOutput)
}

// A ipv4 or ipv6 network route
func (o GerNetworkResultOutput) Routes() GerNetworkRouteArrayOutput {
	return o.ApplyT(func(v GerNetworkResult) []GerNetworkRoute { return v.Routes }).(GerNetworkRouteArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GerNetworkResultOutput{})
}