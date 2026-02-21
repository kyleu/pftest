# VHS for installation of Project Forge
#
# https://github.com/charmbracelet/vhs

Output ../assets/homebrew-install.gif
# Output ../assets/homebrew-install.mp4
# Output ../assets/homebrew-install.webm
# Output ../assets/homebrew-install/frames/

Set TypingSpeed 50ms
Set FontSize 18
Set Width 1000
Set Height 230
Set Padding 20
Set Theme "Atom"

Type "brew install kyleu/kyleu/pftest"

Sleep 500ms

Enter

Wait

Sleep 500ms

Type "pftest version"

Sleep 500ms

Enter

Sleep 1s

Type "Test Project"

Sleep 1s

Type " - A Test application, built with Project Forge"

Sleep 3s
