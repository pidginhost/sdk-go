# Changelog

All notable changes to the PidginHost Go SDK are documented here.

This SDK is auto-generated from the [PidginHost API schema](https://www.pidginhost.com/api/schema/).
Version bumps reflect schema changes, not hand-written code changes.

## v0.7.0

### Added

- Floating IP support. New API groups `cloud_floating_ipv4` and `cloud_floating_ipv6` cover list, create, retrieve, destroy, plus the `authorize`, `unauthorize`, and `authorizations` actions. Floating IPs can be authorized on multiple servers simultaneously for customer-managed HA (keepalived/VRRP inside the guest).
- New model types: `FloatingIPv4`, `FloatingIPv6`, `FloatingIPv4Create`, `FloatingIPv6Create`, `FloatingIPAuthorization`, `FloatingIPAuthorizeRequest`, `FloatingIPv4AuthorizeResponse`, `FloatingIPv4UnauthorizeResponse`, `FloatingIPv6AuthorizeResponse`, `FloatingIPv6UnauthorizeResponse`, `PaginatedFloatingIPv4List`, `PaginatedFloatingIPv6List`.

## v0.6.0

### Added

- `ServerAdd.UserData` field for cloud-init startup scripts (bash with shebang or `#cloud-config` YAML), max 64 KiB, Linux images only.

## v0.5.0

### Changed

- Generated under the rebuilt CI pipeline (validate + smoke + matrix + single-click approval gate); no functional changes.

### Notes

- No SDK API changes — generated from the same schema as 0.4.x.

## v0.4.1

### Fixed

- Removed `decoder.DisallowUnknownFields()` from all model `UnmarshalJSON` methods — API responses with new fields no longer cause silent deserialization failures

## v0.4.0

### Added

- `Server.Generation` field — server hardware generation (e.g. `general-purpose`)
- New API groups: cloud generations, server packages by generation
- `ServerAdd.NoNetworkAcknowledged` field for servers without public networking

### Changed

- Regenerated from latest API schema

## v0.3.0

### Added

- Kubernetes API: clusters, node pools, nodes, HTTP/TCP/UDP routes
- Billing API: funds, deposits, invoices, services, subscriptions
- Dedicated servers API
- FreeDNS API
- Hosting API
- Support tickets API
- Domain API: registrants, transfers, nameservers, TLD listing

### Changed

- Regenerated from latest API schema with full API coverage

## v0.2.0

### Added

- Cloud compute: servers, images, packages, volumes, firewalls, IPs, networks
- Account management: profile, SSH keys, API tokens
- Convenience wrapper (`pidginhost.New(token, apiURL)`)
