Capslock::Esc

F8::
  Run %A_WorkingDir%\bin\fmedia-1.31-windows-x64\fmedia\fmedia.exe --record --out=rec.mp3 --globcmd=listen,, Hide
  return

F9::
  Run %A_WorkingDir%\bin\fmedia-1.31-windows-x64\fmedia\fmedia.exe --globcmd=stop,, Hide
  Sleep, 100
  Run %A_WorkingDir%\whisper-autohotkey.exe,, Hide
  return


  /*Run cmd /k %A_WorkingDir%\whisper-autohotkey.exe,, */