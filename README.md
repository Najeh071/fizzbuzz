FIZZ-BUZZ REST API version `v0.1.0`

API REST en go qui génère une séquence Fizz-buzz

> **Note :** Projet réalisé par **Najeh Abassi** dans le cadre du test technique Backend.

Prérequis
- Go 1.22 ou suppérieur
- Git

Lancement de project:

- git clone https://github.com/Najeh071/fizzbuzz.git
- cd fizzbuzz
- go mod tidy
- go run ./cmd/server # le serveur démarre sur http://localhost:8081 , tu peut changer le port dans main.go
- pour exécuter les testes: go test ./... -v

Exemples d'utilisation:

pour générer un Fizz-Buzz:
- curl "http://localhost:8081/api/v1/fizzbuzz?int1=3&int2=5&limit=15&str1=fizz&str2=buzz"
pour consulter les statestiques:
- curl http://localhost:8081/api/v1/stats

## Auteur 
* **Najeh Abassi** - *Lead Developer*