use "debug"
use "collections"
use "itertools"

actor Main
  new create(env: Env) =>
    env.input(recover LineNotify(A) end, 512)

class A
 fun ref apply(s: String) =>
   Debug(s)

   Iter[U32](s.runes())
    .map[None]({(b) => Debug(String.from_utf32(b)) })
    .run()

  fun dispose() => None 