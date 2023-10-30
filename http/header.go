package http

import (
	"bytes"
	"strings"
)

// Common HTTP headers, these are defined in RFC 7231 section 4.3.
const (
	HeaderInvalid Header = iota //

	// Authentication
	HeaderAuthorization // RFC 7235, 4.2
	HeaderProxyAuthenticate
	HeaderProxyAuthorization
	HeaderWWWAuthenticate

	// Caching
	HeaderAge
	HeaderCacheControl
	HeaderClearSiteData
	HeaderExpires
	HeaderPragma
	HeaderWarning

	// Client hints
	HeaderAcceptCH
	HeaderAcceptCHLifetime
	HeaderContentDPR
	HeaderDPR
	HeaderEarlyData
	HeaderSaveData
	HeaderViewportWidth
	HeaderWidth

	// Conditionals
	HeaderETag
	HeaderIfMatch
	HeaderIfModifiedSince
	HeaderIfNoneMatch
	HeaderIfUnmodifiedSince
	HeaderLastModified
	HeaderVary

	// Connection management
	HeaderConnection
	HeaderKeepAlive
	HeaderProxyConnection

	// Content negotiation
	HeaderAccept
	HeaderAcceptCharset
	HeaderAcceptEncoding
	HeaderAcceptLanguage

	// Controls
	HeaderCookie
	HeaderExpect
	HeaderMaxForwards
	HeaderSetCookie

	// CORS
	HeaderAccessControlAllowCredentials
	HeaderAccessControlAllowHeaders
	HeaderAccessControlAllowMethods
	HeaderAccessControlAllowOrigin
	HeaderAccessControlExposeHeaders
	HeaderAccessControlMaxAge
	HeaderAccessControlRequestHeaders
	HeaderAccessControlRequestMethod
	HeaderOrigin
	HeaderTimingAllowOrigin
	HeaderXPermittedCrossDomainPolicies

	// Do Not Track
	HeaderDNT
	HeaderTk

	// Downloads
	HeaderContentDisposition

	// Message body information
	HeaderContentEncoding
	HeaderContentLanguage
	HeaderContentLength
	HeaderContentLocation
	HeaderContentType

	// Proxies
	HeaderForwarded
	HeaderVia
	HeaderXForwardedFor
	HeaderXForwardedHost
	HeaderXForwardedProto

	// Redirects
	HeaderLocation

	// Request context
	HeaderFrom
	HeaderHost
	HeaderReferer
	HeaderReferrerPolicy
	HeaderUserAgent

	// Response context
	HeaderAllow
	HeaderServer

	// Range requests
	HeaderAcceptRanges
	HeaderContentRange
	HeaderIfRange
	HeaderRange

	// Security
	HeaderContentSecurityPolicy
	HeaderContentSecurityPolicyReportOnly
	HeaderCrossOriginResourcePolicy
	HeaderExpectCT
	HeaderFeaturePolicy
	HeaderPublicKeyPins
	HeaderPublicKeyPinsReportOnly
	HeaderStrictTransportSecurity
	HeaderUpgradeInsecureRequests
	HeaderXContentTypeOptions
	HeaderXDownloadOptions
	HeaderXFrameOptions
	HeaderXPoweredBy
	HeaderXXSSProtection

	// Server-sent event
	HeaderLastEventID
	HeaderNEL
	HeaderPingFrom
	HeaderPingTo
	HeaderReportTo

	// Transfer coding
	HeaderTE
	HeaderTrailer
	HeaderTransferEncoding

	// WebSockets
	HeaderSecWebSocketAccept
	HeaderSecWebSocketExtensions
	HeaderSecWebSocketKey
	HeaderSecWebSocketProtocol
	HeaderSecWebSocketVersion

	// Other
	HeaderAcceptPatch
	HeaderAcceptPushPolicy
	HeaderAcceptSignature
	HeaderAltSvc
	HeaderDate
	HeaderIndex
	HeaderLargeAllocation
	HeaderLink
	HeaderPushPolicy
	HeaderRetryAfter
	HeaderServerTiming
	HeaderSignature
	HeaderSignedHeaders
	HeaderSourceMap
	HeaderUpgrade
	HeaderXDNSPrefetchControl
	HeaderXPingback
	HeaderXRequestedWith
	HeaderXRobotsTag
	HeaderXUACompatible
)

