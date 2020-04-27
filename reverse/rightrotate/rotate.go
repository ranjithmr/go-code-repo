package rightrotate

import m "reverse/reverse"


func RightRot(s []int, n int)[]int {
     ri := m.Reverse(s)
      a := m.Reverse(ri[:n])
     b := m.Reverse(ri[n:])
     c := append(a, b...)
      return c
}
