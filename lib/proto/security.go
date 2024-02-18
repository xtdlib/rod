// This file is generated by "./lib/proto/generate"

package proto

/*

Security

Security

*/

// SecurityCertificateID An internal certificate ID value.
type SecurityCertificateID int

// SecurityMixedContentType A description of mixed content (HTTP resources on HTTPS pages), as defined by
// https://www.w3.org/TR/mixed-content/#categories
type SecurityMixedContentType string

const (
	// SecurityMixedContentTypeBlockable enum const.
	SecurityMixedContentTypeBlockable SecurityMixedContentType = "blockable"

	// SecurityMixedContentTypeOptionallyBlockable enum const.
	SecurityMixedContentTypeOptionallyBlockable SecurityMixedContentType = "optionally-blockable"

	// SecurityMixedContentTypeNone enum const.
	SecurityMixedContentTypeNone SecurityMixedContentType = "none"
)

// SecuritySecurityState The security level of a page or resource.
type SecuritySecurityState string

const (
	// SecuritySecurityStateUnknown enum const.
	SecuritySecurityStateUnknown SecuritySecurityState = "unknown"

	// SecuritySecurityStateNeutral enum const.
	SecuritySecurityStateNeutral SecuritySecurityState = "neutral"

	// SecuritySecurityStateInsecure enum const.
	SecuritySecurityStateInsecure SecuritySecurityState = "insecure"

	// SecuritySecurityStateSecure enum const.
	SecuritySecurityStateSecure SecuritySecurityState = "secure"

	// SecuritySecurityStateInfo enum const.
	SecuritySecurityStateInfo SecuritySecurityState = "info"

	// SecuritySecurityStateInsecureBroken enum const.
	SecuritySecurityStateInsecureBroken SecuritySecurityState = "insecure-broken"
)

// SecurityCertificateSecurityState (experimental) Details about the security state of the page certificate.
type SecurityCertificateSecurityState struct {
	// Protocol name (e.g. "TLS 1.2" or "QUIC").
	Protocol string `json:"protocol"`

	// KeyExchange Key Exchange used by the connection, or the empty string if not applicable.
	KeyExchange string `json:"keyExchange"`

	// KeyExchangeGroup (optional) (EC)DH group used by the connection, if applicable.
	KeyExchangeGroup string `json:"keyExchangeGroup,omitempty"`

	// Cipher name.
	Cipher string `json:"cipher"`

	// Mac (optional) TLS MAC. Note that AEAD ciphers do not have separate MACs.
	Mac string `json:"mac,omitempty"`

	// Certificate Page certificate.
	Certificate []string `json:"certificate"`

	// SubjectName Certificate subject name.
	SubjectName string `json:"subjectName"`

	// Issuer Name of the issuing CA.
	Issuer string `json:"issuer"`

	// ValidFrom Certificate valid from date.
	ValidFrom TimeSinceEpoch `json:"validFrom"`

	// ValidTo Certificate valid to (expiration) date
	ValidTo TimeSinceEpoch `json:"validTo"`

	// CertificateNetworkError (optional) The highest priority network error code, if the certificate has an error.
	CertificateNetworkError string `json:"certificateNetworkError,omitempty"`

	// CertificateHasWeakSignature True if the certificate uses a weak signature aglorithm.
	CertificateHasWeakSignature bool `json:"certificateHasWeakSignature"`

	// CertificateHasSha1Signature True if the certificate has a SHA1 signature in the chain.
	CertificateHasSha1Signature bool `json:"certificateHasSha1Signature"`

	// ModernSSL True if modern SSL
	ModernSSL bool `json:"modernSSL"`

	// ObsoleteSslProtocol True if the connection is using an obsolete SSL protocol.
	ObsoleteSslProtocol bool `json:"obsoleteSslProtocol"`

	// ObsoleteSslKeyExchange True if the connection is using an obsolete SSL key exchange.
	ObsoleteSslKeyExchange bool `json:"obsoleteSslKeyExchange"`

	// ObsoleteSslCipher True if the connection is using an obsolete SSL cipher.
	ObsoleteSslCipher bool `json:"obsoleteSslCipher"`

	// ObsoleteSslSignature True if the connection is using an obsolete SSL signature.
	ObsoleteSslSignature bool `json:"obsoleteSslSignature"`
}

