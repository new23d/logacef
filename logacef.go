package logacef

import (
	"bufio"
	"bytes"
	"os"
	"strconv"
	"time"
)

func NewLogACEF(device_vendor string, device_product string, device_version string, filepath string, min_severity int) *LogACEF {
	var er error

	ev := new(LogACEF)
	ev.cef_version = "0"
	ev.fs = "|"
	ev.efs = " "
	ev.eao = "="

	ev._device_process_id = strconv.Itoa(os.Getpid())

	ev.hostname, er = os.Hostname()
	if er != nil {
		ev.hostname = "?"
	}

	ev.device_vendor = cef_fs_escape(device_vendor, ev.fs)
	ev.device_product = cef_fs_escape(device_product, ev.fs)
	ev.device_version = cef_fs_escape(device_version, ev.fs)

	log_file, er := os.OpenFile(filepath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0640)
	if er != nil {
		panic(er)
	}

	ev._file_handle = bufio.NewWriter(log_file)
	ev._min_severity = min_severity

	return ev
}

func (ev LogACEF) WriteEvent(deci string, desc string, sevr int, extn CEFExtn) {
	if sevr >= ev._min_severity {
		ev.timestamp = time.Now().UTC()

		ev.name = cef_fs_escape(desc, ev.fs)
		ev.device_event_class_id = deci
		ev.severity = sevr

		var extn_string bytes.Buffer

		extn_string.WriteString("dvcpid")
		extn_string.WriteString(ev.eao)
		extn_string.WriteString(ev._device_process_id)

		for k, v := range extn {
			extn_string.WriteString(ev.efs)
			extn_string.WriteString(cef_eao_escape(k, ev.eao))
			extn_string.WriteString(ev.eao)
			extn_string.WriteString(cef_eao_escape(v, ev.eao))
		}

		var file_line bytes.Buffer
		file_line.WriteString(ev.timestamp.Format(time.RFC3339))
		file_line.WriteString(ev.efs)
		file_line.WriteString(ev.hostname)
		file_line.WriteString(ev.efs)
		file_line.WriteString("CEF:")
		file_line.WriteString(ev.cef_version)
		file_line.WriteString(ev.fs)
		file_line.WriteString(ev.device_vendor)
		file_line.WriteString(ev.fs)
		file_line.WriteString(ev.device_product)
		file_line.WriteString(ev.fs)
		file_line.WriteString(ev.device_version)
		file_line.WriteString(ev.fs)
		file_line.WriteString(ev.device_event_class_id)
		file_line.WriteString(ev.fs)
		file_line.WriteString(ev.name)
		file_line.WriteString(ev.fs)
		file_line.WriteString(strconv.Itoa(ev.severity))
		file_line.WriteString(ev.fs)
		file_line.WriteString(extn_string.String())
		file_line.WriteByte('\n')

		ev._file_handle.Write(file_line.Bytes())

		defer ev._file_handle.Flush()
	}
}
