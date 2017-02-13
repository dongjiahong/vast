// Package vast implements IAB VAST 3.0 specification http://www.iab.net/media/file/VASTv3.0.pdf
package vast

// VAST is the root <VAST> tag
type VAST struct {
	// The version of the VAST spec (should be either "2.0" or "3.0")
	Version string `xml:"version,attr" json:"version,omitempty"`
	// One or more Ad elements. Advertisers and video content publishers may
	// associate an <Ad> element with a line item video ad defined in contract
	// documentation, usually an insertion order. These line item ads typically
	// specify the creative to display, price, delivery schedule, targeting,
	// and so on.
	Ads []*Ad `xml:"Ad" json:"ads,omitempty"`
	// Contains a URI to a tracking resource that the video player should request
	// upon receiving a “no ad” response
	Errors []string `xml:"Error" json:"errors,omitempty"`
}

// Ad represent an <Ad> child tag in a VAST document
//
// Each <Ad> contains a single <InLine> element or <Wrapper> element (but never both).
type Ad struct {
	// An ad server-defined identifier string for the ad
	ID string `xml:"id,attr,omitempty" json:"id,omitempty"`
	// A number greater than zero (0) that identifies the sequence in which
	// an ad should play; all <Ad> elements with sequence values are part of
	// a pod and are intended to be played in sequence
	Sequence int      `xml:"sequence,attr,omitempty" json:"sequence,omitempty"`
	InLine   *InLine  `xml:",omitempty" json:"inline,omitempty"`
	Wrapper  *Wrapper `xml:",omitempty" json:"wrapper,omitempty"`
}

// InLine is a vast <InLine> ad element containing actual ad definition
//
// The last ad server in the ad supply chain serves an <InLine> element.
// Within the nested elements of an <InLine> element are all the files and
// URIs necessary to display the ad.
type InLine struct {
	// The name of the ad server that returned the ad
	AdSystem *AdSystem `json:"ad_system,omitempty"`
	// The common name of the ad
	AdTitle string `json:"ad_title,omitempty"`
	// One or more URIs that directs the video player to a tracking resource file that the
	// video player should request when the first frame of the ad is displayed
	Impressions []*Impression `xml:"Impression" json:"impressions,omitempty"`
	// The container for one or more <Creative> elements
	Creatives []*Creative `xml:"Creatives>Creative" json:"creatives,omitempty"`
	// A string value that provides a longer description of the ad.
	Description string `xml:",omitempty" json:"description,omitempty"`
	// The name of the advertiser as defined by the ad serving party.
	// This element can be used to prevent displaying ads with advertiser
	// competitors. Ad serving parties and publishers should identify how
	// to interpret values provided within this element. As with any optional
	// elements, the video player is not required to support it.
	Advertiser string `xml:",omitempty" json:"advertiser,omitempty"`
	// A URI to a survey vendor that could be the survey, a tracking pixel,
	// or anything to do with the survey. Multiple survey elements can be provided.
	// A type attribute is available to specify the MIME type being served.
	// For example, the attribute might be set to type=”text/javascript”.
	// Surveys can be dynamically inserted into the VAST response as long as
	// cross-domain issues are avoided.
	Survey string `xml:",omitempty" json:"survey,omitempty"`
	// A URI representing an error-tracking pixel; this element can occur multiple
	// times.
	Errors []string `xml:"Error,omitempty" json:"errors,omitempty"`
	// Provides a value that represents a price that can be used by real-time bidding
	// (RTB) systems. VAST is not designed to handle RTB since other methods exist,
	// but this element is offered for custom solutions if needed.
	Pricing string `xml:",omitempty" json:"pricing,omitempty"`
	// XML node for custom extensions, as defined by the ad server. When used, a
	// custom element should be nested under <Extensions> to help separate custom
	// XML elements from VAST elements. The following example includes a custom
	// xml element within the Extensions element.
	Extensions *Extensions `xml:",omitempty" json:"extensions,omitempty"`
}

// Impression is a URI that directs the video player to a tracking resource file that
// the video player should request when the first frame of the ad is displayed
type Impression struct {
	ID  string `xml:"id,attr,omitempty" json:"id,omitempty"`
	URI string `xml:",chardata" json:"url,omitempty"`
}

