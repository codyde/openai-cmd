# A GO CLI For Calling OpenAI prompts 

![Go CLI for calling OpenAI](https://user-images.githubusercontent.com/17350652/230741456-a62cc7a5-83eb-44bb-bf0e-1b30ad6a3244.gif)

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
