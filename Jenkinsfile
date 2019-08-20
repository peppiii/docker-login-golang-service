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
        stage('Build') {
            parallel {
                stage('Code review') {
                    environment {
                        scannerHome = tool 'sonarQubeScanner'
                    }
                    when {
                        branch 'develop'
                    }
                    steps {
                        withSonarQubeEnv('sonarQube') {
                            sh "${scannerHome}/bin/sonar-scanner"
                        }
                        timeout(time: 5, unit: 'MINUTES') {
                            waitForQualityGate abortPipeline: true
                        }
                    }
                }