// Pricing provides a value that represents a price that can be used by real-time
// bidding (RTB) systems. VAST is not designed to handle RTB since other methods
// exist,  but this element is offered for custom solutions if needed.
type Pricing struct {
	// Identifies the pricing model as one of "cpm", "cpc", "cpe" or "cpv".
	Model string `xml:"model,attr" json:"model,omitempty"`
	// The 3 letter ISO-4217 currency symbol that identifies the currency of
	// the value provided
	Currency string `xml:"currency,attr" json:"currency,omitempty"`
	// If the value provided is to be obfuscated/encoded, publishers and advertisers
	// must negotiate the appropriate mechanism to do so. When included as part of
	// a VAST Wrapper in a chain of Wrappers, only the value offered in the first
	// Wrapper need be considered.
	Value string `xml:",chardata" json:"value,omitempty"`
}

// Wrapper element contains a URI reference to a vendor ad server (often called
// a third party ad server). The destination ad server either provides the ad
// files within a VAST <InLine> ad element or may provide a secondary Wrapper
// ad, pointing to yet another ad server. Eventually, the final ad server in
// the ad supply chain must contain all the necessary files needed to display
// the ad.
type Wrapper struct {
	// The name of the ad server that returned the ad
	AdSystem *AdSystem `json:"ad_system,omitempty"`
	// URL of ad tag of downstream Secondary Ad Server
	VASTAdTagURI string `json:"vast_ad_tag_url,omitempty"`
	// One or more URIs that directs the video player to a tracking resource file that the
	// video player should request when the first frame of the ad is displayed
	Impressions []*Impression `xml:"Impression" json:"impressions,omitempty"`
	// A URI representing an error-tracking pixel; this element can occur multiple
	// times.
	Errors []string `xml:"Error,omitempty" json:"errors,omitempty"`
	// The container for one or more <Creative> elements
	Creatives []*CreativeWrapper `xml:"Creatives>Creative" json:"creatives,omitempty"`
	// XML node for custom extensions, as defined by the ad server. When used, a
	// custom element should be nested under <Extensions> to help separate custom
	// XML elements from VAST elements. The following example includes a custom
	// xml element within the Extensions element.
	Extensions *Extensions `xml:",omitempty" json:"extensions,omitempty"`
}

// AdSystem contains information about the system that returned the ad
type AdSystem struct {
	Version string `xml:"version,attr,omitempty" json:"version,omitempty"`
	Name    string `xml:",chardata" json:"name,omitempty"`
}

// Creative is a file that is part of a VAST ad.
type Creative struct {
	// An ad server-defined identifier for the creative
	ID string `xml:"id,attr,omitempty" json:"id,omitempty"`
	// The preferred order in which multiple Creatives should be displayed
	Sequence int `xml:"sequence,attr,omitempty" json:"sequence,omitempty"`
	// Identifies the ad with which the creative is served
	AdID string `xml:"AdID,attr,omitempty" json:"adid,omitempty"`
	// The technology used for any included API
	APIFramework string `xml:"apiFramework,attr,omitempty" json:"api_framework,omitempty"`
	// If present, defines a linear creative
	Linear *Linear `xml:",omitempty" json:"linear,omitempty"`
	// If defined, defins companions creatives
	CompanionAds *CompanionAds `xml:",omitempty" json:"companionads,omitempty"`
	// If defined, defins non linear creatives
	NonLinearAds *NonLinearAds `xml:",omitempty" json:"nonlinearads,omitempty"`
}

// CompanionAds contains companions creatives
type CompanionAds struct {
	// Provides information about which companion creative to display.
	// All means that the player must attempt to display all. Any means the player
	// must attempt to play at least one. None means all companions are optional
	Required   string       `xml:"required,attr,omitempty" json:"required,omitempty"`
	Companions []*Companion `xml:"Companion,omitempty" json:"companions,omitempty"`
}

// NonLinearAds contains non linear creatives
type NonLinearAds struct {
	TrackingEvents []*Tracking `xml:"TrackingEvents>Tracking,omitempty" json:"tracking_events,omitempty"`
	// Non linear creatives
	NonLinears []NonLinear `xml:"NonLinear,omitempty" json:"nonlinear,omitempty"`
}

