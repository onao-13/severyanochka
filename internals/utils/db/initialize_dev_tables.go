package db

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
)

// TableUtils Утилиты для работы с бд
type TableUtils struct {
	pool *pgxpool.Pool
	ctx  context.Context
}

var log = logrus.New()

// New Создание утилиты
func New(pool *pgxpool.Pool, ctx context.Context) *TableUtils {
	util := new(TableUtils)
	util.pool = pool
	util.ctx = ctx
	return util
}

// CreateTables Создает таблицы, которые указаны в файле "sql/create-table.sql"
func (util *TableUtils) CreateTables() {
	sql := readSqlFile("sql/create-tables.sql")

	_, err := util.pool.Exec(util.ctx, sql)
	if err != nil {
		log.Fatalln("Error execute. Error: ", err)
		return
	}

	log.Infoln("Successfully create tables")
}

// UploadDevData Загрузка тестовых данных в таблицы
func (util *TableUtils) UploadDevData() {
	sql := readSqlFile("sql/dev/create-dev-data.sql")

	_, err := util.pool.Exec(util.ctx, sql)
	if err != nil {
		log.Infoln("Error execute. Error: ", err)
		return
	}

	log.Infoln("Successfully upload dev data")
}

// DropTables Удаляет таблицы
func (util *TableUtils) DropTables() {
	sql := readSqlFile("sql/dev/drop-dev-data.sql")

	_, err := util.pool.Exec(util.ctx, sql)
	if err != nil {
		log.Infoln("Error execute. Error: ", err)
		return
	}

	log.Infoln("Successfully drop tables")
}

// Метод для чтения файла с sql запросами
func readSqlFile(sqlFilePath string) string {
	file, err := os.Open(sqlFilePath)
	if err != nil {
		log.Fatalln("Error open file. Error: ", err)
	}

	sql, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalln("Error read file. Error: ", err)
	}

	defer file.Close()

	return string(sql)
}
