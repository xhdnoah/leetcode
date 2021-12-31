package main

import "strings"

// Input: "/home/" Output: "/home"; Input: "/../" Output: "/"
// Input: "/home//foo/" Output: "/home/foo"; Input: "/a/./b/../../c/" Output: "/c"
func simplifyPath(path string) string {
	arr := strings.Split(path, "/")
	stack := make([]string, 0)
	for i := 0; i < len(arr); i++ {
		cur := strings.TrimSpace(arr[i])
		if cur == ".." { // 遇到 .. 出栈
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else if cur != "." && len(cur) > 0 { // 非 "." 入栈
			stack = append(stack, arr[i])
		}
	}
	if len(stack) == 0 {
		return "/"
	}
	return "/" + strings.Join(stack, "/")
}

// 状态机
type State int
type CharType int

const (
	CharDir CharType = iota // '/'
	CharDot                 // '.'
	CharPath
)

const (
	StateDir    State = iota
	StateDot          // '.'
	StateDotDot       // '..'
	StateName         // 'filename'
)

// 当前状态-当前字符-下个状态 映射关系
var transfer = map[State]map[CharType]State{
	StateDir: {
		CharDir:  StateDir,
		CharDot:  StateDot,
		CharPath: StateName,
	},
	StateDot: {
		CharDir:  StateDir,
		CharDot:  StateDotDot, // 两个'.'得到'..'
		CharPath: StateName,
	},
	StateDotDot: {
		CharDir:  StateDir,
		CharDot:  StateName,
		CharPath: StateName,
	},
	StateName: {
		CharDir:  StateDir,
		CharDot:  StateName,
		CharPath: StateName,
	},
}

func charType(c rune) CharType {
	switch c {
	case '.':
		return CharDot
	case '/':
		return CharDir
	}
	return CharPath
}

func simplifyPath_stateMachine(path string) string {
	paths := []string{}
	state, idx := StateDir, 0

	// 根据当前状态和下个状态给出字符串操作
	action := map[State]map[State]func(i int){
		StateDir: {
			StateName: func(i int) { idx = i },
			StateDot:  func(i int) { idx = i },
		},
		StateName: {
			StateDir: func(i int) { paths = append(paths, path[idx:i]) },
		},
		StateDot: {},
		StateDotDot: {
			StateDir: func(i int) {
				if len(paths) > 0 {
					paths = paths[:len(paths)-1]
				}
			},
		},
	}

	if path[len(path)-1] != '/' {
		path += "/"
	}
	for i, c := range path {
		newState := transfer[state][charType(c)]
		if f, ok := action[state][newState]; ok {
			f(i)
		}
		state = newState
	}

	return "/" + strings.Join(paths, "/")
}
