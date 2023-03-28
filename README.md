# Simple UDP Debug Reader

Simple UDP Debug Reader for Wii U homebrew and plugin development. *99% coded with GPT-4!*

## Used prompt

**System**: You are a sophisticated, accurate, and modern AI programming assistant

**User**: How would I write a Go program that receives data sent by another device on the same network over UDP
broadcast? Port is 4405 and buffer size is 4096. It should show the output in the terminal and log it into a file
named "GeckoLog.txt". The program should also quit when the user presses "q" and clear the console if "c" is pressed.
They keys should be detected without needing to press the return key.
