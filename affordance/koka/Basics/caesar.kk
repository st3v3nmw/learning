fun encode(s: string, shift: int)
  s.map( fn(c)
    if c < 'a' || c > 'z' then return c
    val base = (c - 'a').int
    val rot = (base + shift) % 26
    (rot.char + 'a')
  )

fun caesar(s: string): string
  s.encode(-3)

fun main()
  println(caesar("this is an example of speech synthesis in english"))
