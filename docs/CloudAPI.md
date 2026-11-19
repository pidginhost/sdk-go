# \CloudAPI

All URIs are relative to *https://www.pidginhost.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudBucketsCreate**](CloudAPI.md#CloudBucketsCreate) | **Post** /api/cloud/buckets/ | 
[**CloudBucketsCredentialsRevealCreate**](CloudAPI.md#CloudBucketsCredentialsRevealCreate) | **Post** /api/cloud/buckets/{id}/credentials/reveal/ | 
[**CloudBucketsCredentialsRotateCreate**](CloudAPI.md#CloudBucketsCredentialsRotateCreate) | **Post** /api/cloud/buckets/{id}/credentials/rotate/ | 
[**CloudBucketsDestroy**](CloudAPI.md#CloudBucketsDestroy) | **Delete** /api/cloud/buckets/{id}/ | 
[**CloudBucketsDestroyProtectionCreate**](CloudAPI.md#CloudBucketsDestroyProtectionCreate) | **Post** /api/cloud/buckets/{id}/destroy-protection/ | 
[**CloudBucketsList**](CloudAPI.md#CloudBucketsList) | **Get** /api/cloud/buckets/ | 
[**CloudBucketsResizeCreate**](CloudAPI.md#CloudBucketsResizeCreate) | **Post** /api/cloud/buckets/{id}/resize/ | 
[**CloudBucketsRetrieve**](CloudAPI.md#CloudBucketsRetrieve) | **Get** /api/cloud/buckets/{id}/ | 
[**CloudBucketsVisibilityCreate**](CloudAPI.md#CloudBucketsVisibilityCreate) | **Post** /api/cloud/buckets/{id}/visibility/ | 
[**CloudFirewallRulesSetCreate**](CloudAPI.md#CloudFirewallRulesSetCreate) | **Post** /api/cloud/firewall-rules-set/ | 
[**CloudFirewallRulesSetDestroy**](CloudAPI.md#CloudFirewallRulesSetDestroy) | **Delete** /api/cloud/firewall-rules-set/{id}/ | 
[**CloudFirewallRulesSetList**](CloudAPI.md#CloudFirewallRulesSetList) | **Get** /api/cloud/firewall-rules-set/ | 
[**CloudFirewallRulesSetPartialUpdate**](CloudAPI.md#CloudFirewallRulesSetPartialUpdate) | **Patch** /api/cloud/firewall-rules-set/{id}/ | 
[**CloudFirewallRulesSetRetrieve**](CloudAPI.md#CloudFirewallRulesSetRetrieve) | **Get** /api/cloud/firewall-rules-set/{id}/ | 
[**CloudFirewallRulesSetRulesCreate**](CloudAPI.md#CloudFirewallRulesSetRulesCreate) | **Post** /api/cloud/firewall-rules-set/{rules_set_id}/rules/ | 
[**CloudFirewallRulesSetRulesDestroy**](CloudAPI.md#CloudFirewallRulesSetRulesDestroy) | **Delete** /api/cloud/firewall-rules-set/{rules_set_id}/rules/{rule_id}/ | 
[**CloudFirewallRulesSetRulesList**](CloudAPI.md#CloudFirewallRulesSetRulesList) | **Get** /api/cloud/firewall-rules-set/{rules_set_id}/rules/ | 
[**CloudFirewallRulesSetRulesPartialUpdate**](CloudAPI.md#CloudFirewallRulesSetRulesPartialUpdate) | **Patch** /api/cloud/firewall-rules-set/{rules_set_id}/rules/{rule_id}/ | 
[**CloudFirewallRulesSetRulesRetrieve**](CloudAPI.md#CloudFirewallRulesSetRulesRetrieve) | **Get** /api/cloud/firewall-rules-set/{rules_set_id}/rules/{rule_id}/ | 
[**CloudFirewallRulesSetRulesUpdate**](CloudAPI.md#CloudFirewallRulesSetRulesUpdate) | **Put** /api/cloud/firewall-rules-set/{rules_set_id}/rules/{rule_id}/ | 
[**CloudFirewallRulesSetUpdate**](CloudAPI.md#CloudFirewallRulesSetUpdate) | **Put** /api/cloud/firewall-rules-set/{id}/ | 
[**CloudFloatingIpv4AuthorizationsList**](CloudAPI.md#CloudFloatingIpv4AuthorizationsList) | **Get** /api/cloud/floating-ipv4/{id}/authorizations/ | 
[**CloudFloatingIpv4AuthorizeCreate**](CloudAPI.md#CloudFloatingIpv4AuthorizeCreate) | **Post** /api/cloud/floating-ipv4/{id}/authorize/ | 
[**CloudFloatingIpv4Create**](CloudAPI.md#CloudFloatingIpv4Create) | **Post** /api/cloud/floating-ipv4/ | 
[**CloudFloatingIpv4Destroy**](CloudAPI.md#CloudFloatingIpv4Destroy) | **Delete** /api/cloud/floating-ipv4/{id}/ | 
[**CloudFloatingIpv4List**](CloudAPI.md#CloudFloatingIpv4List) | **Get** /api/cloud/floating-ipv4/ | 
[**CloudFloatingIpv4RdnsCreate**](CloudAPI.md#CloudFloatingIpv4RdnsCreate) | **Post** /api/cloud/floating-ipv4/{id}/rdns/ | 
[**CloudFloatingIpv4RdnsRetrieve**](CloudAPI.md#CloudFloatingIpv4RdnsRetrieve) | **Get** /api/cloud/floating-ipv4/{id}/rdns/ | 
[**CloudFloatingIpv4Retrieve**](CloudAPI.md#CloudFloatingIpv4Retrieve) | **Get** /api/cloud/floating-ipv4/{id}/ | 
[**CloudFloatingIpv4UnauthorizeCreate**](CloudAPI.md#CloudFloatingIpv4UnauthorizeCreate) | **Post** /api/cloud/floating-ipv4/{id}/unauthorize/ | 
[**CloudFloatingIpv6AuthorizationsList**](CloudAPI.md#CloudFloatingIpv6AuthorizationsList) | **Get** /api/cloud/floating-ipv6/{id}/authorizations/ | 
[**CloudFloatingIpv6AuthorizeCreate**](CloudAPI.md#CloudFloatingIpv6AuthorizeCreate) | **Post** /api/cloud/floating-ipv6/{id}/authorize/ | 
[**CloudFloatingIpv6Create**](CloudAPI.md#CloudFloatingIpv6Create) | **Post** /api/cloud/floating-ipv6/ | 
[**CloudFloatingIpv6Destroy**](CloudAPI.md#CloudFloatingIpv6Destroy) | **Delete** /api/cloud/floating-ipv6/{id}/ | 
[**CloudFloatingIpv6List**](CloudAPI.md#CloudFloatingIpv6List) | **Get** /api/cloud/floating-ipv6/ | 
[**CloudFloatingIpv6RdnsCreate**](CloudAPI.md#CloudFloatingIpv6RdnsCreate) | **Post** /api/cloud/floating-ipv6/{id}/rdns/ | 
[**CloudFloatingIpv6RdnsRetrieve**](CloudAPI.md#CloudFloatingIpv6RdnsRetrieve) | **Get** /api/cloud/floating-ipv6/{id}/rdns/ | 
[**CloudFloatingIpv6Retrieve**](CloudAPI.md#CloudFloatingIpv6Retrieve) | **Get** /api/cloud/floating-ipv6/{id}/ | 
[**CloudFloatingIpv6UnauthorizeCreate**](CloudAPI.md#CloudFloatingIpv6UnauthorizeCreate) | **Post** /api/cloud/floating-ipv6/{id}/unauthorize/ | 
[**CloudGatewaysCreate**](CloudAPI.md#CloudGatewaysCreate) | **Post** /api/cloud/gateways/ | 
[**CloudGatewaysDestroy**](CloudAPI.md#CloudGatewaysDestroy) | **Delete** /api/cloud/gateways/{id}/ | 
[**CloudGatewaysFirewallRulesCreate**](CloudAPI.md#CloudGatewaysFirewallRulesCreate) | **Post** /api/cloud/gateways/{gateway_id}/firewall-rules/ | 
[**CloudGatewaysFirewallRulesDestroy**](CloudAPI.md#CloudGatewaysFirewallRulesDestroy) | **Delete** /api/cloud/gateways/{gateway_id}/firewall-rules/{id}/ | 
[**CloudGatewaysFirewallRulesList**](CloudAPI.md#CloudGatewaysFirewallRulesList) | **Get** /api/cloud/gateways/{gateway_id}/firewall-rules/ | 
[**CloudGatewaysFirewallRulesMoveCreate**](CloudAPI.md#CloudGatewaysFirewallRulesMoveCreate) | **Post** /api/cloud/gateways/{gateway_id}/firewall-rules/{id}/move/ | 
[**CloudGatewaysFirewallRulesPartialUpdate**](CloudAPI.md#CloudGatewaysFirewallRulesPartialUpdate) | **Patch** /api/cloud/gateways/{gateway_id}/firewall-rules/{id}/ | 
[**CloudGatewaysFirewallRulesRetrieve**](CloudAPI.md#CloudGatewaysFirewallRulesRetrieve) | **Get** /api/cloud/gateways/{gateway_id}/firewall-rules/{id}/ | 
[**CloudGatewaysList**](CloudAPI.md#CloudGatewaysList) | **Get** /api/cloud/gateways/ | 
[**CloudGatewaysListenersCreate**](CloudAPI.md#CloudGatewaysListenersCreate) | **Post** /api/cloud/gateways/{gateway_id}/listeners/ | 
[**CloudGatewaysListenersDestroy**](CloudAPI.md#CloudGatewaysListenersDestroy) | **Delete** /api/cloud/gateways/{gateway_id}/listeners/{id}/ | 
[**CloudGatewaysListenersList**](CloudAPI.md#CloudGatewaysListenersList) | **Get** /api/cloud/gateways/{gateway_id}/listeners/ | 
[**CloudGatewaysListenersPartialUpdate**](CloudAPI.md#CloudGatewaysListenersPartialUpdate) | **Patch** /api/cloud/gateways/{gateway_id}/listeners/{id}/ | 
[**CloudGatewaysListenersRetrieve**](CloudAPI.md#CloudGatewaysListenersRetrieve) | **Get** /api/cloud/gateways/{gateway_id}/listeners/{id}/ | 
[**CloudGatewaysListenersRoutesCreate**](CloudAPI.md#CloudGatewaysListenersRoutesCreate) | **Post** /api/cloud/gateways/{gateway_id}/listeners/{listener_id}/routes/ | 
[**CloudGatewaysListenersRoutesDestroy**](CloudAPI.md#CloudGatewaysListenersRoutesDestroy) | **Delete** /api/cloud/gateways/{gateway_id}/listeners/{listener_id}/routes/{id}/ | 
[**CloudGatewaysListenersRoutesList**](CloudAPI.md#CloudGatewaysListenersRoutesList) | **Get** /api/cloud/gateways/{gateway_id}/listeners/{listener_id}/routes/ | 
[**CloudGatewaysListenersRoutesPartialUpdate**](CloudAPI.md#CloudGatewaysListenersRoutesPartialUpdate) | **Patch** /api/cloud/gateways/{gateway_id}/listeners/{listener_id}/routes/{id}/ | 
[**CloudGatewaysListenersRoutesRetrieve**](CloudAPI.md#CloudGatewaysListenersRoutesRetrieve) | **Get** /api/cloud/gateways/{gateway_id}/listeners/{listener_id}/routes/{id}/ | 
[**CloudGatewaysNetworksCreate**](CloudAPI.md#CloudGatewaysNetworksCreate) | **Post** /api/cloud/gateways/{gateway_id}/networks/ | 
[**CloudGatewaysNetworksDestroy**](CloudAPI.md#CloudGatewaysNetworksDestroy) | **Delete** /api/cloud/gateways/{gateway_id}/networks/{network_id}/ | 
[**CloudGatewaysNetworksList**](CloudAPI.md#CloudGatewaysNetworksList) | **Get** /api/cloud/gateways/{gateway_id}/networks/ | 
[**CloudGatewaysNetworksPartialUpdate**](CloudAPI.md#CloudGatewaysNetworksPartialUpdate) | **Patch** /api/cloud/gateways/{gateway_id}/networks/{network_id}/ | 
[**CloudGatewaysNetworksRetrieve**](CloudAPI.md#CloudGatewaysNetworksRetrieve) | **Get** /api/cloud/gateways/{gateway_id}/networks/{network_id}/ | 
[**CloudGatewaysOperationsList**](CloudAPI.md#CloudGatewaysOperationsList) | **Get** /api/cloud/gateways/{gateway_id}/operations/ | 
[**CloudGatewaysOperationsRetrieve**](CloudAPI.md#CloudGatewaysOperationsRetrieve) | **Get** /api/cloud/gateways/{gateway_id}/operations/{id}/ | 
[**CloudGatewaysPartialUpdate**](CloudAPI.md#CloudGatewaysPartialUpdate) | **Patch** /api/cloud/gateways/{id}/ | 
[**CloudGatewaysPoolsBackendsCreate**](CloudAPI.md#CloudGatewaysPoolsBackendsCreate) | **Post** /api/cloud/gateways/{gateway_id}/pools/{pool_id}/backends/ | 
[**CloudGatewaysPoolsBackendsDestroy**](CloudAPI.md#CloudGatewaysPoolsBackendsDestroy) | **Delete** /api/cloud/gateways/{gateway_id}/pools/{pool_id}/backends/{id}/ | 
[**CloudGatewaysPoolsBackendsList**](CloudAPI.md#CloudGatewaysPoolsBackendsList) | **Get** /api/cloud/gateways/{gateway_id}/pools/{pool_id}/backends/ | 
[**CloudGatewaysPoolsBackendsPartialUpdate**](CloudAPI.md#CloudGatewaysPoolsBackendsPartialUpdate) | **Patch** /api/cloud/gateways/{gateway_id}/pools/{pool_id}/backends/{id}/ | 
[**CloudGatewaysPoolsBackendsRetrieve**](CloudAPI.md#CloudGatewaysPoolsBackendsRetrieve) | **Get** /api/cloud/gateways/{gateway_id}/pools/{pool_id}/backends/{id}/ | 
[**CloudGatewaysPoolsCreate**](CloudAPI.md#CloudGatewaysPoolsCreate) | **Post** /api/cloud/gateways/{gateway_id}/pools/ | 
[**CloudGatewaysPoolsDestroy**](CloudAPI.md#CloudGatewaysPoolsDestroy) | **Delete** /api/cloud/gateways/{gateway_id}/pools/{id}/ | 
[**CloudGatewaysPoolsList**](CloudAPI.md#CloudGatewaysPoolsList) | **Get** /api/cloud/gateways/{gateway_id}/pools/ | 
[**CloudGatewaysPoolsPartialUpdate**](CloudAPI.md#CloudGatewaysPoolsPartialUpdate) | **Patch** /api/cloud/gateways/{gateway_id}/pools/{id}/ | 
[**CloudGatewaysPoolsRetrieve**](CloudAPI.md#CloudGatewaysPoolsRetrieve) | **Get** /api/cloud/gateways/{gateway_id}/pools/{id}/ | 
[**CloudGatewaysPortForwardsCreate**](CloudAPI.md#CloudGatewaysPortForwardsCreate) | **Post** /api/cloud/gateways/{gateway_id}/port-forwards/ | 
[**CloudGatewaysPortForwardsDestroy**](CloudAPI.md#CloudGatewaysPortForwardsDestroy) | **Delete** /api/cloud/gateways/{gateway_id}/port-forwards/{id}/ | 
[**CloudGatewaysPortForwardsList**](CloudAPI.md#CloudGatewaysPortForwardsList) | **Get** /api/cloud/gateways/{gateway_id}/port-forwards/ | 
[**CloudGatewaysPortForwardsPartialUpdate**](CloudAPI.md#CloudGatewaysPortForwardsPartialUpdate) | **Patch** /api/cloud/gateways/{gateway_id}/port-forwards/{id}/ | 
[**CloudGatewaysPortForwardsRetrieve**](CloudAPI.md#CloudGatewaysPortForwardsRetrieve) | **Get** /api/cloud/gateways/{gateway_id}/port-forwards/{id}/ | 
[**CloudGatewaysReenrollCreate**](CloudAPI.md#CloudGatewaysReenrollCreate) | **Post** /api/cloud/gateways/{id}/reenroll/ | 
[**CloudGatewaysRetrieve**](CloudAPI.md#CloudGatewaysRetrieve) | **Get** /api/cloud/gateways/{id}/ | 
[**CloudGatewaysStatusRetrieve**](CloudAPI.md#CloudGatewaysStatusRetrieve) | **Get** /api/cloud/gateways/{id}/status/ | 
[**CloudGenerationsList**](CloudAPI.md#CloudGenerationsList) | **Get** /api/cloud/generations/ | List hardware generations
[**CloudGenerationsRetrieve**](CloudAPI.md#CloudGenerationsRetrieve) | **Get** /api/cloud/generations/{slug}/ | 
[**CloudImagesList**](CloudAPI.md#CloudImagesList) | **Get** /api/cloud/images/ | 
[**CloudImagesRetrieve**](CloudAPI.md#CloudImagesRetrieve) | **Get** /api/cloud/images/{id}/ | 
[**CloudIpv4Create**](CloudAPI.md#CloudIpv4Create) | **Post** /api/cloud/ipv4/ | 
[**CloudIpv4Destroy**](CloudAPI.md#CloudIpv4Destroy) | **Delete** /api/cloud/ipv4/{id}/ | 
[**CloudIpv4DetachCreate**](CloudAPI.md#CloudIpv4DetachCreate) | **Post** /api/cloud/ipv4/{id}/detach/ | 
[**CloudIpv4List**](CloudAPI.md#CloudIpv4List) | **Get** /api/cloud/ipv4/ | 
[**CloudIpv4RdnsCreate**](CloudAPI.md#CloudIpv4RdnsCreate) | **Post** /api/cloud/ipv4/{id}/rdns/ | 
[**CloudIpv4RdnsRetrieve**](CloudAPI.md#CloudIpv4RdnsRetrieve) | **Get** /api/cloud/ipv4/{id}/rdns/ | 
[**CloudIpv4Retrieve**](CloudAPI.md#CloudIpv4Retrieve) | **Get** /api/cloud/ipv4/{id}/ | 
[**CloudIpv6Create**](CloudAPI.md#CloudIpv6Create) | **Post** /api/cloud/ipv6/ | 
[**CloudIpv6Destroy**](CloudAPI.md#CloudIpv6Destroy) | **Delete** /api/cloud/ipv6/{id}/ | 
[**CloudIpv6DetachCreate**](CloudAPI.md#CloudIpv6DetachCreate) | **Post** /api/cloud/ipv6/{id}/detach/ | 
[**CloudIpv6List**](CloudAPI.md#CloudIpv6List) | **Get** /api/cloud/ipv6/ | 
[**CloudIpv6RdnsCreate**](CloudAPI.md#CloudIpv6RdnsCreate) | **Post** /api/cloud/ipv6/{id}/rdns/ | 
[**CloudIpv6RdnsRetrieve**](CloudAPI.md#CloudIpv6RdnsRetrieve) | **Get** /api/cloud/ipv6/{id}/rdns/ | 
[**CloudIpv6Retrieve**](CloudAPI.md#CloudIpv6Retrieve) | **Get** /api/cloud/ipv6/{id}/ | 
[**CloudPrivateNetworksAddServerCreate**](CloudAPI.md#CloudPrivateNetworksAddServerCreate) | **Post** /api/cloud/private-networks/{id}/add-server/ | 
[**CloudPrivateNetworksApplyRouteCreate**](CloudAPI.md#CloudPrivateNetworksApplyRouteCreate) | **Post** /api/cloud/private-networks/{id}/apply-route/ | 
[**CloudPrivateNetworksCreate**](CloudAPI.md#CloudPrivateNetworksCreate) | **Post** /api/cloud/private-networks/ | 
[**CloudPrivateNetworksDestroy**](CloudAPI.md#CloudPrivateNetworksDestroy) | **Delete** /api/cloud/private-networks/{id}/ | 
[**CloudPrivateNetworksList**](CloudAPI.md#CloudPrivateNetworksList) | **Get** /api/cloud/private-networks/ | 
[**CloudPrivateNetworksPartialUpdate**](CloudAPI.md#CloudPrivateNetworksPartialUpdate) | **Patch** /api/cloud/private-networks/{id}/ | 
[**CloudPrivateNetworksRemoveServerCreate**](CloudAPI.md#CloudPrivateNetworksRemoveServerCreate) | **Post** /api/cloud/private-networks/{id}/remove-server/ | 
[**CloudPrivateNetworksRetrieve**](CloudAPI.md#CloudPrivateNetworksRetrieve) | **Get** /api/cloud/private-networks/{id}/ | 
[**CloudPrivateNetworksRouteStatusRetrieve**](CloudAPI.md#CloudPrivateNetworksRouteStatusRetrieve) | **Get** /api/cloud/private-networks/{id}/route-status/ | 
[**CloudPrivateNetworksUpdate**](CloudAPI.md#CloudPrivateNetworksUpdate) | **Put** /api/cloud/private-networks/{id}/ | 
[**CloudServerPackagesByGenerationRetrieve**](CloudAPI.md#CloudServerPackagesByGenerationRetrieve) | **Get** /api/cloud/server-packages/by-generation/ | 
[**CloudServerPackagesList**](CloudAPI.md#CloudServerPackagesList) | **Get** /api/cloud/server-packages/ | 
[**CloudServerPackagesRetrieve**](CloudAPI.md#CloudServerPackagesRetrieve) | **Get** /api/cloud/server-packages/{id}/ | 
[**CloudServersActivityRetrieve**](CloudAPI.md#CloudServersActivityRetrieve) | **Get** /api/cloud/servers/{id}/activity/ | 
[**CloudServersAttachIpv4Create**](CloudAPI.md#CloudServersAttachIpv4Create) | **Post** /api/cloud/servers/{id}/attach-ipv4/ | 
[**CloudServersAttachIpv6Create**](CloudAPI.md#CloudServersAttachIpv6Create) | **Post** /api/cloud/servers/{id}/attach-ipv6/ | 
[**CloudServersBootIsosList**](CloudAPI.md#CloudServersBootIsosList) | **Get** /api/cloud/servers/{id}/boot-isos/ | 
[**CloudServersConsoleCreate**](CloudAPI.md#CloudServersConsoleCreate) | **Post** /api/cloud/servers/{id}/console/ | 
[**CloudServersCreate**](CloudAPI.md#CloudServersCreate) | **Post** /api/cloud/servers/ | 
[**CloudServersDestroy**](CloudAPI.md#CloudServersDestroy) | **Delete** /api/cloud/servers/{id}/ | 
[**CloudServersDestroyProtectionCreate**](CloudAPI.md#CloudServersDestroyProtectionCreate) | **Post** /api/cloud/servers/{id}/destroy-protection/ | 
[**CloudServersDetachIpv4Create**](CloudAPI.md#CloudServersDetachIpv4Create) | **Post** /api/cloud/servers/{id}/detach-ipv4/ | 
[**CloudServersDetachIpv6Create**](CloudAPI.md#CloudServersDetachIpv6Create) | **Post** /api/cloud/servers/{id}/detach-ipv6/ | 
[**CloudServersList**](CloudAPI.md#CloudServersList) | **Get** /api/cloud/servers/ | 
[**CloudServersModifyPackageCreate**](CloudAPI.md#CloudServersModifyPackageCreate) | **Post** /api/cloud/servers/{id}/modify-package/ | 
[**CloudServersMonitoringRetrieve**](CloudAPI.md#CloudServersMonitoringRetrieve) | **Get** /api/cloud/servers/{id}/monitoring/ | 
[**CloudServersMonitoringUpdate**](CloudAPI.md#CloudServersMonitoringUpdate) | **Put** /api/cloud/servers/{id}/monitoring/ | 
[**CloudServersPartialUpdate**](CloudAPI.md#CloudServersPartialUpdate) | **Patch** /api/cloud/servers/{id}/ | 
[**CloudServersPowerManagementCreate**](CloudAPI.md#CloudServersPowerManagementCreate) | **Post** /api/cloud/servers/{id}/power-management/ | 
[**CloudServersPowerManagementRetrieve**](CloudAPI.md#CloudServersPowerManagementRetrieve) | **Get** /api/cloud/servers/{id}/power-management/ | 
[**CloudServersPublicInterfaceCreate**](CloudAPI.md#CloudServersPublicInterfaceCreate) | **Post** /api/cloud/servers/{id}/public-interface/ | 
[**CloudServersPublicInterfaceDestroy**](CloudAPI.md#CloudServersPublicInterfaceDestroy) | **Delete** /api/cloud/servers/{id}/public-interface/ | 
[**CloudServersPublicInterfaceRetrieve**](CloudAPI.md#CloudServersPublicInterfaceRetrieve) | **Get** /api/cloud/servers/{id}/public-interface/ | 
[**CloudServersRescueEnterCreate**](CloudAPI.md#CloudServersRescueEnterCreate) | **Post** /api/cloud/servers/{id}/rescue/enter/ | 
[**CloudServersRescueExitCreate**](CloudAPI.md#CloudServersRescueExitCreate) | **Post** /api/cloud/servers/{id}/rescue/exit/ | 
[**CloudServersRetrieve**](CloudAPI.md#CloudServersRetrieve) | **Get** /api/cloud/servers/{id}/ | 
[**CloudServersRetryProvisionCreate**](CloudAPI.md#CloudServersRetryProvisionCreate) | **Post** /api/cloud/servers/{id}/retry-provision/ | 
[**CloudServersSnapshotsCreate**](CloudAPI.md#CloudServersSnapshotsCreate) | **Post** /api/cloud/servers/{id}/snapshots/ | 
[**CloudServersSnapshotsDestroy**](CloudAPI.md#CloudServersSnapshotsDestroy) | **Delete** /api/cloud/servers/{id}/snapshots/{snapshot_name}/ | 
[**CloudServersSnapshotsList**](CloudAPI.md#CloudServersSnapshotsList) | **Get** /api/cloud/servers/{id}/snapshots/ | 
[**CloudServersSnapshotsRollbackCreate**](CloudAPI.md#CloudServersSnapshotsRollbackCreate) | **Post** /api/cloud/servers/{id}/snapshots/{snapshot_name}/rollback/ | 
[**CloudServersTrafficRetrieve**](CloudAPI.md#CloudServersTrafficRetrieve) | **Get** /api/cloud/servers/{id}/traffic/ | 
[**CloudServersUpdate**](CloudAPI.md#CloudServersUpdate) | **Put** /api/cloud/servers/{id}/ | 
[**CloudServersUsageRetrieve**](CloudAPI.md#CloudServersUsageRetrieve) | **Get** /api/cloud/servers/{id}/usage/ | 
[**CloudServersVolumesCreate**](CloudAPI.md#CloudServersVolumesCreate) | **Post** /api/cloud/servers/{server_id}/volumes/ | 
[**CloudServersVolumesDestroy**](CloudAPI.md#CloudServersVolumesDestroy) | **Delete** /api/cloud/servers/{server_id}/volumes/{volume_id}/ | 
[**CloudServersVolumesList**](CloudAPI.md#CloudServersVolumesList) | **Get** /api/cloud/servers/{server_id}/volumes/ | 
[**CloudServersVolumesPartialUpdate**](CloudAPI.md#CloudServersVolumesPartialUpdate) | **Patch** /api/cloud/servers/{server_id}/volumes/{volume_id}/ | 
[**CloudServersVolumesRetrieve**](CloudAPI.md#CloudServersVolumesRetrieve) | **Get** /api/cloud/servers/{server_id}/volumes/{volume_id}/ | 
[**CloudServersVolumesUpdate**](CloudAPI.md#CloudServersVolumesUpdate) | **Put** /api/cloud/servers/{server_id}/volumes/{volume_id}/ | 
[**CloudStorageProductsList**](CloudAPI.md#CloudStorageProductsList) | **Get** /api/cloud/storage-products/ | 
[**CloudStorageProductsRetrieve**](CloudAPI.md#CloudStorageProductsRetrieve) | **Get** /api/cloud/storage-products/{id}/ | 
[**CloudVolumesAttachCreate**](CloudAPI.md#CloudVolumesAttachCreate) | **Post** /api/cloud/volumes/{id}/attach/ | 
[**CloudVolumesDestroy**](CloudAPI.md#CloudVolumesDestroy) | **Delete** /api/cloud/volumes/{id}/ | 
[**CloudVolumesDetachCreate**](CloudAPI.md#CloudVolumesDetachCreate) | **Post** /api/cloud/volumes/{id}/detach/ | 
[**CloudVolumesList**](CloudAPI.md#CloudVolumesList) | **Get** /api/cloud/volumes/ | 
[**CloudVolumesPartialUpdate**](CloudAPI.md#CloudVolumesPartialUpdate) | **Patch** /api/cloud/volumes/{id}/ | 
[**CloudVolumesRetrieve**](CloudAPI.md#CloudVolumesRetrieve) | **Get** /api/cloud/volumes/{id}/ | 
[**CloudVolumesUpdate**](CloudAPI.md#CloudVolumesUpdate) | **Put** /api/cloud/volumes/{id}/ | 



## CloudBucketsCreate

> Bucket CloudBucketsCreate(ctx).BucketCreateRequest(bucketCreateRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	bucketCreateRequest := *openapiclient.NewBucketCreateRequest("Name_example", int32(123)) // BucketCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudBucketsCreate(context.Background()).BucketCreateRequest(bucketCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudBucketsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudBucketsCreate`: Bucket
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudBucketsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudBucketsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bucketCreateRequest** | [**BucketCreateRequest**](BucketCreateRequest.md) |  | 

### Return type

[**Bucket**](Bucket.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudBucketsCredentialsRevealCreate

> BucketCredentials CloudBucketsCredentialsRevealCreate(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this S3 bucket.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudBucketsCredentialsRevealCreate(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudBucketsCredentialsRevealCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudBucketsCredentialsRevealCreate`: BucketCredentials
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudBucketsCredentialsRevealCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this S3 bucket. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudBucketsCredentialsRevealCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BucketCredentials**](BucketCredentials.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudBucketsCredentialsRotateCreate

> BucketCredentials CloudBucketsCredentialsRotateCreate(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this S3 bucket.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudBucketsCredentialsRotateCreate(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudBucketsCredentialsRotateCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudBucketsCredentialsRotateCreate`: BucketCredentials
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudBucketsCredentialsRotateCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this S3 bucket. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudBucketsCredentialsRotateCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BucketCredentials**](BucketCredentials.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudBucketsDestroy

> BucketCancelResponse CloudBucketsDestroy(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this S3 bucket.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudBucketsDestroy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudBucketsDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudBucketsDestroy`: BucketCancelResponse
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudBucketsDestroy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this S3 bucket. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudBucketsDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BucketCancelResponse**](BucketCancelResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudBucketsDestroyProtectionCreate

> DestroyProtection CloudBucketsDestroyProtectionCreate(ctx, id).DestroyProtectionRequest(destroyProtectionRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this S3 bucket.
	destroyProtectionRequest := *openapiclient.NewDestroyProtectionRequest(false) // DestroyProtectionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudBucketsDestroyProtectionCreate(context.Background(), id).DestroyProtectionRequest(destroyProtectionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudBucketsDestroyProtectionCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudBucketsDestroyProtectionCreate`: DestroyProtection
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudBucketsDestroyProtectionCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this S3 bucket. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudBucketsDestroyProtectionCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **destroyProtectionRequest** | [**DestroyProtectionRequest**](DestroyProtectionRequest.md) |  | 

### Return type

[**DestroyProtection**](DestroyProtection.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudBucketsList

> []Bucket CloudBucketsList(ctx).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudBucketsList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudBucketsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudBucketsList`: []Bucket
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudBucketsList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudBucketsListRequest struct via the builder pattern


### Return type

[**[]Bucket**](Bucket.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudBucketsResizeCreate

> Bucket CloudBucketsResizeCreate(ctx, id).BucketResizeRequest(bucketResizeRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this S3 bucket.
	bucketResizeRequest := *openapiclient.NewBucketResizeRequest(int32(123)) // BucketResizeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudBucketsResizeCreate(context.Background(), id).BucketResizeRequest(bucketResizeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudBucketsResizeCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudBucketsResizeCreate`: Bucket
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudBucketsResizeCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this S3 bucket. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudBucketsResizeCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bucketResizeRequest** | [**BucketResizeRequest**](BucketResizeRequest.md) |  | 

### Return type

[**Bucket**](Bucket.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudBucketsRetrieve

> Bucket CloudBucketsRetrieve(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this S3 bucket.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudBucketsRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudBucketsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudBucketsRetrieve`: Bucket
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudBucketsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this S3 bucket. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudBucketsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Bucket**](Bucket.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudBucketsVisibilityCreate

> Bucket CloudBucketsVisibilityCreate(ctx, id).BucketVisibilityRequest(bucketVisibilityRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this S3 bucket.
	bucketVisibilityRequest := *openapiclient.NewBucketVisibilityRequest(false) // BucketVisibilityRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudBucketsVisibilityCreate(context.Background(), id).BucketVisibilityRequest(bucketVisibilityRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudBucketsVisibilityCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudBucketsVisibilityCreate`: Bucket
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudBucketsVisibilityCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this S3 bucket. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudBucketsVisibilityCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bucketVisibilityRequest** | [**BucketVisibilityRequest**](BucketVisibilityRequest.md) |  | 

### Return type

[**Bucket**](Bucket.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudFirewallRulesSetCreate

> FirewallRulesSet CloudFirewallRulesSetCreate(ctx).FirewallRulesSetRequest(firewallRulesSetRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	firewallRulesSetRequest := *openapiclient.NewFirewallRulesSetRequest("Name_example") // FirewallRulesSetRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudFirewallRulesSetCreate(context.Background()).FirewallRulesSetRequest(firewallRulesSetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudFirewallRulesSetCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudFirewallRulesSetCreate`: FirewallRulesSet
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudFirewallRulesSetCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudFirewallRulesSetCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **firewallRulesSetRequest** | [**FirewallRulesSetRequest**](FirewallRulesSetRequest.md) |  | 

### Return type

[**FirewallRulesSet**](FirewallRulesSet.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudFirewallRulesSetDestroy

> CloudFirewallRulesSetDestroy(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this firewall rules set.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CloudAPI.CloudFirewallRulesSetDestroy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudFirewallRulesSetDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this firewall rules set. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudFirewallRulesSetDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudFirewallRulesSetList

> []FirewallRulesSet CloudFirewallRulesSetList(ctx).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudFirewallRulesSetList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudFirewallRulesSetList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudFirewallRulesSetList`: []FirewallRulesSet
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudFirewallRulesSetList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudFirewallRulesSetListRequest struct via the builder pattern


### Return type

[**[]FirewallRulesSet**](FirewallRulesSet.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudFirewallRulesSetPartialUpdate

> FirewallRulesSet CloudFirewallRulesSetPartialUpdate(ctx, id).PatchedFirewallRulesSetRequest(patchedFirewallRulesSetRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this firewall rules set.
	patchedFirewallRulesSetRequest := *openapiclient.NewPatchedFirewallRulesSetRequest() // PatchedFirewallRulesSetRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudFirewallRulesSetPartialUpdate(context.Background(), id).PatchedFirewallRulesSetRequest(patchedFirewallRulesSetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudFirewallRulesSetPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudFirewallRulesSetPartialUpdate`: FirewallRulesSet
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudFirewallRulesSetPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this firewall rules set. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudFirewallRulesSetPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedFirewallRulesSetRequest** | [**PatchedFirewallRulesSetRequest**](PatchedFirewallRulesSetRequest.md) |  | 

### Return type

[**FirewallRulesSet**](FirewallRulesSet.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudFirewallRulesSetRetrieve

> FirewallRulesSet CloudFirewallRulesSetRetrieve(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this firewall rules set.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudFirewallRulesSetRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudFirewallRulesSetRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudFirewallRulesSetRetrieve`: FirewallRulesSet
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudFirewallRulesSetRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this firewall rules set. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudFirewallRulesSetRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FirewallRulesSet**](FirewallRulesSet.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudFirewallRulesSetRulesCreate

> FirewallRule CloudFirewallRulesSetRulesCreate(ctx, rulesSetId).FirewallRuleRequest(firewallRuleRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	rulesSetId := "rulesSetId_example" // string | 
	firewallRuleRequest := *openapiclient.NewFirewallRuleRequest(openapiclient.FirewallRuleDirectionEnum("in"), openapiclient.FwPolicyOutEnum("ACCEPT")) // FirewallRuleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudFirewallRulesSetRulesCreate(context.Background(), rulesSetId).FirewallRuleRequest(firewallRuleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudFirewallRulesSetRulesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudFirewallRulesSetRulesCreate`: FirewallRule
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudFirewallRulesSetRulesCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rulesSetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudFirewallRulesSetRulesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **firewallRuleRequest** | [**FirewallRuleRequest**](FirewallRuleRequest.md) |  | 

### Return type

[**FirewallRule**](FirewallRule.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudFirewallRulesSetRulesDestroy

> CloudFirewallRulesSetRulesDestroy(ctx, ruleId, rulesSetId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	ruleId := "ruleId_example" // string | 
	rulesSetId := "rulesSetId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CloudAPI.CloudFirewallRulesSetRulesDestroy(context.Background(), ruleId, rulesSetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudFirewallRulesSetRulesDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleId** | **string** |  | 
**rulesSetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudFirewallRulesSetRulesDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudFirewallRulesSetRulesList

> []FirewallRule CloudFirewallRulesSetRulesList(ctx, rulesSetId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	rulesSetId := "rulesSetId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudFirewallRulesSetRulesList(context.Background(), rulesSetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudFirewallRulesSetRulesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudFirewallRulesSetRulesList`: []FirewallRule
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudFirewallRulesSetRulesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rulesSetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudFirewallRulesSetRulesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]FirewallRule**](FirewallRule.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudFirewallRulesSetRulesPartialUpdate

> FirewallRule CloudFirewallRulesSetRulesPartialUpdate(ctx, ruleId, rulesSetId).PatchedFirewallRuleRequest(patchedFirewallRuleRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	ruleId := "ruleId_example" // string | 
	rulesSetId := "rulesSetId_example" // string | 
	patchedFirewallRuleRequest := *openapiclient.NewPatchedFirewallRuleRequest() // PatchedFirewallRuleRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudFirewallRulesSetRulesPartialUpdate(context.Background(), ruleId, rulesSetId).PatchedFirewallRuleRequest(patchedFirewallRuleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudFirewallRulesSetRulesPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudFirewallRulesSetRulesPartialUpdate`: FirewallRule
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudFirewallRulesSetRulesPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleId** | **string** |  | 
**rulesSetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudFirewallRulesSetRulesPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchedFirewallRuleRequest** | [**PatchedFirewallRuleRequest**](PatchedFirewallRuleRequest.md) |  | 

### Return type

[**FirewallRule**](FirewallRule.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudFirewallRulesSetRulesRetrieve

> FirewallRule CloudFirewallRulesSetRulesRetrieve(ctx, ruleId, rulesSetId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	ruleId := "ruleId_example" // string | 
	rulesSetId := "rulesSetId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudFirewallRulesSetRulesRetrieve(context.Background(), ruleId, rulesSetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudFirewallRulesSetRulesRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudFirewallRulesSetRulesRetrieve`: FirewallRule
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudFirewallRulesSetRulesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleId** | **string** |  | 
**rulesSetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudFirewallRulesSetRulesRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**FirewallRule**](FirewallRule.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudFirewallRulesSetRulesUpdate

> FirewallRule CloudFirewallRulesSetRulesUpdate(ctx, ruleId, rulesSetId).FirewallRuleRequest(firewallRuleRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	ruleId := "ruleId_example" // string | 
	rulesSetId := "rulesSetId_example" // string | 
	firewallRuleRequest := *openapiclient.NewFirewallRuleRequest(openapiclient.FirewallRuleDirectionEnum("in"), openapiclient.FwPolicyOutEnum("ACCEPT")) // FirewallRuleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudFirewallRulesSetRulesUpdate(context.Background(), ruleId, rulesSetId).FirewallRuleRequest(firewallRuleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudFirewallRulesSetRulesUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudFirewallRulesSetRulesUpdate`: FirewallRule
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudFirewallRulesSetRulesUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleId** | **string** |  | 
**rulesSetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudFirewallRulesSetRulesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **firewallRuleRequest** | [**FirewallRuleRequest**](FirewallRuleRequest.md) |  | 

### Return type

[**FirewallRule**](FirewallRule.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudFirewallRulesSetUpdate

> FirewallRulesSet CloudFirewallRulesSetUpdate(ctx, id).FirewallRulesSetRequest(firewallRulesSetRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this firewall rules set.
	firewallRulesSetRequest := *openapiclient.NewFirewallRulesSetRequest("Name_example") // FirewallRulesSetRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudFirewallRulesSetUpdate(context.Background(), id).FirewallRulesSetRequest(firewallRulesSetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudFirewallRulesSetUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudFirewallRulesSetUpdate`: FirewallRulesSet
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudFirewallRulesSetUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this firewall rules set. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudFirewallRulesSetUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **firewallRulesSetRequest** | [**FirewallRulesSetRequest**](FirewallRulesSetRequest.md) |  | 

### Return type

[**FirewallRulesSet**](FirewallRulesSet.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudFloatingIpv4AuthorizationsList

> PaginatedFloatingIPAuthorizationList CloudFloatingIpv4AuthorizationsList(ctx, id).Page(page).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this floating IPv4.
	page := int32(56) // int32 | A page number within the paginated result set. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudFloatingIpv4AuthorizationsList(context.Background(), id).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudFloatingIpv4AuthorizationsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudFloatingIpv4AuthorizationsList`: PaginatedFloatingIPAuthorizationList
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudFloatingIpv4AuthorizationsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this floating IPv4. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudFloatingIpv4AuthorizationsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | A page number within the paginated result set. | 

### Return type

[**PaginatedFloatingIPAuthorizationList**](PaginatedFloatingIPAuthorizationList.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudFloatingIpv4AuthorizeCreate

> FloatingIPv4AuthorizeResponse CloudFloatingIpv4AuthorizeCreate(ctx, id).FloatingIPAuthorizeRequest(floatingIPAuthorizeRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this floating IPv4.
	floatingIPAuthorizeRequest := *openapiclient.NewFloatingIPAuthorizeRequest(int32(123)) // FloatingIPAuthorizeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudFloatingIpv4AuthorizeCreate(context.Background(), id).FloatingIPAuthorizeRequest(floatingIPAuthorizeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudFloatingIpv4AuthorizeCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudFloatingIpv4AuthorizeCreate`: FloatingIPv4AuthorizeResponse
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudFloatingIpv4AuthorizeCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this floating IPv4. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudFloatingIpv4AuthorizeCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **floatingIPAuthorizeRequest** | [**FloatingIPAuthorizeRequest**](FloatingIPAuthorizeRequest.md) |  | 

### Return type

[**FloatingIPv4AuthorizeResponse**](FloatingIPv4AuthorizeResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudFloatingIpv4Create

> FloatingIPv4 CloudFloatingIpv4Create(ctx).FloatingIPv4CreateRequest(floatingIPv4CreateRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	floatingIPv4CreateRequest := *openapiclient.NewFloatingIPv4CreateRequest() // FloatingIPv4CreateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudFloatingIpv4Create(context.Background()).FloatingIPv4CreateRequest(floatingIPv4CreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudFloatingIpv4Create``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudFloatingIpv4Create`: FloatingIPv4
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudFloatingIpv4Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudFloatingIpv4CreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **floatingIPv4CreateRequest** | [**FloatingIPv4CreateRequest**](FloatingIPv4CreateRequest.md) |  | 

### Return type

[**FloatingIPv4**](FloatingIPv4.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudFloatingIpv4Destroy

> CloudFloatingIpv4Destroy(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this floating IPv4.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CloudAPI.CloudFloatingIpv4Destroy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudFloatingIpv4Destroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this floating IPv4. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudFloatingIpv4DestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudFloatingIpv4List

> PaginatedFloatingIPv4List CloudFloatingIpv4List(ctx).Page(page).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	page := int32(56) // int32 | A page number within the paginated result set. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudFloatingIpv4List(context.Background()).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudFloatingIpv4List``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudFloatingIpv4List`: PaginatedFloatingIPv4List
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudFloatingIpv4List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudFloatingIpv4ListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | A page number within the paginated result set. | 

### Return type

[**PaginatedFloatingIPv4List**](PaginatedFloatingIPv4List.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudFloatingIpv4RdnsCreate

> ReverseDNS CloudFloatingIpv4RdnsCreate(ctx, id).ReverseDNSRequest(reverseDNSRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this floating IPv4.
	reverseDNSRequest := *openapiclient.NewReverseDNSRequest("ReverseDns_example") // ReverseDNSRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudFloatingIpv4RdnsCreate(context.Background(), id).ReverseDNSRequest(reverseDNSRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudFloatingIpv4RdnsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudFloatingIpv4RdnsCreate`: ReverseDNS
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudFloatingIpv4RdnsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this floating IPv4. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudFloatingIpv4RdnsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reverseDNSRequest** | [**ReverseDNSRequest**](ReverseDNSRequest.md) |  | 

### Return type

[**ReverseDNS**](ReverseDNS.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudFloatingIpv4RdnsRetrieve

> ReverseDNS CloudFloatingIpv4RdnsRetrieve(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this floating IPv4.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudFloatingIpv4RdnsRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudFloatingIpv4RdnsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudFloatingIpv4RdnsRetrieve`: ReverseDNS
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudFloatingIpv4RdnsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this floating IPv4. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudFloatingIpv4RdnsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ReverseDNS**](ReverseDNS.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudFloatingIpv4Retrieve

> FloatingIPv4 CloudFloatingIpv4Retrieve(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this floating IPv4.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudFloatingIpv4Retrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudFloatingIpv4Retrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudFloatingIpv4Retrieve`: FloatingIPv4
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudFloatingIpv4Retrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this floating IPv4. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudFloatingIpv4RetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FloatingIPv4**](FloatingIPv4.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudFloatingIpv4UnauthorizeCreate

> FloatingIPv4UnauthorizeResponse CloudFloatingIpv4UnauthorizeCreate(ctx, id).FloatingIPAuthorizeRequest(floatingIPAuthorizeRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this floating IPv4.
	floatingIPAuthorizeRequest := *openapiclient.NewFloatingIPAuthorizeRequest(int32(123)) // FloatingIPAuthorizeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudFloatingIpv4UnauthorizeCreate(context.Background(), id).FloatingIPAuthorizeRequest(floatingIPAuthorizeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudFloatingIpv4UnauthorizeCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudFloatingIpv4UnauthorizeCreate`: FloatingIPv4UnauthorizeResponse
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudFloatingIpv4UnauthorizeCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this floating IPv4. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudFloatingIpv4UnauthorizeCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **floatingIPAuthorizeRequest** | [**FloatingIPAuthorizeRequest**](FloatingIPAuthorizeRequest.md) |  | 

### Return type

[**FloatingIPv4UnauthorizeResponse**](FloatingIPv4UnauthorizeResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudFloatingIpv6AuthorizationsList

> PaginatedFloatingIPAuthorizationList CloudFloatingIpv6AuthorizationsList(ctx, id).Page(page).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this floating IPv6.
	page := int32(56) // int32 | A page number within the paginated result set. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudFloatingIpv6AuthorizationsList(context.Background(), id).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudFloatingIpv6AuthorizationsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudFloatingIpv6AuthorizationsList`: PaginatedFloatingIPAuthorizationList
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudFloatingIpv6AuthorizationsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this floating IPv6. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudFloatingIpv6AuthorizationsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | A page number within the paginated result set. | 

### Return type

[**PaginatedFloatingIPAuthorizationList**](PaginatedFloatingIPAuthorizationList.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudFloatingIpv6AuthorizeCreate

> FloatingIPv6AuthorizeResponse CloudFloatingIpv6AuthorizeCreate(ctx, id).FloatingIPAuthorizeRequest(floatingIPAuthorizeRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this floating IPv6.
	floatingIPAuthorizeRequest := *openapiclient.NewFloatingIPAuthorizeRequest(int32(123)) // FloatingIPAuthorizeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudFloatingIpv6AuthorizeCreate(context.Background(), id).FloatingIPAuthorizeRequest(floatingIPAuthorizeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudFloatingIpv6AuthorizeCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudFloatingIpv6AuthorizeCreate`: FloatingIPv6AuthorizeResponse
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudFloatingIpv6AuthorizeCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this floating IPv6. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudFloatingIpv6AuthorizeCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **floatingIPAuthorizeRequest** | [**FloatingIPAuthorizeRequest**](FloatingIPAuthorizeRequest.md) |  | 

### Return type

[**FloatingIPv6AuthorizeResponse**](FloatingIPv6AuthorizeResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudFloatingIpv6Create

> FloatingIPv6 CloudFloatingIpv6Create(ctx).FloatingIPv6CreateRequest(floatingIPv6CreateRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	floatingIPv6CreateRequest := *openapiclient.NewFloatingIPv6CreateRequest() // FloatingIPv6CreateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudFloatingIpv6Create(context.Background()).FloatingIPv6CreateRequest(floatingIPv6CreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudFloatingIpv6Create``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudFloatingIpv6Create`: FloatingIPv6
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudFloatingIpv6Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudFloatingIpv6CreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **floatingIPv6CreateRequest** | [**FloatingIPv6CreateRequest**](FloatingIPv6CreateRequest.md) |  | 

### Return type

[**FloatingIPv6**](FloatingIPv6.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudFloatingIpv6Destroy

> CloudFloatingIpv6Destroy(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this floating IPv6.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CloudAPI.CloudFloatingIpv6Destroy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudFloatingIpv6Destroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this floating IPv6. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudFloatingIpv6DestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudFloatingIpv6List

> PaginatedFloatingIPv6List CloudFloatingIpv6List(ctx).Page(page).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	page := int32(56) // int32 | A page number within the paginated result set. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudFloatingIpv6List(context.Background()).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudFloatingIpv6List``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudFloatingIpv6List`: PaginatedFloatingIPv6List
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudFloatingIpv6List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudFloatingIpv6ListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | A page number within the paginated result set. | 

### Return type

[**PaginatedFloatingIPv6List**](PaginatedFloatingIPv6List.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudFloatingIpv6RdnsCreate

> ReverseDNS CloudFloatingIpv6RdnsCreate(ctx, id).ReverseDNSRequest(reverseDNSRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this floating IPv6.
	reverseDNSRequest := *openapiclient.NewReverseDNSRequest("ReverseDns_example") // ReverseDNSRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudFloatingIpv6RdnsCreate(context.Background(), id).ReverseDNSRequest(reverseDNSRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudFloatingIpv6RdnsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudFloatingIpv6RdnsCreate`: ReverseDNS
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudFloatingIpv6RdnsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this floating IPv6. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudFloatingIpv6RdnsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reverseDNSRequest** | [**ReverseDNSRequest**](ReverseDNSRequest.md) |  | 

### Return type

[**ReverseDNS**](ReverseDNS.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudFloatingIpv6RdnsRetrieve

> ReverseDNS CloudFloatingIpv6RdnsRetrieve(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this floating IPv6.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudFloatingIpv6RdnsRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudFloatingIpv6RdnsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudFloatingIpv6RdnsRetrieve`: ReverseDNS
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudFloatingIpv6RdnsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this floating IPv6. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudFloatingIpv6RdnsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ReverseDNS**](ReverseDNS.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudFloatingIpv6Retrieve

> FloatingIPv6 CloudFloatingIpv6Retrieve(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this floating IPv6.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudFloatingIpv6Retrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudFloatingIpv6Retrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudFloatingIpv6Retrieve`: FloatingIPv6
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudFloatingIpv6Retrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this floating IPv6. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudFloatingIpv6RetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FloatingIPv6**](FloatingIPv6.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudFloatingIpv6UnauthorizeCreate

> FloatingIPv6UnauthorizeResponse CloudFloatingIpv6UnauthorizeCreate(ctx, id).FloatingIPAuthorizeRequest(floatingIPAuthorizeRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this floating IPv6.
	floatingIPAuthorizeRequest := *openapiclient.NewFloatingIPAuthorizeRequest(int32(123)) // FloatingIPAuthorizeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudFloatingIpv6UnauthorizeCreate(context.Background(), id).FloatingIPAuthorizeRequest(floatingIPAuthorizeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudFloatingIpv6UnauthorizeCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudFloatingIpv6UnauthorizeCreate`: FloatingIPv6UnauthorizeResponse
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudFloatingIpv6UnauthorizeCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this floating IPv6. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudFloatingIpv6UnauthorizeCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **floatingIPAuthorizeRequest** | [**FloatingIPAuthorizeRequest**](FloatingIPAuthorizeRequest.md) |  | 

### Return type

[**FloatingIPv6UnauthorizeResponse**](FloatingIPv6UnauthorizeResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGatewaysCreate

> Gateway CloudGatewaysCreate(ctx).GatewayCreateRequest(gatewayCreateRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	gatewayCreateRequest := *openapiclient.NewGatewayCreateRequest("Name_example") // GatewayCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudGatewaysCreate(context.Background()).GatewayCreateRequest(gatewayCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudGatewaysCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGatewaysCreate`: Gateway
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudGatewaysCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudGatewaysCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gatewayCreateRequest** | [**GatewayCreateRequest**](GatewayCreateRequest.md) |  | 

### Return type

[**Gateway**](Gateway.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGatewaysDestroy

> GatewayOperation CloudGatewaysDestroy(ctx, id).Networks(networks).Revision(revision).Servers(servers).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := "id_example" // string | 
	networks := []int32{int32(123)} // []int32 | The preview's network ids. (optional)
	revision := int32(56) // int32 | The gateway's desired revision from the preview; without it, a preview. (optional)
	servers := []int32{int32(123)} // []int32 | The preview's server ids. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudGatewaysDestroy(context.Background(), id).Networks(networks).Revision(revision).Servers(servers).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudGatewaysDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGatewaysDestroy`: GatewayOperation
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudGatewaysDestroy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGatewaysDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networks** | **[]int32** | The preview&#39;s network ids. | 
 **revision** | **int32** | The gateway&#39;s desired revision from the preview; without it, a preview. | 
 **servers** | **[]int32** | The preview&#39;s server ids. | 

### Return type

[**GatewayOperation**](GatewayOperation.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGatewaysFirewallRulesCreate

> GatewayFirewallRule CloudGatewaysFirewallRulesCreate(ctx, gatewayId).GatewayFirewallRuleWriteRequest(gatewayFirewallRuleWriteRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	gatewayId := int32(56) // int32 | 
	gatewayFirewallRuleWriteRequest := *openapiclient.NewGatewayFirewallRuleWriteRequest(openapiclient.FwPolicyOutEnum("ACCEPT"), *openapiclient.NewGatewayFirewallSourceRequest(), *openapiclient.NewGatewayFirewallDestinationRequest(), openapiclient.GatewayFirewallProtocolEnum("tcp")) // GatewayFirewallRuleWriteRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudGatewaysFirewallRulesCreate(context.Background(), gatewayId).GatewayFirewallRuleWriteRequest(gatewayFirewallRuleWriteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudGatewaysFirewallRulesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGatewaysFirewallRulesCreate`: GatewayFirewallRule
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudGatewaysFirewallRulesCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gatewayId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGatewaysFirewallRulesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gatewayFirewallRuleWriteRequest** | [**GatewayFirewallRuleWriteRequest**](GatewayFirewallRuleWriteRequest.md) |  | 

### Return type

[**GatewayFirewallRule**](GatewayFirewallRule.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGatewaysFirewallRulesDestroy

> GatewayOperation CloudGatewaysFirewallRulesDestroy(ctx, gatewayId, id).Revision(revision).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	gatewayId := int32(56) // int32 | 
	id := int32(56) // int32 | 
	revision := int32(56) // int32 | The gateway's desired revision from the preview; without it, a preview. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudGatewaysFirewallRulesDestroy(context.Background(), gatewayId, id).Revision(revision).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudGatewaysFirewallRulesDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGatewaysFirewallRulesDestroy`: GatewayOperation
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudGatewaysFirewallRulesDestroy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gatewayId** | **int32** |  | 
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGatewaysFirewallRulesDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **revision** | **int32** | The gateway&#39;s desired revision from the preview; without it, a preview. | 

### Return type

[**GatewayOperation**](GatewayOperation.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGatewaysFirewallRulesList

> PaginatedGatewayFirewallRuleList CloudGatewaysFirewallRulesList(ctx, gatewayId).Page(page).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	gatewayId := int32(56) // int32 | 
	page := int32(56) // int32 | A page number within the paginated result set. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudGatewaysFirewallRulesList(context.Background(), gatewayId).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudGatewaysFirewallRulesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGatewaysFirewallRulesList`: PaginatedGatewayFirewallRuleList
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudGatewaysFirewallRulesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gatewayId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGatewaysFirewallRulesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | A page number within the paginated result set. | 

### Return type

[**PaginatedGatewayFirewallRuleList**](PaginatedGatewayFirewallRuleList.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGatewaysFirewallRulesMoveCreate

> GatewayFirewallRule CloudGatewaysFirewallRulesMoveCreate(ctx, gatewayId, id).GatewayFirewallMoveRequest(gatewayFirewallMoveRequest).Revision(revision).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	gatewayId := int32(56) // int32 | 
	id := int32(56) // int32 | 
	gatewayFirewallMoveRequest := *openapiclient.NewGatewayFirewallMoveRequest(int32(123)) // GatewayFirewallMoveRequest | 
	revision := int32(56) // int32 | The gateway's desired revision from the preview; without it, a preview. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudGatewaysFirewallRulesMoveCreate(context.Background(), gatewayId, id).GatewayFirewallMoveRequest(gatewayFirewallMoveRequest).Revision(revision).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudGatewaysFirewallRulesMoveCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGatewaysFirewallRulesMoveCreate`: GatewayFirewallRule
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudGatewaysFirewallRulesMoveCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gatewayId** | **int32** |  | 
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGatewaysFirewallRulesMoveCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **gatewayFirewallMoveRequest** | [**GatewayFirewallMoveRequest**](GatewayFirewallMoveRequest.md) |  | 
 **revision** | **int32** | The gateway&#39;s desired revision from the preview; without it, a preview. | 

### Return type

[**GatewayFirewallRule**](GatewayFirewallRule.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGatewaysFirewallRulesPartialUpdate

> GatewayFirewallRule CloudGatewaysFirewallRulesPartialUpdate(ctx, gatewayId, id).PatchedGatewayFirewallRuleWriteRequest(patchedGatewayFirewallRuleWriteRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	gatewayId := int32(56) // int32 | 
	id := int32(56) // int32 | 
	patchedGatewayFirewallRuleWriteRequest := *openapiclient.NewPatchedGatewayFirewallRuleWriteRequest() // PatchedGatewayFirewallRuleWriteRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudGatewaysFirewallRulesPartialUpdate(context.Background(), gatewayId, id).PatchedGatewayFirewallRuleWriteRequest(patchedGatewayFirewallRuleWriteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudGatewaysFirewallRulesPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGatewaysFirewallRulesPartialUpdate`: GatewayFirewallRule
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudGatewaysFirewallRulesPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gatewayId** | **int32** |  | 
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGatewaysFirewallRulesPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchedGatewayFirewallRuleWriteRequest** | [**PatchedGatewayFirewallRuleWriteRequest**](PatchedGatewayFirewallRuleWriteRequest.md) |  | 

### Return type

[**GatewayFirewallRule**](GatewayFirewallRule.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGatewaysFirewallRulesRetrieve

> GatewayFirewallRule CloudGatewaysFirewallRulesRetrieve(ctx, gatewayId, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	gatewayId := int32(56) // int32 | 
	id := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudGatewaysFirewallRulesRetrieve(context.Background(), gatewayId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudGatewaysFirewallRulesRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGatewaysFirewallRulesRetrieve`: GatewayFirewallRule
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudGatewaysFirewallRulesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gatewayId** | **int32** |  | 
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGatewaysFirewallRulesRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GatewayFirewallRule**](GatewayFirewallRule.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGatewaysList

> PaginatedGatewayList CloudGatewaysList(ctx).Page(page).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	page := int32(56) // int32 | A page number within the paginated result set. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudGatewaysList(context.Background()).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudGatewaysList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGatewaysList`: PaginatedGatewayList
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudGatewaysList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudGatewaysListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | A page number within the paginated result set. | 

### Return type

[**PaginatedGatewayList**](PaginatedGatewayList.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGatewaysListenersCreate

> GatewayLBListener CloudGatewaysListenersCreate(ctx, gatewayId).GatewayLBListenerWriteRequest(gatewayLBListenerWriteRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	gatewayId := int32(56) // int32 | 
	gatewayLBListenerWriteRequest := *openapiclient.NewGatewayLBListenerWriteRequest() // GatewayLBListenerWriteRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudGatewaysListenersCreate(context.Background(), gatewayId).GatewayLBListenerWriteRequest(gatewayLBListenerWriteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudGatewaysListenersCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGatewaysListenersCreate`: GatewayLBListener
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudGatewaysListenersCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gatewayId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGatewaysListenersCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gatewayLBListenerWriteRequest** | [**GatewayLBListenerWriteRequest**](GatewayLBListenerWriteRequest.md) |  | 

### Return type

[**GatewayLBListener**](GatewayLBListener.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGatewaysListenersDestroy

> GatewayOperation CloudGatewaysListenersDestroy(ctx, gatewayId, id).Revision(revision).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	gatewayId := int32(56) // int32 | 
	id := int32(56) // int32 | 
	revision := int32(56) // int32 | The gateway's desired revision from the preview; without it, a preview. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudGatewaysListenersDestroy(context.Background(), gatewayId, id).Revision(revision).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudGatewaysListenersDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGatewaysListenersDestroy`: GatewayOperation
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudGatewaysListenersDestroy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gatewayId** | **int32** |  | 
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGatewaysListenersDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **revision** | **int32** | The gateway&#39;s desired revision from the preview; without it, a preview. | 

### Return type

[**GatewayOperation**](GatewayOperation.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGatewaysListenersList

> PaginatedGatewayLBListenerList CloudGatewaysListenersList(ctx, gatewayId).Page(page).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	gatewayId := int32(56) // int32 | 
	page := int32(56) // int32 | A page number within the paginated result set. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudGatewaysListenersList(context.Background(), gatewayId).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudGatewaysListenersList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGatewaysListenersList`: PaginatedGatewayLBListenerList
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudGatewaysListenersList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gatewayId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGatewaysListenersListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | A page number within the paginated result set. | 

### Return type

[**PaginatedGatewayLBListenerList**](PaginatedGatewayLBListenerList.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGatewaysListenersPartialUpdate

> GatewayLBListener CloudGatewaysListenersPartialUpdate(ctx, gatewayId, id).PatchedGatewayLBListenerWriteRequest(patchedGatewayLBListenerWriteRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	gatewayId := int32(56) // int32 | 
	id := int32(56) // int32 | 
	patchedGatewayLBListenerWriteRequest := *openapiclient.NewPatchedGatewayLBListenerWriteRequest() // PatchedGatewayLBListenerWriteRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudGatewaysListenersPartialUpdate(context.Background(), gatewayId, id).PatchedGatewayLBListenerWriteRequest(patchedGatewayLBListenerWriteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudGatewaysListenersPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGatewaysListenersPartialUpdate`: GatewayLBListener
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudGatewaysListenersPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gatewayId** | **int32** |  | 
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGatewaysListenersPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchedGatewayLBListenerWriteRequest** | [**PatchedGatewayLBListenerWriteRequest**](PatchedGatewayLBListenerWriteRequest.md) |  | 

### Return type

[**GatewayLBListener**](GatewayLBListener.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGatewaysListenersRetrieve

> GatewayLBListener CloudGatewaysListenersRetrieve(ctx, gatewayId, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	gatewayId := int32(56) // int32 | 
	id := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudGatewaysListenersRetrieve(context.Background(), gatewayId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudGatewaysListenersRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGatewaysListenersRetrieve`: GatewayLBListener
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudGatewaysListenersRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gatewayId** | **int32** |  | 
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGatewaysListenersRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GatewayLBListener**](GatewayLBListener.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGatewaysListenersRoutesCreate

> GatewayRoute CloudGatewaysListenersRoutesCreate(ctx, gatewayId, listenerId).GatewayRouteWriteRequest(gatewayRouteWriteRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	gatewayId := int32(56) // int32 | 
	listenerId := int32(56) // int32 | 
	gatewayRouteWriteRequest := *openapiclient.NewGatewayRouteWriteRequest() // GatewayRouteWriteRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudGatewaysListenersRoutesCreate(context.Background(), gatewayId, listenerId).GatewayRouteWriteRequest(gatewayRouteWriteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudGatewaysListenersRoutesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGatewaysListenersRoutesCreate`: GatewayRoute
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudGatewaysListenersRoutesCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gatewayId** | **int32** |  | 
**listenerId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGatewaysListenersRoutesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **gatewayRouteWriteRequest** | [**GatewayRouteWriteRequest**](GatewayRouteWriteRequest.md) |  | 

### Return type

[**GatewayRoute**](GatewayRoute.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGatewaysListenersRoutesDestroy

> GatewayOperation CloudGatewaysListenersRoutesDestroy(ctx, gatewayId, id, listenerId).Revision(revision).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	gatewayId := int32(56) // int32 | 
	id := int32(56) // int32 | 
	listenerId := int32(56) // int32 | 
	revision := int32(56) // int32 | The gateway's desired revision from the preview; without it, a preview. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudGatewaysListenersRoutesDestroy(context.Background(), gatewayId, id, listenerId).Revision(revision).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudGatewaysListenersRoutesDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGatewaysListenersRoutesDestroy`: GatewayOperation
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudGatewaysListenersRoutesDestroy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gatewayId** | **int32** |  | 
**id** | **int32** |  | 
**listenerId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGatewaysListenersRoutesDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **revision** | **int32** | The gateway&#39;s desired revision from the preview; without it, a preview. | 

### Return type

[**GatewayOperation**](GatewayOperation.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGatewaysListenersRoutesList

> PaginatedGatewayRouteList CloudGatewaysListenersRoutesList(ctx, gatewayId, listenerId).Page(page).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	gatewayId := int32(56) // int32 | 
	listenerId := int32(56) // int32 | 
	page := int32(56) // int32 | A page number within the paginated result set. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudGatewaysListenersRoutesList(context.Background(), gatewayId, listenerId).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudGatewaysListenersRoutesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGatewaysListenersRoutesList`: PaginatedGatewayRouteList
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudGatewaysListenersRoutesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gatewayId** | **int32** |  | 
**listenerId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGatewaysListenersRoutesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | A page number within the paginated result set. | 

### Return type

[**PaginatedGatewayRouteList**](PaginatedGatewayRouteList.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGatewaysListenersRoutesPartialUpdate

> GatewayRoute CloudGatewaysListenersRoutesPartialUpdate(ctx, gatewayId, id, listenerId).PatchedGatewayRouteWriteRequest(patchedGatewayRouteWriteRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	gatewayId := int32(56) // int32 | 
	id := int32(56) // int32 | 
	listenerId := int32(56) // int32 | 
	patchedGatewayRouteWriteRequest := *openapiclient.NewPatchedGatewayRouteWriteRequest() // PatchedGatewayRouteWriteRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudGatewaysListenersRoutesPartialUpdate(context.Background(), gatewayId, id, listenerId).PatchedGatewayRouteWriteRequest(patchedGatewayRouteWriteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudGatewaysListenersRoutesPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGatewaysListenersRoutesPartialUpdate`: GatewayRoute
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudGatewaysListenersRoutesPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gatewayId** | **int32** |  | 
**id** | **int32** |  | 
**listenerId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGatewaysListenersRoutesPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **patchedGatewayRouteWriteRequest** | [**PatchedGatewayRouteWriteRequest**](PatchedGatewayRouteWriteRequest.md) |  | 

### Return type

[**GatewayRoute**](GatewayRoute.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGatewaysListenersRoutesRetrieve

> GatewayRoute CloudGatewaysListenersRoutesRetrieve(ctx, gatewayId, id, listenerId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	gatewayId := int32(56) // int32 | 
	id := int32(56) // int32 | 
	listenerId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudGatewaysListenersRoutesRetrieve(context.Background(), gatewayId, id, listenerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudGatewaysListenersRoutesRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGatewaysListenersRoutesRetrieve`: GatewayRoute
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudGatewaysListenersRoutesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gatewayId** | **int32** |  | 
**id** | **int32** |  | 
**listenerId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGatewaysListenersRoutesRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**GatewayRoute**](GatewayRoute.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGatewaysNetworksCreate

> GatewayOperation CloudGatewaysNetworksCreate(ctx, gatewayId).GatewayNetworkCreateRequest(gatewayNetworkCreateRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	gatewayId := int32(56) // int32 | 
	gatewayNetworkCreateRequest := *openapiclient.NewGatewayNetworkCreateRequest(int32(123), openapiclient.GatewayNetworkRoleEnum("router")) // GatewayNetworkCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudGatewaysNetworksCreate(context.Background(), gatewayId).GatewayNetworkCreateRequest(gatewayNetworkCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudGatewaysNetworksCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGatewaysNetworksCreate`: GatewayOperation
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudGatewaysNetworksCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gatewayId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGatewaysNetworksCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gatewayNetworkCreateRequest** | [**GatewayNetworkCreateRequest**](GatewayNetworkCreateRequest.md) |  | 

### Return type

[**GatewayOperation**](GatewayOperation.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGatewaysNetworksDestroy

> GatewayOperation CloudGatewaysNetworksDestroy(ctx, gatewayId, networkId).Revision(revision).Servers(servers).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	gatewayId := int32(56) // int32 | 
	networkId := int32(56) // int32 | 
	revision := int32(56) // int32 | The gateway's desired revision from the preview; without it, a preview. (optional)
	servers := []int32{int32(123)} // []int32 | The preview's server ids. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudGatewaysNetworksDestroy(context.Background(), gatewayId, networkId).Revision(revision).Servers(servers).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudGatewaysNetworksDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGatewaysNetworksDestroy`: GatewayOperation
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudGatewaysNetworksDestroy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gatewayId** | **int32** |  | 
**networkId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGatewaysNetworksDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **revision** | **int32** | The gateway&#39;s desired revision from the preview; without it, a preview. | 
 **servers** | **[]int32** | The preview&#39;s server ids. | 

### Return type

[**GatewayOperation**](GatewayOperation.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGatewaysNetworksList

> PaginatedGatewayNetworkList CloudGatewaysNetworksList(ctx, gatewayId).Page(page).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	gatewayId := int32(56) // int32 | 
	page := int32(56) // int32 | A page number within the paginated result set. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudGatewaysNetworksList(context.Background(), gatewayId).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudGatewaysNetworksList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGatewaysNetworksList`: PaginatedGatewayNetworkList
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudGatewaysNetworksList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gatewayId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGatewaysNetworksListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | A page number within the paginated result set. | 

### Return type

[**PaginatedGatewayNetworkList**](PaginatedGatewayNetworkList.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGatewaysNetworksPartialUpdate

> GatewayNetwork CloudGatewaysNetworksPartialUpdate(ctx, gatewayId, networkId).Revision(revision).Servers(servers).PatchedGatewayNetworkUpdateRequest(patchedGatewayNetworkUpdateRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	gatewayId := int32(56) // int32 | 
	networkId := int32(56) // int32 | 
	revision := int32(56) // int32 | The gateway's desired revision from the preview; without it, a preview. (optional)
	servers := []int32{int32(123)} // []int32 | The preview's server ids. (optional)
	patchedGatewayNetworkUpdateRequest := *openapiclient.NewPatchedGatewayNetworkUpdateRequest() // PatchedGatewayNetworkUpdateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudGatewaysNetworksPartialUpdate(context.Background(), gatewayId, networkId).Revision(revision).Servers(servers).PatchedGatewayNetworkUpdateRequest(patchedGatewayNetworkUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudGatewaysNetworksPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGatewaysNetworksPartialUpdate`: GatewayNetwork
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudGatewaysNetworksPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gatewayId** | **int32** |  | 
**networkId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGatewaysNetworksPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **revision** | **int32** | The gateway&#39;s desired revision from the preview; without it, a preview. | 
 **servers** | **[]int32** | The preview&#39;s server ids. | 
 **patchedGatewayNetworkUpdateRequest** | [**PatchedGatewayNetworkUpdateRequest**](PatchedGatewayNetworkUpdateRequest.md) |  | 

### Return type

[**GatewayNetwork**](GatewayNetwork.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGatewaysNetworksRetrieve

> GatewayNetwork CloudGatewaysNetworksRetrieve(ctx, gatewayId, networkId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	gatewayId := int32(56) // int32 | 
	networkId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudGatewaysNetworksRetrieve(context.Background(), gatewayId, networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudGatewaysNetworksRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGatewaysNetworksRetrieve`: GatewayNetwork
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudGatewaysNetworksRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gatewayId** | **int32** |  | 
**networkId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGatewaysNetworksRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GatewayNetwork**](GatewayNetwork.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGatewaysOperationsList

> PaginatedGatewayOperationList CloudGatewaysOperationsList(ctx, gatewayId).Page(page).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	gatewayId := int32(56) // int32 | 
	page := int32(56) // int32 | A page number within the paginated result set. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudGatewaysOperationsList(context.Background(), gatewayId).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudGatewaysOperationsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGatewaysOperationsList`: PaginatedGatewayOperationList
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudGatewaysOperationsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gatewayId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGatewaysOperationsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | A page number within the paginated result set. | 

### Return type

[**PaginatedGatewayOperationList**](PaginatedGatewayOperationList.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGatewaysOperationsRetrieve

> GatewayOperation CloudGatewaysOperationsRetrieve(ctx, gatewayId, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	gatewayId := int32(56) // int32 | 
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudGatewaysOperationsRetrieve(context.Background(), gatewayId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudGatewaysOperationsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGatewaysOperationsRetrieve`: GatewayOperation
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudGatewaysOperationsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gatewayId** | **int32** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGatewaysOperationsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GatewayOperation**](GatewayOperation.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGatewaysPartialUpdate

> Gateway CloudGatewaysPartialUpdate(ctx, id).PatchedGatewayUpdateRequest(patchedGatewayUpdateRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := "id_example" // string | 
	patchedGatewayUpdateRequest := *openapiclient.NewPatchedGatewayUpdateRequest() // PatchedGatewayUpdateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudGatewaysPartialUpdate(context.Background(), id).PatchedGatewayUpdateRequest(patchedGatewayUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudGatewaysPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGatewaysPartialUpdate`: Gateway
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudGatewaysPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGatewaysPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedGatewayUpdateRequest** | [**PatchedGatewayUpdateRequest**](PatchedGatewayUpdateRequest.md) |  | 

### Return type

[**Gateway**](Gateway.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGatewaysPoolsBackendsCreate

> GatewayBackend CloudGatewaysPoolsBackendsCreate(ctx, gatewayId, poolId).GatewayBackendWriteRequest(gatewayBackendWriteRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	gatewayId := int32(56) // int32 | 
	poolId := int32(56) // int32 | 
	gatewayBackendWriteRequest := *openapiclient.NewGatewayBackendWriteRequest() // GatewayBackendWriteRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudGatewaysPoolsBackendsCreate(context.Background(), gatewayId, poolId).GatewayBackendWriteRequest(gatewayBackendWriteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudGatewaysPoolsBackendsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGatewaysPoolsBackendsCreate`: GatewayBackend
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudGatewaysPoolsBackendsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gatewayId** | **int32** |  | 
**poolId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGatewaysPoolsBackendsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **gatewayBackendWriteRequest** | [**GatewayBackendWriteRequest**](GatewayBackendWriteRequest.md) |  | 

### Return type

[**GatewayBackend**](GatewayBackend.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGatewaysPoolsBackendsDestroy

> GatewayOperation CloudGatewaysPoolsBackendsDestroy(ctx, gatewayId, id, poolId).Revision(revision).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	gatewayId := int32(56) // int32 | 
	id := int32(56) // int32 | 
	poolId := int32(56) // int32 | 
	revision := int32(56) // int32 | The gateway's desired revision from the preview; without it, a preview. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudGatewaysPoolsBackendsDestroy(context.Background(), gatewayId, id, poolId).Revision(revision).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudGatewaysPoolsBackendsDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGatewaysPoolsBackendsDestroy`: GatewayOperation
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudGatewaysPoolsBackendsDestroy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gatewayId** | **int32** |  | 
**id** | **int32** |  | 
**poolId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGatewaysPoolsBackendsDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **revision** | **int32** | The gateway&#39;s desired revision from the preview; without it, a preview. | 

### Return type

[**GatewayOperation**](GatewayOperation.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGatewaysPoolsBackendsList

> PaginatedGatewayBackendList CloudGatewaysPoolsBackendsList(ctx, gatewayId, poolId).Page(page).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	gatewayId := int32(56) // int32 | 
	poolId := int32(56) // int32 | 
	page := int32(56) // int32 | A page number within the paginated result set. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudGatewaysPoolsBackendsList(context.Background(), gatewayId, poolId).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudGatewaysPoolsBackendsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGatewaysPoolsBackendsList`: PaginatedGatewayBackendList
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudGatewaysPoolsBackendsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gatewayId** | **int32** |  | 
**poolId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGatewaysPoolsBackendsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | A page number within the paginated result set. | 

### Return type

[**PaginatedGatewayBackendList**](PaginatedGatewayBackendList.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGatewaysPoolsBackendsPartialUpdate

> GatewayBackend CloudGatewaysPoolsBackendsPartialUpdate(ctx, gatewayId, id, poolId).PatchedGatewayBackendWriteRequest(patchedGatewayBackendWriteRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	gatewayId := int32(56) // int32 | 
	id := int32(56) // int32 | 
	poolId := int32(56) // int32 | 
	patchedGatewayBackendWriteRequest := *openapiclient.NewPatchedGatewayBackendWriteRequest() // PatchedGatewayBackendWriteRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudGatewaysPoolsBackendsPartialUpdate(context.Background(), gatewayId, id, poolId).PatchedGatewayBackendWriteRequest(patchedGatewayBackendWriteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudGatewaysPoolsBackendsPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGatewaysPoolsBackendsPartialUpdate`: GatewayBackend
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudGatewaysPoolsBackendsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gatewayId** | **int32** |  | 
**id** | **int32** |  | 
**poolId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGatewaysPoolsBackendsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **patchedGatewayBackendWriteRequest** | [**PatchedGatewayBackendWriteRequest**](PatchedGatewayBackendWriteRequest.md) |  | 

### Return type

[**GatewayBackend**](GatewayBackend.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGatewaysPoolsBackendsRetrieve

> GatewayBackend CloudGatewaysPoolsBackendsRetrieve(ctx, gatewayId, id, poolId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	gatewayId := int32(56) // int32 | 
	id := int32(56) // int32 | 
	poolId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudGatewaysPoolsBackendsRetrieve(context.Background(), gatewayId, id, poolId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudGatewaysPoolsBackendsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGatewaysPoolsBackendsRetrieve`: GatewayBackend
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudGatewaysPoolsBackendsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gatewayId** | **int32** |  | 
**id** | **int32** |  | 
**poolId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGatewaysPoolsBackendsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**GatewayBackend**](GatewayBackend.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGatewaysPoolsCreate

> GatewayBackendPool CloudGatewaysPoolsCreate(ctx, gatewayId).GatewayBackendPoolWriteRequest(gatewayBackendPoolWriteRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	gatewayId := int32(56) // int32 | 
	gatewayBackendPoolWriteRequest := *openapiclient.NewGatewayBackendPoolWriteRequest() // GatewayBackendPoolWriteRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudGatewaysPoolsCreate(context.Background(), gatewayId).GatewayBackendPoolWriteRequest(gatewayBackendPoolWriteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudGatewaysPoolsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGatewaysPoolsCreate`: GatewayBackendPool
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudGatewaysPoolsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gatewayId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGatewaysPoolsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gatewayBackendPoolWriteRequest** | [**GatewayBackendPoolWriteRequest**](GatewayBackendPoolWriteRequest.md) |  | 

### Return type

[**GatewayBackendPool**](GatewayBackendPool.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGatewaysPoolsDestroy

> GatewayOperation CloudGatewaysPoolsDestroy(ctx, gatewayId, id).Revision(revision).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	gatewayId := int32(56) // int32 | 
	id := int32(56) // int32 | 
	revision := int32(56) // int32 | The gateway's desired revision from the preview; without it, a preview. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudGatewaysPoolsDestroy(context.Background(), gatewayId, id).Revision(revision).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudGatewaysPoolsDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGatewaysPoolsDestroy`: GatewayOperation
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudGatewaysPoolsDestroy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gatewayId** | **int32** |  | 
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGatewaysPoolsDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **revision** | **int32** | The gateway&#39;s desired revision from the preview; without it, a preview. | 

### Return type

[**GatewayOperation**](GatewayOperation.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGatewaysPoolsList

> PaginatedGatewayBackendPoolList CloudGatewaysPoolsList(ctx, gatewayId).Page(page).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	gatewayId := int32(56) // int32 | 
	page := int32(56) // int32 | A page number within the paginated result set. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudGatewaysPoolsList(context.Background(), gatewayId).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudGatewaysPoolsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGatewaysPoolsList`: PaginatedGatewayBackendPoolList
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudGatewaysPoolsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gatewayId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGatewaysPoolsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | A page number within the paginated result set. | 

### Return type

[**PaginatedGatewayBackendPoolList**](PaginatedGatewayBackendPoolList.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGatewaysPoolsPartialUpdate

> GatewayBackendPool CloudGatewaysPoolsPartialUpdate(ctx, gatewayId, id).PatchedGatewayBackendPoolWriteRequest(patchedGatewayBackendPoolWriteRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	gatewayId := int32(56) // int32 | 
	id := int32(56) // int32 | 
	patchedGatewayBackendPoolWriteRequest := *openapiclient.NewPatchedGatewayBackendPoolWriteRequest() // PatchedGatewayBackendPoolWriteRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudGatewaysPoolsPartialUpdate(context.Background(), gatewayId, id).PatchedGatewayBackendPoolWriteRequest(patchedGatewayBackendPoolWriteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudGatewaysPoolsPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGatewaysPoolsPartialUpdate`: GatewayBackendPool
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudGatewaysPoolsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gatewayId** | **int32** |  | 
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGatewaysPoolsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchedGatewayBackendPoolWriteRequest** | [**PatchedGatewayBackendPoolWriteRequest**](PatchedGatewayBackendPoolWriteRequest.md) |  | 

### Return type

[**GatewayBackendPool**](GatewayBackendPool.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGatewaysPoolsRetrieve

> GatewayBackendPool CloudGatewaysPoolsRetrieve(ctx, gatewayId, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	gatewayId := int32(56) // int32 | 
	id := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudGatewaysPoolsRetrieve(context.Background(), gatewayId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudGatewaysPoolsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGatewaysPoolsRetrieve`: GatewayBackendPool
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudGatewaysPoolsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gatewayId** | **int32** |  | 
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGatewaysPoolsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GatewayBackendPool**](GatewayBackendPool.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGatewaysPortForwardsCreate

> GatewayForward CloudGatewaysPortForwardsCreate(ctx, gatewayId).GatewayForwardWriteRequest(gatewayForwardWriteRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	gatewayId := int32(56) // int32 | 
	gatewayForwardWriteRequest := *openapiclient.NewGatewayForwardWriteRequest(openapiclient.GatewayForwardProtocolEnum("tcp"), *openapiclient.NewGatewayPortRangeRequest(int32(123), int32(123)), *openapiclient.NewGatewayForwardTargetRequest(int32(123), "Address_example", int32(123))) // GatewayForwardWriteRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudGatewaysPortForwardsCreate(context.Background(), gatewayId).GatewayForwardWriteRequest(gatewayForwardWriteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudGatewaysPortForwardsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGatewaysPortForwardsCreate`: GatewayForward
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudGatewaysPortForwardsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gatewayId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGatewaysPortForwardsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gatewayForwardWriteRequest** | [**GatewayForwardWriteRequest**](GatewayForwardWriteRequest.md) |  | 

### Return type

[**GatewayForward**](GatewayForward.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGatewaysPortForwardsDestroy

> GatewayOperation CloudGatewaysPortForwardsDestroy(ctx, gatewayId, id).Revision(revision).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	gatewayId := int32(56) // int32 | 
	id := int32(56) // int32 | 
	revision := int32(56) // int32 | The gateway's desired revision from the preview; without it, a preview. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudGatewaysPortForwardsDestroy(context.Background(), gatewayId, id).Revision(revision).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudGatewaysPortForwardsDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGatewaysPortForwardsDestroy`: GatewayOperation
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudGatewaysPortForwardsDestroy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gatewayId** | **int32** |  | 
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGatewaysPortForwardsDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **revision** | **int32** | The gateway&#39;s desired revision from the preview; without it, a preview. | 

### Return type

[**GatewayOperation**](GatewayOperation.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGatewaysPortForwardsList

> PaginatedGatewayForwardList CloudGatewaysPortForwardsList(ctx, gatewayId).Page(page).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	gatewayId := int32(56) // int32 | 
	page := int32(56) // int32 | A page number within the paginated result set. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudGatewaysPortForwardsList(context.Background(), gatewayId).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudGatewaysPortForwardsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGatewaysPortForwardsList`: PaginatedGatewayForwardList
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudGatewaysPortForwardsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gatewayId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGatewaysPortForwardsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | A page number within the paginated result set. | 

### Return type

[**PaginatedGatewayForwardList**](PaginatedGatewayForwardList.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGatewaysPortForwardsPartialUpdate

> GatewayForward CloudGatewaysPortForwardsPartialUpdate(ctx, gatewayId, id).PatchedGatewayForwardWriteRequest(patchedGatewayForwardWriteRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	gatewayId := int32(56) // int32 | 
	id := int32(56) // int32 | 
	patchedGatewayForwardWriteRequest := *openapiclient.NewPatchedGatewayForwardWriteRequest() // PatchedGatewayForwardWriteRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudGatewaysPortForwardsPartialUpdate(context.Background(), gatewayId, id).PatchedGatewayForwardWriteRequest(patchedGatewayForwardWriteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudGatewaysPortForwardsPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGatewaysPortForwardsPartialUpdate`: GatewayForward
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudGatewaysPortForwardsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gatewayId** | **int32** |  | 
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGatewaysPortForwardsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchedGatewayForwardWriteRequest** | [**PatchedGatewayForwardWriteRequest**](PatchedGatewayForwardWriteRequest.md) |  | 

### Return type

[**GatewayForward**](GatewayForward.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGatewaysPortForwardsRetrieve

> GatewayForward CloudGatewaysPortForwardsRetrieve(ctx, gatewayId, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	gatewayId := int32(56) // int32 | 
	id := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudGatewaysPortForwardsRetrieve(context.Background(), gatewayId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudGatewaysPortForwardsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGatewaysPortForwardsRetrieve`: GatewayForward
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudGatewaysPortForwardsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gatewayId** | **int32** |  | 
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGatewaysPortForwardsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GatewayForward**](GatewayForward.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGatewaysReenrollCreate

> GatewayOperation CloudGatewaysReenrollCreate(ctx, id).GatewayReenrollRequest(gatewayReenrollRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := "id_example" // string | 
	gatewayReenrollRequest := *openapiclient.NewGatewayReenrollRequest() // GatewayReenrollRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudGatewaysReenrollCreate(context.Background(), id).GatewayReenrollRequest(gatewayReenrollRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudGatewaysReenrollCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGatewaysReenrollCreate`: GatewayOperation
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudGatewaysReenrollCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGatewaysReenrollCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gatewayReenrollRequest** | [**GatewayReenrollRequest**](GatewayReenrollRequest.md) |  | 

### Return type

[**GatewayOperation**](GatewayOperation.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGatewaysRetrieve

> Gateway CloudGatewaysRetrieve(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudGatewaysRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudGatewaysRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGatewaysRetrieve`: Gateway
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudGatewaysRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGatewaysRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Gateway**](Gateway.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGatewaysStatusRetrieve

> GatewayStatus CloudGatewaysStatusRetrieve(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudGatewaysStatusRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudGatewaysStatusRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGatewaysStatusRetrieve`: GatewayStatus
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudGatewaysStatusRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGatewaysStatusRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GatewayStatus**](GatewayStatus.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGenerationsList

> []HardwareGeneration CloudGenerationsList(ctx).Execute()

List hardware generations



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudGenerationsList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudGenerationsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGenerationsList`: []HardwareGeneration
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudGenerationsList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGenerationsListRequest struct via the builder pattern


### Return type

[**[]HardwareGeneration**](HardwareGeneration.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGenerationsRetrieve

> HardwareGeneration CloudGenerationsRetrieve(ctx, slug).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	slug := "slug_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudGenerationsRetrieve(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudGenerationsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGenerationsRetrieve`: HardwareGeneration
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudGenerationsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGenerationsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**HardwareGeneration**](HardwareGeneration.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudImagesList

> PaginatedOSImageList CloudImagesList(ctx).Page(page).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	page := int32(56) // int32 | A page number within the paginated result set. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudImagesList(context.Background()).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudImagesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudImagesList`: PaginatedOSImageList
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudImagesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudImagesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | A page number within the paginated result set. | 

### Return type

[**PaginatedOSImageList**](PaginatedOSImageList.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudImagesRetrieve

> OSImage CloudImagesRetrieve(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this operating system.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudImagesRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudImagesRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudImagesRetrieve`: OSImage
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudImagesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this operating system. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudImagesRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OSImage**](OSImage.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudIpv4Create

> PublicIPv4 CloudIpv4Create(ctx).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudIpv4Create(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudIpv4Create``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudIpv4Create`: PublicIPv4
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudIpv4Create`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudIpv4CreateRequest struct via the builder pattern


### Return type

[**PublicIPv4**](PublicIPv4.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudIpv4Destroy

> CloudIpv4Destroy(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this Public IPv4.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CloudAPI.CloudIpv4Destroy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudIpv4Destroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this Public IPv4. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudIpv4DestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudIpv4DetachCreate

> DetachIPv4Response CloudIpv4DetachCreate(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this Public IPv4.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudIpv4DetachCreate(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudIpv4DetachCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudIpv4DetachCreate`: DetachIPv4Response
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudIpv4DetachCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this Public IPv4. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudIpv4DetachCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DetachIPv4Response**](DetachIPv4Response.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudIpv4List

> PaginatedPublicIPv4List CloudIpv4List(ctx).Page(page).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	page := int32(56) // int32 | A page number within the paginated result set. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudIpv4List(context.Background()).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudIpv4List``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudIpv4List`: PaginatedPublicIPv4List
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudIpv4List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudIpv4ListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | A page number within the paginated result set. | 

### Return type

[**PaginatedPublicIPv4List**](PaginatedPublicIPv4List.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudIpv4RdnsCreate

> ReverseDNS CloudIpv4RdnsCreate(ctx, id).ReverseDNSRequest(reverseDNSRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this Public IPv4.
	reverseDNSRequest := *openapiclient.NewReverseDNSRequest("ReverseDns_example") // ReverseDNSRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudIpv4RdnsCreate(context.Background(), id).ReverseDNSRequest(reverseDNSRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudIpv4RdnsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudIpv4RdnsCreate`: ReverseDNS
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudIpv4RdnsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this Public IPv4. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudIpv4RdnsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reverseDNSRequest** | [**ReverseDNSRequest**](ReverseDNSRequest.md) |  | 

### Return type

[**ReverseDNS**](ReverseDNS.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudIpv4RdnsRetrieve

> ReverseDNS CloudIpv4RdnsRetrieve(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this Public IPv4.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudIpv4RdnsRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudIpv4RdnsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudIpv4RdnsRetrieve`: ReverseDNS
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudIpv4RdnsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this Public IPv4. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudIpv4RdnsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ReverseDNS**](ReverseDNS.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudIpv4Retrieve

> PublicIPv4 CloudIpv4Retrieve(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this Public IPv4.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudIpv4Retrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudIpv4Retrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudIpv4Retrieve`: PublicIPv4
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudIpv4Retrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this Public IPv4. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudIpv4RetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PublicIPv4**](PublicIPv4.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudIpv6Create

> PublicIPv6 CloudIpv6Create(ctx).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudIpv6Create(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudIpv6Create``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudIpv6Create`: PublicIPv6
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudIpv6Create`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudIpv6CreateRequest struct via the builder pattern


### Return type

[**PublicIPv6**](PublicIPv6.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudIpv6Destroy

> CloudIpv6Destroy(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this Public IPv6.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CloudAPI.CloudIpv6Destroy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudIpv6Destroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this Public IPv6. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudIpv6DestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudIpv6DetachCreate

> DetachIPv6Response CloudIpv6DetachCreate(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this Public IPv6.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudIpv6DetachCreate(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudIpv6DetachCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudIpv6DetachCreate`: DetachIPv6Response
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudIpv6DetachCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this Public IPv6. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudIpv6DetachCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DetachIPv6Response**](DetachIPv6Response.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudIpv6List

> PaginatedPublicIPv6List CloudIpv6List(ctx).Page(page).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	page := int32(56) // int32 | A page number within the paginated result set. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudIpv6List(context.Background()).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudIpv6List``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudIpv6List`: PaginatedPublicIPv6List
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudIpv6List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudIpv6ListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | A page number within the paginated result set. | 

### Return type

[**PaginatedPublicIPv6List**](PaginatedPublicIPv6List.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudIpv6RdnsCreate

> ReverseDNS CloudIpv6RdnsCreate(ctx, id).ReverseDNSRequest(reverseDNSRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this Public IPv6.
	reverseDNSRequest := *openapiclient.NewReverseDNSRequest("ReverseDns_example") // ReverseDNSRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudIpv6RdnsCreate(context.Background(), id).ReverseDNSRequest(reverseDNSRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudIpv6RdnsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudIpv6RdnsCreate`: ReverseDNS
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudIpv6RdnsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this Public IPv6. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudIpv6RdnsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reverseDNSRequest** | [**ReverseDNSRequest**](ReverseDNSRequest.md) |  | 

### Return type

[**ReverseDNS**](ReverseDNS.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudIpv6RdnsRetrieve

> ReverseDNS CloudIpv6RdnsRetrieve(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this Public IPv6.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudIpv6RdnsRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudIpv6RdnsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudIpv6RdnsRetrieve`: ReverseDNS
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudIpv6RdnsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this Public IPv6. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudIpv6RdnsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ReverseDNS**](ReverseDNS.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudIpv6Retrieve

> PublicIPv6 CloudIpv6Retrieve(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this Public IPv6.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudIpv6Retrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudIpv6Retrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudIpv6Retrieve`: PublicIPv6
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudIpv6Retrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this Public IPv6. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudIpv6RetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PublicIPv6**](PublicIPv6.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPrivateNetworksAddServerCreate

> AddServerResponse CloudPrivateNetworksAddServerCreate(ctx, id).PrivateNetworkAddHostRequest(privateNetworkAddHostRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this private network.
	privateNetworkAddHostRequest := *openapiclient.NewPrivateNetworkAddHostRequest("Server_example") // PrivateNetworkAddHostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudPrivateNetworksAddServerCreate(context.Background(), id).PrivateNetworkAddHostRequest(privateNetworkAddHostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudPrivateNetworksAddServerCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPrivateNetworksAddServerCreate`: AddServerResponse
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudPrivateNetworksAddServerCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this private network. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPrivateNetworksAddServerCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **privateNetworkAddHostRequest** | [**PrivateNetworkAddHostRequest**](PrivateNetworkAddHostRequest.md) |  | 

### Return type

[**AddServerResponse**](AddServerResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPrivateNetworksApplyRouteCreate

> GatewayOperation CloudPrivateNetworksApplyRouteCreate(ctx, id).ApplyRouteRequest(applyRouteRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | 
	applyRouteRequest := *openapiclient.NewApplyRouteRequest([]int32{int32(123)}) // ApplyRouteRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudPrivateNetworksApplyRouteCreate(context.Background(), id).ApplyRouteRequest(applyRouteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudPrivateNetworksApplyRouteCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPrivateNetworksApplyRouteCreate`: GatewayOperation
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudPrivateNetworksApplyRouteCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPrivateNetworksApplyRouteCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **applyRouteRequest** | [**ApplyRouteRequest**](ApplyRouteRequest.md) |  | 

### Return type

[**GatewayOperation**](GatewayOperation.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPrivateNetworksCreate

> PrivateNetwork CloudPrivateNetworksCreate(ctx).PrivateNetworkRequest(privateNetworkRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	privateNetworkRequest := *openapiclient.NewPrivateNetworkRequest("Address_example") // PrivateNetworkRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudPrivateNetworksCreate(context.Background()).PrivateNetworkRequest(privateNetworkRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudPrivateNetworksCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPrivateNetworksCreate`: PrivateNetwork
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudPrivateNetworksCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPrivateNetworksCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **privateNetworkRequest** | [**PrivateNetworkRequest**](PrivateNetworkRequest.md) |  | 

### Return type

[**PrivateNetwork**](PrivateNetwork.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPrivateNetworksDestroy

> CloudPrivateNetworksDestroy(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this private network.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CloudAPI.CloudPrivateNetworksDestroy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudPrivateNetworksDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this private network. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPrivateNetworksDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPrivateNetworksList

> PaginatedPrivateNetworkList CloudPrivateNetworksList(ctx).Page(page).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	page := int32(56) // int32 | A page number within the paginated result set. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudPrivateNetworksList(context.Background()).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudPrivateNetworksList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPrivateNetworksList`: PaginatedPrivateNetworkList
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudPrivateNetworksList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPrivateNetworksListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | A page number within the paginated result set. | 

### Return type

[**PaginatedPrivateNetworkList**](PaginatedPrivateNetworkList.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPrivateNetworksPartialUpdate

> PrivateNetwork CloudPrivateNetworksPartialUpdate(ctx, id).PatchedPrivateNetworkUpdateRequest(patchedPrivateNetworkUpdateRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this private network.
	patchedPrivateNetworkUpdateRequest := *openapiclient.NewPatchedPrivateNetworkUpdateRequest() // PatchedPrivateNetworkUpdateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudPrivateNetworksPartialUpdate(context.Background(), id).PatchedPrivateNetworkUpdateRequest(patchedPrivateNetworkUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudPrivateNetworksPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPrivateNetworksPartialUpdate`: PrivateNetwork
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudPrivateNetworksPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this private network. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPrivateNetworksPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedPrivateNetworkUpdateRequest** | [**PatchedPrivateNetworkUpdateRequest**](PatchedPrivateNetworkUpdateRequest.md) |  | 

### Return type

[**PrivateNetwork**](PrivateNetwork.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPrivateNetworksRemoveServerCreate

> RemoveServerResponse CloudPrivateNetworksRemoveServerCreate(ctx, id).PrivateNetworkRemoveHostRequest(privateNetworkRemoveHostRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this private network.
	privateNetworkRemoveHostRequest := *openapiclient.NewPrivateNetworkRemoveHostRequest("Server_example") // PrivateNetworkRemoveHostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudPrivateNetworksRemoveServerCreate(context.Background(), id).PrivateNetworkRemoveHostRequest(privateNetworkRemoveHostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudPrivateNetworksRemoveServerCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPrivateNetworksRemoveServerCreate`: RemoveServerResponse
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudPrivateNetworksRemoveServerCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this private network. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPrivateNetworksRemoveServerCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **privateNetworkRemoveHostRequest** | [**PrivateNetworkRemoveHostRequest**](PrivateNetworkRemoveHostRequest.md) |  | 

### Return type

[**RemoveServerResponse**](RemoveServerResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPrivateNetworksRetrieve

> PrivateNetwork CloudPrivateNetworksRetrieve(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this private network.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudPrivateNetworksRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudPrivateNetworksRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPrivateNetworksRetrieve`: PrivateNetwork
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudPrivateNetworksRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this private network. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPrivateNetworksRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PrivateNetwork**](PrivateNetwork.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPrivateNetworksRouteStatusRetrieve

> RouteStatus CloudPrivateNetworksRouteStatusRetrieve(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudPrivateNetworksRouteStatusRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudPrivateNetworksRouteStatusRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPrivateNetworksRouteStatusRetrieve`: RouteStatus
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudPrivateNetworksRouteStatusRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPrivateNetworksRouteStatusRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RouteStatus**](RouteStatus.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPrivateNetworksUpdate

> PrivateNetwork CloudPrivateNetworksUpdate(ctx, id).PrivateNetworkUpdateRequest(privateNetworkUpdateRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this private network.
	privateNetworkUpdateRequest := *openapiclient.NewPrivateNetworkUpdateRequest() // PrivateNetworkUpdateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudPrivateNetworksUpdate(context.Background(), id).PrivateNetworkUpdateRequest(privateNetworkUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudPrivateNetworksUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPrivateNetworksUpdate`: PrivateNetwork
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudPrivateNetworksUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this private network. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPrivateNetworksUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **privateNetworkUpdateRequest** | [**PrivateNetworkUpdateRequest**](PrivateNetworkUpdateRequest.md) |  | 

### Return type

[**PrivateNetwork**](PrivateNetwork.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudServerPackagesByGenerationRetrieve

> ServerProduct CloudServerPackagesByGenerationRetrieve(ctx).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudServerPackagesByGenerationRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudServerPackagesByGenerationRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudServerPackagesByGenerationRetrieve`: ServerProduct
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudServerPackagesByGenerationRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudServerPackagesByGenerationRetrieveRequest struct via the builder pattern


### Return type

[**ServerProduct**](ServerProduct.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudServerPackagesList

> PaginatedServerProductList CloudServerPackagesList(ctx).Generation(generation).Page(page).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	generation := "generation_example" // string | Filter packages available on the given hardware generation (slug). Excludes free-tier-only packages when the generation is not free-tier eligible. (optional)
	page := int32(56) // int32 | A page number within the paginated result set. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudServerPackagesList(context.Background()).Generation(generation).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudServerPackagesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudServerPackagesList`: PaginatedServerProductList
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudServerPackagesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudServerPackagesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **generation** | **string** | Filter packages available on the given hardware generation (slug). Excludes free-tier-only packages when the generation is not free-tier eligible. | 
 **page** | **int32** | A page number within the paginated result set. | 

### Return type

[**PaginatedServerProductList**](PaginatedServerProductList.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudServerPackagesRetrieve

> ServerProduct CloudServerPackagesRetrieve(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this metered product.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudServerPackagesRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudServerPackagesRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudServerPackagesRetrieve`: ServerProduct
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudServerPackagesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this metered product. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudServerPackagesRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServerProduct**](ServerProduct.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudServersActivityRetrieve

> ActivityLogResponse CloudServersActivityRetrieve(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this virtual machine.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudServersActivityRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudServersActivityRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudServersActivityRetrieve`: ActivityLogResponse
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudServersActivityRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this virtual machine. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudServersActivityRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ActivityLogResponse**](ActivityLogResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudServersAttachIpv4Create

> AttachIPv4Response CloudServersAttachIpv4Create(ctx, id).AttachIPv4Request(attachIPv4Request).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this virtual machine.
	attachIPv4Request := *openapiclient.NewAttachIPv4Request("Ipv4_example") // AttachIPv4Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudServersAttachIpv4Create(context.Background(), id).AttachIPv4Request(attachIPv4Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudServersAttachIpv4Create``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudServersAttachIpv4Create`: AttachIPv4Response
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudServersAttachIpv4Create`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this virtual machine. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudServersAttachIpv4CreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **attachIPv4Request** | [**AttachIPv4Request**](AttachIPv4Request.md) |  | 

### Return type

[**AttachIPv4Response**](AttachIPv4Response.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudServersAttachIpv6Create

> AttachIPv6Response CloudServersAttachIpv6Create(ctx, id).AttachIPv6Request(attachIPv6Request).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this virtual machine.
	attachIPv6Request := *openapiclient.NewAttachIPv6Request("Ipv6_example") // AttachIPv6Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudServersAttachIpv6Create(context.Background(), id).AttachIPv6Request(attachIPv6Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudServersAttachIpv6Create``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudServersAttachIpv6Create`: AttachIPv6Response
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudServersAttachIpv6Create`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this virtual machine. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudServersAttachIpv6CreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **attachIPv6Request** | [**AttachIPv6Request**](AttachIPv6Request.md) |  | 

### Return type

[**AttachIPv6Response**](AttachIPv6Response.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudServersBootIsosList

> []BootISO CloudServersBootIsosList(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this virtual machine.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudServersBootIsosList(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudServersBootIsosList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudServersBootIsosList`: []BootISO
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudServersBootIsosList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this virtual machine. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudServersBootIsosListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]BootISO**](BootISO.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudServersConsoleCreate

> ConsoleToken CloudServersConsoleCreate(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this virtual machine.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudServersConsoleCreate(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudServersConsoleCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudServersConsoleCreate`: ConsoleToken
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudServersConsoleCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this virtual machine. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudServersConsoleCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConsoleToken**](ConsoleToken.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudServersCreate

> ServerAddResponse CloudServersCreate(ctx).ServerAddRequest(serverAddRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	serverAddRequest := *openapiclient.NewServerAddRequest("Image_example", "Package_example") // ServerAddRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudServersCreate(context.Background()).ServerAddRequest(serverAddRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudServersCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudServersCreate`: ServerAddResponse
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudServersCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudServersCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serverAddRequest** | [**ServerAddRequest**](ServerAddRequest.md) |  | 

### Return type

[**ServerAddResponse**](ServerAddResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudServersDestroy

> CloudServersDestroy(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this virtual machine.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CloudAPI.CloudServersDestroy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudServersDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this virtual machine. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudServersDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudServersDestroyProtectionCreate

> DestroyProtectionResponse CloudServersDestroyProtectionCreate(ctx, id).DestroyProtectionRequest(destroyProtectionRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this virtual machine.
	destroyProtectionRequest := *openapiclient.NewDestroyProtectionRequest(false) // DestroyProtectionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudServersDestroyProtectionCreate(context.Background(), id).DestroyProtectionRequest(destroyProtectionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudServersDestroyProtectionCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudServersDestroyProtectionCreate`: DestroyProtectionResponse
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudServersDestroyProtectionCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this virtual machine. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudServersDestroyProtectionCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **destroyProtectionRequest** | [**DestroyProtectionRequest**](DestroyProtectionRequest.md) |  | 

### Return type

[**DestroyProtectionResponse**](DestroyProtectionResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudServersDetachIpv4Create

> ServerDetachIPv4Response CloudServersDetachIpv4Create(ctx, id).Ipv4(ipv4).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this virtual machine.
	ipv4 := "ipv4_example" // string | ID or slug of the IPv4 address to detach. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudServersDetachIpv4Create(context.Background(), id).Ipv4(ipv4).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudServersDetachIpv4Create``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudServersDetachIpv4Create`: ServerDetachIPv4Response
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudServersDetachIpv4Create`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this virtual machine. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudServersDetachIpv4CreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ipv4** | **string** | ID or slug of the IPv4 address to detach. | 

### Return type

[**ServerDetachIPv4Response**](ServerDetachIPv4Response.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudServersDetachIpv6Create

> DetachIPv6 CloudServersDetachIpv6Create(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this virtual machine.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudServersDetachIpv6Create(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudServersDetachIpv6Create``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudServersDetachIpv6Create`: DetachIPv6
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudServersDetachIpv6Create`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this virtual machine. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudServersDetachIpv6CreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DetachIPv6**](DetachIPv6.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudServersList

> PaginatedServerList CloudServersList(ctx).Page(page).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	page := int32(56) // int32 | A page number within the paginated result set. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudServersList(context.Background()).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudServersList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudServersList`: PaginatedServerList
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudServersList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudServersListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | A page number within the paginated result set. | 

### Return type

[**PaginatedServerList**](PaginatedServerList.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudServersModifyPackageCreate

> ServerUpgradeResponse CloudServersModifyPackageCreate(ctx, id).ServerProductUpgradeRequest(serverProductUpgradeRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this virtual machine.
	serverProductUpgradeRequest := *openapiclient.NewServerProductUpgradeRequest("Package_example") // ServerProductUpgradeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudServersModifyPackageCreate(context.Background(), id).ServerProductUpgradeRequest(serverProductUpgradeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudServersModifyPackageCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudServersModifyPackageCreate`: ServerUpgradeResponse
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudServersModifyPackageCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this virtual machine. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudServersModifyPackageCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serverProductUpgradeRequest** | [**ServerProductUpgradeRequest**](ServerProductUpgradeRequest.md) |  | 

### Return type

[**ServerUpgradeResponse**](ServerUpgradeResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudServersMonitoringRetrieve

> ServerMonitoring CloudServersMonitoringRetrieve(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this virtual machine.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudServersMonitoringRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudServersMonitoringRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudServersMonitoringRetrieve`: ServerMonitoring
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudServersMonitoringRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this virtual machine. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudServersMonitoringRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServerMonitoring**](ServerMonitoring.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudServersMonitoringUpdate

> ServerMonitoring CloudServersMonitoringUpdate(ctx, id).ServerMonitoringRequest(serverMonitoringRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this virtual machine.
	serverMonitoringRequest := *openapiclient.NewServerMonitoringRequest(false) // ServerMonitoringRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudServersMonitoringUpdate(context.Background(), id).ServerMonitoringRequest(serverMonitoringRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudServersMonitoringUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudServersMonitoringUpdate`: ServerMonitoring
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudServersMonitoringUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this virtual machine. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudServersMonitoringUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serverMonitoringRequest** | [**ServerMonitoringRequest**](ServerMonitoringRequest.md) |  | 

### Return type

[**ServerMonitoring**](ServerMonitoring.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudServersPartialUpdate

> ServerDetail CloudServersPartialUpdate(ctx, id).PatchedServerDetailRequest(patchedServerDetailRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this virtual machine.
	patchedServerDetailRequest := *openapiclient.NewPatchedServerDetailRequest() // PatchedServerDetailRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudServersPartialUpdate(context.Background(), id).PatchedServerDetailRequest(patchedServerDetailRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudServersPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudServersPartialUpdate`: ServerDetail
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudServersPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this virtual machine. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudServersPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedServerDetailRequest** | [**PatchedServerDetailRequest**](PatchedServerDetailRequest.md) |  | 

### Return type

[**ServerDetail**](ServerDetail.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudServersPowerManagementCreate

> PowerManagement CloudServersPowerManagementCreate(ctx, id).PowerManagementRequest(powerManagementRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this virtual machine.
	powerManagementRequest := *openapiclient.NewPowerManagementRequest(openapiclient.PowerManagementRequestActionEnum("start")) // PowerManagementRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudServersPowerManagementCreate(context.Background(), id).PowerManagementRequest(powerManagementRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudServersPowerManagementCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudServersPowerManagementCreate`: PowerManagement
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudServersPowerManagementCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this virtual machine. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudServersPowerManagementCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **powerManagementRequest** | [**PowerManagementRequest**](PowerManagementRequest.md) |  | 

### Return type

[**PowerManagement**](PowerManagement.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudServersPowerManagementRetrieve

> PowerManagement CloudServersPowerManagementRetrieve(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this virtual machine.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudServersPowerManagementRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudServersPowerManagementRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudServersPowerManagementRetrieve`: PowerManagement
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudServersPowerManagementRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this virtual machine. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudServersPowerManagementRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PowerManagement**](PowerManagement.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudServersPublicInterfaceCreate

> PublicInterface CloudServersPublicInterfaceCreate(ctx, id).PublicInterfaceRequest(publicInterfaceRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this virtual machine.
	publicInterfaceRequest := *openapiclient.NewPublicInterfaceRequest() // PublicInterfaceRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudServersPublicInterfaceCreate(context.Background(), id).PublicInterfaceRequest(publicInterfaceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudServersPublicInterfaceCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudServersPublicInterfaceCreate`: PublicInterface
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudServersPublicInterfaceCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this virtual machine. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudServersPublicInterfaceCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **publicInterfaceRequest** | [**PublicInterfaceRequest**](PublicInterfaceRequest.md) |  | 

### Return type

[**PublicInterface**](PublicInterface.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudServersPublicInterfaceDestroy

> CloudServersPublicInterfaceDestroy(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this virtual machine.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CloudAPI.CloudServersPublicInterfaceDestroy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudServersPublicInterfaceDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this virtual machine. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudServersPublicInterfaceDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudServersPublicInterfaceRetrieve

> PublicInterface CloudServersPublicInterfaceRetrieve(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this virtual machine.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudServersPublicInterfaceRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudServersPublicInterfaceRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudServersPublicInterfaceRetrieve`: PublicInterface
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudServersPublicInterfaceRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this virtual machine. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudServersPublicInterfaceRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PublicInterface**](PublicInterface.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudServersRescueEnterCreate

> RescueEnterQueued CloudServersRescueEnterCreate(ctx, id).IsoBootRequest(isoBootRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this virtual machine.
	isoBootRequest := *openapiclient.NewIsoBootRequest() // IsoBootRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudServersRescueEnterCreate(context.Background(), id).IsoBootRequest(isoBootRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudServersRescueEnterCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudServersRescueEnterCreate`: RescueEnterQueued
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudServersRescueEnterCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this virtual machine. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudServersRescueEnterCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **isoBootRequest** | [**IsoBootRequest**](IsoBootRequest.md) |  | 

### Return type

[**RescueEnterQueued**](RescueEnterQueued.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudServersRescueExitCreate

> RescueExitQueued CloudServersRescueExitCreate(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this virtual machine.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudServersRescueExitCreate(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudServersRescueExitCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudServersRescueExitCreate`: RescueExitQueued
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudServersRescueExitCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this virtual machine. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudServersRescueExitCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RescueExitQueued**](RescueExitQueued.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudServersRetrieve

> ServerDetail CloudServersRetrieve(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this virtual machine.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudServersRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudServersRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudServersRetrieve`: ServerDetail
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudServersRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this virtual machine. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudServersRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServerDetail**](ServerDetail.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudServersRetryProvisionCreate

> RetryProvision CloudServersRetryProvisionCreate(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this virtual machine.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudServersRetryProvisionCreate(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudServersRetryProvisionCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudServersRetryProvisionCreate`: RetryProvision
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudServersRetryProvisionCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this virtual machine. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudServersRetryProvisionCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RetryProvision**](RetryProvision.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudServersSnapshotsCreate

> SnapshotCreateQueued CloudServersSnapshotsCreate(ctx, id).SnapshotCreateRequest(snapshotCreateRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this virtual machine.
	snapshotCreateRequest := *openapiclient.NewSnapshotCreateRequest("Name_example") // SnapshotCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudServersSnapshotsCreate(context.Background(), id).SnapshotCreateRequest(snapshotCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudServersSnapshotsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudServersSnapshotsCreate`: SnapshotCreateQueued
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudServersSnapshotsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this virtual machine. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudServersSnapshotsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **snapshotCreateRequest** | [**SnapshotCreateRequest**](SnapshotCreateRequest.md) |  | 

### Return type

[**SnapshotCreateQueued**](SnapshotCreateQueued.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudServersSnapshotsDestroy

> SnapshotDeleteQueued CloudServersSnapshotsDestroy(ctx, id, snapshotName).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this virtual machine.
	snapshotName := "snapshotName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudServersSnapshotsDestroy(context.Background(), id, snapshotName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudServersSnapshotsDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudServersSnapshotsDestroy`: SnapshotDeleteQueued
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudServersSnapshotsDestroy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this virtual machine. | 
**snapshotName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudServersSnapshotsDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SnapshotDeleteQueued**](SnapshotDeleteQueued.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudServersSnapshotsList

> []Snapshot CloudServersSnapshotsList(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this virtual machine.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudServersSnapshotsList(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudServersSnapshotsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudServersSnapshotsList`: []Snapshot
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudServersSnapshotsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this virtual machine. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudServersSnapshotsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Snapshot**](Snapshot.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudServersSnapshotsRollbackCreate

> SnapshotRollbackQueued CloudServersSnapshotsRollbackCreate(ctx, id, snapshotName).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this virtual machine.
	snapshotName := "snapshotName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudServersSnapshotsRollbackCreate(context.Background(), id, snapshotName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudServersSnapshotsRollbackCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudServersSnapshotsRollbackCreate`: SnapshotRollbackQueued
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudServersSnapshotsRollbackCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this virtual machine. | 
**snapshotName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudServersSnapshotsRollbackCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SnapshotRollbackQueued**](SnapshotRollbackQueued.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudServersTrafficRetrieve

> ServerTrafficResponse CloudServersTrafficRetrieve(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this virtual machine.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudServersTrafficRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudServersTrafficRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudServersTrafficRetrieve`: ServerTrafficResponse
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudServersTrafficRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this virtual machine. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudServersTrafficRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServerTrafficResponse**](ServerTrafficResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudServersUpdate

> ServerDetail CloudServersUpdate(ctx, id).ServerDetailRequest(serverDetailRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this virtual machine.
	serverDetailRequest := *openapiclient.NewServerDetailRequest() // ServerDetailRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudServersUpdate(context.Background(), id).ServerDetailRequest(serverDetailRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudServersUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudServersUpdate`: ServerDetail
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudServersUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this virtual machine. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudServersUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serverDetailRequest** | [**ServerDetailRequest**](ServerDetailRequest.md) |  | 

### Return type

[**ServerDetail**](ServerDetail.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudServersUsageRetrieve

> ServerUsageResponse CloudServersUsageRetrieve(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this virtual machine.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudServersUsageRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudServersUsageRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudServersUsageRetrieve`: ServerUsageResponse
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudServersUsageRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this virtual machine. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudServersUsageRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServerUsageResponse**](ServerUsageResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudServersVolumesCreate

> Volume CloudServersVolumesCreate(ctx, serverId).VolumeRequest(volumeRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	serverId := "serverId_example" // string | 
	volumeRequest := *openapiclient.NewVolumeRequest(int32(123), "Product_example") // VolumeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudServersVolumesCreate(context.Background(), serverId).VolumeRequest(volumeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudServersVolumesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudServersVolumesCreate`: Volume
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudServersVolumesCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudServersVolumesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **volumeRequest** | [**VolumeRequest**](VolumeRequest.md) |  | 

### Return type

[**Volume**](Volume.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudServersVolumesDestroy

> CloudServersVolumesDestroy(ctx, serverId, volumeId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	serverId := "serverId_example" // string | 
	volumeId := "volumeId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CloudAPI.CloudServersVolumesDestroy(context.Background(), serverId, volumeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudServersVolumesDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** |  | 
**volumeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudServersVolumesDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudServersVolumesList

> []Volume CloudServersVolumesList(ctx, serverId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	serverId := "serverId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudServersVolumesList(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudServersVolumesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudServersVolumesList`: []Volume
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudServersVolumesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudServersVolumesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Volume**](Volume.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudServersVolumesPartialUpdate

> Volume CloudServersVolumesPartialUpdate(ctx, serverId, volumeId).PatchedVolumeUpdateRequest(patchedVolumeUpdateRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	serverId := "serverId_example" // string | 
	volumeId := "volumeId_example" // string | 
	patchedVolumeUpdateRequest := *openapiclient.NewPatchedVolumeUpdateRequest() // PatchedVolumeUpdateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudServersVolumesPartialUpdate(context.Background(), serverId, volumeId).PatchedVolumeUpdateRequest(patchedVolumeUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudServersVolumesPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudServersVolumesPartialUpdate`: Volume
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudServersVolumesPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** |  | 
**volumeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudServersVolumesPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchedVolumeUpdateRequest** | [**PatchedVolumeUpdateRequest**](PatchedVolumeUpdateRequest.md) |  | 

### Return type

[**Volume**](Volume.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudServersVolumesRetrieve

> Volume CloudServersVolumesRetrieve(ctx, serverId, volumeId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	serverId := "serverId_example" // string | 
	volumeId := "volumeId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudServersVolumesRetrieve(context.Background(), serverId, volumeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudServersVolumesRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudServersVolumesRetrieve`: Volume
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudServersVolumesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** |  | 
**volumeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudServersVolumesRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Volume**](Volume.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudServersVolumesUpdate

> Volume CloudServersVolumesUpdate(ctx, serverId, volumeId).VolumeUpdateRequest(volumeUpdateRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	serverId := "serverId_example" // string | 
	volumeId := "volumeId_example" // string | 
	volumeUpdateRequest := *openapiclient.NewVolumeUpdateRequest(int32(123)) // VolumeUpdateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudServersVolumesUpdate(context.Background(), serverId, volumeId).VolumeUpdateRequest(volumeUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudServersVolumesUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudServersVolumesUpdate`: Volume
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudServersVolumesUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** |  | 
**volumeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudServersVolumesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **volumeUpdateRequest** | [**VolumeUpdateRequest**](VolumeUpdateRequest.md) |  | 

### Return type

[**Volume**](Volume.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudStorageProductsList

> PaginatedStorageProductList CloudStorageProductsList(ctx).Page(page).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	page := int32(56) // int32 | A page number within the paginated result set. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudStorageProductsList(context.Background()).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudStorageProductsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudStorageProductsList`: PaginatedStorageProductList
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudStorageProductsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudStorageProductsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | A page number within the paginated result set. | 

### Return type

[**PaginatedStorageProductList**](PaginatedStorageProductList.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudStorageProductsRetrieve

> StorageProduct CloudStorageProductsRetrieve(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this metered product.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudStorageProductsRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudStorageProductsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudStorageProductsRetrieve`: StorageProduct
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudStorageProductsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this metered product. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudStorageProductsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StorageProduct**](StorageProduct.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudVolumesAttachCreate

> AttachVolume CloudVolumesAttachCreate(ctx, id).AttachVolumeRequest(attachVolumeRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this storage.
	attachVolumeRequest := *openapiclient.NewAttachVolumeRequest(int32(123)) // AttachVolumeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudVolumesAttachCreate(context.Background(), id).AttachVolumeRequest(attachVolumeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudVolumesAttachCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudVolumesAttachCreate`: AttachVolume
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudVolumesAttachCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this storage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudVolumesAttachCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **attachVolumeRequest** | [**AttachVolumeRequest**](AttachVolumeRequest.md) |  | 

### Return type

[**AttachVolume**](AttachVolume.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudVolumesDestroy

> CloudVolumesDestroy(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this storage.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CloudAPI.CloudVolumesDestroy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudVolumesDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this storage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudVolumesDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudVolumesDetachCreate

> DetachVolume CloudVolumesDetachCreate(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this storage.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudVolumesDetachCreate(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudVolumesDetachCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudVolumesDetachCreate`: DetachVolume
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudVolumesDetachCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this storage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudVolumesDetachCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DetachVolume**](DetachVolume.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudVolumesList

> []Volume CloudVolumesList(ctx).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudVolumesList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudVolumesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudVolumesList`: []Volume
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudVolumesList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudVolumesListRequest struct via the builder pattern


### Return type

[**[]Volume**](Volume.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudVolumesPartialUpdate

> Volume CloudVolumesPartialUpdate(ctx, id).PatchedVolumeUpdateRequest(patchedVolumeUpdateRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this storage.
	patchedVolumeUpdateRequest := *openapiclient.NewPatchedVolumeUpdateRequest() // PatchedVolumeUpdateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudVolumesPartialUpdate(context.Background(), id).PatchedVolumeUpdateRequest(patchedVolumeUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudVolumesPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudVolumesPartialUpdate`: Volume
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudVolumesPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this storage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudVolumesPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedVolumeUpdateRequest** | [**PatchedVolumeUpdateRequest**](PatchedVolumeUpdateRequest.md) |  | 

### Return type

[**Volume**](Volume.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudVolumesRetrieve

> Volume CloudVolumesRetrieve(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this storage.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudVolumesRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudVolumesRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudVolumesRetrieve`: Volume
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudVolumesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this storage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudVolumesRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Volume**](Volume.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudVolumesUpdate

> Volume CloudVolumesUpdate(ctx, id).VolumeUpdateRequest(volumeUpdateRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this storage.
	volumeUpdateRequest := *openapiclient.NewVolumeUpdateRequest(int32(123)) // VolumeUpdateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudVolumesUpdate(context.Background(), id).VolumeUpdateRequest(volumeUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudVolumesUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudVolumesUpdate`: Volume
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudVolumesUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this storage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudVolumesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **volumeUpdateRequest** | [**VolumeUpdateRequest**](VolumeUpdateRequest.md) |  | 

### Return type

[**Volume**](Volume.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

