package cli

type TextPrompter interface {
	Run() (string, error)
}

type SelectPrompter interface {
	Run() (int, string, error)
}
