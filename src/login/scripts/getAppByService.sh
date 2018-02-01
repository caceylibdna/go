#!/bin/sh
dpoutput=`cat /root/demo/duster/src/login/tmp/$1_win32services.txt|grep NormalizeService |sed -rn "s/.Bin.bms.BmsService.exe//gp"`
uxoutput=`cat /root/demo/duster/src/login/tmp/$1_win32services.txt|grep UserConsoleService |sed -rn "s/.bin.BDNA.NormalizeBI.Service.exe//gp"`
dpDir=`echo $dpoutput|sed -rn "s/.*\"(.*)\".*/\1/gp"`
dpStatus=`echo $dpoutput|sed -rn "s/.*\"\|(.*)/\1/gp"`
uxDir=`echo $uxoutput|sed -rn "s/.*\"(.*)\".*/\1/gp"`
uxStatus=`echo $uxoutput|sed -rn "s/.*\"\|(.*)/\1/gp"`

dpinfo=`cat /root/demo/duster/src/login/tmp/$1_dpconf |sed -rn "s/.*<NDB_DBConnection.*<ConnectionType>(.*)<\/ConnectionType>.*<Host>(.*)<\/Host>\s*<User>(.*)<\/User>.*<Password>(.*)<\/Password>.*<PDB_DBConnection.*<ConnectionType>(.*)<\/ConnectionType>.*<Host>(.*)<\/Host>\s*<User>(.*)<\/User>.*<Password>(.*)<\/Passwor.*<\/PDB_DBConnection>.*<Version>(.*)<\/Version>.*/\"version\": \"\9\", \"ndb_type\": \"\1\", \"ndb_host\": \"\2\", \"ndb_user\": \"\3\"，\"ndb_passwd\": \"\4\", \"pdb_type\": \"\5\", \"pdb_host\": \"\6\", \"pdb_user\": \"\7\", \"pdb_passwd\": \"\8\"/gp"`

uiinfo=`cat /root/demo/duster/src/login/tmp/$1_dpconf |sed -rn "s/.*<string>BI_CONFIG_BI_USER<\/string>.*<string>(.*)<\/string>.*<string>BI_CONFIG_BI_PASSWORD<\/string>.*<string>(.*)<\/string>.*<string>BI_CONFIG_HAS_ANALYZE.*/\"ui_user\": \"\1\", \"ui_passwd\": \"\2\"/gp"`

uxinfo=`cat /root/demo/duster/src/login/tmp/$1_uxconf| sed -rn "s/.*<Name>Publish<\/Name>.*<ConnectionType>(.*)<\/ConnectionType>.*<Host>(.*)<\/Host>.*<User>(.*)<\/User>.*<Password>(.*)<\/Password>.*<\/DBConnectionBase>.*/\"pdb_type\": \"\1\", \"pdb_host\": \"\2\", \"pdb_user\": \"\3\", \"pdb_passwd\": \"\4\"/gp"`
dpinfo="{\"name\": \"BDNA Data Platform\", \"installDir\": \"$dpDir\", \"state\": \"$dpStatus\", $dpinfo, $uiinfo}"
uxinfo="{\"name\": \"BDNA User Console\", \"installDir\": \"$uxDir\", \"state\": \"$uxStatus\", $uxinfo}"

echo "$dpinfo\n$uxinfo"
