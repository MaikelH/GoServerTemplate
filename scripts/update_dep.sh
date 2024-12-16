#!/usr/bin/env bash
set -o errexit
set -o pipefail
set -o nounset
if [[ "${TRACE:-0}" == "1" ]]; then set -o xtrace; fi

# This script updates all Go module dependencies in a project to their latest versions, displays a list of updated 
# dependencies with old and new versions, and reports the script execution time.

backup_go_mod() {
  if [ ! -f go.mod ]; then
    echo "Error: go.mod file not found in the current directory. Please run the script in the correct directory."
    exit 1
  fi

  echo "Backing up go.mod..."
  cp go.mod go.mod.bak
  echo "Backup completed."
}

fetch_dependencies() {
  echo "Fetching $1 dependencies..."
  go list -m all 2>/dev/null
  echo "Fetched $1 dependencies."
}

update_dependencies() {
  echo "Updating all dependencies to the latest version..."
  go get -u ./...
  echo "Dependencies updated."
}

compare_dependencies() {
  echo "Comparing dependencies and listing updates:"
  echo "-------------------------------------------"
  diff -y --suppress-common-lines \
    <(echo "$1" | awk '{print $1, $2}') \
    <(echo "$2" | awk '{print $1, $2}') \
    | awk '{printf "%-50s %-15s -> %s\n", $1, $2, $5}'
}

# Check if commands needed for the script are present.
check_commands() {
  local commands=("go" "diff" "awk")

  for cmd in "${commands[@]}"; do
    if ! command -v "$cmd" > /dev/null; then
      echo "Error: '$cmd' command not found. Please make sure it is installed and available in your PATH."
      exit 1
    fi
  done
}

main() {
  local start_time=$SECONDS

  backup_go_mod

  local current_dependencies
  current_dependencies=$(fetch_dependencies "current")
  
  update_dependencies
  
  local updated_dependencies
  updated_dependencies=$(fetch_dependencies "updated")

  compare_dependencies "$current_dependencies" "$updated_dependencies"

  local elapsed_time=$((SECONDS - start_time))
  echo "-------------------------------------------"
  echo "Script execution time: $elapsed_time seconds"
}

check_commands
main
