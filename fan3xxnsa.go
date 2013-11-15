package main

// this will only work when 1 is sent to pwr_enable

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	// path to the sys i2c interface (can be different on various kernel versions)
	IDIR = "/sys/class/i2c-dev/i2c-0/device/0-002e"

	// marginal temperatures and pwm values (at idle 42 board 48 cpu when case is opened)
	LOWTEMP  = 42
	HIGHTEMP = 62
	PWMMIN   = 155
	PWMMAX   = 255
)

func ReadTemp() byte {
	f, err := os.OpenFile(IDIR+"/temp1_input", os.O_RDONLY, 0)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	r := bufio.NewReader(f)
	tempstr, err := r.ReadString(10)
	if err != nil {
		log.Fatal(err)
	}
	temp_i, err := strconv.ParseInt(strings.TrimSpace(tempstr), 10, 32)
	if err != nil {
		log.Fatal(err)
	}
	temp := byte(temp_i / 1000)
	return temp
}

func WritePwm(pwm byte) {
	f, err := os.OpenFile(IDIR+"/pwm1", os.O_WRONLY, 0)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	_, err = f.WriteString(fmt.Sprintf("%d\n", pwm))
	if err != nil {
		log.Fatal(err)
	}
}

func CalculatePwm(temp byte) byte {
	if temp <= LOWTEMP {
		return 0
	} else if temp >= HIGHTEMP {
		return PWMMAX
	}
	return PWMMIN + (5 * (temp - LOWTEMP))
}

func PrintTemp() {
	fmt.Printf("Temp: %d\n", ReadTemp())
}

func PrintPwmValues() {
	for i := 0; i <= 100; i++ {
		fmt.Printf("Temp: %d Pwm: %d\n", i, CalculatePwm(byte(i)))
	}
}

func main() {
	for {
		WritePwm(CalculatePwm(ReadTemp()))
		time.Sleep(120 * time.Second)
	}
}
