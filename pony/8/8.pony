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

   // Part 1
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
   
   // Part 2
   Debug.out("Part 2")
   var top_layer:Array[U8] iso = recover Array[U8] end
   for i in Range(0,150) do
     top_layer.push(top(i,lines)?)
   end
   let str = String.from_iso_array(consume top_layer)
   str.replace("0"," ")
   str.replace("1","#")   

   for i in Range(0,150,25) do
     Debug(str.substring(i.isize(),i.isize()+25))
   end
  end

  fun top(i: USize, a: Array[String]): U8 ? =>
    Iter[USize](Range(0,a.size()))
      .skip_while({(x)? => a(x)?(i)? == '2'} )
      .map[U8]({(x)? => a(x)?(i)?} )
      .nth(0)?   

  fun dispose() => None 