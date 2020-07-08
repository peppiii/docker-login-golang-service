pipeline {
  agent any
  stages {
    stage('Checkout') {
      when {
        anyOf {
          branch 'master'
          branch 'develop'
          branch 'staging'
        }

      }
      steps {
        echo 'Checking out from Git'
        checkout scm
      }
    }

    stage('Sanity check') {
      when {
        branch 'master'
      }
      steps {
        input 'Does the staging environment look ok?'
      }
    }

  }
  environment {
    SERVICE = 'testing-golang'
  }
}