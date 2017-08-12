## c2m - Couchbase Cluster Monitor
A CLI utility to monitor Couchbase cluster. This application can help you get key insights about your Couchbase Cluster. These details can help you to proactively monitor your cluster and take appropriate step. 

## Create Executable / Build Project
$go install github.com/raiskumar/c2m
<br/> or
<br/> $go build

## Bootstrap CLI application
CLI application needs to know certain basic details before it starts spitting cluster insights. Your cluster might have 100s of nodes, does it mean you need to provide details of all nodes ? Certenly NOT!
<br />
<br /> You just need to provide base URL of any one node and credentials to access it (if configured)!
<br />
<br/> $./c2m config http://172.27.0.1:8091 Administrator Password123
<br />
<br/> If you haven't configured User credentials during the Couchbase setup then don't pass those details in above command.
<br /> Note: The application doesn't store the credentials in clear text!

## Supported Commands
$./c2m node
<br/> Gets the node related details of the couchbaase cluster

$./c2m bucket {optional_bucket_name}
<br/> Prints bucket related details of the cluster

$./c2m cluster
<br/> Prints cluster metadata

Above commands can also respond appropriately on usage of --verbose flag. 
<br/> $./c2m bucket {bucket_name} --verbose

## Get Help about any Command
To get help of any command
<br /> $./c2m command --help


## Libraries/References Used
Print the output on console in tabular format: https://github.com/olekukonko/tablewriter
<br />Important Couchbase REST End points: https://developer.couchbase.com/documentation/server/current/rest-api/rest-endpoints-all.html
<br />CLI support: https://github.com/spf13/cobra
<br />Monitoring Couchbase: https://blog.couchbase.com/monitoring-couchbase-cluster/ & https://dzone.com/articles/monitoring-couchbase-cluster
<br />Troubleshooting Issues: https://www.slideshare.net/Couchbase/experience-at-global-scale-powered-by-couchbase-mobile-couchbase-connect-2015?next_slideshow=1

