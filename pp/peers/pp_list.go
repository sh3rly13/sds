package peers

import (
	"github.com/alex023/clock"
	"github.com/stratosnet/sds/pp/client"
	"github.com/stratosnet/sds/pp/event"
	"github.com/stratosnet/sds/pp/setting"
	"github.com/stratosnet/sds/utils"
	"time"
)

// InitPPList
func initPPList() {
	pplist := setting.GetLocalPPList()
	if len(pplist) == 0 {
		event.GetPPList()
	} else {
		for _, ppInfo := range pplist {
			client.PPConn = client.NewClient(ppInfo.NetworkAddress, true)
			if client.PPConn == nil {

				setting.DeletePPList(ppInfo.NetworkAddress)
			} else {
				event.RegisterChain(false)
				return
			}
		}

		event.GetPPList()
	}
}

func startStatusReportToSP() {
	utils.DebugLog("Status will be reported to SP while mining")
	// trigger first report at time-0 immediately
	event.ReportNodeStatus()
	// trigger consecutive reports with interval
	clock := clock.NewClock()
	clock.AddJobRepeat(time.Second*event.ReportIntervalSec, 0, event.ReportNodeStatus)
}
