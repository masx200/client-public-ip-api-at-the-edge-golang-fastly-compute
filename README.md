# client-public-ip-api-at-the-edge-golang-fastly-compute

#### 介绍

client-public-ip-api-at-the-edge-golang-fastly-compute

#### 软件架构

软件架构说明

#### 安装教程

1. xxxx
2. xxxx
3. xxxx

#### 使用说明

1. xxxx
2. xxxx
3. xxxx

# Default Starter Kit for Go

[![Deploy to Fastly](https://deploy.edgecompute.app/button)](https://deploy.edgecompute.app/deploy)

Get to know the Fastly Compute environment with a basic starter that
demonstrates routing, simple synthetic responses and code comments that cover
common patterns.

**For more details about other starter kits for Compute, see the
[Fastly Documentation Hub](https://www.fastly.com/documentation/solutions/starters)**

## Features

- Allow only requests with particular HTTP methods
- Match request URL path and methods for routing
- Build synthetic responses at the edge

## Understanding the code

This starter is intentionally lightweight, and requires no dependencies aside
from the
[`"github.com/fastly/compute-sdk-go/fsthttp"`](https://github.com/fastly/compute-sdk-go)
repo. It will help you understand the basics of processing requests at the edge
using Fastly. This starter includes implementations of common patterns explained
in our [using Compute](https://www.fastly.com/documentation/guides/compute/go/)
and
[VCL migration](https://www.fastly.com/documentation/guides/compute/migrate/)
guides.

The starter doesn't require the use of any backends. Once deployed, you will
have a Fastly service running on Compute that can generate synthetic responses
at the edge.

It is recommended to use the [Fastly CLI](https://github.com/fastly/cli) for
this template. The template uses the `fastly.toml` scripts, to allow for
building the project using your installed Go compiler. The Fastly CLI should
also be used for serving and testing your build output, as well as deploying
your finalized package!

## Security issues

Please see our [SECURITY.md](SECURITY.md) for guidance on reporting
security-related issues.
