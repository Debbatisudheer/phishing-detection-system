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

                echo "========== Files =========="
                ls -la

                echo "========== package.json =========="
                cat package.json

                echo "========== npm install =========="
                npm install

                echo "========== node_modules/.bin =========="
                ls -la node_modules/.bin

                echo "========== Vite Version =========="
                npx vite --version

                echo "========== Build =========="
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
            echo 'CI BUILD SUCCESSFUL'
        }

        failure {
            echo 'CI BUILD FAILED'
        }

        always {
            cleanWs()
        }
    }
}