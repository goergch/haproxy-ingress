/*
Copyright 2019 The HAProxy Ingress Controller Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package types

// Global config
const (
	GlobalAcmeEmails                   = "acme-emails"
	GlobalAcmeEndpoint                 = "acme-endpoint"
	GlobalAcmeExpiring                 = "acme-expiring"
	GlobalAcmeShared                   = "acme-shared"
	GlobalAcmeTermsAgreed              = "acme-terms-agreed"
	GlobalAuthLogFormat                = "auth-log-format"
	GlobalAuthProxy                    = "auth-proxy"
	GlobalBindFrontingProxy            = "bind-fronting-proxy"
	GlobalBindHTTP                     = "bind-http"
	GlobalBindHTTPS                    = "bind-https"
	GlobalBindIPAddrHealthz            = "bind-ip-addr-healthz"
	GlobalBindIPAddrHTTP               = "bind-ip-addr-http"
	GlobalBindIPAddrPrometheus         = "bind-ip-addr-prometheus"
	GlobalBindIPAddrStats              = "bind-ip-addr-stats"
	GlobalBindIPAddrTCP                = "bind-ip-addr-tcp"
	GlobalCloseSessionsDuration        = "close-sessions-duration"
	GlobalConfigDefaults               = "config-defaults"
	GlobalConfigFrontend               = "config-frontend"
	GlobalConfigGlobal                 = "config-global"
	GlobalConfigProxy                  = "config-proxy"
	GlobalConfigSections               = "config-sections"
	GlobalConfigTCP                    = "config-tcp"
	GlobalCookieKey                    = "cookie-key"
	GlobalCPUMap                       = "cpu-map"
	GlobalCrossNamespaceSecretsCA      = "cross-namespace-secrets-ca"
	GlobalCrossNamespaceSecretsCrt     = "cross-namespace-secrets-crt"
	GlobalCrossNamespaceSecretsPasswd  = "cross-namespace-secrets-passwd"
	GlobalCrossNamespaceServices       = "cross-namespace-services"
	GlobalDefaultBackendRedirect       = "default-backend-redirect"
	GlobalDefaultBackendRedirectCode   = "default-backend-redirect-code"
	GlobalDefaultHost                  = "default-host"
	GlobalDNSAcceptedPayloadSize       = "dns-accepted-payload-size"
	GlobalDNSClusterDomain             = "dns-cluster-domain"
	GlobalDNSHoldObsolete              = "dns-hold-obsolete"
	GlobalDNSHoldValid                 = "dns-hold-valid"
	GlobalDNSResolvers                 = "dns-resolvers"
	GlobalDNSTimeoutRetry              = "dns-timeout-retry"
	GlobalDrainSupport                 = "drain-support"
	GlobalDrainSupportRedispatch       = "drain-support-redispatch"
	GlobalExternalHasLua               = "external-has-lua"
	GlobalForwardfor                   = "forwardfor"
	GlobalFrontingProxyPort            = "fronting-proxy-port"
	GlobalGroupname                    = "groupname"
	GlobalHealthzPort                  = "healthz-port"
	GlobalHTTPLogFormat                = "http-log-format"
	GlobalHTTPPort                     = "http-port"
	GlobalHTTPResponse200              = "http-response-200"
	GlobalHTTPResponse400              = "http-response-400"
	GlobalHTTPResponse401              = "http-response-401"
	GlobalHTTPResponse403              = "http-response-403"
	GlobalHTTPResponse404              = "http-response-404"
	GlobalHTTPResponse405              = "http-response-405"
	GlobalHTTPResponse407              = "http-response-407"
	GlobalHTTPResponse408              = "http-response-408"
	GlobalHTTPResponse410              = "http-response-410"
	GlobalHTTPResponse413              = "http-response-413"
	GlobalHTTPResponse421              = "http-response-421"
	GlobalHTTPResponse425              = "http-response-425"
	GlobalHTTPResponse429              = "http-response-429"
	GlobalHTTPResponse495              = "http-response-495"
	GlobalHTTPResponse496              = "http-response-496"
	GlobalHTTPResponse500              = "http-response-500"
	GlobalHTTPResponse501              = "http-response-501"
	GlobalHTTPResponse502              = "http-response-502"
	GlobalHTTPResponse503              = "http-response-503"
	GlobalHTTPResponse504              = "http-response-504"
	GlobalHTTPResponsePrometheusRoot   = "http-response-prometheus-root"
	GlobalHTTPSLogFormat               = "https-log-format"
	GlobalHTTPSPort                    = "https-port"
	GlobalHTTPStoHTTPPort              = "https-to-http-port"
	GlobalLoadServerState              = "load-server-state"
	GlobalMasterExitOnFailure          = "master-exit-on-failure"
	GlobalMaxConnections               = "max-connections"
	GlobalModsecurityEndpoints         = "modsecurity-endpoints"
	GlobalModsecurityTimeoutConnect    = "modsecurity-timeout-connect"
	GlobalModsecurityTimeoutHello      = "modsecurity-timeout-hello"
	GlobalModsecurityTimeoutIdle       = "modsecurity-timeout-idle"
	GlobalModsecurityTimeoutProcessing = "modsecurity-timeout-processing"
	GlobalModsecurityTimeoutServer     = "modsecurity-timeout-server"
	GlobalNbprocBalance                = "nbproc-balance"
	GlobalNbprocSSL                    = "nbproc-ssl"
	GlobalNbthread                     = "nbthread"
	GlobalNoTLSRedirectLocations       = "no-tls-redirect-locations"
	GlobalPathTypeOrder                = "path-type-order"
	GlobalUsername                     = "username"
	GlobalPrometheusPort               = "prometheus-port"
	GlobalRedirectFromCode             = "redirect-from-code"
	GlobalRedirectToCode               = "redirect-to-code"
	GlobalSSLDHDefaultMaxSize          = "ssl-dh-default-max-size"
	GlobalSSLDHParam                   = "ssl-dh-param"
	GlobalSSLEngine                    = "ssl-engine"
	GlobalSSLHeadersPrefix             = "ssl-headers-prefix"
	GlobalSSLModeAsync                 = "ssl-mode-async"
	GlobalSSLOptions                   = "ssl-options"
	GlobalSSLRedirectCode              = "ssl-redirect-code"
	GlobalStatsAuth                    = "stats-auth"
	GlobalStatsPort                    = "stats-port"
	GlobalStatsProxyProtocol           = "stats-proxy-protocol"
	GlobalStatsSSLCert                 = "stats-ssl-cert"
	GlobalStrictHost                   = "strict-host"
	GlobalSyslogEndpoint               = "syslog-endpoint"
	GlobalSyslogFormat                 = "syslog-format"
	GlobalSyslogLength                 = "syslog-length"
	GlobalSyslogTag                    = "syslog-tag"
	GlobalTCPLogFormat                 = "tcp-log-format"
	GlobalTimeoutClient                = "timeout-client"
	GlobalTimeoutClientFin             = "timeout-client-fin"
	GlobalTimeoutStop                  = "timeout-stop"
	GlobalUseChroot                    = "use-chroot"
	GlobalUseCPUMap                    = "use-cpu-map"
	GlobalUseForwardedProto            = "use-forwarded-proto"
	GlobalUseHAProxyUser               = "use-haproxy-user"
	GlobalUseHTX                       = "use-htx"
	GlobalUseProxyProtocol             = "use-proxy-protocol"
	GlobalWorkerMaxReloads             = "worker-max-reloads"
)
