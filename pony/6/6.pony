use "debug"
use "collections"
use "itertools"

actor Main
  new create(env: Env) =>
    env.input(recover LineNotify(A) end)

class A
  let _edges: Map[String, String] = recover Map[String,String] end
  let _nodes: Set[String] ref = Set[String]
  fun ref apply(s: String) =>
    try
      let e = s.split(")")
      _edges.insert(e(1)?,e(0)?)
      _nodes.set(e(1)?)
      _nodes.set(e(0)?)
    end
    
  fun dispose() =>
      //Part 1
      let total = Iter[String](_nodes.values())
        .map[U32](M~height(where h = 0, e = copy()))
        .fold[U32](0,{(acc, x) => acc + x})
      Debug.out("Total heights of all subtrees: " + total.string())

      //Part 2 distance from santa to you in graph
      let santa = M.path("SAN", copy(), Array[String])
      let you = M.path("YOU", copy(), Array[String])
      var common = 
        Iter[USize](Range(0,santa.size()))
         .take_while({(i)? => santa(i)? == you(i)?})
	 .count()
      let leg1 = santa.size() - common.usize()
      let leg2 = you.size() - common.usize()  
      let length = (leg1 + leg2)
      Debug.out("Distance from Santa to You: " + length.string())
  
  fun copy(): Map[String, String] val =>
    try
      let c = recover iso Map[String,String] end
      for k in _edges.keys() do
        c.insert(k,_edges(k)?)
      end
      consume c
    else
      recover val Map[String,String] end
  end
  

primitive M
  fun val height(s: String, e: Map[String,String] val,h: U32): U32 =>
    let parent = e.get_or_else(s,"")
    if parent == "" then return h end
    height(parent, e, h + 1)

  fun val path(s: String, e: Map[String,String] val,p: Array[String]): Array[String] =>
    let parent = e.get_or_else(s,"")
//    Debug.out(parent)
    if parent == "COM" then
      p.unshift(parent)
      return p
    end
    p.unshift(parent)
    path(parent,e,p)
    
    