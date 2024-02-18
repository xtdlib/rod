// This file is generated by "./lib/proto/generate"

package proto

/*

Emulation

This domain emulates different environments for the page.

*/

// EmulationScreenOrientationType enum.
type EmulationScreenOrientationType string

const (
	// EmulationScreenOrientationTypePortraitPrimary enum const.
	EmulationScreenOrientationTypePortraitPrimary EmulationScreenOrientationType = "portraitPrimary"

	// EmulationScreenOrientationTypePortraitSecondary enum const.
	EmulationScreenOrientationTypePortraitSecondary EmulationScreenOrientationType = "portraitSecondary"

	// EmulationScreenOrientationTypeLandscapePrimary enum const.
	EmulationScreenOrientationTypeLandscapePrimary EmulationScreenOrientationType = "landscapePrimary"

	// EmulationScreenOrientationTypeLandscapeSecondary enum const.
	EmulationScreenOrientationTypeLandscapeSecondary EmulationScreenOrientationType = "landscapeSecondary"
)

// EmulationScreenOrientation Screen orientation.
type EmulationScreenOrientation struct {
	// Type Orientation type.
	Type EmulationScreenOrientationType `json:"type"`

	// Angle Orientation angle.
	Angle int `json:"angle"`
}

// EmulationDisplayFeatureOrientation enum.
type EmulationDisplayFeatureOrientation string

const (
	// EmulationDisplayFeatureOrientationVertical enum const.
	EmulationDisplayFeatureOrientationVertical EmulationDisplayFeatureOrientation = "vertical"

	// EmulationDisplayFeatureOrientationHorizontal enum const.
	EmulationDisplayFeatureOrientationHorizontal EmulationDisplayFeatureOrientation = "horizontal"
)

// EmulationDisplayFeature ...
type EmulationDisplayFeature struct {
	// Orientation of a display feature in relation to screen
	Orientation EmulationDisplayFeatureOrientation `json:"orientation"`

	// Offset The offset from the screen origin in either the x (for vertical
	// orientation) or y (for horizontal orientation) direction.
	Offset int `json:"offset"`

	// MaskLength A display feature may mask content such that it is not physically
	// displayed - this length along with the offset describes this area.
	// A display feature that only splits content will have a 0 mask_length.
	MaskLength int `json:"maskLength"`
}

// EmulationMediaFeature ...
type EmulationMediaFeature struct {
	// Name ...
	Name string `json:"name"`

	// Value ...
	Value string `json:"value"`
}

// EmulationVirtualTimePolicy (experimental) advance: If the scheduler runs out of immediate work, the virtual time base may fast forward to
// allow the next delayed task (if any) to run; pause: The virtual time base may not advance;
// pauseIfNetworkFetchesPending: The virtual time base may not advance if there are any pending
// resource fetches.
type EmulationVirtualTimePolicy string

const (
	// EmulationVirtualTimePolicyAdvance enum const.
	EmulationVirtualTimePolicyAdvance EmulationVirtualTimePolicy = "advance"

	// EmulationVirtualTimePolicyPause enum const.
	EmulationVirtualTimePolicyPause EmulationVirtualTimePolicy = "pause"

	// EmulationVirtualTimePolicyPauseIfNetworkFetchesPending enum const.
	EmulationVirtualTimePolicyPauseIfNetworkFetchesPending EmulationVirtualTimePolicy = "pauseIfNetworkFetchesPending"
)

// EmulationUserAgentBrandVersion (experimental) Used to specify User Agent Cient Hints to emulate. See https://wicg.github.io/ua-client-hints
type EmulationUserAgentBrandVersion struct {
	// Brand ...
	Brand string `json:"brand"`

	// Version ...
	Version string `json:"version"`
}

// EmulationUserAgentMetadata (experimental) Used to specify User Agent Cient Hints to emulate. See https://wicg.github.io/ua-client-hints
// Missing optional values will be filled in by the target with what it would normally use.
type EmulationUserAgentMetadata struct {
	// Brands (optional) Brands appearing in Sec-CH-UA.
	Brands []*EmulationUserAgentBrandVersion `json:"brands,omitempty"`

	// FullVersionList (optional) Brands appearing in Sec-CH-UA-Full-Version-List.
	FullVersionList []*EmulationUserAgentBrandVersion `json:"fullVersionList,omitempty"`

	// FullVersion (deprecated) (optional) ...
	FullVersion string `json:"fullVersion,omitempty"`

	// Platform ...
	Platform string `json:"platform"`

	// PlatformVersion ...
	PlatformVersion string `json:"platformVersion"`

	// Architecture ...
	Architecture string `json:"architecture"`

	// Model ...
	Model string `json:"model"`

	// Mobile ...
	Mobile bool `json:"mobile"`

	// Bitness (optional) ...
	Bitness string `json:"bitness,omitempty"`

	// Wow64 (optional) ...
	Wow64 bool `json:"wow64,omitempty"`
}