// SecuritySafetyTipStatus (experimental) ...
type SecuritySafetyTipStatus string

const (
	// SecuritySafetyTipStatusBadReputation enum const.
	SecuritySafetyTipStatusBadReputation SecuritySafetyTipStatus = "badReputation"

	// SecuritySafetyTipStatusLookalike enum const.
	SecuritySafetyTipStatusLookalike SecuritySafetyTipStatus = "lookalike"
)

// SecuritySafetyTipInfo (experimental) ...
type SecuritySafetyTipInfo struct {
	// SafetyTipStatus Describes whether the page triggers any safety tips or reputation warnings. Default is unknown.
	SafetyTipStatus SecuritySafetyTipStatus `json:"safetyTipStatus"`

	// SafeURL (optional) The URL the safety tip suggested ("Did you mean?"). Only filled in for lookalike matches.
	SafeURL string `json:"safeUrl,omitempty"`
}

// SecurityVisibleSecurityState (experimental) Security state information about the page.
type SecurityVisibleSecurityState struct {
	// SecurityState The security level of the page.
	SecurityState SecuritySecurityState `json:"securityState"`

	// CertificateSecurityState (optional) Security state details about the page certificate.
	CertificateSecurityState *SecurityCertificateSecurityState `json:"certificateSecurityState,omitempty"`

	// SafetyTipInfo (optional) The type of Safety Tip triggered on the page. Note that this field will be set even if the Safety Tip UI was not actually shown.
	SafetyTipInfo *SecuritySafetyTipInfo `json:"safetyTipInfo,omitempty"`

	// SecurityStateIssueIDs Array of security state issues ids.
	SecurityStateIssueIDs []string `json:"securityStateIssueIds"`
}

// SecuritySecurityStateExplanation An explanation of an factor contributing to the security state.
type SecuritySecurityStateExplanation struct {
	// SecurityState Security state representing the severity of the factor being explained.
	SecurityState SecuritySecurityState `json:"securityState"`

	// Title describing the type of factor.
	Title string `json:"title"`

	// Summary Short phrase describing the type of factor.
	Summary string `json:"summary"`

	// Description Full text explanation of the factor.
	Description string `json:"description"`

	// MixedContentType The type of mixed content described by the explanation.
	MixedContentType SecurityMixedContentType `json:"mixedContentType"`

	// Certificate Page certificate.
	Certificate []string `json:"certificate"`

	// Recommendations (optional) Recommendations to fix any issues.
	Recommendations []string `json:"recommendations,omitempty"`
}

// SecurityInsecureContentStatus (deprecated) Information about insecure content on the page.
type SecurityInsecureContentStatus struct {
	// RanMixedContent Always false.
	RanMixedContent bool `json:"ranMixedContent"`

	// DisplayedMixedContent Always false.
	DisplayedMixedContent bool `json:"displayedMixedContent"`

	// ContainedMixedForm Always false.
	ContainedMixedForm bool `json:"containedMixedForm"`

	// RanContentWithCertErrors Always false.
	RanContentWithCertErrors bool `json:"ranContentWithCertErrors"`

	// DisplayedContentWithCertErrors Always false.
	DisplayedContentWithCertErrors bool `json:"displayedContentWithCertErrors"`

	// RanInsecureContentStyle Always set to unknown.
	RanInsecureContentStyle SecuritySecurityState `json:"ranInsecureContentStyle"`

	// DisplayedInsecureContentStyle Always set to unknown.
	DisplayedInsecureContentStyle SecuritySecurityState `json:"displayedInsecureContentStyle"`
}

// SecurityCertificateErrorAction The action to take when a certificate error occurs. continue will continue processing the
// request and cancel will cancel the request.
type SecurityCertificateErrorAction string

const (
	// SecurityCertificateErrorActionContinue enum const.
	SecurityCertificateErrorActionContinue SecurityCertificateErrorAction = "continue"

	// SecurityCertificateErrorActionCancel enum const.
	SecurityCertificateErrorActionCancel SecurityCertificateErrorAction = "cancel"
)

// SecurityDisable Disables tracking security state changes.
type SecurityDisable struct{}

