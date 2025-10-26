pipeline {
    agent any

    environment {
        DOCKER_USER = 'haroldbg' //  Tu usuario de Docker Hub
        DOCKER_PASS = credentials('dockerhub-token') //  ID de las credenciales guardadas en Jenkins
        IMAGE_NAME = 'haroldbg/task-movil-api-go' //  Nombre de la imagen que se subirá a Docker Hub
    }

    stages {
        stage('Clonar repositorio') {
            steps {
                echo "📥 Clonando el repositorio del backend Go..."
                // Jenkins ya lo clona automáticamente si el Jenkinsfile está en el repo
                sh 'git pull origin main || true'
            }
        }

        stage('Construir imagen Docker') {
            steps {
                echo "⚙️ Construyendo imagen Docker del backend Go..."
                sh 'docker build -t $IMAGE_NAME:latest .'
            }
        }

        stage('Subir imagen a Docker Hub') {
            steps {
                echo "⬆️ Subiendo imagen a Docker Hub..."
                withCredentials([string(credentialsId: 'dockerhub-token', variable: 'DOCKER_PASS')]) {
                    sh '''
                        echo "$DOCKER_PASS" | docker login -u "$DOCKER_USER" --password-stdin
                        docker push $IMAGE_NAME:latest
                    '''
                }
            }
        }

        stage('Desplegar con Docker Compose') {
            steps {
                echo "🚀 Desplegando contenedores (PostgreSQL + Backend Go)..."
                sh '''
                    docker compose down || true
                    docker compose up -d --build
                '''
            }
        }
    }

    post {
        success {
            echo "✅ Despliegue completado correctamente. El backend Go está en marcha."
        }
        failure {
            echo "❌ Error en el pipeline. Revisa los logs en Jenkins."
        }
    }
}
