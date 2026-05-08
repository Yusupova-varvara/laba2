import java.util.Scanner;

public class task {
    public static String compress(String s) {
        StringBuilder result = new StringBuilder();
        int count = 1;
        
        for (int i = 0; i < s.length(); i++) {
            if (i + 1 < s.length() && s.charAt(i) == s.charAt(i + 1)) {
                count++;
            } else {
                result.append(s.charAt(i)).append(count);
                count = 1;
            }
        }
        
        if (result.length() < s.length()) {
            return result.toString();
        }
        return s;
    }
    
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        System.out.print("Введите строку: ");
        String input = scanner.nextLine();
        System.out.println(compress(input));
        scanner.close();
    }
}