// EmulationDisabledImageType (experimental) Enum of image types that can be disabled.
type EmulationDisabledImageType string

const (
	// EmulationDisabledImageTypeAvif enum const.
	EmulationDisabledImageTypeAvif EmulationDisabledImageType = "avif"

	// EmulationDisabledImageTypeWebp enum const.
	EmulationDisabledImageTypeWebp EmulationDisabledImageType = "webp"
)

// EmulationCanEmulate Tells whether emulation is supported.
type EmulationCanEmulate struct{}

// ProtoReq name.
func (m EmulationCanEmulate) ProtoReq() string { return "Emulation.canEmulate" }

// Call the request.
func (m EmulationCanEmulate) Call(c Client) (*EmulationCanEmulateResult, error) {
	var res EmulationCanEmulateResult
	return &res, call(m.ProtoReq(), m, &res, c)
}

// EmulationCanEmulateResult ...
type EmulationCanEmulateResult struct {
	// Result True if emulation is supported.
	Result bool `json:"result"`
}

// EmulationClearDeviceMetricsOverride Clears the overridden device metrics.
type EmulationClearDeviceMetricsOverride struct{}

// ProtoReq name.
func (m EmulationClearDeviceMetricsOverride) ProtoReq() string {
	return "Emulation.clearDeviceMetricsOverride"
}

// Call sends the request.
func (m EmulationClearDeviceMetricsOverride) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// EmulationClearGeolocationOverride Clears the overridden Geolocation Position and Error.
type EmulationClearGeolocationOverride struct{}

// ProtoReq name.
func (m EmulationClearGeolocationOverride) ProtoReq() string {
	return "Emulation.clearGeolocationOverride"
}

// Call sends the request.
func (m EmulationClearGeolocationOverride) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// EmulationResetPageScaleFactor (experimental) Requests that page scale factor is reset to initial values.
type EmulationResetPageScaleFactor struct{}

// ProtoReq name.
func (m EmulationResetPageScaleFactor) ProtoReq() string { return "Emulation.resetPageScaleFactor" }

// Call sends the request.
func (m EmulationResetPageScaleFactor) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// EmulationSetFocusEmulationEnabled (experimental) Enables or disables simulating a focused and active page.
type EmulationSetFocusEmulationEnabled struct {
	// Enabled Whether to enable to disable focus emulation.
	Enabled bool `json:"enabled"`
}

// ProtoReq name.
func (m EmulationSetFocusEmulationEnabled) ProtoReq() string {
	return "Emulation.setFocusEmulationEnabled"
}

// Call sends the request.
func (m EmulationSetFocusEmulationEnabled) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// EmulationSetAutoDarkModeOverride (experimental) Automatically render all web contents using a dark theme.
type EmulationSetAutoDarkModeOverride struct {
	// Enabled (optional) Whether to enable or disable automatic dark mode.
	// If not specified, any existing override will be cleared.
	Enabled bool `json:"enabled,omitempty"`
}

// ProtoReq name.
func (m EmulationSetAutoDarkModeOverride) ProtoReq() string {
	return "Emulation.setAutoDarkModeOverride"
}

// Call sends the request.
func (m EmulationSetAutoDarkModeOverride) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// EmulationSetCPUThrottlingRate (experimental) Enables CPU throttling to emulate slow CPUs.
type EmulationSetCPUThrottlingRate struct {
	// Rate Throttling rate as a slowdown factor (1 is no throttle, 2 is 2x slowdown, etc).
	Rate float64 `json:"rate"`
}

// ProtoReq name.
func (m EmulationSetCPUThrottlingRate) ProtoReq() string { return "Emulation.setCPUThrottlingRate" }

// Call sends the request.
func (m EmulationSetCPUThrottlingRate) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// EmulationSetDefaultBackgroundColorOverride Sets or clears an override of the default background color of the frame. This override is used
// if the content does not specify one.
type EmulationSetDefaultBackgroundColorOverride struct {
	// Color (optional) RGBA of the default background color. If not specified, any existing override will be
	// cleared.
	Color *DOMRGBA `json:"color,omitempty"`
}

