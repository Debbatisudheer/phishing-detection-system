pipeline {
    agent any

    stages {

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
                    sh 'npm install'
                    sh 'npm run build'
                }
            }
        }

    }
}