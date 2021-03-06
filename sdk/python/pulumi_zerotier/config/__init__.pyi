# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

zerotierCentralToken: Optional[str]
"""
ZeroTier Central API Token; you can generate a new one at https://my.zerotier.com/account.
"""

zerotierCentralUrl: Optional[str]
"""
ZeroTier Central API endpoint. Unlikely you'll need to alter this unless you're testing ZeroTier central itself.
"""

