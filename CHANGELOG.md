# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/).

## [Unreleased]

### Changed

- Updated Go toolchain to 1.26.0 (minimum version remains 1.21)
- Replaced deprecated `ioutil.ReadAll` with `io.ReadAll`
- Updated Dockerfile base image from `golang:1.21-alpine` to `golang:1.26-alpine`
- Updated GitHub Actions: `actions/setup-go@v3` to `v5`, `actions/checkout@v3` to `v4`
- Added Go 1.26 to CI test matrix alongside Go 1.21 for backward compatibility testing
- Updated minimum Go version in README from 1.13 to 1.21
