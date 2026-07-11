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
                sh '''
                    git --version
                    go version
                    node -v
                    npm -v
                    docker --version
                '''
            }
        }

        stage('Database Check') {
            steps {
                sh '''
                    echo "========== DATABASE ENV =========="
                    echo "DB_HOST=$DB_HOST"
                    echo "DB_PORT=$DB_PORT"
                    echo "DB_USER=$DB_USER"
                    echo "DB_PASSWORD=$DB_PASSWORD"
                    echo "DB_NAME=$DB_NAME"

                    echo "========== DNS =========="
                    getent hosts postgres-ci || true

                    echo "========== RUNNING CONTAINERS =========="
                    docker ps
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

        stage('Run Backend Tests') {
            steps {
                sh '''
                    go test ./... -v
                '''
            }
        }
    }

    post {

        success {
            echo '====================================='
            echo 'CI BUILD SUCCESSFUL'
            echo '====================================='
        }

        failure {
            echo '====================================='
            echo 'CI BUILD FAILED'
            echo '====================================='
        }

        always {
            cleanWs()
        }
    }
}