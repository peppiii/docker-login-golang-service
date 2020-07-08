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
        input 'deploy to argocd'
      }
    }

    stage('deploy to argocd') {
      when {
        branch 'master'
      }
      steps {
        echo 'deploy to argocd'
      }
    }

  }

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
    stage('Approval Code Review') {
      when {
        branch 'develop'
      }
      steps {
        input 'Approval Code Review'
      }
    }

    stage('Build Application') {
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

    stage('push to registry') {
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

    stage('Approval') {
      when {
        branch 'master'
      }
      steps {
        input 'deploy to argocd'
      }
    }

    stage('Approval Staging') {
      when {
        branch 'staging'
      }
      steps {
        input 'Approval Staging'
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
    post {
          success {
            slackSend color: '#00FF00', message: "Job ${env.JOB_NAME} build ${env.BUILD_NUMBER} SUCCESS (<${env.BUILD_URL}|Open>)"
          }
          failure {
            slackSend color: '#FF0000', message: "Job ${env.JOB_NAME} build ${env.BUILD_NUMBER} FAILED (<${env.BUILD_URL}|Open>)"
          }
        }
  }
  environment {
    SERVICE = 'testing-golang'
  }
}
