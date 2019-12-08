use "debug"
use "itertools"
use "collections"

actor Main
  new create(env: Env) =>
    env.out.print("Intcode initialising...")
    env.input(recover LineNotify(A) end, 512)
    env.out.print("Intcode shutting down...")

class A
  fun ref apply(s: String) =>
    var mem: Mem = Mem
    Iter[String](s.split(",").values())
     .map[USize](C~usize())
     .map_stateful[None]({(i) => mem.push(i)  } )
     .run()

   let m = to_iso(mem.clone())
   let m' = to_iso(mem.clone())   

   let machine: Machine = Machine(consume m)
   let machine': Machine = Machine(consume m')

   machine.set_noun(12)
   machine.set_verb(2)

   machine'.set_noun(23)
   machine'.set_verb(47)

   machine.exec()
   machine'.exec()

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

 




















