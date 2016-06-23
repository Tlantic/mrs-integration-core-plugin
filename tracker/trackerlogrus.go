package tracker


import (
	"github.com/Sirupsen/logrus"
	"github.com/sebest/logrusly"
)

type TrackerLogrusOptions struct {
	RemoteMode bool
	Token string
}

type TrackerLogrus struct {
	RemoteMode bool
	Token string
	Log *logrus.Logger
	Hook *logrusly.LogglyHook
}

func NewTrackerLogrus(options TrackerLogrusOptions) TrackerLogrus{

	remoteMode   := false
	token := ""

	lg := logrus.New()

	var hook *logrusly.LogglyHook


	if(options.RemoteMode){
		remoteMode = options.RemoteMode
		token = options.Token
		hook = logrusly.NewLogglyHook(token, "www.mrs-api.tlantic.com", logrus.WarnLevel, "integration", "filesender")
		lg.Hooks.Add(hook)
	}

	return TrackerLogrus{
		RemoteMode: remoteMode,
		Token: token,
		Log: lg,
		Hook: hook,
	}
}

func(tracker TrackerLogrus) TrackEvent(event string, params map[string]interface{}) {
	tracker.Log.WithFields(logrus.Fields{"ns": "api.tracker", "at": "TrackEvent"})

	if params == nil {
		params = map[string]interface{}{}
	}

	tracker.Log.WithFields(logrus.Fields{"event": event}).WithFields(logrus.Fields(params)).Error("Information")
	tracker.Hook.Flush()

}

func(tracker TrackerLogrus) TrackLog(event string, params map[string]interface{}) {
	params["state"] = "log"
	tracker.TrackEvent(event, params)
}

func(tracker TrackerLogrus) TrackError(event string, params map[string]interface{}) {
	params["state"] = "error"
	tracker.TrackEvent(event, params)
}

