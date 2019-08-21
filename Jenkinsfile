pipeline {
    agent any 
    
    environment {
        SERVICE = 'testing-golang'
    }
    stages {
        stage('Checkout') {
            when {
                                branch 'master'
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
                        branch 'master'
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
            }
        }
    }
}
