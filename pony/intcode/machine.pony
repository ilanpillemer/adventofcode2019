use "debug"
use "itertools"

type Mem is Array[USize]

class val Address 
  let _address : USize 
  new val create(address: USize) =>
    _address = address

  fun apply() : USize =>
    _address
  
class val Val 
  let _value : USize 
  new val create(value: USize) =>
    _value = value

  fun apply() : USize =>
    _value


type Data is (Address | Val)


actor Machine
  let _mem : Mem ref
  var _pc : USize = 0

  new create(mem: Mem iso) =>
    _mem = consume mem

  be exec() =>
    try
      while true do
      let op = _mem(_pc)?
        match op
        | 1 => add(Address(_mem(_pc+1)?), Address(_mem(_pc+2)?), Address(_mem(_pc+3)?)); _pc = _pc + 4
	| 2 => mul(Address(_mem(_pc+1)?), Address(_mem(_pc+2)?), Address(_mem(_pc+3)?)); _pc = _pc + 4
	| 99 => halt(); break
        end
      end
    end

   be add(op1: Data val, op2: Data val, ad3: Address val) =>
     try
       Debug.out("add:")
       match (op1, op2)
        | (let ad1: Address, let ad2: Address) => _mem(ad3())? =  _mem(ad1())? + _mem(ad2())?
        | (let ad1: Address, let v2: Val) => _mem(ad3())? =  _mem(ad1())? + v2()
        | (let v1: Val, let ad2: Address) => _mem(ad3())? =  v1() + _mem(ad3())?
        | (let v1: Val, let v2: Val) => _mem(ad3())? =  v1() + v2()
       else
        error
       end
     end

   be mul(op1: Data val, op2: Data val, ad3: Address val) =>
     try
       Debug.out("mul:")
       match (op1, op2)
        | (let ad1: Address, let ad2: Address) => _mem(ad3())? =  _mem(ad1())? * _mem(ad2())?
        | (let ad1: Address, let v2: Val) => _mem(ad3())? =  _mem(ad1())? * v2()
        | (let v1: Val, let ad2: Address) => _mem(ad3())? =  v1() * _mem(ad3())?
        | (let v1: Val, let v2: Val) => _mem(ad3())? =  v1() * v2()
       else
        error
       end
     end

   be halt() =>
     // for v in _mem.values() do
     //   Debug.out(v.string())
     // end
     try
       Debug.out("Pos 0: " + _mem(0)?.string())
     end
     Debug.out("halted")
     None


