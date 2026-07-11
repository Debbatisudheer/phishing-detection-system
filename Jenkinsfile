pipeline {
    agent any

    options {
        timestamps()
    }

    environment {
        GO111MODULE = 'on'
        NODE_ENV = 'production'
    }

    stages {

        stage('Checkout Source') {
            steps {
                checkout scm
            }
        }

        stage('Environment Check') {
            steps {
                sh 'git --version'
                sh 'go version'
                sh 'node -v'
                sh 'npm -v'
            }
        }

        stage('Backend Build') {
            steps {
                sh 'go mod download'
                sh 'go build ./...'
            }
        }

        stage('Frontend Build') {
    steps {
        dir('frontend') {
            sh '''
                pwd

                npm config list

                npm root

                npm prefix

                npm install

                ls -la

                ls -la node_modules || true

                ls -la node_modules/.bin || true
            '''
        }
    }
}

        stage('Backend Tests') {
            steps {
                sh 'go test ./... -v'
            }
        }
    }

    post {
        success {
            echo 'CI BUILD SUCCESSFUL'
        }

        failure {
            echo 'CI BUILD FAILED'
        }

    }
}