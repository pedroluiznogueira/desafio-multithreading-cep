# Multithreading CEP

Busca um CEP em duas APIs simultaneamente (BrasilAPI e ViaCEP) e usa a primeira resposta que
chegar, descartando a outra. Timeout de 1 segundo — se nenhuma responder a tempo, exibe erro.

## Rodando

```
go run main.go 01001000
```

O CEP é opcional (`01153000` é usado se nenhum for informado).

## Saída esperada

```
Resposta mais rápida: BrasilAPI
{"cep":"01001000","state":"SP","city":"São Paulo", ...}
```

Ou, em caso de timeout:

```
Erro: timeout - nenhuma API respondeu em 1 segundo
```
