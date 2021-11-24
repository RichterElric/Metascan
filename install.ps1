# Create install directory

# Check if python exist
Write-Host "Make sure python 3 and pip is installed on your system"
# Remove error print
$ErrorActionPreference = 'SilentlyContinue'

if (-not (Test-Path "$PSScriptRoot/bin"))
{
    New-Item -ItemType Directory -Path "$PSScriptRoot/bin" | Out-Null
}
# Install kics
if (-not [System.IO.File]::Exists("$PSScriptRoot/bin/kics/kics.exe"))
{
    Invoke-WebRequest "https://github.com/Checkmarx/kics/releases/download/v1.4.7/kics_1.4.7_windows_x64.zip" -O "$PSScriptRoot/bin/kics.zip"
    Expand-Archive "$PSScriptRoot/bin/kics.zip" -DestinationPath "$PSScriptRoot/bin/kics/"
    Remove-Item -Path "$PSScriptRoot/bin/kics.zip"
}
# Install git-secret
if (-not (Test-Path "$PSScriptRoot/bin/gitsecret"))
{
    Invoke-WebRequest "https://github.com/awslabs/git-secrets/archive/refs/heads/master.zip" -O "$PSScriptRoot/bin/gitsecret.zip"
    Expand-Archive "$PSScriptRoot/bin/gitsecret.zip" -DestinationPath "$PSScriptRoot/bin/gitsecret"
    Get-ChildItem -Path "$PSScriptRoot/bin/gitsecret/git-secrets-master" -Recurse -File | Move-Item -Destination "$PSScriptRoot/bin/gitsecret/"
    Remove-Item -Path "$PSScriptRoot/bin/gitsecret.zip"
    & "$PSScriptRoot/bin/gitsecret/install.ps1" | Out-Null
}
# Install Keyfinder
if (-not -(Test-Path "$PSScriptRoot/bin/keyfinder"))
{
    Invoke-WebRequest "https://github.com/CERTCC/keyfinder/archive/refs/heads/master.zip" -O "$PSScriptRoot/bin/keyfinder.zip"
    Expand-Archive "$PSScriptRoot/bin/keyfinder.zip" -DestinationPath "$PSScriptRoot/bin/keyfinder"
    Get-ChildItem -Path "$PSScriptRoot/bin/keyfinder/keyfinder-master" -Recurse -File | Move-Item -Destination "$PSScriptRoot/bin/keyfinder/"
    Remove-Item -Path "$PSScriptRoot/bin/keyfinder.zip"
    Remove-Item -Path "$PSScriptRoot/bin/keyfinder/keyfinder-master"
    Invoke-Item (start powershell "pip install install androguard python-magic-bin PyOpenSSL") | Out-Null
}