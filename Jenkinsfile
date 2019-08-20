pipeline {
    agent {
        node {
            label 'master'
        }
    }
    environment {
        SERVICE = 'testing-golang'
    }
    stages {
        stage('Checkout') {
            when {
                anyOf { branch 'master';}
            }
            steps {
                echo 'Checking out from Git'
                checkout scm
            }
        }
    }
