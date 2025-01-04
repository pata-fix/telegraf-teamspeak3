# Basis-Image für AMD64, aber wir bauen für ARM64
FROM golang:1.21-alpine AS builder

# Installieren von Build-Essentials und Cross-Compilation-Tools
RUN apk add --no-cache git gcc musl-dev

# Arbeitsverzeichnis erstellen
WORKDIR /build

# Kopieren der Go-Dateien
COPY teamspeak.go .

# Go Modules initialisieren und Abhängigkeiten installieren
RUN go mod init telegraf-teamspeak3 && \
    go get github.com/thannaske/go-ts3 && \
    go mod tidy

# Cross-Kompilierung für ARM64
RUN GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -o teamspeak-monitor teamspeak.go

# Zweite Stage: Finales Image
FROM arm64v8/telegraf:1.28

# Arbeitsverzeichnis
WORKDIR /etc/telegraf

# Kopieren des kompilierten TeamSpeak-Monitors
COPY --from=builder /build/teamspeak-monitor /usr/local/bin/
RUN chmod +x /usr/local/bin/teamspeak-monitor

# Standard-Konfigurationsdatei erstellen
COPY telegraf.conf /etc/telegraf/telegraf.conf

# Ports für Telegraf
EXPOSE 8125/udp 8092/udp 8094

# Startbefehl
ENTRYPOINT ["/entrypoint.sh"]
CMD ["telegraf"]