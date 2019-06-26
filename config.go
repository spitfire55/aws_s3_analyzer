package awss3analyzer

// Config contains settings passed from command line options (or defaults)
type Config struct {
    Region            *string
    Profile           string
    RootBucket        *string
    NumWorkers        int
}