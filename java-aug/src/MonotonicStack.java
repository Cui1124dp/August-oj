import java.io.BufferedInputStream;
import java.util.Scanner;

/**
 * Created with IntelliJ IDEA.
 *
 * @Author: cdx1124dp
 * @Date: 2022/08/09/19:18
 * @Description:
 */
public class MonotonicStack {

    public static void main(String[] args) {
        int top = -1;
        Scanner scanner = new Scanner(new BufferedInputStream(System.in));
        int n = Integer.parseInt(scanner.nextLine());
        int[] list = new int[n];
        for (int i=0;i<n;i++){
            int num = scanner.nextInt();

            while(top!=-1&&list[top]>=num){
                top--;
            }
            if (top==-1){
                System.out.print(-1+" ");
            }else {
                System.out.print(list[top]+" ");
            }
            list[++top] = num;

        }
    }

}
