package response

// APIStatus defines a standard structure for response messages and codes
type APIStatus struct {
	Code    int
	Message string
}

// Define HTTP status constants directly as APIStatus with RFC references
var (
	StatusContinue           = APIStatus{100, "Continue"}            // RFC 9110, 15.2.1
	StatusSwitchingProtocols = APIStatus{101, "Switching Protocols"} // RFC 9110, 15.2.2
	StatusProcessing         = APIStatus{102, "Processing"}          // RFC 2518, 10.1
	StatusEarlyHints         = APIStatus{103, "Early Hints"}         // RFC 8297

	StatusOK                   = APIStatus{200, "OK"}                            // RFC 9110, 15.3.1
	StatusCreated              = APIStatus{201, "Created"}                       // RFC 9110, 15.3.2
	StatusAccepted             = APIStatus{202, "Accepted"}                      // RFC 9110, 15.3.3
	StatusNonAuthoritativeInfo = APIStatus{203, "Non-Authoritative Information"} // RFC 9110, 15.3.4
	StatusNoContent            = APIStatus{204, "No Content"}                    // RFC 9110, 15.3.5
	StatusResetContent         = APIStatus{205, "Reset Content"}                 // RFC 9110, 15.3.6
	StatusPartialContent       = APIStatus{206, "Partial Content"}               // RFC 9110, 15.3.7
	StatusMultiStatus          = APIStatus{207, "Multi-Status"}                  // RFC 4918, 11.1
	StatusAlreadyReported      = APIStatus{208, "Already Reported"}              // RFC 5842, 7.1
	StatusIMUsed               = APIStatus{226, "IM Used"}                       // RFC 3229, 10.4.1

	StatusMultipleChoices   = APIStatus{300, "Multiple Choices"}   // RFC 9110, 15.4.1
	StatusMovedPermanently  = APIStatus{301, "Moved Permanently"}  // RFC 9110, 15.4.2
	StatusFound             = APIStatus{302, "Found"}              // RFC 9110, 15.4.3
	StatusSeeOther          = APIStatus{303, "See Other"}          // RFC 9110, 15.4.4
	StatusNotModified       = APIStatus{304, "Not Modified"}       // RFC 9110, 15.4.5
	StatusUseProxy          = APIStatus{305, "Use Proxy"}          // RFC 9110, 15.4.6
	StatusTemporaryRedirect = APIStatus{307, "Temporary Redirect"} // RFC 9110, 15.4.8
	StatusPermanentRedirect = APIStatus{308, "Permanent Redirect"} // RFC 9110, 15.4.9

	StatusBadRequest                  = APIStatus{400, "Bad Request"}                     // RFC 9110, 15.5.1
	StatusUnauthorized                = APIStatus{401, "Unauthorized"}                    // RFC 9110, 15.5.2
	StatusPaymentRequired             = APIStatus{402, "Payment Required"}                // RFC 9110, 15.5.3
	StatusForbidden                   = APIStatus{403, "Forbidden"}                       // RFC 9110, 15.5.4
	StatusNotFound                    = APIStatus{404, "Not Found"}                       // RFC 9110, 15.5.5
	StatusMethodNotAllowed            = APIStatus{405, "Method Not Allowed"}              // RFC 9110, 15.5.6
	StatusNotAcceptable               = APIStatus{406, "Not Acceptable"}                  // RFC 9110, 15.5.7
	StatusProxyAuthRequired           = APIStatus{407, "Proxy Authentication Required"}   // RFC 9110, 15.5.8
	StatusRequestTimeout              = APIStatus{408, "Request Timeout"}                 // RFC 9110, 15.5.9
	StatusConflict                    = APIStatus{409, "Conflict"}                        // RFC 9110, 15.5.10
	StatusGone                        = APIStatus{410, "Gone"}                            // RFC 9110, 15.5.11
	StatusLengthRequired              = APIStatus{411, "Length Required"}                 // RFC 9110, 15.5.12
	StatusPreconditionFailed          = APIStatus{412, "Precondition Failed"}             // RFC 9110, 15.5.13
	StatusPayloadTooLarge             = APIStatus{413, "Payload Too Large"}               // RFC 9110, 15.5.14
	StatusURITooLong                  = APIStatus{414, "URI Too Long"}                    // RFC 9110, 15.5.15
	StatusUnsupportedMediaType        = APIStatus{415, "Unsupported Media Type"}          // RFC 9110, 15.5.16
	StatusRangeNotSatisfiable         = APIStatus{416, "Range Not Satisfiable"}           // RFC 9110, 15.5.17
	StatusExpectationFailed           = APIStatus{417, "Expectation Failed"}              // RFC 9110, 15.5.18
	StatusTeapot                      = APIStatus{418, "I'm a teapot"}                    // RFC 9110, 15.5.19 (Unused)
	StatusMisdirectedRequest          = APIStatus{421, "Misdirected Request"}             // RFC 9110, 15.5.20
	StatusUnprocessableEntity         = APIStatus{422, "Unprocessable Entity"}            // RFC 9110, 15.5.21
	StatusLocked                      = APIStatus{423, "Locked"}                          // RFC 4918, 11.3
	StatusFailedDependency            = APIStatus{424, "Failed Dependency"}               // RFC 4918, 11.4
	StatusTooEarly                    = APIStatus{425, "Too Early"}                       // RFC 8470, 5.2.
	StatusUpgradeRequired             = APIStatus{426, "Upgrade Required"}                // RFC 9110, 15.5.22
	StatusPreconditionRequired        = APIStatus{428, "Precondition Required"}           // RFC 6585, 3
	StatusTooManyRequests             = APIStatus{429, "Too Many Requests"}               // RFC 6585, 4
	StatusRequestHeaderFieldsTooLarge = APIStatus{431, "Request Header Fields Too Large"} // RFC 6585, 5
	StatusUnavailableForLegalReasons  = APIStatus{451, "Unavailable For Legal Reasons"}   // RFC 7725, 3

	StatusInternalServerError           = APIStatus{500, "Internal Server Error"}           // RFC 9110, 15.6.1
	StatusNotImplemented                = APIStatus{501, "Not Implemented"}                 // RFC 9110, 15.6.2
	StatusBadGateway                    = APIStatus{502, "Bad Gateway"}                     // RFC 9110, 15.6.3
	StatusServiceUnavailable            = APIStatus{503, "Service Unavailable"}             // RFC 9110, 15.6.4
	StatusGatewayTimeout                = APIStatus{504, "Gateway Timeout"}                 // RFC 9110, 15.6.5
	StatusHTTPVersionNotSupported       = APIStatus{505, "HTTP Version Not Supported"}      // RFC 9110, 15.6.6
	StatusVariantAlsoNegotiates         = APIStatus{506, "Variant Also Negotiates"}         // RFC 2295, 8.1
	StatusInsufficientStorage           = APIStatus{507, "Insufficient Storage"}            // RFC 4918, 11.5
	StatusLoopDetected                  = APIStatus{508, "Loop Detected"}                   // RFC 5842, 7.2
	StatusNotExtended                   = APIStatus{510, "Not Extended"}                    // RFC 2774, 7
	StatusNetworkAuthenticationRequired = APIStatus{511, "Network Authentication Required"} // RFC 6585, 6
)
