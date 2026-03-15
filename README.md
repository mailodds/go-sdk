# MailOdds Go SDK

The official Go SDK for high-performance email validation, DMARC monitoring, and deliverability infrastructure through the MailOdds API.

[![Go Reference](https://pkg.go.dev/badge/github.com/mailodds/go-sdk.svg)](https://pkg.go.dev/github.com/mailodds/go-sdk)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)
[![API Docs](https://img.shields.io/badge/docs-api--reference-blue)](https://mailodds.com/api-reference)

## Installation

```sh
go get github.com/mailodds/go-sdk
```

## Quick Start

```go
package main

import (
    "context"
    "fmt"
    "os"

    mailodds "github.com/mailodds/go-sdk"
)

func main() {
    cfg := mailodds.NewConfiguration()
    client := mailodds.NewAPIClient(cfg)

    ctx := context.WithValue(context.Background(), mailodds.ContextAccessToken, os.Getenv("MAILODDS_API_KEY"))

    result, _, err := client.EmailValidationAPI.ValidateEmail(ctx).ValidateRequest(
        mailodds.ValidateRequest{Email: "user@example.com"},
    ).Execute()
    if err != nil {
        panic(err)
    }
    fmt.Printf("Status: %s, Action: %s\n", result.Result.GetStatus(), result.Result.GetAction())
}
```

## MailOdds Platform

This SDK is part of the [MailOdds](https://mailodds.com) email deliverability platform. Use it alongside the dashboard and other tools to validate, send, and monitor email at scale.

- [Email Validation API](https://mailodds.com/email-validation-api) - Single and batch email verification with 25+ real-time checks
- [Bulk Email Validation](https://mailodds.com/bulk-email-validation) - Process lists of up to 500,000 emails per job
- [Email Sending API](https://mailodds.com/email-sending-api) - Transactional email delivery with DKIM dual signing
- [Email Deliverability Platform](https://mailodds.com/email-deliverability-platform) - Full-stack deliverability monitoring and optimization
- [DMARC Monitoring](https://mailodds.com/dmarc-monitoring) - Aggregate report analysis with policy recommendations
- [Sender Reputation](https://mailodds.com/sender-reputation) - Real-time sender health scoring and trend analysis
- [SMTP Server Test](https://mailodds.com/smtp-server-test) - DNS, MX, and SMTP connectivity diagnostics
- [API Reference](https://mailodds.com/api-reference) - Full endpoint documentation with request and response examples
- [Guide: Email Authentication](https://mailodds.com/guides/email-authentication) - SPF, DKIM, and DMARC setup guide
- [Security](https://mailodds.com/security) - Infrastructure security and data protection practices

## Features

- **Context-based auth** - Pass authentication and cancellation through standard Go contexts
- **Pointer helpers** - Utility functions (`PtrString`, `PtrInt64`, `PtrBool`, etc.) for building optional request fields
- **Server URL configuration** - Override the base URL per client or per operation using context values
- **Concurrent-safe** - Thread-safe client suitable for use across multiple goroutines
- **Typed error handling** - Structured API error responses with machine-readable codes
- **Full DMARC support** - Add domains, verify DNS records, analyze sources, and retrieve trend data

## Why MailOdds

MailOdds is a complete email deliverability platform built for developers. Every email validated or sent through MailOdds passes through 25+ real-time checks including syntax verification, DNS and MX validation, SMTP mailbox probing, disposable domain detection, and role account identification.

The platform maintains sub-200ms median response times for single validations, 99.9% API uptime, and processes bulk lists of up to 500,000 emails per job. MailOdds supports 11 language SDKs, an MCP server for AI agent integration, a CLI for local development, and a WordPress plugin for no-code deployments.

All email sending uses DKIM dual signing with automated key rotation, and the deliverability monitoring stack covers DMARC aggregate reports, blacklist surveillance across 80+ DNSBLs, and real-time sender health scoring.

## API Reference

Full documentation is available at the [MailOdds API Reference](https://mailodds.com/api-reference).

All URIs are relative to `https://api.mailodds.com/v1`.

<details>
<summary>All Endpoints</summary>

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AgentControlPlaneAPI* | [**GetMcpCapabilities**](docs/AgentControlPlaneAPI.md#getmcpcapabilities) | **Get** /v1/mcp/capabilities | Get MCP capabilities
*BlacklistMonitoringAPI* | [**AddBlacklistMonitor**](docs/BlacklistMonitoringAPI.md#addblacklistmonitor) | **Post** /v1/blacklist-monitors | Add blacklist monitor
*BlacklistMonitoringAPI* | [**DeleteBlacklistMonitor**](docs/BlacklistMonitoringAPI.md#deleteblacklistmonitor) | **Delete** /v1/blacklist-monitors/{monitor_id} | Delete a blacklist monitor
*BlacklistMonitoringAPI* | [**GetBlacklistHistory**](docs/BlacklistMonitoringAPI.md#getblacklisthistory) | **Get** /v1/blacklist-monitors/{monitor_id}/history | Get blacklist check history
*BlacklistMonitoringAPI* | [**ListBlacklistMonitors**](docs/BlacklistMonitoringAPI.md#listblacklistmonitors) | **Get** /v1/blacklist-monitors | List blacklist monitors
*BlacklistMonitoringAPI* | [**RunBlacklistCheck**](docs/BlacklistMonitoringAPI.md#runblacklistcheck) | **Post** /v1/blacklist-monitors/{monitor_id}/check | Run blacklist check
*BounceAnalysisAPI* | [**CreateBounceAnalysis**](docs/BounceAnalysisAPI.md#createbounceanalysis) | **Post** /v1/bounce-analyses | Analyze bounce logs
*BounceAnalysisAPI* | [**CrossReferenceBounces**](docs/BounceAnalysisAPI.md#crossreferencebounces) | **Get** /v1/bounce-analyses/{analysis_id}/cross-reference | Cross-reference bounces with validation logs
*BounceAnalysisAPI* | [**GetBounceAnalysis**](docs/BounceAnalysisAPI.md#getbounceanalysis) | **Get** /v1/bounce-analyses/{analysis_id} | Get bounce analysis
*BounceAnalysisAPI* | [**GetBounceRecords**](docs/BounceAnalysisAPI.md#getbouncerecords) | **Get** /v1/bounce-analyses/{analysis_id}/records | Get bounce records
*BulkValidationAPI* | [**CancelJob**](docs/BulkValidationAPI.md#canceljob) | **Post** /v1/jobs/{job_id}/cancel | Cancel a job
*BulkValidationAPI* | [**CreateJob**](docs/BulkValidationAPI.md#createjob) | **Post** /v1/jobs | Create bulk validation job (JSON)
*BulkValidationAPI* | [**CreateJobFromS3**](docs/BulkValidationAPI.md#createjobfroms3) | **Post** /v1/jobs/upload/s3 | Create job from S3 upload
*BulkValidationAPI* | [**CreateJobUpload**](docs/BulkValidationAPI.md#createjobupload) | **Post** /v1/jobs/upload | Create bulk validation job (file upload)
*BulkValidationAPI* | [**DeleteJob**](docs/BulkValidationAPI.md#deletejob) | **Delete** /v1/jobs/{job_id} | Delete a job
*BulkValidationAPI* | [**GetJob**](docs/BulkValidationAPI.md#getjob) | **Get** /v1/jobs/{job_id} | Get job status
*BulkValidationAPI* | [**GetJobResults**](docs/BulkValidationAPI.md#getjobresults) | **Get** /v1/jobs/{job_id}/results | Get job results
*BulkValidationAPI* | [**GetPresignedUpload**](docs/BulkValidationAPI.md#getpresignedupload) | **Post** /v1/jobs/upload/presigned | Get S3 presigned upload URL
*BulkValidationAPI* | [**ListJobs**](docs/BulkValidationAPI.md#listjobs) | **Get** /v1/jobs | List validation jobs
*CampaignAnalyticsAPI* | [**GetCampaignABResults**](docs/CampaignAnalyticsAPI.md#getcampaignabresults) | **Get** /v1/campaigns/{campaign_id}/ab-results | Get A/B test results
*CampaignAnalyticsAPI* | [**GetCampaignAttribution**](docs/CampaignAnalyticsAPI.md#getcampaignattribution) | **Get** /v1/campaigns/{campaign_id}/conversions/attribution | Get campaign attribution
*CampaignAnalyticsAPI* | [**GetCampaignDeliveryConfidence**](docs/CampaignAnalyticsAPI.md#getcampaigndeliveryconfidence) | **Get** /v1/campaigns/{campaign_id}/delivery-confidence | Get pre-send delivery confidence
*CampaignAnalyticsAPI* | [**GetCampaignFunnel**](docs/CampaignAnalyticsAPI.md#getcampaignfunnel) | **Get** /v1/campaigns/{campaign_id}/funnel | Get campaign funnel
*CampaignAnalyticsAPI* | [**GetCampaignProviderIntelligence**](docs/CampaignAnalyticsAPI.md#getcampaignproviderintelligence) | **Get** /v1/campaigns/{campaign_id}/provider-intelligence | Get provider intelligence
*CampaignsAPI* | [**CancelCampaign**](docs/CampaignsAPI.md#cancelcampaign) | **Post** /v1/campaigns/{campaign_id}/cancel | Cancel a campaign
*CampaignsAPI* | [**CreateCampaign**](docs/CampaignsAPI.md#createcampaign) | **Post** /v1/campaigns | Create a campaign
*CampaignsAPI* | [**CreateCampaignVariant**](docs/CampaignsAPI.md#createcampaignvariant) | **Post** /v1/campaigns/{campaign_id}/variants | Create A/B variant
*CampaignsAPI* | [**GetCampaign**](docs/CampaignsAPI.md#getcampaign) | **Get** /v1/campaigns/{campaign_id} | Get campaign with stats
*CampaignsAPI* | [**ListCampaigns**](docs/CampaignsAPI.md#listcampaigns) | **Get** /v1/campaigns | List campaigns
*CampaignsAPI* | [**ScheduleCampaign**](docs/CampaignsAPI.md#schedulecampaign) | **Post** /v1/campaigns/{campaign_id}/schedule | Schedule a campaign
*CampaignsAPI* | [**SendCampaign**](docs/CampaignsAPI.md#sendcampaign) | **Post** /v1/campaigns/{campaign_id}/send | Send a campaign
*ContactListsAPI* | [**AppendToContactList**](docs/ContactListsAPI.md#appendtocontactlist) | **Post** /v1/contact-lists/{list_id}/append | Append to contact list
*ContactListsAPI* | [**CreateContactList**](docs/ContactListsAPI.md#createcontactlist) | **Post** /v1/contact-lists | Create contact list
*ContactListsAPI* | [**DeleteContactList**](docs/ContactListsAPI.md#deletecontactlist) | **Delete** /v1/contact-lists/{list_id} | Delete a contact list
*ContactListsAPI* | [**GetInactiveContactsReport**](docs/ContactListsAPI.md#getinactivecontactsreport) | **Get** /v1/contacts/inactive-report | Get inactive contacts report
*ContactListsAPI* | [**ListContactLists**](docs/ContactListsAPI.md#listcontactlists) | **Get** /v1/contact-lists | List contact lists
*ContactListsAPI* | [**QueryContactList**](docs/ContactListsAPI.md#querycontactlist) | **Post** /v1/contact-lists/{list_id}/query | Query contact list
*ContentClassificationAPI* | [**ClassifyContent**](docs/ContentClassificationAPI.md#classifycontent) | **Post** /v1/content-check | Classify email content
*DMARCMonitoringAPI* | [**AddDmarcDomain**](docs/DMARCMonitoringAPI.md#adddmarcdomain) | **Post** /v1/dmarc-domains | Add DMARC domain
*DMARCMonitoringAPI* | [**DeleteDmarcDomain**](docs/DMARCMonitoringAPI.md#deletedmarcdomain) | **Delete** /v1/dmarc-domains/{domain_id} | Delete a DMARC domain
*DMARCMonitoringAPI* | [**GetDmarcDomain**](docs/DMARCMonitoringAPI.md#getdmarcdomain) | **Get** /v1/dmarc-domains/{domain_id} | Get DMARC domain
*DMARCMonitoringAPI* | [**GetDmarcRecommendation**](docs/DMARCMonitoringAPI.md#getdmarcrecommendation) | **Get** /v1/dmarc-domains/{domain_id}/recommendation | Get DMARC policy recommendation
*DMARCMonitoringAPI* | [**GetDmarcSources**](docs/DMARCMonitoringAPI.md#getdmarcsources) | **Get** /v1/dmarc-domains/{domain_id}/sources | Get DMARC sending sources
*DMARCMonitoringAPI* | [**GetDmarcTrend**](docs/DMARCMonitoringAPI.md#getdmarctrend) | **Get** /v1/dmarc-domains/{domain_id}/trend | Get DMARC trend
*DMARCMonitoringAPI* | [**ListDmarcDomains**](docs/DMARCMonitoringAPI.md#listdmarcdomains) | **Get** /v1/dmarc-domains | List DMARC domains
*DMARCMonitoringAPI* | [**VerifyDmarcDomain**](docs/DMARCMonitoringAPI.md#verifydmarcdomain) | **Post** /v1/dmarc-domains/{domain_id}/verify | Verify DMARC DNS records
*EmailSendingAPI* | [**DeliverBatch**](docs/EmailSendingAPI.md#deliverbatch) | **Post** /v1/deliver/batch | Send to multiple recipients (max 100)
*EmailSendingAPI* | [**DeliverEmail**](docs/EmailSendingAPI.md#deliveremail) | **Post** /v1/deliver | Send a single email
*EmailValidationAPI* | [**ValidateBatch**](docs/EmailValidationAPI.md#validatebatch) | **Post** /v1/validate/batch | Validate multiple emails (sync)
*EmailValidationAPI* | [**ValidateEmail**](docs/EmailValidationAPI.md#validateemail) | **Post** /v1/validate | Validate single email
*EventsAPI* | [**TrackEvent**](docs/EventsAPI.md#trackevent) | **Post** /v1/events/track | Track a commerce event
*MessageEventsAPI* | [**GetMessageEvents**](docs/MessageEventsAPI.md#getmessageevents) | **Get** /v1/message-events | Get message events
*OAuth20API* | [**CreateToken**](docs/OAuth20API.md#createtoken) | **Post** /oauth/token | Create token
*OAuth20API* | [**GetJwks**](docs/OAuth20API.md#getjwks) | **Get** /.well-known/jwks.json | Get JSON Web Key Set
*OAuth20API* | [**IntrospectToken**](docs/OAuth20API.md#introspecttoken) | **Post** /oauth/introspect | Introspect token
*OAuth20API* | [**OauthServerMetadata**](docs/OAuth20API.md#oauthservermetadata) | **Get** /.well-known/oauth-authorization-server | OAuth server metadata
*OAuth20API* | [**RevokeToken**](docs/OAuth20API.md#revoketoken) | **Post** /oauth/revoke | Revoke token
*ProductsAPI* | [**BatchProducts**](docs/ProductsAPI.md#batchproducts) | **Post** /v1/stores/{store_id}/products/batch | Batch push products
*ProductsAPI* | [**GetProduct**](docs/ProductsAPI.md#getproduct) | **Get** /v1/store-products/{product_id} | Get a product
*ProductsAPI* | [**QueryProducts**](docs/ProductsAPI.md#queryproducts) | **Get** /v1/store-products | Query products
*SenderHealthAPI* | [**GetSenderHealth**](docs/SenderHealthAPI.md#getsenderhealth) | **Get** /v1/sender-health | Get sender health score
*SenderHealthAPI* | [**GetSenderHealthTrend**](docs/SenderHealthAPI.md#getsenderhealthtrend) | **Get** /v1/sender-health/trend | Get sender health trend
*SendingDomainsAPI* | [**CreateSendingDomain**](docs/SendingDomainsAPI.md#createsendingdomain) | **Post** /v1/sending-domains | Add a sending domain
*SendingDomainsAPI* | [**DeleteSendingDomain**](docs/SendingDomainsAPI.md#deletesendingdomain) | **Delete** /v1/sending-domains/{domain_id} | Delete a sending domain
*SendingDomainsAPI* | [**GetSendingDomain**](docs/SendingDomainsAPI.md#getsendingdomain) | **Get** /v1/sending-domains/{domain_id} | Get a sending domain
*SendingDomainsAPI* | [**GetSendingDomainIdentityScore**](docs/SendingDomainsAPI.md#getsendingdomainidentityscore) | **Get** /v1/sending-domains/{domain_id}/identity-score | Get domain identity score
*SendingDomainsAPI* | [**GetSendingStats**](docs/SendingDomainsAPI.md#getsendingstats) | **Get** /v1/sending-stats | Get sending statistics
*SendingDomainsAPI* | [**ListSendingDomains**](docs/SendingDomainsAPI.md#listsendingdomains) | **Get** /v1/sending-domains | List sending domains
*SendingDomainsAPI* | [**VerifySendingDomain**](docs/SendingDomainsAPI.md#verifysendingdomain) | **Post** /v1/sending-domains/{domain_id}/verify | Verify domain DNS records
*ServerTestsAPI* | [**GetServerTest**](docs/ServerTestsAPI.md#getservertest) | **Get** /v1/server-tests/{test_id} | Get server test
*ServerTestsAPI* | [**ListServerTests**](docs/ServerTestsAPI.md#listservertests) | **Get** /v1/server-tests | List server tests
*ServerTestsAPI* | [**RunServerTest**](docs/ServerTestsAPI.md#runservertest) | **Post** /v1/server-tests | Run server test
*SpamChecksAPI* | [**GetSpamCheck**](docs/SpamChecksAPI.md#getspamcheck) | **Get** /v1/spam-checks/{check_id} | Get spam check
*SpamChecksAPI* | [**ListSpamChecks**](docs/SpamChecksAPI.md#listspamchecks) | **Get** /v1/spam-checks | List spam checks
*SpamChecksAPI* | [**RunSpamCheck**](docs/SpamChecksAPI.md#runspamcheck) | **Post** /v1/spam-checks | Run spam check
*StoreConnectionsAPI* | [**CreateStore**](docs/StoreConnectionsAPI.md#createstore) | **Post** /v1/stores | Create a store connection
*StoreConnectionsAPI* | [**DisconnectStore**](docs/StoreConnectionsAPI.md#disconnectstore) | **Delete** /v1/stores/{store_id} | Disconnect a store
*StoreConnectionsAPI* | [**GetStore**](docs/StoreConnectionsAPI.md#getstore) | **Get** /v1/stores/{store_id} | Get a store connection
*StoreConnectionsAPI* | [**ListStores**](docs/StoreConnectionsAPI.md#liststores) | **Get** /v1/stores | List store connections
*StoreConnectionsAPI* | [**TriggerSync**](docs/StoreConnectionsAPI.md#triggersync) | **Post** /v1/stores/{store_id}/sync | Trigger product sync
*SubscriberListsAPI* | [**ConfirmSubscription**](docs/SubscriberListsAPI.md#confirmsubscription) | **Get** /v1/confirm/{token} | Confirm subscription
*SubscriberListsAPI* | [**CreateList**](docs/SubscriberListsAPI.md#createlist) | **Post** /v1/lists | Create a subscriber list
*SubscriberListsAPI* | [**DeleteList**](docs/SubscriberListsAPI.md#deletelist) | **Delete** /v1/lists/{list_id} | Delete a subscriber list
*SubscriberListsAPI* | [**GetList**](docs/SubscriberListsAPI.md#getlist) | **Get** /v1/lists/{list_id} | Get a subscriber list
*SubscriberListsAPI* | [**GetLists**](docs/SubscriberListsAPI.md#getlists) | **Get** /v1/lists | List subscriber lists
*SubscriberListsAPI* | [**GetSubscribers**](docs/SubscriberListsAPI.md#getsubscribers) | **Get** /v1/lists/{list_id}/subscribers | List subscribers
*SubscriberListsAPI* | [**Subscribe**](docs/SubscriberListsAPI.md#subscribe) | **Post** /v1/subscribe/{list_id} | Subscribe to a list
*SubscriberListsAPI* | [**UnsubscribeSubscriber**](docs/SubscriberListsAPI.md#unsubscribesubscriber) | **Delete** /v1/lists/{list_id}/subscribers/{subscriber_id} | Unsubscribe a subscriber
*SuppressionListsAPI* | [**AddSuppression**](docs/SuppressionListsAPI.md#addsuppression) | **Post** /v1/suppression | Add suppression entries
*SuppressionListsAPI* | [**CheckSuppression**](docs/SuppressionListsAPI.md#checksuppression) | **Post** /v1/suppression/check | Check suppression status
*SuppressionListsAPI* | [**GetSuppressionAuditLog**](docs/SuppressionListsAPI.md#getsuppressionauditlog) | **Get** /v1/suppression/audit | Get suppression audit log
*SuppressionListsAPI* | [**GetSuppressionStats**](docs/SuppressionListsAPI.md#getsuppressionstats) | **Get** /v1/suppression/stats | Get suppression statistics
*SuppressionListsAPI* | [**ListSuppression**](docs/SuppressionListsAPI.md#listsuppression) | **Get** /v1/suppression | List suppression entries
*SuppressionListsAPI* | [**RemoveSuppression**](docs/SuppressionListsAPI.md#removesuppression) | **Delete** /v1/suppression | Remove suppression entries
*SystemAPI* | [**GetTelemetrySummary**](docs/SystemAPI.md#gettelemetrysummary) | **Get** /v1/telemetry/summary | Get validation telemetry
*SystemAPI* | [**HealthCheck**](docs/SystemAPI.md#healthcheck) | **Get** /health | Health check
*ValidationPoliciesAPI* | [**AddPolicyRule**](docs/ValidationPoliciesAPI.md#addpolicyrule) | **Post** /v1/policies/{policy_id}/rules | Add rule to policy
*ValidationPoliciesAPI* | [**CreatePolicy**](docs/ValidationPoliciesAPI.md#createpolicy) | **Post** /v1/policies | Create policy
*ValidationPoliciesAPI* | [**CreatePolicyFromPreset**](docs/ValidationPoliciesAPI.md#createpolicyfrompreset) | **Post** /v1/policies/from-preset | Create policy from preset
*ValidationPoliciesAPI* | [**DeletePolicy**](docs/ValidationPoliciesAPI.md#deletepolicy) | **Delete** /v1/policies/{policy_id} | Delete policy
*ValidationPoliciesAPI* | [**DeletePolicyRule**](docs/ValidationPoliciesAPI.md#deletepolicyrule) | **Delete** /v1/policies/{policy_id}/rules/{rule_id} | Delete rule
*ValidationPoliciesAPI* | [**GetPolicy**](docs/ValidationPoliciesAPI.md#getpolicy) | **Get** /v1/policies/{policy_id} | Get policy
*ValidationPoliciesAPI* | [**GetPolicyPresets**](docs/ValidationPoliciesAPI.md#getpolicypresets) | **Get** /v1/policies/presets | Get policy presets
*ValidationPoliciesAPI* | [**ListPolicies**](docs/ValidationPoliciesAPI.md#listpolicies) | **Get** /v1/policies | List policies
*ValidationPoliciesAPI* | [**TestPolicy**](docs/ValidationPoliciesAPI.md#testpolicy) | **Post** /v1/policies/test | Test policy evaluation
*ValidationPoliciesAPI* | [**UpdatePolicy**](docs/ValidationPoliciesAPI.md#updatepolicy) | **Put** /v1/policies/{policy_id} | Update policy

</details>

<details>
<summary>All Models</summary>

 - [AddBlacklistMonitor201Response](docs/AddBlacklistMonitor201Response.md)
 - [AddBlacklistMonitorRequest](docs/AddBlacklistMonitorRequest.md)
 - [AddDmarcDomain201Response](docs/AddDmarcDomain201Response.md)
 - [AddDmarcDomainRequest](docs/AddDmarcDomainRequest.md)
 - [AddPolicyRule201Response](docs/AddPolicyRule201Response.md)
 - [AddSuppressionRequest](docs/AddSuppressionRequest.md)
 - [AddSuppressionRequestEntriesInner](docs/AddSuppressionRequestEntriesInner.md)
 - [AddSuppressionResponse](docs/AddSuppressionResponse.md)
 - [AppendToContactList200Response](docs/AppendToContactList200Response.md)
 - [AppendToContactListRequest](docs/AppendToContactListRequest.md)
 - [BatchDeliverRequest](docs/BatchDeliverRequest.md)
 - [BatchDeliverRequestStructuredData](docs/BatchDeliverRequestStructuredData.md)
 - [BatchDeliverResponse](docs/BatchDeliverResponse.md)
 - [BatchDeliverResponseDelivery](docs/BatchDeliverResponseDelivery.md)
 - [BatchDeliverResponseRejectedInner](docs/BatchDeliverResponseRejectedInner.md)
 - [BatchProductsRequest](docs/BatchProductsRequest.md)
 - [BatchProductsRequestProductsInner](docs/BatchProductsRequestProductsInner.md)
 - [BatchProductsResponse](docs/BatchProductsResponse.md)
 - [BatchProductsResponseErrorsInner](docs/BatchProductsResponseErrorsInner.md)
 - [BlacklistMonitor](docs/BlacklistMonitor.md)
 - [BlacklistMonitorLatestCheck](docs/BlacklistMonitorLatestCheck.md)
 - [BounceAnalysisResponse](docs/BounceAnalysisResponse.md)
 - [BounceAnalysisResponseAnalysis](docs/BounceAnalysisResponseAnalysis.md)
 - [BounceAnalysisResponseAnalysisCategories](docs/BounceAnalysisResponseAnalysisCategories.md)
 - [BounceAnalysisResponseAnalysisTopDomainsInner](docs/BounceAnalysisResponseAnalysisTopDomainsInner.md)
 - [Campaign](docs/Campaign.md)
 - [CampaignResponse](docs/CampaignResponse.md)
 - [CampaignStats](docs/CampaignStats.md)
 - [CampaignVariant](docs/CampaignVariant.md)
 - [CheckSuppressionRequest](docs/CheckSuppressionRequest.md)
 - [ClassifyContent200Response](docs/ClassifyContent200Response.md)
 - [ClassifyContent200ResponseContentCheck](docs/ClassifyContent200ResponseContentCheck.md)
 - [ClassifyContentRequest](docs/ClassifyContentRequest.md)
 - [ConfirmSubscription200Response](docs/ConfirmSubscription200Response.md)
 - [ContactList](docs/ContactList.md)
 - [CreateBounceAnalysisRequest](docs/CreateBounceAnalysisRequest.md)
 - [CreateCampaignRequest](docs/CreateCampaignRequest.md)
 - [CreateCampaignVariant201Response](docs/CreateCampaignVariant201Response.md)
 - [CreateContactList201Response](docs/CreateContactList201Response.md)
 - [CreateContactListRequest](docs/CreateContactListRequest.md)
 - [CreateJobFromS3Request](docs/CreateJobFromS3Request.md)
 - [CreateJobRequest](docs/CreateJobRequest.md)
 - [CreateList201Response](docs/CreateList201Response.md)
 - [CreateListRequest](docs/CreateListRequest.md)
 - [CreatePolicyFromPresetRequest](docs/CreatePolicyFromPresetRequest.md)
 - [CreatePolicyRequest](docs/CreatePolicyRequest.md)
 - [CreateSendingDomain201Response](docs/CreateSendingDomain201Response.md)
 - [CreateSendingDomainRequest](docs/CreateSendingDomainRequest.md)
 - [CreateStore201Response](docs/CreateStore201Response.md)
 - [CreateStoreRequest](docs/CreateStoreRequest.md)
 - [CreateToken200Response](docs/CreateToken200Response.md)
 - [CreateVariantRequest](docs/CreateVariantRequest.md)
 - [CrossReferenceBounces200Response](docs/CrossReferenceBounces200Response.md)
 - [CrossReferenceBounces200ResponseCrossReference](docs/CrossReferenceBounces200ResponseCrossReference.md)
 - [CrossReferenceBounces200ResponseCrossReferenceEntriesInner](docs/CrossReferenceBounces200ResponseCrossReferenceEntriesInner.md)
 - [DeleteJob200Response](docs/DeleteJob200Response.md)
 - [DeletePolicy200Response](docs/DeletePolicy200Response.md)
 - [DeletePolicyRule200Response](docs/DeletePolicyRule200Response.md)
 - [DeliverRequest](docs/DeliverRequest.md)
 - [DeliverRequestOptions](docs/DeliverRequestOptions.md)
 - [DeliverRequestStructuredData](docs/DeliverRequestStructuredData.md)
 - [DeliverRequestToInner](docs/DeliverRequestToInner.md)
 - [DeliverResponse](docs/DeliverResponse.md)
 - [DeliverResponseDelivery](docs/DeliverResponseDelivery.md)
 - [DisconnectStore200Response](docs/DisconnectStore200Response.md)
 - [DmarcDomain](docs/DmarcDomain.md)
 - [ErrorResponse](docs/ErrorResponse.md)
 - [GetBlacklistHistory200Response](docs/GetBlacklistHistory200Response.md)
 - [GetBlacklistHistory200ResponseChecksInner](docs/GetBlacklistHistory200ResponseChecksInner.md)
 - [GetBounceRecords200Response](docs/GetBounceRecords200Response.md)
 - [GetBounceRecords200ResponseRecordsInner](docs/GetBounceRecords200ResponseRecordsInner.md)
 - [GetCampaignABResults200Response](docs/GetCampaignABResults200Response.md)
 - [GetCampaignABResults200ResponseVariantsInner](docs/GetCampaignABResults200ResponseVariantsInner.md)
 - [GetCampaignABResults200ResponseWinner](docs/GetCampaignABResults200ResponseWinner.md)
 - [GetCampaignAttribution200Response](docs/GetCampaignAttribution200Response.md)
 - [GetCampaignAttribution200ResponseAttribution](docs/GetCampaignAttribution200ResponseAttribution.md)
 - [GetCampaignAttribution200ResponseAttributionFirstTouch](docs/GetCampaignAttribution200ResponseAttributionFirstTouch.md)
 - [GetCampaignDeliveryConfidence200Response](docs/GetCampaignDeliveryConfidence200Response.md)
 - [GetCampaignDeliveryConfidence200ResponseFactors](docs/GetCampaignDeliveryConfidence200ResponseFactors.md)
 - [GetCampaignDeliveryConfidence200ResponseFactorsDomainAuth](docs/GetCampaignDeliveryConfidence200ResponseFactorsDomainAuth.md)
 - [GetCampaignDeliveryConfidence200ResponseFactorsListQuality](docs/GetCampaignDeliveryConfidence200ResponseFactorsListQuality.md)
 - [GetCampaignDeliveryConfidence200ResponseFactorsSenderReputation](docs/GetCampaignDeliveryConfidence200ResponseFactorsSenderReputation.md)
 - [GetCampaignFunnel200Response](docs/GetCampaignFunnel200Response.md)
 - [GetCampaignFunnel200ResponseFunnel](docs/GetCampaignFunnel200ResponseFunnel.md)
 - [GetCampaignFunnel200ResponseRates](docs/GetCampaignFunnel200ResponseRates.md)
 - [GetCampaignProviderIntelligence200Response](docs/GetCampaignProviderIntelligence200Response.md)
 - [GetCampaignProviderIntelligence200ResponseProvidersInner](docs/GetCampaignProviderIntelligence200ResponseProvidersInner.md)
 - [GetDmarcDomain200Response](docs/GetDmarcDomain200Response.md)
 - [GetDmarcDomain200ResponseDomain](docs/GetDmarcDomain200ResponseDomain.md)
 - [GetDmarcDomain200ResponseDomainAllOfSummary](docs/GetDmarcDomain200ResponseDomainAllOfSummary.md)
 - [GetDmarcRecommendation200Response](docs/GetDmarcRecommendation200Response.md)
 - [GetDmarcRecommendation200ResponseRecommendation](docs/GetDmarcRecommendation200ResponseRecommendation.md)
 - [GetDmarcSources200Response](docs/GetDmarcSources200Response.md)
 - [GetDmarcSources200ResponseSourcesInner](docs/GetDmarcSources200ResponseSourcesInner.md)
 - [GetDmarcTrend200Response](docs/GetDmarcTrend200Response.md)
 - [GetDmarcTrend200ResponseTrendInner](docs/GetDmarcTrend200ResponseTrendInner.md)
 - [GetInactiveContactsReport200Response](docs/GetInactiveContactsReport200Response.md)
 - [GetInactiveContactsReport200ResponseByListInner](docs/GetInactiveContactsReport200ResponseByListInner.md)
 - [GetLists200Response](docs/GetLists200Response.md)
 - [GetMessageEvents200Response](docs/GetMessageEvents200Response.md)
 - [GetMessageEvents200ResponseClicksInner](docs/GetMessageEvents200ResponseClicksInner.md)
 - [GetMessageEvents200ResponseEventsInner](docs/GetMessageEvents200ResponseEventsInner.md)
 - [GetMessageEvents200ResponseSummary](docs/GetMessageEvents200ResponseSummary.md)
 - [GetPresignedUploadRequest](docs/GetPresignedUploadRequest.md)
 - [GetProduct200Response](docs/GetProduct200Response.md)
 - [GetSenderHealth200Response](docs/GetSenderHealth200Response.md)
 - [GetSenderHealth200ResponseComponents](docs/GetSenderHealth200ResponseComponents.md)
 - [GetSenderHealth200ResponseComponentsDeliveryRate](docs/GetSenderHealth200ResponseComponentsDeliveryRate.md)
 - [GetSenderHealth200ResponseVolume](docs/GetSenderHealth200ResponseVolume.md)
 - [GetSenderHealthTrend200Response](docs/GetSenderHealthTrend200Response.md)
 - [GetSenderHealthTrend200ResponseDataPointsInner](docs/GetSenderHealthTrend200ResponseDataPointsInner.md)
 - [GetSendingDomainIdentityScore200Response](docs/GetSendingDomainIdentityScore200Response.md)
 - [GetSendingStats200Response](docs/GetSendingStats200Response.md)
 - [GetSendingStats200ResponseStats](docs/GetSendingStats200ResponseStats.md)
 - [GetSubscribers200Response](docs/GetSubscribers200Response.md)
 - [HealthCheck200Response](docs/HealthCheck200Response.md)
 - [IdentityScoreCheck](docs/IdentityScoreCheck.md)
 - [IntrospectToken200Response](docs/IntrospectToken200Response.md)
 - [Job](docs/Job.md)
 - [JobArtifacts](docs/JobArtifacts.md)
 - [JobListResponse](docs/JobListResponse.md)
 - [JobResponse](docs/JobResponse.md)
 - [JobSummary](docs/JobSummary.md)
 - [JwksResponse](docs/JwksResponse.md)
 - [JwksResponseKeysInner](docs/JwksResponseKeysInner.md)
 - [ListBlacklistMonitors200Response](docs/ListBlacklistMonitors200Response.md)
 - [ListCampaigns200Response](docs/ListCampaigns200Response.md)
 - [ListContactLists200Response](docs/ListContactLists200Response.md)
 - [ListDmarcDomains200Response](docs/ListDmarcDomains200Response.md)
 - [ListSendingDomains200Response](docs/ListSendingDomains200Response.md)
 - [ListServerTests200Response](docs/ListServerTests200Response.md)
 - [ListSpamChecks200Response](docs/ListSpamChecks200Response.md)
 - [ListStores200Response](docs/ListStores200Response.md)
 - [McpCapabilities](docs/McpCapabilities.md)
 - [McpCapabilitiesPillarsInner](docs/McpCapabilitiesPillarsInner.md)
 - [McpCapabilitiesPillarsInnerToolsInner](docs/McpCapabilitiesPillarsInnerToolsInner.md)
 - [OAuthServerMetadata](docs/OAuthServerMetadata.md)
 - [Pagination](docs/Pagination.md)
 - [Policy](docs/Policy.md)
 - [PolicyListResponse](docs/PolicyListResponse.md)
 - [PolicyListResponseLimits](docs/PolicyListResponseLimits.md)
 - [PolicyPresetsResponse](docs/PolicyPresetsResponse.md)
 - [PolicyPresetsResponsePresetsInner](docs/PolicyPresetsResponsePresetsInner.md)
 - [PolicyResponse](docs/PolicyResponse.md)
 - [PolicyRule](docs/PolicyRule.md)
 - [PolicyRuleAction](docs/PolicyRuleAction.md)
 - [PolicyTestResponse](docs/PolicyTestResponse.md)
 - [PresignedUploadResponse](docs/PresignedUploadResponse.md)
 - [PresignedUploadResponseUpload](docs/PresignedUploadResponseUpload.md)
 - [ProductFacets](docs/ProductFacets.md)
 - [ProductFacetsCategoriesInner](docs/ProductFacetsCategoriesInner.md)
 - [ProductFacetsPriceRangesInner](docs/ProductFacetsPriceRangesInner.md)
 - [ProductFacetsStoresInner](docs/ProductFacetsStoresInner.md)
 - [QueryContactList200Response](docs/QueryContactList200Response.md)
 - [QueryContactList200ResponseEmailsInner](docs/QueryContactList200ResponseEmailsInner.md)
 - [QueryContactListRequest](docs/QueryContactListRequest.md)
 - [QueryContactListRequestFiltersInner](docs/QueryContactListRequestFiltersInner.md)
 - [QueryProducts200Response](docs/QueryProducts200Response.md)
 - [RemoveSuppression200Response](docs/RemoveSuppression200Response.md)
 - [RemoveSuppressionRequest](docs/RemoveSuppressionRequest.md)
 - [ResultsResponse](docs/ResultsResponse.md)
 - [RunBlacklistCheck200Response](docs/RunBlacklistCheck200Response.md)
 - [RunBlacklistCheck200ResponseCheck](docs/RunBlacklistCheck200ResponseCheck.md)
 - [RunServerTest201Response](docs/RunServerTest201Response.md)
 - [RunServerTestRequest](docs/RunServerTestRequest.md)
 - [RunSpamCheck201Response](docs/RunSpamCheck201Response.md)
 - [RunSpamCheckRequest](docs/RunSpamCheckRequest.md)
 - [ScheduleCampaignRequest](docs/ScheduleCampaignRequest.md)
 - [SendingDomain](docs/SendingDomain.md)
 - [SendingDomainDnsRecords](docs/SendingDomainDnsRecords.md)
 - [SendingDomainDnsRecordsNs](docs/SendingDomainDnsRecordsNs.md)
 - [SendingDomainIdentityScore](docs/SendingDomainIdentityScore.md)
 - [SendingDomainIdentityScoreBreakdown](docs/SendingDomainIdentityScoreBreakdown.md)
 - [ServerTest](docs/ServerTest.md)
 - [ServerTestDnsChecks](docs/ServerTestDnsChecks.md)
 - [ServerTestMxRecordsInner](docs/ServerTestMxRecordsInner.md)
 - [ServerTestSmtpCheck](docs/ServerTestSmtpCheck.md)
 - [SpamCheck](docs/SpamCheck.md)
 - [SpamCheckChecks](docs/SpamCheckChecks.md)
 - [StoreConnection](docs/StoreConnection.md)
 - [StoreProduct](docs/StoreProduct.md)
 - [SubscribeRequest](docs/SubscribeRequest.md)
 - [Subscriber](docs/Subscriber.md)
 - [SubscriberList](docs/SubscriberList.md)
 - [SuppressionAuditResponse](docs/SuppressionAuditResponse.md)
 - [SuppressionAuditResponseEntriesInner](docs/SuppressionAuditResponseEntriesInner.md)
 - [SuppressionCheckResponse](docs/SuppressionCheckResponse.md)
 - [SuppressionEntry](docs/SuppressionEntry.md)
 - [SuppressionListResponse](docs/SuppressionListResponse.md)
 - [SuppressionStatsResponse](docs/SuppressionStatsResponse.md)
 - [SuppressionStatsResponseByType](docs/SuppressionStatsResponseByType.md)
 - [SyncResponse](docs/SyncResponse.md)
 - [TelemetrySummary](docs/TelemetrySummary.md)
 - [TelemetrySummaryRates](docs/TelemetrySummaryRates.md)
 - [TelemetrySummaryTopDomainsInner](docs/TelemetrySummaryTopDomainsInner.md)
 - [TelemetrySummaryTopReasonsInner](docs/TelemetrySummaryTopReasonsInner.md)
 - [TelemetrySummaryTotals](docs/TelemetrySummaryTotals.md)
 - [TestPolicyRequest](docs/TestPolicyRequest.md)
 - [TestPolicyRequestTestResult](docs/TestPolicyRequestTestResult.md)
 - [TrackEventRequest](docs/TrackEventRequest.md)
 - [TrackEventResponse](docs/TrackEventResponse.md)
 - [UnsubscribeSubscriber200Response](docs/UnsubscribeSubscriber200Response.md)
 - [UpdatePolicyRequest](docs/UpdatePolicyRequest.md)
 - [ValidateBatch200Response](docs/ValidateBatch200Response.md)
 - [ValidateBatch200ResponseSummary](docs/ValidateBatch200ResponseSummary.md)
 - [ValidateBatchRequest](docs/ValidateBatchRequest.md)
 - [ValidateRequest](docs/ValidateRequest.md)
 - [ValidationResponse](docs/ValidationResponse.md)
 - [ValidationResponsePolicyApplied](docs/ValidationResponsePolicyApplied.md)
 - [ValidationResponseSuppressionMatch](docs/ValidationResponseSuppressionMatch.md)
 - [ValidationResult](docs/ValidationResult.md)
 - [ValidationResultSuppression](docs/ValidationResultSuppression.md)
 - [WebhookEvent](docs/WebhookEvent.md)

</details>

## Other SDKs

| Language | Package | Source |
|----------|---------|--------|
| [Python](https://mailodds.com/sdks) | [PyPI](https://pypi.org/project/mailodds/) | [GitHub](https://github.com/mailodds/python-sdk) |
| [TypeScript](https://mailodds.com/sdks) | [npm](https://www.npmjs.com/package/@mailodds/sdk) | [GitHub](https://github.com/mailodds/typescript-sdk) |
| [PHP](https://mailodds.com/sdks) | [Packagist](https://packagist.org/packages/mailodds/mailodds-php) | [GitHub](https://github.com/mailodds/php-sdk) |
| [Java](https://mailodds.com/sdks) | [GitHub](https://github.com/mailodds/java-sdk) | [GitHub](https://github.com/mailodds/java-sdk) |
| [Go](https://mailodds.com/sdks) | [pkg.go.dev](https://pkg.go.dev/github.com/mailodds/go-sdk) | [GitHub](https://github.com/mailodds/go-sdk) |
| [C# / .NET](https://mailodds.com/sdks) | [GitHub](https://github.com/mailodds/csharp-sdk) | [GitHub](https://github.com/mailodds/csharp-sdk) |
| [Ruby](https://mailodds.com/sdks) | [RubyGems](https://rubygems.org/gems/mailodds) | [GitHub](https://github.com/mailodds/ruby-sdk) |
| [Kotlin](https://mailodds.com/sdks) | [GitHub](https://github.com/mailodds/kotlin-sdk) | [GitHub](https://github.com/mailodds/kotlin-sdk) |
| [Rust](https://mailodds.com/sdks) | [crates.io](https://crates.io/crates/mailodds) | [GitHub](https://github.com/mailodds/rust-sdk) |
| [Swift](https://mailodds.com/sdks) | [GitHub](https://github.com/mailodds/swift-sdk) | [GitHub](https://github.com/mailodds/swift-sdk) |
| [Dart / Flutter](https://mailodds.com/sdks) | [pub.dev](https://pub.dev/packages/mailodds) | [GitHub](https://github.com/mailodds/dart-sdk) |

## Resources

- [Documentation](https://mailodds.com/docs)
- [Developer Quickstart](https://mailodds.com/developers)
- [All SDKs](https://mailodds.com/sdks)
- [Security](https://mailodds.com/security)
- [Guide: Email Authentication](https://mailodds.com/guides/email-authentication)

## License

MIT
