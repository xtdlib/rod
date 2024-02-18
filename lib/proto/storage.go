// This file is generated by "./lib/proto/generate"

package proto

/*

Storage

*/

// StorageSerializedStorageKey ...
type StorageSerializedStorageKey string

// StorageStorageType Enum of possible storage types.
type StorageStorageType string

const (
	// StorageStorageTypeAppcache enum const.
	StorageStorageTypeAppcache StorageStorageType = "appcache"

	// StorageStorageTypeCookies enum const.
	StorageStorageTypeCookies StorageStorageType = "cookies"

	// StorageStorageTypeFileSystems enum const.
	StorageStorageTypeFileSystems StorageStorageType = "file_systems"

	// StorageStorageTypeIndexeddb enum const.
	StorageStorageTypeIndexeddb StorageStorageType = "indexeddb"

	// StorageStorageTypeLocalStorage enum const.
	StorageStorageTypeLocalStorage StorageStorageType = "local_storage"

	// StorageStorageTypeShaderCache enum const.
	StorageStorageTypeShaderCache StorageStorageType = "shader_cache"

	// StorageStorageTypeWebsql enum const.
	StorageStorageTypeWebsql StorageStorageType = "websql"

	// StorageStorageTypeServiceWorkers enum const.
	StorageStorageTypeServiceWorkers StorageStorageType = "service_workers"

	// StorageStorageTypeCacheStorage enum const.
	StorageStorageTypeCacheStorage StorageStorageType = "cache_storage"

	// StorageStorageTypeInterestGroups enum const.
	StorageStorageTypeInterestGroups StorageStorageType = "interest_groups"

	// StorageStorageTypeSharedStorage enum const.
	StorageStorageTypeSharedStorage StorageStorageType = "shared_storage"

	// StorageStorageTypeStorageBuckets enum const.
	StorageStorageTypeStorageBuckets StorageStorageType = "storage_buckets"

	// StorageStorageTypeAll enum const.
	StorageStorageTypeAll StorageStorageType = "all"

	// StorageStorageTypeOther enum const.
	StorageStorageTypeOther StorageStorageType = "other"
)

// StorageUsageForType Usage for a storage type.
type StorageUsageForType struct {
	// StorageType Name of storage type.
	StorageType StorageStorageType `json:"storageType"`

	// Usage Storage usage (bytes).
	Usage float64 `json:"usage"`
}

// StorageTrustTokens (experimental) Pair of issuer origin and number of available (signed, but not used) Trust
// Tokens from that issuer.
type StorageTrustTokens struct {
	// IssuerOrigin ...
	IssuerOrigin string `json:"issuerOrigin"`

	// Count ...
	Count float64 `json:"count"`
}

// StorageInterestGroupAccessType Enum of interest group access types.
type StorageInterestGroupAccessType string

const (
	// StorageInterestGroupAccessTypeJoin enum const.
	StorageInterestGroupAccessTypeJoin StorageInterestGroupAccessType = "join"

	// StorageInterestGroupAccessTypeLeave enum const.
	StorageInterestGroupAccessTypeLeave StorageInterestGroupAccessType = "leave"

	// StorageInterestGroupAccessTypeUpdate enum const.
	StorageInterestGroupAccessTypeUpdate StorageInterestGroupAccessType = "update"

	// StorageInterestGroupAccessTypeLoaded enum const.
	StorageInterestGroupAccessTypeLoaded StorageInterestGroupAccessType = "loaded"

	// StorageInterestGroupAccessTypeBid enum const.
	StorageInterestGroupAccessTypeBid StorageInterestGroupAccessType = "bid"

	// StorageInterestGroupAccessTypeWin enum const.
	StorageInterestGroupAccessTypeWin StorageInterestGroupAccessType = "win"
)

// StorageInterestGroupAd Ad advertising element inside an interest group.
type StorageInterestGroupAd struct {
	// RenderURL ...
	RenderURL string `json:"renderUrl"`

	// Metadata (optional) ...
	Metadata string `json:"metadata,omitempty"`
}

