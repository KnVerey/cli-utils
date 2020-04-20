// Copyright 2020 The Kubernetes Authors.
// SPDX-License-Identifier: Apache-2.0

package printers

import (
	"fmt"

	"k8s.io/cli-runtime/pkg/genericclioptions"
	"sigs.k8s.io/cli-utils/cmd/status/print"
	"sigs.k8s.io/cli-utils/pkg/kstatus/polling/collector"
)

// CreatePrinter return an implementation of the Printer interface. The
// actual implementation is based on the printerType requested.
func CreatePrinter(printerType string, collector *collector.ResourceStatusCollector,
	ioStreams genericclioptions.IOStreams) (print.Printer, error) {
	switch printerType {
	case "table":
		return NewTablePrinter(collector, ioStreams), nil
	default:
		return nil, fmt.Errorf("no printer available for output %q", printerType)
	}
}
