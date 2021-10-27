package models

import (
	"database/sql"
	"errors"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/prometheus/common/log"
	"strconv"
)

var Db *sql.DB

type Status int8

const (
	Init    Status = 0
	OK      Status = 1
	Pending Status = 2
	Close   Status = 3
	Dispute Status = 4
)

type Payment struct {
	Id     int
	Amount int    `json:"amount"`
	Expiry int    `json:"expiry"`
	From   string `json:"from"`
	To     string `json:"to"`
}
type Channel struct {
	ChannelAddress string              `json:"channel_address"`
	PlayerCount    int                 `json:"player_count"`
	BestRound      int                 `json:"best_round"`
	Status         Status              `json:"status"`
	OpenTime       timestamp.Timestamp `json:"open_time"`
	Deadline       int                 `json:"deadline"`
}
type Player struct {
	Uid        int    `json:"Uid"`
	Addr       string `json:"Addr"`
	Credit     int    `json:"Credit"`
	Withdrawal int    `json:"Withdrawal"`
	Withdrawn  int    `json:"Withdrawn"`
	Deposit    int    `json:"Deposit"`
}

func InitDB() {
	dbconn, _ := beego.AppConfig.String("DBConn")
	db, err := sql.Open("mysql", dbconn)
	if err != nil {
		fmt.Println(err)
		return

	}
	db.SetMaxOpenConns(2000)
	db.SetMaxIdleConns(0)
	db.Ping()
	Db = db
}
func CLose() {
	if Db != nil {
		Db.Close()
	}
}
func AddPayment(pay Payment) error {
	var isql = "INSERT payment SET account_id=?,partner_id=?,user_id=?,amount=?,outer_tradeno=?,remark=?"
	if Db == nil {
		return errors.New("AddPaymenRec connect mysql failed")
	}
	stmt, _ := Db.Prepare(isql)
	defer stmt.Close()
	log.Info("Addpayment rec=%#v", pay)

	res, err := stmt.Exec(pay.From, pay.To, pay.Amount, pay.Expiry)
	if err != nil {
		log.Info("res-%#v", res)
		return errors.New("insert payment failed\r\n")
	}
	return nil
}

func GetPayment(payId int) (Payment, error) {
	var qsql = "SELECT * FROM payment WHERE  account_id=?"
	stmt, _ := Db.Prepare(qsql)
	rows, err := stmt.Query(payId)
	var response Payment
	defer rows.Close()
	if err != nil {
		return Payment{}, errors.New("Query payment failed\r\n")
	}
	for rows.Next() {
		err = rows.Scan(&response.Id, &response.Amount, &response.Expiry, &response.From, &response.To)
		if err != nil {
			return response, errors.New("read payment failed \r\n")
		}
		return response, err
	}
	return Payment{}, errors.New("null data")
}

func AddPlayer(player Player, channelId string) (int, error) {
	var isql = "INSERT INTO " + channelId + "(uid, addr, credit, withdrawn, deposit) VALUES( ?, ?, ?, ?, ? )"
	if Db == nil {
		return 0, errors.New("AddPlayer connect mysql failed")
	}
	stmt, err := Db.Prepare(isql)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	log.Info("AddPLayer rec=%#v", player)

	res, err := stmt.Exec(nil, player.Addr, player.Credit, player.Withdrawn, player.Deposit)
	if err != nil {
		log.Info("res-%#v", res, err)
		return 0, err
	}
	id, _ := res.LastInsertId()
	return int(id), nil

}
func GetPlayerById(id int, channelId string) (Player, error) {
	var qsql = "SELECT * FROM CHANNEL. "+channelId+" where uid=?"
	stmt, _ := Db.Prepare(qsql)
	rows, err := stmt.Query(id)
	var response Player
	defer stmt.Close()
	if err != nil {
		return Player{}, errors.New("Query payment failed\r\n")
	}
	for rows.Next() {
		err = rows.Scan(&response.Uid, &response.Addr, &response.Credit, &response.Withdrawn, &response.Deposit)
		if err != nil {
			return response, errors.New("read payment failed \r\n")
		}
		return response, err
	}
	return Player{}, errors.New("null data")
}

func GetPlayerByAddr(addr string, channelId string) (Player, error) {
	var qsql = "SELECT * FROM CHANNEL." + channelId + " where addr=?"
	stmt, _ := Db.Prepare(qsql)
	rows, err := stmt.Query(addr)
	var response Player
	defer rows.Close()
	if err != nil {
		return Player{}, errors.New("Query payment failed\r\n")
	}
	for rows.Next() {
		err = rows.Scan(&response.Uid, &response.Addr, &response.Credit, &response.Withdrawn, &response.Deposit)
		if err != nil {
			return response, errors.New("read payment failed \r\n")
		}
		return response, err
	}
	return Player{}, errors.New("null data")

}

//each table for each channel
func CreateChannel(channelId string) error {
	sql := `CREATE TABLE ` + channelId + `(
		uid int unsigned NOT NULL AUTO_INCREMENT,
  		addr text,
  		credit int DEFAULT NULL,
  		withdrawn int DEFAULT NULL,
  		deposit int DEFAULT NULL,
  		PRIMARY KEY (uid)
	);`
	fmt.Println("\n" + sql + "\n")
	smt, err := Db.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = smt.Exec()
	if err != nil {
		return err
	}
	return nil
}

func UpdatePLayer(player Player, channelId string) error {
	playID := strconv.Itoa(player.Uid)
	var isql = "UPDATE " + channelId + " SET credit=?, withdrawn=?, deposit=? where uid=" + playID
	if Db == nil {
		return errors.New("Null DB")
	}
	stmt, err := Db.Prepare(isql)
	if err != nil {
		return err
	}
	defer stmt.Close()
	log.Info("UpdatePLayer rec=%#v", player)

	res, err := stmt.Exec(player.Credit, player.Withdrawn, player.Deposit)
	if err != nil {
		log.Info("res-%#v", res, err)
		return err
	}
	return nil
}

func DeletePlayer(playerId int, channelId string) error {
	sql := "DELETE FROM " + channelId + " WHERE uid = ?"
	stmtOut, err := Db.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmtOut.Close()

	result, err := stmtOut.Exec(playerId)
	if err != nil {
		return err
	}
	if rowNum, err := result.RowsAffected(); err != nil || rowNum != int64(1) {
		return errors.New("delete error")
	}
	return nil
}
