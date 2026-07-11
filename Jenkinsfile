pipeline {
    agent any

    options {
        timestamps()
        ansiColor('xterm')
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
                sh 'echo "===== Environment ====="'
                sh 'git --version'
                sh 'go version'
                sh 'node -v'
                sh 'npm -v'
            }
        }

        stage('Backend Build') {
            steps {
                sh 'echo "===== Building Backend ====="'
                sh 'go mod download'
                sh 'go build ./...'
            }
        }

        stage('Frontend Build') {
            steps {
                dir('frontend') {
                    sh 'echo "===== Installing Frontend Dependencies ====="'
                    sh 'npm ci'

                    sh 'echo "===== Building Frontend ====="'
                    sh 'npm run build'
                }
            }
        }

        stage('Backend Tests') {
            steps {
                sh 'echo "===== Running Go Tests ====="'
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