// ProtoReq name.
func (m EmulationSetDefaultBackgroundColorOverride) ProtoReq() string {
	return "Emulation.setDefaultBackgroundColorOverride"
}

// Call sends the request.
func (m EmulationSetDefaultBackgroundColorOverride) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// EmulationSetDeviceMetricsOverride Overrides the values of device screen dimensions (window.screen.width, window.screen.height,
// window.innerWidth, window.innerHeight, and "device-width"/"device-height"-related CSS media
// query results).
type EmulationSetDeviceMetricsOverride struct {
	// Width Overriding width value in pixels (minimum 0, maximum 10000000). 0 disables the override.
	Width int `json:"width"`

	// Height Overriding height value in pixels (minimum 0, maximum 10000000). 0 disables the override.
	Height int `json:"height"`

	// DeviceScaleFactor Overriding device scale factor value. 0 disables the override.
	DeviceScaleFactor float64 `json:"deviceScaleFactor"`

	// Mobile Whether to emulate mobile device. This includes viewport meta tag, overlay scrollbars, text
	// autosizing and more.
	Mobile bool `json:"mobile"`

	// Scale (experimental) (optional) Scale to apply to resulting view image.
	Scale *float64 `json:"scale,omitempty"`

	// ScreenWidth (experimental) (optional) Overriding screen width value in pixels (minimum 0, maximum 10000000).
	ScreenWidth *int `json:"screenWidth,omitempty"`

	// ScreenHeight (experimental) (optional) Overriding screen height value in pixels (minimum 0, maximum 10000000).
	ScreenHeight *int `json:"screenHeight,omitempty"`

	// PositionX (experimental) (optional) Overriding view X position on screen in pixels (minimum 0, maximum 10000000).
	PositionX *int `json:"positionX,omitempty"`

	// PositionY (experimental) (optional) Overriding view Y position on screen in pixels (minimum 0, maximum 10000000).
	PositionY *int `json:"positionY,omitempty"`

	// DontSetVisibleSize (experimental) (optional) Do not set visible view size, rely upon explicit setVisibleSize call.
	DontSetVisibleSize bool `json:"dontSetVisibleSize,omitempty"`

	// ScreenOrientation (optional) Screen orientation override.
	ScreenOrientation *EmulationScreenOrientation `json:"screenOrientation,omitempty"`

	// Viewport (experimental) (optional) If set, the visible area of the page will be overridden to this viewport. This viewport
	// change is not observed by the page, e.g. viewport-relative elements do not change positions.
	Viewport *PageViewport `json:"viewport,omitempty"`

	// DisplayFeature (experimental) (optional) If set, the display feature of a multi-segment screen. If not set, multi-segment support
	// is turned-off.
	DisplayFeature *EmulationDisplayFeature `json:"displayFeature,omitempty"`
}

// ProtoReq name.
func (m EmulationSetDeviceMetricsOverride) ProtoReq() string {
	return "Emulation.setDeviceMetricsOverride"
}

// Call sends the request.
func (m EmulationSetDeviceMetricsOverride) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// EmulationSetScrollbarsHidden (experimental) ...
type EmulationSetScrollbarsHidden struct {
	// Hidden Whether scrollbars should be always hidden.
	Hidden bool `json:"hidden"`
}

// ProtoReq name.
func (m EmulationSetScrollbarsHidden) ProtoReq() string { return "Emulation.setScrollbarsHidden" }

// Call sends the request.
func (m EmulationSetScrollbarsHidden) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// EmulationSetDocumentCookieDisabled (experimental) ...
type EmulationSetDocumentCookieDisabled struct {
	// Disabled Whether document.coookie API should be disabled.
	Disabled bool `json:"disabled"`
}

// ProtoReq name.
func (m EmulationSetDocumentCookieDisabled) ProtoReq() string {
	return "Emulation.setDocumentCookieDisabled"
}

// Call sends the request.
func (m EmulationSetDocumentCookieDisabled) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// EmulationSetEmitTouchEventsForMouseConfiguration enum.
type EmulationSetEmitTouchEventsForMouseConfiguration string

const (
	// EmulationSetEmitTouchEventsForMouseConfigurationMobile enum const.
	EmulationSetEmitTouchEventsForMouseConfigurationMobile EmulationSetEmitTouchEventsForMouseConfiguration = "mobile"

	// EmulationSetEmitTouchEventsForMouseConfigurationDesktop enum const.
	EmulationSetEmitTouchEventsForMouseConfigurationDesktop EmulationSetEmitTouchEventsForMouseConfiguration = "desktop"
)

