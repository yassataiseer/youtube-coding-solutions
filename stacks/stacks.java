
import java.util.Stack;

public class stacks {
    public static void main(String[] args) {
        // Creating a Stack
        Stack<String> stackOfdata = new Stack<>();

        // Adding new items to the Stack
        stackOfdata.push("Like");
        stackOfdata.push("Share");
        stackOfdata.push("Subscribe");
        System.out.println(stackOfdata);//Viewing current data add

        // Removing the top item from the Stack
        String dataAtTop = stackOfdata.pop();  
        System.out.println("The card on the top originally" + dataAtTop);
        System.out.println(stackOfdata);
        System.out.println();

        // Get the item at the top of the stack without removing it
        dataAtTop = stackOfdata.peek();
        System.out.println("Stack.peek() => " + dataAtTop);
        System.out.println("Current Stack => " + stackOfdata);

    }
}
