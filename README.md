# Autohotkey ChatGPT

Use your voice to control Windows 📢

| You say                    | Windows does                                     |
| -------------------------- | ------------------------------------------------ |
| Open Firefox               | Opens Firefox                                    |
| Search for cupcake recipes | Opens browser and searches for 'Cupcake Recipes' |
| Tell me the first 20 digits of Pi                            | Shows a window with the 20 first digits of Pi                                                 |
| Paste a poem | Pastes a poem |

⚠ WARNING ⚠ This is an experimental application. ChatGPT can end up doing random stuff. So please use with caution!

## How does this work?

Check out my blog post: TBD

## Install

- Download and install AutoHotKey V1 from [autohotkey.com](https://www.autohotkey.com/)
- Download `AutoHotKey-ChatGPT.zip` from the [Releases](https://github.com/mxro/autohotkey-chatgpt-voice/releases) for the latest release.
- Extract `AutoHotKey-ChatGPT.zip`
- Edit `config.json` from the extracted files. Provide your [Open API Key](https://www.howtogeek.com/885918/how-to-get-an-openai-api-key/) for the property `OpenapiKey`.

```json
{
  "OpenapiKey": "",
  "AutoHotKeyExec": ".\\bin\\autohotkey-1.1.37.01\\AutoHotkeyU64.exe"
}
```

## Usage

- Double click on `watch.ahk` from the extracted files
- Press F8
- Speak into your microphone what you would like to do
- Press F8
- Wait for Open AI and AutoHotKey to do their magic

## Customise

### Trigger Hotkey

The hotkey to start/stop a voice command is defined in `watch.ahk`. You can replace the following with a hotkey of your choice:

```
F8::
```

### Prompt

Among the extracted files, there is a `prompt.txt`. You can edit this to customise it to yor own needs.

For instance, the current prompt file defaults to using the [DuckDuckGo](https://duckduckgo.com/) search engine. You can change this easily to any search engine you like, by modifying the following line in `prompt.txt`:

```
Unless otherwise specified, assume:
...
- the default search engine is DuckDuckGo
...
```

## Prior Art

- [ChatGPT-AutoHotkey-Utility](https://github.com/kdalanon/ChatGPT-AutoHotkey-Utility): Uses AutoHotKey to perform a number of actions, such as translate
- [ChatGPT Voice Assistant](https://github.com/DonGuillotine/chatGPT_whisper_AI_voice_assistant): Provides a Windows based assistant driven by ChatGPT
- [How to Make Your Own Windows Transcription App With Whisper and AutoHotkey](https://www.makeuseof.com/make-transcription-app-whisper-autohotkey/): Step by step tutorial to make a transcription app using AutoHotKey (added as per [reddit](https://www.reddit.com/r/AutoHotkey/comments/16ork8y/combining_ahk_with_chatgpt_to_automated_windows/))


## Develop

### Build Source Code

`task build`

### Package Executable

`task package`

### Run Locally

```
go run ./cmd/whisper-autohotkey/.
```