// StorageInterestGroupDetails The full details of an interest group.
type StorageInterestGroupDetails struct {
	// OwnerOrigin ...
	OwnerOrigin string `json:"ownerOrigin"`

	// Name ...
	Name string `json:"name"`

	// ExpirationTime ...
	ExpirationTime TimeSinceEpoch `json:"expirationTime"`

	// JoiningOrigin ...
	JoiningOrigin string `json:"joiningOrigin"`

	// BiddingURL (optional) ...
	BiddingURL string `json:"biddingUrl,omitempty"`

	// BiddingWasmHelperURL (optional) ...
	BiddingWasmHelperURL string `json:"biddingWasmHelperUrl,omitempty"`

	// UpdateURL (optional) ...
	UpdateURL string `json:"updateUrl,omitempty"`

	// TrustedBiddingSignalsURL (optional) ...
	TrustedBiddingSignalsURL string `json:"trustedBiddingSignalsUrl,omitempty"`

	// TrustedBiddingSignalsKeys ...
	TrustedBiddingSignalsKeys []string `json:"trustedBiddingSignalsKeys"`

	// UserBiddingSignals (optional) ...
	UserBiddingSignals string `json:"userBiddingSignals,omitempty"`

	// Ads ...
	Ads []*StorageInterestGroupAd `json:"ads"`

	// AdComponents ...
	AdComponents []*StorageInterestGroupAd `json:"adComponents"`
}

// StorageSharedStorageAccessType Enum of shared storage access types.
type StorageSharedStorageAccessType string

const (
	// StorageSharedStorageAccessTypeDocumentAddModule enum const.
	StorageSharedStorageAccessTypeDocumentAddModule StorageSharedStorageAccessType = "documentAddModule"

	// StorageSharedStorageAccessTypeDocumentSelectURL enum const.
	StorageSharedStorageAccessTypeDocumentSelectURL StorageSharedStorageAccessType = "documentSelectURL"

	// StorageSharedStorageAccessTypeDocumentRun enum const.
	StorageSharedStorageAccessTypeDocumentRun StorageSharedStorageAccessType = "documentRun"

	// StorageSharedStorageAccessTypeDocumentSet enum const.
	StorageSharedStorageAccessTypeDocumentSet StorageSharedStorageAccessType = "documentSet"

	// StorageSharedStorageAccessTypeDocumentAppend enum const.
	StorageSharedStorageAccessTypeDocumentAppend StorageSharedStorageAccessType = "documentAppend"

	// StorageSharedStorageAccessTypeDocumentDelete enum const.
	StorageSharedStorageAccessTypeDocumentDelete StorageSharedStorageAccessType = "documentDelete"

	// StorageSharedStorageAccessTypeDocumentClear enum const.
	StorageSharedStorageAccessTypeDocumentClear StorageSharedStorageAccessType = "documentClear"

	// StorageSharedStorageAccessTypeWorkletSet enum const.
	StorageSharedStorageAccessTypeWorkletSet StorageSharedStorageAccessType = "workletSet"

	// StorageSharedStorageAccessTypeWorkletAppend enum const.
	StorageSharedStorageAccessTypeWorkletAppend StorageSharedStorageAccessType = "workletAppend"

	// StorageSharedStorageAccessTypeWorkletDelete enum const.
	StorageSharedStorageAccessTypeWorkletDelete StorageSharedStorageAccessType = "workletDelete"

	// StorageSharedStorageAccessTypeWorkletClear enum const.
	StorageSharedStorageAccessTypeWorkletClear StorageSharedStorageAccessType = "workletClear"

	// StorageSharedStorageAccessTypeWorkletGet enum const.
	StorageSharedStorageAccessTypeWorkletGet StorageSharedStorageAccessType = "workletGet"

	// StorageSharedStorageAccessTypeWorkletKeys enum const.
	StorageSharedStorageAccessTypeWorkletKeys StorageSharedStorageAccessType = "workletKeys"

	// StorageSharedStorageAccessTypeWorkletEntries enum const.
	StorageSharedStorageAccessTypeWorkletEntries StorageSharedStorageAccessType = "workletEntries"

	// StorageSharedStorageAccessTypeWorkletLength enum const.
	StorageSharedStorageAccessTypeWorkletLength StorageSharedStorageAccessType = "workletLength"

	// StorageSharedStorageAccessTypeWorkletRemainingBudget enum const.
	StorageSharedStorageAccessTypeWorkletRemainingBudget StorageSharedStorageAccessType = "workletRemainingBudget"
)

// StorageSharedStorageEntry Struct for a single key-value pair in an origin's shared storage.
type StorageSharedStorageEntry struct {
	// Key ...
	Key string `json:"key"`

	// Value ...
	Value string `json:"value"`
}

