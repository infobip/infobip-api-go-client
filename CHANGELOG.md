# Change Log of `infobip-api-go-client`

All notable changes to the library will be documented in this file.

The format of the file is based on [Keep a Changelog](http://keepachangelog.com/)
and this library adheres to [Semantic Versioning](http://semver.org/) as mentioned in [README.md][readme] file.

## [ [3.1.0](https://github.com/infobip/infobip-api-go-client/releases/tag/3.1.0)] - 2025-01-20

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
