library 'ops-jenkins-libs@master'

pipeline {
  agent none
  stages {
    stage ('PrepBuild') {
      agent {
        docker {
          args '--net=host'
          image "${env.BUILDPREP}"
          label 'dind'
          registryCredentialsId "${env.WEAVEREGISTRYCREDS}"
          registryUrl "${env.WEAVEREGISTRY}"
        }
      }
      steps {
        script {
          weave = readYaml file: '.weave.yaml'
          sh '/buildprep'
        }
        stash name: 'buildenv', includes: "buildprep.env"
      }
    }
    stage('Build') {
      agent any
        environment {
          BOOT_WAIT = "30"
          POSTGRES_DB = "insys_onboarding_test"
          POSTGRES_SEARCH_PATH = "insys_onboarding"
          POSTGRES_USER = "postgres"
        }
      steps {
        script {
          // Log into Weave container registry
          docker.withRegistry("${env.WEAVEREGISTRY}", "${env.WEAVEREGISTRYCREDS}") {
            // Bootstrap our sidecar
            docker.image('postgres:11-alpine').withRun("--env POSTGRES_HOST_AUTH_METHOD=trust --env POSTGRES_DB=${env.POSTGRES_DB} --env POSTGRES_USER=${env.POSTGRES_USER}") { psql ->
              // Resolve IP address for our service container since we can't rely on /etc/hosts modifications
              psqlIP = sh(script: "docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' ${psql.id}", returnStdout: true).trim()
                // Test network connectivity / wait for sidecar to be ready (optional)
                docker.image('postgres:11-alpine').inside("-e 'PSQL=${psqlIP}' -e 'WAIT_SECONDS=${env.BOOT_WAIT}'") {
                  sh '''
                    /usr/local/bin/pg_isready -h \$PSQL -t \$WAIT_SECONDS
                    psql -U \$POSTGRES_USER -h \$PSQL -d \$POSTGRES_DB -c "CREATE SCHEMA \$POSTGRES_SCHEMA;"
                  '''
                }
              psqlDSN = "postgresql://${POSTGRES_USER}@${psqlIP}/${env.POSTGRES_DB}?sslmode=disable&search_path=${env.POSTGRES_SCHEMA}"
              docker.image("${env.WEAVEBUILDER}").inside("-e PG_PRIMARY_CONNECT_STRING=${psqlDSN}") {
                unstash "buildenv"
                withCredentials([file(credentialsId: 'weavelabbotkey', variable: 'WEAVELABBOT')]) {
                  sh '''
                    . ./buildprep.env
                    cat $WEAVELABBOT > ~/.ssh/id_rsa

                    # This allows us to install goose without installing all go mod dependencies
                    MY_CWD=$(pwd)
                    cd ~
                    go get -u github.com/pressly/goose/cmd/goose
                    cd $MY_CWD

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
    }
    stage('Package') {
      agent {
        docker {
          args "--privileged"
          label "dind"
          image "${env.CDTOOLS}"
          registryCredentialsId "${env.WEAVEREGISTRYCREDS}"
          registryUrl "${env.WEAVEREGISTRY}"
        }
      }
      steps {
        unstash "bins"
        unstash "buildenv"
        withCredentials([file(credentialsId: 'build-secrets', variable: 'BUILDSECRETS')]) {
          sh '''
            . ./buildprep.env
            /bin/cd-tools
          '''
        }
      }
    }
  }
  post {
    failure { slackSend color: '#DD0000', channel: "${weave.slack}", message: "${env.JOB_NAME} build failed. ${env.RUN_DISPLAY_URL}"}
  }
}