# Secrets Operator
This program is designed for analyzing, reporting and collecting statistics
about hard coded secrets in source code repositories(git). Gitleaks used for
report generation engine 


---
## Tools
- **Programming Language**: Golang (go version >go1.19.3)  
- **Web Framework**: gin
- **Database**: MongoDB  


## Client side flow:
1. pipeline will fetch config.toml(configuration file for gitleaks) and 
base-findings.json(already found secrets) from api
2. it will check for new secrets. 
3. it will post findings to secrets-operator, then program will store \
them in database and send notification.


## API Matrix
