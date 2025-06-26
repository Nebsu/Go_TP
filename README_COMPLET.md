# GoLog Analyzer

Un outil en ligne de commande (CLI) avancé en Go pour l'analyse distribuée de fichiers de logs. Développé avec une approche TDD (Test-Driven Development).

## 👥 Équipe de développement

- **[Theau Nicolas]**
- **[Zhou Sébastien]**
- **[Rozier Yanis]**

## 📋 Fonctionnalités

### ✅ Fonctionnalités principales
- 🔍 **Analyse concurrente** de multiples fichiers de logs via goroutines
- 📊 **Export JSON** des résultats d'analyse  
- 🛡️ **Gestion avancée des erreurs** avec types d'erreurs personnalisées
- 🎯 **Interface CLI** intuitive avec Cobra
- 📁 **Création automatique** des dossiers d'export
- 🗓️ **Horodatage automatique** des fichiers de rapport

### 🎁 Fonctionnalités bonus implémentées
- ➕ **Commande `add-log`** pour ajouter des configurations
- 🔍 **Filtrage par statut** (`--status OK|FAILED`)
- 📅 **Noms de fichiers horodatés** (format AAMMJJ)
- 📂 **Création automatique des répertoires** d'export

## 🏗️ Architecture

Le projet suit une architecture modulaire claire :

```
cmd/                    # Commandes CLI
├── root.go            # Commande racine
├── analyze.go         # Commande d'analyse principale  
└── add_log.go         # Commande d'ajout de configuration

internal/              # Packages internes
├── config/            # Gestion des configurations JSON
├── analyzer/          # Logique d'analyse et erreurs personnalisées
└── reporter/          # Export des rapports JSON
```

## 🚀 Installation et utilisation

### Prérequis
- Go 1.24+
- Module Cobra installé

### Installation
```bash
git clone <votre-repo>
cd Go_TP
go mod download
go build -o loganalyzer main.go
```

### Utilisation

#### 1. Commande `analyze` - Analyse de logs

```bash
# Analyse basique avec configuration JSON
./loganalyzer analyze -c config.json

# Analyse avec export vers fichier de rapport
./loganalyzer analyze -c config.json -o rapport.json

# Analyse avec filtrage par statut
./loganalyzer analyze -c config.json --status FAILED

# Les trois options combinées
./loganalyzer analyze -c config.json -o rapports/mon_rapport.json --status OK
```

#### 2. Commande `add-log` - Ajout de configuration

```bash
./loganalyzer add-log --id "nouveau-log" --path "/var/log/app.log" --type "application" --file config.json
```

## 📄 Format des fichiers

### Configuration JSON (`config.json`)
```json
[
  {
    "id": "web-server-1",
    "path": "/var/log/nginx/access.log", 
    "type": "nginx-access"
  },
  {
    "id": "app-backend-2",
    "path": "/var/log/my_app/errors.log",
    "type": "custom-app"
  }
]
```

### Rapport de sortie (`report.json`)
```json
[
  {
    "log_id": "web-server-1",
    "file_path": "/var/log/nginx/access.log",
    "status": "OK", 
    "message": "Analyse terminée avec succès.",
    "error_details": ""
  },
  {
    "log_id": "invalid-path",
    "file_path": "/non/existent/log.log", 
    "status": "FAILED",
    "message": "Fichier introuvable.",
    "error_details": "open /non/existent/log.log: no such file or directory"
  }
]
```

## 🔧 Gestion des erreurs

### Erreurs personnalisées implémentées :

```go
// Fichier introuvable ou inaccessible
type FileNotFoundError struct { Path string }

// Erreur de parsing (simulée à 10%)  
type ParseError struct { Path string }
```

### Utilisation avec errors.Is et errors.As :
```go
result, err := analyzer.AnalyzeLog(id, path, typ)
if err != nil {
    var fnfErr *FileNotFoundError
    var parseErr *ParseError
    
    if errors.As(err, &fnfErr) {
        // Gestion spécifique fichier non trouvé
    } else if errors.As(err, &parseErr) {
        // Gestion spécifique erreur parsing
    }
}
```

## 🌟 Exemples d'utilisation

### Exemple complet d'analyse
```bash
# 1. Créer une configuration
echo '[{"id":"test","path":"./test.log","type":"app"}]' > config.json

# 2. Créer un fichier de log de test  
echo "log content" > test.log

# 3. Analyser
./loganalyzer analyze -c config.json -o rapport_250613.json

# 4. Vérifier le résultat
cat rapport_250613.json
```

### Exemple avec filtrage
```bash
# Voir seulement les échecs
./loganalyzer analyze -c config.json --status FAILED

# Exporter seulement les succès
./loganalyzer analyze -c config.json -o succès.json --status OK
```