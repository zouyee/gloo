
---
title: "hcm.proto"
weight: 5
---

<!-- Code generated by solo-kit. DO NOT EDIT. -->


### Package: `hcm.options.gloo.solo.io` 
#### Types:


- [HttpConnectionManagerSettings](#httpconnectionmanagersettings)
- [SetCurrentClientCertDetails](#setcurrentclientcertdetails)
- [ForwardClientCertDetails](#forwardclientcertdetails)
  



##### Source File: [github.com/solo-io/gloo/projects/gloo/api/v1/options/hcm/hcm.proto](https://github.com/solo-io/gloo/blob/master/projects/gloo/api/v1/options/hcm/hcm.proto)





---
### HttpConnectionManagerSettings

 
Contains various settings for Envoy's http connection manager.
See here for more information: https://www.envoyproxy.io/docs/envoy/v1.9.0/configuration/http_conn_man/http_conn_man

```yaml
"skipXffAppend": bool
"via": string
"xffNumTrustedHops": int
"useRemoteAddress": .google.protobuf.BoolValue
"generateRequestId": .google.protobuf.BoolValue
"proxy100Continue": bool
"streamIdleTimeout": .google.protobuf.Duration
"idleTimeout": .google.protobuf.Duration
"maxRequestHeadersKb": .google.protobuf.UInt32Value
"requestTimeout": .google.protobuf.Duration
"drainTimeout": .google.protobuf.Duration
"delayedCloseTimeout": .google.protobuf.Duration
"serverName": string
"acceptHttp10": bool
"defaultHostForHttp10": string
"tracing": .tracing.options.gloo.solo.io.ListenerTracingSettings
"forwardClientCertDetails": .hcm.options.gloo.solo.io.HttpConnectionManagerSettings.ForwardClientCertDetails
"setCurrentClientCertDetails": .hcm.options.gloo.solo.io.HttpConnectionManagerSettings.SetCurrentClientCertDetails
"preserveExternalRequestId": bool
"upgrades": []protocol_upgrade.options.gloo.solo.io.ProtocolUpgradeConfig

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `skipXffAppend` | `bool` |  |  |
| `via` | `string` |  |  |
| `xffNumTrustedHops` | `int` |  |  |
| `useRemoteAddress` | [.google.protobuf.BoolValue](https://developers.google.com/protocol-buffers/docs/reference/csharp/class/google/protobuf/well-known-types/bool-value) |  |  |
| `generateRequestId` | [.google.protobuf.BoolValue](https://developers.google.com/protocol-buffers/docs/reference/csharp/class/google/protobuf/well-known-types/bool-value) |  |  |
| `proxy100Continue` | `bool` |  |  |
| `streamIdleTimeout` | [.google.protobuf.Duration](https://developers.google.com/protocol-buffers/docs/reference/csharp/class/google/protobuf/well-known-types/duration) |  |  |
| `idleTimeout` | [.google.protobuf.Duration](https://developers.google.com/protocol-buffers/docs/reference/csharp/class/google/protobuf/well-known-types/duration) |  |  |
| `maxRequestHeadersKb` | [.google.protobuf.UInt32Value](https://developers.google.com/protocol-buffers/docs/reference/csharp/class/google/protobuf/well-known-types/u-int-32-value) |  |  |
| `requestTimeout` | [.google.protobuf.Duration](https://developers.google.com/protocol-buffers/docs/reference/csharp/class/google/protobuf/well-known-types/duration) |  |  |
| `drainTimeout` | [.google.protobuf.Duration](https://developers.google.com/protocol-buffers/docs/reference/csharp/class/google/protobuf/well-known-types/duration) |  |  |
| `delayedCloseTimeout` | [.google.protobuf.Duration](https://developers.google.com/protocol-buffers/docs/reference/csharp/class/google/protobuf/well-known-types/duration) |  |  |
| `serverName` | `string` |  |  |
| `acceptHttp10` | `bool` | For explanation of these settings see: https://www.envoyproxy.io/docs/envoy/latest/api-v2/api/v2/core/protocol.proto#envoy-api-msg-core-http1protocoloptions. |  |
| `defaultHostForHttp10` | `string` |  |  |
| `tracing` | [.tracing.options.gloo.solo.io.ListenerTracingSettings](../../tracing/tracing.proto.sk/#listenertracingsettings) |  |  |
| `forwardClientCertDetails` | [.hcm.options.gloo.solo.io.HttpConnectionManagerSettings.ForwardClientCertDetails](../hcm.proto.sk/#forwardclientcertdetails) |  |  |
| `setCurrentClientCertDetails` | [.hcm.options.gloo.solo.io.HttpConnectionManagerSettings.SetCurrentClientCertDetails](../hcm.proto.sk/#setcurrentclientcertdetails) |  |  |
| `preserveExternalRequestId` | `bool` |  |  |
| `upgrades` | [[]protocol_upgrade.options.gloo.solo.io.ProtocolUpgradeConfig](../../protocol_upgrade/protocol_upgrade.proto.sk/#protocolupgradeconfig) | HttpConnectionManager configuration for protocol upgrade requests. Note: WebSocket upgrades are enabled by default on the HTTP Connection Manager and must be explicitly disabled. |  |




---
### SetCurrentClientCertDetails



```yaml
"subject": .google.protobuf.BoolValue
"cert": bool
"chain": bool
"dns": bool
"uri": bool

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `subject` | [.google.protobuf.BoolValue](https://developers.google.com/protocol-buffers/docs/reference/csharp/class/google/protobuf/well-known-types/bool-value) |  |  |
| `cert` | `bool` |  |  |
| `chain` | `bool` |  |  |
| `dns` | `bool` |  |  |
| `uri` | `bool` |  |  |




---
### ForwardClientCertDetails



| Name | Description |
| ----- | ----------- | 
| `SANITIZE` |  |
| `FORWARD_ONLY` |  |
| `APPEND_FORWARD` |  |
| `SANITIZE_SET` |  |
| `ALWAYS_FORWARD_ONLY` |  |





<!-- Start of HubSpot Embed Code -->
<script type="text/javascript" id="hs-script-loader" async defer src="//js.hs-scripts.com/5130874.js"></script>
<!-- End of HubSpot Embed Code -->
