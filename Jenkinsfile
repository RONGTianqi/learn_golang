// Run on an agent where we want to use Go<font></font>
node {
    // Install the desired Go version<font></font>
    def root = tool name: 'Go 1.8', type: 'go'
    // Export environment variables pointing to the directory where Go was installed<font></font>
    withEnv(["GOROOT=${root}", "PATH+GO=${root}/bin"]) {
        sh 'go version'
    }
}