// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package neverror

func PanicError(err error) {
	if err != nil {
		panic(err)
	}
}

func HandleError(err *error) {
	if r := recover(); r != nil {
		if e, ok := r.(error); ok {
			*err = e
		} else {
			panic(r)
		}
	}
}
