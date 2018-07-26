package Request

type QueryParameters struct {
	timeZone           string
	geoLocation        interface{}
	contexts           interface{}
	resetContexts      bool
	sessionEntityTypes []SessionEntityType
	payload            map[string]interface{}
}
