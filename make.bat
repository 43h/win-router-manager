del *.exe
del rsrc.syso

tools\rsrc -manifest main.manifest -o rsrc.syso

go build -ldflags "-w -s -H windowsgui" -o wrm.exe
