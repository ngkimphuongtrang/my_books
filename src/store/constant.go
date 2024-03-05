package store

type ReadSource int

const (
	SourceHardCopy ReadSource = iota
	SourceSoftCopy
	SourceAudio
)

func (s ReadSource) String() string {
	return [...]string{"hard_copy", "soft_copy", "audio"}[s]
}

type BookLanguage int

const (
	LangVI BookLanguage = iota
	LangEN
)

func (l BookLanguage) String() string {
	return [...]string{"LangVI", "EN"}[l]
}
