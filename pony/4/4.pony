use "debug"
use "collections"
use "regex"
use "itertools"

actor Main
  new create(env: Env) =>
    env.out.print("2019 Day 4")
    var count:USize = 
    Iter[USize](Range(347312,805915))
     .map[String]({(i) => i.string()})
     .filter(R~rule1())
     .filter(R~rule2())
     .count()
    Debug.out("total " +  count.string())

primitive R
   fun rule1 (s: String ): Bool =>
     try 
       Iter[Match](MatchIterator(Regex("(\\d)\\1")?,s))
	.any({(m)? => not s.contains(m(1)?.mul(3)) })
     else
       false
     end

  fun rule2 (s: String): Bool =>
    not Iter[USize](Range(0,s.size()-1))
        .any({(c: USize)? => s(c)? > s(c+1)? })
     


// Part 1
// It is a six-digit number.
// The value is within the range given in your puzzle input.
// Two adjacent digits are the same (like 22 in 122345).

// Going from left to right, the digits never decrease; they only ever
// increase or stay the same (like 111123 or 135679)

// Part 2
// An Elf just remembered one more important detail: the two adjacent
// matching digits are not part of a larger group of matching digits
