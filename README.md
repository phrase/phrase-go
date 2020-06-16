# Go API client for phrase

Phrase is a translation management platform for software projects. You can collaborate on language file translation with your team or order translations through our platform. The API allows you to import locale files, download locale files, tag keys or interact in other ways with the localization data stored in Phrase for your account.

## API Endpoint

```
https://api.phrase.com/v2/
```

The API is only accessible via HTTPS, the base URL is <code>https://api.phrase.com/</code>, and the current version is <code>v2</code> which results in the base URL for all requests: <code>https://api.phrase.com/v2/</code>.

## Usage

[curl](http://curl.haxx.se/) is used primarily to send requests to Phrase in the examples. On most you'll find a second variant using the [Phrase API v2 client](https://phrase.com/cli/) that might be more convenient to handle. For further information check its [documentation](https://help.phrase.com/help/phrase-in-your-terminal).

## Use of HTTP Verbs

Phrase API v2 tries to use the appropriate HTTP verb for accessing each endpoint according to REST specification where possible:

<div class=\"table-responsive\">
  <table class=\"basic-table\">
    <thead>
      <tr class=\"basic-table__row basic-table__row--header\">
        <th class=\"basic-table__cell basic-table__cell--header\">Verb</th>
        <th class=\"basic-table__cell basic-table__cell--header\">Description</th>
      </tr>
    </thead>
    <tbody>
      <tr>
        <td class=\"basic-table__cell\">GET</td>
        <td class=\"basic-table__cell\">Retrieve one or multiple resources</td>
      </tr>
      <tr>
        <td class=\"basic-table__cell\">POST</td>
        <td class=\"basic-table__cell\">Create a resource</td>
      </tr>
      <tr>
        <td class=\"basic-table__cell\">PUT</td>
        <td class=\"basic-table__cell\">Update a resource</td>
      </tr>
      <tr>
        <td class=\"basic-table__cell\">PATCH</td>
        <td class=\"basic-table__cell\">Update a resource (partially)</td>
      </tr>
      <tr>
        <td class=\"basic-table__cell\">DELETE</td>
        <td class=\"basic-table__cell\">Delete a resource</td>
      </tr>
    </tbody>
  </table>
</div>


## Identification via User-Agent

You must include the User-Agent header with the name of your application or project. It might be a good idea to include some sort of contact information  as well, so that we can get in touch if necessary (e.g. to warn you about Rate-Limiting or badly formed requests). Examples of excellent User-Agent headers:

```
User-Agent: Example Mobile App (example@phrase.com)
User-Agent: ACME Inc Python Client (http://example.com/contact)
```

If you don't send this header, you will receive a response with 400 Bad Request.


## Lists

When you request a list of resources, the API will typically only return an array of resources including their most important attributes. For a detailed representation of the resource you should request its detailed representation.

Lists are usually [paginated](#pagination).


## Parameters

Many endpoints support additional parameters, e.g. for pagination. When passing them in a GET request you can send them as HTTP query string parameters:

```
$ curl -u EMAIL_OR_ACCESS_TOKEN \"https://api.phrase.com/v2/projects?page=2\"
```

When performing a POST, PUT, PATCH or DELETE request, we recommend sending parameters that are not already included in the URL, as JSON body:

```
$ curl -H 'Content-Type: application/json' -d '{\"name\":\"My new project\"}' -u EMAIL_OR_ACCESS_TOKEN https://api.phrase.com/v2/projects
```

Encoding parameters as JSON means better support for types (boolean, integer) and usually better readability. Don't forget to set the correct Content-Type for your request.

*The Content-Type header is omitted in some of the following examples for better readbility.*


## Errors


### Request Errors

If a request contains invalid JSON or is missing a required parameter (besides resource attributes), the status `400 Bad Request` is returned:

```
{
  \"message\": \"JSON could not be parsed\"
}
```


### Validation Errors

When the validation for a resource fails, the status `422 Unprocessable Entity` is returned, along with information on the affected fields:

```
{
  \"message\": \"Validation Failed\",
  \"errors\": [
    {
      \"resource\": \"Project\",
      \"field\": \"name\",
      \"message\": \"can't be blank\"
    }
  ]
}
```


## Date Format

Times and dates are returned and expected in [ISO 8601](http://en.wikipedia.org/wiki/ISO_8601) date format:

```
YYYY-MM-DDTHH:MM:SSZ
```

Instead of 'Z' for UTC time zone you can specify your time zone's locale offset using the following notation:

```
YYYY-MM-DDTHH:MM:SS¬±hh:mm
```

Example for CET (1 hour behind UTC):

```
2015-03-31T13:00+01:00
```

Please note that in HTTP headers, we will use the appropriate recommended date formats instead of ISO 8601.


## Authentication

<div class=\"alert alert-info\">For more detailed information on authentication, check out the <a href=\"#authentication\">API v2 Authentication Guide</a>.</div>

There are two different ways to authenticate when performing API requests:

* E-Mail and password
* Oauth Access Token


### E-Mail and password

To get started easily, you can use HTTP Basic authentication with your email and password:

```
$ curl -u username:password \"https://api.phrase.com/v2/projects\"
```


### OAuth via Access Tokens

You can create and manage access tokens in your [profile settings](https://app.phrase.com/settings/oauth_access_tokens) in Translation Center or via the [Authorizations API](#authorizations).

Simply pass the access token as the username of your request:

```
$ curl -u ACCESS_TOKEN: \"https://api.phrase.com/v2/projects\"
```

or send the access token via the `Authorization` header field:

```
$ curl -H \"Authorization: token ACCESS_TOKEN\" https://api.phrase.com/v2/projects
```

For more detailed information on authentication, check out the <a href=\"#authentication\">API v2 Authentication Guide</a>.

#### Send via parameter

As JSONP (and other) requests cannot send HTTP Basic Auth credentials, a special query parameter `access_token` can be used:

```
curl \"https://api.phrase.com/v2/projects?access_token=ACCESS_TOKEN\"
```

You should only use this transport method if sending the authentication via header or Basic authentication is not possible.

### Two-Factor-Authentication

Users with Two-Factor-Authentication enabled have to send a valid token along their request with certain authentication methods (such as Basic authentication). The necessity of a Two-Factor-Authentication token is indicated by the `X-PhraseApp-OTP: required; :MFA-type` header in the response. The `:MFA-type`field indicates the source of the token, e.g. `app` (refers to your Authenticator application):

```
X-PhraseApp-OTP: required; app
```

To provide a Two-Factor-Authentication token you can simply send it in the header of the request:

```
curl -H \"X-PhraseApp-OTP: MFA-TOKEN\" -u EMAIL https://api.phrase.com/v2/projects
```

Since Two-Factor-Authentication tokens usually expire quickly, we recommend using an alternative authentication method such as OAuth access tokens.

### Multiple Accounts

Some endpoints require the account ID to be specified if the authenticated user is a member of multiple accounts. You can find the eight-digit account ID inside <a href=\"https://app.phrase.com/\" target=\"_blank\">Translation Center</a> by switching to the desired account and then visiting the account details page. If required, you can specify the account just like a normal parameter within the request.

## Pagination

Endpoints that return a list or resources will usually return paginated results and include 25 items by default. To access further pages, use the `page` parameter:

```
$ curl -u EMAIL_OR_ACCESS_TOKEN \"https://api.phrase.com/v2/projects?page=2\"
```

Some endpoints also allow a custom page size by using the `per_page` parameter:

```
$ curl -u EMAIL_OR_ACCESS_TOKEN \"https://api.phrase.com/v2/projects?page=2&per_page=50\"
```

Unless specified otherwise in the description of the respective endpoint, `per_page` allows you to specify a page size up to 100 items.


## Link-Headers

We provide you with pagination URLs in the [Link Header field](http://tools.ietf.org/html/rfc5988). Make use of this information to avoid building pagination URLs yourself.

```
Link: <https://api.phrase.com/v2/projects?page=1>; rel=\"first\", <https://api.phrase.com/v2/projects?page=3>; rel=\"prev\", <https://api.phrase.com/v2/projects?page=5>; rel=\"next\", <https://api.phrase.com/v2/projects?page=9>; rel=\"last\"
```

Possible `rel` values are:

<div class=\"table-responsive\">
  <table class=\"basic-table\">
    <thead>
      <tr class=\"basic-table__row basic-table__row--header\">
        <th class=\"basic-table__cell basic-table__cell--header\">Value</th>
        <th class=\"basic-table__cell basic-table__cell--header\">Description</th>
      </tr>
    </thead>
    <tbody>
      <tr>
        <td class=\"basic-table__cell\">next</td>
        <td class=\"basic-table__cell\">URL of the next page of results</td>
      </tr>
      <tr>
        <td class=\"basic-table__cell\">last</td>
        <td class=\"basic-table__cell\">URL of the last page of results</td>
      </tr>
      <tr>
        <td class=\"basic-table__cell\">first</td>
        <td class=\"basic-table__cell\">URL of the first page of results</td>
      </tr>
      <tr>
        <td class=\"basic-table__cell\">prev</td>
        <td class=\"basic-table__cell\">URL of the previous page of results</td>
      </tr>
    </tbody>
  </table>
</div>

## Rate Limiting

All API endpoints are subject to rate limiting to ensure good performance for all customers. The rate limit is calculated per user:

* 1000 requests per 5 minutes
* 4 concurrent (parallel) requests

For your convenience we send information on the current rate limit within the response headers:

<div class=\"table-responsive\">
  <table class=\"basic-table\">
    <thead>
      <tr class=\"basic-table__row basic-table__row--header\">
        <th class=\"basic-table__cell basic-table__cell--header\">Header</th>
        <th class=\"basic-table__cell basic-table__cell--header\">Description</th>
      </tr>
    </thead>
    <tbody>
      <tr>
        <td class=\"basic-table__cell\" style=\"white-space: nowrap;\"><code>X-Rate-Limit-Limit</code></td>
        <td class=\"basic-table__cell\">Number of max requests allowed in the current time period</td>
      </tr>
      <tr>
        <td class=\"basic-table__cell\" style=\"white-space: nowrap;\"><code>X-Rate-Limit-Remaining</code></td>
        <td class=\"basic-table__cell\">Number of remaining requests in the current time period</td>
      </tr>
      <tr>
        <td class=\"basic-table__cell\" style=\"white-space: nowrap;\"><code>X-Rate-Limit-Reset</code></td>
        <td class=\"basic-table__cell\">Timestamp of end of current time period as UNIX timestamp</td>
      </tr>
    </tbody>
  </table>
</div>

If you should run into the rate limit, you will receive the HTTP status code `429: Too many requests`.

If you should need higher rate limits, [contact us](https://phrase.com/contact).


## Conditional GET requests / HTTP Caching

<div class=\"alert alert-info\"><p><strong>Note:</strong> Conditional GET requests are currently only supported for <a href=\"#locales_download\">locales#download</a> and <a href=\"#translations_index\">translations#index</a></p></div>

We will return an ETag or Last-Modified header with most GET requests. When you request a resource we recommend to store this value and submit them on subsequent requests as `If-Modified-Since` and `If-None-Match` headers. If the resource has not changed in the meantime, we will return the status `304 Not Modified` instead of rendering and returning the resource again. In most cases this is less time-consuming and makes your application/integration faster.

Please note that all conditional requests that return a response with status 304 don't count against your rate limits.

```
$ curl -i -u EMAIL_OR_ACCESS_TOKEN \"https://api.phrase.com/v2/projects/1234abcd1234abcdefefabcd1234efab/locales/en/download\"
HTTP/1.1 200 OK
ETag: \"abcd1234abcdefefabcd1234efab1234\"
Last-Modified: Wed, 28 Jan 2015 15:31:30 UTC
Status: 200 OK

$ curl -i -u EMAIL_OR_ACCESS_TOKEN \"https://api.phrase.com/v2/projects/1234abcd1234abcdefefabcd1234efab/locales/en/download\" -H 'If-None-Match: \"abcd1234abcdefefabcd1234efab1234\"'
HTTP/1.1 304 Not Modified
ETag: \"abcd1234abcdefefabcd1234efab1234\"
Last-Modified: Wed, 28 Jan 2015 15:31:30 UTC
Status: 304 Not Modified

$ curl -i -u EMAIL_OR_ACCESS_TOKEN \"https://api.phrase.com/v2/projects/1234abcd1234abcdefefabcd1234efab/locales/en/download\" -H \"If-Modified-Since: Wed, 28 Jan 2015 15:31:30 UTC\"
HTTP/1.1 304 Not Modified
Last-Modified: Wed, 28 Jan 2015 15:31:30 UTC
Status: 304 Not Modified
```


## JSONP

The Phrase API supports [JSONP](http://en.wikipedia.org/wiki/JSONP) for all GET requests in order to deal with cross-domain request issues. Just send a `?callback` parameter along with the request to specify the Javascript function name to be called with the response content:

```
$ curl \"https://api.phrase.com/v2/projects?callback=myFunction\"
```

The response will include the normal output for that endpoint, along with a `meta` section including header data:

```
myFunction({
  {
    \"meta\": {
      \"status\": 200,
      ...
    },
    \"data\": [
      {
        \"id\": \"1234abcd1234abc1234abcd1234abc\"
        ...
      }
    ]
  }
});
```

To authenticate a JSONP request, you can send a valid [access token](#authentication) as the `?access_token` parameter along the request:

```
$ curl \"https://api.phrase.com/v2/projects?callback=myFunction&access_token=ACCESS-TOKEN\"
```


## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.

- API version: 2.0.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://developers.phrase.com/api/](https://developers.phrase.com/api/)

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
go get github.com/antihax/optional
```

Put the package under your project folder and add the following in import:

```golang
import "./phrase"
```

## Documentation For Authorization



## Basic

- **Type**: HTTP basic authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextBasicAuth, sw.BasicAuth{
    UserName: "username",
    Password: "password",
})
r, err := client.Service.Operation(auth, args)
```


## Token

- **Type**: API key

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAPIKey, sw.APIKey{
    Key: "APIKEY",
    Prefix: "token",
})
r, err := client.Service.Operation(auth, args)
```



## Getting Started

```golang
package main

import (
	"context"
	"fmt"

	phrase "github.com/phrase/phrase-go"
)

func main() {
	auth := context.WithValue(context.Background(), phrase.ContextAPIKey, phrase.APIKey{
		Key:    "YOUR_API_TOKEN",
		Prefix: "token",
	})

	cfg := phrase.NewConfiguration()
	client := phrase.NewAPIClient(cfg)

	var projectsListOpts phrase.ProjectsListOpts

	data, api_response, _ := client.ProjectsApi.ProjectsList(auth, &projectsListOpts)

	fmt.Printf("%+v\n", data)
	fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
}
```

## Documentation for API Endpoints

All URIs are relative to *https://api.phrase.com/v2*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AccountsApi* | [**AccountShow**](docs/AccountsApi.md#accountshow) | **Get** /accounts/{id} | Get a single account
*AccountsApi* | [**AccountsList**](docs/AccountsApi.md#accountslist) | **Get** /accounts | List accounts
*AuthorizationsApi* | [**AuthorizationCreate**](docs/AuthorizationsApi.md#authorizationcreate) | **Post** /authorizations | Create an authorization
*AuthorizationsApi* | [**AuthorizationDelete**](docs/AuthorizationsApi.md#authorizationdelete) | **Delete** /authorizations/{id} | Delete an authorization
*AuthorizationsApi* | [**AuthorizationShow**](docs/AuthorizationsApi.md#authorizationshow) | **Get** /authorizations/{id} | Get a single authorization
*AuthorizationsApi* | [**AuthorizationUpdate**](docs/AuthorizationsApi.md#authorizationupdate) | **Patch** /authorizations/{id} | Update an authorization
*AuthorizationsApi* | [**AuthorizationsList**](docs/AuthorizationsApi.md#authorizationslist) | **Get** /authorizations | List authorizations
*BitbucketSyncApi* | [**BitbucketSyncExport**](docs/BitbucketSyncApi.md#bitbucketsyncexport) | **Post** /bitbucket_syncs/{id}/export | Export from Phrase to Bitbucket
*BitbucketSyncApi* | [**BitbucketSyncImport**](docs/BitbucketSyncApi.md#bitbucketsyncimport) | **Post** /bitbucket_syncs/{id}/import | Import to Phrase from Bitbucket
*BitbucketSyncApi* | [**BitbucketSyncsList**](docs/BitbucketSyncApi.md#bitbucketsyncslist) | **Get** /bitbucket_syncs | List Bitbucket syncs
*BlacklistedKeysApi* | [**BlacklistedKeyCreate**](docs/BlacklistedKeysApi.md#blacklistedkeycreate) | **Post** /projects/{project_id}/blacklisted_keys | Create a blacklisted key
*BlacklistedKeysApi* | [**BlacklistedKeyDelete**](docs/BlacklistedKeysApi.md#blacklistedkeydelete) | **Delete** /projects/{project_id}/blacklisted_keys/{id} | Delete a blacklisted key
*BlacklistedKeysApi* | [**BlacklistedKeyShow**](docs/BlacklistedKeysApi.md#blacklistedkeyshow) | **Get** /projects/{project_id}/blacklisted_keys/{id} | Get a single blacklisted key
*BlacklistedKeysApi* | [**BlacklistedKeyUpdate**](docs/BlacklistedKeysApi.md#blacklistedkeyupdate) | **Patch** /projects/{project_id}/blacklisted_keys/{id} | Update a blacklisted key
*BlacklistedKeysApi* | [**BlacklistedKeysList**](docs/BlacklistedKeysApi.md#blacklistedkeyslist) | **Get** /projects/{project_id}/blacklisted_keys | List blacklisted keys
*BranchesApi* | [**BranchCompare**](docs/BranchesApi.md#branchcompare) | **Get** /projects/{project_id}/branches/{name}/compare | Compare branches
*BranchesApi* | [**BranchCreate**](docs/BranchesApi.md#branchcreate) | **Post** /projects/{project_id}/branches | Create a branch
*BranchesApi* | [**BranchDelete**](docs/BranchesApi.md#branchdelete) | **Delete** /projects/{project_id}/branches/{name} | Delete a branch
*BranchesApi* | [**BranchMerge**](docs/BranchesApi.md#branchmerge) | **Patch** /projects/{project_id}/branches/{name}/merge | Merge a branch
*BranchesApi* | [**BranchShow**](docs/BranchesApi.md#branchshow) | **Get** /projects/{project_id}/branches/{name} | Get a single branch
*BranchesApi* | [**BranchUpdate**](docs/BranchesApi.md#branchupdate) | **Patch** /projects/{project_id}/branches/{name} | Update a branch
*BranchesApi* | [**BranchesList**](docs/BranchesApi.md#brancheslist) | **Get** /projects/{project_id}/branches | List branches
*CommentsApi* | [**CommentCreate**](docs/CommentsApi.md#commentcreate) | **Post** /projects/{project_id}/keys/{key_id}/comments | Create a comment
*CommentsApi* | [**CommentDelete**](docs/CommentsApi.md#commentdelete) | **Delete** /projects/{project_id}/keys/{key_id}/comments/{id} | Delete a comment
*CommentsApi* | [**CommentMarkCheck**](docs/CommentsApi.md#commentmarkcheck) | **Get** /projects/{project_id}/keys/{key_id}/comments/{id}/read | Check if comment is read
*CommentsApi* | [**CommentMarkRead**](docs/CommentsApi.md#commentmarkread) | **Patch** /projects/{project_id}/keys/{key_id}/comments/{id}/read | Mark a comment as read
*CommentsApi* | [**CommentMarkUnread**](docs/CommentsApi.md#commentmarkunread) | **Delete** /projects/{project_id}/keys/{key_id}/comments/{id}/read | Mark a comment as unread
*CommentsApi* | [**CommentShow**](docs/CommentsApi.md#commentshow) | **Get** /projects/{project_id}/keys/{key_id}/comments/{id} | Get a single comment
*CommentsApi* | [**CommentUpdate**](docs/CommentsApi.md#commentupdate) | **Patch** /projects/{project_id}/keys/{key_id}/comments/{id} | Update a comment
*CommentsApi* | [**CommentsList**](docs/CommentsApi.md#commentslist) | **Get** /projects/{project_id}/keys/{key_id}/comments | List comments
*DistributionsApi* | [**DistributionCreate**](docs/DistributionsApi.md#distributioncreate) | **Post** /accounts/{account_id}/distributions | Create a distribution
*DistributionsApi* | [**DistributionDelete**](docs/DistributionsApi.md#distributiondelete) | **Delete** /accounts/{account_id}/distributions/{id} | Delete a distribution
*DistributionsApi* | [**DistributionShow**](docs/DistributionsApi.md#distributionshow) | **Get** /accounts/{account_id}/distributions/{id} | Get a single distribution
*DistributionsApi* | [**DistributionUpdate**](docs/DistributionsApi.md#distributionupdate) | **Patch** /accounts/{account_id}/distributions/{id} | Update a distribution
*DistributionsApi* | [**DistributionsList**](docs/DistributionsApi.md#distributionslist) | **Get** /accounts/{account_id}/distributions | List distributions
*FormatsApi* | [**FormatsList**](docs/FormatsApi.md#formatslist) | **Get** /formats | List formats
*GitLabSyncApi* | [**GitlabSyncDelete**](docs/GitLabSyncApi.md#gitlabsyncdelete) | **Delete** /gitlab_syncs/{id} | Delete single Sync Setting
*GitLabSyncApi* | [**GitlabSyncExport**](docs/GitLabSyncApi.md#gitlabsyncexport) | **Post** /gitlab_syncs/{gitlab_sync_id}/export | Export from Phrase to GitLab
*GitLabSyncApi* | [**GitlabSyncHistory**](docs/GitLabSyncApi.md#gitlabsynchistory) | **Get** /gitlab_syncs/{gitlab_sync_id}/history | History of single Sync Setting
*GitLabSyncApi* | [**GitlabSyncImport**](docs/GitLabSyncApi.md#gitlabsyncimport) | **Post** /gitlab_syncs/{gitlab_sync_id}/import | Import from GitLab to Phrase
*GitLabSyncApi* | [**GitlabSyncList**](docs/GitLabSyncApi.md#gitlabsynclist) | **Get** /gitlab_syncs | List GitLab syncs
*GitLabSyncApi* | [**GitlabSyncShow**](docs/GitLabSyncApi.md#gitlabsyncshow) | **Get** /gitlab_syncs/{id} | Get single Sync Setting
*GitLabSyncApi* | [**GitlabSyncUpdate**](docs/GitLabSyncApi.md#gitlabsyncupdate) | **Put** /gitlab_syncs/{id} | Update single Sync Setting
*GlossariesApi* | [**GlossariesList**](docs/GlossariesApi.md#glossarieslist) | **Get** /accounts/{account_id}/glossaries | List glossaries
*GlossariesApi* | [**GlossaryCreate**](docs/GlossariesApi.md#glossarycreate) | **Post** /accounts/{account_id}/glossaries | Create a glossary
*GlossariesApi* | [**GlossaryDelete**](docs/GlossariesApi.md#glossarydelete) | **Delete** /accounts/{account_id}/glossaries/{id} | Delete a glossary
*GlossariesApi* | [**GlossaryShow**](docs/GlossariesApi.md#glossaryshow) | **Get** /accounts/{account_id}/glossaries/{id} | Get a single glossary
*GlossariesApi* | [**GlossaryUpdate**](docs/GlossariesApi.md#glossaryupdate) | **Patch** /accounts/{account_id}/glossaries/{id} | Update a glossary
*GlossaryTermTranslationsApi* | [**GlossaryTermTranslationCreate**](docs/GlossaryTermTranslationsApi.md#glossarytermtranslationcreate) | **Post** /accounts/{account_id}/glossaries/{glossary_id}/terms/{term_id}/translations | Create a glossary term translation
*GlossaryTermTranslationsApi* | [**GlossaryTermTranslationDelete**](docs/GlossaryTermTranslationsApi.md#glossarytermtranslationdelete) | **Delete** /accounts/{account_id}/glossaries/{glossary_id}/terms/{term_id}/translations/{id} | Delete a glossary term translation
*GlossaryTermTranslationsApi* | [**GlossaryTermTranslationUpdate**](docs/GlossaryTermTranslationsApi.md#glossarytermtranslationupdate) | **Patch** /accounts/{account_id}/glossaries/{glossary_id}/terms/{term_id}/translations/{id} | Update a glossary term translation
*GlossaryTermsApi* | [**GlossaryTermCreate**](docs/GlossaryTermsApi.md#glossarytermcreate) | **Post** /accounts/{account_id}/glossaries/{glossary_id}/terms | Create a glossary term
*GlossaryTermsApi* | [**GlossaryTermDelete**](docs/GlossaryTermsApi.md#glossarytermdelete) | **Delete** /accounts/{account_id}/glossaries/{glossary_id}/terms/{id} | Delete a glossary term
*GlossaryTermsApi* | [**GlossaryTermShow**](docs/GlossaryTermsApi.md#glossarytermshow) | **Get** /accounts/{account_id}/glossaries/{glossary_id}/terms/{id} | Get a single glossary term
*GlossaryTermsApi* | [**GlossaryTermUpdate**](docs/GlossaryTermsApi.md#glossarytermupdate) | **Patch** /accounts/{account_id}/glossaries/{glossary_id}/terms/{id} | Update a glossary term
*GlossaryTermsApi* | [**GlossaryTermsList**](docs/GlossaryTermsApi.md#glossarytermslist) | **Get** /accounts/{account_id}/glossaries/{glossary_id}/terms | List glossary terms
*InvitationsApi* | [**InvitationCreate**](docs/InvitationsApi.md#invitationcreate) | **Post** /accounts/{account_id}/invitations | Create a new invitation
*InvitationsApi* | [**InvitationDelete**](docs/InvitationsApi.md#invitationdelete) | **Delete** /accounts/{account_id}/invitations/{id} | Delete an invitation
*InvitationsApi* | [**InvitationResend**](docs/InvitationsApi.md#invitationresend) | **Post** /accounts/{account_id}/invitations/{id}/resend | Resend an invitation
*InvitationsApi* | [**InvitationShow**](docs/InvitationsApi.md#invitationshow) | **Get** /accounts/{account_id}/invitations/{id} | Get a single invitation
*InvitationsApi* | [**InvitationUpdate**](docs/InvitationsApi.md#invitationupdate) | **Patch** /accounts/{account_id}/invitations/{id} | Update an invitation
*InvitationsApi* | [**InvitationsList**](docs/InvitationsApi.md#invitationslist) | **Get** /accounts/{account_id}/invitations | List invitations
*JobLocalesApi* | [**JobLocaleComplete**](docs/JobLocalesApi.md#joblocalecomplete) | **Post** /projects/{project_id}/jobs/{job_id}/locales/{id}/complete | Complete a job locale
*JobLocalesApi* | [**JobLocaleDelete**](docs/JobLocalesApi.md#joblocaledelete) | **Delete** /projects/{project_id}/jobs/{job_id}/locales/{id} | Delete a job locale
*JobLocalesApi* | [**JobLocaleReopen**](docs/JobLocalesApi.md#joblocalereopen) | **Post** /projects/{project_id}/jobs/{job_id}/locales/{id}/reopen | Reopen a job locale
*JobLocalesApi* | [**JobLocaleShow**](docs/JobLocalesApi.md#joblocaleshow) | **Get** /projects/{project_id}/jobs/{job_id}/locale/{id} | Get a single job locale
*JobLocalesApi* | [**JobLocaleUpdate**](docs/JobLocalesApi.md#joblocaleupdate) | **Patch** /projects/{project_id}/jobs/{job_id}/locales/{id} | Update a job locale
*JobLocalesApi* | [**JobLocalesCreate**](docs/JobLocalesApi.md#joblocalescreate) | **Post** /projects/{project_id}/jobs/{job_id}/locales | Create a job locale
*JobLocalesApi* | [**JobLocalesList**](docs/JobLocalesApi.md#joblocaleslist) | **Get** /projects/{project_id}/jobs/{job_id}/locales | List job locales
*JobsApi* | [**JobComplete**](docs/JobsApi.md#jobcomplete) | **Post** /projects/{project_id}/jobs/{id}/complete | Complete a job
*JobsApi* | [**JobCreate**](docs/JobsApi.md#jobcreate) | **Post** /projects/{project_id}/jobs | Create a job
*JobsApi* | [**JobDelete**](docs/JobsApi.md#jobdelete) | **Delete** /projects/{project_id}/jobs/{id} | Delete a job
*JobsApi* | [**JobKeysCreate**](docs/JobsApi.md#jobkeyscreate) | **Post** /projects/{project_id}/jobs/{id}/keys | Add keys to job
*JobsApi* | [**JobKeysDelete**](docs/JobsApi.md#jobkeysdelete) | **Delete** /projects/{project_id}/jobs/{id}/keys | Remove keys from job
*JobsApi* | [**JobReopen**](docs/JobsApi.md#jobreopen) | **Post** /projects/{project_id}/jobs/{id}/reopen | Reopen a job
*JobsApi* | [**JobShow**](docs/JobsApi.md#jobshow) | **Get** /projects/{project_id}/jobs/{id} | Get a single job
*JobsApi* | [**JobStart**](docs/JobsApi.md#jobstart) | **Post** /projects/{project_id}/jobs/{id}/start | Start a job
*JobsApi* | [**JobUpdate**](docs/JobsApi.md#jobupdate) | **Patch** /projects/{project_id}/jobs/{id} | Update a job
*JobsApi* | [**JobsList**](docs/JobsApi.md#jobslist) | **Get** /projects/{project_id}/jobs | List jobs
*KeysApi* | [**KeyCreate**](docs/KeysApi.md#keycreate) | **Post** /projects/{project_id}/keys | Create a key
*KeysApi* | [**KeyDelete**](docs/KeysApi.md#keydelete) | **Delete** /projects/{project_id}/keys/{id} | Delete a key
*KeysApi* | [**KeyShow**](docs/KeysApi.md#keyshow) | **Get** /projects/{project_id}/keys/{id} | Get a single key
*KeysApi* | [**KeyUpdate**](docs/KeysApi.md#keyupdate) | **Patch** /projects/{project_id}/keys/{id} | Update a key
*KeysApi* | [**KeysDelete**](docs/KeysApi.md#keysdelete) | **Delete** /projects/{project_id}/keys | Delete collection of keys
*KeysApi* | [**KeysList**](docs/KeysApi.md#keyslist) | **Get** /projects/{project_id}/keys | List keys
*KeysApi* | [**KeysSearch**](docs/KeysApi.md#keyssearch) | **Post** /projects/{project_id}/keys/search | Search keys
*KeysApi* | [**KeysTag**](docs/KeysApi.md#keystag) | **Patch** /projects/{project_id}/keys/tag | Add tags to collection of keys
*KeysApi* | [**KeysUntag**](docs/KeysApi.md#keysuntag) | **Patch** /projects/{project_id}/keys/untag | Remove tags from collection of keys
*LocalesApi* | [**LocaleCreate**](docs/LocalesApi.md#localecreate) | **Post** /projects/{project_id}/locales | Create a locale
*LocalesApi* | [**LocaleDelete**](docs/LocalesApi.md#localedelete) | **Delete** /projects/{project_id}/locales/{id} | Delete a locale
*LocalesApi* | [**LocaleDownload**](docs/LocalesApi.md#localedownload) | **Get** /projects/{project_id}/locales/{id}/download | Download a locale
*LocalesApi* | [**LocaleShow**](docs/LocalesApi.md#localeshow) | **Get** /projects/{project_id}/locales/{id} | Get a single locale
*LocalesApi* | [**LocaleUpdate**](docs/LocalesApi.md#localeupdate) | **Patch** /projects/{project_id}/locales/{id} | Update a locale
*LocalesApi* | [**LocalesList**](docs/LocalesApi.md#localeslist) | **Get** /projects/{project_id}/locales | List locales
*MembersApi* | [**MemberDelete**](docs/MembersApi.md#memberdelete) | **Delete** /accounts/{account_id}/members/{id} | Remove a user from the account
*MembersApi* | [**MemberShow**](docs/MembersApi.md#membershow) | **Get** /accounts/{account_id}/members/{id} | Get single member
*MembersApi* | [**MemberUpdate**](docs/MembersApi.md#memberupdate) | **Patch** /accounts/{account_id}/members/{id} | Update a member
*MembersApi* | [**MembersList**](docs/MembersApi.md#memberslist) | **Get** /accounts/{account_id}/members | List members
*OrdersApi* | [**OrderConfirm**](docs/OrdersApi.md#orderconfirm) | **Patch** /projects/{project_id}/orders/{id}/confirm | Confirm an order
*OrdersApi* | [**OrderCreate**](docs/OrdersApi.md#ordercreate) | **Post** /projects/{project_id}/orders | Create a new order
*OrdersApi* | [**OrderDelete**](docs/OrdersApi.md#orderdelete) | **Delete** /projects/{project_id}/orders/{id} | Cancel an order
*OrdersApi* | [**OrderShow**](docs/OrdersApi.md#ordershow) | **Get** /projects/{project_id}/orders/{id} | Get a single order
*OrdersApi* | [**OrdersList**](docs/OrdersApi.md#orderslist) | **Get** /projects/{project_id}/orders | List orders
*ProjectsApi* | [**ProjectCreate**](docs/ProjectsApi.md#projectcreate) | **Post** /projects | Create a project
*ProjectsApi* | [**ProjectDelete**](docs/ProjectsApi.md#projectdelete) | **Delete** /projects/{id} | Delete a project
*ProjectsApi* | [**ProjectShow**](docs/ProjectsApi.md#projectshow) | **Get** /projects/{id} | Get a single project
*ProjectsApi* | [**ProjectUpdate**](docs/ProjectsApi.md#projectupdate) | **Patch** /projects/{id} | Update a project
*ProjectsApi* | [**ProjectsList**](docs/ProjectsApi.md#projectslist) | **Get** /projects | List projects
*ReleasesApi* | [**ReleaseCreate**](docs/ReleasesApi.md#releasecreate) | **Post** /accounts/{account_id}/distributions/{distribution_id}/releases | Create a release
*ReleasesApi* | [**ReleaseDelete**](docs/ReleasesApi.md#releasedelete) | **Delete** /accounts/{account_id}/distributions/{distribution_id}/releases/{id} | Delete a release
*ReleasesApi* | [**ReleasePublish**](docs/ReleasesApi.md#releasepublish) | **Post** /accounts/{account_id}/distributions/{distribution_id}/releases/{id}/publish | Publish a release
*ReleasesApi* | [**ReleaseShow**](docs/ReleasesApi.md#releaseshow) | **Get** /accounts/{account_id}/distributions/{distribution_id}/releases/{id} | Get a single release
*ReleasesApi* | [**ReleaseUpdate**](docs/ReleasesApi.md#releaseupdate) | **Patch** /accounts/{account_id}/distributions/{distribution_id}/releases/{id} | Update a release
*ReleasesApi* | [**ReleasesList**](docs/ReleasesApi.md#releaseslist) | **Get** /accounts/{account_id}/distributions/{distribution_id}/releases | List releases
*ScreenshotMarkersApi* | [**ScreenshotMarkerCreate**](docs/ScreenshotMarkersApi.md#screenshotmarkercreate) | **Post** /projects/{project_id}/screenshots/{screenshot_id}/markers | Create a screenshot marker
*ScreenshotMarkersApi* | [**ScreenshotMarkerDelete**](docs/ScreenshotMarkersApi.md#screenshotmarkerdelete) | **Delete** /projects/{project_id}/screenshots/{screenshot_id}/markers | Delete a screenshot marker
*ScreenshotMarkersApi* | [**ScreenshotMarkerShow**](docs/ScreenshotMarkersApi.md#screenshotmarkershow) | **Get** /projects/{project_id}/screenshots/{screenshot_id}/markers/{id} | Get a single screenshot marker
*ScreenshotMarkersApi* | [**ScreenshotMarkerUpdate**](docs/ScreenshotMarkersApi.md#screenshotmarkerupdate) | **Patch** /projects/{project_id}/screenshots/{screenshot_id}/markers | Update a screenshot marker
*ScreenshotMarkersApi* | [**ScreenshotMarkersList**](docs/ScreenshotMarkersApi.md#screenshotmarkerslist) | **Get** /projects/{project_id}/screenshots/{id}/markers | List screenshot markers
*ScreenshotsApi* | [**ScreenshotCreate**](docs/ScreenshotsApi.md#screenshotcreate) | **Post** /projects/{project_id}/screenshots | Create a screenshot
*ScreenshotsApi* | [**ScreenshotDelete**](docs/ScreenshotsApi.md#screenshotdelete) | **Delete** /projects/{project_id}/screenshots/{id} | Delete a screenshot
*ScreenshotsApi* | [**ScreenshotShow**](docs/ScreenshotsApi.md#screenshotshow) | **Get** /projects/{project_id}/screenshots/{id} | Get a single screenshot
*ScreenshotsApi* | [**ScreenshotUpdate**](docs/ScreenshotsApi.md#screenshotupdate) | **Patch** /projects/{project_id}/screenshots/{id} | Update a screenshot
*ScreenshotsApi* | [**ScreenshotsList**](docs/ScreenshotsApi.md#screenshotslist) | **Get** /projects/{project_id}/screenshots | List screenshots
*SpacesApi* | [**SpaceCreate**](docs/SpacesApi.md#spacecreate) | **Post** /accounts/{account_id}/spaces | Create a Space
*SpacesApi* | [**SpaceDelete**](docs/SpacesApi.md#spacedelete) | **Delete** /accounts/{account_id}/spaces/{id} | Delete Space
*SpacesApi* | [**SpaceShow**](docs/SpacesApi.md#spaceshow) | **Get** /accounts/{account_id}/spaces/{id} | Get Space
*SpacesApi* | [**SpaceUpdate**](docs/SpacesApi.md#spaceupdate) | **Patch** /accounts/{account_id}/spaces/{id} | Update Space
*SpacesApi* | [**SpacesList**](docs/SpacesApi.md#spaceslist) | **Get** /accounts/{account_id}/spaces | List Spaces
*SpacesApi* | [**SpacesProjectsCreate**](docs/SpacesApi.md#spacesprojectscreate) | **Post** /accounts/{account_id}/spaces/{space_id}/projects | Add Project
*SpacesApi* | [**SpacesProjectsDelete**](docs/SpacesApi.md#spacesprojectsdelete) | **Delete** /accounts/{account_id}/spaces/{space_id}/projects/{id} | Remove Project
*SpacesApi* | [**SpacesProjectsList**](docs/SpacesApi.md#spacesprojectslist) | **Get** /accounts/{account_id}/spaces/{space_id}/projects | List Projects
*StyleGuidesApi* | [**StyleguideCreate**](docs/StyleGuidesApi.md#styleguidecreate) | **Post** /projects/{project_id}/styleguides | Create a style guide
*StyleGuidesApi* | [**StyleguideDelete**](docs/StyleGuidesApi.md#styleguidedelete) | **Delete** /projects/{project_id}/styleguides/{id} | Delete a style guide
*StyleGuidesApi* | [**StyleguideShow**](docs/StyleGuidesApi.md#styleguideshow) | **Get** /projects/{project_id}/styleguides/{id} | Get a single style guide
*StyleGuidesApi* | [**StyleguideUpdate**](docs/StyleGuidesApi.md#styleguideupdate) | **Patch** /projects/{project_id}/styleguides/{id} | Update a style guide
*StyleGuidesApi* | [**StyleguidesList**](docs/StyleGuidesApi.md#styleguideslist) | **Get** /projects/{project_id}/styleguides | List style guides
*TagsApi* | [**TagCreate**](docs/TagsApi.md#tagcreate) | **Post** /projects/{project_id}/tags | Create a tag
*TagsApi* | [**TagDelete**](docs/TagsApi.md#tagdelete) | **Delete** /projects/{project_id}/tags/{name} | Delete a tag
*TagsApi* | [**TagShow**](docs/TagsApi.md#tagshow) | **Get** /projects/{project_id}/tags/{name} | Get a single tag
*TagsApi* | [**TagsList**](docs/TagsApi.md#tagslist) | **Get** /projects/{project_id}/tags | List tags
*TranslationsApi* | [**TranslationCreate**](docs/TranslationsApi.md#translationcreate) | **Post** /projects/{project_id}/translations | Create a translation
*TranslationsApi* | [**TranslationExclude**](docs/TranslationsApi.md#translationexclude) | **Patch** /projects/{project_id}/translations/{id}/exclude | Exclude a translation from export
*TranslationsApi* | [**TranslationInclude**](docs/TranslationsApi.md#translationinclude) | **Patch** /projects/{project_id}/translations/{id}/include | Revoke exclusion of a translation in export
*TranslationsApi* | [**TranslationReview**](docs/TranslationsApi.md#translationreview) | **Patch** /projects/{project_id}/translations/{id}/review | Review a translation
*TranslationsApi* | [**TranslationShow**](docs/TranslationsApi.md#translationshow) | **Get** /projects/{project_id}/translations/{id} | Get a single translation
*TranslationsApi* | [**TranslationUnverify**](docs/TranslationsApi.md#translationunverify) | **Patch** /projects/{project_id}/translations/{id}/unverify | Mark a translation as unverified
*TranslationsApi* | [**TranslationUpdate**](docs/TranslationsApi.md#translationupdate) | **Patch** /projects/{project_id}/translations/{id} | Update a translation
*TranslationsApi* | [**TranslationVerify**](docs/TranslationsApi.md#translationverify) | **Patch** /projects/{project_id}/translations/{id}/verify | Verify a translation
*TranslationsApi* | [**TranslationsByKey**](docs/TranslationsApi.md#translationsbykey) | **Get** /projects/{project_id}/keys/{key_id}/translations | List translations by key
*TranslationsApi* | [**TranslationsByLocale**](docs/TranslationsApi.md#translationsbylocale) | **Get** /projects/{project_id}/locales/{locale_id}/translations | List translations by locale
*TranslationsApi* | [**TranslationsExclude**](docs/TranslationsApi.md#translationsexclude) | **Patch** /projects/{project_id}/translations/exclude | Set exclude from export flag on translations selected by query
*TranslationsApi* | [**TranslationsInclude**](docs/TranslationsApi.md#translationsinclude) | **Patch** /projects/{project_id}/translations/include | Remove exlude from import flag from translations selected by query
*TranslationsApi* | [**TranslationsList**](docs/TranslationsApi.md#translationslist) | **Get** /projects/{project_id}/translations | List all translations
*TranslationsApi* | [**TranslationsReview**](docs/TranslationsApi.md#translationsreview) | **Patch** /projects/{project_id}/translations/review | Review translations selected by query
*TranslationsApi* | [**TranslationsSearch**](docs/TranslationsApi.md#translationssearch) | **Post** /projects/{project_id}/translations/search | Search translations
*TranslationsApi* | [**TranslationsUnverify**](docs/TranslationsApi.md#translationsunverify) | **Patch** /projects/{project_id}/translations/unverify | Mark translations selected by query as unverified
*TranslationsApi* | [**TranslationsVerify**](docs/TranslationsApi.md#translationsverify) | **Patch** /projects/{project_id}/translations/verify | Verify translations selected by query
*UploadsApi* | [**UploadCreate**](docs/UploadsApi.md#uploadcreate) | **Post** /projects/{project_id}/uploads | Upload a new file
*UploadsApi* | [**UploadShow**](docs/UploadsApi.md#uploadshow) | **Get** /projects/{project_id}/uploads/{id} | View upload details
*UploadsApi* | [**UploadsList**](docs/UploadsApi.md#uploadslist) | **Get** /projects/{project_id}/uploads | List uploads
*UsersApi* | [**ShowUser**](docs/UsersApi.md#showuser) | **Get** /user | Show current User
*VersionsHistoryApi* | [**VersionShow**](docs/VersionsHistoryApi.md#versionshow) | **Get** /projects/{project_id}/translations/{translation_id}/versions/{id} | Get a single version
*VersionsHistoryApi* | [**VersionsList**](docs/VersionsHistoryApi.md#versionslist) | **Get** /projects/{project_id}/translations/{translation_id}/versions | List all versions
*WebhooksApi* | [**WebhookCreate**](docs/WebhooksApi.md#webhookcreate) | **Post** /projects/{project_id}/webhooks | Create a webhook
*WebhooksApi* | [**WebhookDelete**](docs/WebhooksApi.md#webhookdelete) | **Delete** /projects/{project_id}/webhooks/{id} | Delete a webhook
*WebhooksApi* | [**WebhookShow**](docs/WebhooksApi.md#webhookshow) | **Get** /projects/{project_id}/webhooks/{id} | Get a single webhook
*WebhooksApi* | [**WebhookTest**](docs/WebhooksApi.md#webhooktest) | **Post** /projects/{project_id}/webhooks/{id}/test | Test a webhook
*WebhooksApi* | [**WebhookUpdate**](docs/WebhooksApi.md#webhookupdate) | **Patch** /projects/{project_id}/webhooks/{id} | Update a webhook
*WebhooksApi* | [**WebhooksList**](docs/WebhooksApi.md#webhookslist) | **Get** /projects/{project_id}/webhooks | List webhooks


## Documentation For Models

 - [Account](docs/Account.md)
 - [AccountDetails](docs/AccountDetails.md)
 - [AccountDetails1](docs/AccountDetails1.md)
 - [AffectedCount](docs/AffectedCount.md)
 - [AffectedResources](docs/AffectedResources.md)
 - [Authorization](docs/Authorization.md)
 - [AuthorizationCreateParameters](docs/AuthorizationCreateParameters.md)
 - [AuthorizationUpdateParameters](docs/AuthorizationUpdateParameters.md)
 - [AuthorizationWithToken](docs/AuthorizationWithToken.md)
 - [AuthorizationWithToken1](docs/AuthorizationWithToken1.md)
 - [BitbucketSync](docs/BitbucketSync.md)
 - [BitbucketSyncExportParameters](docs/BitbucketSyncExportParameters.md)
 - [BitbucketSyncExportResponse](docs/BitbucketSyncExportResponse.md)
 - [BitbucketSyncImportParameters](docs/BitbucketSyncImportParameters.md)
 - [BlacklistedKey](docs/BlacklistedKey.md)
 - [BlacklistedKeyCreateParameters](docs/BlacklistedKeyCreateParameters.md)
 - [BlacklistedKeyUpdateParameters](docs/BlacklistedKeyUpdateParameters.md)
 - [Branch](docs/Branch.md)
 - [BranchCreateParameters](docs/BranchCreateParameters.md)
 - [BranchMergeParameters](docs/BranchMergeParameters.md)
 - [BranchUpdateParameters](docs/BranchUpdateParameters.md)
 - [Comment](docs/Comment.md)
 - [CommentCreateParameters](docs/CommentCreateParameters.md)
 - [CommentMarkReadParameters](docs/CommentMarkReadParameters.md)
 - [CommentUpdateParameters](docs/CommentUpdateParameters.md)
 - [Distribution](docs/Distribution.md)
 - [DistributionCreateParameters](docs/DistributionCreateParameters.md)
 - [DistributionPreview](docs/DistributionPreview.md)
 - [DistributionUpdateParameters](docs/DistributionUpdateParameters.md)
 - [Format](docs/Format.md)
 - [GitlabSync](docs/GitlabSync.md)
 - [GitlabSyncExport](docs/GitlabSyncExport.md)
 - [GitlabSyncExportParameters](docs/GitlabSyncExportParameters.md)
 - [GitlabSyncHistory](docs/GitlabSyncHistory.md)
 - [GitlabSyncImportParameters](docs/GitlabSyncImportParameters.md)
 - [Glossary](docs/Glossary.md)
 - [GlossaryCreateParameters](docs/GlossaryCreateParameters.md)
 - [GlossaryTerm](docs/GlossaryTerm.md)
 - [GlossaryTermCreateParameters](docs/GlossaryTermCreateParameters.md)
 - [GlossaryTermTranslation](docs/GlossaryTermTranslation.md)
 - [GlossaryTermTranslationCreateParameters](docs/GlossaryTermTranslationCreateParameters.md)
 - [GlossaryTermTranslationUpdateParameters](docs/GlossaryTermTranslationUpdateParameters.md)
 - [GlossaryTermUpdateParameters](docs/GlossaryTermUpdateParameters.md)
 - [GlossaryUpdateParameters](docs/GlossaryUpdateParameters.md)
 - [Invitation](docs/Invitation.md)
 - [InvitationCreateParameters](docs/InvitationCreateParameters.md)
 - [InvitationUpdateParameters](docs/InvitationUpdateParameters.md)
 - [Job](docs/Job.md)
 - [JobCompleteParameters](docs/JobCompleteParameters.md)
 - [JobCreateParameters](docs/JobCreateParameters.md)
 - [JobDetails](docs/JobDetails.md)
 - [JobDetails1](docs/JobDetails1.md)
 - [JobKeysCreateParameters](docs/JobKeysCreateParameters.md)
 - [JobLocale](docs/JobLocale.md)
 - [JobLocaleCompleteParameters](docs/JobLocaleCompleteParameters.md)
 - [JobLocaleReopenParameters](docs/JobLocaleReopenParameters.md)
 - [JobLocaleUpdateParameters](docs/JobLocaleUpdateParameters.md)
 - [JobLocalesCreateParameters](docs/JobLocalesCreateParameters.md)
 - [JobPreview](docs/JobPreview.md)
 - [JobReopenParameters](docs/JobReopenParameters.md)
 - [JobStartParameters](docs/JobStartParameters.md)
 - [JobUpdateParameters](docs/JobUpdateParameters.md)
 - [KeyCreateParameters](docs/KeyCreateParameters.md)
 - [KeyPreview](docs/KeyPreview.md)
 - [KeyUpdateParameters](docs/KeyUpdateParameters.md)
 - [KeysSearchParameters](docs/KeysSearchParameters.md)
 - [KeysTagParameters](docs/KeysTagParameters.md)
 - [KeysUntagParameters](docs/KeysUntagParameters.md)
 - [Locale](docs/Locale.md)
 - [LocaleCreateParameters](docs/LocaleCreateParameters.md)
 - [LocaleDetails](docs/LocaleDetails.md)
 - [LocaleDetails1](docs/LocaleDetails1.md)
 - [LocalePreview](docs/LocalePreview.md)
 - [LocaleStatistics](docs/LocaleStatistics.md)
 - [LocaleUpdateParameters](docs/LocaleUpdateParameters.md)
 - [Member](docs/Member.md)
 - [MemberUpdateParameters](docs/MemberUpdateParameters.md)
 - [OrderConfirmParameters](docs/OrderConfirmParameters.md)
 - [OrderCreateParameters](docs/OrderCreateParameters.md)
 - [Project](docs/Project.md)
 - [ProjectCreateParameters](docs/ProjectCreateParameters.md)
 - [ProjectDetails](docs/ProjectDetails.md)
 - [ProjectDetails1](docs/ProjectDetails1.md)
 - [ProjectLocales](docs/ProjectLocales.md)
 - [ProjectLocales1](docs/ProjectLocales1.md)
 - [ProjectShort](docs/ProjectShort.md)
 - [ProjectUpdateParameters](docs/ProjectUpdateParameters.md)
 - [Release](docs/Release.md)
 - [ReleaseCreateParameters](docs/ReleaseCreateParameters.md)
 - [ReleasePreview](docs/ReleasePreview.md)
 - [ReleaseUpdateParameters](docs/ReleaseUpdateParameters.md)
 - [Screenshot](docs/Screenshot.md)
 - [ScreenshotCreateParameters](docs/ScreenshotCreateParameters.md)
 - [ScreenshotMarker](docs/ScreenshotMarker.md)
 - [ScreenshotMarkerCreateParameters](docs/ScreenshotMarkerCreateParameters.md)
 - [ScreenshotMarkerUpdateParameters](docs/ScreenshotMarkerUpdateParameters.md)
 - [ScreenshotUpdateParameters](docs/ScreenshotUpdateParameters.md)
 - [Space](docs/Space.md)
 - [SpaceCreateParameters](docs/SpaceCreateParameters.md)
 - [SpaceUpdateParameters](docs/SpaceUpdateParameters.md)
 - [SpacesProjectsCreateParameters](docs/SpacesProjectsCreateParameters.md)
 - [Styleguide](docs/Styleguide.md)
 - [StyleguideCreateParameters](docs/StyleguideCreateParameters.md)
 - [StyleguideDetails](docs/StyleguideDetails.md)
 - [StyleguideDetails1](docs/StyleguideDetails1.md)
 - [StyleguidePreview](docs/StyleguidePreview.md)
 - [StyleguideUpdateParameters](docs/StyleguideUpdateParameters.md)
 - [Tag](docs/Tag.md)
 - [TagCreateParameters](docs/TagCreateParameters.md)
 - [TagWithStats](docs/TagWithStats.md)
 - [TagWithStats1](docs/TagWithStats1.md)
 - [TagWithStats1Statistics](docs/TagWithStats1Statistics.md)
 - [TagWithStats1Statistics1](docs/TagWithStats1Statistics1.md)
 - [Translation](docs/Translation.md)
 - [TranslationCreateParameters](docs/TranslationCreateParameters.md)
 - [TranslationDetails](docs/TranslationDetails.md)
 - [TranslationDetails1](docs/TranslationDetails1.md)
 - [TranslationExcludeParameters](docs/TranslationExcludeParameters.md)
 - [TranslationIncludeParameters](docs/TranslationIncludeParameters.md)
 - [TranslationKey](docs/TranslationKey.md)
 - [TranslationKeyDetails](docs/TranslationKeyDetails.md)
 - [TranslationKeyDetails1](docs/TranslationKeyDetails1.md)
 - [TranslationOrder](docs/TranslationOrder.md)
 - [TranslationReviewParameters](docs/TranslationReviewParameters.md)
 - [TranslationUnverifyParameters](docs/TranslationUnverifyParameters.md)
 - [TranslationUpdateParameters](docs/TranslationUpdateParameters.md)
 - [TranslationVerifyParameters](docs/TranslationVerifyParameters.md)
 - [TranslationVersion](docs/TranslationVersion.md)
 - [TranslationVersionWithUser](docs/TranslationVersionWithUser.md)
 - [TranslationVersionWithUser1](docs/TranslationVersionWithUser1.md)
 - [TranslationsExcludeParameters](docs/TranslationsExcludeParameters.md)
 - [TranslationsIncludeParameters](docs/TranslationsIncludeParameters.md)
 - [TranslationsReviewParameters](docs/TranslationsReviewParameters.md)
 - [TranslationsSearchParameters](docs/TranslationsSearchParameters.md)
 - [TranslationsUnverifyParameters](docs/TranslationsUnverifyParameters.md)
 - [TranslationsVerifyParameters](docs/TranslationsVerifyParameters.md)
 - [Upload](docs/Upload.md)
 - [UploadCreateParameters](docs/UploadCreateParameters.md)
 - [UploadSummary](docs/UploadSummary.md)
 - [User](docs/User.md)
 - [UserPreview](docs/UserPreview.md)
 - [Webhook](docs/Webhook.md)
 - [WebhookCreateParameters](docs/WebhookCreateParameters.md)
 - [WebhookUpdateParameters](docs/WebhookUpdateParameters.md)


## Author

support@phrase.com

