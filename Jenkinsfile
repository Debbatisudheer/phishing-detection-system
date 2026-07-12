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

        stage('Verify Database') {
            steps {
                sh '''
                    echo "========== DATABASE ENV =========="
                    echo "DB_HOST=$DB_HOST"
                    echo "DB_PORT=$DB_PORT"
                    echo "DB_USER=$DB_USER"
                    echo "DB_NAME=$DB_NAME"

                    echo "========== DNS =========="
                    getent hosts $DB_HOST || true

                    echo "========== RUNNING CONTAINERS =========="
                    docker ps
                '''
            }
        }

        stage('Import Database Schema') {
            steps {
                sh '''
                    echo "Waiting for PostgreSQL..."

                    until docker exec postgres-ci pg_isready -U postgres
                    do
                        sleep 2
                    done

                    echo "Importing schema..."

                    docker exec -i postgres-ci \
                        psql \
                        -U postgres \
                        -d phishing_platform \
                        < database/schema.sql
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

        stage('Backend Tests') {
            steps {
                sh '''
                    go test ./... -v
                '''
            }
        }

        stage('UI Tests (Playwright)') {
            steps {
                dir('qa') {
                    sh '''
                        npm install

                        npx playwright install --with-deps

                        npx playwright test
                    '''
                }
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
            cleanWs()
        }
    }
}