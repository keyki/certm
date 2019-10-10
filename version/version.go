package version

var (
	// Version should be updated by hand at each release
	Version = "0.2.0"
)

func FullVersion() string {
	return Version
}
