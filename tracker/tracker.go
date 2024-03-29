package tracker


type Tracker interface {
	TrackEvent(event string, params map[string]interface{})
	TrackLog(event string, params map[string]interface{})
	TrackError(event string, err error, params map[string]interface{})
}
