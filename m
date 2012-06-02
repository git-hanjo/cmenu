#!/bin/sh

cmd_1='git push -u origin'
cmd_2='git clone git@github.com:git-hanjo/p001.git'
cmd_3='$GOPATH=$GOPATH'
cmd_4='alias m=`. ./m`'

show_cmd(){
  i=0
  while [ $((i=i+1)) -le 10 ]; do
#    echo $i
    tmp="echo $i \$cmd_"$i
    tmp="\$cmd_"$i
#    echo $tmp
    if [ $i -le 9 ];      then echo -n "  "
      else echo -n " "
    fi
    echo $tmp
    #eval $tmp
    echo ${#tmp}
    echo eval $tmp
  done
}
echo "\n\n\n"
echo "-------------------------------------------"
echo "WICHTIG:"
echo ". ./r.sh"
export GOPATH=`pwd`
echo "-------------------------------------------"
echo "git push -u origin"
echo "git clone git@github.com:git-hanjo/p001.git"
echo "-------------------------------------------"

show_cmd
