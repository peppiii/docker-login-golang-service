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
    stage('build') {
      when {
        branch 'master'
      }
      steps {
        echo 'build docker'
      }
    }
    stage('nexus') {
      when {
        branch 'master'
      }
      steps {
        echo 'post build to nexus'
      }
    }
    stage('Approval') {
      when {
        branch 'master'
      }
      steps {
        input 'deploy to k8s'
      }
    }
  }
  environment {
    SERVICE = 'testing-golang'
  }
}
