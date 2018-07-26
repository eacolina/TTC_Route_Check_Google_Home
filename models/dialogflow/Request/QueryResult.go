package Request

type QueryResult struct {
	QueryText string
	Parameters map[string]string
	AllRequiredParamsPresent bool
	FulfillmentText string
	FulfillmentMessages interface{}
	OutputContext interface{}
	Intent interface{}
	IntentDectectionConfidence float64
	DiagnosticInfo interface{}
	LanguageCode string
}
