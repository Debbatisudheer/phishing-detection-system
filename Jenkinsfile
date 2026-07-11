pipeline {
    agent any

    options {
        timestamps()
    }

    environment {
        GO111MODULE = 'on'
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
                        npm install --include=dev
                        npm run build
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
            echo '====================================='
            echo ' CI BUILD SUCCESSFUL '
            echo '====================================='
        }

        failure {
            echo '====================================='
            echo ' CI BUILD FAILED '
            echo '====================================='
        }

        always {
            cleanWs()
        }
    }
}