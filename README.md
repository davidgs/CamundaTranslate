# Camunda Translator

This is designed to translate the Camunda Platform web apps to almost any language.

## Usage

```
$ git clone https://github.com/davidgs/CamundaTranslator
$ cd CamundaTranslator
$ go build
$ go install
$ CamundaTranslator -lang=de
```

That will get you a full German translation, with the relevant JSON files in the `output` directory.

## Requirements

You **must** have a google natural language translation 'secret' available. You can set up a project by following these [instructions](https://cloud.google.com/translate/docs/setup). That should produce the required `google-secret.json` file as well as the required `Project ID`. Enter those in the `config.yml` file before attempting to run the program.

The program will look in the the current working directory for the `config.yml` file, and will create the `output` directory in the same directory. 