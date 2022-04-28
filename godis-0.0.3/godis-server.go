package main

import (
	"fmt"
	"godis/core"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//服务端实例
var godis new
