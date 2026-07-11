pipeline {
    agent any

    options {
        timestamps()
    }

    environment {
        GO111MODULE = 'on'

        DB_HOST = 'postgres-ci'
        DB_PORT = '5432'
        DB_USER = 'postgres'
        DB_PASSWORD = 'postgres'
        DB_NAME = 'phishing_platform'
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
                sh 'docker --version'
            }
        }

        stage('Start PostgreSQL') {
            steps {
                sh '''
                    docker rm -f postgres-ci || true

                    docker run -d \
                      --name postgres-ci \
                      -e POSTGRES_USER=postgres \
                      -e POSTGRES_PASSWORD=postgres \
                      -e POSTGRES_DB=phishing_platform \
                      -p 5432:5432 \
                      postgres:17

                    echo "Waiting for PostgreSQL..."

                    sleep 15
                '''
            }
        }

        stage('Backend Build') {
            steps {
                sh '''
                    go mod download
                    go build ./...
                '''
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

        stage('Run All Tests') {
            steps {
                sh '''
                    go test ./... -v
                '''
            }
        }
    }

    post {

        success {
            echo "====================================="
            echo "CI BUILD SUCCESSFUL"
            echo "====================================="
        }

        failure {
            echo "====================================="
            echo "CI BUILD FAILED"
            echo "====================================="
        }

        always {

            sh '''
                docker rm -f postgres-ci || true
            '''

            cleanWs()
        }
    }
}