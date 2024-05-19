# Benchmarks

Dit is een benchmark repo voor verschillende talen, ik houd het hier heel simpel:
- Haal een records op uit een steam vanity URL op uit een mySQL database
- Strip je vanity url hier uit
- doe een api call naar de steam api om het steamID op te halen en geef dit terug in het volgende format
- increment en decrement een counter in een mySQL database voor elke request
- Cache dit resultaat in een Redis database
```json
{
  [
    {
      "steamID": "76561198000000000",
      "name": "Naam"
    },
    {
      "steamID": "76561198000000001",
      "name": "Naam"
    }
  ]
}
```

## De benchmark
De benchmark wordt gedraaid met k6 volgens script.js dus 60s van 200 requests per seconde.
Ik doe hier maar 200 omdat Laravel bij meer requests begon te droppen vanaf het begin.

De graceful stop staat op 30s zodat de requests die nog bezig zijn ook nog afgehandeld kunnen worden, dit zal je terug zien in de runtime kolom; requests langer dan dat worden als failed beschouwd.

De docker image voor Redis en Mysql worden elke run volume en al verwijderd en opnieuw aangemaakt.

## Resultaten
CPU: Ryzen 9 5900x
RAM: 32GB 3200Mhz
SSD:Lexar NM790
OS: Windows 10 Pro 22H2 build 19045.4412
Base steam api time: ~165ms
Internet speed: 1000/1000mbps (up/down)
50 ping to steam api average: 2ms

| Taal    | Requests/s | Min response tijd | Gem. responsetijd | Requests failed | Requests handled | runtime |
|---------|------------|-------------------|-------------------|-----------------|------------------|---------|
| Go      |            |                   |                   |                 |                  |         |
| Elixir  |            |                   |                   |                 |                  |         |
| Laravel | 5.9        | 3.01s             | 27.75s            | 7 (1,3%)        | 533              | 90s     |