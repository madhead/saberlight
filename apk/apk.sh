#!/usr/bin/env bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
APKTOOL_VERSION=${APKTOOL_VERSION-2.0.3}
DEX2JAR_VERSION=${DEX2JAR_VERSION-2.0}
PROCYON_VERSION=${PROCYON_VERSION-0.5.30}
CFR_VERSION=${CFR_VERSION-0_110}
ANT_VERSION=${ANT_VERSION-1.9.6}

find $DIR ! -path $DIR ! -name "apk.sh" ! -name "README.md" -delete
wget -O $DIR/app.apk http://down.mumayi.com/881461

mkdir $DIR/tools
wget https://bitbucket.org/iBotPeaches/apktool/downloads/apktool_${APKTOOL_VERSION}.jar -O $DIR/tools/apktool_${APKTOOL_VERSION}.jar
wget https://github.com/pxb1988/dex2jar/releases/download/${DEX2JAR_VERSION}/dex-tools-${DEX2JAR_VERSION}.zip -O $DIR/tools/dex-tools-${DEX2JAR_VERSION}.zip
unzip $DIR/tools/dex-tools-${DEX2JAR_VERSION}.zip -d $DIR/tools
chmod +x $DIR/tools/dex2jar-${DEX2JAR_VERSION}/*.sh
wget https://bitbucket.org/mstrobel/procyon/downloads/procyon-decompiler-${PROCYON_VERSION}.jar -O $DIR/tools/procyon-decompiler-${PROCYON_VERSION}.jar
wget http://www.benf.org/other/cfr/cfr_${CFR_VERSION}.jar -O $DIR/tools/cfr_${CFR_VERSION}.jar
wget http://www.us.apache.org/dist/ant/binaries/apache-ant-${ANT_VERSION}-bin.zip -O $DIR/tools/apache-ant-${ANT_VERSION}-bin.zip
unzip $DIR/tools/apache-ant-${ANT_VERSION}-bin.zip -d $DIR/tools
chmod +x $DIR/tools/apache-ant-${ANT_VERSION}/bin/*.sh
git clone https://github.com/fesh0r/fernflower.git $DIR/tools/fernflower
$DIR/tools/apache-ant-${ANT_VERSION}/bin/*.sh
$DIR/tools/apache-ant-${ANT_VERSION}/bin/ant -f $DIR/tools/fernflower/build.xml

mkdir $DIR/interim
java -jar $DIR/tools/apktool_${APKTOOL_VERSION}.jar -o $DIR/interim/apktool decode $DIR/app.apk
$DIR/tools/dex2jar-${DEX2JAR_VERSION}/d2j-dex2jar.sh -o $DIR/interim/app.jar $DIR/app.apk

mkdir $DIR/src
mkdir $DIR/src/procyon
java -jar $DIR/tools/procyon-decompiler-${PROCYON_VERSION}.jar -o $DIR/src/procyon $DIR/interim/app.jar
mkdir $DIR/src/cfr
java -jar $DIR/tools/cfr_${CFR_VERSION}.jar --outputdir $DIR/src/cfr $DIR/interim/app.jar
mkdir $DIR/src/fernflower
java -jar $DIR/tools/fernflower/fernflower.jar $DIR/interim/app.jar $DIR/src/fernflower
unzip $DIR/src/fernflower/app.jar -d $DIR/src/fernflower