// EmulationSetEmitTouchEventsForMouse (experimental) ...
type EmulationSetEmitTouchEventsForMouse struct {
	// Enabled Whether touch emulation based on mouse input should be enabled.
	Enabled bool `json:"enabled"`

	// Configuration (optional) Touch/gesture events configuration. Default: current platform.
	Configuration EmulationSetEmitTouchEventsForMouseConfiguration `json:"configuration,omitempty"`
}

// ProtoReq name.
func (m EmulationSetEmitTouchEventsForMouse) ProtoReq() string {
	return "Emulation.setEmitTouchEventsForMouse"
}

// Call sends the request.
func (m EmulationSetEmitTouchEventsForMouse) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// EmulationSetEmulatedMedia Emulates the given media type or media feature for CSS media queries.
type EmulationSetEmulatedMedia struct {
	// Media (optional) Media type to emulate. Empty string disables the override.
	Media string `json:"media,omitempty"`

	// Features (optional) Media features to emulate.
	Features []*EmulationMediaFeature `json:"features,omitempty"`
}

// ProtoReq name.
func (m EmulationSetEmulatedMedia) ProtoReq() string { return "Emulation.setEmulatedMedia" }

// Call sends the request.
func (m EmulationSetEmulatedMedia) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// EmulationSetEmulatedVisionDeficiencyType enum.
type EmulationSetEmulatedVisionDeficiencyType string

const (
	// EmulationSetEmulatedVisionDeficiencyTypeNone enum const.
	EmulationSetEmulatedVisionDeficiencyTypeNone EmulationSetEmulatedVisionDeficiencyType = "none"

	// EmulationSetEmulatedVisionDeficiencyTypeBlurredVision enum const.
	EmulationSetEmulatedVisionDeficiencyTypeBlurredVision EmulationSetEmulatedVisionDeficiencyType = "blurredVision"

	// EmulationSetEmulatedVisionDeficiencyTypeReducedContrast enum const.
	EmulationSetEmulatedVisionDeficiencyTypeReducedContrast EmulationSetEmulatedVisionDeficiencyType = "reducedContrast"

	// EmulationSetEmulatedVisionDeficiencyTypeAchromatopsia enum const.
	EmulationSetEmulatedVisionDeficiencyTypeAchromatopsia EmulationSetEmulatedVisionDeficiencyType = "achromatopsia"

	// EmulationSetEmulatedVisionDeficiencyTypeDeuteranopia enum const.
	EmulationSetEmulatedVisionDeficiencyTypeDeuteranopia EmulationSetEmulatedVisionDeficiencyType = "deuteranopia"

	// EmulationSetEmulatedVisionDeficiencyTypeProtanopia enum const.
	EmulationSetEmulatedVisionDeficiencyTypeProtanopia EmulationSetEmulatedVisionDeficiencyType = "protanopia"

	// EmulationSetEmulatedVisionDeficiencyTypeTritanopia enum const.
	EmulationSetEmulatedVisionDeficiencyTypeTritanopia EmulationSetEmulatedVisionDeficiencyType = "tritanopia"
)

// EmulationSetEmulatedVisionDeficiency (experimental) Emulates the given vision deficiency.
type EmulationSetEmulatedVisionDeficiency struct {
	// Type Vision deficiency to emulate. Order: best-effort emulations come first, followed by any
	// physiologically accurate emulations for medically recognized color vision deficiencies.
	Type EmulationSetEmulatedVisionDeficiencyType `json:"type"`
}

// ProtoReq name.
func (m EmulationSetEmulatedVisionDeficiency) ProtoReq() string {
	return "Emulation.setEmulatedVisionDeficiency"
}

// Call sends the request.
func (m EmulationSetEmulatedVisionDeficiency) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// EmulationSetGeolocationOverride Overrides the Geolocation Position or Error. Omitting any of the parameters emulates position
// unavailable.
type EmulationSetGeolocationOverride struct {
	// Latitude (optional) Mock latitude
	Latitude *float64 `json:"latitude,omitempty"`

	// Longitude (optional) Mock longitude
	Longitude *float64 `json:"longitude,omitempty"`

	// Accuracy (optional) Mock accuracy
	Accuracy *float64 `json:"accuracy,omitempty"`
}

// ProtoReq name.
func (m EmulationSetGeolocationOverride) ProtoReq() string { return "Emulation.setGeolocationOverride" }

