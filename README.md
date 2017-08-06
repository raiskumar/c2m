# c2m Couchbase Cluster Manager

To Create Executable
$go install github.com/raiskumar/c2m

To run command
./c2m Cluster

How to build the project:
$go build

Run commands
$./c2m cluster



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
