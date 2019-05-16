package version

const Maj = "0"
const Min = "4"
const Fix = "3"

var (
	// The full version string
	Version = "0.4.3"

	// GitCommit is set with --ldflags "-X main.gitCommit=$(git rev-parse HEAD)"
	GitCommit string
)

func init() {
	if GitCommit != "" {
		Version += "-" + GitCommit[:8]
	}
}
