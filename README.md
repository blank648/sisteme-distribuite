# Sisteme Distribuite

Rezolvări pentru laboratoarele de **Sisteme Distribuite** implementate în Go (Golang).

## Structura repository-ului

- **lab1/**: Noțiuni introductive de Go (sintaxă, array-uri, slice-uri, map-uri, I/O fișiere)
  - `ex1.go`: Verificare număr prim
  - `ex2.go`: Sortare array crescător (Bubble Sort)
  - `ex3.go`: Filtrare numere pare dintr-un șir
  - `ex4.go`: Eliminare element după index din slice `float32`
  - `ex5.go`: Inversare in-place a unui slice de string-uri
  - `ex6.go`: Sortare descrescătoare a unui map după valori (înălțimi)
  - `ex7.go`: Citire cuvinte din fișier și calculul frecvenței de apariție
  - `main.go`: Meniu consolă pentru rularea exercițiilor 1-7
- **lab2/**: Canale și concurență în Go (Goroutines & Channels)

## Rulare

Pentru a rula laboratorul 1:
```bash
cd lab1
go run .
```

Pentru a rula laboratorul 2:
```bash
cd lab2
go run main.go
```
