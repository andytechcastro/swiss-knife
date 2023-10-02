mini-host = $(shell minikube service sonar-svc --url -n sonar)

test:
	mkdir -p cover && \
	go test ./kubernetes/builders -v -cover -coverprofile=cover/cov_builders.out && \
	go test ./kubernetes/builders/core/v1 -v -cover -coverprofile=cover/cov_builders_corev1.out && \
	go test ./kubernetes/builders/batch/v1 -v -cover -coverprofile=cover/cov_builders_batchv1.out && \
	go test ./kubernetes/builders/apps/v1 -v -cover -coverprofile=cover/cov_builders_appsv1.out && \
	go test ./kubernetes/actions -v -cover -coverprofile=cover/cov_actions.out && \
	go test ./kubernetes/actions/core/v1 -v -cover -coverprofile=cover/cov_actions_corev1.out && \
	go test ./kubernetes/actions/batch/v1 -v -cover -coverprofile=cover/cov_actions_batchv1.out && \
	go test ./kubernetes/actions/apps/v1 -v -cover -coverprofile=cover/cov_actions_appsv1.out

scan:
	@echo $(mini-host)
	mkdir -p cover && \
	go test ./kubernetes/builders -v -cover -coverprofile=cover/cov_builders.out && \
	go test ./kubernetes/builders/core/v1 -v -cover -coverprofile=cover/cov_builders_corev1.out && \
	go test ./kubernetes/builders/batch/v1 -v -cover -coverprofile=cover/cov_builders_batchv1.out && \
	go test ./kubernetes/builders/apps/v1 -v -cover -coverprofile=cover/cov_builders_appsv1.out && \
	go test ./kubernetes/actions -v -cover -coverprofile=cover/cov_actions.out && \
	go test ./kubernetes/actions/core/v1 -v -cover -coverprofile=cover/cov_actions_corev1.out && \
	go test ./kubernetes/actions/batch/v1 -v -cover -coverprofile=cover/cov_actions_batchv1.out && \
	go test ./kubernetes/actions/apps/v1 -v -cover -coverprofile=cover/cov_actions_appsv1.out && \
	sonar-scanner -Dsonar.projectKey=swiss-knife -Dsonar.sources=. -Dsonar.host.url=${mini-host} -Dsonar.login=${SONARQUBE_TOKEN} && rm -r cover
