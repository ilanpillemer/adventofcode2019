use "debug"
use "collections"
use "regex"

actor Main
  new create(env: Env) =>
    env.out.print("2019 Day 4")
    var count:U32 = 0
    for i in Range(347312,805915) do
      if rule1(i.string()) and rule2(i.string())then
	count = count + 1
      end
    end
    Debug.out("total " +  count.string())


  fun ref rule1(s: String): Bool =>
     try
       let r = Regex("(\\d)\\1")?
       r(s)? 
       true
     else
       false
     end

   fun ref rule2(s: String): Bool =>
   try
     for c in Range(0,s.size()-1) do
      if s(c)? > s(c+1)? then
        return false
      end
     end
     return true
   else
     Debug.out("wtf " + s)
     true 
   end


//  It is a six-digit number.
// The value is within the range given in your puzzle input.
//Two adjacent digits are the same (like 22 in 122345).

// Going from left to right, the digits never decrease; they only ever
// increase or stay the same (like 111123 or 135679)