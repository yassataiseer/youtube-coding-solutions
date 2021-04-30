
public class solution2
{
	public static void main(String[] args) {
		int i = 0;
        while (i!=100){
            String Answer = "";
            i++;
            if (i%3==0){
                Answer+="Fizz";
            }
            if (i%5==0){
                Answer+="Buzz";
            }
            if(Answer.length()==0){
                Answer+=Integer.toString(i);
            }
            System.out.println(Answer);
        }
	}
}