def pos_and_kw(foo, bar: true)
  if bar
    puts foo
  end
end

pos_and_kw("x")
pos_and_kw("x", bar: false)

def all_kw(foo: "y", bar: true)
  if bar
    puts foo
  end
end

all_kw(bar: false)
all_kw(bar: false, foo: "z")

def defaults(foo = "x", bar = "y")
  "foo: #{foo}, bar: #{bar}"
end

defaults
defaults("z")
defaults("z", "a")

class Foo
  def initialize(foo = 10)
    @foo = foo
  end
end

Foo.new

def splat(a, *b, c: false)
  if c
    b[0]
  else
    a
  end
end

splat(9, 2, 3)
splat(9, 2, c: true)
splat(9)
splat(9, *[1, 2])
splat(9, 5, *[1, 2])
