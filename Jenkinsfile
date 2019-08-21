pipeline {
    agent any 
    
    environment {
        SERVICE = 'testing-golang'
    }
    stages {
        stage('Checkout') {
            when {
                anyOf { branch 'master'; branch 'develop'; branch 'staging' }
            }
            steps {
                echo 'Checking out from Git'
                checkout scm
            }
        }
    }
}
