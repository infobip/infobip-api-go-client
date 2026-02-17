# Change Log of `infobip-api-go-client`

All notable changes to the library will be documented in this file.

The format of the file is based on [Keep a Changelog](http://keepachangelog.com/)
and this library adheres to [Semantic Versioning](http://semver.org/) as mentioned in [README.md][readme] file.

## [ [3.2.0](https://github.com/infobip/infobip-api-go-client/releases/tag/3.2.0)] - 2026-02-17
### Added
* Support for [Infobip RCS API](https://www.infobip.com/docs/api/channels/rcs)
* Most recent feature set for:
    * [Infobip SMS API](https://www.infobip.com/docs/api/channels/sms)
    * [Infobip 2FA API](https://www.infobip.com/docs/api/platform/2fa)
    * [Infobip Messages API](https://www.infobip.com/docs/api/platform/messages-api)
    * [Infobip Email API](https://www.infobip.com/docs/api/channels/email)
    * [Infobip Voice API](https://www.infobip.com/docs/api/channels/voice)
    * [Infobip Moments](https://www.infobip.com/docs/api/customer-engagement/moments)
* Additional mock tests to verify the correctness of request payloads and response handling.

### Changed

* **General:**
    * Deprecated `Prefix` in client configuration when setting API key authentication, as the SDK now automatically prepends "App " to the API key, simplifying the configuration process and reducing potential errors in authentication setup. Alternatively you can provide single API key directly instead of a map.
    * Enforcing usage of HTTPS when setting up base URL in client configuration.

* **SMS:**
    * Unified existing `InboundMessage` model with `MoReport`.
    * Unified existing `InboundMessageResult` model with `ReportResponse`.
    * Extended `CreateEmailMessageRequest`, `EmailMessage` with `LandingPageId`.
    * Extended `LogsResponse` with `CursorPageInfo` to support cursor-based pagination in the Infobip SMS API.

* **2FA:**
    * Extended `ResendPinRequest`, `StartAuthenticationRequest` with `TrackDelivery`.
    * Extended `StartAuthenticationResponse` with `ExternalMessageId`.
    * Extended `StartEmailAuthenticationRequest` with `LandingPagePlaceholders`.
    * Extended `TfaMessage`, `UpdateEmailMessageRequest` with `LandingPageId`.
    * Extended `TrackReport` with `EventId`, `Sender`, `CampaignReferenceId`, `EntityId` and `ApplicationId`.

* **Messages API:**
    * Extended `ChannelsDestination` with `MessageId`.
    * Extended `DefaultMessageRequestOptions` with `MessageOrdering`.
    * Extended `MessageBodyType` with `PRODUCT`, `MIXED`, `FLOW`, `TIME_PICKER`, `ORDER_REQUEST`, `ORDER_STATUS` and `FORM` types.
    * Extended `MessageButtonType` with `DIAL_PHONE` and `SHOW_LOCATION` types.
    * Extended `MessageContent` with `SenderDisplayInfo`.
    * Extended `MessageDeliveryReporting` with `ReceiveTriggeredFailoverReports`.
    * Extended `MessageOpenUrlButton` with `PostbackData` and `OpenIn`.
    * Extended `MessageOptions` with `TransliterationCode`, `CorrelationData`, `TrafficType`, `SessionRate` and `PrimaryDevice`.
    * Extended `MoEvent` with `MessageCount` and `Metadata`.
    * Extended `MoEventContentType` with `URL`, `FLOW_RESPONSE`, `PAYMENT_RESPONSE`, `FORM_RESPONSE` and `REACTION` types.
    * Extended `OutboundTemplateChannel` with `VIBER_BM` type.
    * Changed `Messages` array type from `RequestMessagesInner` to `BaseMessage` in `Request`.

* **Email:**
    * Extended `ApiReport` with `AttemptCount` and `TimeToFirstAttempt`.
    * Extended `DeliveryReport` with `AttemptCount`, `TimeToFirstAttempt`, `CampaignReferenceId`, `EntityId` and `ApplicationId`.
    * Extended `DomainResponse` with `BlocklistConfigurationLevel`.
    * Extended `IpDetailResponse`, `IpResponse` with `IpAddresses` array of strings.

* **Voice:**
* Adjusted IVR models in script processing. Scenario scripting is now implemented as a raw string to increase usability of the feature. Scripts should be passed as strings to the IVR request model in all upcoming SDK versions. All related scripts models are removed.

    * Extended `AddExistingCallRequest` with `Role` and `CustomData`.
    * Extended `AddNewCallRequest` with `Role`.
    * Extended `Call`, `CallLog`, `CallRequest` with `ExternalId`.
    * Extended `CallEndpointType` with `WHATSAPP` and `WEBSOCKET` types.
    * Extended `CallRecordingRequest` with `Channels`.
    * Extended `CallRoutingCriteriaType` with `APPLICATION` type.
    * Extended `CallRoutingEndpointType` with `WHATSAPP` and `WEBSOCKET` types.
    * Extended `CallRoutingPhoneEndpoint`, `CallRoutingSipEndpoint` with `From` and `RingbackGeneration`.
    * Extended `CallRoutingRouteResponse` with `Status` and `Order`.
    * Changed `Text` field to `Message` in `ConferenceBroadcastWebrtcTextRequest` model due to previous endpoint `/calls/1/dialogs/{dialogId}/broadcast-webrtc-text` being sunset and migrated to `/calls/1/dialogs/{dialogId}/send-message`
    * Changed `Text` field to `Message` in `DialogBroadcastWebrtcTextRequest` model due to previous endpoint `/calls/1/conferences/{conferenceId}/broadcast-webrtc-text` being sunset and migrated to `/calls/1/conferences/{conferenceId}/send-message`
    * Changed `Platform` field to `ApplicationId` and `CallRecordings` array type to `CallRecording` in `ConferenceRecording`.
    * Changed `ComposedFiles` array type from `RecordingFile` to `PublicRecordingFile` in `ConferenceRecordingLog`, `DialogRecordingLog`, `DialogRecordingResponse`.
    * Changed `Results` array type from `ConferenceRecording` to `PublicConferenceRecording` in `ConferenceRecordingLog`.
    * Changed `Location` field type in `CreateProviderSipTrunkResponse`, `CreateRegisteredSipTrunkResponse`, `CreateStaticSipTrunkResponse`, `ProviderSipTrunkRequest`, `ProviderSipTrunkResponse`, `RegisteredSipTrunkRequest`, `RegisteredSipTrunkResponse`, `RegisteredSipTrunkUpdateRequest`, `StaticSipTrunkRequest`, `StaticSipTrunkResponse`, `StaticSipTrunkUpdateRequest`  to `string`.
    * Extended `DialCallbackResponse` with `MachineDetection`.
    * Extended `DialogLogResponse` with `HangupSource`.
    * Extended `DialogState` with `TRANSFERRING` type.
    * Extended `Language` enum with `kk-kz`, `es-ar`, `es-ar`, `uz-uz`, `mr-in`, `sw-ke` and `sw-tz` types.
    * Extended `MachineDetectionProperties` with `ConfidenceRating`.
    * Extended `MachineDetectionRequest` with `DetectionTime`.
    * Changed `MediaStreamConfigRequest` to superclass with `MediaStreamingConfigRequest` and `WebsocketEndpointConfigRequest` subclasses.
    * Changed `MediaStreamConfigResponse` to superclass with `MediaStreamingConfigResponse` and `WebsocketEndpointConfigResponse` subclasses.
    * Extended `NumberMaskingStatusRequest` with `MachineDetectionResult`.
    * Extended `Participant` with `Role`.
    * Extended `ProviderTrunkType` with `OPENAI_REALTIME` type.
    * Changed `Files` array type from `RecordingFile` to `PublicRecordingFile` in `Recording`.
    * Extended `RecordingFile` with `ExpirationTime`.
    * Changed `CustomData` to `MultiChannelMappingData` in `RecordingFile`.
    * Changed `Script` field type in `SearchResponse`, `UpdateScenarioRequest` to `string`.
    * Extended `SearchResponse`, `UpdateScenarioResponse` with `NotifyUrl`, `NotifyContentType`, `Record` and `LastUsageDate`.
    * Extended `SpeechCaptureRequest` with `TerminateOnKeyPhrase`, `CustomDictionary` and `AdvancedFormatting`.
    * Extended `Transcription` with `CustomDictionary` and `AdvancedFormatting`.
    * Extended `UpdateRequest`, `VideoMediaProperties` with `Blind`.
    * Extended `UpdateScenarioRequest` with `NotifyUrl`, `NotifyContentType`, `Record`
    * Extended `VoiceData` with `Direction`, `AnsweredBy` and `CallRecordingField`.
    * Extended `VoiceName` enum with additional voices.

### Removed

* **Email:**
    * Removed `ReturnPathAddress` field from `AddDomainRequest`, `DomainResponse`.

* **Voice:**
    * Removed `DISCONNECTED` enum value from `CallState`.
    * Removed `UNKNOWN` enum value from `DetectionResult`.
    * Removed `CreationMethod` field from `File`.
    * Removed `Record` field from `IvrMessage`.
    * Removed `JOHANNESBURG_1` from `RecordingLocation`.

## [ [3.1.2](https://github.com/infobip/infobip-api-go-client/releases/tag/3.1.2)] - 2025-07-01

## Fixed
* Added custom `UnmarshalJSON` function to the [MoEvent](pkg/infobip/models/messagesapi/mo_event.go) struct to properly handle the `callbackData` field when it contains either a string or a JSON object, ensuring correct deserialization in both cases.

## [ [3.1.1](https://github.com/infobip/infobip-api-go-client/releases/tag/3.1.1)] - 2025-04-28

## Added
* `AdditionalProperties map[string]interface{}` to the `TemplateTextBody` struct to align with the current state of the API.
* `AdditionalProperties map[string]interface{}` to the `TemplateTextHeader` struct to align with the current state of the API.

## Changed
*  **Renamed Field**: `body` ➜ `requestBody` in `ApiSubmitFormDataRequest` (`FormsApi`) struct to improve clarity and naming consistency.
*  **Fixed Incorrect Field Type**: Updated `Data` field type from `map[string]map[string]interface{}` to `map[string]interface{}` to align with the current state of the API.

## [ [3.1.0](https://github.com/infobip/infobip-api-go-client/releases/tag/3.1.0)] - 2025-01-15

⚠️ IMPORTANT NOTE: This release contains compile time breaking changes.
All changes, including breaking changes, are addressed and explained in the list bellow.
If you find out that something was not addressed properly, please submit an issue.

### Added
* Most recent [Infobip SMS API](https://www.infobip.com/docs/api/channels/sms) feature set.
* Most recent [Infobip Messages API](https://www.infobip.com/docs/api/platform/messages-api) feature set.
* Support for [Infobip Email API](https://www.infobip.com/docs/api/channels/email)
* Support for [Infobip Voice API](https://www.infobip.com/docs/api/channels/voice)
* Support for [Infobip Moments](https://www.infobip.com/docs/api/customer-engagement/moments)
* Added mock tests to verify the correctness of request payloads and response handling.

### Changed

#### Renamed Classes
* Sms:
  * **LanguageV1** renamed to **PreviewLanguage**
  * **LanguageV3** renamed to **Language**
  * **LogContent** renamed to **MessageContent**
* Messages API:
  * **MessageResponseMessageResponseDetails** renamed to **MessageResponseDetails**
  * **ResponseEnvelopeMessageResponseMessageResponseDetails** renamed to **MessageResponse**
* Tfa:
  * **CreateSmsOrVoiceMessageRequest** renamed to **CreateMessageRequest**
  * **UpdateSmsOrVoiceMessageRequest** renamed to **UpdateMessageRequest**
  * **SmsOrVoiceMessageRequest** renamed to **TfaTemplateMessage**

#### Documentation
* Updated descriptions for models.
* Fixed examples in the documentation (README.md, two-factor-authentication.md).

#### Fixes
* Fixed default function for creating instances with default discriminator values in the following components:
  - `carousel_template_phone_number_button`
  - `carousel_template_open_url_button`
  - `carousel_template_quick_reply_button`
  - `inbound_typing_starting_event`
  - `inbound_typing_stopped_event`
  - `message_authentication_request_body`
  - `message_carousel_body`
  - `message_document_body`
  - `message_image_body`
  - `message_list_body`
  - `message_open_url_button`
  - `message_reply_button`
  - `message_rick_link_body`
  - `message_text_body`
  - `message_text_header`
  - `message_video_body`
  - `mo_event_audio_content`
  - `mo_event_authentication_response_content`
  - `mo_event_button_reply_content`
  - `mo_event_document_content`
  - `mo_event_file_content`
  - `mo_event_image_content`
  - `mo_event_list_reply_content`
  - `mo_event_location_content`
  - `mo_event_subject_content`
  - `mo_event_text_content`
  - `mo_event_video_content`
  - `mo_event`
  - `outbound_typing_started_event`
  - `outbound_typing_stopped_event`
  - `template_carousel_body`
  - `template_carousel_card_image_header`
  - `template_carousel_card_video_header`
  - `template_catalog_button`
  - `template_copy_code_button`
  - `template_document_header`
  - `template_flow_button`
  - `template_image_header`
  - `template_location_header`
  - `template_multi_product_button`
  - `template_open_url_button`
  - `template_phone_number_button`
  - `template_quick_reply_button`
  - `template_text_body`
  - `template_text_header`
  - `template_video_header`

### Removed
* `delivery_result`, `price`, `status`, `error` unused classes, replaced by the existing `delivery_reports`, `message_price`, `message_status`, `message_error` class.

## [[3.0.0](https://github.com/infobip/infobip-api-go-client/releases/tag/3.0.0)] - 2024-10-24

🎉 **NEW Major Version of `infobip-api-go-client`.**

⚠️ IMPORTANT NOTE: This release contains breaking changes! From this point onward `Go` 1.13 is no longer supported. The minimum supported version is `Go` 1.18.

In this release, we updated and modernized the infobip-api-go-client library. It is auto-generated and different from the previous version.

### Added
* Most recent [Infobip SMS API](https://www.infobip.com/docs/api/channels/sms) feature set.
* Support for [Infobip Messages API](https://www.infobip.com/docs/api/platform/messages-api).
* Added mock tests to verify the correctness of request payloads and response handling.

### Changed
* To improve maintainability and organization, we have separated the models for each Infobip product into their own directories. This means that each product —such as `SMS`, `Messages API`, or other services— will now have its own dedicated folder for models, making it easier to manage dependencies and clearly see which models are tied to each specific product. Users will now find a more structured, modular design, allowing easier integration and upgrades in the future.

## [[2.1.0](https://github.com/infobip/infobip-api-go-client/releases/tag/2.1.0)] - 2023-04-12

### Added
- Support for [Confirm Conversion API](https://www.infobip.com/docs/api/channels/sms/sms-messaging/outbound-sms/log-end-tag)

### Changed
- SMS models:
    - Additional fields related to [CPaaS X](https://www.infobip.com/docs/cpaas-x): `EntityId` and `ApplicationId`.
    - `UrlOptions` and `IncludeSMSCountInResponse` fields in `SmsAdvancedTextualRequest` model.
    - `ValidationErrors` field in `SmsApiRequestErrorDetails` model.
    - New regional options for sending promotional SMS to phone numbers registered in Turkey.
    - New model `UrlOptions` for setting up URL shortening and tracking feature.
- TFA models:
    - New languages support for sending text to speech messages: `hr`, `sl` and `ro`.


## [ [2.0.0](https://github.com/infobip/infobip-api-go-client/releases/tag/2.0.0) ] - 2021-06-30

🎉 **NEW Major Version of `infobip-api-go-client`.**

⚠ **IMPORTANT NOTE:** This release is the first major release of a given library.
The changes listed below refer to the diff between the previous `pre-release` `0.x.x` based version.

In this release, we updated and modernized the infobip-api-go-client library. It is auto-generated and completely different from the previous version.

### Added
- Support for [Infobip Two-factor Authentication API](https://www.infobip.com/docs/api#channels/sms/send-2fa-pin-code-over-sms)
- `CONTRIBUTING.md` which contains guidelines for creating GitHub issues

### Changed
- Targeting minimum Go version 1.13
- Models, structure, examples, etc. for [Infobip SMS API](https://www.infobip.com/docs/api#channels/sms)
- Library dependencies
- `README.md` which contains necessary data and examples for quickstart as well as some other important pieces of information on versioning, licensing, etc.

[readme]: README.mustache