// StorageSharedStorageMetadata Details for an origin's shared storage.
type StorageSharedStorageMetadata struct {
	// CreationTime ...
	CreationTime TimeSinceEpoch `json:"creationTime"`

	// Length ...
	Length int `json:"length"`

	// RemainingBudget ...
	RemainingBudget float64 `json:"remainingBudget"`
}

// StorageSharedStorageReportingMetadata Pair of reporting metadata details for a candidate URL for `selectURL()`.
type StorageSharedStorageReportingMetadata struct {
	// EventType ...
	EventType string `json:"eventType"`

	// ReportingURL ...
	ReportingURL string `json:"reportingUrl"`
}

// StorageSharedStorageURLWithMetadata Bundles a candidate URL with its reporting metadata.
type StorageSharedStorageURLWithMetadata struct {
	// URL Spec of candidate URL.
	URL string `json:"url"`

	// ReportingMetadata Any associated reporting metadata.
	ReportingMetadata []*StorageSharedStorageReportingMetadata `json:"reportingMetadata"`
}

// StorageSharedStorageAccessParams Bundles the parameters for shared storage access events whose
// presence/absence can vary according to SharedStorageAccessType.
type StorageSharedStorageAccessParams struct {
	// ScriptSourceURL (optional) Spec of the module script URL.
	// Present only for SharedStorageAccessType.documentAddModule.
	ScriptSourceURL string `json:"scriptSourceUrl,omitempty"`

	// OperationName (optional) Name of the registered operation to be run.
	// Present only for SharedStorageAccessType.documentRun and
	// SharedStorageAccessType.documentSelectURL.
	OperationName string `json:"operationName,omitempty"`

	// SerializedData (optional) The operation's serialized data in bytes (converted to a string).
	// Present only for SharedStorageAccessType.documentRun and
	// SharedStorageAccessType.documentSelectURL.
	SerializedData string `json:"serializedData,omitempty"`

	// UrlsWithMetadata (optional) Array of candidate URLs' specs, along with any associated metadata.
	// Present only for SharedStorageAccessType.documentSelectURL.
	UrlsWithMetadata []*StorageSharedStorageURLWithMetadata `json:"urlsWithMetadata,omitempty"`

	// Key (optional) Key for a specific entry in an origin's shared storage.
	// Present only for SharedStorageAccessType.documentSet,
	// SharedStorageAccessType.documentAppend,
	// SharedStorageAccessType.documentDelete,
	// SharedStorageAccessType.workletSet,
	// SharedStorageAccessType.workletAppend,
	// SharedStorageAccessType.workletDelete, and
	// SharedStorageAccessType.workletGet.
	Key string `json:"key,omitempty"`

	// Value (optional) Value for a specific entry in an origin's shared storage.
	// Present only for SharedStorageAccessType.documentSet,
	// SharedStorageAccessType.documentAppend,
	// SharedStorageAccessType.workletSet, and
	// SharedStorageAccessType.workletAppend.
	Value string `json:"value,omitempty"`

	// IgnoreIfPresent (optional) Whether or not to set an entry for a key if that key is already present.
	// Present only for SharedStorageAccessType.documentSet and
	// SharedStorageAccessType.workletSet.
	IgnoreIfPresent bool `json:"ignoreIfPresent,omitempty"`
}

// StorageStorageBucketsDurability ...
type StorageStorageBucketsDurability string

const (
	// StorageStorageBucketsDurabilityRelaxed enum const.
	StorageStorageBucketsDurabilityRelaxed StorageStorageBucketsDurability = "relaxed"

	// StorageStorageBucketsDurabilityStrict enum const.
	StorageStorageBucketsDurabilityStrict StorageStorageBucketsDurability = "strict"
)

// StorageStorageBucketInfo ...
type StorageStorageBucketInfo struct {
	// StorageKey ...
	StorageKey StorageSerializedStorageKey `json:"storageKey"`

	// ID ...
	ID string `json:"id"`

	// Name ...
	Name string `json:"name"`

	// IsDefault ...
	IsDefault bool `json:"isDefault"`

	// Expiration ...
	Expiration TimeSinceEpoch `json:"expiration"`

	// Quota Storage quota (bytes).
	Quota float64 `json:"quota"`

	// Persistent ...
	Persistent bool `json:"persistent"`

	// Durability ...
	Durability StorageStorageBucketsDurability `json:"durability"`
}

