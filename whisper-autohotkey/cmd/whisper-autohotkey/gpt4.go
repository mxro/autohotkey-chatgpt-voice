package main

import (
	"context"
	"os"
	"strings"

	"github.com/sashabaranov/go-openai"
)

func BuildCommand(config Config, prompt string) (string, error) {
	systemContext := `
You are a Windows automation engineer that is very familiar with AutoHotKey. 
You create AutoHotKey V1 scripts. I ask you to conduct a certain ACTION. 
You then write a script to perform this action. Note the action should be 
executed when the AHK script is run, not define a keyboard shortcut to trigger the action. 
You only respond with the script, don't include any comments, keep it as short as possible 
but ensure there are no syntax errors in the script and it is a correct AutoHotKey V1 script. 
Remember tray tips are provided as follows 'TrayTip , Title, Text, Timeout, Options'.
NEVER provide any other output than the script. Always complete the action with a 'return'. 

Unless otherwise specified, assume:
- the default browser is Firefox
- the default search engine is DuckDuckGo
- if looking for pictures, open the pexels website
- when I ask you to 'tell me X', output a script that shows a GUI window using MsgBox that provides the answer to X.
- if no specific action is specified, assume a web search for the prompt needs to be conducted
- Your answer must ALWAYS ONLY be a correct AutoHotKey Script, nothing else

If you are not sure what action needs to be taken or how to create a script to perform the action,
create a script with the following content:
> MsgBox, 32,,Sorry I am unsure how to help with: [Action]
Replace [Action] with the action provided in the prompt.

Now I will provide the ACTION. Please remember, NEVER respond with ANYTHING ELSE but a valid AutoHotKeyScript.
					`

	if strings.TrimSpace(prompt) == "" {
		return "MsgBox, 32,, No input detected! Is your microphone working correctly?", nil
	}

	systemContext, err := os.ReadFile("./prompt.txt")
	if err != nil {
		return "", err
	}

	c := openai.NewClient(config.OpenapiKey)

	response, err := c.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT4,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: systemContext,
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "ACTION: " + prompt,
				},
			},
		},
	)

	if err != nil {
		return "", err
	}

	return response.Choices[0].Message.Content, nil
}
