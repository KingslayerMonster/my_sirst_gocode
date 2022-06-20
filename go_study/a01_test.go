package go_study

import (
	"fmt"
	"testing"
)

//test方式运行函数
func Test_a01_01(t *testing.T) {
	fmt.Println("okokokokok")
	fmt.Println("1234")
}
func Test_a01_02(t *testing.T) {
	var index int = 3
	if index == 1 {
		fmt.Println("赵文杰的第一个go程序")

	}
	if index == 2 {
		fmt.Println("赵文杰的第二个go程序")

	}
	if index == 3 {
		fmt.Println("赵文杰的第三个go程序")

	}
	if index == 4 {
		fmt.Println("赵文杰的第四个go程序")

	}
	if index == 5 {
		fmt.Println("赵文杰的第五个go程序")

	}

}
func Test_a01_03(t *testing.T) {
	var zwj string = "1"
	switch zwj {
	case "1":
		fmt.Println("zwj1")
	case "3":
		fmt.Println("zwj3")
	case "2":
		fmt.Println("zwj2")
	case "4":
		fmt.Println("zw4j")
	default:
		fmt.Println("p dou bu shi")
	}
}

func Test_a01_04(t *testing.T) {
	zhaozhenjie_girlfriends:=[]bool{
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		true,
		false,

	}

	for position,girl:=range zhaozhenjie_girlfriends{
		if girl==true{

			fmt.Printf("卑微的赵文杰追第%d个女娃终于到手，结第三次婚，带四个娃，大概五百斤\n",position)
			break
		}else{
			fmt.Printf("卑微的赵文杰追第%d个女娃失手 \n",position)
		}
	}
}
func Test_a01_05(t *testing.T) {

	zhaozhenjie_of_life := map[int]string{
		12:  "裤衩扯了 f缝裤c产生感情",
		14:  "w王者荣耀y遇到个n女的 j结果s是n男的 s伤心欲绝",
		17:  "t谈了一周 f放屁太臭 n你太嫌弃",
		21:  "因为人家太优秀，没有勇气表白，错过",
		27:  "单身六年，一直用手解决，手速一流",
		28:  "遇到一个挚爱，结果结了三次婚，因了四次哇，由于失败的感情经历，暴饮暴食五百斤",
		30:  "当初最喜欢的女神嫁给一个肥猪一样的富二代 郁闷 伤心欲绝",
		35:  "如狼似虎的年龄已经过去 手速无人能敌 至今还是处男",
		41:  "终于和一个60岁的富婆成婚",
		43:  "富婆陨落归西 你继承财产",
		46:  "至今依然单身",
		47:  "单身",
		50:  "回忆当初当初自己的手速 真是怀念 想当年 老子手速 如今老yi",
		53:  "看到年轻人在比手速 自己不服气 和年轻人在路边一起打起了 飞机 手速过快 但已经远不及当年 赢得了年轻人的尊重 爷爷就是厉害",
		57:  "腰身疼 但是每天还要打一遍飞机 吃两个猪腰子",
		59:  "做生意有点起色 成了一个千万富翁",
		71:  "和一个20岁的小姑娘谈起来恋爱 但是已经 没有欲望了",
		91:  "结婚20年 没有进行过一次 因为你不行了 生了一堆娃 但不是你的",
		101: "一群儿子和媳妇在争夺你的财产中 归西 安静点",
	}

	zhaozhenjie_story := []string{
		"被男人骗了身体",
		"被女人仙人跳",
		"捐精赚了点外快",
		"比赛打飞机 比手速",
		"被甩",
		"渴望得到一个soulmate",
		"打飞机",
		"玩绅士游戏",
		"写爬虫下载小电影 手速上一层楼",
		"手速过快 手骨折 降速",
		"右手由于骨折 换左手",
		"左手骨折 换右手",
		"进入贤者时间",
	}
	for pos := 0; pos <= 102; pos++ {
		if v, ok := zhaozhenjie_of_life[pos]; ok {
			important_thing := fmt.Sprintf("在赵文杰%d岁,主要事件:%s\n", pos, v)
			Red(important_thing)
		} else {
			x:=fmt.Sprintf("在赵文杰%d岁,主要事件:%s\n", pos, zhaozhenjie_story[pos%len(zhaozhenjie_story)])
			Green(x)
		}
	}
}
