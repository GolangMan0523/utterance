# Utterance:

Write a Go application which processes a time series of transcript events called utterances, analyzes them, fixes some attribution issues, and constructs a complete transcript.

Each utterance will contain a “speaker” property which identifies the speaker detected during the utterance. 

Sometimes a speaker change event happens after a word or two has already been spoken by a new speaker, so we need to correct the data.  

## For instance:

```json
    {
        "speaker": "bob@domain.com",
        "text": "How is the product launch going?  It is",
        "timestampMs": 1705947091584
    }

    {
        "speaker": "alice@domain.com",
        "text": "going well.  Thanks for asking.",
        "timestampMs": 1705947112584
    }
```

In this example, the complete transcript will reflect:

```json
    {"utterances": [
        {
            "speaker": "bob@domain.com",
            "text": "How is the product launch going?",
            "timestampMs": 1705947091584
        },
        {
             "speaker": "alice@domain.com",
             "text": "It is going well.  Thanks for asking.",
             "timestampMs": 1705947112584
        }]}
```

To achieve this, we need to first assess if the first word in a given utterance is a sentence fragment or not, and if so look back for the preceding words to construct the sentence correctly and modify both objects.

As an output, generate a complete.json file with an array of utterances with the properties in the above example (speaker, text, and timestampMs).

Here are the utterances to process.

## Requirements:

Use Go.
Provide a readme outlining any steps needed to run the application.
Code should be tested, well formatted, and documented so that it is readable and maintainable.


## How to run

To run the application, you need to have Go installed on your system. You can download and install Go.

You also need to have an input JSON file that contains a slice of utterances in `data` folder, each with a speaker, a text, and a timestamp in milliseconds. The input file should be named `input.json` and placed in the `data` folder. 
And you can see `output.json` as a result in the `data` folder after running application.

To build and run the application, open a terminal and navigate to the application directory. Then run the following command:

## run

```shell
go run main.go
```

## test

```shell
go test -v .\utterance\
```