// StorageGetStorageKeyForFrame Returns a storage key given a frame id.
type StorageGetStorageKeyForFrame struct {
	// FrameID ...
	FrameID PageFrameID `json:"frameId"`
}

// ProtoReq name.
func (m StorageGetStorageKeyForFrame) ProtoReq() string { return "Storage.getStorageKeyForFrame" }

// Call the request.
func (m StorageGetStorageKeyForFrame) Call(c Client) (*StorageGetStorageKeyForFrameResult, error) {
	var res StorageGetStorageKeyForFrameResult
	return &res, call(m.ProtoReq(), m, &res, c)
}

// StorageGetStorageKeyForFrameResult ...
type StorageGetStorageKeyForFrameResult struct {
	// StorageKey ...
	StorageKey StorageSerializedStorageKey `json:"storageKey"`
}

// StorageClearDataForOrigin Clears storage for origin.
type StorageClearDataForOrigin struct {
	// Origin Security origin.
	Origin string `json:"origin"`

	// StorageTypes Comma separated list of StorageType to clear.
	StorageTypes string `json:"storageTypes"`
}

// ProtoReq name.
func (m StorageClearDataForOrigin) ProtoReq() string { return "Storage.clearDataForOrigin" }

// Call sends the request.
func (m StorageClearDataForOrigin) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// StorageClearDataForStorageKey Clears storage for storage key.
type StorageClearDataForStorageKey struct {
	// StorageKey Storage key.
	StorageKey string `json:"storageKey"`

	// StorageTypes Comma separated list of StorageType to clear.
	StorageTypes string `json:"storageTypes"`
}

// ProtoReq name.
func (m StorageClearDataForStorageKey) ProtoReq() string { return "Storage.clearDataForStorageKey" }

// Call sends the request.
func (m StorageClearDataForStorageKey) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// StorageGetCookies Returns all browser cookies.
type StorageGetCookies struct {
	// BrowserContextID (optional) Browser context to use when called on the browser endpoint.
	BrowserContextID BrowserBrowserContextID `json:"browserContextId,omitempty"`
}

// ProtoReq name.
func (m StorageGetCookies) ProtoReq() string { return "Storage.getCookies" }

// Call the request.
func (m StorageGetCookies) Call(c Client) (*StorageGetCookiesResult, error) {
	var res StorageGetCookiesResult
	return &res, call(m.ProtoReq(), m, &res, c)
}

// StorageGetCookiesResult ...
type StorageGetCookiesResult struct {
	// Cookies Array of cookie objects.
	Cookies []*NetworkCookie `json:"cookies"`
}

// StorageSetCookies Sets given cookies.
type StorageSetCookies struct {
	// Cookies to be set.
	Cookies []*NetworkCookieParam `json:"cookies"`

	// BrowserContextID (optional) Browser context to use when called on the browser endpoint.
	BrowserContextID BrowserBrowserContextID `json:"browserContextId,omitempty"`
}

// ProtoReq name.
func (m StorageSetCookies) ProtoReq() string { return "Storage.setCookies" }

// Call sends the request.
func (m StorageSetCookies) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// StorageClearCookies Clears cookies.
type StorageClearCookies struct {
	// BrowserContextID (optional) Browser context to use when called on the browser endpoint.
	BrowserContextID BrowserBrowserContextID `json:"browserContextId,omitempty"`
}

// ProtoReq name.
func (m StorageClearCookies) ProtoReq() string { return "Storage.clearCookies" }

// Call sends the request.
func (m StorageClearCookies) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// StorageGetUsageAndQuota Returns usage and quota in bytes.
type StorageGetUsageAndQuota struct {
	// Origin Security origin.
	Origin string `json:"origin"`
}

// ProtoReq name.
func (m StorageGetUsageAndQuota) ProtoReq() string { return "Storage.getUsageAndQuota" }

// Call the request.
func (m StorageGetUsageAndQuota) Call(c Client) (*StorageGetUsageAndQuotaResult, error) {
	var res StorageGetUsageAndQuotaResult
	return &res, call(m.ProtoReq(), m, &res, c)
}

// StorageGetUsageAndQuotaResult ...
type StorageGetUsageAndQuotaResult struct {
	// Usage Storage usage (bytes).
	Usage float64 `json:"usage"`

	// Quota Storage quota (bytes).
	Quota float64 `json:"quota"`

	// OverrideActive Whether or not the origin has an active storage quota override
	OverrideActive bool `json:"overrideActive"`

	// UsageBreakdown Storage usage per type (bytes).
	UsageBreakdown []*StorageUsageForType `json:"usageBreakdown"`
}

