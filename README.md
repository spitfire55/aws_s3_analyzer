# AWS S3 Analyzer

This code helps you enumerate an Amazon S3 Bucket and parse/analyze
the files within the bucket. How you parse or analyze the files is left up to
you.

### Options

* region - Set the region of your S3 bucket. Default is us-east-1
* profile - Set the name of the profile in your credentials file to access the
  S3 bucketAWS profile in your `~/.aws/credentials` file.
* rootBucket - Name of the bucket in your S3. Default is `example_bucket`.
* numWorkers - Number of threads to parse/analyze files. Default is 20.

