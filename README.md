winston-logrus-formatter
========================

a simplified JSON logging formatter

The general message format has three fields:

* s: The logging status/level
* m: The logging message itself
* t: The recorded log time

Example:

`{"m":"starting standalone server...","s":"info","t":"2014-11-16T10:19:33-08:00"}`
