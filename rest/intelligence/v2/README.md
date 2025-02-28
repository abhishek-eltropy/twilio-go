# Go API client for openapi

This is the public Twilio REST API.

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project from the OpenAPI specs located at [twilio/twilio-oai](https://github.com/twilio/twilio-oai/tree/main/spec).  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.50.1
- Package version: 1.0.0
- Build package: com.twilio.oai.TwilioGoGenerator
For more information, please visit [https://support.twilio.com](https://support.twilio.com)

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import "./openapi"
```

## Documentation for API Endpoints

All URIs are relative to *https://intelligence.twilio.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ServicesApi* | [**CreateService**](docs/ServicesApi.md#createservice) | **Post** /v2/Services | 
*ServicesApi* | [**DeleteService**](docs/ServicesApi.md#deleteservice) | **Delete** /v2/Services/{Sid} | 
*ServicesApi* | [**FetchService**](docs/ServicesApi.md#fetchservice) | **Get** /v2/Services/{Sid} | 
*ServicesApi* | [**ListService**](docs/ServicesApi.md#listservice) | **Get** /v2/Services | 
*ServicesApi* | [**UpdateService**](docs/ServicesApi.md#updateservice) | **Post** /v2/Services/{Sid} | 
*TranscriptsApi* | [**CreateTranscript**](docs/TranscriptsApi.md#createtranscript) | **Post** /v2/Transcripts | 
*TranscriptsApi* | [**DeleteTranscript**](docs/TranscriptsApi.md#deletetranscript) | **Delete** /v2/Transcripts/{Sid} | 
*TranscriptsApi* | [**FetchTranscript**](docs/TranscriptsApi.md#fetchtranscript) | **Get** /v2/Transcripts/{Sid} | 
*TranscriptsApi* | [**ListTranscript**](docs/TranscriptsApi.md#listtranscript) | **Get** /v2/Transcripts | 
*TranscriptsMediaApi* | [**FetchMedia**](docs/TranscriptsMediaApi.md#fetchmedia) | **Get** /v2/Transcripts/{Sid}/Media | 
*TranscriptsOperatorResultsApi* | [**FetchOperatorResult**](docs/TranscriptsOperatorResultsApi.md#fetchoperatorresult) | **Get** /v2/Transcripts/{TranscriptSid}/OperatorResults/{OperatorSid} | 
*TranscriptsOperatorResultsApi* | [**ListOperatorResult**](docs/TranscriptsOperatorResultsApi.md#listoperatorresult) | **Get** /v2/Transcripts/{TranscriptSid}/OperatorResults | 
*TranscriptsSentencesApi* | [**ListSentence**](docs/TranscriptsSentencesApi.md#listsentence) | **Get** /v2/Transcripts/{TranscriptSid}/Sentences | 


## Documentation For Models

 - [ListOperatorResultResponseMeta](docs/ListOperatorResultResponseMeta.md)
 - [IntelligenceV2OperatorResult](docs/IntelligenceV2OperatorResult.md)
 - [IntelligenceV2Transcript](docs/IntelligenceV2Transcript.md)
 - [ListTranscriptResponse](docs/ListTranscriptResponse.md)
 - [IntelligenceV2Service](docs/IntelligenceV2Service.md)
 - [IntelligenceV2Sentence](docs/IntelligenceV2Sentence.md)
 - [ListSentenceResponse](docs/ListSentenceResponse.md)
 - [IntelligenceV2Media](docs/IntelligenceV2Media.md)
 - [ListOperatorResultResponse](docs/ListOperatorResultResponse.md)
 - [ListServiceResponse](docs/ListServiceResponse.md)


## Documentation For Authorization



## accountSid_authToken

- **Type**: HTTP basic authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextBasicAuth, sw.BasicAuth{
    UserName: "username",
    Password: "password",
})
r, err := client.Service.Operation(auth, args)
```

