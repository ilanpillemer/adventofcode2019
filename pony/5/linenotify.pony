use "buffered"

interface Applyable 
  fun ref apply(l: String)
  fun dispose() => None

class LineNotify is InputNotify
  let _rb: Reader
  let _p: Applyable

  new create(p: Applyable) =>
    _rb = Reader
    _p = p

  fun ref apply(data: Array[U8] iso) =>
    _rb.append(consume data)
    while true do
      try
        let l = _rb.line()?
	_p(consume l)
      else
        break
      end
    end

  fun ref dispose() =>
    try
      if _rb.size() > 0 then
        let rest: Array[U8] val = _rb.block(_rb.size())?
        _p(String.from_array(rest))
      end
    end
    _p.dispose()

