### JOGO DA ADIVINHAÇÃO EM GO

Este é um jogo simples de adivinhação em Go. O jogo gera um número aleatório entre 1 e 100 e o jogador deve adivinhar qual é esse número. O jogo dá dicas de "maior" ou "menor" para ajudar o jogador a adivinhar o número.

Para jogar, basta executar o arquivo `main.go` com o comando `go run main.go`.

### Exemplo de uso

```
$ go run main.go
Tente adivinhar o número que estou pensando entre 1 e 100...
Digite um número: 50
O número que estou pensando é menor
Digite um número: 25
O número que estou pensando é maior
Digite um número: 37
O número que estou pensando é menor
Digite um número: 31
O número que estou pensando é maior
Digite um número: 34
Parabéns! Você acertou o número que eu estava pensando em 34 tentativas.
```

