package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Options struct {
	HelpFlag   bool
	DropFlag   bool
	ToDrop     []string
	CreateFlag bool
	ToCreate   []string
	SpawnFlag  bool
	StarSteps  int
	TestFlag   bool
}

func ParseOpts() (*Options, error) {
	o := new(Options)
	var state int
	const DFLAG = 1
	const CFLAG = 2
	const SFLAG = 3
	var anyFlag bool
	for _, arg := range os.Args[1:] {
		switch arg {
		case "-h":
			o.HelpFlag = true
			return o, nil
		case "-d":
			anyFlag = true
			state = DFLAG
			o.DropFlag = true
		case "-c":
			anyFlag = true
			state = CFLAG
			o.CreateFlag = true
		case "-s":
			anyFlag = true
			state = SFLAG
			o.SpawnFlag = true
		case "-t":
			anyFlag = true
			o.TestFlag = true
		default:
			switch state {
			case DFLAG:
				o.ToDrop = append(o.ToDrop, arg)
			case CFLAG:
				o.ToCreate = append(o.ToCreate, arg)
			case SFLAG:
				if x, err := strconv.Atoi(arg); err != nil {
					return nil, fmt.Errorf("unparsable starnum %s", arg)
				} else {
					o.StarSteps = x
				}
			default:
				if strings.HasPrefix(arg, "-") {
					for _, char := range arg {
						switch char {
						case '-':
						case 'h':
							anyFlag = true
							o.HelpFlag = true
						case 's':
							anyFlag = true
							o.SpawnFlag = true
						case 'd':
							anyFlag = true
							o.DropFlag = true
						case 't':
							anyFlag = true
							o.TestFlag = true
						case 'c':
							anyFlag = true
							o.CreateFlag = true
						default:
							o.HelpFlag = true
							return o, fmt.Errorf("unknown flag '%s'", string(char))
						}
					}
				} else {
					o.HelpFlag = true
					return o, fmt.Errorf("unknown arg %s", arg)
				}
			}
		}
	}
	if !anyFlag {
		o.HelpFlag = true
		return o, nil
	}
	return o, nil
}

func PrintHelp() {
	fmt.Printf(`
manage [-h][-d [TABLENAME...]][-c [TABLENAME...]][-s [STEPNUM]]
  -h	Print this help message and do nothing else
  -d	Drop tables listed, default list if none listed
  -c	Create tables listed, default list if none listed
  -s	Repopulate star table, regenerating star file if STEPNUM is present and >0
  -t	Test star table size
You may combine flags that have no arguments (e.g. 'manage -csdt')

`)
}
