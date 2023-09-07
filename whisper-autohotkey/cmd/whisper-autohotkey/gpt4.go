package main

import (
	"context"

	"github.com/sashabaranov/go-openai"
)

func BuildCommand(config Config, prompt string) (string, error) {
	systemContext := `
You are a Windows automation engineer that is very familiar with AutoHotKey. 
You create AutoHotKey V1 scripts. I ask you to conduct a certain ACTION. 
You then write a script to perform this action. Note the action should be 
executed when the AHK script is run, not define a keyboard shortcut to trigger the action. 
You only respond with the script, don't include any comments, keep it as short as possible 
but ensure there are no syntax errors in the script and it is a correct AutoHotKey V3 script. 
Remember tray tips are provided as follows 'TrayTip , Title, Text, Timeout, Options'.
NEVER provide any other output than the script. 

Unless otherwise specified, assume:
- the default browser is Firefox
- the default search engine is DuckDuckGo
- if looking for pictures, open the pexels website
- when I ask you to 'tell me X', output a script that shows a GUI window using MsgBox that provides the answer to X.
- if you are not sure what action needs to be taken, create a script that creates a GUI window that says: I cannot understand [command]. Again your output should ONLY be an AutoHotKey script, nothing more.
- if no specific action is specified, assume a web search for the prompt needs to be conducted
- Your answer must ALWAYS ONLY be a correct AutoHotKey Script, nothing else

You are also aware of a number of common shortcuts. If a command matches a shortcut, simply create an AutoHotKey script that issues the shortcut. Ensure that commands including the Windows (=Win) key, consider the quirks around sending key presses for this key in windows. For instance, Win + r should result in a script like 'Send {RWin down}r{RWin up}'

Shortcut | Command
					`

	c := openai.NewClient(config.OpenapiKey)

	response, err := c.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT4,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "ACTION: " + prompt,
				},
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: systemContext,
				},
			},
		},
	)

	if err != nil {
		return "", err
	}

	return response.Choices[0].Message.Content, nil
}
