# Tür 2.0

Die bestehende elektronische Türschließanalge für den Krautspace soll überarbeitet werden da das momentane Setup an einigen Stellen nicht zufriedenstellend ist. Dies hier ist die Spec dazu.

Auf http://schuylertowne.com/blog/smart-locks gibt es eine Übersicht an kommerziellen Schliessystemen und deren Zustand.

## Allgemeines

Sowohl die Haustür als auch die Tür zu den Vereinsräumen (Wohnungstür) sollen mit der elektronischen Schließanlage auf- und abgeschlossen werden können. 

### Featurewünsche

Wenn der Raum besetzt ist, dann soll beim Klingeln an der Haustür automatisch eine akustische Nachricht "Raum ist offen" ausgegeben werden und dann der Türschnapper geöffnet werden. Wenn der Raum nicht besetzt ist, dann soll auch das akustisch mitgeteilt werden und es soll die möglichkeit bestehen per Klingelzeichen trotzdem das Öffnen der Haustür zu veranlassen. Priorität hat, dass die Haustür per Klingelzeichen geöffnet werden kann, dann dass sie geöffnet wird wenn der Raum besetzt ist und dann die akustische Rückmeldung. 

## Schließmethoden (für die Tür zum Raum)

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

### Selbstgelötets Widerstandsdevice

Es werden n (ca. 4) Widerstände auf eine kleine Platine mit Stecker gelötet. Der Stecker wird in eine Buchse an der Tür gesteckt, ein kleiner Computer misst die Widerstände, wenn diese einem hinterlegen Profil entsprechen wird geöffnet. Dies sollte eine sehr kostengünstige und einfache Variante sein. Man muss aber die Buchse an der Tür vandalitätssicher machen. Probleme können auftreten wenn dort eine Spannungsquelle angeschlossen wird.

### Schlüsselschalter

Ein Schalter, welcher mit einem Schließzylinder und Schlüssel geschaltet wird. Das ist eine sehr einfache Möglichkeit, muss aber physisch gut abgesichert werden. 

### Richtiger Schlüssel für das Schloss der Tür

Der Schlüssel aus Metall soll zumindest für die Wohnungstür nur noch als Rückfalllösung benutzt werden. 

## Benutzermanagement

Für die Zugangsdaten muss es eine Verwaltungsmöglichkeit geben auf die mehrere Vereinsmitglieder Zugang haben können. 

## Sonstige Anforderungen

### Stromverbrauch

Die Anlage sollte im Leerlauf möglichst wenig Strom (zusätzlich zu bereits vorhandener Hardware wie den Freifunkroutern) verbrauchen.

### Akkubetrieb

Es sollten Akkus verbaut werden, so dass die Anlage auch bis zu 24 Stunden ohne feste Stromversorgung funktioniert.

