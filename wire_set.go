//go:build wireinject
// +build wireinject

package main

import (
	"fmt"

	"github.com/google/wire"
)

func InitApp() *App {
	panic(wire.Build(wire.NewSet(
		// ---- 切换 MySQL 源 START -----
		// GetMySQLConfig,
		// wire.Bind(new(DataSource), new(*MySQL)),
		// NewMySQL,
		// ---- 切换 MySQL 源  END  -----

		// ---- 切换 Redis 源 START -----
		GetRedisConf,
		wire.Bind(new(DataSource), new(*Redis)),
		NewRedis,
		// ---- 切换 Redis 源  END  -----
		NewApp,
	)))
}

// 封装MySQL Client
type MySQLClient struct {
}

func NewMySQLClient(conf *MySQLConfig) *MySQLClient {
	return &MySQLClient{}
}

func (db *MySQLClient) Exec(id string) string {
	return fmt.Sprintf("MySQL: some data %v", id)
}

type MySQLConfig struct {
}

func GetMySQLConfig() *MySQLConfig {
	return &MySQLConfig{}
}

// 实现DataSource接口
type MySQL struct {
	m *MySQLClient
}

func NewMySQL(conf *MySQLConfig) *MySQL {
	return &MySQL{m: NewMySQLClient(conf)}
}

func (m *MySQL) GetById(id string) string {
	return m.m.Exec(id)
}

// 封装Reids Client
type RedisClient struct {
}

func NewRedisClient(conf *RedisConfig) *RedisClient {
	return &RedisClient{}
}

func (db *RedisClient) Do(id string) string {
	return fmt.Sprintf("Redis: some data %v", id)
}

type RedisConfig struct {
}

func GetRedisConf() *RedisConfig {
	return &RedisConfig{}
}

// 实现DataSource接口
type Redis struct {
	r *RedisClient
}

func NewRedis(conf *RedisConfig) *Redis {
	return &Redis{r: NewRedisClient(conf)}
}

func (r *Redis) GetById(id string) string {
	return r.r.Do(id)
}

// 抽象出 DataSource
type DataSource interface {
	GetById(id string) string
}

// ---------------
func main() {
	app := InitApp()
	app.Run()
}

type App struct {
	ds DataSource
}

func NewApp(ds DataSource) *App {
	return &App{ds: ds}
}

func (a *App) Run() {
	a.GetData("1")
}

func (a *App) GetData(id string) string {
	data := a.ds.GetById(id)
	fmt.Println(data)
	return data
}