// CreativeWrapper defines wrapped creative's parent trackers
type CreativeWrapper struct {
	// An ad server-defined identifier for the creative
	ID string `xml:"id,attr,omitempty" json:"id,omitempty"`
	// The preferred order in which multiple Creatives should be displayed
	Sequence int `xml:"sequence,attr,omitempty" json:"sequence,omitempty"`
	// Identifies the ad with which the creative is served
	AdID string `xml:"AdID,attr,omitempty" json:"adid,omitempty"`
	// If present, defines a linear creative
	Linear *LinearWrapper `xml:",omitempty" json:"linear,omitempty"`
	// If defined, defins companions creatives
	CompanionAds *CompanionAdsWrapper `xml:"CompanionAds,omitempty" json:"compaionads,omitempty"`
	// If defined, defines non linear creatives
	NonLinearAds *NonLinearAdsWrapper `xml:"NonLinearAds,omitempty" json:"nonlinearads,omitempty"`
}

// CompanionAdsWrapper contains companions creatives in a wrapper
type CompanionAdsWrapper struct {
	// Provides information about which companion creative to display.
	// All means that the player must attempt to display all. Any means the player
	// must attempt to play at least one. None means all companions are optional
	Required   string              `xml:"required,attr,omitempty" json:"required,omitempty"`
	Companions []*CompanionWrapper `xml:"Companion,omitempty" json:"companions,omitempty"`
}

// NonLinearAdsWrapper contains non linear creatives in a wrapper
type NonLinearAdsWrapper struct {
	TrackingEvents []*Tracking `xml:"TrackingEvents>Tracking,omitempty" json:"tracking_events,omitempty"`
	// Non linear creatives
	NonLinears []*NonLinearWrapper `xml:"NonLinear,omitempty" json:"nonlinears,omitempty"`
}

// Linear is the most common type of video advertisement trafficked in the
// industry is a “linear ad”, which is an ad that displays in the same area
// as the content but not at the same time as the content. In fact, the video
// player must interrupt the content before displaying a linear ad.
// Linear ads are often displayed right before the video content plays.
// This ad position is called a “pre-roll” position. For this reason, a linear
// ad is often called a “pre-roll.”
type Linear struct {
	// To specify that a Linear creative can be skipped, the ad server must
	// include the skipoffset attribute in the <Linear> element. The value
	// for skipoffset is a time value in the format HH:MM:SS or HH:MM:SS.mmm
	// or a percentage in the format n%. The .mmm value in the time offset
	// represents milliseconds and is optional. This skipoffset value
	// indicates when the skip control should be provided after the creative
	// begins playing.
	SkipOffset *Offset `xml:"skipoffset,attr,omitempty" json:"skip_offset,omitempty"`
	// Duration in standard time format, hh:mm:ss
	Duration           *Duration           `json:"duration,omitempty"`
	AdParameters       *AdParameters       `xml:",omitempty" json:"ad_parameters,omitempty"`
	Icons              []*Icon             `json:"icons,omitempty"`
	TrackingEvents     []*Tracking         `xml:"TrackingEvents>Tracking,omitempty" json:"tracking_events,omitempty"`
	VideoClicks        *VideoClicks        `xml:",omitempty" json:"video_click,omitempty"`
	MediaFiles         []*MediaFile        `xml:"MediaFiles>MediaFile,omitempty" json:"media_files,omitempty"`
	CreativeExtensions *CreativeExtensions `xml:",omitempty" json:"creative_extension,omitempty"`
}

// LinearWrapper defines a wrapped linear creative
type LinearWrapper struct {
	Icons              []*Icon             `json:"icons,omitempty"`
	TrackingEvents     []*Tracking         `xml:"TrackingEvents>Tracking,omitempty" josn:"tracking_events,omitempty"`
	VideoClicks        *VideoClicks        `xml:",omitempty" json:"video_click,omitempty"`
	CreativeExtensions *CreativeExtensions `xml:",omitempty" json:"creative_extension,omitempty"`
}

