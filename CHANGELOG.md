# Changelog

## [Unreleased]

### Changed
- Added `POSTGRES_HOST_AUTH_METHOD=trust` environment variable to work with modern PostgreSQL images that require authentication configuration
- Moved nil config check before container creation
- Updated to use Go modules
