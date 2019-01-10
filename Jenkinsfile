pipeline {
    agent any
    tools{
          go 'Go 1.8' 
    }
    
    stages  {
        
        stage("checkout") {
            steps {
                // sh 'ci-init'
                // checkout(
                // [$class: 'GitSCM', branches: [[name: env.GIT_BUILD_REF]], 
                // userRemoteConfigs: [[url: env.GIT_REPO_URL]]]
                //   [$class: 'GitSCM', branches: [[name: '*/master']], 
                //   userRemoteConfigs: [[url: 'git@git.dev.tencent.com:RONGTianqi/dev-demo.git']]]
                )
            }
        }

        stage("build") {
            steps {
                echo "构建中..."
                sh 'go version'
                // sh 'export GOPATH=$WORKSPACE/gopath'
                // sh 'export PATH=$GOPATH/bin:$PATH'
                // sh 'mkdir -p $GOPATH/bin'
                // sh 'mkdir -p $GOPATH/src/github.com/'
                // sh 'go get github.com/gin-gonic/gin'
                sh 'go build'
                // sh 'go install'
                // 请在这里放置您项目代码的单元测试调用过程，例如:
                // sh 'mvn package' // mvn 示例
                // sh 'make' // make 示例
                echo "构建完成."
                // archiveArtifacts artifacts: '**/target/*.jar', fingerprint: true // 收集构建产物
            }
        }

        stage("测试") {
            steps {
                echo "单元测试中..."
                sh 'go run main.go'
                // 请在这里放置您项目代码的单元测试调用过程，例如:
                // sh 'mvn test' // mvn 示例
                // sh 'make test' // make 示例
                echo "单元测试完成."
                // junit 'target/surefire-reports/*.xml' // 收集单元测试报告的调用过程
            }
        }

        stage("部署") {
            steps {
                echo "部署中..."
                // 请在这里放置收集单元测试报告的调用过程，例如:
                // sh 'mvn tomcat7:deploy' // Maven tomcat7 插件示例：
                // sh './deploy.sh' // 自研部署脚本
                echo "部署完成"
            }
        }
    }
}