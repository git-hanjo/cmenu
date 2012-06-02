package main
import ("fmt";"strconv")

const ver="0.0.0"
type tMenuItem struct{
	cmd    string
	note   string
	group  int
      }
type  tMenu struct{
        item []tMenuItem
	len  int
      }
var   vMenu  tMenu

func (v *tMenu) listItems(){
  len1 := 0
  len2 := 0
  for i := 0; i<=v.len;i++ {
    if len1 < len(v.item[i].cmd)  { len1 = len(v.item[i].cmd) }
    if len2 < len(v.item[i].note) { len2 = len(v.item[i].note) }
  }
  for i := 0; i<=v.len; i++ {
    fmt.Printf("│%3d %-"+strconv.FormatInt(int64(len1),10)+"s %-"+strconv.FormatInt(int64(len2),10)+"s %2d\n", i, v.item[i].cmd, v.item[i].note, v.item[i].group)
  }
}

func (v *tMenu) init(){
  v.item = make([]tMenuItem, 10)
  i:=0; v.item[i]=tMenuItem{"git push -u origin master"                   ,"Repo auf Server pushen"     ,0}
  i++;  v.item[i]=tMenuItem{"git clone git@github.com:git-hanjo/cmenu.git","Repo vom Server clonen"     ,0}
  i++;  v.item[i]=tMenuItem{"$GOPATH=pwd"                                 ,"GOPATH für currentProj"     ,0}
  i++;  v.item[i]=tMenuItem{"alias m=`. ./m`"                             ,"Alias für calAlias für call",0}
  v.len = i
}
func screenLine(){
  println("         1         2         3         4         5         6         7         8")
  println("12345678901234567890123456789012345678901234567890123456789012345678901234567890")
}

func main() {
  screenLine()
  fmt.Printf("┌───────────────────────────────────────────────────── cMenu (v)%s ───\n",ver)
  vMenu.init()
  vMenu.listItems()
  var i int
  fmt.Scan(&i)
  fmt.Println("read number", i, "from stdin")
}
/*
//─│┌┬┐├┼┤└┴┘
 tMenu
 ├──[]tMenuItem
 │  ├── cmdi
 │  ├── note
 │  └── group
 └── len
 */
