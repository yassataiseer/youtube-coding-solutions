// Import the HashMap class
import java.util.HashMap;
import java.util.Arrays;

public class Main {
  public static void main(String[] args) {

    // Create a HashMap object called people
    HashMap<String, Integer[]> people = new HashMap<String, Integer[]>();

    // Add keys and values (Name, Age)
    people.put("John", new Integer[] {1, 3, 2});
    people.put("Steve", new Integer[] {1, 3, 2});
    people.put("Angie", new Integer[] {1, 3, 2});

    System.out.println(people);
  }
}
