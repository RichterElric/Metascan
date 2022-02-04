# Metascan
Outil de scan de projet avec analyse des dépendances, des fichiers de configuration
et analyse statique de code.

# Build le projet

- Deux scripts sont présents dans le projet : install.ps1 & install.sh pour l'installation des dépendances.
- Pour une installation plus pratique : un dockerfile est à disposition dans le dossier racine Metascan
```bash
docker build . -t metascan
```
Sur Windows :
```bash
docker run --rm -v C:\chemin\absolu\vers\le\dossier:/opt/scan metascan -dc=false
```
Sur Linux :
```bash
docker run --rm -v /chemin/absolu/vers/le/dossier:/opt/scan metascan -dc=false
```

# Outils pour les scans

(liste en cours de développement)

| File type | Scanner |
| :--- | :---: |
| Docker, docker compose, kubernetes ... | [Kics](https://github.com/Checkmarx/kics)|
| PMD : analyse statique de code Java & XML | [PMD](https://pmd.github.io/)| 
| pyLint : analyse statique de code Python| [Pylint](https://pylint.org/)|
| dotenv-linter : analyse .env| [dotenv-linter](https://github.com/dotenv-linter/dotenv-linter)|
| Passwords/keys | [git-secret](https://github.com/awslabs/git-secrets)|
| C/C++ analyse statique de code | [cppcheck](https://cppcheck.sourceforge.io/)|
| Dependency checker : analyse de dépendances|[dependency checker](https://jeremylong.github.io/DependencyCheck/)|

# Notes sur les scanners 

| Scanner | Note |
| :--- | :--- |
|cppcheck| A utiliser à la racine d'un projet Cpp/C |
|pyLint| A utiliser à la racine d'un projet python contenant un \_\_init\_\_.py |
