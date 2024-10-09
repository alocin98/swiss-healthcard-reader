# Swiss healthcard reader

Dies ist eine minimalistische Implementation zum Auslesen
der Daten einer Schweizer Versichertenkarte (Krankenversicherung)

Die Implementierung folgt der Implementierungsanleitung, welche
beim BAG verfügbar ist:
[bag.admin.ch](https://www.bag.admin.ch/bag/de/home/versicherungen/krankenversicherung/krankenversicherung-versicherte-mit-wohnsitz-in-der-schweiz/versichertenkarte.html)

## Anleitung

### Voraussetzung (Entwicklung)

- Go ist auf dem System installiert
  [https://go.dev/doc/install](https://go.dev/doc/install)

Für den Devserver wird air verwendet:
[https://github.com/air-verse/air](https://github.com/air-verse/air)

Um die Anwendung zu starten, musst du folgende Schritte ausführen:

1. Module installieren

```
go mod download
```

2. Dev server starten

```
air
```

2. Build erstellen (für windows)

```
go build -o ./build/server.exe .
```
