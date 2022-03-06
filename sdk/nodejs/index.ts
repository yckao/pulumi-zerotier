// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

// Export members:
export * from "./getNetwork";
export * from "./identity";
export * from "./member";
export * from "./network";
export * from "./provider";
export * from "./token";

// Export sub-modules:
import * as config from "./config";
import * as types from "./types";

export {
    config,
    types,
};

// Import resources to register:
import { Identity } from "./identity";
import { Member } from "./member";
import { Network } from "./network";
import { Token } from "./token";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "zerotier:index/identity:Identity":
                return new Identity(name, <any>undefined, { urn })
            case "zerotier:index/member:Member":
                return new Member(name, <any>undefined, { urn })
            case "zerotier:index/network:Network":
                return new Network(name, <any>undefined, { urn })
            case "zerotier:index/token:Token":
                return new Token(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("zerotier", "index/identity", _module)
pulumi.runtime.registerResourceModule("zerotier", "index/member", _module)
pulumi.runtime.registerResourceModule("zerotier", "index/network", _module)
pulumi.runtime.registerResourceModule("zerotier", "index/token", _module)

import { Provider } from "./provider";

pulumi.runtime.registerResourcePackage("zerotier", {
    version: utilities.getVersion(),
    constructProvider: (name: string, type: string, urn: string): pulumi.ProviderResource => {
        if (type !== "pulumi:providers:zerotier") {
            throw new Error(`unknown provider type ${type}`);
        }
        return new Provider(name, <any>undefined, { urn });
    },
});
