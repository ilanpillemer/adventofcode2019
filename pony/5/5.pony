use "debug"

actor Main
  new create(env: Env) =>
    env.input(recover LineNotify(A) end,512)

class A
  fun ref apply(s: String) =>
    Debug.out(s)

  fun dispose() =>
    None
  