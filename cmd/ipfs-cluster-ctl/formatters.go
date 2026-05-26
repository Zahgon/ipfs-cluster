package main

import (
	"github.com/ipfs-cluster/ipfs-cluster/api"
)

type addedOutputQuiet struct {
	api.AddedOutput
	quiet bool
}

func jsonFormatObject(resp interface{}) { _ = "STUB: not implemented"; return }

// print original objects as in JSON it makes
// no sense to have a human "quiet" output

func jsonFormatPrint(obj interface{}) { _ = "STUB: not implemented"; return }

func textFormatObject(resp interface{}) { _ = "STUB: not implemented"; return }

func textFormatPrintID(obj api.ID) { _ = "STUB: not implemented"; return }

func textFormatPrintGPInfo(obj api.GlobalPinInfo) { _ = "STUB: not implemented"; return }

func textFormatPrintVersion(obj api.Version) { _ = "STUB: not implemented"; return }

func textFormatPrintPin(obj api.Pin) { _ = "STUB: not implemented"; return }

func textFormatPrintAddedOutput(obj api.AddedOutput) { _ = "STUB: not implemented"; return }

func textFormatPrintAddedOutputQuiet(obj addedOutputQuiet) { _ = "STUB: not implemented"; return }

func textFormatPrintMetric(obj api.Metric) { _ = "STUB: not implemented"; return }

func textFormatPrintAlert(obj api.Alert) { _ = "STUB: not implemented"; return }

func textFormatPrintBandwidthByProtocol(obj api.BandwidthByProtocol) {
	_ = "STUB: not implemented"
	return
}

func textFormatPrintGlobalRepoGC(obj api.GlobalRepoGC) { _ = "STUB: not implemented"; return }

// If peer name is set, use it instead of peer ID.

// key.Key will be empty

func textFormatPrintError(obj api.Error) { _ = "STUB: not implemented"; return }

func trackerStatusAllString() string { _ = "STUB: not implemented"; return "" }
