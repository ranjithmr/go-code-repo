package leftrotate

import m "reverse/reverse"


func LeftRot(s []int, n int)[]int {
      a := m.Reverse(s[:n])
     b := m.Reverse(s[n:])
     c := append(a, b...)
      return m.Reverse(c)
}
