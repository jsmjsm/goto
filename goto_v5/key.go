/*
 * @Description  :
 * @Author       : jsmjsm
 * @Github       : https://github.com/jsmjsm
 * @Date         : 2021-03-09 11:32:36
 * @LastEditors  : jsmjsm
 * @LastEditTime : 2021-03-09 11:38:24
 * @FilePath     : /goto/goto_v5/key.go
 */
package main

var keyChar = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func genKey(n int) string {
	if n == 0 {
		return string(keyChar[0])
	}
	l := len(keyChar)
	s := make([]byte, 20) // FIXME: will overflow. eventually.
	i := len(s)
	for n > 0 && i >= 0 {
		i--
		j := n % l
		n = (n - j) / l
		s[i] = keyChar[j]
	}
	return string(s[i:])
}
