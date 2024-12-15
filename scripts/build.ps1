function DockerBuild {
    param(
        [string]$imageName,
        [string]$dockerFile,
        [string]$context
    )

    $changelogPath = $dockerFile.Replace("Dockerfile", "changelog.md")
    $version = Get-Version -changelogPath $changelogPath

    $imageName = "$($imageName):$version"
    
    Write-Host "Building $imageName"

    docker build -t $imageName -f $dockerFile $context --progress=plain
}

function Get-Version {
    param(
        [string]$changelogPath
    )

    $changelog = Get-Content $changelogPath
    $version = $changelog | Select-String -Pattern "#" | Select-Object -First 1
    $version = $version -replace "#", ""
    $version = $version.Trim()
    
    return $version
}

DockerBuild -imageName "scraper-service" -dockerFile "./cmd/scraper/Dockerfile" -context "."
DockerBuild -imageName "parser-service" -dockerFile "./cmd/parser/Dockerfile" -context "."