package Request

type WebhookRequest struct {
	Session                     string
	ResponseID                  string
	QueryResult                 QueryResult
	OriginalDetectIntentRequest interface{}
}
