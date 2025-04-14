pipeline {
    agent any

    environment {
        HARBOR_DOMAIN = "harbor.steveay.com"
        DOCKER_REGISTRY = "${HARBOR_DOMAIN}/lottery/go-zero-lottery"
        IMAGE_TAG = "alpha"
    }

    stages {
        stage('Configure Jenkins') {
            steps {
                script {
                    System.setProperty("org.jenkinsci.plugins.durabletask.BourneShellScript.HEARTBEAT_CHECK_INTERVAL", "86400")
                }
            }
        }

        stage('Login to Harbor') {
            steps {
                script {
                    withCredentials([usernamePassword(credentialsId: 'harbor', usernameVariable: 'USERNAME', passwordVariable: 'PASSWORD')]) {
                        sh '''
                        echo "$PASSWORD" | docker login --username "$USERNAME" --password-stdin $HARBOR_DOMAIN
                        '''
                    }
                }
            }
        }

        stage('Build Docker Image') {
            steps {
                sh "docker build --platform=linux/amd64 -t ${DOCKER_REGISTRY}:${IMAGE_TAG} ."
            }
        }

        stage('Push Docker Image') {
            steps {
                sh "docker push ${DOCKER_REGISTRY}:${IMAGE_TAG}"
            }
        }
    }
}
