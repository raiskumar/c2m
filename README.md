## c2m - Couchbase Cluster Monitor
A CLI utility to monitor Couchbase cluster.

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
<br/> $./c2m node
<br/> Gets the node related details of the couchbaase cluster

get cluster details - ip:8091/pools/default
get cluster details and also the name of bucket and vBucketMap as well - ip:8091/pools/default/buckets (it gives everyting which pervious one gives)


## Get Help about any Command
To get help of any command
<br /> $./c2m command --help



### ------ 

------https://forums.couchbase.com/t/monitoring-alerts-through-api-or-cli/893
The health check of the whole cluster can be gotten by doing a curl against http://(ip_address):8091/pools/default/ ... remember to pass username and password.

if you want to see possible auto failover or logs go to:
http://(ip_address):8091/logs

if you want very fine grain detail.
http://(ip_address):8091/pools/default/buckets/(name_of_bucket)/nodes/(ip_address):8091/status
This is a minute stat and you will get lots of data from it.

------

For this discussion, we want to focus on monitoring but it’s important to note we do provide verbose Couchbase logging to facilitate application troubleshooting. These logs are stored in ‘/opt/couchbase/var/lib/couchbase/logs’. 

-----
how to monitor
https://blog.couchbase.com/monitoring-couchbase-cluster/
https://dzone.com/articles/monitoring-couchbase-cluster

-----
https://www.slideshare.net/Couchbase/best-practices-troubleshooting-your-couchbase-application-couchbase-connect-2015
----



----Libraries Used
Print the output on console in tabular format: https://github.com/olekukonko/tablewriter
Important Couchbase Urls: https://developer.couchbase.com/documentation/server/current/rest-api/rest-endpoints-all.html
CLI support: https://github.com/spf13/cobra

