package artifact

type Type int

const (
	UploadableArchive Type = iota
	UploadableBinary
	DockerImage
)

type Artifact struct {
	Name    string
	Version string
}
