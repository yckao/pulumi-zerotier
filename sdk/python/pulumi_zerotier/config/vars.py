# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

import types

__config__ = pulumi.Config('zerotier')


class _ExportableConfig(types.ModuleType):
    @property
    def zerotier_central_token(self) -> Optional[str]:
        """
        ZeroTier Central API Token; you can generate a new one at https://my.zerotier.com/account.
        """
        return __config__.get('zerotierCentralToken')

    @property
    def zerotier_central_url(self) -> Optional[str]:
        """
        ZeroTier Central API endpoint. Unlikely you'll need to alter this unless you're testing ZeroTier central itself.
        """
        return __config__.get('zerotierCentralUrl')

