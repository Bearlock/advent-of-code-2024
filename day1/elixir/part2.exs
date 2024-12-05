defmodule Stuff do
  # Turns the list of lists into a map with left list as keys and counts the
  # number of times items in the right list match a key on the left
  def to_count_map([left | [right | []]]) do
    default = Enum.reduce(left, %{}, fn e, acc ->
      Map.put(acc, String.to_integer(e), 0) 
    end)

    mp = Enum.reduce(right, %{}, fn e, acc ->
      intKey = String.to_integer(e)
      if Map.has_key?(default, intKey) do
        Map.put(acc, intKey, Map.get(acc, intKey, 0) + 1)
      else
        acc
      end
    end)

    mp
  end
end


File.stream!("../input.txt", :line)
|> Stream.map(&String.trim/1)
|> Stream.map(&String.split/1)
|> Enum.reduce([[], []], fn [left | [right | []]], [lacc, racc] -> [[left | lacc], [right | racc]] end)
|> Stuff.to_count_map() 
|> Enum.reduce(0, fn {k, v}, acc -> acc + (k * v) end)
|> (fn answer -> IO.puts("The answer is: #{answer}") end).()
