pipeline {
    agent any

    stages {
        stage('Setup Env') {
            steps {
                sh '''
                cat <<EOF > .env
                AUTH_DB_USER=asep
                AUTH_DB_PASSWORD=123456
                AUTH_DB_NAME=auth_db
                JWT_SECRET=supersecret
                EOF
                '''
            }
        }

        stage('Build Docker') {
            steps {
                sh 'docker-compose build'
            }
        }

        stage('Force Clean Old Containers') {
            steps {
                sh '''
                docker rm -f auth-service || true
                docker rm -f article-service || true
                docker rm -f gateway-service || true
                docker rm -f mongodb || true
                docker rm -f postgres || true
                '''
            }
        }

        stage('Deploy') {
            steps {
                sh '''
                docker-compose up -d
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