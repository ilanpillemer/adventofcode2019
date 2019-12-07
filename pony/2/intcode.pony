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
    let mem : Array[USize] iso =
     recover
       var mem': Array[USize] = Array[USize]
       Iter[String](s.split(",").values())
         .map[USize](C~usize())
	 .map_stateful[None]({(i) => mem'.push(i)  } )
	 .run()
      mem'(1)? = 12
      mem'(2)? = 2
      mem'
     end

    let m: Machine = Machine(consume mem)
    m.exec()
  end
      

  fun dispose() =>
    None

primitive C
  fun usize(s: String): USize ?  =>
   s.usize()?

 




















