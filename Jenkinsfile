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
    }
}