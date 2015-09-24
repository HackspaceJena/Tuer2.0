# Tür Server

Öffnet die Tür zum Hackerspace. Dieser Server bietet eine Rest-API für die verschiedenen Funktionen:

 - `/door/buzzer/open`: Öffnet die Tür unten.
 - `/door/lock/lock`: Schliest die Tür zu.
 - `/door/lock/unlock`: Schliest die Tür auf
 - `/door/lock/open`: Öffnet die Tür

## Bauprozess

 - Go installieren.
 - GOPATH auf "$REPO/src/server" setzen.
 - go get server
 - go install server
 - bin/server ausführen


## TODO

 - Authentifizierung einbauen. 
 - Ein Angular/Whatever-Interface schreiben.
