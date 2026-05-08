import Foundation

func compress(_ s: String) -> String {
    var result = ""
    var count = 1
    let chars = Array(s)
    
    for i in 0..<chars.count {
        if i + 1 < chars.count && chars[i] == chars[i + 1] {
            count += 1
        } else {
            result += "\(chars[i])\(count)"
            count = 1
        }
    }
    
    return result.count < s.count ? result : s
}

print("Введите строку: ", terminator: "")
if let input = readLine() {
    print(compress(input))
}
