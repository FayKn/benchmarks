# Benchmarks

Dit is een benchmark repo voor verschillende talen, ik houd het hier heel simpel:
- Haal een records op uit een steam vanity URL op uit een mySQL database
- Strip je vanity url hier uit
- doe een api call naar de steam api om het steamID op te halen en geef dit terug in het volgende format
- update de SteamUsers tabel met het steamID
- increment en decrement een counter in een mySQL database voor elke request
- Cache dit resultaat in een Redis database
```json
[
  "gabelogannewell",
  "knol_gamer",
  "olbmaphlee",
  "SandsVR",
  "SilverHazel",
  "valvesoftware"
]
```

## De benchmark
De benchmark wordt gedraaid met k6 volgens script.js dus 60s van 200 requests per seconde.
Ik doe hier maar 200 omdat Laravel bij meer requests begon te droppen vanaf het begin.

De graceful stop staat op 30s zodat de requests die nog bezig zijn ook nog afgehandeld kunnen worden, dit zal je terug zien in de runtime kolom; requests langer dan dat worden als failed beschouwd.

De docker image voor Redis en Mysql worden elke run volume en al verwijderd en opnieuw aangemaakt.

## Resultaten
CPU: Ryzen 9 5900x | RAM: 32GB 3200Mhz | SSD:Lexar NM790

OS: Windows 10 Pro 22H2 build 19045.4412 | Docker: 26.1.1 | WSL: 2.0.14.0

Base steam api time: ~165ms | Internet speed: 1000/1000mbps (up/down) | 50 ping to steam api average: 3ms

| Taal    | Requests/s | Min request duration | Avr. request duration | Max request duration | Requests failed | Requests completed | runtime |
|---------|------------|----------------------|-----------------------|----------------------|-----------------|--------------------|---------|
| Go      | 188.9      | 516.2µs              | 46.56ms               | 2.46s                | 0 (0,00%)       | 11520              | 61s     |
| Elixir  | 188,1      | 999.4µs              | 53.03ms               | 4.18s                | 11 (0,09%)      | 11472              | 61s     |
| Laravel | 5.5        | 9.3s                 | 28.85s                | 43.5s                | 7 (1,3%)        | 499                | 90s     |