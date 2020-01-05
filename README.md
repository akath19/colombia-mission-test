[![Build Status](https://img.shields.io/endpoint.svg?url=https%3A%2F%2Factions-badge.atrox.dev%2Fakath19%2Fcolombia-mission-test%2Fbadge%3Fref%3Dmaster&style=popout)](https://actions-badge.atrox.dev/akath19/colombia-mission-test/goto?ref=master)
# Colombia Mission Technical Test
This repo holds all files requested by the Colombia Mission Technical Operations Test

## Configuration
The following environment variables must be passed to the Docker container for the app to start properly

| Property | Description |
| -------- | ----------- |
| POSTGRES_ADDR | PostgreSQL DB to connect to, this address must not contain a protocol |
| POSTGRES_PORT | PostgreSQL DB port to connect to |
| POSTGRES_USER | PostgreSQL user to connect as |
| POSTGRES_PASS | PostgreSQL password for the above user |
| POSTGRES_DB | PostgreSQL database to use |
| HTTP_PORT | http port to expose service in |

Please remember to create a user with `SELECT, INSERT, UPDATE, DELETE & CONNECT` privileges to the database selected in `POSTGRES_DB`

## How To Run
1. Clone this repo
1. Download & install golang from [here](https://golang.org/)
1. Run `go get` to download all dependencies
1. Run `go build` to create the final executable
1. Run `colombia-mission-test` executable after setting up the above environment variables & PostgreSQL instance

## Ansible Automation
The `ansible` folder includes an ansible playbook to do the following:

1. Create a Kubernetes Cluster in Google Cloud (GKE)
1. Save the associated kubeconfig file for the cluster
1. Create a Node Pool in said cluster
1. Create Database inside PostgreSQL
1. Deploy the application as described in the `deployment.yaml` file

To run the playbook, the following libraries must be installed:

1. requests >= 2.18.4
1. google_auth >= 1.3.0
1. pyyaml >= 5.0

The folder includes a `requirements.txt` file to install all required libraries.

Once all libraries are installed, a service account must be created in the project with the following roles:

1. Service Account User (to allow the service account to run tasks)
1. Kubernetes Engine Admin (to create the cluster)
1. Compute Engine Admin (to create the cluster)
1. CloudSQL Admin (to create the CloudSQL instance)

Download a Service Account JSON authentication file and place it in your local disk, the location of the file must be referenced in the ansible `vars.yaml` file.

Also, create a CloudSQL instance in the project, this instance is not created as part of the playbook due to data consistency.

Finally, update the `k8s/configmap.yaml` & `k8s/deployment.yaml` to add your PostgreSQL/CloudSQL values 