node {
    def scmVars = checkout scm
    def image
    def imageName = "willwangkelda/hotrod-driver:release-${scmVars.GIT_COMMIT}"

    stage('Build') {
        image = docker.build(imageName)
    }
    stage('Push') {
        image.push()
    }
}
