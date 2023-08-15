test:
	mkdir cover && go test ./kubernetes/builders -v -cover -coverprofile=cover/cov_builders.out && go test ./kubernetes/actions -v -cover -coverprofile=cover/cov_actions.out

scan:
	mkdir cover && go test ./kubernetes/builders -v -cover -coverprofile=cover/cov_builders.out && go test ./kubernetes/actions -v -cover -coverprofile=cover/cov_actions.out && sonar-scanner -Dsonar.projectKey=swiss-knife -Dsonar.sources=. -Dsonar.host.url=http://localhost:9000 -Dsonar.login=${SONARQUBE_TOKEN} && rm -r cover