// StorageOverrideQuotaForOrigin (experimental) Override quota for the specified origin.
type StorageOverrideQuotaForOrigin struct {
	// Origin Security origin.
	Origin string `json:"origin"`

	// QuotaSize (optional) The quota size (in bytes) to override the original quota with.
	// If this is called multiple times, the overridden quota will be equal to
	// the quotaSize provided in the final call. If this is called without
	// specifying a quotaSize, the quota will be reset to the default value for
	// the specified origin. If this is called multiple times with different
	// origins, the override will be maintained for each origin until it is
	// disabled (called without a quotaSize).
	QuotaSize *float64 `json:"quotaSize,omitempty"`
}

// ProtoReq name.
func (m StorageOverrideQuotaForOrigin) ProtoReq() string { return "Storage.overrideQuotaForOrigin" }

// Call sends the request.
func (m StorageOverrideQuotaForOrigin) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// StorageTrackCacheStorageForOrigin Registers origin to be notified when an update occurs to its cache storage list.
type StorageTrackCacheStorageForOrigin struct {
	// Origin Security origin.
	Origin string `json:"origin"`
}

// ProtoReq name.
func (m StorageTrackCacheStorageForOrigin) ProtoReq() string {
	return "Storage.trackCacheStorageForOrigin"
}

// Call sends the request.
func (m StorageTrackCacheStorageForOrigin) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// StorageTrackCacheStorageForStorageKey Registers storage key to be notified when an update occurs to its cache storage list.
type StorageTrackCacheStorageForStorageKey struct {
	// StorageKey Storage key.
	StorageKey string `json:"storageKey"`
}

// ProtoReq name.
func (m StorageTrackCacheStorageForStorageKey) ProtoReq() string {
	return "Storage.trackCacheStorageForStorageKey"
}

// Call sends the request.
func (m StorageTrackCacheStorageForStorageKey) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// StorageTrackIndexedDBForOrigin Registers origin to be notified when an update occurs to its IndexedDB.
type StorageTrackIndexedDBForOrigin struct {
	// Origin Security origin.
	Origin string `json:"origin"`
}

// ProtoReq name.
func (m StorageTrackIndexedDBForOrigin) ProtoReq() string { return "Storage.trackIndexedDBForOrigin" }

// Call sends the request.
func (m StorageTrackIndexedDBForOrigin) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// StorageTrackIndexedDBForStorageKey Registers storage key to be notified when an update occurs to its IndexedDB.
type StorageTrackIndexedDBForStorageKey struct {
	// StorageKey Storage key.
	StorageKey string `json:"storageKey"`
}

// ProtoReq name.
func (m StorageTrackIndexedDBForStorageKey) ProtoReq() string {
	return "Storage.trackIndexedDBForStorageKey"
}

// Call sends the request.
func (m StorageTrackIndexedDBForStorageKey) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// StorageUntrackCacheStorageForOrigin Unregisters origin from receiving notifications for cache storage.
type StorageUntrackCacheStorageForOrigin struct {
	// Origin Security origin.
	Origin string `json:"origin"`
}

// ProtoReq name.
func (m StorageUntrackCacheStorageForOrigin) ProtoReq() string {
	return "Storage.untrackCacheStorageForOrigin"
}

// Call sends the request.
func (m StorageUntrackCacheStorageForOrigin) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// StorageUntrackCacheStorageForStorageKey Unregisters storage key from receiving notifications for cache storage.
type StorageUntrackCacheStorageForStorageKey struct {
	// StorageKey Storage key.
	StorageKey string `json:"storageKey"`
}

// ProtoReq name.
func (m StorageUntrackCacheStorageForStorageKey) ProtoReq() string {
	return "Storage.untrackCacheStorageForStorageKey"
}

// Call sends the request.
func (m StorageUntrackCacheStorageForStorageKey) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// StorageUntrackIndexedDBForOrigin Unregisters origin from receiving notifications for IndexedDB.
type StorageUntrackIndexedDBForOrigin struct {
	// Origin Security origin.
	Origin string `json:"origin"`
}

// ProtoReq name.
func (m StorageUntrackIndexedDBForOrigin) ProtoReq() string {
	return "Storage.untrackIndexedDBForOrigin"
}

