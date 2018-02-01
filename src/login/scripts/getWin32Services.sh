#!/bin/sh
echo "aaaaa"
`wmic -U $1%$2 //$3 "Select PathName,state from Win32_Service" >./tmp/$3_win32services.txt`
echo "bbbbbb"
`winexe -U $1%$2 //$3 'cmd /C type "C:\BDNA\Data Platform\Conf\Norm.Configuration.config"' >./tmp/$3_dpconf1`
echo "cccccc"
`winexe -U $1%$2 //$3 'cmd /C type "C:\BDNA\User Console\Conf\NBI.Configuration.config"' >./tmp/$3_uxconf1`
echo "dddddd"
`vi -e -s -c "%s/\r//g" -c "wq" ./tmp/$3_dpconf`
`vi -e -s -c "%s/\r//g" -c "wq" ./tmp/$3_uxconf`
echo "tets"
`awk 'BEGIN{ORS=" "}{print $0}' ./tmp/$3_dpconf1 >./tmp/$3_dpconf`
echo "test2"
`awk 'BEGIN{ORS=" "}{print $0}' ./tmp/$3_uxconf1 >./tmp/$3_uxconf`
`rm -rf ./tmp/$3_dpconf1 ./tmp/$3_uxconf1`
