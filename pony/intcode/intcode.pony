actor Main
  new create(env: Env) =>
    env.out.print("Intcode initialising...")
      //quicktest
      let mem: Array[USize] iso = recover iso [1;9;10;70;2;3;11;0;99;30;40;50] end
      let m : Machine = Machine(consume mem)
      m.exec()

 




















