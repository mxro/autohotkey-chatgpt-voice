#persistent

Capslock::Esc

#s::Send ^s  ; Win + S to Ctrl + S
#v::Send ^v  ; Win + V to Ctrl + V
#c::Send ^c  ; Win + C to Ctrl + C
#z::Send ^z  ; Win + Z to Ctrl + Z

F8::
  NotRecording := !NotRecording
  If NotRecording
  {
    Run %A_WorkingDir%\bin\fmedia-1.31-windows-x64\fmedia\fmedia.exe --record --overwrite --mpeg-quality=16 --rate=12000 --out=rec.mp3 --globcmd=listen,, Hide
  }
  Else
  {  
    Run %A_WorkingDir%\bin\fmedia-1.31-windows-x64\fmedia\fmedia.exe --globcmd=stop,, Hide
    Sleep, 100
    Run %A_WorkingDir%\bin\whisper-autohotkey\whisper-autohotkey.exe,, Hide
  }
  return