// ProtoReq name.
func (m SecurityDisable) ProtoReq() string { return "Security.disable" }

// Call sends the request.
func (m SecurityDisable) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// SecurityEnable Enables tracking security state changes.
type SecurityEnable struct{}

// ProtoReq name.
func (m SecurityEnable) ProtoReq() string { return "Security.enable" }

// Call sends the request.
func (m SecurityEnable) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// SecuritySetIgnoreCertificateErrors (experimental) Enable/disable whether all certificate errors should be ignored.
type SecuritySetIgnoreCertificateErrors struct {
	// Ignore If true, all certificate errors will be ignored.
	Ignore bool `json:"ignore"`
}

// ProtoReq name.
func (m SecuritySetIgnoreCertificateErrors) ProtoReq() string {
	return "Security.setIgnoreCertificateErrors"
}

// Call sends the request.
func (m SecuritySetIgnoreCertificateErrors) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// SecurityHandleCertificateError (deprecated) Handles a certificate error that fired a certificateError event.
type SecurityHandleCertificateError struct {
	// EventID The ID of the event.
	EventID int `json:"eventId"`

	// Action The action to take on the certificate error.
	Action SecurityCertificateErrorAction `json:"action"`
}

// ProtoReq name.
func (m SecurityHandleCertificateError) ProtoReq() string { return "Security.handleCertificateError" }

// Call sends the request.
func (m SecurityHandleCertificateError) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// SecuritySetOverrideCertificateErrors (deprecated) Enable/disable overriding certificate errors. If enabled, all certificate error events need to
// be handled by the DevTools client and should be answered with `handleCertificateError` commands.
type SecuritySetOverrideCertificateErrors struct {
	// Override If true, certificate errors will be overridden.
	Override bool `json:"override"`
}

// ProtoReq name.
func (m SecuritySetOverrideCertificateErrors) ProtoReq() string {
	return "Security.setOverrideCertificateErrors"
}

// Call sends the request.
func (m SecuritySetOverrideCertificateErrors) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// SecurityCertificateError (deprecated) There is a certificate error. If overriding certificate errors is enabled, then it should be
// handled with the `handleCertificateError` command. Note: this event does not fire if the
// certificate error has been allowed internally. Only one client per target should override
// certificate errors at the same time.
type SecurityCertificateError struct {
	// EventID The ID of the event.
	EventID int `json:"eventId"`

	// ErrorType The type of the error.
	ErrorType string `json:"errorType"`

	// RequestURL The url that was requested.
	RequestURL string `json:"requestURL"`
}

// ProtoEvent name.
func (evt SecurityCertificateError) ProtoEvent() string {
	return "Security.certificateError"
}

// SecurityVisibleSecurityStateChanged (experimental) The security state of the page changed.
type SecurityVisibleSecurityStateChanged struct {
	// VisibleSecurityState Security state information about the page.
	VisibleSecurityState *SecurityVisibleSecurityState `json:"visibleSecurityState"`
}

// ProtoEvent name.
func (evt SecurityVisibleSecurityStateChanged) ProtoEvent() string {
	return "Security.visibleSecurityStateChanged"
}

// SecuritySecurityStateChanged (deprecated) The security state of the page changed. No longer being sent.
type SecuritySecurityStateChanged struct {
	// SecurityState Security state.
	SecurityState SecuritySecurityState `json:"securityState"`

	// SchemeIsCryptographic (deprecated) True if the page was loaded over cryptographic transport such as HTTPS.
	SchemeIsCryptographic bool `json:"schemeIsCryptographic"`

	// Explanations (deprecated) Previously a list of explanations for the security state. Now always
	// empty.
	Explanations []*SecuritySecurityStateExplanation `json:"explanations"`

	// InsecureContentStatus (deprecated) Information about insecure content on the page.
	InsecureContentStatus *SecurityInsecureContentStatus `json:"insecureContentStatus"`

	// Summary (deprecated) (optional) Overrides user-visible description of the state. Always omitted.
	Summary string `json:"summary,omitempty"`
}

// ProtoEvent name.
func (evt SecuritySecurityStateChanged) ProtoEvent() string {
	return "Security.securityStateChanged"
}
