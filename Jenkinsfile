pipeline {
    agent any

    stages {
        stage('Pull Code') {
            steps {
                git 'https://github.com/AsepNurdinDev/microservice-app.git'
            }
        }

        stage('Build Docker') {
            steps {
                sh 'docker compose build'
            }
        }

        stage('Deploy') {
            steps {
                sh '''
                docker compose down
                docker compose up -d
                '''
            }
        }

        stage('Cleanup') {
            steps {
                sh 'docker system prune -f'
            }
        }
    }
}
