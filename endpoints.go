package igdb

type endpoint string

// These are the available IGDB API endpoints.
const (
	CharacterEndpoint   endpoint = "characters/"
	CollectionEndpoint  endpoint = "collections/"
	CompanyEndpoint     endpoint = "companies/"
	CreditEndpoint      endpoint = "credits/"
	EngineEndpoint      endpoint = "game_engines/"
	FeedEndpoint        endpoint = "feeds/"
	FranchiseEndpoint   endpoint = "franchises/"
	GameEndpoint        endpoint = "games/"
	GameModeEndpoint    endpoint = "game_modes/"
	GenreEndpoint       endpoint = "genres/"
	KeywordEndpoint     endpoint = "keywords/"
	PageEndpoint        endpoint = "pages/"
	PersonEndpoint      endpoint = "people/"
	PlatformEndpoint    endpoint = "platforms/"
	PerspectiveEndpoint endpoint = "player_perspectives/"
	PulseEndpoint       endpoint = "pulses/"
	PulseGroupEndpoint  endpoint = "pulse_groups/"
	PulseSourceEndpoint endpoint = "pulse_sources/"
	ReleaseDateEndpoint endpoint = "release_dates/"
	ReviewEndpoint      endpoint = "reviews/"
	ThemeEndpoint       endpoint = "themes/"
	TitleEndpoint       endpoint = "titles/"
	VersionEndpoint     endpoint = "game_versions/"
)

// Count contains the number of objects
// of a certain type counted in the IGDB.
type Count struct {
	Count int `json:"count"`
}

// GetEndpointFieldList returns a list of fields that represent the
// model of the data available at the given IGDB endpoint.
func (c *Client) GetEndpointFieldList(end endpoint) ([]string, error) {
	url := c.rootURL + string(end) + "meta"

	var f []string

	err := c.get(url, &f)
	if err != nil && err != ErrNoResults {
		return nil, err
	}

	return f, nil
}

// GetEndpointCount returns the count of entities available for the given IGDB endpoint.
func (c *Client) GetEndpointCount(end endpoint, opts ...OptionFunc) (int, error) {
	url, err := c.countURL(end, opts...)
	if err != nil {
		return 0, err
	}

	var ct Count

	err = c.get(url, &ct)
	if err != nil {
		return 0, err
	}

	return ct.Count, nil
}
