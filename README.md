# mailcow Go API Client

This repo contains a generated API client to talk with mailcow's API from Go.

## Specials

The following API endpoints require an array as response body:

* /delete/domain
* /delete/mailbox

Therfore the template client.mustache include a special handling, e.g:

```
		if strings.HasSuffix(path, "/api/v1/delete/domain") {
			var deleteDomainRequest *DeleteDomainRequest = postBody.(*DeleteDomainRequest)
			postBody = *&deleteDomainRequest.Items
		}

```

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.0

## Installation

Install the following dependencies:

```shell
go install github.com/l.with/mailcow-go
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AddressRewritingApi* | [**CreateBCCMap**](docs/AddressRewritingApi.md#createbccmap) | **Post** /api/v1/add/bcc | Create BCC Map
*AddressRewritingApi* | [**CreateRecipientMap**](docs/AddressRewritingApi.md#createrecipientmap) | **Post** /api/v1/add/recipient_map | Create Recipient Map
*AddressRewritingApi* | [**DeleteBCCMap**](docs/AddressRewritingApi.md#deletebccmap) | **Post** /api/v1/delete/bcc | Delete BCC Map
*AddressRewritingApi* | [**DeleteRecipientMap**](docs/AddressRewritingApi.md#deleterecipientmap) | **Post** /api/v1/delete/recipient_map | Delete Recipient Map
*AddressRewritingApi* | [**GetBCCMap**](docs/AddressRewritingApi.md#getbccmap) | **Get** /api/v1/get/bcc/{id} | Get BCC Map
*AddressRewritingApi* | [**GetRecipientMap**](docs/AddressRewritingApi.md#getrecipientmap) | **Get** /api/v1/get/recipient_map/{id} | Get Recipient Map
*AliasesApi* | [**CreateAlias**](docs/AliasesApi.md#createalias) | **Post** /api/v1/add/alias | Create alias
*AliasesApi* | [**CreateTimeLimitedAlias**](docs/AliasesApi.md#createtimelimitedalias) | **Post** /api/v1/add/time_limited_alias | Create time limited alias
*AliasesApi* | [**DeleteAlias**](docs/AliasesApi.md#deletealias) | **Post** /api/v1/delete/alias | Delete alias
*AliasesApi* | [**GetAliases**](docs/AliasesApi.md#getaliases) | **Get** /api/v1/get/alias/{id} | Get aliases
*AliasesApi* | [**GetTimeLimitedAliases**](docs/AliasesApi.md#gettimelimitedaliases) | **Get** /api/v1/get/time_limited_aliases/{mailbox} | Get time limited aliases
*AliasesApi* | [**UpdateAlias**](docs/AliasesApi.md#updatealias) | **Post** /api/v1/edit/alias | Update alias
*AppPasswordsApi* | [**CreateAppPassword**](docs/AppPasswordsApi.md#createapppassword) | **Post** /api/v1/add/app-passwd | Create App Password
*AppPasswordsApi* | [**DeleteAppPassword**](docs/AppPasswordsApi.md#deleteapppassword) | **Post** /api/v1/delete/app-passwd | Delete App Password
*AppPasswordsApi* | [**GetAppPassword**](docs/AppPasswordsApi.md#getapppassword) | **Get** /api/v1/get/app-passwd/all/{mailbox} | Get App Password
*DKIMApi* | [**DeleteDKIMKey**](docs/DKIMApi.md#deletedkimkey) | **Post** /api/v1/delete/dkim | Delete DKIM Key
*DKIMApi* | [**DuplicateDKIMKey**](docs/DKIMApi.md#duplicatedkimkey) | **Post** /api/v1/add/dkim_duplicate | Duplicate DKIM Key
*DKIMApi* | [**GenerateDKIMKey**](docs/DKIMApi.md#generatedkimkey) | **Post** /api/v1/add/dkim | Generate DKIM Key
*DKIMApi* | [**GetDKIMKey**](docs/DKIMApi.md#getdkimkey) | **Get** /api/v1/get/dkim/{domain} | Get DKIM Key
*DomainAdminApi* | [**CreateDomainAdminUser**](docs/DomainAdminApi.md#createdomainadminuser) | **Post** /api/v1/add/domain-admin | Create Domain Admin user
*DomainAdminApi* | [**DeleteDomainAdmin**](docs/DomainAdminApi.md#deletedomainadmin) | **Post** /api/v1/delete/domain-admin | Delete Domain Admin
*DomainAdminApi* | [**EditDomainAdminACL**](docs/DomainAdminApi.md#editdomainadminacl) | **Post** /api/v1/edit/da-acl | Edit Domain Admin ACL
*DomainAdminApi* | [**EditDomainAdminUser**](docs/DomainAdminApi.md#editdomainadminuser) | **Post** /api/v1/edit/domain-admin | Edit Domain Admin user
*DomainAdminApi* | [**GetDomainAdmins**](docs/DomainAdminApi.md#getdomainadmins) | **Get** /api/v1/get/domain-admin/all | Get Domain Admins
*DomainAntispamPoliciesApi* | [**CreateDomainPolicy**](docs/DomainAntispamPoliciesApi.md#createdomainpolicy) | **Post** /api/v1/add/domain-policy | Create domain policy
*DomainAntispamPoliciesApi* | [**DeleteDomainPolicy**](docs/DomainAntispamPoliciesApi.md#deletedomainpolicy) | **Post** /api/v1/delete/domain-policy | Delete domain policy
*DomainAntispamPoliciesApi* | [**ListBlacklistDomainPolicy**](docs/DomainAntispamPoliciesApi.md#listblacklistdomainpolicy) | **Get** /api/v1/get/policy_bl_domain/{domain} | List blacklist domain policy
*DomainAntispamPoliciesApi* | [**ListWhitelistDomainPolicy**](docs/DomainAntispamPoliciesApi.md#listwhitelistdomainpolicy) | **Get** /api/v1/get/policy_wl_domain/{domain} | List whitelist domain policy
*DomainsApi* | [**CreateDomain**](docs/DomainsApi.md#createdomain) | **Post** /api/v1/add/domain | Create domain
*DomainsApi* | [**DeleteDomain**](docs/DomainsApi.md#deletedomain) | **Post** /api/v1/delete/domain | Delete domain
*DomainsApi* | [**DeleteDomainTags**](docs/DomainsApi.md#deletedomaintags) | **Post** /api/v1/delete/domain/tag/{domain} | Delete domain tags
*DomainsApi* | [**GetDomains**](docs/DomainsApi.md#getdomains) | **Get** /api/v1/get/domain/{id} | Get domains
*DomainsApi* | [**UpdateDomain**](docs/DomainsApi.md#updatedomain) | **Post** /api/v1/edit/domain | Update domain
*Fail2BanApi* | [**EditFail2Ban**](docs/Fail2BanApi.md#editfail2ban) | **Post** /api/v1/edit/fail2ban | Edit Fail2Ban
*Fail2BanApi* | [**GetFail2BanConfig**](docs/Fail2BanApi.md#getfail2banconfig) | **Get** /api/v1/get/fail2ban | Get Fail2Ban Config
*FordwardingHostsApi* | [**AddForwardHost**](docs/FordwardingHostsApi.md#addforwardhost) | **Post** /api/v1/add/fwdhost | Add Forward Host
*FordwardingHostsApi* | [**DeleteForwardHost**](docs/FordwardingHostsApi.md#deleteforwardhost) | **Post** /api/v1/delete/fwdhost | Delete Forward Host
*FordwardingHostsApi* | [**GetForwardingHosts**](docs/FordwardingHostsApi.md#getforwardinghosts) | **Get** /api/v1/get/fwdhost/all | Get Forwarding Hosts
*LogsApi* | [**GetACMELogs**](docs/LogsApi.md#getacmelogs) | **Get** /api/v1/get/logs/acme/{count} | Get ACME logs
*LogsApi* | [**GetApiLogs**](docs/LogsApi.md#getapilogs) | **Get** /api/v1/get/logs/api/{count} | Get Api logs
*LogsApi* | [**GetAutodiscoverLogs**](docs/LogsApi.md#getautodiscoverlogs) | **Get** /api/v1/get/logs/autodiscover/{count} | Get Autodiscover logs
*LogsApi* | [**GetDovecotLogs**](docs/LogsApi.md#getdovecotlogs) | **Get** /api/v1/get/logs/dovecot/{count} | Get Dovecot logs
*LogsApi* | [**GetNetfilterLogs**](docs/LogsApi.md#getnetfilterlogs) | **Get** /api/v1/get/logs/netfilter/{count} | Get Netfilter logs
*LogsApi* | [**GetPostfixLogs**](docs/LogsApi.md#getpostfixlogs) | **Get** /api/v1/get/logs/postfix/{count} | Get Postfix logs
*LogsApi* | [**GetRatelimitLogs**](docs/LogsApi.md#getratelimitlogs) | **Get** /api/v1/get/logs/ratelimited/{count} | Get Ratelimit logs
*LogsApi* | [**GetRspamdLogs**](docs/LogsApi.md#getrspamdlogs) | **Get** /api/v1/get/logs/rspamd-history/{count} | Get Rspamd logs
*LogsApi* | [**GetSOGoLogs**](docs/LogsApi.md#getsogologs) | **Get** /api/v1/get/logs/sogo/{count} | Get SOGo logs
*LogsApi* | [**GetWatchdogLogs**](docs/LogsApi.md#getwatchdoglogs) | **Get** /api/v1/get/logs/watchdog/{count} | Get Watchdog logs
*MailboxesApi* | [**CreateMailbox**](docs/MailboxesApi.md#createmailbox) | **Post** /api/v1/add/mailbox | Create mailbox
*MailboxesApi* | [**DeleteMailbox**](docs/MailboxesApi.md#deletemailbox) | **Post** /api/v1/delete/mailbox | Delete mailbox
*MailboxesApi* | [**DeleteMailboxTags**](docs/MailboxesApi.md#deletemailboxtags) | **Post** /api/v1/delete/mailbox/tag/{mailbox} | Delete mailbox tags
*MailboxesApi* | [**EditMailboxSpamFilterScore**](docs/MailboxesApi.md#editmailboxspamfilterscore) | **Post** /api/v1/edit/spam-score/ | Edit mailbox spam filter score
*MailboxesApi* | [**GetMailboxes**](docs/MailboxesApi.md#getmailboxes) | **Get** /api/v1/get/mailbox/{id} | Get mailboxes
*MailboxesApi* | [**QuarantineNotifications**](docs/MailboxesApi.md#quarantinenotifications) | **Post** /api/v1/edit/quarantine_notification | Quarantine Notifications
*MailboxesApi* | [**UpdateMailbox**](docs/MailboxesApi.md#updatemailbox) | **Post** /api/v1/edit/mailbox | Update mailbox
*MailboxesApi* | [**UpdateMailboxACL**](docs/MailboxesApi.md#updatemailboxacl) | **Post** /api/v1/edit/user-acl | Update mailbox ACL
*MailboxesApi* | [**UpdatePushoverSettings**](docs/MailboxesApi.md#updatepushoversettings) | **Post** /api/v1/edit/pushover | Update Pushover settings
*OAuthClientsApi* | [**CreateOAuthClient**](docs/OAuthClientsApi.md#createoauthclient) | **Post** /api/v1/add/oauth2-client | Create oAuth Client
*OAuthClientsApi* | [**DeleteOAuthClient**](docs/OAuthClientsApi.md#deleteoauthclient) | **Post** /api/v1/delete/oauth2-client | Delete oAuth Client
*OAuthClientsApi* | [**GetOAuthClients**](docs/OAuthClientsApi.md#getoauthclients) | **Get** /api/v1/get/oauth2-client/{id} | Get oAuth Clients
*OutgoingTLSPolicyMapOverridesApi* | [**CreateTLSPolicyMap**](docs/OutgoingTLSPolicyMapOverridesApi.md#createtlspolicymap) | **Post** /api/v1/add/tls-policy-map | Create TLS Policy Map
*OutgoingTLSPolicyMapOverridesApi* | [**DeleteTLSPolicyMap**](docs/OutgoingTLSPolicyMapOverridesApi.md#deletetlspolicymap) | **Post** /api/v1/delete/tls-policy-map | Delete TLS Policy Map
*OutgoingTLSPolicyMapOverridesApi* | [**GetTLSPolicyMap**](docs/OutgoingTLSPolicyMapOverridesApi.md#gettlspolicymap) | **Get** /api/v1/get/tls-policy-map/{id} | Get TLS Policy Map
*QuarantineApi* | [**DeleteMailsInQuarantine**](docs/QuarantineApi.md#deletemailsinquarantine) | **Post** /api/v1/delete/qitem | Delete mails in Quarantine
*QuarantineApi* | [**GetMailsInQuarantine**](docs/QuarantineApi.md#getmailsinquarantine) | **Get** /api/v1/get/quarantine/all | Get mails in Quarantine
*QueueManagerApi* | [**DeleteQueue**](docs/QueueManagerApi.md#deletequeue) | **Post** /api/v1/delete/mailq | Delete Queue
*QueueManagerApi* | [**FlushQueue**](docs/QueueManagerApi.md#flushqueue) | **Post** /api/v1/edit/mailq | Flush Queue
*QueueManagerApi* | [**GetQueue**](docs/QueueManagerApi.md#getqueue) | **Get** /api/v1/get/mailq/all | Get Queue
*RatelimitsApi* | [**EditDomainRatelimits**](docs/RatelimitsApi.md#editdomainratelimits) | **Post** /api/v1/edit/rl-domain/ | Edit domain ratelimits
*RatelimitsApi* | [**EditMailboxRatelimits**](docs/RatelimitsApi.md#editmailboxratelimits) | **Post** /api/v1/edit/rl-mbox/ | Edit mailbox ratelimits
*RatelimitsApi* | [**GetDomainRatelimits**](docs/RatelimitsApi.md#getdomainratelimits) | **Get** /api/v1/get/rl-domain/{domain} | Get domain ratelimits
*RatelimitsApi* | [**GetMailboxRatelimits**](docs/RatelimitsApi.md#getmailboxratelimits) | **Get** /api/v1/get/rl-mbox/{mailbox} | Get mailbox ratelimits
*ResourcesApi* | [**CreateResources**](docs/ResourcesApi.md#createresources) | **Post** /api/v1/add/resource | Create Resources
*ResourcesApi* | [**DeleteResources**](docs/ResourcesApi.md#deleteresources) | **Post** /api/v1/delete/resource | Delete Resources
*ResourcesApi* | [**GetResources**](docs/ResourcesApi.md#getresources) | **Get** /api/v1/get/resource/all | Get Resources
*RoutingApi* | [**CreateSenderDependentTransports**](docs/RoutingApi.md#createsenderdependenttransports) | **Post** /api/v1/add/relayhost | Create Sender-Dependent Transports
*RoutingApi* | [**CreateTransportMaps**](docs/RoutingApi.md#createtransportmaps) | **Post** /api/v1/add/transport | Create Transport Maps
*RoutingApi* | [**DeleteSenderDependentTransports**](docs/RoutingApi.md#deletesenderdependenttransports) | **Post** /api/v1/delete/relayhost | Delete Sender-Dependent Transports
*RoutingApi* | [**DeleteTransportMaps**](docs/RoutingApi.md#deletetransportmaps) | **Post** /api/v1/delete/transport | Delete Transport Maps
*RoutingApi* | [**GetSenderDependentTransports**](docs/RoutingApi.md#getsenderdependenttransports) | **Get** /api/v1/get/relayhost/{id} | Get Sender-Dependent Transports
*RoutingApi* | [**GetTransportMaps**](docs/RoutingApi.md#gettransportmaps) | **Get** /api/v1/get/transport/{id} | Get Transport Maps
*StatusApi* | [**GetContainerStatus**](docs/StatusApi.md#getcontainerstatus) | **Get** /api/v1/get/status/containers | Get container status
*StatusApi* | [**GetSolrStatus**](docs/StatusApi.md#getsolrstatus) | **Get** /api/v1/get/status/solr | Get solr status
*StatusApi* | [**GetVersionStatus**](docs/StatusApi.md#getversionstatus) | **Get** /api/v1/get/status/version | Get version status
*StatusApi* | [**GetVmailStatus**](docs/StatusApi.md#getvmailstatus) | **Get** /api/v1/get/status/vmail | Get vmail status
*SyncJobsApi* | [**CreateSyncJob**](docs/SyncJobsApi.md#createsyncjob) | **Post** /api/v1/add/syncjob | Create sync job
*SyncJobsApi* | [**DeleteSyncJob**](docs/SyncJobsApi.md#deletesyncjob) | **Post** /api/v1/delete/syncjob | Delete sync job
*SyncJobsApi* | [**GetSyncJobs**](docs/SyncJobsApi.md#getsyncjobs) | **Get** /api/v1/get/syncjobs/all/no_log | Get sync jobs
*SyncJobsApi* | [**UpdateSyncJob**](docs/SyncJobsApi.md#updatesyncjob) | **Post** /api/v1/edit/syncjob | Update sync job


## Documentation For Models

 - [AddForwardHostRequest](docs/AddForwardHostRequest.md)
 - [CreateAlias200Response](docs/CreateAlias200Response.md)
 - [CreateAlias401Response](docs/CreateAlias401Response.md)
 - [CreateAliasRequest](docs/CreateAliasRequest.md)
 - [CreateAppPasswordRequest](docs/CreateAppPasswordRequest.md)
 - [CreateBCCMapRequest](docs/CreateBCCMapRequest.md)
 - [CreateDomainAdminUserRequest](docs/CreateDomainAdminUserRequest.md)
 - [CreateDomainPolicyRequest](docs/CreateDomainPolicyRequest.md)
 - [CreateDomainRequest](docs/CreateDomainRequest.md)
 - [CreateMailboxRequest](docs/CreateMailboxRequest.md)
 - [CreateOAuthClientRequest](docs/CreateOAuthClientRequest.md)
 - [CreateRecipientMapRequest](docs/CreateRecipientMapRequest.md)
 - [CreateResourcesRequest](docs/CreateResourcesRequest.md)
 - [CreateSenderDependentTransportsRequest](docs/CreateSenderDependentTransportsRequest.md)
 - [CreateSyncJobRequest](docs/CreateSyncJobRequest.md)
 - [CreateTLSPolicyMapRequest](docs/CreateTLSPolicyMapRequest.md)
 - [CreateTimeLimitedAliasRequest](docs/CreateTimeLimitedAliasRequest.md)
 - [CreateTransportMapsRequest](docs/CreateTransportMapsRequest.md)
 - [DeleteAppPasswordRequest](docs/DeleteAppPasswordRequest.md)
 - [DeleteBCCMapRequest](docs/DeleteBCCMapRequest.md)
 - [DeleteDomainAdminRequest](docs/DeleteDomainAdminRequest.md)
 - [DeleteDomainPolicyRequest](docs/DeleteDomainPolicyRequest.md)
 - [DeleteDomainRequest](docs/DeleteDomainRequest.md)
 - [DeleteDomainTagsRequest](docs/DeleteDomainTagsRequest.md)
 - [DeleteForwardHostRequest](docs/DeleteForwardHostRequest.md)
 - [DeleteMailboxRequest](docs/DeleteMailboxRequest.md)
 - [DeleteMailboxTagsRequest](docs/DeleteMailboxTagsRequest.md)
 - [DeleteMailsInQuarantineRequest](docs/DeleteMailsInQuarantineRequest.md)
 - [DeleteOAuthClientRequest](docs/DeleteOAuthClientRequest.md)
 - [DeleteQueueRequest](docs/DeleteQueueRequest.md)
 - [DeleteRecipientMapRequest](docs/DeleteRecipientMapRequest.md)
 - [DeleteResourcesRequest](docs/DeleteResourcesRequest.md)
 - [DeleteSenderDependentTransportsRequest](docs/DeleteSenderDependentTransportsRequest.md)
 - [DeleteSyncJobRequest](docs/DeleteSyncJobRequest.md)
 - [DeleteTLSPolicyMapRequest](docs/DeleteTLSPolicyMapRequest.md)
 - [DeleteTransportMapsRequest](docs/DeleteTransportMapsRequest.md)
 - [DuplicateDKIMKeyRequest](docs/DuplicateDKIMKeyRequest.md)
 - [EditDomainAdminACLRequest](docs/EditDomainAdminACLRequest.md)
 - [EditDomainAdminACLRequestAttr](docs/EditDomainAdminACLRequestAttr.md)
 - [EditDomainAdminUser200Response](docs/EditDomainAdminUser200Response.md)
 - [EditDomainAdminUserRequest](docs/EditDomainAdminUserRequest.md)
 - [EditDomainAdminUserRequestAttr](docs/EditDomainAdminUserRequestAttr.md)
 - [EditDomainRatelimitsRequest](docs/EditDomainRatelimitsRequest.md)
 - [EditFail2BanRequest](docs/EditFail2BanRequest.md)
 - [EditFail2BanRequestAttr](docs/EditFail2BanRequestAttr.md)
 - [EditMailboxRatelimitsRequest](docs/EditMailboxRatelimitsRequest.md)
 - [EditMailboxRatelimitsRequestAttr](docs/EditMailboxRatelimitsRequestAttr.md)
 - [FlushQueueRequest](docs/FlushQueueRequest.md)
 - [GenerateDKIMKeyRequest](docs/GenerateDKIMKeyRequest.md)
 - [QuarantineNotificationsRequest](docs/QuarantineNotificationsRequest.md)
 - [QuarantineNotificationsRequestAttr](docs/QuarantineNotificationsRequestAttr.md)
 - [UpdateAliasRequest](docs/UpdateAliasRequest.md)
 - [UpdateAliasRequestAttr](docs/UpdateAliasRequestAttr.md)
 - [UpdateDomainRequest](docs/UpdateDomainRequest.md)
 - [UpdateDomainRequestAttr](docs/UpdateDomainRequestAttr.md)
 - [UpdateMailboxACLRequest](docs/UpdateMailboxACLRequest.md)
 - [UpdateMailboxACLRequestAttr](docs/UpdateMailboxACLRequestAttr.md)
 - [UpdateMailboxRequest](docs/UpdateMailboxRequest.md)
 - [UpdateMailboxRequestAttr](docs/UpdateMailboxRequestAttr.md)
 - [UpdatePushoverSettingsRequest](docs/UpdatePushoverSettingsRequest.md)
 - [UpdatePushoverSettingsRequestAttr](docs/UpdatePushoverSettingsRequestAttr.md)
 - [UpdateSyncJobRequest](docs/UpdateSyncJobRequest.md)
 - [UpdateSyncJobRequestAttr](docs/UpdateSyncJobRequestAttr.md)


## Documentation For Authorization



### ApiKeyAuth

- **Type**: API key
- **API key parameter name**: X-API-Key
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: X-API-Key and passed in as the auth context for each request.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`
