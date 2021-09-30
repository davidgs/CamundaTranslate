package main

import (
	"context"
	"log"
	"io/ioutil"
	"fmt"

	"cloud.google.com/go/translate"
	"golang.org/x/text/language"
	"google.golang.org/api/option"
	"gopkg.in/yaml.v2"
)

// Configuration details
type Configuration struct {
	Credentials string `yaml:"Google Credentials"`
	ProjectId string `yaml:"Project ID"`
}


type Translator struct {}
// this is directly copy/pasted from Google example
func (t *Translator) TranslateTextWithModel(targetLanguage, text, model string) (string, error) {
	config := Configuration{}
  dat, err := ioutil.ReadFile("./config.yml")
	if err != nil {
		log.Fatal("No startup file: ", err)
	}
	err = yaml.Unmarshal(dat, &config)
	if err != nil {
		log.Fatal(err)
	}
	lang, err := language.Parse(targetLanguage)
	if err != nil {
		return "", fmt.Errorf("language.Parse: %v", err)
	}
	client, ctx, err := t.authTranslate(config.Credentials, config.ProjectId)
	if err != nil {
		return "", fmt.Errorf("translate.NewClient: %v", err)
	}
	defer client.Close()
	resp, err := client.Translate(ctx, []string{text}, lang, &translate.Options{
		Model: model, // Either "nmt" or "base".
	})
	if err != nil {
		return "", fmt.Errorf("Translate: %v", err)
	}
	if len(resp) == 0 {
		return "", nil
	}
	return resp[0].Text, nil
}

func (t *Translator) authTranslate(jsonPath, projectID string) (*translate.Client, context.Context, error) {
	ctx := context.Background()
	client, err := translate.NewClient(ctx, option.WithCredentialsFile(jsonPath))
	if err != nil {
		fmt.Println(err)
		return client, ctx, err
	}
	return client, ctx, nil

}