// Call sends the request.
func (m EmulationSetGeolocationOverride) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// EmulationSetIdleOverride (experimental) Overrides the Idle state.
type EmulationSetIdleOverride struct {
	// IsUserActive Mock isUserActive
	IsUserActive bool `json:"isUserActive"`

	// IsScreenUnlocked Mock isScreenUnlocked
	IsScreenUnlocked bool `json:"isScreenUnlocked"`
}

// ProtoReq name.
func (m EmulationSetIdleOverride) ProtoReq() string { return "Emulation.setIdleOverride" }

// Call sends the request.
func (m EmulationSetIdleOverride) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// EmulationClearIdleOverride (experimental) Clears Idle state overrides.
type EmulationClearIdleOverride struct{}

// ProtoReq name.
func (m EmulationClearIdleOverride) ProtoReq() string { return "Emulation.clearIdleOverride" }

// Call sends the request.
func (m EmulationClearIdleOverride) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// EmulationSetNavigatorOverrides (deprecated) (experimental) Overrides value returned by the javascript navigator object.
type EmulationSetNavigatorOverrides struct {
	// Platform The platform navigator.platform should return.
	Platform string `json:"platform"`
}

// ProtoReq name.
func (m EmulationSetNavigatorOverrides) ProtoReq() string { return "Emulation.setNavigatorOverrides" }

// Call sends the request.
func (m EmulationSetNavigatorOverrides) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// EmulationSetPageScaleFactor (experimental) Sets a specified page scale factor.
type EmulationSetPageScaleFactor struct {
	// PageScaleFactor Page scale factor.
	PageScaleFactor float64 `json:"pageScaleFactor"`
}

// ProtoReq name.
func (m EmulationSetPageScaleFactor) ProtoReq() string { return "Emulation.setPageScaleFactor" }

// Call sends the request.
func (m EmulationSetPageScaleFactor) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// EmulationSetScriptExecutionDisabled Switches script execution in the page.
type EmulationSetScriptExecutionDisabled struct {
	// Value Whether script execution should be disabled in the page.
	Value bool `json:"value"`
}

// ProtoReq name.
func (m EmulationSetScriptExecutionDisabled) ProtoReq() string {
	return "Emulation.setScriptExecutionDisabled"
}

// Call sends the request.
func (m EmulationSetScriptExecutionDisabled) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// EmulationSetTouchEmulationEnabled Enables touch on platforms which do not support them.
type EmulationSetTouchEmulationEnabled struct {
	// Enabled Whether the touch event emulation should be enabled.
	Enabled bool `json:"enabled"`

	// MaxTouchPoints (optional) Maximum touch points supported. Defaults to one.
	MaxTouchPoints *int `json:"maxTouchPoints,omitempty"`
}

// ProtoReq name.
func (m EmulationSetTouchEmulationEnabled) ProtoReq() string {
	return "Emulation.setTouchEmulationEnabled"
}

// Call sends the request.
func (m EmulationSetTouchEmulationEnabled) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// EmulationSetVirtualTimePolicy (experimental) Turns on virtual time for all frames (replacing real-time with a synthetic time source) and sets
// the current virtual time policy.  Note this supersedes any previous time budget.
type EmulationSetVirtualTimePolicy struct {
	// Policy ...
	Policy EmulationVirtualTimePolicy `json:"policy"`

	// Budget (optional) If set, after this many virtual milliseconds have elapsed virtual time will be paused and a
	// virtualTimeBudgetExpired event is sent.
	Budget *float64 `json:"budget,omitempty"`

	// MaxVirtualTimeTaskStarvationCount (optional) If set this specifies the maximum number of tasks that can be run before virtual is forced
	// forwards to prevent deadlock.
	MaxVirtualTimeTaskStarvationCount *int `json:"maxVirtualTimeTaskStarvationCount,omitempty"`

	// InitialVirtualTime (optional) If set, base::Time::Now will be overridden to initially return this value.
	InitialVirtualTime TimeSinceEpoch `json:"initialVirtualTime,omitempty"`
}

// ProtoReq name.
func (m EmulationSetVirtualTimePolicy) ProtoReq() string { return "Emulation.setVirtualTimePolicy" }

// Call the request.
func (m EmulationSetVirtualTimePolicy) Call(c Client) (*EmulationSetVirtualTimePolicyResult, error) {
	var res EmulationSetVirtualTimePolicyResult
	return &res, call(m.ProtoReq(), m, &res, c)
}

