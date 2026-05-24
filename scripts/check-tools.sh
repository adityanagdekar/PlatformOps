#!/bin/sh
set -eu

missing=0

check_tool() {
  if ! command -v "$1" >/dev/null 2>&1; then
    printf 'missing: %s\n' "$1"
    missing=1
    return
  fi

  printf 'found: %s (%s)\n' "$1" "$(command -v "$1")"
}

check_tool docker
check_tool kubectl
check_tool kind
check_tool go

if [ "$missing" -ne 0 ]; then
  printf '\nInstall the missing tools before running the full local workflow.\n'
  exit 1
fi

printf '\nAll required Phase 1 tools are available.\n'

