package dao

/*
import (
	"../global"
	"../models"
	"../tools"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"sync"
	"time"
)

var (
	jobdbname   = global.Config.DbInfo.JobDb     //工作数据库
	jobdbtye    = global.Config.DbInfo.JobDbType //数据库类型
	jobdb       *gorm.DB                         //数据库连接
	jobtx       *gorm.DB                         //事务
	jobmodified = false                          //被修改
	wgjobop     *sync.WaitGroup                  //正在操作数据库
	wgjobcommit *sync.WaitGroup                  //正在提交事务
)

func init() {
	jobdbname = global.CurrPath + jobdbname
	tdb, err := gorm.Open(jobdbtye, jobdbname)
	tools.PanicErr(err, "工作数据库初始化")
	jobdb = tdb
	if !jobdb.HasTable(&models.Job{}) {
		jobdb.CreateTable(&models.Job{})
	}
	jobtx = jobdb.Begin()
	wgjobcommit=new(sync.WaitGroup)
	wgjobop=new(sync.WaitGroup)
	go StartAutoJobCommit()
	fmt.Println("工作数据库初始化完成")
}

type tid struct {
	Id int
}

func PublishJob(job *models.Job) (id int, err error) {
	//如果需要提交的话先等待提交
	wgjobcommit.Wait()
	//正在操作
	wgjobop.Add(1)

	jobtx.Create(job)
	sql := `SELECT last_insert_rowid() as id;`
	var lid tid
	jobtx.Raw(sql).Select("id").Scan(&lid)
	id = lid.Id

	jobmodified = true
	//完成本次操作
	wgjobop.Done()
	return
}

func ShowJob(id int) (job *models.Job, err error) {
	job = new(models.Job)
	jobdb.Where(&models.Job{Id: id}).First(job)
	if job.Name == "" {
		err = global.NoSuchJob
		return
	}
	return
}

func QueryJobCount(job *models.Job) (count int) {
	jobs := new([]models.Job)
	jobdb.Where(job).Find(jobs).Count(&count)
	return
}

func QueryJob(job *models.Job, limit, page int) (jobs *[]models.Job) {
	jobs = new([]models.Job)
	jobdb.Where(job).Offset((page - 1) * limit).Limit(limit).Find(jobs)
	return
}

func UpdataJob(id int, newjob *models.Job) (err error) {
	job := new(models.Job)
	c := 0
	jobtx.Where(&models.Job{Id: id}).First(job).Count(&c)
	if c == 0 {
		err = global.NoSuchJob
		return
	}
	job.CopyJobFromEId(newjob)
	jobtx.Save(job)
	jobmodified = true
	return
}

func DeleteJob(id int) (err error) {
	job := new(models.Job)
	c := 0
	jobtx.Where(&models.Job{Id: id}).First(job).Count(&c)
	if c == 0 {
		err = global.NoSuchJob
		return
	}
	jobtx.Delete(job)
	jobmodified = true
	return
}

//根据ID查询工作
func GetJobById(jid int) (job *models.Job, err error) {
	qjob := &models.Job{Id: jid}
	job = new(models.Job)
	jobdb.Where(qjob).Find(job)
	if job.PublisherId == 0 {
		err = global.NoSuchJob
		return
	}
	return
}

func StartAutoJobCommit() {
	for {
		time.Sleep(500)
		if jobmodified {
			//先表示需要提交事务
			wgjobcommit.Add(1)
			//等待其它进程完成
			wgjobop.Wait()

			jobtx.Commit()
			jobmodified = false

			jobtx = jobdb.Begin()
			wgjobcommit.Done()
		}
	}
}
*/
