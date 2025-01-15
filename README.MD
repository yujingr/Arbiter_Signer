# Arbiter Signer (Community Edition)

[![Go Report Card](https://goreportcard.com/badge/github.com/yujingr/Arbiter_Signer)](https://goreportcard.com/report/github.com/yujingr/Arbiter_Signer)
[![Build Status](https://github.com/yujingr/Arbiter_Signer/workflows/Build%20and%20Test/badge.svg)](https://github.com/yujingr/Arbiter_Signer/actions)
[![MIT License](https://img.shields.io/badge/license-MIT-blue.svg)](http://copyfree.org)

## Introduction

This is a community-maintained fork of the official [BeL2 Arbiter Signer](https://github.com/BeL2Labs/Arbiter_Signer) that simplifies the setup process for Elastos/BeL2 community members. It provides pre-built executables for various operating systems, allowing users to quickly test and participate in the BeL2 ecosystem.

## ⚠️ Risk Disclosure

Please read the following risk disclosures carefully before running an arbiter node:

1. **Uptime Requirements**:

   - Arbiters must maintain high availability (24/7 uptime)
   - If selected for arbitration while inactive, you risk losing part or all of your staked assets
   - The exact time window for response is to be determined, but is likely measured in hours

2. **Recommended Setup**:

   - A cloud server with 24/7 uptime is strongly recommended
   - Implement robust monitoring systems to alert you of any downtime
   - Have backup systems and procedures in place

3. **Community Edition Notice**:
   - This is a community-maintained version focused on ease of setup
   - While we strive for reliability, use at your own discretion
   - For production deployments, consider using the [official repository](https://github.com/BeL2Labs/Arbiter_Signer)

## Quick Start

1. Download the latest release for your operating system from the [Releases](https://github.com/yujingr/Arbiter_Signer/releases) section
2. Run the executable to start the interactive setup process
3. Follow the prompts to configure your arbiter node

## Demo

Watch our setup demonstration video to see how to configure and run your arbiter node:

[![Arbiter Signer Setup Demo](https://img.youtube.com/vi/VYMq8LSK0iI/0.jpg)](https://youtu.be/VYMq8LSK0iI)

## Requirements

- A stable internet connection
- 24/7 system availability
- Monitoring system (recommended)
- Sufficient stake as required by the BeL2 protocol

## Configuration

The setup process will guide you through creating a `config.yaml` file with the following key parameters, there are some default values, you can change them as you need:

### Chain API

1. **esc**: ESC chain API endpoint (default: "https://api.elastos.io/esc")

### Arbiter Settings

1. **listener**: Controls event monitoring from the ESC arbiter contract (default: true)
2. **signer**: Controls BTC transaction signing and ESC contract submission (default: true)
3. **network**: Specifies "mainnet" or "testnet" for BTC address generation (default: "mainnet")
4. **escStartHeight**: Starting block height for event monitoring (default: 28437808)
5. **escArbiterContractAddress**: Current arbiter contract address (default: "0xA10b92006743Ef3B12077da67e465963743b03D3")
6. **escArbiterManagerContractAddress**: Arbiter manager contract address (default: "0x9963b5214434776D043A4e98Bc7f33321F6aaCfc")
7. **dataPath**: Path for event file storage (default: "./app/arbiter/data")
8. **keyFilePath**: Path for BTC and ESC keystores (default: "./app/arbiter/data/keys/")
9. **escArbiterAddress**: Your arbiter wallet address (required)
10. **escPrivateKey**: Your ESC private key (required)
11. **btcPrivateKey**: Your BTC private key (required)

## Advanced Setup

For production deployments or advanced configurations, please refer to:

- [Official Repository](https://github.com/BeL2Labs/Arbiter_Signer)
- [Deployment Guide](https://github.com/BeL2Labs/Arbiter_Signer/blob/main/docs/deploy_loan_arbiter.md)

## License

This project is licensed under the [copyfree](http://copyfree.org) MIT License.
