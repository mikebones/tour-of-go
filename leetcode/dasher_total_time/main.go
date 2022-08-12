package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type DashTime struct {
	room string
	timestamp int
	function string
}

type DashTimeStack []DashTime 

func (dts *DashTimeStack) push(dt DashTime) {
	*dts = append(*dts, dt)
}

func (dts *DashTimeStack) pop()(DashTime, bool) {
	var dashTime DashTime 
	if !dts.isStackEmpty() {
		dashTime = (*dts)[len(*dts)-1]
		*dts = (*dts)[:len(*dts)-1]
		return dashTime, true
	} else {
		return dashTime, false
	}
}

func (dts *DashTimeStack) isStackEmpty() bool {
	return len(*dts) == 0 
}

func getDashTime(filePath string) {
	var dts DashTimeStack
	roomTime := make(map[string]int)
	file, err := os.Open(filePath)	
	check(err)
	scanner := bufio.NewScanner(file)
	var previousTime int
	for scanner.Scan() {
		
		values := strings.Split(scanner.Text(),",")
		var dt DashTime 
		dt.room = values[0]
		dt.timestamp, err = strconv.Atoi(values[1])
		dt.function = values[2]
		if dt.function == "Enter" {
			roomTime[dt.room] = roomTime[dt.room] + (dt.timestamp - previousTime)
			dts.push(dt)
		}
		if dt.function == "Exit" {
			if elem, ok := dts.pop(); ok {
				roomTime[dt.room] = roomTime[dt.room] + (dt.timestamp - elem.timestamp)
			}
		}
		previousTime = dt.timestamp
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}


