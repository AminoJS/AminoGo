# Security Policy

## Supported Versions

| Version | Supported          |
| ------- | ------------------ |
| 0.0.0 (Current)   | :white_check_mark: |

## Reporting a Vulnerability

**ALL VULNERABILITY SHOULD ONLY BE DISCUSSED OR DISCLOSED VIA A SECURE SOURCE, SUCH AS A SECURE EMAIL CHANNEL. DO NOT REPORT VIA THE GITHUB ISSUE TRACKING SYSTEM**

If you believe you have found any vulnerability, and would like to report to our core-developers, please do so by sending us an email at aminogo.security@secret.fyi. It would be much appreciated, and we will handle the issue as soon as possible with minimum disruption.

## Counter measures

### `1)` G304: File path provided as taint input

References:

Code: [GitHub](https://github.com/AminoJS/AminoGo/blob/eaa8f18a7c2fc029994951d357fcf9b59e84cd64/aminogo/upload_media.go#L118-L137)

Gosec: [Rules](https://github.com/securego/gosec#available-rules)

---

Ageist: Gosec's G304

Measures: 

    Added a snippet of code for checking whether is the original base directory path is identical to the one where
    AminoGo is going to operate on
    
Examples:

| Pass | Operate Path | Origin Path  |
|---|---|---|
| ✔️ | /usr/app/images/mock.jpg | /usr/app/images/  |
  ❌  | /usr/app/images/  | /etc/  |