#!/bin/bash

# Force PATH to include go/bin
export PATH=$PATH:$HOME/go/bin

# Run gomobile with full toolchain context
$HOME/go/bin/gomobile bind -target=ios -o $HOME/tailscale/TailscaleBridge.xcframework $HOME/tailscale/mobilebridge

