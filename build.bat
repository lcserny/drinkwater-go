@echo off

echo Creating registry entry

powershell -Command "& {set-location 'HKCU:\SOFTWARE\Microsoft\Windows\CurrentVersion\Notifications\Settings'; new-item -Force DrinkwaterGo; set-location DrinkwaterGo; new-itemproperty . -Name ShowInActionCenter -Value 1 -Type DWORD;}"

echo Building drinkwater.exe

go build -ldflags="-H windowsgui" -o drinkwater.exe main.go

echo Done!
