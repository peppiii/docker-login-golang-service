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
    stage('build application') {
      when {
        anyOf {
          branch 'master'
          branch 'develop'
          branch 'staging'
        }

      }
      steps {
        echo 'build docker'
      }
    }

    stage('post to registry') {
      when {
        anyOf {
          branch 'master'
          branch 'develop'
          branch 'staging'
        }

      }
      steps {
        echo 'post build to nexus'
      }
    }

    stage('Approval Code Review') {
      when {
        branch 'develop'
      }
      steps {
        input 'approval code review'
      }
    }

    stage('Approval Qa') {
      when {
        branch 'staging'
      }
      steps {
        input 'Approval qa'
      }
    }

    stage('Approval to k8s') {
      when {
        anyOf {
          branch 'master'
          branch 'develop'
          branch 'staging'
        }
      }
      steps {
        input 'deploy to argocd'
      }
    }

    stage('deploy to argocd') {
      when {
        anyOf {
          branch 'master'
          branch 'develop'
          branch 'staging'
        }
      }
      steps {
        echo 'deploy to argocd'
      }
    }
  }
  environment {
    SERVICE = 'testing-golang'
  }
}
