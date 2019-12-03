use "debug"
use "itertools"
use "regex"
use "collections"

actor Main
  new create(env: Env) =>
    env.out.print("2019 in Pony.. Day 3")
    env.out.print("====================")
    env.input(recover LineNotify(A) end,512)
 
class A

  var counts: Map[String,U32] = Map[String,U32] 
  
  fun ref apply(s: String) =>
  try
    var x:I32 = 0
    var y:I32 = 0

    var set:Set[String] = Set[String]
    let dir = recover ref s.split(",") end
    for s' in dir.values() do
//      Debug.out(s')
      let r = Regex("(\\w)(\\d+)")?
      let matched = r(s')?
      let n = matched(2)?
      match matched(1)?
      | "U" =>
	 for i in Range(0,n.usize()?) do
	 y = y - 1
	 let here = x.string() + "," + y.string()
	 set.set(here)
         end
      | "D" =>
	 for i in Range(0,n.usize()?) do
	 y = y + 1
	 let here = x.string() + "," + y.string()
	 set.set(here)
         end
      | "L" =>
	 for i in Range(0,n.usize()?) do
	 x = x - 1
	 let here = x.string() + "," + y.string()
	 set.set(here)
         end
      | "R" =>
	 for i in Range(0,n.usize()?) do
	 x = x + 1
	 let here = x.string() + "," + y.string()
	 set.set(here)
         end
       end
    end
    for s' in set.values() do
      var value = counts.get_or_else(s',0)
      value = value + 1
      if value > 1 then
       let p = s'.split(",")
       let man = p(0)?.i32()?.abs() + p(1)?.i32()?.abs()
       Debug.out(man.string() + " " + s' + " " + value.string())
      end
      counts.insert(s', value)
    end
   end
     


  fun dispose() =>
//    for v in counts.values() do
//     Debug.out(v.string())
//    end
   None

