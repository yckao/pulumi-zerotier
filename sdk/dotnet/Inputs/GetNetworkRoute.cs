// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Zerotier.Inputs
{

    public sealed class GetNetworkRouteArgs : Pulumi.InvokeArgs
    {
        [Input("target", required: true)]
        public string Target { get; set; } = null!;

        [Input("via")]
        public string? Via { get; set; }

        public GetNetworkRouteArgs()
        {
        }
    }
}