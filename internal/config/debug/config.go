package debug

const (
	// BetaVersion beta version name
	BetaVersion = "beta"
)

var (
	// AppName service name
	AppName string

	// Version describes version of application dev/beta/latest
	Version string

	// GithubSHA describes the commit on which the build was built
	GithubSHA string

	// GithubSHAShort describes the commit on which the build was built
	GithubSHAShort string

	// BuildTime describes date of build
	BuildTime string
)
