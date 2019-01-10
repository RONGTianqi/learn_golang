
// Run on an agent where we want to use Go
node {
    // Install the desired Go version
    def root = tool name: 'go-1.11', type: 'go'
 
    // Export environment variables pointing to the directory where Go was installed
    withEnv(["GOROOT=${root}", "PATH+GO=${root}/bin"]) {
        sh 'go version'
        stage ('build'){
            echo "构建中..."
            sh 'go build'
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
}
