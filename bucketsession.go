package awss3analyzer

import(
    "fmt"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/s3"
)

type bucketSession struct {
    Session *session.Session
    Config  *Config
}

func (bs *bucketSession) getObjects(cToken *string) (
    []*s3.Object,
    *string,
    error,
) {
    svc := s3.New(bs.Session)
    input := &s3.ListObjectsV2Input{
        Bucket: bs.Config.RootBucket,
    }
    if cToken != nil {
        input.ContinuationToken = cToken
    }
    result, err := svc.ListObjectsV2(input)
    if err != nil {
        return nil, nil, err
    }
    if *result.IsTruncated {
        return result.Contents, result.NextContinuationToken, nil
    }
    return result.Contents, nil, nil
}

func (bs *bucketSession) doSomething(fileName string) error {
    // IMPLEMENT FILE ANALYZER LOGIC HERE
    return fmt.Errorf("unimplemented analyzer")
}