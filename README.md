# Go API client for phrase

Phrase Strings is a translation management platform for software projects. You can collaborate on language file translation with your team or order translations through our platform. The API allows you to import locale files, download locale files, tag keys or interact in other ways with the localization data stored in Phrase Strings for your account.

## Installation

```shell
go get github.com/phrase/phrase-go
```

## Getting Started

```golang
package main

import (
	"context"
	"fmt"

	phrase "github.com/phrase/phrase-go/v3" // x-release-please-major
)

func main() {
	auth := context.WithValue(context.Background(), phrase.ContextAPIKey, phrase.APIKey{
		Key:    "YOUR_API_TOKEN",
		Prefix: "token",
	})

	cfg := phrase.NewConfiguration()
	client := phrase.NewAPIClient(cfg)

	var projectsListOpts phrase.ProjectsListOpts

	projects, apiResponse, err := client.ProjectsApi.ProjectsList(auth, &projectsListOpts)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%+v\n", projects)
		fmt.Printf("%+v\n", apiResponse) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
	}
}
```

## Datacenters

The API is only accessible via HTTPS and the current version is <code>v2</code>, which results in a base URL like: <code>https://api.phrase.com/v2</code> depending on the datacenter.

### EU Datacenter
```
https://api.phrase.com/v2
```

This is the default datacenter.

### US Datacenter
```
https://api.us.app.phrase.com/v2/
```

