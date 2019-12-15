use "debug"
use "collections"
use "itertools"

actor Main
  new create (env: Env) =>
    Debug("Pony Day 14")
    env.input(recover LineNotify(A) end, 512)

class A

  let f : Map[String,(ISize, Array[(ISize,String)])] = Map[String,(ISize, Array[(ISize,String)])]
  let have : Map[String,ISize] = Map[String,ISize]

  fun ref apply(s: String) =>
    try
      let line = s.trim().split_by(" => ")
      let need =
        Iter[String](line(0)?.split_by(", ").values())
	  .map[(ISize,String)](this~p1())
	  .collect(Array[(ISize,String)])
      var nget: ISize = 0
      var get: String = recover String() end
      (nget, get) = p1(line(1)?.trim())?
      f.insert(get,(nget,need))
    end      


  fun p1(s: String): (ISize, String) ?  =>
   let p = s.trim().split(" ")
   (p(0)?.isize()?,p(1)?)

  fun ref ore(n': ISize, item: String): ISize ? =>
    var n = n'
    let h' = have.get_or_else(item,0)
    if h' >= n then
      have.insert(item,h'-n)
      return 0
    end
    if h' < n then
      n = n - h'
      have.insert(item,0)
    end

    if item == "ORE" then
      return n
    end
      var ans: ISize = 0
      var nget: ISize = 0
      var need: Array[(ISize, String)] = Array[(ISize, String)]
      (nget, need) = f(item)?

      (var amt: ISize, var extra: ISize) = getAmt(n,nget)
      let rem = (max(n,nget)) % nget
      let h = have.get_or_else(item,0)
      have.insert(item,h + extra)
      for v in need.values() do
        ans = ans + (ore(amt * v._1,v._2)?) 
      end
      ans

  fun getAmt(n: ISize, nget: ISize): (ISize,ISize) =>
    var amt = nget
    var inc = nget
    while true do
      if amt >= n then
        return (amt/nget, amt-n)
      end
      amt = amt + inc
    end
    (0,0)
    
 
  fun max(i: ISize, j: ISize): ISize => if i > j then i else j end
  fun abs(i: ISize): ISize => if i < 0 then -i else i end   
  
  fun ref dispose() =>
  try
    Debug("fuel needed")
    //PART 1 ===> Debug(ore(1,"FUEL")?)    
    //3_000_000 low
    //3_750_000 high
    //3_500_000 high
    //3_250_000 low 952400547108
    //3_375_000 low 989031282477
    //3_400_000 low 996357504862
    //3_450_000 high 1011009785481
    //3_425_000 high 1003683702949
    //3_423_750 high 1003317310866
    //3_415_000 high 1000753184586
    //3_407_500 low 998555385948
    //3_412_500 high 1000020551802
    //3_410_000 low 999288019814
    //3_411_000 low 999580957077
    //3_412_000 low 999874041409
    //3_412_250 low 999947395060
    //3_412_375 low 999983983687
    //3_412_415 low 999995748554
    //3_412_450 high 1000005937334
    //3_412_425 low 999998587534
    //3_412_437 high 1000002118971
    //3_412_432 high 1000000645130
    //3_412_430 high 1000000133741
    //3_412_428 low  999999469719
    //3_412_429 best 999999808842
    Debug(ore(3_412_429,"FUEL")?)
  end