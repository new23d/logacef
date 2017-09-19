# LogACEF

An opinionated and simple implementation for writing **Log** files in the **A**rcSight **C**ommon **E**vent **F**ormat (CEF), in Go, with ingestion of logs into indexing backends in mind.

_Based on Version 24 of the format dated August 22, 2017 found at ArcSight Common Event Format (CEF) Guide [1]. The direct link to the specification at the time is at [2]. The specification document is also included with this source for persistence._

[1] https://community.saas.hpe.com/t5/ArcSight-Connectors/ArcSight-Common-Event-Format-CEF-Guide/ta-p/1589306

[2] https://community.saas.hpe.com/dcvta86296/attachments/dcvta86296/connector-documentation/1116/2/CommonEventFormatV24.pdf

# Usage
```
package main

import "github.com/new23d/logacef"

func main() {
	/* Instantiating a new logger,
	   with NewLogACEF(device_vendor string, device_product string, device_version string, filepath string, min_severity int) */
	myAppLog := logacef.NewLogACEF("AcmeInc", "myApp", "2.1.4", "/var/log/myApp/myApp.log", 5)

	/* Writing a new event into the log file,
	   with WriteEvent(Device_Event_Class_ID/deci string, Name/desc string, Severity/sevr int, Extension/extn CEFExtn/map[string]string) */
	/* This event will NOT be written since the minimum severity level set in the instance is 5 */
	myAppLog.WriteEvent("-", "config file parsed", 3, logacef.CEFExtn{"spt": "8080"})

	/* This event will be written since the severity level of 7 is greater than or equal to the minimum severity level set in the instance */
	myAppLog.WriteEvent("-", "user authentication failed", 7, logacef.CEFExtn{"duser": "joe.bloggs", "dpriv": "guest"})
}
```
# Notes

* Several fields such as the _timestamp_, _hostname_, _PID_, etc. are determined automatically for accuracy and convenience.
* CEF Key Names For Event Producers are not prescibed and left up to the user to choose from the specification document or coin their own as they see fit.
* Log file mode will be `0640`.
* Log file is opened in append mode and flushed on a fully formatted line, allowing lightweight log-rotation to work.
* Concurrency behaviour is currently unknown.
* Timestamp will be in the [RFC3339](https://www.ietf.org/rfc/rfc3339.txt) profile of ISO8601, and in UTC.
