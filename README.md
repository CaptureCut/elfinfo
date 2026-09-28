2026-09-28

Goal:
Создать аналог readelf на Go.

Done:
- инициализирован go module
- реализован parser.go
- реализован вывод ELF Header
- реализован вывод секций
- протестировано на /bin/ls и /bin/bash

Next:
- ImportedLibraries()
- JSON output
- SHA256
