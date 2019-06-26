package awss3analyzer

import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/s3"
    "log"
    "sync"
)

// Max 1k file names in the queue. File consumer will only emit new file name
// to this channel if there are under 1k in the queue. The bottleneck will be 
// the file consumer threads, not the S3 directory traversal thread.
var files = make(chan string, 1000)

// Start is the entry point for S3 analyzer
func Start(c *Config) error {
    var wg sync.WaitGroup
    sess := createSession(c.Region, c.Profile)
    bucketSession := &bucketSession{
        Session: sess,
        Config: c,
    }
    go bucketSession.allocateWork()

    for i := 0; i < c.NumWorkers; i++ {
        wg.Add(1)
        go func() {
            for file := range files {
                err := bucketSession.doSomething(file)
                if err != nil {
                    log.Fatal(err)
                }
        	}
            wg.Done()
        }()
    }
    wg.Wait()
    return nil
}

func (bs *bucketSession) allocateWork() {
    var nextToken  *string
    var objects    []*s3.Object
    var err        error

    for ok := true; ok; ok = nextToken != nil {
        objects, nextToken, err = bs.getObjects(nextToken)
        if err != nil {
            log.Fatal(err)
        }
        for _, object := range objects {
            fileName := *object.Key
            files <- fileName
        }
    }
    close(files)
}

func createSession(region *string, profile string) *session.Session {
    return session.Must(session.NewSessionWithOptions(session.Options{
        Config: aws.Config{
            Region: region,
        },
        Profile: profile,
    }))
}
