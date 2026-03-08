# qrcodefyi-go — Go API Client Demo
Output demo.gif

Set FontSize 20
Set Width 1000
Set Height 580
Set Shell "zsh"
Set Theme "Catppuccin Mocha"
Set TypingSpeed 30ms
Set Padding 24

Type "# QRCodeFYI Go API Client"
Enter
Sleep 0.5s

Type "curl -s 'https://qrcodefyi.com/api/search/?q=micro' | jq '.results[:3]'"
Enter
Sleep 3s

Type "curl -s 'https://qrcodefyi.com/api/version/1/' | jq '{version, modules}'"
Enter
Sleep 3s

Sleep 1s
