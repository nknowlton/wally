package main

import (
	"github.com/fdidron/webview"
	"wally/wally"
)

func handleRPC(w webview.WebView, data string, s *wally.State) {
	switch {
	case data == "close":
		w.Terminate()
	case data == "openFirmwareFile":
		var filter string
		if s.Device.Model == 0 {
			filter = "*.bin"
		}
		if s.Device.Model == 1 {
			filter = "*.hex"
		}

		firmwarePath := w.Dialog(webview.DialogTypeOpen, 0, "Select firmware file", "", filter)

		if firmwarePath != "" {
			s.SelectFirmware(firmwarePath)
			s.FlashFirmware()
			w.Dispatch(func() {
				w.Bind("state", s)
			})
		}
	}
}
