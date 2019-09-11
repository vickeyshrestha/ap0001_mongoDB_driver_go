// 1. define the node where Jenkins will be running
node ('vic_node && server_1') {
    def pwd = pwd()
    try {
        stage("clean and prepare environment") {
            env.MONGO_ENGINE_VERSION = "00.00.07"
            env.GOROOT = "/usr/local/go"
            env.GOPATH = "${pwd}/go"
            env.PATH = "${env.GOROOT}/bin;${pwd}/go/bin:/usr/local/bin:${env.PATH}"
            sh "rm -rf ${pwd}/go"
            sh "mkdir ${pwd}/go"
            dir ("${pwd}/go") {
                sh "mkdir src; mkdir bin; cd src"
            }
        }

        stage ('pull code') {
            dir("${pwd}/go/src/ap0001_mongo_engine") {
                git branch: "${env.BRANCH_NAME}", credentialsId: 'e1W5csqvaS-T81qOcbewANuHvHBkQvhKu-UcAQ5Cgxkz6FCr76', url: 'https://github.com/vickeyshrestha/ap0001_mongo_engine.git'
            }
        }

        stage ('Test') {
            dir("${pwd}/go/src/ap0001_mongo_engine") {
                sh "go get -u github.com/jstemmer/go-juni-report"
                sh "go test -v ./... > output.out"
                sh "cat output.out | go-junit-report > report.xml"
                junit 'report.xml'
            }
        }

        stage ('Test Coverage') {
            dir ("${pwd}/go/src/ap0001_mongo_engine") {
                sh '''  go get github.com/axw/gocov/gocov;
                        go get github.com/math/gocov-html;
                        go get github.com/AlekSi/gocov-xml;
                    '''
                def coverageHtml = sh(returnStdout: true, script: '''gocov test -v $(go list ./... | grep -v /vendor) | gocov-html''')
                writeFile file: 'coverage.html', text: "$coverageHtml"
                checkCoverage()
                publishHtml([allowMissing: true, alwaysLinkToLastBuild: true, keepAll: true, reportDir: '', reportFiles: 'coverage.html', reportName: 'Test Coverage Report'])
            }
        }

        stage ('Static analysis') {
            dir ("${pwd}/go/src/ap0001_mongo_engine") {
                sh "go vet \$(go list ./... | grep -v /vendor)"
            }
        }

        stage ('build') {
             dir ("${pwd}/go/src/ap0001_mongo_engine") {
                sh "CGO_ENABLED=0 GOOS=linux go build -o ap0001_mongo_engine -ldflags '-w -s' cmd/main.go"
             }
        }

        stage ('Build Docker Image') {
             dir ("${pwd}/go/src/ap0001_mongo_engine") {
                sh "docker build -t vickeyshrestha/ap0001_mongo_engine:${MONGO_ENGINE_VERSION} ."
             }
        }

        // NEW - We will do some performance testing as well
        stage ('save and stash image in Jenkins node') {
            dir ("${pwd}/go/src/ap0001_mongo_engine") {
                sh "docker save vickeyshrestha/ap0001_mongo_engine:${MONGO_ENGINE_VERSION} > mongoEng.tar"
                stash name: "DockerImage", includes: "mongoEng.tar"
                stash name: "JMeter", includes: "performanceTest/MongoEngine_performance_10TPS.jmx"
            }
        }

        // Now we will unstash in server_2, our server for exclusively running performance test only, and run JMeter over it
        node('server_2') {
            stage ('Load and Unstash on performance node') {
                unstash "DockerImage"
                unstash "JMeter"
                sh "docker load < mongoEng.tar"
                sh "ls -lart"
                sh "docker images"
            }

            stage ('Performance Test') {
                // Looking into JMeter and will update this portion later
            }
        }

        // if everything looks fine, time to push into docker repository
        stage ('Push to Dockerhub') {
            if (env.BRANCH_NAME == 'master') {
                dir ("${pwd}/go/src/ap0001_mongo_engine") {
                    sh "docker push vickeyshrestha/ap0001_mongo_engine:${MONGO_ENGINE_VERSION}"
                }
            } else {
                echo "skipping image push because it is not a master branch."
            }
        }

    } catch (e) {
        currentBuild.result = "FAILED"
        throw e
    } finally {
        notifyDevelopers(currentBuild.result)
    }
}

def notifyDevelopers(String buildStatus) {
    subject = "${buildStatus}: Build ${env.BRANCH_NAME} ${env.BUILD_NUMBER}"
    details = "${buildStatus}: \n Job: ${env.JOB_NAME} \n Build#: ${env.BUILD_NUMBER} \n The latest build failed. Please check attached console output or view console output at ${env.BUILD_URL}"

    if ((buildStatus == "FAILURE") && (env.BRANCH_NAME == 'master')) {
        emailext attachLog: true, body: details, subject: subject, to: 'vickey.shrestha.1987@gmail.com'
    } else if (buildStatus == "FAILURE") {
        emailext attachLog: true, body:details, recipientProviders: [[$Class:'culpritsRecipientProvider'], [$class: 'DevelopersRecipientProvider']], subject: subject
    }
}

def checkCoverage() {
    def total = sh(returnStdout: true, script: '''totalCoverage=$(awk -F "[><]" '/totalcov/{print $3}' coverage.html)
    totalCoverage = ${totalCoverage%?}
    echo "${totalCoverage}"''')

    double coverageThreshold = 72
    double totalCoverage = total.trim() as Double
    if (totalCoverage < coverageThreshold) {
        error "Test coverage not enough. Currently at ${totalCoverage}%. Requires ${coverageThreshold}%"
    }
}