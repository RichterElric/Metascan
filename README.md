# Metascan
Outil de scan de projet avec analyse des dépendances, des fichiers de configuration
et analyse statique de code.

# Build le projet

- deux scripts sont présent dans le projet : install.ps1 & install.sh pour l'installation des dépendances.
- Pour une installation plus pratique : un dockerfile est à disposition : dans le dossier racine Metascan
```bash
docker build . -t metascan
```
puis :
sur windows
```bash
docker run --rm -v C:\chemin\absolu\vers\le\dossier:/opt/scan metascan -dc=false
```
sur linux :
```bash
docker run --rm -v /chemin/absolu/vers/le/dossier:/opt/scan metascan -dc=false
```

# Outils pour les scans

(liste en cours de développement)

| File type | scanner |
| :--- | :---: |
| Docker, docker compose, kubernetes ... | [Kics](https://github.com/Checkmarx/kics)|
| PMD : analyse static de code java & XML | [PMD](https://pmd.github.io/)| 
|pyLint : analyse static de code python| [Pylint](https://pylint.org/)|
|dotenv linter : analyse .env| [dotenv-linter](https://github.com/dotenv-linter/dotenv-linter)|
| password / keys | [git-secret](https://github.com/awslabs/git-secrets)|
|c / cpp analyse static de code | [cppcheck](https://cppcheck.sourceforge.io/)|
|Dependency checker :  analyse de dépendances|[dependency checker](https://jeremylong.github.io/DependencyCheck/)|

# Note sur les scanners 

| Scanner | Note |
| :--- | :--- |
|cppcheck| A utiliser à la racine d'un projet Cpp/C |
|pyLint| A utiliser à la racine d'un projet python contenant un \_\_init\_\_.py |