// Companion defines a companion ad
type Companion struct {
	// Optional identifier
	ID string `xml:"id,attr,omitempty" json:"id,omitempty"`
	// Pixel dimensions of companion slot.
	Width int `xml:"width,attr" json:"width,omitempty"`
	// Pixel dimensions of companion slot.
	Height int `xml:"height,attr" json:"height,omitempty"`
	// Pixel dimensions of the companion asset.
	AssetWidth int `xml:"assetWidth,attr" json:"asset_width,omitempty"`
	// Pixel dimensions of the companion asset.
	AssetHeight int `xml:"assetHeight,attr" json:"asset_height,omitempty"`
	// Pixel dimensions of expanding companion ad when in expanded state.
	ExpandedWidth int `xml:"expandedWidth,attr" json:"expanded_width,omitempty"`
	// Pixel dimensions of expanding companion ad when in expanded state.
	ExpandeHeight int `xml:"expandedHeight,attr" json:"expanded_height,omitempty"`
	// The apiFramework defines the method to use for communication with the companion.
	APIFramework string `xml:"apiFramework,attr,omitempty" json:"api_framework,omitempty"`
	// Used to match companion creative to publisher placement areas on the page.
	AdSlotID string `xml:"adSlotId,attr,omitempty" json:"ad_slot_id,omitempty"`
	// URL to open as destination page when user clicks on the the companion banner ad.
	CompanionClickThrough string `xml:",omitempty" json:"companion_click_through,omitempty"`
	// Alt text to be displayed when companion is rendered in HTML environment.
	AltText string `xml:",omitempty" json:"alt_text,omitempty"`
	// The creativeView should always be requested when present. For Companions
	// creativeView is the only supported event.
	TrackingEvents []*Tracking `xml:"TrackingEvents>Tracking,omitempty" json:"tracking_events,omitempty"`
	// Data to be passed into the companion ads. The apiFramework defines the method
	// to use for communication (e.g. “FlashVar”)
	AdParameters *AdParameters `xml:",omitempty" json:"ad_parameters,omitempty"`
	// URL to a static file, such as an image or SWF file
	StaticResource *StaticResource `xml:",omitempty" json:"static_resource,omitempty"`
	// URL source for an IFrame to display the companion element
	IFrameResource string `xml:",omitempty" json:"iframe_resource,omitempty"`
	// HTML to display the companion element
	HTMLResource *HTMLResource `xml:",omitempty" json:"html_resource,omitempty"`
	// Extensions
	CreativeExtensions *CreativeExtensions `xml:",omitempty" json:"creative_extension,omitempty"`
}

// CompanionWrapper defines a companion ad in a wrapper
type CompanionWrapper struct {
	// Optional identifier
	ID string `xml:"id,attr,omitempty" json:"id,omitempty"`
	// Pixel dimensions of companion slot.
	Width int `xml:"width,attr" json:"width,omitempty"`
	// Pixel dimensions of companion slot.
	Height int `xml:"height,attr" json:"height,omitempty"`
	// Pixel dimensions of the companion asset.
	AssetWidth int `xml:"assetWidth,attr" json:"asset_width,omitempty"`
	// Pixel dimensions of the companion asset.
	AssetHeight int `xml:"assetHeight,attr" json:"asset_height,omitempty"`
	// Pixel dimensions of expanding companion ad when in expanded state.
	ExpandedWidth int `xml:"expandedWidth,attr" json:"expanded_width,omitempty"`
	// Pixel dimensions of expanding companion ad when in expanded state.
	ExpandeHeight int `xml:"expandedHeight,attr" json:"expanded_height,omitempty"`
	// The apiFramework defines the method to use for communication with the companion.
	APIFramework string `xml:"apiFramework,attr,omitempty" json:"api_framework,omitempty"`
	// Used to match companion creative to publisher placement areas on the page.
	AdSlotID string `xml:"adSlotId,attr,omitempty" json:"ad_slot_id,omitempty"`
	// URL to open as destination page when user clicks on the the companion banner ad.
	CompanionClickThrough string `xml:",omitempty" json:"companion_click_through,omitempty"`
	// URLs to ping when user clicks on the the companion banner ad.
	CompanionClickTracking []string `xml:",omitempty" json:"companion_click_trackings,omitempty"`
	// Alt text to be displayed when companion is rendered in HTML environment.
	AltText string `xml:",omitempty" json:"alt_text,omitempty"`
	// The creativeView should always be requested when present. For Companions
	// creativeView is the only supported event.
	TrackingEvents []*Tracking `xml:"TrackingEvents>Tracking,omitempty" json:"tracking_events,omitempty"`
	// Data to be passed into the companion ads. The apiFramework defines the method
	// to use for communication (e.g. “FlashVar”)
	AdParameters *AdParameters `xml:",omitempty" json:"ad_parameters,omitempty"`
	// URL to a static file, such as an image or SWF file
	StaticResource *StaticResource `xml:",omitempty" json:"static_resource,omitempty"`
	// URL source for an IFrame to display the companion element
	IFrameResource string `xml:",omitempty" json:"iframe_resource,omitempty"`
	// HTML to display the companion element
	HTMLResource       *HTMLResource       `xml:",omitempty" json:"html_resource,omitempty"`
	CreativeExtensions *CreativeExtensions `xml:",omitempty" json:"creative_extension,omitempty"`
}

