use "debug"
use "itertools"

actor Main
  new create(env: Env) =>
    env.out.print("Intcode initialising...")
    env.input(recover LineNotify(A) end, 512)
    env.out.print("Intcode shutting down...")

class A
  fun ref apply(s: String) =>
  try
    var mem': Mem = Mem
    Iter[String](s.split(",").values())
     .map[USize](C~usize())
     .map_stateful[None]({(i) => mem'.push(i)  } )
     .run()
    let mem'' = mem'.clone()
    
    mem'(1)? = 12
    mem'(2)? = 2

    mem''(1)? = 23
    mem''(2)? = 47

   let mem: Array[Mem] = Array[Mem]
   mem.push(mem')
   mem.push(mem'')

   let m = to_iso(mem(0)?)
   let m' = to_iso(mem(1)?)   

   let machine: Machine = Machine(consume m)
   let machine2: Machine = Machine(consume m')
   machine.exec()
   machine2.exec()
  end
      

  fun dispose() =>
    None

  fun to_iso(mem: Mem ref): Mem iso^ =>
    let mem': Mem iso = recover Mem end
    for m in mem.values() do
      mem'.push(m)
    end
    consume mem'

primitive C
  fun usize(s: String): USize ?  =>
   s.usize()?

 




















