#!/usr/bin/env pwsh

param
(
	[Parameter(Mandatory=$true)][boolean]$isLocalBuild,
   [Parameter(Mandatory=$true)][string]$versionFile,
   [Parameter(Mandatory=$false)][string]$eventName
)

function ComputeVersionTimeComponent{
    [System.DateTime]$nowx = [System.DateTime]::Now
    [System.DateTime]$then = [System.DateTime]::Parse("2018-1-1")
    [System.TimeSpan]$diff = $nowx - $then
    $sec = [System.Math]::Floor($diff.TotalSeconds)
    $build = $sec.ToString("000000000")
    return $build
}

# read target product version from version.json file
$Product_Version_Object = (get-content -Path $versionFile -Raw)  | ConvertFrom-Json
$Product_Version = $Product_Version_Object.version


# Product version for ReleaseVersion (master or release branch): <SemVer>
# Otherwise: <SemVer>-<prefix>-<time>-<branchSha>
$branch_hash = git rev-parse --short HEAD

if ($isLocalBuild) {
    $branch = git rev-parse --abbrev-ref HEAD
}elseif ($eventName -like "pull*") {
    $branch = $env:GITHUB_HEAD_REF
}else {
    $branch = $env:GITHUB_REF
}

$prefix="";
if ($eventName -like "pull*") {
    $prefix = "prq-"
}elseif ($null -ne $branch) {
   if ($branch -like "*develop*") {
      $prefix = "dev-";
   } else {
      $prefix = "fea-";
   }
}

$buildTime = ComputeVersionTimeComponent

#determine if release product version
$IsReleaseVersion = $false;
if ((-not $isLocalBuild) -and ($eventName -eq "push") -and (($branch -eq "master") -or ($branch -like "release*"))) {
    $IsReleaseVersion = $true;
}

#condition on branch name
if (-not $IsReleaseVersion) {
    $Product_Version = "$Product_Version-$prefix$buildTime-$branch_hash"
}

$o = @{}
$o.ProductVersion = $Product_Version
$o.Branch = $branch
$o.IsReleaseVersion = $IsReleaseVersion


return $o