// NonLinear defines a non linear ad
type NonLinear struct {
	// Optional identifier
	ID string `xml:"id,attr,omitempty" json:"id,omitempty"`
	// Pixel dimensions of companion.
	Width int `xml:"width,attr" json:"width,omitempty"`
	// Pixel dimensions of companion.
	Height int `xml:"height,attr" json:"height,omitempty"`
	// Pixel dimensions of expanding nonlinear ad when in expanded state.
	ExpandedWidth int `xml:"expandedWidth,attr" json:"expanded_width,omitempty"`
	// Pixel dimensions of expanding nonlinear ad when in expanded state.
	ExpandeHeight int `xml:"expandedHeight,attr" json:"expanded_height,omitempty"`
	// Whether it is acceptable to scale the image.
	Scalable bool `xml:"scalable,attr,omitempty" json:"scalable,omitempty"`
	// Whether the ad must have its aspect ratio maintained when scales.
	MaintainAspectRatio bool `xml:"maintainAspectRatio,attr,omitempty" json:"maintain_aspect_ratio,omitempty"`
	// Suggested duration to display non-linear ad, typically for animation to complete.
	// Expressed in standard time format hh:mm:ss.
	MinSuggestedDuration *Duration `xml:"minSuggestedDuration,attr,omitempty" json:"min_suggested_duration,omitempty"`
	// The apiFramework defines the method to use for communication with the nonlinear element.
	APIFramework string `xml:"apiFramework,attr,omitempty" json:"api_framework,omitempty"`
	// URLs to ping when user clicks on the the non-linear ad.
	NonLinearClickTracking []string `xml:",omitempty" json:"nonlinear_click_trackings,omitempty"`
	// URL to open as destination page when user clicks on the non-linear ad unit.
	NonLinearClickThrough string `xml:",omitempty" json:"nonlinear_click_through,omitempty"`
	// Data to be passed into the video ad.
	AdParameters *AdParameters `xml:",omitempty" json:"ad_parameters,omitempty"`
	// URL to a static file, such as an image or SWF file
	StaticResource *StaticResource `xml:",omitempty" json:"static_resource,omitempty"`
	// URL source for an IFrame to display the companion element
	IFrameResource string `xml:",omitempty" json:"iframe_resource,omitempty"`
	// HTML to display the companion element
	HTMLResource       *HTMLResource       `xml:",omitempty" json:"html_resource,omitempty"`
	CreativeExtensions *CreativeExtensions `xml:",omitempty" json:"creative_extension,omitempty"`
}

// NonLinearWrapper defines a non linear ad in a wrapper
type NonLinearWrapper struct {
	// Optional identifier
	ID string `xml:"id,attr,omitempty" json:"id,omitempty"`
	// Pixel dimensions of companion.
	Width int `xml:"width,attr" json:"width,omitempty"`
	// Pixel dimensions of companion.
	Height int `xml:"height,attr" json:"height,omitempty"`
	// Pixel dimensions of expanding nonlinear ad when in expanded state.
	ExpandedWidth int `xml:"expandedWidth,attr" json:"expanded_width,omitempty"`
	// Pixel dimensions of expanding nonlinear ad when in expanded state.
	ExpandeHeight int `xml:"expandedHeight,attr" json:"expanded_height,omitempty"`
	// Whether it is acceptable to scale the image.
	Scalable bool `xml:"scalable,attr,omitempty" json:"scalable,omitempty"`
	// Whether the ad must have its aspect ratio maintained when scales.
	MaintainAspectRatio bool `xml:"maintainAspectRatio,attr,omitempty" json:"maintain_aspect_ratio,omitempty"`
	// Suggested duration to display non-linear ad, typically for animation to complete.
	// Expressed in standard time format hh:mm:ss.
	MinSuggestedDuration *Duration `xml:"minSuggestedDuration,attr,omitempty" json"min_suggested_duration,omitempty"`
	// The apiFramework defines the method to use for communication with the nonlinear element.
	APIFramework string `xml:"apiFramework,attr,omitempty" json:"api_framework,omitempty"`
	// The creativeView should always be requested when present.
	TrackingEvents []*Tracking `xml:"TrackingEvents>Tracking,omitempty" json:"tracking_events,omitempty"`
	// URLs to ping when user clicks on the the non-linear ad.
	NonLinearClickTracking []string            `xml:",omitempty" json:"nonlinear_click_trackings,omitempty"`
	CreativeExtensions     *CreativeExtensions `xml:",omitempty" json:"creative_extension,omitempty"`
}

