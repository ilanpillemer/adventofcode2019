use "debug"

actor Main
  new create(env: Env) =>
    env.out.print("2019 in Pony.. Day 1")
    env.out.print("====================")
    env.input(recover LineNotify(A) end,512)

class A 
  var total:I32 = 0

  fun ref apply(s: String) =>
   try
     var n:I32 = s.i32()?
     while n >= 0 do   
       let fuel:I32 = (n / 3) - 2
       if fuel > 0 then
         total = total + fuel
       end
       n = fuel
     end
   end
  
  fun dispose() =>
    Debug.out("final: " + total.string())
    
