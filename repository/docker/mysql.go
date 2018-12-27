package docker

const mysqlStartTimeout = 10
const mysqlPortOpenTimeout = 10

type MysqlDocker struct {
	Docker        Docker
	ContainerName string
}

func (m *MysqlDocker) StartMysqlDocker() {
	mysqlOptions := map[string]string{
		"MYSQL_ROOT_PASSWORD": "root",
		"MYSQL_USER":          "go",
		"MYSQL_PASSWORD":      "root",
		"MYSQL_DATABASE":      "godb",
	}
	containerOption := ContainerOption{
		Name:              m.ContainerName,
		Options:           mysqlOptions,
		MountVolumePath:   "/var/lib/mysql",
		PortExpose:        "3306",
		ContainerFileName: "mysql:5.7",
	}
	m.Docker = Docker{
		ContainerName: m.ContainerName,
	}
	m.Docker.StopByName()
	m.Docker.Start(containerOption)
	m.Docker.WaitForStartOrKill(mysqlStartTimeout)
	m.Docker.WaitForPortOpen(mysqlPortOpenTimeout)
}

func (m *MysqlDocker) Stop() {
	m.Docker.Stop()
}
