# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/).

## [2.0.0] - 2026-02-21

### Added

- New `Ohlc` endpoint for OHLC (Open, High, Low, Close) candlestick data (requires Tier 3 subscription)
  - Parameters: `currency` (required), `date` (required), `base` (optional, default USD), `interval` (optional: `5m`, `15m`, `30m`, `1h`, `4h`, `12h`, `1d`, default `1d`), `output` (optional)

### Changed

- Migrated from API v1 to v2 (`/api/v1` → `/api/v2`)
- Added deprecation notice for API v1 (retired 31st July 2026, redirected to v2)

## [1.2.0]

### Changed

- Updated Go toolchain to 1.26.0 (minimum version remains 1.21)
- Replaced deprecated `ioutil.ReadAll` with `io.ReadAll`
- Updated Dockerfile base image from `golang:1.21-alpine` to `golang:1.26-alpine`
- Updated GitHub Actions: `actions/setup-go@v3` to `v5`, `actions/checkout@v3` to `v4`
- Added Go 1.26 to CI test matrix alongside Go 1.21 for backward compatibility testing
- Updated minimum Go version in README from 1.13 to 1.21
