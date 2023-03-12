package main

type Section struct {
	Type    string            `yaml:"type" json:"type"`
	Content map[string]string `yaml:"content" json:"content"`
}

func (s *Section) Render() error {
	return nil
}
