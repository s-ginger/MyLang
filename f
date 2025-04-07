


class Greeter
  def initialize(name)
    @name = name.capitalize
  end

  def salute
    puts "Привет, #{@name}!"
  end
end

g = Greeter.new("мир")

g.salute

