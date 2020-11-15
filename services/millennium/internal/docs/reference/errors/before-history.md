---
title: Before History
replacement: https://developers.aiblocks.io/api/errors/http-status-codes/millennium-specific/
---

A millennium server may be configured to only keep a portion of the aiblocks network's history stored
within its database.  This error will be returned when a client requests a piece of information
(such as a page of transactions or a single operation) that the server can positively identify as
falling outside the range of recorded history.

This error returns a
[HTTP 410 Error](https://developer.mozilla.org/en-US/docs/Web/HTTP/Response_codes).

## Attributes

As with all errors Millennium returns, `before_history` follows the
[Problem Details for HTTP APIs](https://tools.ietf.org/html/draft-ietf-appsawg-http-problem-00)
draft specification guide and thus has the following attributes:

| Attribute   | Type   | Description                                                                     |
| ----------- | ------ | ------------------------------------------------------------------------------- |
| `type`      | URL    | The identifier for the error.  This is a URL that can be visited in the browser.|
| `title`     | String | A short title describing the error.                                             |
| `status`    | Number | An HTTP status code that maps to the error.                                     |
| `detail`    | String | A more detailed description of the error.                                       |

## Example

```shell
$ curl -X GET "https://millennium-testnet.aiblocks.io/transactions?cursor=1&order=desc"
{
  "type": "https://aiblocks.io/millennium-errors/before_history",
  "title": "Data Requested Is Before Recorded History",
  "status": 410,
  "detail": "This millennium instance is configured to only track a portion of the aiblocks network's latest history. This request is asking for results prior to the recorded history known to this millennium instance."
}
```

## Related

- [Not Found](./not-found.md)