#!/bin/bash

set -e pipefail

BRIDGE=br0                # name for the bridge we will create to share network with the raspbian img
BINARY_PATH=/usr/bin      # path prefix for binaries
LOCK_FILE="/.net_prepared"

if [[ -f $LOCK_FILE ]]; then
  echo "network already configured, skipping"
  exit 0
fi

IFACE="$( ip r | grep "default via" | awk '{ print $5 }' | head -1 )"

if [[ "$IFACE" == "" ]] || [[ "$BRIDGE" == "" ]]; then
  echo "failed getting iface name"
  exit 1
fi

IP=$( ip address show dev "$IFACE" | grep global | grep -oP '\d{1,3}(.\d{1,3}){3}' | head -1 )
if [[ "$IP" == "" ]]; then
  echo "no IP found for $IFACE";
  exit 1
fi

type brctl &>/dev/null || {
  echo "brctl is not installed"
  exit 1
}

modprobe tun &>/dev/null
grep -q tun <(lsmod) || {
  echo "need tun module"
  exit 1
}

# network configuration
cat > /etc/qemu-ifup <<EOF
#!/bin/sh
echo "Executing /etc/qemu-ifup"
echo "Bringing up \$1 for bridged mode..."
ip link set \$1 up promisc on
echo "Adding \$1 to $BRIDGE..."
brctl addif $BRIDGE \$1
sleep 2
EOF

cat > /etc/qemu-ifdown <<EOF
#!/bin/sh
echo "Executing /etc/qemu-ifdown"
ip link set \$1 down
brctl delif $BRIDGE \$1
ip link delete dev \$1
EOF

chmod 750 /etc/qemu-ifdown /etc/qemu-ifup

IPFW=$( sysctl net.ipv4.ip_forward | cut -d= -f2 )
sysctl net.ipv4.ip_forward=1

echo "Getting routes for interface: $IFACE"
ROUTES=$( ip route | grep $IFACE )

echo "Changing those routes to bridge interface: $BRIDGE"
BRROUT=$( echo "$ROUTES" | sed "s=$IFACE=$BRIDGE=" )

echo "Creating new bridge: $BRIDGE"
brctl addbr $BRIDGE

echo "Adding $IFACE interface to bridge $BRIDGE"
brctl addif $BRIDGE $IFACE

echo "Setting link up for: $BRIDGE"
ip link set up dev $BRIDGE

echo "Flusing routes to interface: $IFACE"
ip route flush dev $IFACE

echo "Adding IP address to bridge: $BRIDGE"
ip address add $IP dev $BRIDGE

echo "Adding routes to bridge: $BRIDGE"
echo "$BRROUT" | tac | while read l; do ip route add $l; done
echo "Routes to bridge $BRIDGE added"

precreationg=$(ip tuntap list | cut -d: -f1 | sort)
ip tuntap add user root mode tap
postcreation=$(ip tuntap list | cut -d: -f1 | sort)
TAPIF=$(comm -13 <(echo "$precreationg") <(echo "$postcreation"))
NET_ARGS="-net nic -net tap,ifname=$TAPIF"

touch $LOCK_FILE