var (
	// HeaderNames is a map of Header to string.
	HeaderNames = map[Header]string{
		// Authentication
		HeaderAuthorization:      "Authorization",
		HeaderProxyAuthenticate:  "Proxy-Authenticate",
		HeaderProxyAuthorization: "Proxy-Authorization",
		HeaderWWWAuthenticate:    "WWW-Authenticate",

		// Caching
		HeaderAge:           "Age",
		HeaderCacheControl:  "Cache-Control",
		HeaderClearSiteData: "Clear-Site-Payload",
		HeaderExpires:       "Expires",
		HeaderPragma:        "Pragma",
		HeaderWarning:       "Warning",

		// Client hints
		HeaderAcceptCH:         "Accept-CH",
		HeaderAcceptCHLifetime: "Accept-CH-Lifetime",
		HeaderContentDPR:       "Content-DPR",
		HeaderDPR:              "DPR",
		HeaderEarlyData:        "Early-Payload",
		HeaderSaveData:         "Save-Payload",
		HeaderViewportWidth:    "Viewport-Width",
		HeaderWidth:            "Width",

		// Conditionals
		HeaderETag:              "ETag",
		HeaderIfMatch:           "If-Match",
		HeaderIfModifiedSince:   "If-Modified-Since",
		HeaderIfNoneMatch:       "If-None-Match",
		HeaderIfUnmodifiedSince: "If-Unmodified-Since",
		HeaderLastModified:      "Last-Modified",
		HeaderVary:              "Vary",

		// Connection management
		HeaderConnection:      "Connection",
		HeaderKeepAlive:       "Keep-Alive",
		HeaderProxyConnection: "Proxy-Connection",

		// Content negotiation
		HeaderAccept:         "Accept",
		HeaderAcceptCharset:  "Accept-Charset",
		HeaderAcceptEncoding: "Accept-Encoding",
		HeaderAcceptLanguage: "Accept-Language",

		// Controls
		HeaderCookie:      "Cookie",
		HeaderExpect:      "Expect",
		HeaderMaxForwards: "Max-Forwards",
		HeaderSetCookie:   "Set-Cookie",

		// CORS
		HeaderAccessControlAllowCredentials: "Access-Control-Allow-Credentials",
		HeaderAccessControlAllowHeaders:     "Access-Control-Allow-Headers",
		HeaderAccessControlAllowMethods:     "Access-Control-Allow-Methods",
		HeaderAccessControlAllowOrigin:      "Access-Control-Allow-Origin",
		HeaderAccessControlExposeHeaders:    "Access-Control-Expose-Headers",
		HeaderAccessControlMaxAge:           "Access-Control-Max-Age",
		HeaderAccessControlRequestHeaders:   "Access-Control-Request-Headers",
		HeaderAccessControlRequestMethod:    "Access-Control-Request-Method",
		HeaderOrigin:                        "Origin",
		HeaderTimingAllowOrigin:             "Timing-Allow-Origin",
		HeaderXPermittedCrossDomainPolicies: "X-Permitted-Cross-Domain-Policies",

		// Do Not Track
		HeaderDNT: "DNT",
		HeaderTk:  "Tk",

		// Downloads
		HeaderContentDisposition: "Content-Disposition",

		// Message body information
		HeaderContentEncoding: "Content-Encoding",
		HeaderContentLanguage: "Content-Language",
		HeaderContentLength:   "Content-Length",
		HeaderContentLocation: "Content-Location",
		HeaderContentType:     "Content-Type",

		// Proxies
		HeaderForwarded:       "Forwarded",
		HeaderVia:             "Via",
		HeaderXForwardedFor:   "X-Forwarded-For",
		HeaderXForwardedHost:  "X-Forwarded-Host",
		HeaderXForwardedProto: "X-Forwarded-Proto",

		// Redirects
		HeaderLocation: "Location",

		// Request context
		HeaderFrom:           "From",
		HeaderHost:           "Host",
		HeaderReferer:        "Referer",
		HeaderReferrerPolicy: "Referrer-Policy",
		HeaderUserAgent:      "User-Agent",

		// Response context
		HeaderAllow:  "Allow",
		HeaderServer: "Server",

		// Range requests
		HeaderAcceptRanges: "Accept-Ranges",
		HeaderContentRange: "Content-Range",
		HeaderIfRange:      "If-Range",
		HeaderRange:        "Range",

		// Security
		HeaderContentSecurityPolicy:           "Content-Security-Policy",
		HeaderContentSecurityPolicyReportOnly: "Content-Security-Policy-Report-Only",
		HeaderCrossOriginResourcePolicy:       "Cross-Origin-Resource-Policy",
		HeaderExpectCT:                        "Expect-CT",
		HeaderFeaturePolicy:                   "Feature-Policy",
		HeaderPublicKeyPins:                   "Public-Key-Pins",
		HeaderPublicKeyPinsReportOnly:         "Public-Key-Pins-Report-Only",
		HeaderStrictTransportSecurity:         "Strict-Transport-Security",
		HeaderUpgradeInsecureRequests:         "Upgrade-Insecure-Requests",
		HeaderXContentTypeOptions:             "X-Content-Definitions-Options",
		HeaderXDownloadOptions:                "X-Download-Options",
		HeaderXFrameOptions:                   "X-Frame-Options",
		HeaderXPoweredBy:                      "X-Powered-By",
		HeaderXXSSProtection:                  "X-XSS-Protection",

		// Server-sent event
		HeaderLastEventID: "Last-Event-ID",
		HeaderNEL:         "NEL",
		HeaderPingFrom:    "Ping-From",
		HeaderPingTo:      "Ping-To",
		HeaderReportTo:    "Report-To",

		// Transfer coding
		HeaderTE:               "TE",
		HeaderTrailer:          "Trailer",
		HeaderTransferEncoding: "Transfer-Encoding",

		// WebSockets
		HeaderSecWebSocketAccept:     "Sec-WebSocket-Accept",
		HeaderSecWebSocketExtensions: "Sec-WebSocket-Extensions", /* #nosec G101 */
		HeaderSecWebSocketKey:        "Sec-WebSocket-Key",
		HeaderSecWebSocketProtocol:   "Sec-WebSocket-Protocol",
		HeaderSecWebSocketVersion:    "Sec-WebSocket-Version",

		// Other
		HeaderAcceptPatch:         "Accept-Patch",
		HeaderAcceptPushPolicy:    "Accept-Push-Policy",
		HeaderAcceptSignature:     "Accept-Signature",
		HeaderAltSvc:              "Alt-Svc",
		HeaderDate:                "Date",
		HeaderIndex:               "Index",
		HeaderLargeAllocation:     "Large-Allocation",
		HeaderLink:                "Link",
		HeaderPushPolicy:          "Push-Policy",
		HeaderRetryAfter:          "Retry-After",
		HeaderServerTiming:        "Server-Timing",
		HeaderSignature:           "Signature",
		HeaderSignedHeaders:       "Signed-Headers",
		HeaderSourceMap:           "SourceMap",
		HeaderUpgrade:             "Upgrade",
		HeaderXDNSPrefetchControl: "X-DNS-Prefetch-Control",
		HeaderXPingback:           "X-Pingback",
		HeaderXRequestedWith:      "X-Requested-With",
		HeaderXRobotsTag:          "X-Robots-Tag",
		HeaderXUACompatible:       "X-UA-Compatible",
	}
)

// String outputs the Header as a string.
func (h Header) String() string {
	return HeaderNames[h]
}

// MarshalJSON outputs the Header as a json.
func (h Header) MarshalJSON() ([]byte, error) {
	if !h.Validate() {
		return []byte(`""`), nil
	}

	return []byte(`"` + h.String() + `"`), nil
}

// UnmarshalJSON parses the Header from json.
func (h *Header) UnmarshalJSON(data []byte) error {
	str := string(bytes.Trim(data, `"`))
	if header := ParseHeader(str); header.Validate() {
		*h = header
	}

	return nil
}

// Validate returns true if the Header is valid.
func (h Header) Validate() bool {
	return h != HeaderInvalid
}

// ParseHeader parses the Header from string.
func ParseHeader(value string) Header {
	value = strings.ToLower(value)
	for k, v := range HeaderNames {
		if strings.ToLower(v) == value {
			return k
		}
	}

	return HeaderInvalid
}
