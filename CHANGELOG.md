# Change Log of `infobip-api-go-client`

All notable changes to the library will be documented in this file.

The format of the file is based on [Keep a Changelog](http://keepachangelog.com/)
and this library adheres to [Semantic Versioning](http://semver.org/) as mentioned in [README.md][readme] file.

## [[3.0.0](https://github.com/infobip/infobip-api-go-client/releases/tag/3.0.0)] - 2024-10-24

🎉 **NEW Major Version of `infobip-api-go-client`.**

⚠ IMPORTANT NOTE: This release contains breaking changes! From this point onward `Go` 1.13 is no longer supported. The minimum supported version is `Go` 1.18.

In this release, we updated and modernized the infobip library. It is auto-generated and different from the previous version.

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

In this release, we updated and modernized the infobip library. It is auto-generated and completely different from the previous version.

### Added
- Support for [Infobip Two-factor Authentication API](https://www.infobip.com/docs/api#channels/sms/send-2fa-pin-code-over-sms)
- `CONTRIBUTING.md` which contains guidelines for creating GitHub issues

### Changed
- Targeting minimum Go version 1.13
- Models, structure, examples, etc. for [Infobip SMS API](https://www.infobip.com/docs/api#channels/sms)
- Library dependencies
- `README.md` which contains necessary data and examples for quickstart as well as some other important pieces of information on versioning, licensing, etc.

[readme]: README.mustache

