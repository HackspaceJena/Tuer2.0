# Tür 2.0

Die bestehende elektronische Türschließanalge für den Krautspace soll überarbeitet werden da das momentane Setup an einigen Stellen nicht zufriedenstellend ist. Dies hier ist die Spec dazu. 

## Allgemeines

Sowohl die Haustür als auch die Tür zu den Vereinsräumen (Wohnungstür) sollen mit der elektronischen Schließanlage auf- und abgeschlossen werden können. 

## Schließmethoden

Es ist gewünscht, dass preislich günstige Tokens zum Schließen benutzt werden können die kein weiteres Gerät auf der Seite des Benutzers erfordern. 
Es ist nicht möglich, dass Hardware an der Hastür angebracht wird. 

### NFC-Tokens

### RFID-Tokens

Es sollen möglichst beliebige RFID-Tokens benutzbar sein um bei den Benutzern bereits vorhandene Tokens integrieren zu können. 
RFID ist leider nur bedingt eine Lösung für besonders vorsichtige Vereinsmitglieder. 

### BT-Tokens

Sowas: http://www.imazeconnect.com/products/pocket-mate-bluetooth-4-0-security-tag

Wenn man mit den Token in der Nähe ist, kann man per Knopf die Tür aufmachen.

### Pin-Pad

Eventuell mit [RFC 6238](https://tools.ietf.org/html/rfc6238)? (Das ist im Prinzip der [Google Authenticator](http://de.wikipedia.org/wiki/Google_Authenticator))

### Tastatur

### Karte mit Magnetstreifen zum durchziehen

### WLAN-Interface

Das WLAN-Interface sollte über Freifunk erreichbar sein, allerdings nicht im gesamten Freifunknetz, sondern nur "in der Nähe" der Vereinsräume um versehentliches aufschließen aus großer räumlicher Entfernung zu verhindern. 

Eventuell mit [RFC 6238](https://tools.ietf.org/html/rfc6238)? (Das ist im Prinzip der [Google Authenticator](http://de.wikipedia.org/wiki/Google_Authenticator))

### Rotary combination lock

Nachbau eines Rotary combination locks wie an Tresoren, beispielsweise mit einem Potentiometer dessen Widerstand mehrfach nacheinander ausgelesen wird.

### Richtiger Schlüssel 

Der Schlüssel aus Metall soll zumindest für die Wohnungstür nur noch als Rückfalllösung benutzt werden. 

# Benutzermanagement

Für die Zugangsdaten muss es eine Verwaltungsmöglichkeit geben auf die mehrere Vereinsmitglieder Zugang haben können. 

# Sonstige Anforderungen

## Stromverbrauch

Die Anlage sollte im Leerlauf möglichst wenig Strom (zusätzlich zu bereits vorhandener Hardware wie den Freifunkroutern) verbrauchen.

## Akkubetrieb

Es sollten Akkus verbaut werden, so dass die Anlage auch bis zu 24 Stunden ohne feste Stromversorgung funktioniert.

