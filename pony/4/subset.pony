use "debug"
use "collections"

class Subset
  let _len: U8
  var _subset: Array[U8]
  let _f: SubSetter
  new create(len: U8, f: SubSetter) =>
    _len = len
    _subset = Array[U8]
    _f = f

  fun ref search() =>
    _search(1)

  fun ref _search(k: U8)  =>
   try
     if (k == (_len + 1)) then
       //Debug.out("hello at " + _subset(0)?.string())
       _f(_subset)
       return
     end
      _subset.push(k)
      _search(k + 1)
      _subset.pop()?
      _search(k + 1)
   end

interface SubSetter
  fun ref apply(data: Array[U8])

