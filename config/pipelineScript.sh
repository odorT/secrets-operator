#!/usr/bin/env bash

## step 1: get base findings and config
## step 2: gitleaks detect with basepath
## step 3: post new findings

SECRETS_OPERATOR_URL=http://secrets-operator-dev.apps.test.ocp.ibar.az
PREFIX=/api/v1/findings

# step 1
echo Fetching baseline findings and gitleaks config.toml ...
curl "${SECRETS_OPERATOR_URL}${PREFIX}/${CI_PROJECT_ID}" | jq '.findings' > base-findings.json
curl "${SECRETS_OPERATOR_URL}/api/v1/config.toml" > config.toml

# step 2
echo Running gitleaks ...
gitleaks detect --config config.toml --baseline-path base-findings.json --source . --report-path findings.json

#step3
echo Publishing findings report to secrets operator ...
curl --location --request POST "${SECRETS_OPERATOR_URL}${PREFIX}/upload" \
--header "pipelineId: ${CI_PIPELINE_ID}" \
--header "repoName: ${CI_PROJECT_NAME}" \
--header "repoId: ${CI_PROJECT_ID}" \
--header "repoURL: ${CI_PROJECT_URL}" \
--header "commitAuthor: ${CI_COMMIT_AUTHOR}" \
--header "commitSHA: ${CI_COMMIT_SHA}" \
--header "timestamp: $(date +%s)" \
--header "notify: true" \
--header "Content-Type: application/json" \
-d @findings.json

echo -e "\nDone."
echo Check details here: $SECRETS_OPERATOR_URL