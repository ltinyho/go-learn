module github.com/ltinyho/go-learn

go 1.14

require (
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-gonic/gin v1.5.0
	github.com/go-redis/redis v6.15.5+incompatible
	github.com/go-zookeeper/zk v1.0.2
	github.com/google/uuid v1.1.1
	github.com/onsi/ginkgo v1.14.1 // indirect
	github.com/onsi/gomega v1.10.2 // indirect
	github.com/sirupsen/logrus v1.6.0
	github.com/spf13/viper v1.4.0
	github.com/stretchr/testify v1.5.1
	gogs.sharkgulf.cn/sg/library v0.1.40
	gorm.io/driver/mysql v1.0.1
	gorm.io/gorm v1.20.0
)

replace gogs.sharkgulf.cn/sg/library v0.1.40 => ../../sg/library
