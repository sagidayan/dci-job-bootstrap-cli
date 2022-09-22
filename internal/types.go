package internal

import "errors"

type JobSettings struct {
	Topic     DCI_Topic `yaml:"dci_topic"`
	JobName   string    `yaml:"dci_name"`
	Tags      []string  `yaml:"dci_tags"`
	ConfigDir string    `yaml:"dci_config_dir"`
}

type RunScriptData struct {
	ConfigDir      string
	KubeconfigPath string
	JobName        string
}

type DCI_Topic string

const (
	OCP4_9  DCI_Topic = "OCP-4.9"
	OCP4_10 DCI_Topic = "OCP-4.10"
	OCP4_11 DCI_Topic = "OCP-4.11"
	OCP4_12 DCI_Topic = "OCP-4.12"
)

func (e *DCI_Topic) String() string {
	return string(*e)
}

func (e *DCI_Topic) Set(v string) error {
	switch v {
	case "OCP-4.9", "OCP-4.10", "OCP-4.11", "OCP-4.12":
		*e = DCI_Topic(v)
		return nil
	default:
		return errors.New(`Topic must be one of "OCP-4.9", "OCP-4.10", "OCP-4.11" or "OCP-4.12"`)
	}
}

// This function is used to print out CLI help
func (e *DCI_Topic) Type() string {
	return "DCI_Topic"
}