// Icon represents advertising industry initiatives like AdChoices.
type Icon struct {
	// Identifies the industry initiative that the icon supports.
	Program string `xml:"program,attr" json:"program,omitempty"`
	// Pixel dimensions of icon.
	Width int `xml:"width,attr" json:"width,omitempty"`
	// Pixel dimensions of icon.
	Height int `xml:"height,attr" json:"height,omitempty"`
	// The horizontal alignment location (in pixels) or a specific alignment.
	// Must match ([0-9]*|left|right)
	XPosition string `xml:"xPosition,attr" json:"x_position,omitempty"`
	// The vertical alignment location (in pixels) or a specific alignment.
	// Must match ([0-9]*|top|bottom)
	YPosition string `xml:"xPosition,attr" json:"y_position,omitempty"`
	// Start time at which the player should display the icon. Expressed in standard time format hh:mm:ss.
	Offset *Offset `xml:"offset,attr" json:"offset,omitempty"`
	// duration for which the player must display the icon. Expressed in standard time format hh:mm:ss.
	Duration *Duration `xml:"duration,attr" json:"duration,omitempty"`
	// The apiFramework defines the method to use for communication with the icon element
	APIFramework string `xml:"apiFramework,attr,omitempty" json:"api_framework,omitempty"`
	// URL to open as destination page when user clicks on the icon.
	IconClickThrough string `xml:"IconClicks>IconClickThrough,omitempty" json:"icon_click_through,omitempty"`
	// URLs to ping when user clicks on the the icon.
	IconClickTrackings []string `xml:"IconClicks>IconClickTracking,omitempty" json:"icon_click_trackings,omitempty"`
	// URL to a static file, such as an image or SWF file
	StaticResource *StaticResource `xml:",omitempty" json:"static_resource,omitempty"`
	// URL source for an IFrame to display the companion element
	IFrameResource string `xml:",omitempty" json:"iframe_resource,omitempty"`
	// HTML to display the companion element
	HTMLResource *HTMLResource `xml:",omitempty" json:"html_resource,omitempty"`
}

// Tracking defines an event tracking URL
type Tracking struct {
	// The name of the event to track for the element. The creativeView should
	// always be requested when present.
	//
	// Possible values are creativeView, start, firstQuartile, midpoint, thirdQuartile,
	// complete, mute, unmute, pause, rewind, resume, fullscreen, exitFullscreen, expand,
	// collapse, acceptInvitation, close, skip, progress.
	Event string `xml:"event,attr" json:"event,omitempty"`
	// The time during the video at which this url should be pinged. Must be present for
	// progress event. Must match (\d{2}:[0-5]\d:[0-5]\d(\.\d\d\d)?|1?\d?\d(\.?\d)*%)
	Offset *Offset `xml:"offset,attr,omitempty" json:"offset,omitempty"`
	URI    string  `xml:",chardata" json:"url,omitempty"`
}

// StaticResource is the URL to a static file, such as an image or SWF file
type StaticResource struct {
	// Mime type of static resource
	CreativeType string `xml:"creativeType,attr,omitempty" json:"creative_type,omitempty"`
	// URL to a static file, such as an image or SWF file
	URI string `xml:",chardata" json:"url,omitempty"`
}

