pipeline {
  agent none
  stages {
    stage ('readWeave') {
      agent any
      steps {
        script {
          weave = readYaml file: '.weave.yaml'
        }
      }
    }
    stage('goBuild') {
      agent {
        docker {
          image "${env.WEAVEBUILDER}"
          registryUrl "${env.WEAVEREGISTRY}"
          registryCredentialsId "${env.WEAVEREGISTRYCREDS}"
          label 'dind'
        }
      }
      steps {
        sh '/usr/local/bin/gobuilder'
        stash name: 'bins', includes: "${weave.slug}"
      }
    }
    stage('cdTools') {
      agent {
        docker {
          image "${env.CDTOOLS}"
          registryUrl "${env.WEAVEREGISTRY}"
          registryCredentialsId "${env.WEAVEREGISTRYCREDS}"
          label 'dind'
        }
      }
      steps {
        unstash 'bins'
        sh '/bin/cd-tools'
      }
    }
  }
  post {
    success { slackSend color: '#00DD00', channel: "${weave.slack}", message: "${env.JOB_NAME} build success. ${env.RUN_DISPLAY_URL}" }
    failure { slackSend color: '#DD0000', channel: "${weave.slack}", message: "${env.JOB_NAME} build failed. ${env.RUN_DISPLAY_URL}"}
  }
}