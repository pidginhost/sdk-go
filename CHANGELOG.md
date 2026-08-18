# Changelog

All notable changes to the PidginHost Go SDK are documented here.

This SDK is auto-generated from the [PidginHost API schema](https://www.pidginhost.com/api/schema/).
Version bumps reflect schema changes, not hand-written code changes.

## v0.13.0

### Added

- **Kubernetes port forwards can be called at all now.** All six routes under
  `/api/kubernetes/clusters/{cluster_id}/port-forwards/` generated with no
  request body and no response model, so this SDK could reach none of them.
  They now expose `KubernetesClustersPortForwardsList`, `...Retrieve`,
  `...Create`, `...Update`, `...PartialUpdate` and `...Destroy`, with the new
  `K8sPortForward`, `PatchedK8sPortForward`, `PaginatedK8sPortForwardList` and
  `ProtocolEnum` models.
- **Gateway route listings are typed and paginated**: `PaginatedHTTPRouteList`,
  `PaginatedTCPRouteList` and `PaginatedUDPRouteList`. The list and detail
  reads previously had no response schema at all.
- `NodeMetricsResponse` gained `disk` and `maxdisk`, which the endpoint has
  always returned but never documented.
- The node RRD route documents the `timeframe` query parameter it accepts
  (`hour`, `day`, `week`, `month`, `year`), so a window other than the default
  is now reachable.

### Changed

**This release is breaking.** Every change below corrects a type the API was
already sending -- the previous declarations did not describe the wire, so the
affected responses could not be decoded.

- `KubernetesClusters{Httproutes,Tcproutes,Udproutes,PortForwards}Retrieve2`
  are **removed**. Those names existed only because the list and detail reads
  collided on one operation id while the schema was incomplete. The list reads
  are now `...List`; the detail reads are `...Retrieve`.
- `KubernetesClusters{Httproutes,Tcproutes,Udproutes}Retrieve` **changed
  meaning**: it was the list read and is now the detail read. The signature
  changed with it, so this surfaces at compile time rather than silently.
- `NodeMetricsResponse.mem`, `.maxmem`, `.netin` and `.netout` widened from `int32` to `int64`.
  They are byte counts: a node with 2 GiB of memory reports 2147483648, one
  past the 32-bit range, so these responses failed to decode on essentially
  every real node.
- `ClusterDetail.dual_stack` and `.talos_upgrade_available` are **booleans**,
  `.storage_quota_gb` is an **integer**, `.last_pool_used_bytes` is a **64-bit
  integer** and `.price_per_hour` is a **number** -- all five were declared as
  strings while sending a bool or a number, which made the whole cluster object
  undecodable.

## v0.12.1

### Fixed

- **Release plumbing only -- no API or generated-code changes from v0.12.0.**
  v0.12.0 never reached PyPI: `publish-python` deleted the GitHub Actions
  workflow that uploads there, then logged an instruction to go and trigger it.
  The workflow is now versioned in the sdk repo and injected during generation,
  and the job fails if it is missing rather than assuming.

  `verify-published.py` also asked Packagist about `pidginhost/sdk` while the
  generator publishes `pidginhost/sdk-php`, so a correct PHP release was
  reported as a failure. Registry names are now read from the generated
  manifests instead of being written down twice.

  This version exists so every registry serves the same one.

## v0.12.0

### Changed (breaking)

- **The eight `/api/cloud/buckets/` routes now describe their real request and
  response bodies.** The bucket viewset advertised its fully read-only `Bucket`
  serializer as the body of every route, so `create`, `resize` and `visibility`
  documented no writable field to send, and `credentials/reveal` and
  `credentials/rotate` claimed to return `Bucket` while really returning an
  access key and secret. This SDK could not decode either credential
  response at all -- `Bucket` requires eleven properties the payload does not
  carry, so it failed with `no value given for required property id`. On
  `rotate` that lands *after* the old keys are invalidated, making the new ones
  unrecoverable.

  New components: `BucketCreate` (`name`, `quota_gb`, `public_read`),
  `BucketResize` (`quota_gb`), `BucketVisibility` (`public_read`),
  `BucketCredentials` (`bucket`, `endpoint`, `region`, `access_key`,
  `secret_key`) and `BucketCancelResponse` (`id`, `status`). `create` documents
  the `202` it really answers; `DELETE` returns `202` with a body rather than
  `204`; the two credential routes take no request body and declare
  `Cache-Control: no-store`.

  Signature changes: `CloudBucketsCredentialsRevealCreate` and
  `...RotateCreate` now return `*BucketCredentials`; `CloudBucketsDestroy`
  returns `*BucketCancelResponse` instead of only an `*http.Response`; and the
  `create`/`resize`/`visibility` request builders take `.BucketCreate(...)`,
  `.BucketResize(...)` and `.BucketVisibility(...)` in place of `.Bucket(...)`.

- **Decimal fields are now `string` instead of `float64`.** The API serialises every
  `DecimalField` as a JSON string (Django REST Framework's
  `COERCE_DECIMAL_TO_STRING` default), and the schema says so:
  `type: string, format: decimal`. This generator mapped that to `float64` anyway,
  which cannot decode `"12.50"` -- so **every response carrying a price failed
  to parse**, including `Deposit`, `Invoice`, `Subscription`, `Service`,
  `Profile.funds`, `TLD` pricing, `ClusterDetail.price_per_month` and
  `FundsBalanceResponse.balance`. 35 properties across 22 components were
  affected. The Python, JS and Rust SDKs already produced strings; only this
  one and PHP did not.

  Read the value as a string and parse it with a decimal type of your choosing.
  Do not use a binary float for money.

### Fixed

- Generator bumped to openapi-generator 7.24.0, which fixes the escaping of
  `validate:"regexp=..."` struct tags on pattern-constrained fields. No API
  shapes changed.

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
