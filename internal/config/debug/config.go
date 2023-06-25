package debug

const (
	// BetaVersion beta version name
	BetaVersion string = "beta"
)

var (
	// AppName service name
	AppName string

	// Version describes version of application dev/beta/latest
	Version string = "beta"

	// GithubSHA describes the commit on which the build was built
	GithubSHA string

	// GithubSHAShort describes the commit on which the build was built
	GithubSHAShort string

	// BuildedAt describes date of build
	BuildedAt string
)
