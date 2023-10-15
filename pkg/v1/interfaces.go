package v1

import (
	"lms/internal/models"


)

type RepoInterface interface{
SaveLog(models.Log) (models.Log,error)

}
type UsecaseInterface interface{
SaveLog(models.Log) (models.Log,error)
}