#!/usr/bin/env pwsh

$versionInfo = ./ci/calculate-buildInfo.ps1 $true version.json
Write-Host "ProductVersion: $($versionInfo.ProductVersion)"
Write-Host "Branch: $($versionInfo.Branch)"
Write-Host "Release product: $($versionInfo.IsReleaseVersion)"

# Docker login to the repo registry
./ci/login.ps1

# Build and tag the Docker Image
./ci/build.ps1 $versionInfo.ProductVersion $versionInfo.Branch $versionInfo.IsReleaseVersion

# Push Docker Image to Artifactory
./ci/publish.ps1 $versionInfo.ProductVersion $versionInfo.Branch $versionInfo.IsReleaseVersion

# Clean Up Artifactory and Docker

./ci/cleanup.ps1 $versionInfo.ProductVersion $versionInfo.Branch $versionInfo.IsReleaseVersion