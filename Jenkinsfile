pipeline {
    agent any

    stages {

        stage('Checkout') {
            steps {
                checkout scm
            }
        }

        stage('Environment') {
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
                sh 'go build -v ./...'
            }
        }

        stage('Frontend Build') {
            steps {
                dir('frontend') {
                    sh 'npm install'
                    sh 'npm run build'
                }
            }
        }

    }

    post {
        success {
            echo '✅ Build Successful!'
        }

        failure {
            echo '❌ Build Failed!'
        }
    }
}