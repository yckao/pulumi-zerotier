// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Zerotier.Inputs
{

    public sealed class GerNetworkAssignIpv6InputArgs : Pulumi.ResourceArgs
    {
        [Input("rfc4193")]
        public Input<bool>? Rfc4193 { get; set; }

        [Input("sixplane")]
        public Input<bool>? Sixplane { get; set; }

        [Input("zerotier")]
        public Input<bool>? Zerotier { get; set; }

        public GerNetworkAssignIpv6InputArgs()
        {
        }
    }
}
