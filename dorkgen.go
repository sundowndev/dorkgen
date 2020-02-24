package dorkgen

// EngineFactory ...
type EngineFactory interface {
	Site(string) *GoogleSearch
	ToString() string
	ToUrl() string
	Intext(string) *GoogleSearch
	Inurl(string) *GoogleSearch
	Filetype(string) *GoogleSearch
	Cache(string) *GoogleSearch
	Related(string) *GoogleSearch
	Ext(string) *GoogleSearch
	Exclude(string) *GoogleSearch
}
