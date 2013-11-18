package main

import (
  "flag"
  "fmt"
  "net"
  "os"
)

func getAddresses() (addrs []net.Addr) {
  addrs, err := net.InterfaceAddrs()
  if err != nil {
    os.Exit(1)
  }
  return
}

func succeed(addr string) {
  fmt.Println(addr)
  os.Exit(0)
}

func main() {
  only6 := flag.Bool("6", false, "Limit to IPv6")
  flag.Parse()

  addrs := getAddresses()
  for _, addr := range addrs {
    network, ok := addr.(*net.IPNet)
    if !ok || network.IP.IsLoopback() {
      continue
    }

    // skip IPv6 unless we want it
    is6 := network.IP.To4() == nil
    if (!*only6 && is6) || (*only6 && !is6) {
      continue
    }

    // skip Docker address
    if network.Contains(net.IPv4(172, 17, 0, 1)) {
      continue
    }

    succeed(network.IP.String())
  }

  // print the IPv6 or Docker address if it's the only one
  if len(addrs) > 0 {
    succeed(addrs[0].String())
  }

  os.Exit(1)
}