// EmulationSetVirtualTimePolicyResult (experimental) ...
type EmulationSetVirtualTimePolicyResult struct {
	// VirtualTimeTicksBase Absolute timestamp at which virtual time was first enabled (up time in milliseconds).
	VirtualTimeTicksBase float64 `json:"virtualTimeTicksBase"`
}

// EmulationSetLocaleOverride (experimental) Overrides default host system locale with the specified one.
type EmulationSetLocaleOverride struct {
	// Locale (optional) ICU style C locale (e.g. "en_US"). If not specified or empty, disables the override and
	// restores default host system locale.
	Locale string `json:"locale,omitempty"`
}

// ProtoReq name.
func (m EmulationSetLocaleOverride) ProtoReq() string { return "Emulation.setLocaleOverride" }

// Call sends the request.
func (m EmulationSetLocaleOverride) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// EmulationSetTimezoneOverride (experimental) Overrides default host system timezone with the specified one.
type EmulationSetTimezoneOverride struct {
	// TimezoneID The timezone identifier. If empty, disables the override and
	// restores default host system timezone.
	TimezoneID string `json:"timezoneId"`
}

// ProtoReq name.
func (m EmulationSetTimezoneOverride) ProtoReq() string { return "Emulation.setTimezoneOverride" }

// Call sends the request.
func (m EmulationSetTimezoneOverride) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// EmulationSetVisibleSize (deprecated) (experimental) Resizes the frame/viewport of the page. Note that this does not affect the frame's container
// (e.g. browser window). Can be used to produce screenshots of the specified size. Not supported
// on Android.
type EmulationSetVisibleSize struct {
	// Width Frame width (DIP).
	Width int `json:"width"`

	// Height Frame height (DIP).
	Height int `json:"height"`
}

// ProtoReq name.
func (m EmulationSetVisibleSize) ProtoReq() string { return "Emulation.setVisibleSize" }

// Call sends the request.
func (m EmulationSetVisibleSize) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// EmulationSetDisabledImageTypes (experimental) ...
type EmulationSetDisabledImageTypes struct {
	// ImageTypes Image types to disable.
	ImageTypes []EmulationDisabledImageType `json:"imageTypes"`
}

// ProtoReq name.
func (m EmulationSetDisabledImageTypes) ProtoReq() string { return "Emulation.setDisabledImageTypes" }

// Call sends the request.
func (m EmulationSetDisabledImageTypes) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// EmulationSetHardwareConcurrencyOverride (experimental) ...
type EmulationSetHardwareConcurrencyOverride struct {
	// HardwareConcurrency Hardware concurrency to report
	HardwareConcurrency int `json:"hardwareConcurrency"`
}

// ProtoReq name.
func (m EmulationSetHardwareConcurrencyOverride) ProtoReq() string {
	return "Emulation.setHardwareConcurrencyOverride"
}

// Call sends the request.
func (m EmulationSetHardwareConcurrencyOverride) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// EmulationSetUserAgentOverride Allows overriding user agent with the given string.
type EmulationSetUserAgentOverride struct {
	// UserAgent User agent to use.
	UserAgent string `json:"userAgent"`

	// AcceptLanguage (optional) Browser langugage to emulate.
	AcceptLanguage string `json:"acceptLanguage,omitempty"`

	// Platform (optional) The platform navigator.platform should return.
	Platform string `json:"platform,omitempty"`

	// UserAgentMetadata (experimental) (optional) To be sent in Sec-CH-UA-* headers and returned in navigator.userAgentData
	UserAgentMetadata *EmulationUserAgentMetadata `json:"userAgentMetadata,omitempty"`
}

// ProtoReq name.
func (m EmulationSetUserAgentOverride) ProtoReq() string { return "Emulation.setUserAgentOverride" }

// Call sends the request.
func (m EmulationSetUserAgentOverride) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// EmulationSetAutomationOverride (experimental) Allows overriding the automation flag.
type EmulationSetAutomationOverride struct {
	// Enabled Whether the override should be enabled.
	Enabled bool `json:"enabled"`
}

// ProtoReq name.
func (m EmulationSetAutomationOverride) ProtoReq() string { return "Emulation.setAutomationOverride" }

// Call sends the request.
func (m EmulationSetAutomationOverride) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// EmulationVirtualTimeBudgetExpired (experimental) Notification sent after the virtual time budget for the current VirtualTimePolicy has run out.
type EmulationVirtualTimeBudgetExpired struct{}

// ProtoEvent name.
func (evt EmulationVirtualTimeBudgetExpired) ProtoEvent() string {
	return "Emulation.virtualTimeBudgetExpired"
}
