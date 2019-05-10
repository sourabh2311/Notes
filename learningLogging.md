# Learning Logging

## **glog**

* Error logs to the ERROR, WARNING, and INFO logs.
* Warning logs to the WARNING and INFO logs.
* V reports whether verbosity at the call site is at least the requested level. The returned value is a boolean of type Verbose, which implements Info, Infoln and Infof. These methods will write to the Info log if called. Thus, one may write either

  `if glog.V(2) { glog.Info("log this") }`

  or

  `glog.V(2).Info("log this")`
  
  The second form is shorter but the first is cheaper if logging is off because it does not evaluate its arguments.

  Whether an individual call to V generates a log record depends on the setting of the -v and --vmodule flags; both are off by default. If the level in the call to V is at least the value of -v, or of -vmodule for the source file containing the call, the V call will log.

  Basically -v thing should be >= V thing

* `-logtostderr` log to standard error instead of files
* `-alsologtostderr` log to standard error as well as files
* `-stderrthreshold value` - logs at or above this threshold go to stderr

  So, if the value is `FATAL` then all error level of less degree i.e. INFO, WARNING, ERROR (incl FATAL) will go to stderr.

* ```go
  /*
  glog-example
  ------------
  background
  ---
  setup
  ---
    $ go get github.com/golang/glog
    $ mkdir log
  run
  ---
    $ go run example.go -stderrthreshold=FATAL -log_dir=./log
  or
    $ go run example.go -stderrthreshold=FATAL -log_dir=./log -v=2
  or
    $ go run example.go -logtostderr // still wont print INFO
  or
    $ go run example.go -logtostderr -v=2 // even -v=3 will print info logs as mentioned before
  or
    $ go run example.go -log_dir='/Users/sourabhaggarwal/Desktop/' -v=2 // even -v=3 will print info logs as mentioned before but this time only error logs get printed on terminal and to see info logs we need to navigate to the log_dir and check the info file
  or
    $ go run example.go -v=3 -stderrthreshold=FATAL // this will not print anything as mentioned before.
  */

  package main

  import (
    "flag"
    "fmt"
    "os"

    "github.com/golang/glog"
  )

  func usage() {
    fmt.Fprintf(os.Stderr, "usage: example -stderrthreshold=[INFO|WARN|FATAL] -log_dir=[string]\n")
    flag.PrintDefaults()
    os.Exit(2)
  }

  func init() {
    flag.Usage = usage
    // NOTE: This next line is key you have to call flag.Parse() for the command line
    // options or "flags" that are defined in the glog module to be picked up.
    flag.Parse()
  }

  func main() {
    // usage()
    number_of_lines := 10
    for i := 0; i < number_of_lines; i++ {
      glog.V(2).Infof("LINE: %d", i)
      message := fmt.Sprintf("TEST LINE: %d", i)
      glog.Error(message)
    }
    glog.Flush()
  }

  ```

## **logrus**
