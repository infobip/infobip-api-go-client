# Change Log of `infobip`

All notable changes to the library will be documented in this file.

The format of the file is based on [Keep a Changelog](http://keepachangelog.com/)
and this library adheres to [Semantic Versioning](http://semver.org/) as mentioned in [README.md][readme] file.

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

🎉 **NEW Major Version of `infobip`.**

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
