open:
	docker exec -it newpg-container bash
#psql -U hp prithvi
#\dt
#\d students
#docker push hp2411/new_student_data:1

install:
	helm install student-release ./helm-chart
uninstall:
	helm uninstall student-release
get:
	kubectl get pods -n default
forward:
	kubectl port-forward student-app-deployment-587f8dc4fc-wpzjg 8080:8080 -n default
build:
	docker build -t hp2411/new_student_data:1 .
push:
	docker push hp2411/new_student_data:1
gen:
	protoc --proto_path=. ./api/*.proto --go_out=./api/pb --go-grpc_out=./api/pb