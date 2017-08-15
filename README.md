[![Go Report Card](https://goreportcard.com/badge/github.com/raiskumar/c2m)](https://goreportcard.com/report/github.com/raiskumar/c2m)
[![GoDoc](https://godoc.org/github.com/raiskumar/c2m?status.svg)](https://godoc.org/github.com/raiskumar/c2m)

## c2m :<i>Couchbase Cluster Monitor</i>

A CLI utility to monitor Couchbase cluster. Tool helps you get key insights about your Couchbase Cluster. These details can help you to proactively monitor your cluster and take appropriate step. 

## Create Executable / Build Project
$go install github.com/raiskumar/c2m
<br/> or
<br/> $go build

## Bootstrap CLI application
CLI application needs to know certain basic details before it starts spitting cluster insights. Your cluster might have 100s of nodes, <b>does it mean you need to provide details of all nodes ? Certenly NOT!</b>
<br />
<br /> You just need to provide base URL of any one node and credentials to access it (if configured)!
<br />
<br/> $./c2m config http://172.27.0.1:8091 Administrator Password123
<br />
<br /> <b>Note:</b> The application doesn't store the credentials in clear text!

## Supported Commands
<b>$./c2m node</b>
<br/> Gets the node related details of the couchbaase cluster

<b>$./c2m bucket {optional_bucket_name}</b>
<br/> Prints bucket related details of the cluster

<b>$./c2m cluster</b>
<br/> Prints cluster metadata

<b>$./c2m index</b>
<br/> Prints Index details of the Cluster

Above commands can also respond appropriately on usage of --verbose flag. 
<br/> $./c2m bucket {bucket_name} --verbose

## Get Help about any Command
To get help of any command
<br /> <b> $./c2m {command} --help </b>
<br /> or
<br /> <b> $./c2m --help</b>


## Libraries/References Used for the Tool
Print the output on console in tabular format: https://github.com/olekukonko/tablewriter
<br />Important Couchbase REST End points: https://developer.couchbase.com/documentation/server/current/rest-api/rest-endpoints-all.html
<br />CLI support: https://github.com/spf13/cobra
<br />Monitoring Couchbase: https://blog.couchbase.com/monitoring-couchbase-cluster/ 
<br />& https://blog.couchbase.com/top-10-things-ops-sys-admin-must-know-about-couchbase/
<br />Troubleshooting Issues: https://www.youtube.com/watch?v=88NHDWz52aY