// Call sends the request.
func (m StorageUntrackIndexedDBForOrigin) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// StorageUntrackIndexedDBForStorageKey Unregisters storage key from receiving notifications for IndexedDB.
type StorageUntrackIndexedDBForStorageKey struct {
	// StorageKey Storage key.
	StorageKey string `json:"storageKey"`
}

// ProtoReq name.
func (m StorageUntrackIndexedDBForStorageKey) ProtoReq() string {
	return "Storage.untrackIndexedDBForStorageKey"
}

// Call sends the request.
func (m StorageUntrackIndexedDBForStorageKey) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// StorageGetTrustTokens (experimental) Returns the number of stored Trust Tokens per issuer for the
// current browsing context.
type StorageGetTrustTokens struct{}

// ProtoReq name.
func (m StorageGetTrustTokens) ProtoReq() string { return "Storage.getTrustTokens" }

// Call the request.
func (m StorageGetTrustTokens) Call(c Client) (*StorageGetTrustTokensResult, error) {
	var res StorageGetTrustTokensResult
	return &res, call(m.ProtoReq(), m, &res, c)
}

// StorageGetTrustTokensResult (experimental) ...
type StorageGetTrustTokensResult struct {
	// Tokens ...
	Tokens []*StorageTrustTokens `json:"tokens"`
}

// StorageClearTrustTokens (experimental) Removes all Trust Tokens issued by the provided issuerOrigin.
// Leaves other stored data, including the issuer's Redemption Records, intact.
type StorageClearTrustTokens struct {
	// IssuerOrigin ...
	IssuerOrigin string `json:"issuerOrigin"`
}

// ProtoReq name.
func (m StorageClearTrustTokens) ProtoReq() string { return "Storage.clearTrustTokens" }

// Call the request.
func (m StorageClearTrustTokens) Call(c Client) (*StorageClearTrustTokensResult, error) {
	var res StorageClearTrustTokensResult
	return &res, call(m.ProtoReq(), m, &res, c)
}

// StorageClearTrustTokensResult (experimental) ...
type StorageClearTrustTokensResult struct {
	// DidDeleteTokens True if any tokens were deleted, false otherwise.
	DidDeleteTokens bool `json:"didDeleteTokens"`
}

// StorageGetInterestGroupDetails (experimental) Gets details for a named interest group.
type StorageGetInterestGroupDetails struct {
	// OwnerOrigin ...
	OwnerOrigin string `json:"ownerOrigin"`

	// Name ...
	Name string `json:"name"`
}

// ProtoReq name.
func (m StorageGetInterestGroupDetails) ProtoReq() string { return "Storage.getInterestGroupDetails" }

// Call the request.
func (m StorageGetInterestGroupDetails) Call(c Client) (*StorageGetInterestGroupDetailsResult, error) {
	var res StorageGetInterestGroupDetailsResult
	return &res, call(m.ProtoReq(), m, &res, c)
}

// StorageGetInterestGroupDetailsResult (experimental) ...
type StorageGetInterestGroupDetailsResult struct {
	// Details ...
	Details *StorageInterestGroupDetails `json:"details"`
}

// StorageSetInterestGroupTracking (experimental) Enables/Disables issuing of interestGroupAccessed events.
type StorageSetInterestGroupTracking struct {
	// Enable ...
	Enable bool `json:"enable"`
}

// ProtoReq name.
func (m StorageSetInterestGroupTracking) ProtoReq() string { return "Storage.setInterestGroupTracking" }

// Call sends the request.
func (m StorageSetInterestGroupTracking) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// StorageGetSharedStorageMetadata (experimental) Gets metadata for an origin's shared storage.
type StorageGetSharedStorageMetadata struct {
	// OwnerOrigin ...
	OwnerOrigin string `json:"ownerOrigin"`
}

// ProtoReq name.
func (m StorageGetSharedStorageMetadata) ProtoReq() string { return "Storage.getSharedStorageMetadata" }

// Call the request.
func (m StorageGetSharedStorageMetadata) Call(c Client) (*StorageGetSharedStorageMetadataResult, error) {
	var res StorageGetSharedStorageMetadataResult
	return &res, call(m.ProtoReq(), m, &res, c)
}

// StorageGetSharedStorageMetadataResult (experimental) ...
type StorageGetSharedStorageMetadataResult struct {
	// Metadata ...
	Metadata *StorageSharedStorageMetadata `json:"metadata"`
}