// HTMLResource is a container for HTML data
type HTMLResource struct {
	// Specifies whether the HTML is XML-encoded
	XMLEncoded bool   `xml:"xmlEncoded,attr,omitempty" json:"xml_encoded,omitempty"`
	HTML       []byte `xml:",chardata" json:"html,omitempty"`
}

// AdParameters defines arbitrary ad parameters
type AdParameters struct {
	// Specifies whether the parameters are XML-encoded
	XMLEncoded bool   `xml:"xmlEncoded,attr,omitempty" json:"xml_encoded,omitempty"`
	Parameters []byte `xml:",chardata" json:"parameters,omitempty"`
}

// VideoClicks contains types of video clicks
type VideoClicks struct {
	ClickThroughs  []*VideoClick `xml:"ClickThrough,omitempty" json:"click_throughs,omitempty"`
	ClickTrackings []*VideoClick `xml:"ClickTracking,omitempty" json:"click_trackings,omitempty"`
	CustomClicks   []*VideoClick `xml:"CustomClick,omitempty" json:"custom_clicks,omitempty"`
}

// VideoClick defines a click URL for a linear creative
type VideoClick struct {
	ID  string `xml:"id,attr,omitempty" json:"id,omitempty"`
	URI string `xml:",chardata" json:"url,omitempty"`
}

// MediaFile defines a reference to a linear creative asset
type MediaFile struct {
	// Optional identifier
	ID string `xml:"id,attr,omitempty" json:"id,omitempty"`
	// Method of delivery of ad (either "streaming" or "progressive")
	Delivery string `xml:"delivery,attr" json:"delivery,omitempty"`
	// MIME type. Popular MIME types include, but are not limited to
	// “video/x-ms-wmv” for Windows Media, and “video/x-flv” for Flash
	// Video. Image ads or interactive ads can be included in the
	// MediaFiles section with appropriate Mime types
	Type string `xml:"type,attr" json:"type,omitempty"`
	// The codec used to produce the media file.
	Codec string `xml:"codec,attr,omitempty" json:"codec,omitempty"`
	// Bitrate of encoded video in Kbps. If bitrate is supplied, MinBitrate
	// and MaxBitrate should not be supplied.
	Bitrate int `xml:"bitrate,attr,omitempty" json:"bitrate"`
	// Minimum bitrate of an adaptive stream in Kbps. If MinBitrate is supplied,
	// MaxBitrate must be supplied and Bitrate should not be supplied.
	MinBitrate int `xml:"minBitrate,attr,omitempty" json:"min_bitrate"`
	// Maximum bitrate of an adaptive stream in Kbps. If MaxBitrate is supplied,
	// MinBitrate must be supplied and Bitrate should not be supplied.
	MaxBitrate int `xml:"maxBitrate,attr,omitempty" json:"max_bitrate"`
	// Pixel dimensions of video.
	Width int `xml:"width,attr" json:"width,omitempty"`
	// Pixel dimensions of video.
	Height int `xml:"height,attr" json:"height,omitempty"`
	// Whether it is acceptable to scale the image.
	Scalable bool `xml:"scalable,attr,omitempty" json:"scalable,omitempty"`
	// Whether the ad must have its aspect ratio maintained when scales.
	MaintainAspectRatio bool `xml:"maintainAspectRatio,attr,omitempty" json:"maintain_aspect_ratio,omitempty"`
	// The APIFramework defines the method to use for communication if the MediaFile
	// is interactive. Suggested values for this element are “VPAID”, “FlashVars”
	// (for Flash/Flex), “initParams” (for Silverlight) and “GetVariables” (variables
	// placed in key/value pairs on the asset request).
	APIFramework string `xml:"apiFramework,attr,omitempty" json:"api_framework,omitempty"`
	URI          string `xml:",chardata" json:"url,omitempty"`
}

// Extensions defines extensions
type Extensions struct {
	Extensions []*Extension `xml:",omitempty" json:"extensions,omitempty"`
}

// CreativeExtensions defines extensions for creatives
type CreativeExtensions struct {
	Extensions []*Extension `xml:"CreativeExtension,omitempty" json:"extensions,omitempty"`
}

// Extension represent aribtrary XML provided by the platform to extend the VAST response
type Extension struct {
	Data []byte `xml:",innerxml" json:"data,omitempty"`
}
