# SaberLight

Collection of tools to control some "smart" BLE bulbs.

## Building and installation

Install from [AUR](https://aur.archlinux.org/packages/saberlight-git/).

...or build it with `make build`. You'll need to place this repo inside your `GOPATH` and set `GO15VENDOREXPERIMENT` to `1`. Building outside `GOPATH`, even with `GO15VENDOREXPERIMENT` [does not work](https://github.com/golang/go/issues/12511).
