use "debug"
actor Main
  new create(env: Env) =>
    env.out.print("2019 in Pony.. Day 1")
    env.out.print("====================")
    env.input(recover LineNotify(A) end,512)

class A 

  fun ref apply(s: String) =>
    Debug.out(s)
  
  fun dispose() =>
    None

  