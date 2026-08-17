# Changelog

All notable changes to the PidginHost Go SDK are documented here.

This SDK is auto-generated from the [PidginHost API schema](https://www.pidginhost.com/api/schema/).
Version bumps reflect schema changes, not hand-written code changes.

## v0.11.0

### Changed (breaking)

- **`AttachIPv4` → `AttachIPv4Request`, `AttachIPv6` → `AttachIPv6Request`.** The server attach routes now have distinct request and response components. `NewAttachIPv4`/`NewAttachIPv6` become `NewAttachIPv4Request`/`NewAttachIPv6Request`, and the request builders become `.AttachIPv4Request(body)`/`.AttachIPv6Request(body)`.
- **`POST /api/cloud/servers/{id}/detach-ipv4/` returns `ServerDetachIPv4Response`** instead of `DetachIPv4Response`. The field (`detached`) is unchanged, so code that binds the result and reads `.Detached` keeps compiling. `DetachIPv4Response` still serves the standalone `POST /api/cloud/ipv4/{id}/detach/` route, so `compute`-style IPv4 detach callers are unaffected.
- **Six enum components renamed off content-derived names.** Where several serializers expose a field called `status`, `type` or `priority` over different choice sets, drf-spectacular was naming the component by hashing its values, so adding a single choice silently renamed the type. These are now pinned:

  | Was | Now |
  |-----|-----|
  | `Priority3cdEnum` | `TicketPriorityEnum` |
  | `Status03cEnum` | `InvoiceStatusEnum` |
  | `Status63aEnum` | `ServiceStatusEnum` |
  | `StatusA57Enum` | `ResourceStatusEnum` |
  | `StatusEf2Enum` | `TicketStatusEnum` |
  | `Type85bEnum` | `ClusterTypeEnum` |

  Generated constants follow the type, so `STATUSA57ENUM_ACTIVE` becomes `RESOURCESTATUSENUM_ACTIVE`, and so on. This is a one-time rename: the names no longer move when a choice set gains a value. `ClusterType.type` had already been renamed once this way (`Type2faEnum` → `Type85bEnum` when `beta` was added) and now shares the pre-existing `ClusterTypeEnum` that carried the identical `dev`/`prod`/`beta` set under a second name.

### Added

- **`AttachIPv4Request.Reboot` and `AttachIPv6Request.Reboot`** (`*bool`, default `false`). A public address is written to the machine config and the guest OS only reads it while booting, so on a running server the address is inert until a restart. Set `Reboot` to have the API issue that restart.
- **`AttachIPv4Response` and `AttachIPv6Response` gain `RebootRequired` and `Rebooted`.** `Attached` alone could describe a server that answers neither ping nor SSH on the new address; these say whether a restart is still owed and whether one was performed. A stopped server reports `RebootRequired: false, Rebooted: false` and is never booted just to pick the address up, and `Reboot: true` on a rescue-mode or transitioning server keeps `RebootRequired: true` rather than rebooting out of the rescue ISO.
- **`AttachIPv6Response`** is new. `POST /api/cloud/servers/{id}/attach-ipv6/` previously echoed the request body back, so there was no `attached` flag to check.
- **Rescue mode**: `POST /api/cloud/servers/{id}/rescue/enter/` and `.../rescue/exit/`, typed as `RescueEnterQueued` / `RescueExitQueued`.
- **Boot ISOs**: `GET /api/cloud/servers/{id}/boot-isos/` with `BootISO`, `PaginatedBootISOList`, and `IsoBootRequest`.
- **Glue / personal-DNS records**: `/api/domain/domain/{domain}/dns/` and `/api/domain/domain/{domain}/dns/{name}/`, typed as `DNSGlue` and `PaginatedDNSGlueList`.
- `TicketMessage`, `ServiceCompany`, `DestroyProtectionResponse`, and `CategoryEnum`.

### Fixed

- **`attach-ipv4` reports a missing field as a missing field.** Omitting `ipv4` returned `Object with address=None does not exist`, which reads as a bad address rather than a bad field name; it is now `400 {"ipv4": ["This field is required."]}`.
- **`attach-ipv6` rejects a conflicting address with a `400`** instead of `200` plus a message, so a failed attach no longer looks like a success. Re-attaching the address a server already has is idempotent.

## v0.10.2

### Added

- `ServerProduct.AvailableGenerations []string` — which hardware generations a SKU is enabled on. Useful to render a generations column in package listings or pre-validate `--generation` choices.
- `FloatingIPv4.ReverseDns` and `FloatingIPv6.ReverseDns` — PTR record on the listing endpoints, no extra rdns fetch needed.
- `ServerDetail.FloatingIps []FloatingIPSummary` — floating IPs this VM is authorized to claim (id, version, address, label, reverse_dns). Populated from FloatingIPv{4,6}Authorization rows; complements the regular `Networks` block which only covers attached NICs.

## v0.10.1

### Fixed

- `CloudFloatingIpv4AuthorizationsList` and `CloudFloatingIpv6AuthorizationsList` restored after a regression in v0.10.0. The backend's schema annotation was dropped earlier so drf-spectacular fell back to the viewset's primary serializer, making the SDK type the response as `*FloatingIPv4`/`*FloatingIPv6` (single object) under a `Retrieve` operation. Schema now correctly advertises `PaginatedFloatingIPAuthorizationList` and the SDK methods take a `Page()` builder like other paginated endpoints.

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
