package Request
//
//{
//
//// Union field input can be only one of the following:
//"audioConfig": {
//object(InputAudioConfig)
//},
//"text": {
//object(TextInput)
//},
//"event": {
//object(EventInput)
//}
//// End of list of possible types for union field input.
//}

type QueryInput struct {
	audioConfig interface{}
	text interface{}
	event interface{}
}