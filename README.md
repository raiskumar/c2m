## c2m - Couchbase Cluster Monitor
Go based utility to monitor your couchbase cluster. 

## Create Executable / Build Project
$go install github.com/raiskumar/c2m

or

$go build

## Run Commands
### 1. Configure the cluster by passing ip address, user id and password
`$./c2m config
$ go run main.go node`

### 2. Get Node details of the cluster
$./c2m node


get cluster details - ip:8091/pools/default
get cluster details and also the name of bucket and vBucketMap as well - ip:8091/pools/default/buckets (it gives everyting which pervious one gives)



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

