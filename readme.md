# A GO CLI For Calling OpenAI prompts 

Another #BuildInPublic / #LearnInPublic thing - throwing together a lightweight CLI for asking questions to OpenAI's API. Currently setup to work with 3.5. 

You need to export your OpenAI API key in order for it to work

```
export OPENAI_API_KEY=<key>
```

You can build the binary with

```
go build main.go
``` 

Open to suggestions! Need to add some better styling to the return to make it pop a bit more, and I'd love to add some additional helper prompts in instead of it being from scratch every time. It's a start though! 