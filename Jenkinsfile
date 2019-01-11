// Run on an agent where we want to use Go<font></font>
node {
    // Install the desired Go version<font></font>
    def root = tool name: 'go-1.11', type: 'go'
    // Export environment variables pointing to the directory where Go was installed<font></font>
    // withEnv(["PATH+GO=${root}/bin"]) {
    // ws("${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}/") {
        withEnv(["GOPATH=${WORKSPACE}","GOROOT=${root}"]) {
        env.PATH="${GOPATH}/bin:$PATH"
        env.PATH="${GOROOT}/bin:$PATH"
        sh 'go version'
        stage ('checkout'){
        checkout([                      
                $class: 'GitSCM', 
                branches: [[name: '*/master']], 
                userRemoteConfigs: [[
                    credentialsId: 'e4d4cf21-2d28-4212-809c-960b68ff5c6f', 
                    url: 'git@github.com:RONGTianqi/learn_golang.git'
                ]]
            ])
        }

        stage ('build'){
            echo "构建中..."
            sh 'export GO111MODULE=on && go build'
            echo "构建完成."
        }
        stage ('test'){
            echo "单元测试中..."
            sh 'go run main.go'
            echo "单元测试完成."
        }
        stage("部署") {
            echo "部署中..."
            echo "部署完成"
        }
        }
    // }
}
