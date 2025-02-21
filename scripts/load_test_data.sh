#!/usr/bin/env bash
set -euo pipefail
# explicitly find and set the script and repository directories
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
# include helper functions. load before trace to avoid debug spam
source "$SCRIPT_DIR/util.sh"
# Set for script debugging
if [[ "${TRACE:-0}" == "1" ]]; then set -o xtrace; fi

# if not set, use default values
_DB="${DB_NAME:-retrolink}"
_ROOT_USER="${DB_ROOT_USER:-postgres}"
_ROOT_PASSWORD="${DB_ROOT_PASSWORD:-testpass}"
_HOST="${DB_HOST:-127.0.0.1}"
_PORT="${DB_PORT:-5432}"

REPO_ROOT_DIR=$(cd "${SCRIPT_DIR}"/.. && pwd)

# save the current directory and change to the repository root directory
pushd "$REPO_ROOT_DIR" > /dev/null

printf "Loading Test Data\n"
# load the test data
export PGPASSWORD="${_ROOT_PASSWORD}"
#mysql --defaults-extra-file=<(printf "[client]\nuser = %s\npassword = %s" "${_ROOT_USER}" "${_ROOT_PASSWORD}") --host="${_HOST}" --port="${_PORT}" "${_DB}" < ./support/sql/company.test_data.sql
psql -h "${_HOST}" -p "${_PORT}" -U "${_ROOT_USER}" -d "${_DB}" -f ./support/test_data.sql

printf "Successfully loaded Test Data into Database ${GOOD}%s${FG_CLEAR} on ${GOOD}%s:%s${FG_CLEAR}\n" "${_DB}" "${_HOST}" "${_PORT}"

# restore the caller's directory
popd > /dev/null
