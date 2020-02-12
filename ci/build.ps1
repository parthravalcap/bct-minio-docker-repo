#!/usr/bin/env pwsh

param
(
	[Parameter(Mandatory=$true)][string]$version,
	[Parameter(Mandatory=$true)][string]$branch,
	[Parameter(Mandatory=$true)][boolean]$isReleaseVersion
)

# Assigning Global Variables
$buildName="minio-build"

$dockerRepo="tbctdevops-docker-local-poc.jfrog.io"
$dockerImage_collector="minio/minio"

# Adding remove the out folder step to resolve an error when consecutively publishing images
# docker build -t $buildName -t "$dockerRepo/${dockerImage_cassandra}:$version" -t "$dockerRepo/${dockerImage_cassandra}:latest" -f ./src/Dockerfile .
docker build -t $buildName -t "$dockerRepo/${dockerImage_collector}:$version" -f ./src/collector/Dockerfile -t "$dockerRepo/${dockerImage_collector}:$version" -f ./src/collector/Dockerfile .