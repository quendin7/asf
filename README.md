# asf — ASol System Fetcher (lekko odklejony)

asf to mały, niepoważny fetcher systemowy z opcjami, które robią mniej więcej to, czego się spodziewasz — albo kompletnie nic sensownego.  
Naprawdę. Uruchom i zobacz.

## Funkcje (wybrane, bo reszta to magia)
- Wyświetla użytkownika i hosta, info o systemie, jądro, pakiety.
- Hardware: CPU, GPU, RAM, swap, bateria, dyski, sieć.
- Desktop: DE/WM, motywy GTK/ikon, fonty, rozdzielczość.
- Dodatki: uptime, muzyka, shell, mały parasol do TTS (tak, parasol).
- Flagi: włącz wszystko (-a), tylko logo (-L), wyłącz logo (-l), parasol :3 (-tts), JSON (--json), domyślny config (-c) i kilka innych.

## Przykładowe użycie
- Normalnie:
```
./asf
```
- Wszystko włączone:
```
./asf -a
```
- Default config:
```
./asf -c
```
- WIELKI POTĘŻNY ASRIEL:
```
./asf -L
```

## Dziwactwa i reguły życiowe projektu
- Jeśli włączysz `-L` i `-tts`, program ci powie to co Oli Sykes w 2015.
- Zrobić coś głupiego? Prawdopodobnie już to zrobi.  
- Masz bug? Przyjmij go jako nową funkcję i stwórz issue, jeżeli masz ochotę.
## Instalacja
no se ogólnie skompiluj
```
go build asfetch.go
```
i wklej do np ```/usr/bin/asf```
lub jak jesteś pro koks i używasz archa btw, to użyj PKGBUILD
```makepkg -si```
## Konfiguracja
Plik konfigu zapisuje się w katalogu użytkownika (domyślnie `~/.config/asf/config.json`).  
Masz flagę `-c`, żeby wymusić użycie domyślnego configu.
