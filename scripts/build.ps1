function DockerBuild {
    param(
        [string]$imageName,
        [string]$dockerFile,
        [string]$context
    )

    docker build -t $imageName -f $dockerFile $context --progress=plain
}

DockerBuild -imageName "scraper-service" -dockerFile "./cmd/scraper/Dockerfile" -context "."
DockerBuild -imageName "parser-service" -dockerFile "./cmd/parser/Dockerfile" -context "."