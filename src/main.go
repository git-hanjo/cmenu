package main
import ("fmt";"strconv")

type tMnuStruct struct{
	cmd    string
	note   string
	group  int
     }
type  tslCmd []tMnuStruct
var   slCmd  tslCmd

func (sl tslCmd) listMenu(){
  len1 := 0
  len2 := 0
  for i := 0; i<len(sl);i++{
    if len1 < len(sl[i].cmd)  { len1 = len(sl[i].cmd) }
    if len2 < len(sl[i].note) { len2 = len(sl[i].note) }
  }
  for i := 0; i<10; i++ {
//    if len(slCmd[i]) > nil {
      fmt.Printf("%3d %-"+strconv.FormatInt(int64(len1),10)+"s %-"+strconv.FormatInt(int64(len2),10)+"s %2d\n", i, sl[i].cmd, sl[i].note, sl[i].group)
//      fmt.Printf("%v\n", sl[i])
//    }
  }
}

func init(){
  slCmd = make(tslCmd, 10)
  i:=0; slCmd[i]=tMnuStruct{"git push -u origin master"                   ,"Repo pushen"                ,0}
  i++;  slCmd[i]=tMnuStruct{"git clone git@github.com:git-hanjo/cmenu.git","Repo vom Server holen"      ,0}
  i++;  slCmd[i]=tMnuStruct{"$GOPATH=pwd"                                 ,"GOPATH für currentProj"     ,0}
  i++;  slCmd[i]=tMnuStruct{"alias m=`. ./m`"                             ,"Alias für calAlias für call",0}
}

func main() {
  fmt.Printf("%s\n","----------------- cMenu")
  slCmd.listMenu()
}
