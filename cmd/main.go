package main

import (
    ".."
    "flag"
    "log"
)

var (
    flagRegion = flag.String("region", "us-east-1",
        "The region that your S3 bucket lives in.")
    flagProfile = flag.String("profile", "default",
        "The profile name you gave your AWS credentials.")
    flagRootBucket = flag.String("rootBucket", "example_bucket",
        "The name of the root bucket to enumerate.")
    flagNumWorkers = flag.Int("numWorkers", 20,
        "The number of concurrent workers to parse files.")
)

func main() {
    flag.Parse()

    config := &awss3analyzer.Config{
        Region: flagRegion,
        Profile: *flagProfile,
        RootBucket: flagRootBucket,
        NumWorkers: *flagNumWorkers,
    }

    err := awss3analyzer.Start(config)
    if err != nil {
        log.Fatal(err)
    }
}
