# OVH refresh DNS zone

This action refreshes a DNS zone using the OVH API.

## Inputs

## `application-key`

**Required** Application key.

## `application-secret`

**Required** Application secret.

## `consumer-key`

**Required** Consumer key.

## `api-endpoint`

API endpoint. Defaults to `ovh-eu`.

## `dns-zone`

**Required** DNS zone.

## Example usage

```yaml
uses: CrispyBaguette/ovh-dns-refresh-action@v1.0
with:
  application-key: foo
  application-secret: bar
  consumer-key: far
  api-endpoint: ovh-us
  dns-zone: example.com
```
