use "debug"
use "collections"
use "itertools"

actor Main
  new create(env: Env) =>
    env.input(recover LineNotify(A) end, 512)

class A
 fun ref apply(s': String) =>
 try
   var s: String ref = recover s'.clone() end
   var lines: Array[String] = Array[String](100) 
   for i in Range(0,100) do
     lines.push(s.substring(0,150))
     s.cut_in_place(0,150)
   end

   var min = USize.max_value()
   var minprod: USize = 0 

   for freqs in Range(0,100) do
    let temp = lines(freqs)?.count("0")
    if temp < min then
      min = temp
      minprod = lines(freqs)?.count("1") * lines(freqs)?.count("2")
    end
   end
   Debug("Part 1: " + minprod.string())

  // String is an Array of 100 X 150 lines
  // [0....99]
  // [100.199]
  // [200.199]
  
  // Iter[U32](s.runes())
  //  .map[None]({(b) => Debug(String.from_utf32(b)) })
  //  .run()
  end

  fun dispose() => None 