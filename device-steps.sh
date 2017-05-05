#!/bin/bash

DIR=/etc/zededa
PROVDIR=/home/nordmark/go-provision
BINDIR=/home/nordmark/go-provision/bin
WAIT=1

while [ $# != 0 ]; do
    if [ "$1" == -w ]; then
	WAIT=0
    else
	DIR=$1
    fi
    shift
done

echo "Configuration from factory/install:"
(cd $DIR; ls -l)
echo

if [ ! \( -f $DIR/device.cert.pem -a -f $DIR/device.key.pem \) ]; then
    echo "Generating a device key pair and self-signed cert (using TPM/TEE if available)"
    $PROVDIR/generate-dc.sh $DIR/device
    SELF_REGISTER=1
else
    echo "Using existing device key pair and self-signed cert"
    SELF_REGISTER=0
fi
if [ ! -f $DIR/server -o ! -f $DIR/root-certificate.pem ]; then
    echo "No server or root-certificate to connect to. Done"
    exit 0
fi

if [ $WAIT == 1 ]; then
    echo; read -n 1 -s -p "Press any key to continue"; echo; echo
fi

# XXX should we harden/remove any Linux network services at this point?

echo "Check for WiFi config"
if [ -f $DIR/wifi_ssid ]; then
    echo -n "SSID: "
    cat $DIR/wifi_ssid
    if [ -f $DIR/wifi_credentials ]; then
	echo -n "Wifi credentials: "
	cat $DIR/wifi_credentials
    fi
    # XXX actually configure wifi
fi
if [ $WAIT == 1 ]; then
    echo; read -n 1 -s -p "Press any key to continue"; echo; echo
fi

echo "Check for NTP config"
if [ -f $DIR/ntp-server ]; then
    echo -n "Using "
    cat $DIR/ntp-server
    # XXX is ntp service running/installed?
    # XXX actually configure ntp
    # Ubuntu has /usr/bin/timedatectl; ditto Debian
    # ntpdate pool.ntp.org
    # Not installed on Ubuntu
    #
    if [ -f /usr/bin/ntpdate ]; then
	/usr/bin/ntpdate `cat $DIR/ntp-server`
    elif [ -f /usr/bin/timedatectl ]; then
	echo "NTP might already be running. Check"
	/usr/bin/timedatectl status
    else
	echo "NTP not installed. Giving up"
	exit 1
    fi
fi
if [ $WAIT == 1 ]; then
    echo; read -n 1 -s -p "Press any key to continue"; echo; echo
fi

if [ $SELF_REGISTER = 1 ]; then
    echo "Self-registering our device certificate"
    if [ ! \( -f $DIR/prov.cert.pem -a -f $DIR/prov.key.pem \) ]; then
	echo "Missing provisioning certificate. Giving up"
	exit 1
    fi
    $BINDIR/client $DIR selfRegister
    if [ $WAIT == 1 ]; then
	echo; read -n 1 -s -p "Press any key to continue"; echo; echo
    fi
fi

if [ ! -f $DIR/lisp.config ]; then
    echo "Retrieving device and overlay network config"
    $BINDIR/client $DIR lookupParam
    if [ $WAIT == 1 ]; then
	echo; read -n 1 -s -p "Press any key to continue"; echo; echo
    fi
fi

echo "Starting domZ, ZedRouter and overlay network"
# Do something
if [ $WAIT == 1 ]; then
    echo; read -n 1 -s -p "Press any key to continue"; echo; echo
fi

echo "Starting ZedManager"
# Do something
if [ $WAIT == 1 ]; then
    echo; read -n 1 -s -p "Press any key to continue"; echo; echo
fi

echo "Uploading device (hardware) status"
machine=`uname -m`
processor=`uname -p`
platform=`uname -i`
if [ -f /proc/device-tree/compatible ]; then
    compatible=`cat /proc/device-tree/compatible`
else
    compatible=""
fi
memory=`awk '/MemTotal/ {print $2}' /proc/meminfo`
storage=`df -kl --output=size / | tail -n +2| awk '{print $1}'`
cat >$DIR/hwstatus.json <<EOF
{
	"Machine": "$machine",
	"Processor": "$processor",
	"Platform": "$platform",
	"Compatible": "$compatible",
	"Memory": $memory,
	"Storage": $storage
}
EOF
$BINDIR/client $DIR updateHwStatus

if [ $WAIT == 1 ]; then
    echo; read -n 1 -s -p "Press any key to continue"; echo; echo
fi

echo "Uploading software status"
# Only report the Linux info for now
name=`uname -o`
version=`uname -r`
description=`uname -v`
cat >$DIR/swstatus.json <<EOF
{
	"ApplicationStatus": [
		{
			"EID": "::",
			"Name": "$name",
			"Version": "$version",
			"Description": "$description",
			"State": 5,
			"Activated": true
		}
	]
}
EOF
$BINDIR/client $DIR updateSwStatus

echo "Initial setup done!"


