# IntelligenceV2Transcript

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The unique SID identifier of the Account. |
**ServiceSid** | Pointer to **string** | The unique SID identifier of the Service. |
**Sid** | Pointer to **string** | A 34 character string that uniquely identifies this Transcript. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date that this Transcript was created, given in ISO 8601 format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date that this Transcript was updated, given in ISO 8601 format. |
**Status** | Pointer to [**string**](TranscriptEnumStatus.md) |  |
**Channel** | Pointer to **interface{}** | Media Channel describing Transcript Source and Participant Mapping |
**DataLogging** | Pointer to **bool** | Data logging allows Twilio to improve the quality of the speech recognition through using customer data to refine its speech recognition models. |
**LanguageCode** | Pointer to **string** | The default language code of the audio. |
**CustomerKey** | Pointer to **string** |  |
**MediaStartTime** | Pointer to [**time.Time**](time.Time.md) | The date that this Transcript's media was started, given in ISO 8601 format. |
**Duration** | Pointer to **int** | The duration of this Transcript's source |
**Url** | Pointer to **string** | The URL of this resource. |
**Redaction** | Pointer to **bool** | If the transcript has been redacted, a redacted alternative of the transcript will be available. |
**Links** | Pointer to **map[string]interface{}** |  |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


