package twitter

import "gopkg.in/mgo.v2/bson"

type Tweet struct {
	Id                          bson.ObjectId `bson:"_id"`
	Contributors                []int64       `bson:"contributors"`
	CreatedAt                   string        `bson:"created_at"`
	Entities                    Entities      `bson:"entities"`
	ExtendedEntities            Entities      `bson:"extended_entities"`
	FavoriteCount               int           `bson:"favorite_count"`
	Favorited                   bool          `bson:"favorited"`
	HasExtendedProfile          bool          `bson:"has_extended_profile"`
	InReplyToScreenName         string        `bson:"in_reply_to_screen_name"`
	InReplyToStatusID           int64         `bson:"in_reply_to_status_id"`
	InReplyToUserID             int64         `bson:"in_reply_to_user_id"`
	Lang                        string        `bson:"lang"`
	QuotedStatusID              int64         `bson:"quoted_status_id"`
	QuotedStatus                *Tweet        `bson:"quoted_status"`
	PossiblySensitive           bool          `bson:"possibly_sensitive"`
	PossiblySensitiveAppealable bool          `bson:"possibly_sensitive_appealable"`
	RetweetCount                int           `bson:"retweet_count"`
	Retweeted                   bool          `bson:"retweeted"`
	RetweetedStatus             *Tweet        `bson:"retweeted_status"`
	Source                      string        `bson:"source"`
	Text                        string        `bson:"text"`
	Truncated                   bool          `bson:"truncated"`
	User                        User          `bson:"user"`
}

type User struct {
	Id                             bson.ObjectId `bson:"_id"`
	ContributorsEnabled            bool          `bson:"contributors_enabled"`
	CreatedAt                      string        `bson:"created_at"`
	DefaultProfile                 bool          `bson:"default_profile"`
	DefaultProfileImage            bool          `bson:"default_profile_image"`
	Description                    string        `bson:"description"`
	FavouritesCount                int           `bson:"favourites_count"`
	FollowRequestSent              bool          `bson:"follow_request_sent"`
	FollowersCount                 int           `bson:"followers_count"`
	Following                      bool          `bson:"following"`
	FriendsCount                   int           `bson:"friends_count"`
	GeoEnabled                     bool          `bson:"geo_enabled"`
	HasExtendedProfile             bool          `bson:"has_extended_profile"`
	IsTranslator                   bool          `bson:"is_translator"`
	IsTranslationEnabled           bool          `bson:"is_translation_enabled"`
	Lang                           string        `bson:"lang"` // BCP-47 code of user defined language
	ListedCount                    int64         `bson:"listed_count"`
	Location                       string        `bson:"location"` // User defined location
	Name                           string        `bson:"name"`
	ProfileBackgroundImageURL      string        `bson:"profile_background_image_url"`
	ProfileBackgroundImageUrlHttps string        `bson:"profile_background_image_url_https"`
	ProfileBannerURL               string        `bson:"profile_banner_url"`
	ProfileImageURL                string        `bson:"profile_image_url"`
	ProfileImageUrlHttps           string        `bson:"profile_image_url_https"`
	Protected                      bool          `bson:"protected"`
	ScreenName                     string        `bson:"screen_name"`
	ShowAllInlineMedia             bool          `bson:"show_all_inline_media"`
	StatusesCount                  int64         `bson:"statuses_count"`
	TimeZone                       string        `bson:"time_zone"`
	URL                            string        `bson:"url"`
	UtcOffset                      int           `bson:"utc_offset"`
	Verified                       bool          `bson:"verified"`
}

type UrlEntity struct {
	Urls []struct {
		Indices      []int
		Url          string
		Display_url  string
		Expanded_url string
	}
}

type Entities struct {
	Hashtags []struct {
		Indices []int
		Text    string
	}
	Urls []struct {
		Indices      []int
		Url          string
		Display_url  string
		Expanded_url string
	}
	Url           UrlEntity
	User_mentions []struct {
		Name        string
		Indices     []int
		Screen_name string
		Id          int64
		Id_str      string
	}
	Media []EntityMedia
}

type EntityMedia struct {
	Id                   int64
	Id_str               string
	Media_url            string
	Media_url_https      string
	Url                  string
	Display_url          string
	Expanded_url         string
	Sizes                MediaSizes
	Source_status_id     int64
	Source_status_id_str string
	Type                 string
	Indices              []int
	VideoInfo            VideoInfo `json:"video_info"`
}

type MediaSizes struct {
	Medium MediaSize `bson:"medium"`
	Thumb  MediaSize `bson:"thumb"`
	Small  MediaSize `bson:"small"`
	Large  MediaSize `bson:"large"`
}

type VideoInfo struct {
	AspectRatio    []int     `json:"aspect_ratio"`
	DurationMillis int64     `json:"duration_millis"`
	Variants       []Variant `json:"variants"`
}

type MediaSize struct {
	W      int    `json:"w"`
	H      int    `json:"h"`
	Resize string `json:"resize"`
}

type Variant struct {
	Bitrate     int    `json:"bitrate"`
	ContentType string `json:"content_type"`
	Url         string `json:"url"`
}
