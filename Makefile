test:
	mkdir -p cover && \
	go test ./kubernetes/builders -v -cover -coverprofile=cover/cov_builders.out && \
	go test ./kubernetes/builders/core/v1 -v -cover -coverprofile=cover/cov_builders_corev1.out && \
	go test ./kubernetes/builders/batch/v1 -v -cover -coverprofile=cover/cov_builders_batchv1.out && \
	go test ./kubernetes/builders/apps/v1 -v -cover -coverprofile=cover/cov_builders_appsv1.out && \
	go test ./kubernetes/actions -v -cover -coverprofile=cover/cov_actions.out

scan:
	mkdir -p cover && \
	go test ./kubernetes/builders -v -cover -coverprofile=cover/cov_builders.out && \
	go test ./kubernetes/builders/core/v1 -v -cover -coverprofile=cover/cov_builders_corev1.out && \
	go test ./kubernetes/builders/batch/v1 -v -cover -coverprofile=cover/cov_builders_batchv1.out && \
	go test ./kubernetes/builders/apps/v1 -v -cover -coverprofile=cover/cov_builders_appsv1.out && \
	go test ./kubernetes/actions -v -cover -coverprofile=cover/cov_actions.out && \
	sonar-scanner -Dsonar.projectKey=swiss-knife -Dsonar.sources=. -Dsonar.host.url=http://localhost:9000 -Dsonar.login=${SONARQUBE_TOKEN} && rm -r cover
