use "debug"
use "itertools"
use "collections"

actor Main
  new create(env: Env) =>
    try
    Debug("Day 16")
//    var str = "12345678"
//    var str = "80871224585914546619083218645595"
//    var str = "69317163492948606335995924319873"
      var str = "59773775883217736423759431065821647306353420453506194937477909478357279250527959717515453593953697526882172199401149893789300695782513381578519607082690498042922082853622468730359031443128253024761190886248093463388723595869794965934425375464188783095560858698959059899665146868388800302242666479679987279787144346712262803568907779621609528347260035619258134850854741360515089631116920328622677237196915348865412336221562817250057035898092525020837239100456855355496177944747496249192354750965666437121797601987523473707071258599440525572142300549600825381432815592726865051526418740875442413571535945830954724825314675166862626566783107780527347044"
    var start =
     Iter[U8](str.values())
      .map[ISize](M~atoi())
      .collect(Array[ISize])

    var next: Array[ISize] = start.clone()
    //Debug(start)
    //Debug(pattern2(0,start.size())?)
    let nphases:USize = 100
    //phase
    for j in Range(0,nphases) do
      for i in Range(0,start.size()) do
        next(i)? =
        Iter[ISize](start.values())
          .zip[ISize](pattern2(i,start.size())?.values())
          .map[ISize](M~mul())
          .fold[ISize](0, {(sum, x) => sum + x })
          .abs()
          .mod(10)
	  .isize()
       end
       start = next.clone()
       //Debug(next,"")
     end
     Debug(next.slice(0,8),"")
//     Debug("tests")
//     Debug(pattern2(0,8)?)
//     Debug(pattern2(1,8)?)
//     Debug(pattern2(2,8)?)
//     Debug(pattern2(7,8)?)          
    end

  fun ref pattern2(i:USize, len: USize): Array[ISize]? =>
    var i' = i + 1
    var l: Array[ISize] = Array[ISize].init(0,len + 1)
    let b = base(0,len)?
    var pos:USize = 0
    var c:USize = 0
    while c < (len + 1) do
      for j in Range(c,c+i') do
       try l(j)? = b(pos)? end
      end
    c = c + i'
    pos = pos + 1
    //Debug([c;len])
    end
    b
    l.>shift()?

  fun ref base(i: USize, len: USize): Array[ISize] ? =>
   var b: Array[ISize] = Array[ISize].init(0,len)
   var switch = true
   var c:USize = 0
   for i' in Range(0,len) do
//     Debug([i';c])
     if i'.mod(2) == 0 then
       if switch then b(c)? = 1 else b(c)? = -1 end
       switch = not switch
     end
     c = c + i + 1
   end
   b.>unshift(0)
    
    
  fun ref pattern(i: USize): Array[ISize] ?  =>
    match i
    | 0 => [1;0;-1;0;1;0;-1;0]
    | 1 => [0;1;1;0;0;-1;-1;0]
    | 2 => [0;0;1;1;1;0;0;0]
    | 3 => [0;0;0;1;1;1;1;0]
    | 4 => [0;0;0;0;1;1;1;1]
    | 5 => [0;0;0;0;0;1;1;1]
    | 6 => [0;0;0;0;0;0;1;1]
    | 7 => [0;0;0;0;0;0;0;1]
    else
     error
    end

primitive M
  fun mul(p: (ISize,ISize)): ISize => p._1 * p._2
  fun abs(i: ISize): ISize => if i < 0 then -i else i end
  fun atoi(a: U8): ISize => (a - '0').isize() 