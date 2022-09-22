/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"text/template"

	"github.com/rh-ecosystem-edge/dci-bootstrap/internal"
	"github.com/spf13/cobra"
)

// jobCmd represents the job command
var jobCmd = &cobra.Command{
	Use:   "job",
	Short: "Create a new Job folder",
	RunE: func(cmd *cobra.Command, args []string) error {
		path, err := os.Getwd()
		if err != nil {
			log.Println(err)
			return err
		}
		jobFolderPath = fmt.Sprintf("%v/dci_%v", path, name)
		hooksFolderPath := fmt.Sprintf("%v/hooks", jobFolderPath)
		if _, err := os.Stat(jobFolderPath); err != nil {
			// Job not exists...
			log.Printf("Generating new job '%v'", name)
			err = os.MkdirAll(hooksFolderPath, 0755)
			if err != nil {
				return err
			}
			return createJob()
		} else {
			if force {
				log.Printf("Job %v already exists at '%v'. Forcing creation...", name, jobFolderPath)
				return createJob()
			} else {
				log.Printf("Job %v already exists at '%v'. Use -f to force. Aborting.", name, jobFolderPath)
			}
		}
		return nil
	},
}
var (
	name          string
	force         bool
	jobFolderPath string
	dciTopic      internal.DCI_Topic
	tags          []string
)

func createJob() error {
	if err := generateHooks(); err != nil {
		return err
	}
	if err := generateSettings(); err != nil {
		return err
	}
	if err := generateRunScript(); err != nil {
		return err
	}
	log.Printf(`Successfully generated job in '%v'.
Please edit settings.yml and hooks as you see fit.

To run the job execute %v/run_job.sh`, jobFolderPath, jobFolderPath)
	return nil
}

func generateHooks() error {
	fileNames := []string{
		"install",
		"pre-run",
		"tests",
		"post-run",
		"teardown",
	}
	err := os.MkdirAll(fmt.Sprintf("%v/hooks", jobFolderPath), 0755)
	if err != nil {
		return err
	}
	for _, file := range fileNames {
		if err := internal.GenerateEmptyYML(fmt.Sprintf("%v/hooks/%v.yml", jobFolderPath, file)); err != nil {
			return err
		}
	}
	return nil
}

func generateSettings() error {

	settings := internal.JobSettings{
		JobName:   name,
		Tags:      append(tags, "debug"),
		Topic:     internal.DCI_Topic(dciTopic),
		ConfigDir: jobFolderPath,
	}
	return internal.GenerateSettings(settings, fmt.Sprintf("%v/settings.yml", jobFolderPath))
}

func generateRunScript() error {
	data := internal.RunScriptData{
		ConfigDir:      jobFolderPath,
		KubeconfigPath: fmt.Sprintf("%v/kubeconfig", jobFolderPath),
		JobName:        name,
	}
	scriptPath := fmt.Sprintf("%v/run_job.sh", jobFolderPath)
	t, err := template.New("runscript").Parse(internal.RunScriptTemplate)
	if err != nil {
		return err
	}
	file, err := os.Create(scriptPath)
	if err != nil {
		return err
	}
	defer file.Close()
	if err = t.Execute(file, data); err != nil {
		return err
	}
	err = os.Chmod(scriptPath, 0755)
	if err != nil {
		return err
	}
	return nil
}

func init() {
	newCmd.AddCommand(jobCmd)
	jobCmd.Flags().BoolVarP(&force, "force", "f", false, "Force create. Might override existing jobs")
	jobCmd.Flags().StringVarP(&name, "name", "n", "", "Job Name")
	jobCmd.MarkFlagRequired("name")
	jobCmd.Flags().VarP(&dciTopic, "topic", "t", "DCI topic. i.e. OCP-4.11")
	jobCmd.MarkFlagRequired("topic")
	jobCmd.Flags().StringSliceVarP(&tags, "tags", "", []string{}, "Job Tags. Comma seperated.")
}
