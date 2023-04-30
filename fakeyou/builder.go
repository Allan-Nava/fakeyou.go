package fakeyou

type IBuilder interface {
	Build() []string
}

/*type Builder fakeyou

func (f *Builder) Build() []string {
	var args []string
	return args
}*/
