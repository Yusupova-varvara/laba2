fun isPangram(s: String): String {
    val letters = BooleanArray(26)
    var count = 0
    
    for (c in s.lowercase()) {
        if (c in 'a'..'z') {
            val index = c - 'a'
            if (!letters[index]) {
                letters[index] = true
                count++
            }
        }
    }
    
    return if (count == 26) "Yes" else "No"
}

fun main() {
    print("Введите строку: ")
    val input = readLine() ?: ""
    println(isPangram(input))
}