// StorageGetSharedStorageEntries (experimental) Gets the entries in an given origin's shared storage.
type StorageGetSharedStorageEntries struct {
	// OwnerOrigin ...
	OwnerOrigin string `json:"ownerOrigin"`
}

// ProtoReq name.
func (m StorageGetSharedStorageEntries) ProtoReq() string { return "Storage.getSharedStorageEntries" }

// Call the request.
func (m StorageGetSharedStorageEntries) Call(c Client) (*StorageGetSharedStorageEntriesResult, error) {
	var res StorageGetSharedStorageEntriesResult
	return &res, call(m.ProtoReq(), m, &res, c)
}

// StorageGetSharedStorageEntriesResult (experimental) ...
type StorageGetSharedStorageEntriesResult struct {
	// Entries ...
	Entries []*StorageSharedStorageEntry `json:"entries"`
}

// StorageSetSharedStorageEntry (experimental) Sets entry with `key` and `value` for a given origin's shared storage.
type StorageSetSharedStorageEntry struct {
	// OwnerOrigin ...
	OwnerOrigin string `json:"ownerOrigin"`

	// Key ...
	Key string `json:"key"`

	// Value ...
	Value string `json:"value"`

	// IgnoreIfPresent (optional) If `ignoreIfPresent` is included and true, then only sets the entry if
	// `key` doesn't already exist.
	IgnoreIfPresent bool `json:"ignoreIfPresent,omitempty"`
}

// ProtoReq name.
func (m StorageSetSharedStorageEntry) ProtoReq() string { return "Storage.setSharedStorageEntry" }

// Call sends the request.
func (m StorageSetSharedStorageEntry) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// StorageDeleteSharedStorageEntry (experimental) Deletes entry for `key` (if it exists) for a given origin's shared storage.
type StorageDeleteSharedStorageEntry struct {
	// OwnerOrigin ...
	OwnerOrigin string `json:"ownerOrigin"`

	// Key ...
	Key string `json:"key"`
}

// ProtoReq name.
func (m StorageDeleteSharedStorageEntry) ProtoReq() string { return "Storage.deleteSharedStorageEntry" }

// Call sends the request.
func (m StorageDeleteSharedStorageEntry) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// StorageClearSharedStorageEntries (experimental) Clears all entries for a given origin's shared storage.
type StorageClearSharedStorageEntries struct {
	// OwnerOrigin ...
	OwnerOrigin string `json:"ownerOrigin"`
}

// ProtoReq name.
func (m StorageClearSharedStorageEntries) ProtoReq() string {
	return "Storage.clearSharedStorageEntries"
}

// Call sends the request.
func (m StorageClearSharedStorageEntries) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// StorageResetSharedStorageBudget (experimental) Resets the budget for `ownerOrigin` by clearing all budget withdrawals.
type StorageResetSharedStorageBudget struct {
	// OwnerOrigin ...
	OwnerOrigin string `json:"ownerOrigin"`
}

// ProtoReq name.
func (m StorageResetSharedStorageBudget) ProtoReq() string { return "Storage.resetSharedStorageBudget" }

// Call sends the request.
func (m StorageResetSharedStorageBudget) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// StorageSetSharedStorageTracking (experimental) Enables/disables issuing of sharedStorageAccessed events.
type StorageSetSharedStorageTracking struct {
	// Enable ...
	Enable bool `json:"enable"`
}

// ProtoReq name.
func (m StorageSetSharedStorageTracking) ProtoReq() string { return "Storage.setSharedStorageTracking" }

// Call sends the request.
func (m StorageSetSharedStorageTracking) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// StorageSetStorageBucketTracking (experimental) Set tracking for a storage key's buckets.
type StorageSetStorageBucketTracking struct {
	// StorageKey ...
	StorageKey string `json:"storageKey"`

	// Enable ...
	Enable bool `json:"enable"`
}

// ProtoReq name.
func (m StorageSetStorageBucketTracking) ProtoReq() string { return "Storage.setStorageBucketTracking" }

// Call sends the request.
func (m StorageSetStorageBucketTracking) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// StorageDeleteStorageBucket (experimental) Deletes the Storage Bucket with the given storage key and bucket name.
type StorageDeleteStorageBucket struct {
	// StorageKey ...
	StorageKey string `json:"storageKey"`

	// BucketName ...
	BucketName string `json:"bucketName"`
}

// ProtoReq name.
func (m StorageDeleteStorageBucket) ProtoReq() string { return "Storage.deleteStorageBucket" }

