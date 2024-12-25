
F8::
  NotRecording := !NotRecording
  If NotRecording
  {
    Run %A_ScriptDir%\bin\fmedia-1.31-windows-x64\fmedia\fmedia.exe --record --overwrite --mpeg-quality=16 --rate=12000 --out=rec.mp3 --globcmd=listen,, Hide
  }
  Else
  {  
    Run %A_ScriptDir%\bin\fmedia-1.31-windows-x64\fmedia\fmedia.exe --globcmd=stop,, Hide
    Sleep, 100
    Run %A_ScriptDir%\bin\whisper-autohotkey\whisper-autohotkey.exe,, Hide
  }
  return

