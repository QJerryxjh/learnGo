function test1() {
  var b = 1
  test2()
}

function test2() {
  console.log(b)
}

test1()