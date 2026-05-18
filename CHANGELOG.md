# Changelog

All notable changes to the PidginHost Go SDK are documented here.

This SDK is auto-generated from the [PidginHost API schema](https://www.pidginhost.com/api/schema/).
Version bumps reflect schema changes, not hand-written code changes.

## v0.10.0

### Added

- Floating IP rDNS: `CloudFloatingIpv4Rdns{Retrieve,Create}` and `CloudFloatingIpv6Rdns{Retrieve,Create}` for reading and setting the PTR record on a floating IPv4/IPv6 address. Floating-backed IPs used to be hidden from the regular `PublicIPv4`/`PublicIPv6` viewsets, so this is the only API path that can set their reverse DNS.

## v0.9.0

### Added

- `ApiCloudServersDetachIpv4CreateRequest.Ipv4(slug string)` builder for the new `?ipv4=<id|slug>` query parameter on `POST /cloud/servers/{id}/detach-ipv4/`. Required when the server has more than one IPv4 attached (multi-NIC support); omit to detach the primary NIC's IPv4 on single-IPv4 servers.
- New typed response models `AttachIPv4Response`, `DetachIPv4Response`, `AttachIPv6Response`, `DetachIPv6Response` carrying the boolean `Attached`/`Detached` field that the backend was already returning but the SDK previously dropped due to a schema name collision with the request body type.

### Changed

- **BREAKING**: `ApiCloudServersAttachIpv4CreateRequest.Execute()` now returns `*AttachIPv4Response` instead of `*AttachIPv4` (the same change for `AttachIpv6` / `DetachIpv4` / `DetachIpv6`). Callers that only checked the error are unaffected; callers that read `.Ipv4` on the previous return value need to switch to `.Attached` / `.Detached`.

## v0.8.0

### Added

- `ServerProduct` now carries `Cpus`, `Memory`, `DiskSize`, `Traffic` (int32) so package sizing is visible without probing an existing server.
- `ApiCloudServerPackagesListRequest.Generation(slug string)` filters the package list by hardware generation. Backend excludes free-tier-only packages when the generation is not flagged free-tier eligible.

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
