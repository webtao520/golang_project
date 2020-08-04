package models

import (
	"time"

	"github.com/astaxie/beego/logs"
)

/* type SecKillActivity struct {
	ActivityId   	int 		活动id
	ActivityName 	string		活动名称
	ProductId    	int			商品id
	StartTime    	int			活动开始时间
	EndTime      	int			活动结束时间
	Total        	int			活动商品数量
	Status       	int			活动状态（1.未开始，2.进行中，3.已结束）
	SecondLimit  	int			每秒可售商品个数
	EveryoneLimit   int			每人限购商品数
	BuyRate      	int			购买到的概率 （百分比）
}

create table `sec_kill_activity`
    -- --------------------------------------------------
    --  Table Structure for `newSecKill/admin/models.SecKillActivity`
    -- --------------------------------------------------
    CREATE TABLE IF NOT EXISTS `sec_kill_activity` (
        `activity_id` integer AUTO_INCREMENT NOT NULL PRIMARY KEY,
        `activity_name` varchar(60) NOT NULL DEFAULT ''  UNIQUE,
        `product_id` integer NOT NULL DEFAULT 0 ,
        `start_time` datetime NOT NULL,
        `end_time` datetime NOT NULL,
        `total` integer NOT NULL DEFAULT 0 ,
        `status` integer NOT NULL DEFAULT 2 ,
        `second_limit` integer NOT NULL DEFAULT 0 ,
        `everyone_limit` integer NOT NULL DEFAULT 0 ,
        `buy_rate` integer NOT NULL DEFAULT 0
    ) ENGINE=MyISAM;
*/

type SecKillActivity struct {
	ActivityId    int       `orm:"pk;auto"`
	ActivityName  string    `orm:"unique;size(60);default('')" form:"ActivityName"  valid:"Required"`
	ProductId     int       `orm:"size(10);default(0)" form:"ProductId" valid:"Min(0)"`
	StartTime     time.Time `orm:"size(10);default(0)" form:"StartTime" valid:"Min(0)"`
	EndTime       time.Time `orm:"size(10);default(0)" form:"EndTime" valid:"Min(0)"`
	Total         int       `orm:"size(10);default(0)" form:"Total" valid:"Min(0)"`
	Status        int       `orm:"default(2)" form:"Status" valid:"Range(1,3)"`
	SecondLimit   int       `orm:"size(10);default(0)" form:"SecondLimit" valid:"Min(0)"`
	EveryoneLimit int       `orm:"size(10);default(0)" form:"EveryoneLimit" valid:"Min(0)"`
	BuyRate       int       `orm:"size(10);default(0)" form:"BuyRate" valid:"Min(0)"`
}

// 实例化模型
func NewActivityModel() *SecKillActivity {
	return &SecKillActivity{}
}

// 保存活动数据
func (this *SecKillActivity) InsertActivity(activity *SecKillActivity) (num int64, err error) {
	num, err = DB.Insert(activity)
	if err != nil {
		logs.Warn("add activity err : %v", err)
		return
	}
	err = this.syncActivityToEtcd("add", *activity)
	if err != nil {
		return
	}
	return
}

func (this *SecKillActivity) GetActivityList(slimit, elimit int) (lists []SecKillActivity, err error) {
	if slimit == 0 && elimit == 0 {
		slimit = 0
		elimit = 20
	}
	num, err := DB.Raw("select * from sec_kill_activity limit ?,?", slimit, elimit).QueryRows(&lists)
	if err != nil && num < 0 {
		logs.Warn("get Activity list err : %v ", err)
		return
	}
	return
}

func (this *SecKillActivity) GetActivityById(ActivityId int) (activity SecKillActivity, err error) {
	err = DB.Raw("select * from sec_kill_activity where activity_id = ? limit 1", ActivityId).QueryRow(&activity) // 取出一条结果集
	if err != nil {
		logs.Warn("get Activity by activity_id failed : activity_id = %d , err : %v ", ActivityId, err)
		return
	}
	return
}

func (this *SecKillActivity) UpdateActivity(Activity *SecKillActivity) (num int64, err error) {
	if num, err = DB.Update(Activity); err != nil {
		logs.Warn("update activity err : %v", err)
		return
	}
	err = this.syncActivityToEtcd("update", *Activity)
	if err != nil {
		return
	}
	return
}

// 删除活动
func (this *SecKillActivity) DelActivity(Activity *SecKillActivity) (num int64, err error) {
	// 先删除etcd
	// 因为 DB.Delete(Activity) 删除成功后 Activity.ActivityId 将变成 0
	err = this.syncActivityToEtcd("del", *Activity)
	if err != nil {
		return
	}

	if num, err = DB.Delete(Activity); err != nil {
		logs.Warn("del Activity err : %v ", err)
		return
	}
	return
}

// // 监听 SecKillActivityList 变化
func (this *SecKillActivity) WatchActivity() {

}
