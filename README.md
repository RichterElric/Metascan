# Metascan
Outil de scan de code

# Build le projet

- deux scripts sont présent dans le projet : install.ps1 & install.sh pour l'installation des dépendances.
- Pour une installation plus pratique : un dockerfile est à disposition : dans le dossier racine Metascan
```docker
docker build . -t metascan
```
puis :
```docker
docker run metascan "args"
```


# Documentation Go

- https://dl.hiva-network.com/Library/security/Black-Hat-Go_Go-Programming-For-Hackers-and-Pentesters.pdf
- https://gobyexample.com/
- https://gowebexamples.com/

# Outils pour les scans

(liste en cours de développement)

| File type | scanner |
| :--- | :---: |
| Docker, docker compose, kubernetes ... | [Kics](https://github.com/Checkmarx/kics)|
| certificates | [keyfinder](https://github.com/CERTCC/keyfinder)|
| password / keys | [git secret](https://github.com/awslabs/git-secrets)|

# Repo pour le dev

Repo avec pleins de types de programmes : https://github.com/TheRenegadeCoder/sample-programs
