package ffmpeg

type IBuilder interface {
	Build() []string
}

type Builder ffmpeg

func (f *Builder) Build() []string {
	var args []string
	return args
}
