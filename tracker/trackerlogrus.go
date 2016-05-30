package tracker


import (
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/stvp/rollbar"
)

type TrackerLogrusOptions struct {
	RemoteMode bool
	RollBarToken string
}

type TrackerLogrus struct {
	RemoteMode bool
	RollBarToken string
}

func NewTrackerLogrus(options TrackerLogrusOptions) TrackerLogrus{

	remoteMode   := false
	rollBarToken := ""

	if(options.RemoteMode){
		remoteMode = options.RemoteMode
		rollBarToken = options.RollBarToken
	}

	rollbar.Token = rollBarToken

	return TrackerLogrus{
		RemoteMode: remoteMode,
		RollBarToken: rollBarToken,
	}
}

func(tracker TrackerLogrus) TrackEvent(event string, params map[string]interface{}) {
	log := logrus.WithFields(logrus.Fields{"ns": "api.tracker", "at": "TrackEvent"})

	if params == nil {
		params = map[string]interface{}{}
	}

	log.WithFields(logrus.Fields{"event": event}).WithFields(logrus.Fields(params)).Info()

}

func(tracker TrackerLogrus) TrackLog(event string, params map[string]interface{}) {
	params["state"] = "log"

	if(tracker.RemoteMode){
		rollbar.Message("info", fmt.Sprintf("%s - %v", event, params))
	}

	tracker.TrackEvent(event, params)
}

func(tracker TrackerLogrus) TrackError(event string, err error, params map[string]interface{}) {
	params["error"] = fmt.Sprintf("%v", err)
	params["state"] = "error"


	if(tracker.RemoteMode){
		extraField := &rollbar.Field{"env", params}
		rollbar.Error(rollbar.ERR, err, extraField)
	}


	tracker.TrackEvent(event, params)
}

