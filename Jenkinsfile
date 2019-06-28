pipeline {
  agent none

    environment {
      CI = 'true'
    }
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
      agent any
        environment {
          POSTGRES_DB = "insys_onboarding_test"
          POSTGRES_USER = "postgres"
          POSTGRES_SEARCH_PATH = "insys_onboarding"
        }
      steps {
        script {
          // Log into Weave container registry
          docker.withRegistry("${env.WEAVEREGISTRY}", "${env.WEAVEREGISTRYCREDS}") {
            // Bootstrap our sidecar
            docker.image('postgres:11-alpine').withRun("--env POSTGRES_DB=${env.POSTGRES_DB} --env POSTGRES_USER=${env.POSTGRES_USER}") { psql ->
              // Resolve IP address for our service container since we can't rely on /etc/hosts modifications
              psqlIP = sh(script: "docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' ${psql.id}", returnStdout: true).trim()
                // Test network connectivity / wait for sidecar to be ready (optional)
                docker.image('postgres:11-alpine').inside("-e 'PSQL=${psqlIP}' -e 'WAIT_SECONDS=${env.BOOT_WAIT}'") {
                  sh '''
                    /usr/local/bin/pg_isready -h \$PSQL -t \$WAIT_SECONDS
                    psql -U \$POSTGRES_USER -h \$PSQL -d insys_onboarding_test -c "CREATE SCHEMA insys_onboarding;"
                  '''
                }
              psqlDSN = "postgresql://${POSTGRES_USER}@${psqlIP}/${env.POSTGRES_DB}?sslmode=disable&search_path=${env.POSTGRES_SEARCH_PATH}"
                docker.image("${env.WEAVEBUILDER}").inside("-e PG_PRIMARY_CONNECT_STRING=${psqlDSN}") {
                  sh '''
                    go get -u github.com/pressly/goose/cmd/goose
                    goose -dir ./dbconfig/migrations postgres $PG_PRIMARY_CONNECT_STRING up
                    /usr/local/bin/gobuilder
                    '''
                    stash name: 'bins', includes: "${weave.slug}"
                }
            }
          }
        }
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
