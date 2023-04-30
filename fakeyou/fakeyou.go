package fakeyou

type IFakeYou interface {
}

type fakeyou struct {
}

func NewFakeYou() IFakeYou {
	return &fakeyou{}
}
