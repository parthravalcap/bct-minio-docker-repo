#!/usr/bin/env pwsh

param
(
	[Parameter(Mandatory=$true)][string]$username,
	[Parameter(Mandatory=$true)][string]$apikey
)

$dockerRepo="tbctdevops-docker-local-poc.jfrog.io"
	  
# Docker login to the repo registry
docker login $dockerRepo --username $username --password $apikey