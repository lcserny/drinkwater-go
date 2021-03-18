@echo off

echo Building drinkwater.exe

go build -ldflags="-H windowsgui" -o "out/drinkwater.exe" main.go

echo Copying required resources
copy drinkwater.exe.manifest out
copy glass_original.ico out

echo Done! Check "out" folder
