defmodule Pangram do
  def is_pangram(s) do
    unique_count = s
    |> String.downcase()
    |> String.replace(~r/[^a-z]/, "")
    |> String.graphemes()
    |> Enum.uniq()
    |> Enum.count()
    
    if unique_count == 26 do
      "Yes"
    else
      "No"
    end
  end
end

IO.puts("Введите строку: ")
input = IO.gets("") |> String.trim()
IO.puts(Pangram.is_pangram(input))

