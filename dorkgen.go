package dorkgen

// EngineFactory is the main interface for
// search engine implementations.
type EngineFactory struct {
	tags []string
}

func (e *EngineFactory) concat(tag string, value string, quotes bool) string {
	if quotes {
		return tag + "\"" + value + "\""
	}

	return tag + value
}
