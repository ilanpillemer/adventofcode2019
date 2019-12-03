use "debug"
use "itertools"
use "regex"
use "collections"

// dependency of regex must be installed with stable or otherwise
// dependency of LineNotify is in the same folder as linenotify.pony

// sends to stdout and unordered list of all the step distances
// to each intersection.

// you can then get what you want with some some simple unix tools
// the correct answer for part 2 was deduced by
// `stable env ponyc -d && ./3 <in | sort -n | head -1`

// you could add a couple more lines to do this with Sort. I did not.

actor Main
  new create(env: Env) =>
    env.input(recover LineNotify(A) end,512)

class A

  var counts: Map[String,U32] = Map[String,U32] 
  
  fun ref apply(s: String) =>
  try
    var x:I32 = 0
    var y:I32 = 0
    var steps:U32 = 0

    var m:Map[String,U32] = Map[String,U32]
    let dir = recover ref s.split(",") end
    for s' in dir.values() do
      let r = Regex("(\\w)(\\d+)")?
      let matched = r(s')?
      let n = matched(2)?
      match matched(1)?
      | "U" =>
	 for i in Range(0,n.usize()?) do
	 steps = steps + 1
	 y = y - 1
	 let here = x.string() + "," + y.string()
	 m.insert_if_absent(here,steps)
         end
      | "D" =>
	 for i in Range(0,n.usize()?) do
	 steps = steps + 1
	 y = y + 1
	 let here = x.string() + "," + y.string()
	 m.insert_if_absent(here,steps)
         end
      | "L" =>
	 for i in Range(0,n.usize()?) do
	 steps = steps + 1	 
	 x = x - 1
	 let here = x.string() + "," + y.string()
	 m.insert_if_absent(here,steps)
         end
      | "R" =>
	 for i in Range(0,n.usize()?) do
	 steps = steps + 1	 
	 x = x + 1
	 let here = x.string() + "," + y.string()
	 m.insert_if_absent(here,steps) 
         end
       end
    end
    for s' in m.keys() do
      if counts.contains(s') then
        var v1 = m.get_or_else(s',0)
	var v2 = counts(s')?
	var value = v1 + v2
        if value > 0 then
          Debug.out(value.string())
        end
      end
      counts.insert(s',m.get_or_else(s',0))
    end
   end

  fun dispose() =>
   None

