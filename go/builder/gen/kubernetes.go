package gen

import (
	"fmt"
)

func Kubernetes(param GenParameters) {
	path, err := CreateFolder(param, "kubernetes")
	if err != nil {
		errMsg := fmt.Errorf("%s", err)
		fmt.Println(errMsg)
	}
	fmt.Println("create sample valina k8s yaml files:", path)

	const readmeContent = `
# This is a Readme file HELM
`
	err = CreateTextFile(path, "README.md", readmeContent)
	if err != nil {
		errMsg := fmt.Errorf("%s", err)
		fmt.Println(errMsg)
	}

	const helloApp = `
# Copyright 2020 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

apiVersion: apps/v1
kind: Deployment
metadata:
  name: hello-app
  labels:
    app: hello
spec:
  selector:
    matchLabels:
      app: hello
      tier: web
  template:
    metadata:
      labels:
        app: hello
        tier: web
    spec:
      containers:
      - name: hello-app
        image: gcr.io/google-samples/hello-app:1.0
        imagePullPolicy: Never
        ports:
        - containerPort: 8080
`
	err = CreateTextFile(path, "hello-app.yaml", helloApp)
	if err != nil {
		errMsg := fmt.Errorf("%s", err)
		fmt.Println(errMsg)
	}
}
