package gits

const (
	KindBitBucket = "bitbucket"
	KindGitea     = "gitea"
	KindGitlab    = "gitlab"
	KindGitHub    = "github"

	DateFormat = "January 2 2006"
)

var (
	KindGits = []string{KindBitBucket, KindGitea, KindGitHub, KindGitlab}
)
