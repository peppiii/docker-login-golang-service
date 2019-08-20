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
    }
}
