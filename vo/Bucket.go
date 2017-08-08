package vo

type Bucket struct {
	name       string // Name of the bucket, if any
	bucketType string //Once a memcached or couchbase bucket has been created, its type cannot be changed.
	ops        int    //Operations per second
	dfps       int    // Disk fetches per second
	//Indicates how frequently Couchbase Server is reaching to disk to retrieve information instead of using the information stored in RAM.
}