// Call sends the request.
func (m StorageDeleteStorageBucket) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// StorageCacheStorageContentUpdated A cache's contents have been modified.
type StorageCacheStorageContentUpdated struct {
	// Origin to update.
	Origin string `json:"origin"`

	// StorageKey Storage key to update.
	StorageKey string `json:"storageKey"`

	// CacheName Name of cache in origin.
	CacheName string `json:"cacheName"`
}

// ProtoEvent name.
func (evt StorageCacheStorageContentUpdated) ProtoEvent() string {
	return "Storage.cacheStorageContentUpdated"
}

// StorageCacheStorageListUpdated A cache has been added/deleted.
type StorageCacheStorageListUpdated struct {
	// Origin to update.
	Origin string `json:"origin"`

	// StorageKey Storage key to update.
	StorageKey string `json:"storageKey"`
}

// ProtoEvent name.
func (evt StorageCacheStorageListUpdated) ProtoEvent() string {
	return "Storage.cacheStorageListUpdated"
}

// StorageIndexedDBContentUpdated The origin's IndexedDB object store has been modified.
type StorageIndexedDBContentUpdated struct {
	// Origin to update.
	Origin string `json:"origin"`

	// StorageKey Storage key to update.
	StorageKey string `json:"storageKey"`

	// DatabaseName Database to update.
	DatabaseName string `json:"databaseName"`

	// ObjectStoreName ObjectStore to update.
	ObjectStoreName string `json:"objectStoreName"`
}

// ProtoEvent name.
func (evt StorageIndexedDBContentUpdated) ProtoEvent() string {
	return "Storage.indexedDBContentUpdated"
}

// StorageIndexedDBListUpdated The origin's IndexedDB database list has been modified.
type StorageIndexedDBListUpdated struct {
	// Origin to update.
	Origin string `json:"origin"`

	// StorageKey Storage key to update.
	StorageKey string `json:"storageKey"`
}

// ProtoEvent name.
func (evt StorageIndexedDBListUpdated) ProtoEvent() string {
	return "Storage.indexedDBListUpdated"
}

// StorageInterestGroupAccessed One of the interest groups was accessed by the associated page.
type StorageInterestGroupAccessed struct {
	// AccessTime ...
	AccessTime TimeSinceEpoch `json:"accessTime"`

	// Type ...
	Type StorageInterestGroupAccessType `json:"type"`

	// OwnerOrigin ...
	OwnerOrigin string `json:"ownerOrigin"`

	// Name ...
	Name string `json:"name"`
}

// ProtoEvent name.
func (evt StorageInterestGroupAccessed) ProtoEvent() string {
	return "Storage.interestGroupAccessed"
}

// StorageSharedStorageAccessed Shared storage was accessed by the associated page.
// The following parameters are included in all events.
type StorageSharedStorageAccessed struct {
	// AccessTime Time of the access.
	AccessTime TimeSinceEpoch `json:"accessTime"`

	// Type Enum value indicating the Shared Storage API method invoked.
	Type StorageSharedStorageAccessType `json:"type"`

	// MainFrameID DevTools Frame Token for the primary frame tree's root.
	MainFrameID PageFrameID `json:"mainFrameId"`

	// OwnerOrigin Serialized origin for the context that invoked the Shared Storage API.
	OwnerOrigin string `json:"ownerOrigin"`

	// Params The sub-parameters warapped by `params` are all optional and their
	// presence/absence depends on `type`.
	Params *StorageSharedStorageAccessParams `json:"params"`
}

// ProtoEvent name.
func (evt StorageSharedStorageAccessed) ProtoEvent() string {
	return "Storage.sharedStorageAccessed"
}

// StorageStorageBucketCreatedOrUpdated ...
type StorageStorageBucketCreatedOrUpdated struct {
	// Bucket ...
	Bucket *StorageStorageBucketInfo `json:"bucket"`
}

// ProtoEvent name.
func (evt StorageStorageBucketCreatedOrUpdated) ProtoEvent() string {
	return "Storage.storageBucketCreatedOrUpdated"
}

// StorageStorageBucketDeleted ...
type StorageStorageBucketDeleted struct {
	// BucketID ...
	BucketID string `json:"bucketId"`
}

// ProtoEvent name.
func (evt StorageStorageBucketDeleted) ProtoEvent() string {
	return "Storage.storageBucketDeleted"
}
