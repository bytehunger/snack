package main

type Section struct {
	Type     string                 `yaml:"type" json:"type"`
	Content  map[string]interface{} `yaml:"content" json:"content"`
	Settings map[string]interface{} `yaml:"settings" json:"settings"`
	Inactive bool                   `yaml:"inactive" json:"inactive"`
}

func (s *Section) Render() error {
	return nil
}