### Specifying US Datacenter
You can use the US datacenter by setting the following:
```
client.ChangeBasePath("https://api.us.app.phrase.com/v2/")
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
*BitbucketSyncApi* | [**BitbucketSyncExport**](docs/BitbucketSyncApi.md#bitbucketsyncexport) | **Post** /bitbucket_syncs/{id}/export | Export from Phrase Strings to Bitbucket
*BitbucketSyncApi* | [**BitbucketSyncImport**](docs/BitbucketSyncApi.md#bitbucketsyncimport) | **Post** /bitbucket_syncs/{id}/import | Import to Phrase Strings from Bitbucket
*BitbucketSyncApi* | [**BitbucketSyncsList**](docs/BitbucketSyncApi.md#bitbucketsyncslist) | **Get** /bitbucket_syncs | List Bitbucket syncs
*BlacklistedKeysApi* | [**BlacklistedKeyCreate**](docs/BlacklistedKeysApi.md#blacklistedkeycreate) | **Post** /projects/{project_id}/blacklisted_keys | Create a blocked key
*BlacklistedKeysApi* | [**BlacklistedKeyDelete**](docs/BlacklistedKeysApi.md#blacklistedkeydelete) | **Delete** /projects/{project_id}/blacklisted_keys/{id} | Delete a blocked key
*BlacklistedKeysApi* | [**BlacklistedKeyShow**](docs/BlacklistedKeysApi.md#blacklistedkeyshow) | **Get** /projects/{project_id}/blacklisted_keys/{id} | Get a single blocked key
*BlacklistedKeysApi* | [**BlacklistedKeyUpdate**](docs/BlacklistedKeysApi.md#blacklistedkeyupdate) | **Patch** /projects/{project_id}/blacklisted_keys/{id} | Update a blocked key
*BlacklistedKeysApi* | [**BlacklistedKeysList**](docs/BlacklistedKeysApi.md#blacklistedkeyslist) | **Get** /projects/{project_id}/blacklisted_keys | List blocked keys
*BranchesApi* | [**BranchCompare**](docs/BranchesApi.md#branchcompare) | **Get** /projects/{project_id}/branches/{name}/compare | Compare branches
*BranchesApi* | [**BranchCreate**](docs/BranchesApi.md#branchcreate) | **Post** /projects/{project_id}/branches | Create a branch
*BranchesApi* | [**BranchDelete**](docs/BranchesApi.md#branchdelete) | **Delete** /projects/{project_id}/branches/{name} | Delete a branch
*BranchesApi* | [**BranchMerge**](docs/BranchesApi.md#branchmerge) | **Patch** /projects/{project_id}/branches/{name}/merge | Merge a branch
*BranchesApi* | [**BranchShow**](docs/BranchesApi.md#branchshow) | **Get** /projects/{project_id}/branches/{name} | Get a single branch
*BranchesApi* | [**BranchUpdate**](docs/BranchesApi.md#branchupdate) | **Patch** /projects/{project_id}/branches/{name} | Update a branch
*BranchesApi* | [**BranchesList**](docs/BranchesApi.md#brancheslist) | **Get** /projects/{project_id}/branches | List branches
*CommentReactionsApi* | [**ReactionCreate**](docs/CommentReactionsApi.md#reactioncreate) | **Post** /projects/{project_id}/keys/{key_id}/comments/{comment_id}/reactions | Create a reaction
*CommentReactionsApi* | [**ReactionDelete**](docs/CommentReactionsApi.md#reactiondelete) | **Delete** /projects/{project_id}/keys/{key_id}/comments/{comment_id}/reactions/{id} | Delete a reaction
*CommentReactionsApi* | [**ReactionShow**](docs/CommentReactionsApi.md#reactionshow) | **Get** /projects/{project_id}/keys/{key_id}/comments/{comment_id}/reactions/{id} | Get a single reaction
*CommentReactionsApi* | [**ReactionsList**](docs/CommentReactionsApi.md#reactionslist) | **Get** /projects/{project_id}/keys/{key_id}/comments/{comment_id}/reactions | List reactions
*CommentRepliesApi* | [**RepliesList**](docs/CommentRepliesApi.md#replieslist) | **Get** /projects/{project_id}/keys/{key_id}/comments/{comment_id}/replies | List replies
*CommentRepliesApi* | [**ReplyCreate**](docs/CommentRepliesApi.md#replycreate) | **Post** /projects/{project_id}/keys/{key_id}/comments/{comment_id}/replies | Create a reply
*CommentRepliesApi* | [**ReplyDelete**](docs/CommentRepliesApi.md#replydelete) | **Delete** /projects/{project_id}/keys/{key_id}/comments/{comment_id}/replies/{id} | Delete a reply
*CommentRepliesApi* | [**ReplyMarkAsRead**](docs/CommentRepliesApi.md#replymarkasread) | **Patch** /projects/{project_id}/keys/{key_id}/comments/{comment_id}/replies/{id}/mark_as_read | Mark a reply as read
*CommentRepliesApi* | [**ReplyMarkAsUnread**](docs/CommentRepliesApi.md#replymarkasunread) | **Patch** /projects/{project_id}/keys/{key_id}/comments/{comment_id}/replies/{id}/mark_as_unread | Mark a reply as unread
*CommentRepliesApi* | [**ReplyShow**](docs/CommentRepliesApi.md#replyshow) | **Get** /projects/{project_id}/keys/{key_id}/comments/{comment_id}/replies/{id} | Get a single reply
*CommentsApi* | [**CommentCreate**](docs/CommentsApi.md#commentcreate) | **Post** /projects/{project_id}/keys/{key_id}/comments | Create a comment
*CommentsApi* | [**CommentDelete**](docs/CommentsApi.md#commentdelete) | **Delete** /projects/{project_id}/keys/{key_id}/comments/{id} | Delete a comment
*CommentsApi* | [**CommentMarkCheck**](docs/CommentsApi.md#commentmarkcheck) | **Get** /projects/{project_id}/keys/{key_id}/comments/{id}/read | Check if comment is read
*CommentsApi* | [**CommentMarkRead**](docs/CommentsApi.md#commentmarkread) | **Patch** /projects/{project_id}/keys/{key_id}/comments/{id}/read | Mark a comment as read
*CommentsApi* | [**CommentMarkUnread**](docs/CommentsApi.md#commentmarkunread) | **Delete** /projects/{project_id}/keys/{key_id}/comments/{id}/read | Mark a comment as unread
*CommentsApi* | [**CommentShow**](docs/CommentsApi.md#commentshow) | **Get** /projects/{project_id}/keys/{key_id}/comments/{id} | Get a single comment
*CommentsApi* | [**CommentUpdate**](docs/CommentsApi.md#commentupdate) | **Patch** /projects/{project_id}/keys/{key_id}/comments/{id} | Update a comment
*CommentsApi* | [**CommentsList**](docs/CommentsApi.md#commentslist) | **Get** /projects/{project_id}/keys/{key_id}/comments | List comments
*CustomMetadataApi* | [**CustomMetadataPropertiesDelete**](docs/CustomMetadataApi.md#custommetadatapropertiesdelete) | **Delete** /accounts/{account_id}/custom_metadata/properties/{id} | Destroy property
*CustomMetadataApi* | [**CustomMetadataPropertiesList**](docs/CustomMetadataApi.md#custommetadatapropertieslist) | **Get** /accounts/{account_id}/custom_metadata/properties | List properties
*CustomMetadataApi* | [**CustomMetadataPropertyCreate**](docs/CustomMetadataApi.md#custommetadatapropertycreate) | **Post** /accounts/{account_id}/custom_metadata/properties | Create a property
*CustomMetadataApi* | [**CustomMetadataPropertyShow**](docs/CustomMetadataApi.md#custommetadatapropertyshow) | **Get** /accounts/{account_id}/custom_metadata/properties/{id} | Get a single property
*CustomMetadataApi* | [**CustomMetadataPropertyUpdate**](docs/CustomMetadataApi.md#custommetadatapropertyupdate) | **Patch** /accounts/{account_id}/custom_metadata/properties/{id} | Update a property
*DistributionsApi* | [**DistributionCreate**](docs/DistributionsApi.md#distributioncreate) | **Post** /accounts/{account_id}/distributions | Create a distribution
*DistributionsApi* | [**DistributionDelete**](docs/DistributionsApi.md#distributiondelete) | **Delete** /accounts/{account_id}/distributions/{id} | Delete a distribution
*DistributionsApi* | [**DistributionShow**](docs/DistributionsApi.md#distributionshow) | **Get** /accounts/{account_id}/distributions/{id} | Get a single distribution
*DistributionsApi* | [**DistributionUpdate**](docs/DistributionsApi.md#distributionupdate) | **Patch** /accounts/{account_id}/distributions/{id} | Update a distribution
*DistributionsApi* | [**DistributionsList**](docs/DistributionsApi.md#distributionslist) | **Get** /accounts/{account_id}/distributions | List distributions
*DocumentsApi* | [**DocumentDelete**](docs/DocumentsApi.md#documentdelete) | **Delete** /projects/{project_id}/documents/{id} | Delete document
*DocumentsApi* | [**DocumentsList**](docs/DocumentsApi.md#documentslist) | **Get** /projects/{project_id}/documents | List documents
*FigmaAttachmentsApi* | [**FigmaAttachmentCreate**](docs/FigmaAttachmentsApi.md#figmaattachmentcreate) | **Post** /projects/{project_id}/figma_attachments | Create a Figma attachment
*FigmaAttachmentsApi* | [**FigmaAttachmentDelete**](docs/FigmaAttachmentsApi.md#figmaattachmentdelete) | **Delete** /projects/{project_id}/figma_attachments/{id} | Delete a Figma attachment
*FigmaAttachmentsApi* | [**FigmaAttachmentShow**](docs/FigmaAttachmentsApi.md#figmaattachmentshow) | **Get** /projects/{project_id}/figma_attachments/{id} | Get a single Figma attachment
*FigmaAttachmentsApi* | [**FigmaAttachmentUpdate**](docs/FigmaAttachmentsApi.md#figmaattachmentupdate) | **Patch** /projects/{project_id}/figma_attachments/{id} | Update a Figma attachment
*FigmaAttachmentsApi* | [**FigmaAttachmentsList**](docs/FigmaAttachmentsApi.md#figmaattachmentslist) | **Get** /projects/{project_id}/figma_attachments | List Figma attachments
*FormatsApi* | [**FormatsList**](docs/FormatsApi.md#formatslist) | **Get** /formats | List formats
*GitHubSyncApi* | [**GithubSyncExport**](docs/GitHubSyncApi.md#githubsyncexport) | **Post** /github_syncs/export | Export from Phrase Strings to GitHub
*GitHubSyncApi* | [**GithubSyncImport**](docs/GitHubSyncApi.md#githubsyncimport) | **Post** /github_syncs/import | Import to Phrase Strings from GitHub
*GitLabSyncApi* | [**GitlabSyncDelete**](docs/GitLabSyncApi.md#gitlabsyncdelete) | **Delete** /gitlab_syncs/{id} | Delete single Sync Setting
*GitLabSyncApi* | [**GitlabSyncExport**](docs/GitLabSyncApi.md#gitlabsyncexport) | **Post** /gitlab_syncs/{gitlab_sync_id}/export | Export from Phrase Strings to GitLab
*GitLabSyncApi* | [**GitlabSyncHistory**](docs/GitLabSyncApi.md#gitlabsynchistory) | **Get** /gitlab_syncs/{gitlab_sync_id}/history | History of single Sync Setting
*GitLabSyncApi* | [**GitlabSyncImport**](docs/GitLabSyncApi.md#gitlabsyncimport) | **Post** /gitlab_syncs/{gitlab_sync_id}/import | Import from GitLab to Phrase
*GitLabSyncApi* | [**GitlabSyncList**](docs/GitLabSyncApi.md#gitlabsynclist) | **Get** /gitlab_syncs | List GitLab syncs
*GitLabSyncApi* | [**GitlabSyncShow**](docs/GitLabSyncApi.md#gitlabsyncshow) | **Get** /gitlab_syncs/{id} | Get single Sync Setting
*GitLabSyncApi* | [**GitlabSyncUpdate**](docs/GitLabSyncApi.md#gitlabsyncupdate) | **Put** /gitlab_syncs/{id} | Update single Sync Setting
*GlossariesApi* | [**GlossariesList**](docs/GlossariesApi.md#glossarieslist) | **Get** /accounts/{account_id}/glossaries | List term bases
*GlossariesApi* | [**GlossaryCreate**](docs/GlossariesApi.md#glossarycreate) | **Post** /accounts/{account_id}/glossaries | Create a term base
*GlossariesApi* | [**GlossaryDelete**](docs/GlossariesApi.md#glossarydelete) | **Delete** /accounts/{account_id}/glossaries/{id} | Delete a term base
*GlossariesApi* | [**GlossaryShow**](docs/GlossariesApi.md#glossaryshow) | **Get** /accounts/{account_id}/glossaries/{id} | Get a single term base
*GlossariesApi* | [**GlossaryUpdate**](docs/GlossariesApi.md#glossaryupdate) | **Patch** /accounts/{account_id}/glossaries/{id} | Update a term base
*GlossaryTermTranslationsApi* | [**GlossaryTermTranslationCreate**](docs/GlossaryTermTranslationsApi.md#glossarytermtranslationcreate) | **Post** /accounts/{account_id}/glossaries/{glossary_id}/terms/{term_id}/translations | Create a translation for a term
*GlossaryTermTranslationsApi* | [**GlossaryTermTranslationDelete**](docs/GlossaryTermTranslationsApi.md#glossarytermtranslationdelete) | **Delete** /accounts/{account_id}/glossaries/{glossary_id}/terms/{term_id}/translations/{id} | Delete a translation for a term
*GlossaryTermTranslationsApi* | [**GlossaryTermTranslationUpdate**](docs/GlossaryTermTranslationsApi.md#glossarytermtranslationupdate) | **Patch** /accounts/{account_id}/glossaries/{glossary_id}/terms/{term_id}/translations/{id} | Update a translation for a term
*GlossaryTermsApi* | [**GlossaryTermCreate**](docs/GlossaryTermsApi.md#glossarytermcreate) | **Post** /accounts/{account_id}/glossaries/{glossary_id}/terms | Create a term
*GlossaryTermsApi* | [**GlossaryTermDelete**](docs/GlossaryTermsApi.md#glossarytermdelete) | **Delete** /accounts/{account_id}/glossaries/{glossary_id}/terms/{id} | Delete a term
*GlossaryTermsApi* | [**GlossaryTermShow**](docs/GlossaryTermsApi.md#glossarytermshow) | **Get** /accounts/{account_id}/glossaries/{glossary_id}/terms/{id} | Get a single term
*GlossaryTermsApi* | [**GlossaryTermUpdate**](docs/GlossaryTermsApi.md#glossarytermupdate) | **Patch** /accounts/{account_id}/glossaries/{glossary_id}/terms/{id} | Update a term
*GlossaryTermsApi* | [**GlossaryTermsList**](docs/GlossaryTermsApi.md#glossarytermslist) | **Get** /accounts/{account_id}/glossaries/{glossary_id}/terms | List terms
*ICUApi* | [**IcuSkeleton**](docs/ICUApi.md#icuskeleton) | **Post** /icu/skeleton | Build ICU skeletons
*InvitationsApi* | [**InvitationCreate**](docs/InvitationsApi.md#invitationcreate) | **Post** /accounts/{account_id}/invitations | Create a new invitation
*InvitationsApi* | [**InvitationDelete**](docs/InvitationsApi.md#invitationdelete) | **Delete** /accounts/{account_id}/invitations/{id} | Delete an invitation
*InvitationsApi* | [**InvitationResend**](docs/InvitationsApi.md#invitationresend) | **Post** /accounts/{account_id}/invitations/{id}/resend | Resend an invitation
*InvitationsApi* | [**InvitationShow**](docs/InvitationsApi.md#invitationshow) | **Get** /accounts/{account_id}/invitations/{id} | Get a single invitation
*InvitationsApi* | [**InvitationUpdate**](docs/InvitationsApi.md#invitationupdate) | **Patch** /accounts/{account_id}/invitations/{id} | Update an invitation
*InvitationsApi* | [**InvitationUpdateSettings**](docs/InvitationsApi.md#invitationupdatesettings) | **Patch** /projects/{project_id}/invitations/{id} | Update a member&#39;s invitation access
*InvitationsApi* | [**InvitationsList**](docs/InvitationsApi.md#invitationslist) | **Get** /accounts/{account_id}/invitations | List invitations
*JobCommentsApi* | [**JobCommentCreate**](docs/JobCommentsApi.md#jobcommentcreate) | **Post** /projects/{project_id}/jobs/{job_id}/comments | Create a job comment
*JobCommentsApi* | [**JobCommentDelete**](docs/JobCommentsApi.md#jobcommentdelete) | **Delete** /projects/{project_id}/jobs/{job_id}/comments/{id} | Delete a job comment
*JobCommentsApi* | [**JobCommentShow**](docs/JobCommentsApi.md#jobcommentshow) | **Get** /projects/{project_id}/jobs/{job_id}/comments/{id} | Get a single job comment
*JobCommentsApi* | [**JobCommentUpdate**](docs/JobCommentsApi.md#jobcommentupdate) | **Patch** /projects/{project_id}/jobs/{job_id}/comments/{id} | Update a job comment
*JobCommentsApi* | [**JobCommentsList**](docs/JobCommentsApi.md#jobcommentslist) | **Get** /projects/{project_id}/jobs/{job_id}/comments | List job comments
*JobLocalesApi* | [**JobLocaleComplete**](docs/JobLocalesApi.md#joblocalecomplete) | **Post** /projects/{project_id}/jobs/{job_id}/locales/{id}/complete | Complete a job locale
*JobLocalesApi* | [**JobLocaleCompleteReview**](docs/JobLocalesApi.md#joblocalecompletereview) | **Post** /projects/{project_id}/jobs/{job_id}/locales/{id}/complete_review | Review a job locale
*JobLocalesApi* | [**JobLocaleDelete**](docs/JobLocalesApi.md#joblocaledelete) | **Delete** /projects/{project_id}/jobs/{job_id}/locales/{id} | Remove a target locale from a job
*JobLocalesApi* | [**JobLocaleReopen**](docs/JobLocalesApi.md#joblocalereopen) | **Post** /projects/{project_id}/jobs/{job_id}/locales/{id}/reopen | Reopen a job locale
*JobLocalesApi* | [**JobLocaleShow**](docs/JobLocalesApi.md#joblocaleshow) | **Get** /projects/{project_id}/jobs/{job_id}/locales/{id} | Show single job target locale
*JobLocalesApi* | [**JobLocaleUpdate**](docs/JobLocalesApi.md#joblocaleupdate) | **Patch** /projects/{project_id}/jobs/{job_id}/locales/{id} | Update a job target locale
*JobLocalesApi* | [**JobLocalesCreate**](docs/JobLocalesApi.md#joblocalescreate) | **Post** /projects/{project_id}/jobs/{job_id}/locales | Add a target locale to a job
*JobLocalesApi* | [**JobLocalesList**](docs/JobLocalesApi.md#joblocaleslist) | **Get** /projects/{project_id}/jobs/{job_id}/locales | List job target locales
*JobTemplateLocalesApi* | [**JobTemplateLocaleDelete**](docs/JobTemplateLocalesApi.md#jobtemplatelocaledelete) | **Delete** /projects/{project_id}/job_templates/{job_template_id}/locales/{job_template_locale_id} | Delete a job template locale
*JobTemplateLocalesApi* | [**JobTemplateLocaleShow**](docs/JobTemplateLocalesApi.md#jobtemplatelocaleshow) | **Get** /projects/{project_id}/job_templates/{job_template_id}/locales/{job_template_locale_id} | Get a single job template locale
*JobTemplateLocalesApi* | [**JobTemplateLocaleUpdate**](docs/JobTemplateLocalesApi.md#jobtemplatelocaleupdate) | **Patch** /projects/{project_id}/job_templates/{job_template_id}/locales/{job_template_locale_id} | Update a job template locale
*JobTemplateLocalesApi* | [**JobTemplateLocalesCreate**](docs/JobTemplateLocalesApi.md#jobtemplatelocalescreate) | **Post** /projects/{project_id}/job_templates/{job_template_id}/locales | Create a job template locale
*JobTemplateLocalesApi* | [**JobTemplateLocalesList**](docs/JobTemplateLocalesApi.md#jobtemplatelocaleslist) | **Get** /projects/{project_id}/job_templates/{job_template_id}/locales | List job template locales
*JobTemplatesApi* | [**JobTemplateCreate**](docs/JobTemplatesApi.md#jobtemplatecreate) | **Post** /projects/{project_id}/job_templates | Create a job template
*JobTemplatesApi* | [**JobTemplateDelete**](docs/JobTemplatesApi.md#jobtemplatedelete) | **Delete** /projects/{project_id}/job_templates/{id} | Delete a job template
*JobTemplatesApi* | [**JobTemplateUpdate**](docs/JobTemplatesApi.md#jobtemplateupdate) | **Patch** /projects/{project_id}/job_templates/{id} | Update a job template
*JobTemplatesApi* | [**JobTemplatesList**](docs/JobTemplatesApi.md#jobtemplateslist) | **Get** /projects/{project_id}/job_templates | List job templates
*JobTemplatesApi* | [**JobTemplatesShow**](docs/JobTemplatesApi.md#jobtemplatesshow) | **Get** /projects/{project_id}/job_templates/{id} | Get a single job template
*JobsApi* | [**JobComplete**](docs/JobsApi.md#jobcomplete) | **Post** /projects/{project_id}/jobs/{id}/complete | Complete a job
*JobsApi* | [**JobCreate**](docs/JobsApi.md#jobcreate) | **Post** /projects/{project_id}/jobs | Create a job
*JobsApi* | [**JobDelete**](docs/JobsApi.md#jobdelete) | **Delete** /projects/{project_id}/jobs/{id} | Delete a job
*JobsApi* | [**JobKeysCreate**](docs/JobsApi.md#jobkeyscreate) | **Post** /projects/{project_id}/jobs/{id}/keys | Add keys to job
*JobsApi* | [**JobKeysDelete**](docs/JobsApi.md#jobkeysdelete) | **Delete** /projects/{project_id}/jobs/{id}/keys | Remove keys from job
*JobsApi* | [**JobLock**](docs/JobsApi.md#joblock) | **Post** /projects/{project_id}/jobs/{id}/lock | Lock a job
*JobsApi* | [**JobReopen**](docs/JobsApi.md#jobreopen) | **Post** /projects/{project_id}/jobs/{id}/reopen | Reopen a job
*JobsApi* | [**JobShow**](docs/JobsApi.md#jobshow) | **Get** /projects/{project_id}/jobs/{id} | Get a single job
*JobsApi* | [**JobStart**](docs/JobsApi.md#jobstart) | **Post** /projects/{project_id}/jobs/{id}/start | Start a job
*JobsApi* | [**JobUnlock**](docs/JobsApi.md#jobunlock) | **Post** /projects/{project_id}/jobs/{id}/unlock | Unlock a job
*JobsApi* | [**JobUpdate**](docs/JobsApi.md#jobupdate) | **Patch** /projects/{project_id}/jobs/{id} | Update a job
*JobsApi* | [**JobsByAccount**](docs/JobsApi.md#jobsbyaccount) | **Get** /accounts/{account_id}/jobs | List account jobs
*JobsApi* | [**JobsList**](docs/JobsApi.md#jobslist) | **Get** /projects/{project_id}/jobs | List jobs
*KeysApi* | [**KeyCreate**](docs/KeysApi.md#keycreate) | **Post** /projects/{project_id}/keys | Create a key
*KeysApi* | [**KeyDelete**](docs/KeysApi.md#keydelete) | **Delete** /projects/{project_id}/keys/{id} | Delete a key
*KeysApi* | [**KeyShow**](docs/KeysApi.md#keyshow) | **Get** /projects/{project_id}/keys/{id} | Get a single key
*KeysApi* | [**KeyUpdate**](docs/KeysApi.md#keyupdate) | **Patch** /projects/{project_id}/keys/{id} | Update a key
*KeysApi* | [**KeysDeleteCollection**](docs/KeysApi.md#keysdeletecollection) | **Delete** /projects/{project_id}/keys | Delete collection of keys
*KeysApi* | [**KeysExclude**](docs/KeysApi.md#keysexclude) | **Patch** /projects/{project_id}/keys/exclude | Exclude a locale on a collection of keys
*KeysApi* | [**KeysInclude**](docs/KeysApi.md#keysinclude) | **Patch** /projects/{project_id}/keys/include | Include a locale on a collection of keys
*KeysApi* | [**KeysList**](docs/KeysApi.md#keyslist) | **Get** /projects/{project_id}/keys | List keys
*KeysApi* | [**KeysSearch**](docs/KeysApi.md#keyssearch) | **Post** /projects/{project_id}/keys/search | Search keys
*KeysApi* | [**KeysTag**](docs/KeysApi.md#keystag) | **Patch** /projects/{project_id}/keys/tag | Add tags to collection of keys
*KeysApi* | [**KeysUntag**](docs/KeysApi.md#keysuntag) | **Patch** /projects/{project_id}/keys/untag | Remove tags from collection of keys
*KeysFigmaAttachmentsApi* | [**FigmaAttachmentAttachToKey**](docs/KeysFigmaAttachmentsApi.md#figmaattachmentattachtokey) | **Post** /projects/{project_id}/figma_attachments/{figma_attachment_id}/keys | Attach the Figma attachment to a key
*KeysFigmaAttachmentsApi* | [**FigmaAttachmentDetachFromKey**](docs/KeysFigmaAttachmentsApi.md#figmaattachmentdetachfromkey) | **Delete** /projects/{project_id}/figma_attachments/{figma_attachment_id}/keys/{id} | Detach the Figma attachment from a key
*LinkedKeysApi* | [**KeyLinksBatchDestroy**](docs/LinkedKeysApi.md#keylinksbatchdestroy) | **Delete** /projects/{project_id}/keys/{id}/key_links | Batch unlink child keys from a parent key
*LinkedKeysApi* | [**KeyLinksCreate**](docs/LinkedKeysApi.md#keylinkscreate) | **Post** /projects/{project_id}/keys/{id}/key_links | Link child keys to a parent key
*LinkedKeysApi* | [**KeyLinksDestroy**](docs/LinkedKeysApi.md#keylinksdestroy) | **Delete** /projects/{project_id}/keys/{id}/key_links/{child_key_id} | Unlink a child key from a parent key
*LinkedKeysApi* | [**KeyLinksIndex**](docs/LinkedKeysApi.md#keylinksindex) | **Get** /projects/{project_id}/keys/{id}/key_links | List child keys of a parent key
*LocalesApi* | [**AccountLocales**](docs/LocalesApi.md#accountlocales) | **Get** /accounts/{id}/locales | List locales used in account
*LocalesApi* | [**LocaleCreate**](docs/LocalesApi.md#localecreate) | **Post** /projects/{project_id}/locales | Create a locale
*LocalesApi* | [**LocaleDelete**](docs/LocalesApi.md#localedelete) | **Delete** /projects/{project_id}/locales/{id} | Delete a locale
*LocalesApi* | [**LocaleDownload**](docs/LocalesApi.md#localedownload) | **Get** /projects/{project_id}/locales/{id}/download | Download a locale
*LocalesApi* | [**LocaleShow**](docs/LocalesApi.md#localeshow) | **Get** /projects/{project_id}/locales/{id} | Get a single locale
*LocalesApi* | [**LocaleUpdate**](docs/LocalesApi.md#localeupdate) | **Patch** /projects/{project_id}/locales/{id} | Update a locale
*LocalesApi* | [**LocalesList**](docs/LocalesApi.md#localeslist) | **Get** /projects/{project_id}/locales | List locales
*MembersApi* | [**MemberDelete**](docs/MembersApi.md#memberdelete) | **Delete** /accounts/{account_id}/members/{id} | Remove a user from the account
*MembersApi* | [**MemberShow**](docs/MembersApi.md#membershow) | **Get** /accounts/{account_id}/members/{id} | Get single member
*MembersApi* | [**MemberUpdate**](docs/MembersApi.md#memberupdate) | **Patch** /accounts/{account_id}/members/{id} | Update a member
*MembersApi* | [**MemberUpdateSettings**](docs/MembersApi.md#memberupdatesettings) | **Patch** /projects/{project_id}/members/{id} | Update a member&#39;s project settings
*MembersApi* | [**MembersList**](docs/MembersApi.md#memberslist) | **Get** /accounts/{account_id}/members | List members
*NotificationGroupsApi* | [**NotificationGroupsList**](docs/NotificationGroupsApi.md#notificationgroupslist) | **Get** /notification_groups | List notification groups
*NotificationGroupsApi* | [**NotificationGroupsMarkAllAsRead**](docs/NotificationGroupsApi.md#notificationgroupsmarkallasread) | **Patch** /notification_groups/mark_all_as_read | Mark all notification groups as read
*NotificationGroupsApi* | [**NotificationGroupsMarkAsRead**](docs/NotificationGroupsApi.md#notificationgroupsmarkasread) | **Patch** /notification_groups/{id}/mark_as_read | Mark a notification group as read
*NotificationsApi* | [**NotificationsList**](docs/NotificationsApi.md#notificationslist) | **Get** /notifications | List notifications
*NotificationsApi* | [**NotificationsMarkAllAsRead**](docs/NotificationsApi.md#notificationsmarkallasread) | **Post** /notifications/mark_all_as_read | Mark all notifications as read
*NotificationsApi* | [**NotificationsShow**](docs/NotificationsApi.md#notificationsshow) | **Get** /notifications/{id} | Get a single notification
*OrdersApi* | [**OrderConfirm**](docs/OrdersApi.md#orderconfirm) | **Patch** /projects/{project_id}/orders/{id}/confirm | Confirm an order
*OrdersApi* | [**OrderCreate**](docs/OrdersApi.md#ordercreate) | **Post** /projects/{project_id}/orders | Create a new order
*OrdersApi* | [**OrderDelete**](docs/OrdersApi.md#orderdelete) | **Delete** /projects/{project_id}/orders/{id} | Cancel an order
*OrdersApi* | [**OrderShow**](docs/OrdersApi.md#ordershow) | **Get** /projects/{project_id}/orders/{id} | Get a single order
*OrdersApi* | [**OrdersList**](docs/OrdersApi.md#orderslist) | **Get** /projects/{project_id}/orders | List orders
*OrganizationJobTemplateLocalesApi* | [**OrganizationJobTemplateLocaleDelete**](docs/OrganizationJobTemplateLocalesApi.md#organizationjobtemplatelocaledelete) | **Delete** /accounts/{account_id}/job_templates/{job_template_id}/locales/{job_template_locale_id} | Delete an organization job template locale
*OrganizationJobTemplateLocalesApi* | [**OrganizationJobTemplateLocaleShow**](docs/OrganizationJobTemplateLocalesApi.md#organizationjobtemplatelocaleshow) | **Get** /accounts/{account_id}/job_templates/{job_template_id}/locales/{job_template_locale_id} | Get a single organization job template locale
*OrganizationJobTemplateLocalesApi* | [**OrganizationJobTemplateLocaleUpdate**](docs/OrganizationJobTemplateLocalesApi.md#organizationjobtemplatelocaleupdate) | **Patch** /accounts/{account_id}/job_templates/{job_template_id}/locales/{job_template_locale_id} | Update an organization job template locale
*OrganizationJobTemplateLocalesApi* | [**OrganizationJobTemplateLocalesCreate**](docs/OrganizationJobTemplateLocalesApi.md#organizationjobtemplatelocalescreate) | **Post** /accounts/{account_id}/job_templates/{job_template_id}/locales | Create an organization job template locale
*OrganizationJobTemplateLocalesApi* | [**OrganizationJobTemplateLocalesList**](docs/OrganizationJobTemplateLocalesApi.md#organizationjobtemplatelocaleslist) | **Get** /accounts/{account_id}/job_templates/{job_template_id}/locales | List organization job template locales
*OrganizationJobTemplatesApi* | [**OrganizationJobTemplateCreate**](docs/OrganizationJobTemplatesApi.md#organizationjobtemplatecreate) | **Post** /accounts/{account_id}/job_templates | Create an organization job template
*OrganizationJobTemplatesApi* | [**OrganizationJobTemplateDelete**](docs/OrganizationJobTemplatesApi.md#organizationjobtemplatedelete) | **Delete** /accounts/{account_id}/job_templates/{id} | Delete an organization job template
*OrganizationJobTemplatesApi* | [**OrganizationJobTemplateUpdate**](docs/OrganizationJobTemplatesApi.md#organizationjobtemplateupdate) | **Patch** /accounts/{account_id}/job_templates/{id} | Update an organization job template
*OrganizationJobTemplatesApi* | [**OrganizationJobTemplatesList**](docs/OrganizationJobTemplatesApi.md#organizationjobtemplateslist) | **Get** /accounts/{account_id}/job_templates | List organization job templates
*OrganizationJobTemplatesApi* | [**OrganizationJobTemplatesShow**](docs/OrganizationJobTemplatesApi.md#organizationjobtemplatesshow) | **Get** /accounts/{account_id}/job_templates/{id} | Get a single organization job template
*ProjectsApi* | [**ProjectCreate**](docs/ProjectsApi.md#projectcreate) | **Post** /projects | Create a project
*ProjectsApi* | [**ProjectDelete**](docs/ProjectsApi.md#projectdelete) | **Delete** /projects/{id} | Delete a project
*ProjectsApi* | [**ProjectShow**](docs/ProjectsApi.md#projectshow) | **Get** /projects/{id} | Get a single project
*ProjectsApi* | [**ProjectUpdate**](docs/ProjectsApi.md#projectupdate) | **Patch** /projects/{id} | Update a project
*ProjectsApi* | [**ProjectsList**](docs/ProjectsApi.md#projectslist) | **Get** /projects | List projects
*QualityPerformanceScoreApi* | [**ProjectsQualityPerformanceScore**](docs/QualityPerformanceScoreApi.md#projectsqualityperformancescore) | **Post** /projects/{project_id}/quality_performance_score | Get Translation Quality
*ReleasesApi* | [**ReleaseCreate**](docs/ReleasesApi.md#releasecreate) | **Post** /accounts/{account_id}/distributions/{distribution_id}/releases | Create a release
*ReleasesApi* | [**ReleaseDelete**](docs/ReleasesApi.md#releasedelete) | **Delete** /accounts/{account_id}/distributions/{distribution_id}/releases/{id} | Delete a release
*ReleasesApi* | [**ReleasePublish**](docs/ReleasesApi.md#releasepublish) | **Post** /accounts/{account_id}/distributions/{distribution_id}/releases/{id}/publish | Publish a release
*ReleasesApi* | [**ReleaseShow**](docs/ReleasesApi.md#releaseshow) | **Get** /accounts/{account_id}/distributions/{distribution_id}/releases/{id} | Get a single release
*ReleasesApi* | [**ReleaseUpdate**](docs/ReleasesApi.md#releaseupdate) | **Patch** /accounts/{account_id}/distributions/{distribution_id}/releases/{id} | Update a release
*ReleasesApi* | [**ReleasesList**](docs/ReleasesApi.md#releaseslist) | **Get** /accounts/{account_id}/distributions/{distribution_id}/releases | List releases
*RepoSyncsApi* | [**RepoSyncActivate**](docs/RepoSyncsApi.md#reposyncactivate) | **Post** /accounts/{account_id}/repo_syncs/{repo_sync_id}/activate | Activate a Repo Sync
*RepoSyncsApi* | [**RepoSyncDeactivate**](docs/RepoSyncsApi.md#reposyncdeactivate) | **Post** /accounts/{account_id}/repo_syncs/{repo_sync_id}/deactivate | Deactivate a Repo Sync
*RepoSyncsApi* | [**RepoSyncEvents**](docs/RepoSyncsApi.md#reposyncevents) | **Get** /accounts/{account_id}/repo_syncs/{repo_sync_id}/events | Repository Syncs History
*RepoSyncsApi* | [**RepoSyncExport**](docs/RepoSyncsApi.md#reposyncexport) | **Post** /accounts/{account_id}/repo_syncs/{repo_sync_id}/export | Export to code repository
*RepoSyncsApi* | [**RepoSyncImport**](docs/RepoSyncsApi.md#reposyncimport) | **Post** /accounts/{account_id}/repo_syncs/{repo_sync_id}/import | Import from code repository
*RepoSyncsApi* | [**RepoSyncList**](docs/RepoSyncsApi.md#reposynclist) | **Get** /accounts/{account_id}/repo_syncs | Get Repo Syncs
*RepoSyncsApi* | [**RepoSyncShow**](docs/RepoSyncsApi.md#reposyncshow) | **Get** /accounts/{account_id}/repo_syncs/{repo_sync_id} | Get a single Repo Sync
*ReportsApi* | [**ReportLocalesList**](docs/ReportsApi.md#reportlocaleslist) | **Get** /projects/{project_id}/report/locales | List Locale Reports
*ReportsApi* | [**ReportShow**](docs/ReportsApi.md#reportshow) | **Get** /projects/{project_id}/report | Get Project Report
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
*SearchApi* | [**SearchInAccount**](docs/SearchApi.md#searchinaccount) | **Post** /accounts/{account_id}/search | Search across projects
*SpacesApi* | [**SpaceCreate**](docs/SpacesApi.md#spacecreate) | **Post** /accounts/{account_id}/spaces | Create a Space
*SpacesApi* | [**SpaceDelete**](docs/SpacesApi.md#spacedelete) | **Delete** /accounts/{account_id}/spaces/{id} | Delete Space
*SpacesApi* | [**SpaceShow**](docs/SpacesApi.md#spaceshow) | **Get** /accounts/{account_id}/spaces/{id} | Get Space
*SpacesApi* | [**SpaceUpdate**](docs/SpacesApi.md#spaceupdate) | **Patch** /accounts/{account_id}/spaces/{id} | Update Space
*SpacesApi* | [**SpacesList**](docs/SpacesApi.md#spaceslist) | **Get** /accounts/{account_id}/spaces | List Spaces
*SpacesApi* | [**SpacesProjectsCreate**](docs/SpacesApi.md#spacesprojectscreate) | **Post** /accounts/{account_id}/spaces/{space_id}/projects | Add Project to Space
*SpacesApi* | [**SpacesProjectsDelete**](docs/SpacesApi.md#spacesprojectsdelete) | **Delete** /accounts/{account_id}/spaces/{space_id}/projects/{id} | Remove Project from Space
*SpacesApi* | [**SpacesProjectsList**](docs/SpacesApi.md#spacesprojectslist) | **Get** /accounts/{account_id}/spaces/{space_id}/projects | List Projects in Space
*StyleGuidesApi* | [**StyleguideCreate**](docs/StyleGuidesApi.md#styleguidecreate) | **Post** /projects/{project_id}/styleguides | Create a style guide
*StyleGuidesApi* | [**StyleguideDelete**](docs/StyleGuidesApi.md#styleguidedelete) | **Delete** /projects/{project_id}/styleguides/{id} | Delete a style guide
*StyleGuidesApi* | [**StyleguideShow**](docs/StyleGuidesApi.md#styleguideshow) | **Get** /projects/{project_id}/styleguides/{id} | Get a single style guide
*StyleGuidesApi* | [**StyleguideUpdate**](docs/StyleGuidesApi.md#styleguideupdate) | **Patch** /projects/{project_id}/styleguides/{id} | Update a style guide
*StyleGuidesApi* | [**StyleguidesList**](docs/StyleGuidesApi.md#styleguideslist) | **Get** /projects/{project_id}/styleguides | List style guides
*TagsApi* | [**TagCreate**](docs/TagsApi.md#tagcreate) | **Post** /projects/{project_id}/tags | Create a tag
*TagsApi* | [**TagDelete**](docs/TagsApi.md#tagdelete) | **Delete** /projects/{project_id}/tags/{name} | Delete a tag
*TagsApi* | [**TagShow**](docs/TagsApi.md#tagshow) | **Get** /projects/{project_id}/tags/{name} | Get a single tag
*TagsApi* | [**TagsList**](docs/TagsApi.md#tagslist) | **Get** /projects/{project_id}/tags | List tags
*TeamsApi* | [**TeamCreate**](docs/TeamsApi.md#teamcreate) | **Post** /accounts/{account_id}/teams | Create a Team
*TeamsApi* | [**TeamDelete**](docs/TeamsApi.md#teamdelete) | **Delete** /accounts/{account_id}/teams/{id} | Delete Team
*TeamsApi* | [**TeamShow**](docs/TeamsApi.md#teamshow) | **Get** /accounts/{account_id}/teams/{id} | Get Team
*TeamsApi* | [**TeamUpdate**](docs/TeamsApi.md#teamupdate) | **Patch** /accounts/{account_id}/teams/{id} | Update Team
*TeamsApi* | [**TeamsList**](docs/TeamsApi.md#teamslist) | **Get** /accounts/{account_id}/teams | List Teams
*TeamsApi* | [**TeamsProjectsCreate**](docs/TeamsApi.md#teamsprojectscreate) | **Post** /accounts/{account_id}/teams/{team_id}/projects | Add Project to Team
*TeamsApi* | [**TeamsProjectsDelete**](docs/TeamsApi.md#teamsprojectsdelete) | **Delete** /accounts/{account_id}/teams/{team_id}/projects/{id} | Remove Project from Team
*TeamsApi* | [**TeamsSpacesCreate**](docs/TeamsApi.md#teamsspacescreate) | **Post** /accounts/{account_id}/teams/{team_id}/spaces | Add Space
*TeamsApi* | [**TeamsSpacesDelete**](docs/TeamsApi.md#teamsspacesdelete) | **Delete** /accounts/{account_id}/teams/{team_id}/spaces/{id} | Remove Space
*TeamsApi* | [**TeamsUsersCreate**](docs/TeamsApi.md#teamsuserscreate) | **Post** /accounts/{account_id}/teams/{team_id}/users | Add User
*TeamsApi* | [**TeamsUsersDelete**](docs/TeamsApi.md#teamsusersdelete) | **Delete** /accounts/{account_id}/teams/{team_id}/users/{id} | Remove User
*TranslationsApi* | [**TranslationCreate**](docs/TranslationsApi.md#translationcreate) | **Post** /projects/{project_id}/translations | Create a translation
*TranslationsApi* | [**TranslationExclude**](docs/TranslationsApi.md#translationexclude) | **Patch** /projects/{project_id}/translations/{id}/exclude | Exclude a translation from export
*TranslationsApi* | [**TranslationInclude**](docs/TranslationsApi.md#translationinclude) | **Patch** /projects/{project_id}/translations/{id}/include | Include a translation
*TranslationsApi* | [**TranslationReview**](docs/TranslationsApi.md#translationreview) | **Patch** /projects/{project_id}/translations/{id}/review | Review a translation
*TranslationsApi* | [**TranslationShow**](docs/TranslationsApi.md#translationshow) | **Get** /projects/{project_id}/translations/{id} | Get a single translation
*TranslationsApi* | [**TranslationUnverify**](docs/TranslationsApi.md#translationunverify) | **Patch** /projects/{project_id}/translations/{id}/unverify | Mark a translation as unverified
*TranslationsApi* | [**TranslationUpdate**](docs/TranslationsApi.md#translationupdate) | **Patch** /projects/{project_id}/translations/{id} | Update a translation
*TranslationsApi* | [**TranslationVerify**](docs/TranslationsApi.md#translationverify) | **Patch** /projects/{project_id}/translations/{id}/verify | Verify a translation
*TranslationsApi* | [**TranslationsByKey**](docs/TranslationsApi.md#translationsbykey) | **Get** /projects/{project_id}/keys/{key_id}/translations | List translations by key
*TranslationsApi* | [**TranslationsByLocale**](docs/TranslationsApi.md#translationsbylocale) | **Get** /projects/{project_id}/locales/{locale_id}/translations | List translations by locale
*TranslationsApi* | [**TranslationsExcludeCollection**](docs/TranslationsApi.md#translationsexcludecollection) | **Patch** /projects/{project_id}/translations/exclude | Exclude translations by query
*TranslationsApi* | [**TranslationsIncludeCollection**](docs/TranslationsApi.md#translationsincludecollection) | **Patch** /projects/{project_id}/translations/include | Include translations by query
*TranslationsApi* | [**TranslationsList**](docs/TranslationsApi.md#translationslist) | **Get** /projects/{project_id}/translations | List all translations
*TranslationsApi* | [**TranslationsReviewCollection**](docs/TranslationsApi.md#translationsreviewcollection) | **Patch** /projects/{project_id}/translations/review | Review translations selected by query
*TranslationsApi* | [**TranslationsSearch**](docs/TranslationsApi.md#translationssearch) | **Post** /projects/{project_id}/translations/search | Search translations
*TranslationsApi* | [**TranslationsUnverifyCollection**](docs/TranslationsApi.md#translationsunverifycollection) | **Patch** /projects/{project_id}/translations/unverify | Unverify translations by query
*TranslationsApi* | [**TranslationsVerifyCollection**](docs/TranslationsApi.md#translationsverifycollection) | **Patch** /projects/{project_id}/translations/verify | Verify translations by query
*UploadsApi* | [**UploadCreate**](docs/UploadsApi.md#uploadcreate) | **Post** /projects/{project_id}/uploads | Upload a new file
*UploadsApi* | [**UploadShow**](docs/UploadsApi.md#uploadshow) | **Get** /projects/{project_id}/uploads/{id} | Get a single upload
*UploadsApi* | [**UploadsList**](docs/UploadsApi.md#uploadslist) | **Get** /projects/{project_id}/uploads | List uploads
*UsersApi* | [**ShowUser**](docs/UsersApi.md#showuser) | **Get** /user | Show current User
*VariablesApi* | [**VariableCreate**](docs/VariablesApi.md#variablecreate) | **Post** /projects/{project_id}/variables | Create a variable
*VariablesApi* | [**VariableDelete**](docs/VariablesApi.md#variabledelete) | **Delete** /projects/{project_id}/variables/{name} | Delete a variable
*VariablesApi* | [**VariableShow**](docs/VariablesApi.md#variableshow) | **Get** /projects/{project_id}/variables/{name} | Get a single variable
*VariablesApi* | [**VariableUpdate**](docs/VariablesApi.md#variableupdate) | **Patch** /projects/{project_id}/variables/{name} | Update a variable
*VariablesApi* | [**VariablesList**](docs/VariablesApi.md#variableslist) | **Get** /projects/{project_id}/variables | List variables
*VersionsHistoryApi* | [**VersionShow**](docs/VersionsHistoryApi.md#versionshow) | **Get** /projects/{project_id}/translations/{translation_id}/versions/{id} | Get a single version
*VersionsHistoryApi* | [**VersionsList**](docs/VersionsHistoryApi.md#versionslist) | **Get** /projects/{project_id}/translations/{translation_id}/versions | List all versions
*WebhookDeliveriesApi* | [**WebhookDeliveriesList**](docs/WebhookDeliveriesApi.md#webhookdeliverieslist) | **Get** /projects/{project_id}/webhooks/{webhook_id}/deliveries | List webhook deliveries
*WebhookDeliveriesApi* | [**WebhookDeliveriesRedeliver**](docs/WebhookDeliveriesApi.md#webhookdeliveriesredeliver) | **Post** /projects/{project_id}/webhooks/{webhook_id}/deliveries/{id}/redeliver | Redeliver a single webhook delivery
*WebhookDeliveriesApi* | [**WebhookDeliveriesShow**](docs/WebhookDeliveriesApi.md#webhookdeliveriesshow) | **Get** /projects/{project_id}/webhooks/{webhook_id}/deliveries/{id} | Get a single webhook delivery
*WebhooksApi* | [**WebhookCreate**](docs/WebhooksApi.md#webhookcreate) | **Post** /projects/{project_id}/webhooks | Create a webhook
*WebhooksApi* | [**WebhookDelete**](docs/WebhooksApi.md#webhookdelete) | **Delete** /projects/{project_id}/webhooks/{id} | Delete a webhook
*WebhooksApi* | [**WebhookShow**](docs/WebhooksApi.md#webhookshow) | **Get** /projects/{project_id}/webhooks/{id} | Get a single webhook
*WebhooksApi* | [**WebhookTest**](docs/WebhooksApi.md#webhooktest) | **Post** /projects/{project_id}/webhooks/{id}/test | Test a webhook
*WebhooksApi* | [**WebhookUpdate**](docs/WebhooksApi.md#webhookupdate) | **Patch** /projects/{project_id}/webhooks/{id} | Update a webhook
*WebhooksApi* | [**WebhooksList**](docs/WebhooksApi.md#webhookslist) | **Get** /projects/{project_id}/webhooks | List webhooks


## Documentation For Models

 - [Account](docs/Account.md)
 - [AccountDetails](docs/AccountDetails.md)
 - [AccountSearchResult](docs/AccountSearchResult.md)
 - [AffectedCount](docs/AffectedCount.md)
 - [AffectedResources](docs/AffectedResources.md)
 - [Authorization](docs/Authorization.md)
 - [AuthorizationCreateParameters](docs/AuthorizationCreateParameters.md)
 - [AuthorizationUpdateParameters](docs/AuthorizationUpdateParameters.md)
 - [AuthorizationWithToken](docs/AuthorizationWithToken.md)
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
 - [BranchName](docs/BranchName.md)
 - [BranchUpdateParameters](docs/BranchUpdateParameters.md)
 - [Comment](docs/Comment.md)
 - [CommentCreateParameters](docs/CommentCreateParameters.md)
 - [CommentMarkReadParameters](docs/CommentMarkReadParameters.md)
 - [CommentReaction](docs/CommentReaction.md)
 - [CommentUpdateParameters](docs/CommentUpdateParameters.md)
 - [CommentsListParameters](docs/CommentsListParameters.md)
 - [CurrentUser](docs/CurrentUser.md)
 - [CustomMetadataDataType](docs/CustomMetadataDataType.md)
 - [CustomMetadataPropertiesCreateParameters](docs/CustomMetadataPropertiesCreateParameters.md)
 - [CustomMetadataPropertiesUpdateParameters](docs/CustomMetadataPropertiesUpdateParameters.md)
 - [CustomMetadataProperty](docs/CustomMetadataProperty.md)
 - [CustomMetadataPropertyCreate422Response](docs/CustomMetadataPropertyCreate422Response.md)
 - [CustomMetadataPropertyCreate422ResponseErrorsInner](docs/CustomMetadataPropertyCreate422ResponseErrorsInner.md)
 - [Distribution](docs/Distribution.md)
 - [DistributionCreateParameters](docs/DistributionCreateParameters.md)
 - [DistributionPreview](docs/DistributionPreview.md)
 - [DistributionUpdateParameters](docs/DistributionUpdateParameters.md)
 - [Document](docs/Document.md)
 - [ErrorError](docs/ErrorError.md)
 - [FigmaAttachment](docs/FigmaAttachment.md)
 - [FigmaAttachmentCreateParameters](docs/FigmaAttachmentCreateParameters.md)
 - [FigmaAttachmentUpdateParameters](docs/FigmaAttachmentUpdateParameters.md)
 - [Format](docs/Format.md)
 - [GithubSyncExportParameters](docs/GithubSyncExportParameters.md)
 - [GithubSyncImportParameters](docs/GithubSyncImportParameters.md)
 - [GitlabSync](docs/GitlabSync.md)
 - [GitlabSyncExport](docs/GitlabSyncExport.md)
 - [GitlabSyncExportParameters](docs/GitlabSyncExportParameters.md)
 - [GitlabSyncHistory](docs/GitlabSyncHistory.md)
 - [GitlabSyncHistoryErrorsInner](docs/GitlabSyncHistoryErrorsInner.md)
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
 - [Icu](docs/Icu.md)
 - [IcuSkeletonParameters](docs/IcuSkeletonParameters.md)
 - [Invitation](docs/Invitation.md)
 - [InvitationCreateParameters](docs/InvitationCreateParameters.md)
 - [InvitationUpdateParameters](docs/InvitationUpdateParameters.md)
 - [InvitationUpdateSettingsParameters](docs/InvitationUpdateSettingsParameters.md)
 - [Job](docs/Job.md)
 - [JobComment](docs/JobComment.md)
 - [JobCommentCreateParameters](docs/JobCommentCreateParameters.md)
 - [JobCommentUpdateParameters](docs/JobCommentUpdateParameters.md)
 - [JobCompleteParameters](docs/JobCompleteParameters.md)
 - [JobCreateParameters](docs/JobCreateParameters.md)
 - [JobDetails](docs/JobDetails.md)
 - [JobKeysCreateParameters](docs/JobKeysCreateParameters.md)
 - [JobLocale](docs/JobLocale.md)
 - [JobLocaleCompleteParameters](docs/JobLocaleCompleteParameters.md)
 - [JobLocaleCompleteReviewParameters](docs/JobLocaleCompleteReviewParameters.md)
 - [JobLocaleReopenParameters](docs/JobLocaleReopenParameters.md)
 - [JobLocaleUpdateParameters](docs/JobLocaleUpdateParameters.md)
 - [JobLocalesCreateParameters](docs/JobLocalesCreateParameters.md)
 - [JobPreview](docs/JobPreview.md)
 - [JobReopenParameters](docs/JobReopenParameters.md)
 - [JobStartParameters](docs/JobStartParameters.md)
 - [JobTemplate](docs/JobTemplate.md)
 - [JobTemplateCreateParameters](docs/JobTemplateCreateParameters.md)
 - [JobTemplateDetails](docs/JobTemplateDetails.md)
 - [JobTemplateLocaleUpdateParameters](docs/JobTemplateLocaleUpdateParameters.md)
 - [JobTemplateLocales](docs/JobTemplateLocales.md)
 - [JobTemplateLocalesCreateParameters](docs/JobTemplateLocalesCreateParameters.md)
 - [JobTemplatePreview](docs/JobTemplatePreview.md)
 - [JobTemplateUpdateParameters](docs/JobTemplateUpdateParameters.md)
 - [JobUpdateParameters](docs/JobUpdateParameters.md)
 - [KeyCreateParameters](docs/KeyCreateParameters.md)
 - [KeyLink](docs/KeyLink.md)
 - [KeyLinksBatchDestroyParameters](docs/KeyLinksBatchDestroyParameters.md)
 - [KeyLinksCreateParameters](docs/KeyLinksCreateParameters.md)
 - [KeyLinksIndex400Response](docs/KeyLinksIndex400Response.md)
 - [KeyPreview](docs/KeyPreview.md)
 - [KeyUpdateParameters](docs/KeyUpdateParameters.md)
 - [KeysExcludeParameters](docs/KeysExcludeParameters.md)
 - [KeysIncludeParameters](docs/KeysIncludeParameters.md)
 - [KeysSearchParameters](docs/KeysSearchParameters.md)
 - [KeysTagParameters](docs/KeysTagParameters.md)
 - [KeysUntagParameters](docs/KeysUntagParameters.md)
 - [Locale](docs/Locale.md)
 - [LocaleCreateParameters](docs/LocaleCreateParameters.md)
 - [LocaleDetails](docs/LocaleDetails.md)
 - [LocalePreview](docs/LocalePreview.md)
 - [LocalePreview1](docs/LocalePreview1.md)
 - [LocaleReport](docs/LocaleReport.md)
 - [LocaleStatistics](docs/LocaleStatistics.md)
 - [LocaleTeamPreview](docs/LocaleTeamPreview.md)
 - [LocaleUpdateParameters](docs/LocaleUpdateParameters.md)
 - [LocaleUserPreview](docs/LocaleUserPreview.md)
 - [Member](docs/Member.md)
 - [MemberProjectDetail](docs/MemberProjectDetail.md)
 - [MemberProjectDetailProjectRolesInner](docs/MemberProjectDetailProjectRolesInner.md)
 - [MemberSpacesInner](docs/MemberSpacesInner.md)
 - [MemberUpdateParameters](docs/MemberUpdateParameters.md)
 - [MemberUpdateSettingsParameters](docs/MemberUpdateSettingsParameters.md)
 - [ModelError](docs/ModelError.md)
 - [Notification](docs/Notification.md)
 - [NotificationGroup](docs/NotificationGroup.md)
 - [NotificationGroupDetail](docs/NotificationGroupDetail.md)
 - [OrderConfirmParameters](docs/OrderConfirmParameters.md)
 - [OrderCreateParameters](docs/OrderCreateParameters.md)
 - [OrganizationJobTemplate](docs/OrganizationJobTemplate.md)
 - [OrganizationJobTemplateCreateParameters](docs/OrganizationJobTemplateCreateParameters.md)
 - [OrganizationJobTemplateDetails](docs/OrganizationJobTemplateDetails.md)
 - [OrganizationJobTemplateLocaleUpdateParameters](docs/OrganizationJobTemplateLocaleUpdateParameters.md)
 - [OrganizationJobTemplateLocalesCreateParameters](docs/OrganizationJobTemplateLocalesCreateParameters.md)
 - [OrganizationJobTemplateUpdateParameters](docs/OrganizationJobTemplateUpdateParameters.md)
 - [Project](docs/Project.md)
 - [ProjectCreateParameters](docs/ProjectCreateParameters.md)
 - [ProjectDetails](docs/ProjectDetails.md)
 - [ProjectLocales](docs/ProjectLocales.md)
 - [ProjectReport](docs/ProjectReport.md)
 - [ProjectShort](docs/ProjectShort.md)
 - [ProjectUpdateParameters](docs/ProjectUpdateParameters.md)
 - [ProjectsQualityPerformanceScore200Response](docs/ProjectsQualityPerformanceScore200Response.md)
 - [ProjectsQualityPerformanceScore200ResponseAnyOf](docs/ProjectsQualityPerformanceScore200ResponseAnyOf.md)
 - [ProjectsQualityPerformanceScore200ResponseAnyOfData](docs/ProjectsQualityPerformanceScore200ResponseAnyOfData.md)
 - [ProjectsQualityPerformanceScore200ResponseAnyOfDataTranslationsInner](docs/ProjectsQualityPerformanceScore200ResponseAnyOfDataTranslationsInner.md)
 - [ProjectsQualityPerformanceScore200ResponseAnyOfErrorsInner](docs/ProjectsQualityPerformanceScore200ResponseAnyOfErrorsInner.md)
 - [ProjectsQualityPerformanceScoreRequest](docs/ProjectsQualityPerformanceScoreRequest.md)
 - [Release](docs/Release.md)
 - [ReleaseCreateParameters](docs/ReleaseCreateParameters.md)
 - [ReleasePreview](docs/ReleasePreview.md)
 - [ReleaseUpdateParameters](docs/ReleaseUpdateParameters.md)
 - [RepliesListParameters](docs/RepliesListParameters.md)
 - [RepoSync](docs/RepoSync.md)
 - [RepoSyncEvent](docs/RepoSyncEvent.md)
 - [RepoSyncEventErrorsInner](docs/RepoSyncEventErrorsInner.md)
 - [RepoSyncExport](docs/RepoSyncExport.md)
 - [RepoSyncImport](docs/RepoSyncImport.md)
 - [Screenshot](docs/Screenshot.md)
 - [ScreenshotMarker](docs/ScreenshotMarker.md)
 - [ScreenshotMarkerCreateParameters](docs/ScreenshotMarkerCreateParameters.md)
 - [ScreenshotMarkerUpdateParameters](docs/ScreenshotMarkerUpdateParameters.md)
 - [ScreenshotUpdateParameters](docs/ScreenshotUpdateParameters.md)
 - [SearchInAccountParameters](docs/SearchInAccountParameters.md)
 - [Space](docs/Space.md)
 - [Space1](docs/Space1.md)
 - [SpaceCreateParameters](docs/SpaceCreateParameters.md)
 - [SpaceUpdateParameters](docs/SpaceUpdateParameters.md)
 - [SpacesProjectsCreateParameters](docs/SpacesProjectsCreateParameters.md)
 - [Styleguide](docs/Styleguide.md)
 - [StyleguideCreateParameters](docs/StyleguideCreateParameters.md)
 - [StyleguideDetails](docs/StyleguideDetails.md)
 - [StyleguidePreview](docs/StyleguidePreview.md)
 - [StyleguideUpdateParameters](docs/StyleguideUpdateParameters.md)
 - [Subscription](docs/Subscription.md)
 - [Tag](docs/Tag.md)
 - [TagCreateParameters](docs/TagCreateParameters.md)
 - [TagWithStats](docs/TagWithStats.md)
 - [TagWithStats1Statistics](docs/TagWithStats1Statistics.md)
 - [TagWithStats1Statistics1](docs/TagWithStats1Statistics1.md)
 - [Team](docs/Team.md)
 - [TeamCreateParameters](docs/TeamCreateParameters.md)
 - [TeamDetail](docs/TeamDetail.md)
 - [TeamShort](docs/TeamShort.md)
 - [TeamUpdateParameters](docs/TeamUpdateParameters.md)
 - [TeamsProjectsCreateParameters](docs/TeamsProjectsCreateParameters.md)
 - [TeamsSpacesCreateParameters](docs/TeamsSpacesCreateParameters.md)
 - [TeamsUsersCreateParameters](docs/TeamsUsersCreateParameters.md)
 - [Translation](docs/Translation.md)
 - [TranslationCreateParameters](docs/TranslationCreateParameters.md)
 - [TranslationDetails](docs/TranslationDetails.md)
 - [TranslationExcludeParameters](docs/TranslationExcludeParameters.md)
 - [TranslationIncludeParameters](docs/TranslationIncludeParameters.md)
 - [TranslationKey](docs/TranslationKey.md)
 - [TranslationKeyDetails](docs/TranslationKeyDetails.md)
 - [TranslationOrder](docs/TranslationOrder.md)
 - [TranslationParent](docs/TranslationParent.md)
 - [TranslationReviewParameters](docs/TranslationReviewParameters.md)
 - [TranslationUnverifyParameters](docs/TranslationUnverifyParameters.md)
 - [TranslationUpdateParameters](docs/TranslationUpdateParameters.md)
 - [TranslationVerifyParameters](docs/TranslationVerifyParameters.md)
 - [TranslationVersion](docs/TranslationVersion.md)
 - [TranslationVersionWithUser](docs/TranslationVersionWithUser.md)
 - [TranslationsExcludeParameters](docs/TranslationsExcludeParameters.md)
 - [TranslationsIncludeParameters](docs/TranslationsIncludeParameters.md)
 - [TranslationsReviewParameters](docs/TranslationsReviewParameters.md)
 - [TranslationsSearchParameters](docs/TranslationsSearchParameters.md)
 - [TranslationsUnverifyParameters](docs/TranslationsUnverifyParameters.md)
 - [TranslationsVerifyParameters](docs/TranslationsVerifyParameters.md)
 - [Upload](docs/Upload.md)
 - [UploadSummary](docs/UploadSummary.md)
 - [User](docs/User.md)
 - [UserPreview](docs/UserPreview.md)
 - [Variable](docs/Variable.md)
 - [VariableCreateParameters](docs/VariableCreateParameters.md)
 - [VariableUpdateParameters](docs/VariableUpdateParameters.md)
 - [Webhook](docs/Webhook.md)
 - [WebhookCreateParameters](docs/WebhookCreateParameters.md)
 - [WebhookDelivery](docs/WebhookDelivery.md)
 - [WebhookUpdateParameters](docs/WebhookUpdateParameters.md)


## More Information
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.

- API version: 2.0.0
- Package version: 3.1.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://developers.phrase.com/api/](https://developers.phrase.com/api/)

## Author

support@phrase.com
support@phrase.com
support@phrase.com
support@phrase.com
support@phrase.com
support@phrase.com
support@phrase.com
support@phrase.com
support@phrase.com
support@phrase.com
support@phrase.com
support@phrase.com
support@phrase.com
support@phrase.com
support@phrase.com
support@phrase.com
support@phrase.com
support@phrase.com
support@phrase.com
support@phrase.com
support@phrase.com
support@phrase.com
support@phrase.com
support@phrase.com
support@phrase.com
support@phrase.com
support@phrase.com
support@phrase.com
support@phrase.com
support@phrase.com
support@phrase.com
support@phrase.com
support@phrase.com
support@phrase.com
support@phrase.com
support@phrase.com
support@phrase.com
support@phrase.com
support@phrase.com
support@phrase.com
support@phrase.com
support@phrase.com
support@phrase.com
support@phrase.com
support@phrase.com
support@phrase.com
support@phrase.com
support@phrase.com
support@phrase.com
support@phrase.com
support@phrase.com
support@phrase.com
support@phrase.com


## Get help / support

Please contact [support@phrase.com](mailto:support@phrase.com?subject=[GitHub]%20phrase-go) and we can take more direct action toward finding a solution.
