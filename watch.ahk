﻿Capslock::Esc

F8::
  NotRecording := !NotRecording
  If NotRecording
  {
    Run %A_WorkingDir%\bin\fmedia-1.31-windows-x64\fmedia\fmedia.exe --record --overwrite --out=rec.mp3 --globcmd=listen,, Hide
  }
  Else
  {  
    Run %A_WorkingDir%\bin\fmedia-1.31-windows-x64\fmedia\fmedia.exe --globcmd=stop,, Hide
    Sleep, 100
    Run %A_WorkingDir%\bin\whisper-autohotkey\whisper-autohotkey.exe,, Hide
  }
  return

