package logacef

import (
	"bufio"
	"time"
)

type LogACEF struct {
	fs                    string
	efs                   string
	eao                   string
	timestamp             time.Time
	hostname              string
	cef_version           string
	device_vendor         string
	device_product        string
	device_version        string
	device_event_class_id string
	name                  string
	severity              int

	_device_process_id string
	_file_handle       *bufio.Writer
	_min_severity      int
}

type CEFExtn map[string]string
